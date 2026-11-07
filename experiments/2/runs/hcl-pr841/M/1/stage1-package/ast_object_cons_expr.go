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
// { hat = "derby" }. Use Expression.AsObjectConsExpr to get one from an
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
	ret := make([]*ObjectConsItem, len(nodes))
	for i, n := range nodes {
		ret[i] = n.content.(*ObjectConsItem)
	}
	return ret
}

// ItemFor returns the item of the object that has the given key, or nil if
// there is no such item.
//
// Only keys that are an identifier or a quoted string literal can match. The
// two forms are equivalent, so ItemFor("hat") finds both hat = ... and
// "hat" = .... A quoted key is compared by its value after escape sequences
// are decoded. Keys that are taken from a variable or another expression,
// such as (hat) or "${hat}", never match.
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
	if item := o.ItemFor(key); item != nil {
		return item.ValueExpr()
	}
	return nil
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, or adds a new item to the end of the object. An existing key is
// left as written. A new key is written as an identifier if it is a valid
// identifier, and as a quoted string literal otherwise.
//
// The expression must not already be part of a syntax tree, or SetItem
// panics. Where possible, use SetItemValue or SetItemTraversal instead.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	expr.assertUnattached()

	item := o.ItemFor(key)
	if item == nil {
		item = o.appendItem(newObjectConsKeyExprForName(key))
	} else {
		old := item.ValueExpr()
		old.Expr().parent = nil
		old.children.Clear()
	}

	value := item.ValueExpr()
	value.expr = value.children.Append(expr)
	expr.parent = item.value
	return item.KeyExpr(), value
}

// SetItemRaw is like SetItem, but uses the given tokens verbatim as the
// value.
//
// The same caveats apply to this function as for NewExpressionRaw on which
// it is based. If possible, prefer to use SetItemValue or SetItemTraversal.
func (o *ObjectConsExpr) SetItemRaw(key string, tokens Tokens) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionRaw(tokens))
}

// SetItemValue is like SetItem, but the value is given as a cty.Value, and
// must therefore be a literal. To set a variable reference or other
// traversal, use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but the value is given as a
// hcl.Traversal, which must be an absolute traversal. To set a literal value,
// use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// appendItem adds a new item with the given key and an empty value to the
// end of the object, and returns it.
//
// In an object written on a single line, the item goes on the same line
// after a comma. Otherwise it goes on its own line after the last item. A
// comma or a line comment that follows the last item on its line stays with
// that item. Anything after that, such as comments on the following lines,
// stays with the closing brace.
func (o *ObjectConsExpr) appendItem(key *ObjectConsKeyExpr) *ObjectConsItem {
	// The parser always puts the tokens after the last item, up to and
	// including the closing brace, in the last child node.
	closing := o.children.last
	rest := closing.content.(Tokens)

	var lead, tail Tokens
	if o.isMultiLine() {
		// Keep everything up to the end of the last item's line before the
		// new item. If the closing brace is on that line, keep only a comma.
		end := 0
		for i, t := range rest {
			if isLineEnd(t) {
				end = i + 1
				break
			}
		}
		if end == 0 && len(rest) > 0 && rest[0].Type == hclsyntax.TokenComma {
			end = 1
		}
		lead = append(lead, rest[:end]...)
		if end == 0 || !isLineEnd(rest[end-1]) {
			lead = append(lead, newNewlineToken())
		}
		tail = append(Tokens{newNewlineToken()}, rest[end:]...)
	} else {
		tail = rest
		if len(tail) > 0 && tail[0].Type == hclsyntax.TokenComma {
			tail = tail[1:]
		}
		if len(o.items) > 0 {
			lead = Tokens{{Type: hclsyntax.TokenComma, Bytes: []byte{','}}}
		}
	}

	item := newObjectConsItem()
	item.children.AppendUnstructuredTokens(lead)
	item.key = item.children.Append(key)
	item.children.AppendUnstructuredTokens(Tokens{
		{Type: hclsyntax.TokenEqual, Bytes: []byte{'='}, SpacesBefore: 1},
	})
	item.value = item.children.Append(newObjectConsValue())

	closing.Detach()
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(tail)
	return item
}

// isMultiLine returns true if the object is written on several lines. Only
// the tokens between the items and the braces count, so line breaks inside
// a key or a value do not make the object multi-line.
func (o *ObjectConsExpr) isMultiLine() bool {
	if hasLineEnd(o.children) {
		return true
	}
	for _, item := range o.Items() {
		if hasLineEnd(item.children) {
			return true
		}
	}
	return false
}

// hasLineEnd returns true if any of the unstructured tokens directly in ns
// ends a line.
func hasLineEnd(ns *nodes) bool {
	for n := ns.first; n != nil; n = n.after {
		if tokens, ok := n.content.(Tokens); ok {
			for _, t := range tokens {
				if isLineEnd(t) {
					return true
				}
			}
		}
	}
	return false
}

// isLineEnd returns true if the token ends a line. A line comment includes
// the newline that ends it.
func isLineEnd(t *Token) bool {
	return t.Type == hclsyntax.TokenNewline ||
		(t.Type == hclsyntax.TokenComment && bytes.HasSuffix(t.Bytes, []byte{'\n'}))
}

func newNewlineToken() *Token {
	return &Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}
}

// ObjectConsItem is one key and value pair of an object-construct
// expression.
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

// ObjectConsKeyExpr is the key of an item in an object-construct expression.
// It is an identifier, such as hat, or an expression, such as "hat" or
// (hat).
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node

	// literal is true when the key is a naked identifier, which hclsyntax
	// treats as a literal name rather than as a variable reference.
	literal bool
}

func newObjectConsKeyExpr(wrapped *node) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// newObjectConsKeyExprForName returns a key that evaluates to the given name.
// It is an identifier if name is a valid identifier, and a quoted string
// literal otherwise. "for" is always quoted, because { for = ... } starts a
// for expression.
func newObjectConsKeyExprForName(name string) *ObjectConsKeyExpr {
	if hclsyntax.ValidIdentifier(name) && name != "for" {
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: name}})
		key := newObjectConsKeyExpr(newNode(expr))
		key.literal = true
		return key
	}

	// TokensForValue escapes the string so that it evaluates back to name.
	expr := newExpression()
	expr.wrapped = expr.children.Append(newQuoted(TokensForValue(cty.StringVal(name))))
	return newObjectConsKeyExpr(newNode(expr))
}

// name returns the string that the key evaluates to, if the key is an
// identifier or a quoted string literal. Otherwise it returns false.
func (k *ObjectConsKeyExpr) name() (string, bool) {
	expr := k.wrapped.content.(*Expression)

	if k.literal {
		tokens := expr.BuildTokens(nil)
		if len(tokens) == 1 && tokens[0].Type == hclsyntax.TokenIdent {
			return string(tokens[0].Bytes), true
		}
		return "", false
	}

	quoted := expr.asQuoted()
	if quoted == nil {
		return "", false
	}
	// The parser only makes a quoted node for a string literal, so it is
	// safe to evaluate it without a context to decode its escape sequences.
	native, diags := hclsyntax.ParseExpression(quoted.tokens.Bytes(), "", hcl.InitialPos)
	if diags.HasErrors() {
		return "", false
	}
	val, diags := native.Value(nil)
	if diags.HasErrors() || !val.IsKnown() || val.IsNull() || !val.Type().Equals(cty.String) {
		return "", false
	}
	return val.AsString(), true
}

// variables returns the absolute traversals in the keys and values of the
// object's items, in source order. Literal keys contribute no traversals.
func (o *ObjectConsExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, n := range o.items.List() {
		item := n.content.(*ObjectConsItem)
		key := item.key.content.(*ObjectConsKeyExpr)
		if !key.literal {
			ret = append(ret, key.wrapped.content.(*Expression).Variables()...)
		}
		ret = append(ret, item.value.content.(*ObjectConsValue).Expr().Variables()...)
	}
	return ret
}

// ObjectConsValue is the value of an item in an object-construct expression.
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
