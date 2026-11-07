// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object constructor expression, such as
// { hat = "derby", (cat) = "calico" }. Use Expression.AsObjectConsExpr to get
// the ObjectConsExpr of an expression.
//
// Its methods look up items by key and set item values, in a similar way to
// how the methods of Body look up and set attributes.
type ObjectConsExpr struct {
	inTree

	items nodeSet

	// close is the node holding the tokens after the last item, up to and
	// including the closing brace.
	close *node
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

// ItemFor returns the first item whose key evaluates to the given string, or
// nil if there is no such item.
//
// Only keys whose value is known without evaluating an expression are
// considered: a naked identifier such as hat, or a quoted string literal such
// as "hat". The two forms are equivalent, and a quoted key is compared after
// its escape sequences are decoded. Keys that are expressions, such as (hat),
// hat.band or "${hat}", never match.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if item.KeyExpr().hasName(key) {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor returns for the
// given key, or nil if there is no such item.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	item := o.ItemFor(key)
	if item == nil {
		return nil
	}
	return item.ValueExpr()
}

// SetItem either replaces the value of the item that ItemFor returns for the
// given key, or adds a new item with the given key to the end of the object.
// The key of an existing item is left as written.
//
// A new key is written as a naked identifier if it is a valid identifier, and
// otherwise as a quoted string literal that evaluates to the given key.
//
// The given expression must not already be part of a syntax tree, or this
// method will panic.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	item := o.ItemFor(key)
	if item != nil {
		item.ValueExpr().setExpr(expr)
	} else {
		item = o.appendItem(newObjectConsKeyExprForName(key), expr)
	}
	return item.KeyExpr(), item.ValueExpr()
}

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the value.
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

// SetItemTraversal is like SetItem, but sets the value to the given traversal,
// which must be absolute, as NewExpressionAbsTraversal does. To set a literal
// value, use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// appendItem adds a new item at the end of the object, laid out in the style
// of the object.
//
// In an object written on a single line, the new item goes on the same line,
// separated from the previous item by a comma. Otherwise, it goes on its own
// line. A comma or line comment that ends the line of the previous item stays
// with that item. Everything after that stays with the closing brace, which
// is moved to its own line if needed.
func (o *ObjectConsExpr) appendItem(key *ObjectConsKeyExpr, expr *Expression) *ObjectConsItem {
	lead, rest := splitLineEnd(o.close.content.(Tokens))
	if o.isMultiline() {
		if len(lead) == 0 || !isLineBreak(lead[len(lead)-1]) {
			lead = append(lead, newNewlineToken())
		}
		rest = append(Tokens{newNewlineToken()}, rest...)
	} else if len(o.items) > 0 && !hasComma(lead) {
		lead = append(lead, newCommaToken())
	}

	item := newObjectConsItem()
	item.init(lead, key, expr)
	o.items.Add(o.children.Insert(o.close, item))
	o.close = o.close.ReplaceWith(rest)
	return item
}

// isMultiline reports whether the object is written on more than one line.
// Only the tokens that belong to the object itself count, so line breaks
// inside a key or value, such as in a nested object, do not.
func (o *ObjectConsExpr) isMultiline() bool {
	for n := o.children.first; n != nil; n = n.after {
		switch c := n.content.(type) {
		case Tokens:
			if hasLineBreak(c) {
				return true
			}
		case *ObjectConsItem:
			if c.hasLineBreak() {
				return true
			}
		}
	}
	return false
}

// traversalNodes returns the nodes of the absolute traversals in the keys and
// values of the object's items, in source order.
func (o *ObjectConsExpr) traversalNodes() []*node {
	var ret []*node
	for _, item := range o.Items() {
		ret = append(ret, item.KeyExpr().traversalNodes()...)
		ret = append(ret, item.ValueExpr().Expr().traversalNodes()...)
	}
	return ret
}

// ObjectConsItem is one item of an ObjectConsExpr: a key and a value, such as
// hat = "derby".
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

// init fills a new item with the given key and value. The lead tokens are
// the separators that come before the key, such as a comma or a newline.
func (item *ObjectConsItem) init(lead Tokens, key *ObjectConsKeyExpr, expr *Expression) {
	item.children.AppendUnstructuredTokens(lead)
	item.key = item.children.Append(key)
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

// hasLineBreak reports whether the tokens of the item itself, outside of its
// key and value, contain a line break.
func (item *ObjectConsItem) hasLineBreak() bool {
	for n := item.children.first; n != nil; n = n.after {
		if toks, ok := n.content.(Tokens); ok && hasLineBreak(toks) {
			return true
		}
	}
	return false
}

// objectConsKeyKind is the syntactic form of an object-construct key.
type objectConsKeyKind int

const (
	// objectConsKeyExpression is a key that is evaluated as an expression,
	// such as (hat), hat.band or "${hat}".
	objectConsKeyExpression objectConsKeyKind = iota

	// objectConsKeyLiteralName is a naked identifier, such as hat. It is a
	// literal name, not a variable.
	objectConsKeyLiteralName

	// objectConsKeyQuotedLiteral is a quoted string literal, such as "hat".
	objectConsKeyQuotedLiteral
)

// ObjectConsKeyExpr is the key of an ObjectConsItem, such as hat in
// hat = "derby".
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node

	kind objectConsKeyKind

	// name is the string that the key evaluates to, with any escape
	// sequences decoded. It is set only if kind is not
	// objectConsKeyExpression.
	name string
}

// newObjectConsKeyExpr wraps the given node, which must have content of type
// *Expression.
func newObjectConsKeyExpr(wrapped *node, kind objectConsKeyKind, name string) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
		kind:    kind,
		name:    name,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// newObjectConsKeyExprForName returns a new key that evaluates to the given
// name. The key is a naked identifier if the name is a valid identifier, and
// a quoted string literal otherwise.
func newObjectConsKeyExprForName(name string) *ObjectConsKeyExpr {
	var tokens Tokens
	var kind objectConsKeyKind
	// A naked "for" as the first key would start a for expression instead of
	// an object, so it is always quoted.
	if hclsyntax.ValidIdentifier(name) && name != "for" {
		tokens = TokensForIdentifier(name)
		kind = objectConsKeyLiteralName
	} else {
		tokens = TokensForValue(cty.StringVal(name))
		kind = objectConsKeyQuotedLiteral
	}
	tokens[0].SpacesBefore = 1

	return newObjectConsKeyExpr(newNode(NewExpressionRaw(tokens)), kind, name)
}

// hasName reports whether the key is a literal name or a quoted string
// literal that evaluates to the given name.
func (k *ObjectConsKeyExpr) hasName(name string) bool {
	return k.kind != objectConsKeyExpression && k.name == name
}

// traversalNodes returns the nodes of the absolute traversals in the key, in
// source order. A literal name key has none.
func (k *ObjectConsKeyExpr) traversalNodes() []*node {
	if k.kind == objectConsKeyLiteralName {
		return nil
	}
	return k.wrapped.content.(*Expression).traversalNodes()
}

// ObjectConsValue is the value of an ObjectConsItem, such as "derby" in
// hat = "derby".
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

// setExpr replaces the expression of the value with the given one, which must
// not be part of a syntax tree.
func (o *ObjectConsValue) setExpr(expr *Expression) {
	o.expr = o.expr.ReplaceWith(expr)
}

// splitLineEnd splits the tokens that follow the last item of an object into
// the tokens that end the line of that item, and the rest.
//
// The line end is an optional comma and any inline comments, up to and
// including a newline or a line comment. If there is no line break before
// the next token that is not a comma or a comment, the line end is just a
// leading comma, if any.
//
// The returned line end can be appended to without changing the rest.
func splitLineEnd(tokens Tokens) (lineEnd, rest Tokens) {
	for i, tok := range tokens {
		if isLineBreak(tok) {
			return tokens[: i+1 : i+1], tokens[i+1:]
		}
		if tok.Type != hclsyntax.TokenComma && tok.Type != hclsyntax.TokenComment {
			break
		}
	}
	if len(tokens) > 0 && tokens[0].Type == hclsyntax.TokenComma {
		return tokens[:1:1], tokens[1:]
	}
	return nil, tokens
}

// isLineBreak reports whether the token ends a line: a newline, or a line
// comment, which includes its trailing newline.
func isLineBreak(tok *Token) bool {
	switch tok.Type {
	case hclsyntax.TokenNewline:
		return true
	case hclsyntax.TokenComment:
		return len(tok.Bytes) > 0 && tok.Bytes[len(tok.Bytes)-1] == '\n'
	default:
		return false
	}
}

func hasLineBreak(tokens Tokens) bool {
	for _, tok := range tokens {
		if isLineBreak(tok) {
			return true
		}
	}
	return false
}

func hasComma(tokens Tokens) bool {
	for _, tok := range tokens {
		if tok.Type == hclsyntax.TokenComma {
			return true
		}
	}
	return false
}

func newNewlineToken() *Token {
	return &Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte{'\n'},
	}
}

func newCommaToken() *Token {
	return &Token{
		Type:  hclsyntax.TokenComma,
		Bytes: []byte{','},
	}
}
