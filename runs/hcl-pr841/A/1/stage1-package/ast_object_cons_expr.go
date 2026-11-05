// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"bytes"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object-construct expression, such as
// { name = "value" }. Use Expression.AsObjectConsExpr to get one.
//
// Its items can be looked up by key and changed with the SetItem methods,
// in the same way as the attributes of a Body.
type ObjectConsExpr struct {
	inTree

	items nodeSet
}

func newObjectConsExpr() *ObjectConsExpr {
	return &ObjectConsExpr{
		inTree: newInTree(),
		items:  newNodeSet(),
	}
}

// Items returns the items of the object, in source order.
func (o *ObjectConsExpr) Items() []*ObjectConsItem {
	nodes := o.items.List()
	ret := make([]*ObjectConsItem, len(nodes))
	for i, n := range nodes {
		ret[i] = n.content.(*ObjectConsItem)
	}
	return ret
}

// ItemFor returns the first item whose key is the given string, or nil if
// there is no such item.
//
// Only a key written as an identifier, such as hat, or as a quoted string
// literal, such as "hat", can match. A quoted key is compared by its value
// after escape sequences are decoded. A key that is computed from another
// expression, such as (hat), hat.band or "${hat}", never matches.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if name, ok := item.KeyExpr().literalName(); ok && name == key {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor finds for the
// given key, or nil if there is no such item.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	item := o.ItemFor(key)
	if item == nil {
		return nil
	}
	return item.ValueExpr()
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, or adds a new item with that key to the end of the object.
//
// The key of an existing item is left as it is written. A new key is written
// as an identifier if it is a valid identifier, or as a quoted string literal
// otherwise.
//
// The given expression must not already be part of a syntax tree, or this
// method panics.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	if item := o.ItemFor(key); item != nil {
		value := item.ValueExpr()
		value.expr = value.expr.ReplaceWith(expr)
		return item.KeyExpr(), value
	}

	item := newObjectConsItem()
	item.init(key, expr)
	o.appendItem(item)
	return item.KeyExpr(), item.ValueExpr()
}

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the
// value expression.
//
// The same caveats apply to this function as for NewExpressionRaw on which
// it is based. If possible, prefer to use SetItemValue or SetItemTraversal.
func (o *ObjectConsExpr) SetItemRaw(key string, tokens Tokens) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionRaw(tokens))
}

// SetItemValue is like SetItem, but sets the value to a literal that
// represents the given value. It is based on NewExpressionLiteral.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but sets the value to the given
// traversal, which must be absolute. It is based on NewExpressionAbsTraversal.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// appendItem adds the given new item to the end of the object, laid out
// as described in appendItemTokens.
func (o *ObjectConsExpr) appendItem(item *ObjectConsItem) {
	multiline := o.isMultiline()

	// Collect the tokens after the last item, or after the opening brace
	// if there are no items, so that they can be split around the new item.
	var rest Tokens
	var start *node
	if items := o.items.List(); len(items) > 0 {
		start = items[len(items)-1].after
	} else {
		start = o.children.first
	}
	for n := start; n != nil; {
		next := n.after
		rest = append(rest, n.content.(Tokens)...)
		n.Detach()
		n = next
	}
	if len(o.items) == 0 {
		open := 0
		for open < len(rest) && rest[open].Type != hclsyntax.TokenOBrace {
			open++
		}
		o.children.AppendUnstructuredTokens(rest[:open+1])
		rest = rest[open+1:]
	}

	before, after := splitObjectTail(rest, multiline, len(o.items) > 0)
	o.children.AppendUnstructuredTokens(before)
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(after)
}

// splitObjectTail splits the tokens that follow the last item of an object,
// up to and including the closing brace, into the tokens to put before and
// after a new item that is added at the end.
//
// A comma or a line comment that follows the last item on its line stays
// before the new item. Everything else, including comments on the following
// lines and an inline comment directly before the brace, stays with the
// closing brace after the new item. Separators are added so that the new
// item goes on the same line in a single-line object, or on its own line in
// a multi-line object.
func splitObjectTail(tail Tokens, multiline, hasItems bool) (before, after Tokens) {
	split := 0
	for i, tok := range tail {
		if multiline && isLineEnd(tok) {
			split = i + 1
			break
		}
		if tok.Type == hclsyntax.TokenComma {
			split = i + 1
			if !multiline {
				break
			}
		}
	}
	before = append(Tokens{}, tail[:split]...)
	after = tail[split:]

	switch {
	case multiline:
		if len(before) == 0 || !isLineEnd(before[len(before)-1]) {
			before = append(before, newNewlineToken())
		}
		after = append(Tokens{newNewlineToken()}, after...)
	case hasItems && split == 0:
		before = append(before, &Token{
			Type:  hclsyntax.TokenComma,
			Bytes: []byte{','},
		})
	}
	return before, after
}

// isMultiline reports whether the object's own tokens contain a line break.
// Line breaks inside the keys and values of its items are not counted.
func (o *ObjectConsExpr) isMultiline() bool {
	for n := o.children.first; n != nil; n = n.after {
		switch c := n.content.(type) {
		case Tokens:
			if hasLineEnd(c) {
				return true
			}
		case *ObjectConsItem:
			for in := c.children.first; in != nil; in = in.after {
				if tokens, ok := in.content.(Tokens); ok && hasLineEnd(tokens) {
					return true
				}
			}
		}
	}
	return false
}

func hasLineEnd(tokens Tokens) bool {
	for _, tok := range tokens {
		if isLineEnd(tok) {
			return true
		}
	}
	return false
}

// isLineEnd reports whether the token ends a line. A line comment includes
// the newline that ends it.
func isLineEnd(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		return bytes.HasSuffix(tok.Bytes, []byte{'\n'})
	}
	return false
}

func newNewlineToken() *Token {
	return &Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte{'\n'},
	}
}

// variables returns the traversals in the keys and values of the items,
// in source order.
func (o *ObjectConsExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, item := range o.Items() {
		ret = append(ret, item.KeyExpr().variables()...)
		ret = append(ret, item.ValueExpr().Expr().Variables()...)
	}
	return ret
}

// ObjectConsItem is a single key = value item of an ObjectConsExpr.
type ObjectConsItem struct {
	inTree
	key   *node
	value *node
}

func newObjectConsItem() *ObjectConsItem {
	return &ObjectConsItem{
		inTree: newInTree(),
	}
}

func (item *ObjectConsItem) init(key string, expr *Expression) {
	item.key = item.children.Append(newObjectConsKeyExprFor(key))
	item.children.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte{'='},
		},
	})
	value := newObjectConsValue()
	value.expr = value.children.Append(expr)
	item.value = item.children.Append(value)
}

// KeyExpr returns the key of the item.
func (item *ObjectConsItem) KeyExpr() *ObjectConsKeyExpr {
	return item.key.content.(*ObjectConsKeyExpr)
}

// ValueExpr returns the value of the item.
func (item *ObjectConsItem) ValueExpr() *ObjectConsValue {
	return item.value.content.(*ObjectConsValue)
}

// ObjectConsKeyExpr is the key of an ObjectConsItem. It is either a literal
// name, written as an identifier or a quoted string, or an expression whose
// result is the key.
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node
}

func newObjectConsKeyExpr(wrapped *node) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// newObjectConsKeyExprFor returns a key that evaluates to the given string.
// It is an identifier if the string is a valid identifier, or a quoted string
// literal otherwise.
func newObjectConsKeyExprFor(key string) *ObjectConsKeyExpr {
	expr := newExpression()
	if hclsyntax.ValidIdentifier(key) {
		expr.children.AppendUnstructuredTokens(TokensForIdentifier(key))
	} else {
		expr.wrapped = expr.children.Append(newQuoted(TokensForValue(cty.StringVal(key))))
	}
	return newObjectConsKeyExpr(newNode(expr))
}

// literalName returns the name that the key stands for if it is an
// identifier or a quoted string literal. Otherwise, ok is false.
//
// The rules match hclsyntax.ObjectConsKeyExpr: the keywords null, true and
// false are names too, and a key in parentheses is always an expression.
func (k *ObjectConsKeyExpr) literalName() (name string, ok bool) {
	src := k.BuildTokens(nil).Bytes()
	expr, diags := hclsyntax.ParseExpression(src, "", hcl.InitialPos)
	if diags.HasErrors() {
		return "", false
	}

	switch expr := expr.(type) {
	case *hclsyntax.ScopeTraversalExpr, *hclsyntax.LiteralValueExpr:
		name := hcl.ExprAsKeyword(expr)
		return name, name != ""
	case *hclsyntax.TemplateExpr:
		if !expr.IsStringLiteral() {
			return "", false
		}
		val, diags := expr.Value(nil)
		if diags.HasErrors() {
			return "", false
		}
		return val.AsString(), true
	default:
		return "", false
	}
}

func (k *ObjectConsKeyExpr) variables() []*Traversal {
	if k.wrapped == nil {
		return nil
	}
	return k.wrapped.content.(*Expression).Variables()
}

// ObjectConsValue is the value of an ObjectConsItem.
type ObjectConsValue struct {
	inTree
	expr *node
}

func newObjectConsValue() *ObjectConsValue {
	return &ObjectConsValue{
		inTree: newInTree(),
	}
}

// Expr returns the expression of the value.
func (o *ObjectConsValue) Expr() *Expression {
	return o.expr.content.(*Expression)
}
