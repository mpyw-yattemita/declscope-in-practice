// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object-construct expression, such as
// { hat = "derby" }. Get one from an expression with
// Expression.AsObjectConsExpr.
//
// Its methods find an item by its key only when the key is an identifier or
// a quoted string literal. The two forms are equivalent: the key "hat" finds
// both hat = ... and "hat" = .... A key that is taken from a variable or from
// another expression, such as (hat), is never found.
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
// there is none.
//
// A quoted key is compared by its value after its escape sequences are
// decoded, so ItemFor("${hat}") finds the key "$${hat}".
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if name, ok := item.KeyExpr().staticName(); ok && name == key {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor finds for the
// given key, or nil if there is none.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	item := o.ItemFor(key)
	if item == nil {
		return nil
	}
	return item.ValueExpr()
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, leaving its key as written, or adds a new item at the end of
// the object.
//
// A new key is written as an identifier when it is a valid identifier other
// than "for", and otherwise as a quoted string literal that evaluates to the
// given key.
//
// SetItem panics if the expression is already part of a syntax tree, such as
// one returned by Attribute.Expr or ObjectConsValue.Expr. Two items sharing
// one expression would change together.
//
// The return values are the key and the value of the item that was either
// modified in place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	item := o.ItemFor(key)
	if item != nil {
		item.setValueExpr(expr)
	} else {
		item = newObjectConsItem()
		item.init(key, expr)
		o.appendItem(item)
	}
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

// SetItemValue is like SetItem, but sets the value to the given literal, as
// NewExpressionLiteral does. To set a variable reference or other traversal,
// use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but sets the value to the given
// traversal, as NewExpressionAbsTraversal does. The traversal must be
// absolute. To set a literal value, use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// variableNodes returns the traversal nodes in the keys and values of the
// items, in source order. A key that is a literal name contributes none.
//
func (o *ObjectConsExpr) variableNodes() []*node {
	var ret []*node
	for _, n := range o.items.List() {
		item := n.content.(*ObjectConsItem)
		key := item.key.content.(*ObjectConsKeyExpr)
		if !key.literalName {
			ret = append(ret, key.wrapped.content.(*Expression).variableNodes()...)
		}
		value := item.value.content.(*ObjectConsValue)
		ret = append(ret, value.expr.content.(*Expression).variableNodes()...)
	}
	return ret
}

// appendItem adds the item after the last one, laid out as the object
// already is: on the same line after a comma when the object is written on
// one line, and otherwise on a line of its own before the closing brace.
func (o *ObjectConsExpr) appendItem(item *ObjectConsItem) {
	// The parser leaves everything after the last item, up to and including
	// the closing brace, in one run of tokens at the end of the children.
	var closing *node
	o.walkChildNodes(func(n *node) { closing = n })
	closingTokens, ok := closing.content.(Tokens)
	if !ok || len(closingTokens) == 0 || closingTokens[len(closingTokens)-1].Type != hclsyntax.TokenCBrace {
		panic(fmt.Sprintf("object-construct expression does not end with a closing brace, but with %T", closing.content))
	}
	// Each use gets its own token, since formatting changes tokens in place.
	newline := func() *Token {
		return &Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte{'\n'},
		}
	}

	multiline := o.isMultiline()
	kept, rest := o.splitClosing(closingTokens)
	// Copy kept, since appending to it would overwrite the start of rest.
	separator := append(Tokens(nil), kept...)
	var afterItem Tokens
	switch {
	case multiline:
		if len(kept) == 0 || !o.tokenEndsLine(kept[len(kept)-1]) {
			separator = append(separator, newline())
		}
		afterItem = Tokens{newline()}
	case len(o.items) > 0 && len(kept) == 0:
		separator = append(separator, &Token{
			Type:  hclsyntax.TokenComma,
			Bytes: []byte{','},
		})
	}
	afterItem = append(afterItem, rest...)

	closing.Detach()
	o.children.AppendUnstructuredTokens(separator)
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(afterItem)
}

// isMultiline reports whether the object is written on several lines. Line
// breaks inside the keys and values of its items do not count.
func (o *ObjectConsExpr) isMultiline() bool {
	multiline := false
	checkTokens := func(n *node) {
		if tokens, ok := n.content.(Tokens); ok {
			for _, tok := range tokens {
				multiline = multiline || o.tokenEndsLine(tok)
			}
		}
	}
	o.walkChildNodes(func(n *node) {
		checkTokens(n)
		if item, ok := n.content.(*ObjectConsItem); ok {
			item.walkChildNodes(checkTokens)
		}
	})
	return multiline
}

// splitClosing splits the tokens after the last item into those that stay
// with that item and the rest, which end with the closing brace. The item
// keeps a comma that follows it, and the comments up to the end of its line
// when the closing brace is not on that line.
func (o *ObjectConsExpr) splitClosing(tokens Tokens) (kept, rest Tokens) {
	afterComma := 0
	if tokens[0].Type == hclsyntax.TokenComma {
		afterComma = 1
	}
	for i := afterComma; i < len(tokens); i++ {
		if o.tokenEndsLine(tokens[i]) {
			return tokens[:i+1], tokens[i+1:]
		}
		if tokens[i].Type != hclsyntax.TokenComment {
			break
		}
	}
	return tokens[:afterComma], tokens[afterComma:]
}

// tokenEndsLine reports whether the token is a newline, or a line comment,
// which includes its newline.
func (o *ObjectConsExpr) tokenEndsLine(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		return len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n'
	}
	return false
}

// ObjectConsItem is one item of an object-construct expression: a key and a
// value, as in hat = "derby".
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

// KeyExpr returns the key of the item.
func (item *ObjectConsItem) KeyExpr() *ObjectConsKeyExpr {
	return item.key.content.(*ObjectConsKeyExpr)
}

// ValueExpr returns the value of the item.
func (item *ObjectConsItem) ValueExpr() *ObjectConsValue {
	return item.value.content.(*ObjectConsValue)
}

// init builds a new item from the key and the value expression, which must
// not be attached yet. The key is written as an identifier when it is a
// valid one, and otherwise as a quoted string literal.
func (item *ObjectConsItem) init(key string, expr *Expression) {
	// A naked "for" as the first key would start a for expression.
	literalName := hclsyntax.ValidIdentifier(key) && key != "for"
	var keyTokens Tokens
	if literalName {
		keyTokens = TokensForIdentifier(key)
	} else {
		keyTokens = TokensForValue(cty.StringVal(key))
	}
	item.key = item.children.Append(newObjectConsKeyExpr(newNode(NewExpressionRaw(keyTokens)), literalName))
	item.children.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte{'='},
		},
	})

	item.value = item.children.Append(newObjectConsValue())
	item.setValueExpr(expr)
}

// setValueExpr makes the given expression, which must not be attached yet,
// the expression of the item's value. The expression it replaces is
// detached.
func (item *ObjectConsItem) setValueExpr(expr *Expression) {
	value := item.ValueExpr()
	if value.expr != nil {
		value.Expr().setParent(nil)
	}
	// The value has no children other than its expression.
	value.children.Clear()
	value.expr = value.children.Append(expr)
	expr.setParent(item.value)
}

// ObjectConsKeyExpr is the key of an item in an object-construct
// expression. Get one from ObjectConsItem.KeyExpr.
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node

	// literalName is true when the key is a naked name, such as "k" in
	// { k = v }. As in hclsyntax, such a key is a literal string, not a
	// reference, so it has no variables.
	literalName bool
}

func newObjectConsKeyExpr(wrapped *node, literalName bool) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:      newInTree(),
		wrapped:     wrapped,
		literalName: literalName,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// staticName returns the string that the key stands for, and true, when the
// key is an identifier or a quoted string literal. Otherwise it returns false.
//
// It reads the key from its tokens as hclsyntax does, so that a quoted key
// is compared by its value after its escape sequences are decoded.
func (k *ObjectConsKeyExpr) staticName() (string, bool) {
	tokens := k.BuildTokens(nil)
	if len(tokens) == 0 {
		return "", false
	}
	native, diags := hclsyntax.ParseExpression(tokens.Bytes(), "", hcl.InitialPos)
	if diags.HasErrors() {
		return "", false
	}

	// As in hclsyntax, a naked identifier is a literal name. A parenthesized
	// one is a ParenthesesExpr, which this does not unwrap.
	if name := hcl.ExprAsKeyword(native); name != "" {
		return name, true
	}

	template, ok := native.(*hclsyntax.TemplateExpr)
	if !ok || tokens[0].Type != hclsyntax.TokenOQuote {
		return "", false
	}
	for _, part := range template.Parts {
		if _, ok := part.(*hclsyntax.LiteralValueExpr); !ok {
			return "", false
		}
	}
	val, diags := template.Value(nil)
	if diags.HasErrors() || !val.IsKnown() || val.IsNull() || val.Type() != cty.String {
		return "", false
	}
	return val.AsString(), true
}

// ObjectConsValue is the value of an item in an object-construct
// expression. Get one from ObjectConsItem.ValueExpr.
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
