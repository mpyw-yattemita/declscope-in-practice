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
// { hat = "derby", (cat) = "calico" }. Get one from an expression with
// Expression.AsObjectConsExpr.
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
	var ret []*ObjectConsItem
	for _, n := range o.items.List() {
		ret = append(ret, n.content.(*ObjectConsItem))
	}
	return ret
}

// ItemFor returns the first item whose key is the given string, or nil if
// there is none.
//
// Only an item whose key is an identifier or a quoted string literal can be
// found, and the two forms are equivalent: "hat" finds both hat = ... and
// "hat" = .... A quoted key is compared after its escape sequences are
// decoded. A key that is taken from a variable or another expression, such
// as (hat), hat.band or "${hat}", is never found.
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem {
	for _, item := range o.Items() {
		if name, ok := item.KeyExpr().literalName(); ok && name == key {
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
// given key, leaving its key as written, or adds a new item to the end of
// the object.
//
// A new key is written as an identifier when it is a valid one, and as a
// quoted string literal otherwise. A new item follows the layout of the
// object: on the same line after a comma in a single-line object, and on
// its own line before the closing brace in a multi-line one.
//
// SetItem panics if the given expression is already part of a syntax tree,
// since two items would then share it.
//
// The return values are the key and the value of the item that was either
// modified in-place or created.
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue) {
	if expr.inList() {
		panic("SetItem called with an expression that is already part of a syntax tree")
	}

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

// SetItemValue is like SetItem, but the value is given as a cty.Value, and
// must therefore be a literal, as for NewExpressionLiteral. To set a
// variable reference or other traversal, use SetItemTraversal.
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionLiteral(val))
}

// SetItemTraversal is like SetItem, but the value is given as a
// hcl.Traversal, which must be absolute, as for NewExpressionAbsTraversal.
// To set a literal value, use SetItemValue.
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue) {
	return o.SetItem(key, NewExpressionAbsTraversal(traversal))
}

// itemExpressions returns the key and value expressions of each item, in
// source order.
//
func (o *ObjectConsExpr) itemExpressions() []*Expression {
	var ret []*Expression
	for _, n := range o.items.List() {
		item := n.content.(*ObjectConsItem)
		key := item.key.content.(*ObjectConsKeyExpr)
		ret = append(ret, key.wrapped.content.(*Expression), item.value.content.(*ObjectConsValue).Expr())
	}
	return ret
}

// appendItem adds a new item with the given key and value after the last
// item of the object, laid out as the object already is.
//
// The tokens after the last item are split in two. A comma or a line comment
// that follows the last item on its line stays with that item. The rest,
// which ends with the closing brace, stays after the new item.
func (o *ObjectConsExpr) appendItem(key string, expr *Expression) *ObjectConsItem {
	rest := o.detachTrailingTokens()
	multiLine := o.hasLineBreak(rest)

	cut, lineEnded, hasComma := 0, false, false
	for i, tok := range rest {
		if tokenIsNewline(tok) {
			cut, lineEnded = i+1, true
			break
		}
		if tok.Type == hclsyntax.TokenComma {
			// Keep a trailing comma with the last item, and anything after
			// it on the line, such as an inline comment, with the closing
			// brace.
			cut, hasComma = i+1, true
		}
	}
	head, tail := rest[:cut], append(Tokens(nil), rest[cut:]...)

	var lead Tokens
	switch {
	case multiLine && !lineEnded:
		// The closing brace shares a line with the last item, so the new
		// item starts a line of its own.
		lead = Tokens{{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}}
	case !multiLine && len(o.items) > 0 && !hasComma:
		lead = Tokens{{Type: hclsyntax.TokenComma, Bytes: []byte{','}}}
	}
	if multiLine {
		tail = append(Tokens{{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}}, tail...)
	} else if len(tail) > 0 && tail[0].SpacesBefore == 0 {
		spaced := *tail[0]
		spaced.SpacesBefore = 1
		tail[0] = &spaced
	}

	item := newObjectConsItem()
	item.init(key, expr, lead)

	o.children.AppendUnstructuredTokens(head)
	o.items.Add(o.children.Append(item))
	o.children.AppendUnstructuredTokens(tail)
	return item
}

// detachTrailingTokens removes and returns the tokens after the last item,
// or after the opening brace if the object has no items.
func (o *ObjectConsExpr) detachTrailingTokens() Tokens {
	var trailing []*node
	o.walkChildNodes(func(n *node) {
		if _, ok := n.content.(Tokens); ok {
			trailing = append(trailing, n)
		} else {
			trailing = nil
		}
	})

	var rest Tokens
	for _, n := range trailing {
		rest = append(rest, n.content.(Tokens)...)
		n.Detach()
	}

	if len(o.items) == 0 {
		for i, tok := range rest {
			if tok.Type == hclsyntax.TokenOBrace {
				o.children.AppendUnstructuredTokens(rest[:i+1])
				return rest[i+1:]
			}
		}
	}
	return rest
}

// hasLineBreak reports whether the object is written on several lines. Only
// the tokens between its items count, along with the given tokens that were
// detached from it, and not those inside an item's key or value.
func (o *ObjectConsExpr) hasLineBreak(detached Tokens) bool {
	layout := append(Tokens(nil), detached...)
	o.walkChildNodes(func(n *node) {
		switch c := n.content.(type) {
		case Tokens:
			layout = append(layout, c...)
		case *ObjectConsItem:
			c.walkChildNodes(func(in *node) {
				if toks, ok := in.content.(Tokens); ok {
					layout = append(layout, toks...)
				}
			})
		}
	})
	for _, tok := range layout {
		if tokenIsNewline(tok) {
			return true
		}
	}
	return false
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

// init fills a new item with the given key and value, after the given tokens
// that separate it from what comes before it.
//
// The key is written as an identifier if it is a valid one, and as a quoted
// string literal otherwise, escaped so that it evaluates back to the key.
// The identifier for is quoted too, since an object that starts with it
// would be read as a for expression.
func (item *ObjectConsItem) init(key string, expr *Expression, lead Tokens) {
	var keyTokens Tokens
	if hclsyntax.ValidIdentifier(key) && key != "for" {
		keyTokens = TokensForIdentifier(key)
	} else {
		keyTokens = TokensForValue(cty.StringVal(key))
	}
	keyTokens[0].SpacesBefore = 1

	item.children.AppendUnstructuredTokens(lead)
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

// ObjectConsKeyExpr is the key of an item of an object-construct
// expression. It is an identifier, a quoted string, or another expression
// whose value is the key.
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

// literalName returns the key as a string, if it is an identifier or a
// quoted string literal. A quoted key has its escape sequences decoded.
func (k *ObjectConsKeyExpr) literalName() (string, bool) {
	toks := k.wrapped.BuildTokens(nil)
	if len(toks) == 1 && toks[0].Type == hclsyntax.TokenIdent {
		return string(toks[0].Bytes), true
	}

	last := len(toks) - 1
	if last < 1 || toks[0].Type != hclsyntax.TokenOQuote || toks[last].Type != hclsyntax.TokenCQuote {
		return "", false
	}
	var b strings.Builder
	for _, tok := range toks[1:last] {
		if tok.Type != hclsyntax.TokenQuotedLit {
			// A template interpolation or directive.
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
