// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ObjectConsExpr is an object-construct expression, such as
// { hat = "derby", (cat) = "calico" }. Use Expression.AsObjectConsExpr to
// get one from an expression.
//
// An item can be looked up and set by its key only when the key is written
// as an identifier or as a quoted string literal. The two forms are
// equivalent: the key "hat" names both hat = ... and "hat" = .... A key
// that is taken from a variable or another expression, such as (hat), is
// never found by its name.
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
	items := make([]*ObjectConsItem, 0, len(nodes))
	for _, n := range nodes {
		items = append(items, n.content.(*ObjectConsItem))
	}
	return items
}

// ItemFor returns the first item whose key is the given string, or nil if
// there is no such item.
//
// A quoted key is compared by its value after its escape sequences are
// decoded, so ItemFor("${hat}") finds "$${hat}" = .... Keys that are not an
// identifier or a quoted string literal never match.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if item.KeyExpr().hasName(key) {
			return item
		}
	}
	return nil
}

// ValueExprFor returns the value of the item that ItemFor finds for the
// given key, or nil if there is no such item.
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue {
	if item := o.ItemFor(key); item != nil {
		return item.ValueExpr()
	}
	return nil
}

// SetItem either replaces the value of the item that ItemFor finds for the
// given key, or adds a new item to the end of the object, with the given
// expression as its value. The key of an existing item is left as written.
//
// A new key is written as an identifier if it is a valid one, and otherwise
// as a quoted string literal that evaluates to the given key. The key "for"
// is always quoted, since HCL would read it as a for expression.
//
// SetItem panics if the expression is already part of a syntax tree, since
// two items would then share it.
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
// NewExpressionLiteral does. To set a variable reference or other
// traversal, use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but writes the given traversal, which
// must be absolute, as NewExpressionAbsTraversal does. To set a literal
// value, use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// traversals appends the absolute traversals in the items' keys and values to
// to, in source order, and returns the result.
//
func (o *ObjectConsExpr) traversals(to []*Traversal) []*Traversal {
	for _, item := range o.Items() {
		to = item.KeyExpr().traversals(to)
		to = item.ValueExpr().Expr().traversals(to)
	}
	return to
}

// appendItem adds a new item with the given key and value after the last
// item of the object, and returns it.
//
// The tokens that follow the last item on its line, such as a comma or a
// line comment, stay with that item. Everything after them stays with the
// closing brace. In an object written on a single line the new item follows
// a comma, and otherwise it goes on its own line.
func (o *ObjectConsExpr) appendItem(key string, expr *Expression) *ObjectConsItem {
	closing := o.children.Last()
	tail, ok := closing.content.(Tokens)
	if !ok {
		panic("object-construct expression does not end with its closing brace")
	}
	multiline := o.multiline()

	split := 0
	for split < len(tail) && tail[split].Type == hclsyntax.TokenComma {
		split++
	}
	if multiline {
		for i, tok := range tail {
			if tokenIsNewline(tok) {
				split = i + 1
				break
			}
		}
	}
	before := append(Tokens{}, tail[:split]...)
	rest := append(Tokens{}, tail[split:]...)

	switch {
	case multiline:
		if len(before) == 0 || !tokenIsNewline(before[len(before)-1]) {
			before = append(before, &Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
		}
		rest = append(Tokens{{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}}, rest...)
	case len(before) == 0 && len(o.items) > 0:
		before = Tokens{{Type: hclsyntax.TokenComma, Bytes: []byte{','}}}
	}

	item := newObjectConsItem()
	item.init(before, key, expr)
	o.items.Add(o.children.Insert(closing, item))
	closing.ReplaceWith(rest)
	return item
}

// multiline reports whether the object is written on several lines. Only
// the line breaks between its items count, not those inside a key or a
// value.
func (o *ObjectConsExpr) multiline() bool {
	for _, n := range o.children.List() {
		switch content := n.content.(type) {
		case Tokens:
			if objectConsBreaksLine(content) {
				return true
			}
		case *ObjectConsItem:
			for _, m := range content.children.List() {
				if tokens, ok := m.content.(Tokens); ok && objectConsBreaksLine(tokens) {
					return true
				}
			}
		}
	}
	return false
}

// objectConsBreaksLine reports whether the tokens between the parts of an
// object contain a line break.
func objectConsBreaksLine(tokens Tokens) bool {
	for _, tok := range tokens {
		if tokenIsNewline(tok) {
			return true
		}
	}
	return false
}

// ObjectConsItem is one item of an object-construct expression: a key and
// the value assigned to it.
type ObjectConsItem struct {
	inTree

	key *node
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

// init builds a new item with the given key and value, after the given
// tokens that separate it from what comes before.
func (item *ObjectConsItem) init(before Tokens, key string, expr *Expression) {
	item.children.AppendUnstructuredTokens(before)
	item.key = item.children.Append(newObjectConsKeyExprNamed(key))
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

// ObjectConsKeyExpr is the key of an item of an object-construct
// expression. It is written as an identifier, as a quoted string, or as
// another expression, such as a variable in parentheses.
type ObjectConsKeyExpr struct {
	inTree

	wrapped *node

	// name is the string the key evaluates to, when literal is true.
	name string

	// literal is true when the key is a naked identifier or a quoted string
	// literal. Such a key names the item instead of referring to a variable.
	literal bool
}

func newObjectConsKeyExpr(wrapped *node, name string, literal bool) *ObjectConsKeyExpr {
	expr := &ObjectConsKeyExpr{
		inTree:  newInTree(),
		wrapped: wrapped,
		name:    name,
		literal: literal,
	}
	expr.children.AppendNode(wrapped)

	return expr
}

// newObjectConsKeyExprNamed builds a key that evaluates to name. It is an
// identifier if name is a valid one, and a quoted string literal otherwise.
//
// The identifier for is quoted too: as the first key of an object, HCL would
// read it as the start of a for expression.
func newObjectConsKeyExprNamed(name string) *ObjectConsKeyExpr {
	var expr *Expression
	if hclsyntax.ValidIdentifier(name) && name != "for" {
		expr = NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: name}})
	} else {
		expr = NewExpressionLiteral(cty.StringVal(name))
	}
	return newObjectConsKeyExpr(newNode(expr), name, true)
}

// hasName reports whether the key is a literal name equal to name.
func (k *ObjectConsKeyExpr) hasName(name string) bool {
	return k.literal && k.name == name
}

// traversals appends the absolute traversals in the key to to, and returns
// the result. A literal name contributes none.
func (k *ObjectConsKeyExpr) traversals(to []*Traversal) []*Traversal {
	if k.literal {
		return to
	}
	return k.wrapped.content.(*Expression).traversals(to)
}

// ObjectConsValue is the value of an item of an object-construct
// expression.
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
