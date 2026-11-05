// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object-construct expression, such as
// { hat = "derby" }. Use Expression.AsObjectConsExpr to get one from a parsed
// expression.
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
	ret := make([]*ObjectConsItem, 0, len(nodes))
	for _, n := range nodes {
		ret = append(ret, n.content.(*ObjectConsItem))
	}
	return ret
}

// ItemFor returns the first item whose key is the given string, or nil if
// there is no such item.
//
// Only a key written as an identifier or as a quoted string literal can
// match, and the two forms are equivalent: "hat" finds both hat = ... and
// "hat" = .... A quoted key is compared by its value, after its escape
// sequences are decoded. A key taken from a variable or another expression,
// such as (hat), hat.band or "${hat}", never matches.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if name, ok := item.KeyExpr().name(); ok && name == key {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor finds for the given
// key, or nil if there is no such item.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	item := o.ItemFor(key)
	if item == nil {
		return nil
	}
	return item.ValueExpr()
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, or adds a new item with the given key to the end of the object.
// The key of an existing item is left as written. A new key is written as an
// identifier if it is a valid one, and as a quoted string literal otherwise,
// escaped so that it evaluates to the given string. The key "for" is always
// quoted, since { for = ... } would read as a for expression.
//
// The given expression must not already be part of a syntax tree, or SetItem
// panics. Reusing an attached expression would make two items share it.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
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

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the
// value.
//
// The same caveats apply to this function as for NewExpressionRaw on which
// it is based. If possible, prefer to use SetItemValue or SetItemTraversal.
func (o *ObjectConsExpr) SetItemRaw(key string, tokens Tokens) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionRaw(tokens))
}

// SetItemValue is like SetItem, but writes the given value as a literal, as
// NewExpressionLiteral does. To set a variable reference or other traversal,
// use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but writes the given traversal, which
// must be absolute, as NewExpressionAbsTraversal does. To set a literal value,
// use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// variables returns the absolute traversals in the keys and values of the
// object's items, in source order. A naked identifier key is a literal name,
// not a variable, so it contributes nothing.
//
func (o *ObjectConsExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, n := range o.items.List() {
		item, ok := n.content.(*ObjectConsItem)
		if !ok {
			continue
		}
		if key := item.key.content.(*ObjectConsKeyExpr); !key.literal {
			ret = append(ret, key.wrapped.content.(*Expression).Variables()...)
		}
		ret = append(ret, item.value.content.(*ObjectConsValue).Expr().Variables()...)
	}
	return ret
}

// appendItem adds a new item with the given key and value after the last
// item of the object, and returns it.
//
// In an object written on one line, the item goes on the same line, after a
// comma. Otherwise it goes on its own line. A comma or a line comment that
// follows the last item on its line stays with that item. Whatever follows
// on later lines, and an inline comment directly before the closing brace,
// stays with the brace.
func (o *ObjectConsExpr) appendItem(key string, expr *Expression) *ObjectConsItem {
	itemNodes := o.items.List()
	stream := o.BuildTokens(nil)

	// Line breaks inside the keys and values of the items do not make the
	// object multi-line.
	inner := make(map[*Token]bool)
	for _, n := range itemNodes {
		item := n.content.(*ObjectConsItem)
		for _, tok := range item.KeyExpr().BuildTokens(nil) {
			inner[tok] = true
		}
		for _, tok := range item.ValueExpr().BuildTokens(nil) {
			inner[tok] = true
		}
	}
	multiline := false
	for _, tok := range stream {
		if !inner[tok] && o.isLineEnd(tok) {
			multiline = true
			break
		}
	}

	// Split the tokens of the object into the gaps before each item and the
	// tail after the last item, or after the opening brace if there are no
	// items. The tail ends with the closing brace.
	gaps := make([]Tokens, 0, len(itemNodes)+1)
	pos := 0
	for _, n := range itemNodes {
		first := n.BuildTokens(nil)[0]
		start := pos
		for stream[start] != first {
			start++
		}
		gaps = append(gaps, stream[pos:start:start])
		pos = start + len(n.BuildTokens(nil))
	}
	tail := stream[pos:]
	if len(itemNodes) == 0 {
		for i, tok := range tail {
			if tok.Type == hclsyntax.TokenOBrace {
				gaps = append(gaps, tail[:i+1:i+1])
				tail = tail[i+1:]
				break
			}
		}
	}

	// Rebuild the children from the gaps and the items, so that the tail
	// can be rearranged around the new item.
	for _, n := range itemNodes {
		n.Detach()
	}
	o.children.Clear()
	for i, gap := range gaps {
		o.children.AppendUnstructuredTokens(gap)
		if i < len(itemNodes) {
			o.children.AppendNode(itemNodes[i])
		}
	}

	// The tokens before split stay with the last item, and the rest with
	// the closing brace.
	split := len(tail) - 1
	for split > 0 && tail[split-1].Type == hclsyntax.TokenComment && !o.isLineEnd(tail[split-1]) {
		split--
	}
	if multiline {
		for i, tok := range tail[:split] {
			if o.isLineEnd(tok) {
				split = i + 1
				break
			}
		}
	}
	before, after := tail[:split], tail[split:]

	item := newObjectConsItem()
	switch {
	case multiline && (len(before) == 0 || !o.isLineEnd(before[len(before)-1])):
		// The closing brace shared a line with the last item.
		item.children.AppendUnstructuredTokens(Tokens{o.newlineToken()})
	case !multiline && len(itemNodes) > 0 && !o.hasComma(before):
		item.children.AppendUnstructuredTokens(Tokens{{
			Type:  hclsyntax.TokenComma,
			Bytes: []byte{','},
		}})
	}

	// An object that starts with the identifier for is read as a for
	// expression, so that key is quoted too.
	var keyExpr *Expression
	literal := hclsyntax.ValidIdentifier(key) && key != "for"
	if literal {
		keyExpr = NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: key}})
	} else {
		keyExpr = NewExpressionRaw(TokensForValue(cty.StringVal(key)))
	}
	if !multiline {
		keyExpr.BuildTokens(nil)[0].SpacesBefore = 1
	}
	item.key = item.children.Append(newObjectConsKeyExpr(newNode(keyExpr), literal))
	item.children.AppendUnstructuredTokens(Tokens{{
		Type:         hclsyntax.TokenEqual,
		Bytes:        []byte{'='},
		SpacesBefore: 1,
	}})
	value := newObjectConsValue()
	value.expr = value.children.Append(expr)
	item.value = item.children.Append(value)

	if multiline {
		after = append(Tokens{o.newlineToken()}, after...)
	} else {
		after[0].SpacesBefore = 1
	}

	o.children.AppendUnstructuredTokens(before)
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(after)
	return item
}

// isLineEnd reports whether tok ends a line. A line comment ends the line
// it is on, since it includes its newline.
func (o *ObjectConsExpr) isLineEnd(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		return len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n'
	}
	return false
}

func (o *ObjectConsExpr) hasComma(tokens Tokens) bool {
	for _, tok := range tokens {
		if tok.Type == hclsyntax.TokenComma {
			return true
		}
	}
	return false
}

func (o *ObjectConsExpr) newlineToken() *Token {
	return &Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte{'\n'},
	}
}

// ObjectConsItem is one item of an object-construct expression: a key and
// its value.
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

// ObjectConsKeyExpr is the key of an item of an object-construct expression.
// It is an identifier, a quoted string, or another expression whose value
// gives the key.
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node
	// literal is set when the key is a naked identifier, which names the
	// attribute literally instead of referring to a variable.
	literal bool
}

func newObjectConsKeyExpr(wrapped *node, literal bool) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
		literal: literal,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// name returns the string the key gives, if the key is an identifier or a
// quoted string literal. Otherwise it returns false.
func (k *ObjectConsKeyExpr) name() (string, bool) {
	tokens := k.BuildTokens(nil)
	if len(tokens) == 1 && tokens[0].Type == hclsyntax.TokenIdent {
		return string(tokens[0].Bytes), true
	}
	if len(tokens) < 2 || tokens[0].Type != hclsyntax.TokenOQuote || tokens[len(tokens)-1].Type != hclsyntax.TokenCQuote {
		return "", false
	}
	var b strings.Builder
	for _, tok := range tokens[1 : len(tokens)-1] {
		if tok.Type != hclsyntax.TokenQuotedLit {
			return "", false
		}
		s, diags := hclsyntax.ParseStringLiteralToken(hclsyntax.Token{Type: tok.Type, Bytes: tok.Bytes})
		if diags.HasErrors() {
			return "", false
		}
		b.WriteString(s)
	}
	return b.String(), true
}

// ObjectConsValue is the value of an item of an object-construct expression.
type ObjectConsValue struct {
	inTree
	expr *node
}

func newObjectConsValue() *ObjectConsValue {
	return &ObjectConsValue{
		inTree: newInTree(),
	}
}

func (o *ObjectConsValue) Expr() *Expression {
	return o.expr.content.(*Expression)
}
