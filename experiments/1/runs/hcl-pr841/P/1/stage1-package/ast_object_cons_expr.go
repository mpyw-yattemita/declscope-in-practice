// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object-construct expression, such as
// { hat = "derby" }. Use Expression.AsObjectConsExpr to get one.
//
// Its methods look up and rewrite the items of the object in place, in the
// same way as the methods of Body do for attributes.
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
	items := make([]*ObjectConsItem, len(nodes))
	for i, n := range nodes {
		items[i] = n.content.(*ObjectConsItem)
	}
	return items
}

// ItemFor returns the item whose key is the given string, or nil if there
// is no such item.
//
// Only a key written as an identifier or as a quoted string literal can be
// found, and the two forms are equivalent: ItemFor("hat") finds both
// hat = ... and "hat" = .... A quoted key is compared by its value, after
// its escape sequences are decoded. A key that is taken from a variable or
// from another expression, such as (hat), hat.band or "${hat}", is never
// found.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if name, ok := item.KeyExpr().name(); ok && name == key {
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
// given key, or adds a new item to the end of the object, using the given
// expression as the value.
//
// The key of an existing item is left as it is written. A new key is written
// as an identifier if it is a valid identifier, and as a quoted string
// literal otherwise.
//
// The expression must not be part of a syntax tree yet, or this method will
// panic. That includes an expression that came from the parser and one that
// was given to an earlier call of SetItem.
//
// The return values are the key and the value of the item that was either
// modified in place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	item := o.ItemFor(key)
	if item != nil {
		value := item.ValueExpr()
		value.expr = value.expr.ReplaceWith(expr)
	} else {
		item = o.appendItem(key, expr)
	}
	return item.KeyExpr(), item.ValueExpr()
}

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the value.
//
// The same caveats apply to this method as for NewExpressionRaw on which it
// is based. If possible, prefer to use SetItemValue or SetItemTraversal.
func (o *ObjectConsExpr) SetItemRaw(key string, tokens Tokens) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionRaw(tokens))
}

// SetItemValue is like SetItem, but writes the given value as a literal, in
// the same way as NewExpressionLiteral.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but writes the given traversal, which
// must be absolute, in the same way as NewExpressionAbsTraversal.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// appendItem adds a new item to the end of the object, laid out in the
// style of the object.
//
// In an object written on a single line, the new item goes on the same line,
// after a comma. In an object written on several lines, it goes on its own
// line after the last item, before any comments that precede the closing
// brace.
func (o *ObjectConsExpr) appendItem(key string, expr *Expression) *ObjectConsItem {
	multiline := o.isMultiline()
	hasItems := len(o.items) > 0
	lineEnd, rest := o.splitTail(o.takeTail())

	newline := &Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}
	comma := &Token{Type: hclsyntax.TokenComma, Bytes: []byte{','}}

	var leading Tokens
	switch {
	case multiline:
		leading = lineEnd
		if _, _, ok := o.splitAtLineEnd(lineEnd); !ok {
			leading = append(leading, newline)
		}
		rest = append(Tokens{newline}, rest...)
	case len(lineEnd) > 0:
		// A single-line object that ends with a trailing comma already has
		// the separator that the new item needs.
		leading = lineEnd
	case hasItems:
		leading = Tokens{comma}
	}

	item := newObjectConsItem()
	item.init(leading, key, expr)
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(rest)
	return item
}

// isMultiline reports whether the object is written on several lines. Only
// the line breaks of the object itself count, not those inside the keys and
// values of its items.
func (o *ObjectConsExpr) isMultiline() bool {
	for n := o.children.first; n != nil; n = n.after {
		var toks Tokens
		switch c := n.content.(type) {
		case Tokens:
			toks = c
		case *ObjectConsItem:
			toks = c.separatorTokens()
		}
		if _, _, ok := o.splitAtLineEnd(toks); ok {
			return true
		}
	}
	return false
}

// takeTail removes the tokens that follow the last item, or that follow the
// opening brace if there are no items, and returns them. They end with the
// closing brace.
func (o *ObjectConsExpr) takeTail() Tokens {
	if items := o.items.List(); len(items) > 0 {
		var tail Tokens
		for n := items[len(items)-1].after; n != nil; {
			next := n.after
			tail = n.BuildTokens(tail)
			n.Detach()
			n = next
		}
		return tail
	}

	all := o.BuildTokens(nil)
	o.children.Clear()
	for i, tok := range all {
		if tok.Type == hclsyntax.TokenOBrace {
			o.children.AppendUnstructuredTokens(append(Tokens(nil), all[:i+1]...))
			return append(Tokens(nil), all[i+1:]...)
		}
	}
	// An object always has an opening brace, so this is unreachable.
	o.children.AppendUnstructuredTokens(all)
	return nil
}

// splitTail splits the tail of the object into the tokens that stay with
// the last item, and the tokens that stay with the closing brace.
//
// The tokens that stay with the last item are those up to the end of its
// line, such as a comma and a line comment. If the closing brace is on the
// same line, only a comma directly after the last item stays with it.
func (o *ObjectConsExpr) splitTail(tail Tokens) (lineEnd, rest Tokens) {
	if lineEnd, rest, ok := o.splitAtLineEnd(tail); ok {
		return lineEnd, rest
	}
	if len(tail) > 0 && tail[0].Type == hclsyntax.TokenComma {
		return append(Tokens(nil), tail[:1]...), append(Tokens(nil), tail[1:]...)
	}
	return nil, tail
}

// splitAtLineEnd splits the given tokens after the first token that ends a
// line, which is a newline or a line comment. It returns false if no token
// ends a line.
//
// The halves are copies, so appending to the first one does not overwrite
// the second.
func (o *ObjectConsExpr) splitAtLineEnd(toks Tokens) (Tokens, Tokens, bool) {
	for i, tok := range toks {
		switch tok.Type {
		case hclsyntax.TokenNewline, hclsyntax.TokenComment:
			// Line comments (# and //) include their terminating newline.
			if len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n' {
				return append(Tokens(nil), toks[:i+1]...), append(Tokens(nil), toks[i+1:]...), true
			}
		}
	}
	return nil, toks, false
}

// ObjectConsItem is one item of an ObjectConsExpr, made of a key and a value.
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

// init builds a new item with the given key and value, preceded by the given
// tokens, which separate it from the item before it.
func (item *ObjectConsItem) init(leading Tokens, key string, expr *Expression) {
	var keyTokens Tokens
	if hclsyntax.ValidIdentifier(key) {
		keyTokens = TokensForIdentifier(key)
	} else {
		keyTokens = TokensForValue(cty.StringVal(key))
	}
	keyTokens[0].SpacesBefore = 1

	item.children.AppendUnstructuredTokens(leading)
	item.key = item.children.Append(newObjectConsKeyExpr(newNode(NewExpressionRaw(keyTokens))))
	item.children.AppendUnstructuredTokens(Tokens{
		{
			Type:         hclsyntax.TokenEqual,
			Bytes:        []byte{'='},
			SpacesBefore: 1,
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

// separatorTokens returns the tokens of the item that are neither in its key
// nor in its value, such as the separator before the item and the equals
// sign.
func (item *ObjectConsItem) separatorTokens() Tokens {
	var toks Tokens
	for n := item.children.first; n != nil; n = n.after {
		if c, ok := n.content.(Tokens); ok {
			toks = append(toks, c...)
		}
	}
	return toks
}

// ObjectConsKeyExpr is the key of an ObjectConsItem. It is an identifier, or
// an expression whose value is the key.
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

// name returns the string that the key stands for, if the key is an
// identifier or a quoted string literal. It returns false for any other key.
func (k *ObjectConsKeyExpr) name() (string, bool) {
	expr, diags := hclsyntax.ParseExpression(k.BuildTokens(nil).Bytes(), "", hcl.InitialPos)
	if diags.HasErrors() {
		return "", false
	}

	if name := hcl.ExprAsKeyword(expr); name != "" {
		return name, true
	}

	tmpl, ok := expr.(*hclsyntax.TemplateExpr)
	if !ok {
		return "", false
	}
	for _, part := range tmpl.Parts {
		if _, ok := part.(*hclsyntax.LiteralValueExpr); !ok {
			return "", false
		}
	}
	val, diags := tmpl.Value(nil)
	if diags.HasErrors() || !val.IsKnown() || val.IsNull() || !val.Type().Equals(cty.String) {
		return "", false
	}
	return val.AsString(), true
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
