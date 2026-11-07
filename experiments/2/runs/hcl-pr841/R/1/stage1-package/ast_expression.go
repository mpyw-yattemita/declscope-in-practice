// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

type Expression struct {
	inTree

	absTraversals nodeSet
	wrapped       *node
}

func newExpression() *Expression {
	return &Expression{
		inTree:        newInTree(),
		absTraversals: newNodeSet(),
	}
}

func (e *Expression) wrap(c nodeContent) {
	if e.wrapped != nil {
		panic("wrap called on already-wrapped Expression")
	}

	e.wrapped = newNode(c)
	e.children.AppendNode(e.wrapped)
}

// NewExpressionRaw constructs an expression containing the given raw tokens.
//
// There is no automatic validation that the given tokens produce a valid
// expression. Callers of thus function must take care to produce invalid
// expression tokens. Where possible, use the higher-level functions
// NewExpressionLiteral or NewExpressionAbsTraversal instead.
//
// Because NewExpressionRaw does not interpret the given tokens in any way,
// an expression created by NewExpressionRaw will produce an empty result
// for calls to its method Variables, even if the given token sequence
// contains a subslice that would normally be interpreted as a traversal under
// parsing.
func NewExpressionRaw(tokens Tokens) *Expression {
	expr := newExpression()
	// We copy the tokens here in order to make sure that later mutations
	// by the caller don't inadvertently cause our expression to become
	// invalid.
	copyTokens := make(Tokens, len(tokens))
	copy(copyTokens, tokens)
	expr.children.AppendUnstructuredTokens(copyTokens)
	return expr
}

// NewExpressionLiteral constructs an an expression that represents the given
// literal value.
//
// Since an unknown value cannot be represented in source code, this function
// will panic if the given value is unknown or contains a nested unknown value.
// Use val.IsWhollyKnown before calling to be sure.
//
// HCL native syntax does not directly represent lists, maps, and sets, and
// instead relies on the automatic conversions to those collection types from
// either list or tuple constructor syntax. Therefore converting collection
// values to source code and re-reading them will lose type information, and
// the reader must provide a suitable type at decode time to recover the
// original value.
func NewExpressionLiteral(val cty.Value) *Expression {
	toks := TokensForValue(val)
	expr := newExpression()
	expr.children.AppendUnstructuredTokens(toks)
	return expr
}

// NewExpressionAbsTraversal constructs an expression that represents the
// given traversal, which must be absolute or this function will panic.
func NewExpressionAbsTraversal(traversal hcl.Traversal) *Expression {
	if traversal.IsRelative() {
		panic("can't construct expression from relative traversal")
	}

	physT := newTraversal()
	rootName := traversal.RootName()
	steps := traversal[1:]

	{
		tn := newTraverseName()
		tn.name = tn.children.Append(newIdentifier(&Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(rootName),
		}))
		physT.steps.Add(physT.children.Append(tn))
	}

	for _, step := range steps {
		switch ts := step.(type) {
		case hcl.TraverseAttr:
			tn := newTraverseName()
			tn.children.AppendUnstructuredTokens(Tokens{
				{
					Type:  hclsyntax.TokenDot,
					Bytes: []byte{'.'},
				},
			})
			tn.name = tn.children.Append(newIdentifier(&Token{
				Type:  hclsyntax.TokenIdent,
				Bytes: []byte(ts.Name),
			}))
			physT.steps.Add(physT.children.Append(tn))
		case hcl.TraverseIndex:
			ti := newTraverseIndex()
			ti.children.AppendUnstructuredTokens(Tokens{
				{
					Type:  hclsyntax.TokenOBrack,
					Bytes: []byte{'['},
				},
			})
			indexExpr := NewExpressionLiteral(ts.Key)
			ti.key = ti.children.Append(indexExpr)
			ti.children.AppendUnstructuredTokens(Tokens{
				{
					Type:  hclsyntax.TokenCBrack,
					Bytes: []byte{']'},
				},
			})
			physT.steps.Add(physT.children.Append(ti))
		}
	}

	expr := newExpression()
	expr.absTraversals.Add(expr.children.Append(physT))
	return expr
}

// Variables returns the absolute traversals that exist within the receiving
// expression, in source order.
//
// This includes the traversals nested in the arguments of a function call,
// the elements of a tuple-construct expression, and the keys and values of
// the items of an object-construct expression. A naked identifier used as an
// object key is a literal name rather than a variable, so it is not included.
func (e *Expression) Variables() []*Traversal {
	nodes := e.traversalNodes()
	ret := make([]*Traversal, len(nodes))
	for i, node := range nodes {
		ret[i] = node.content.(*Traversal)
	}
	return ret
}

// nestedTraversals is implemented by the structured content that an
// Expression can wrap and that contains nested expressions, such as an
// object-construct expression.
type nestedTraversals interface {
	// traversalNodes returns the nodes of the absolute traversals in the
	// nested expressions, in source order.
	traversalNodes() []*node
}

// traversalNodes returns the nodes of the absolute traversals within the
// receiving expression, including those in nested expressions, in source
// order. Each returned node has content of type *Traversal.
func (e *Expression) traversalNodes() []*node {
	ret := e.absTraversals.List()
	if e.wrapped != nil {
		if nested, ok := e.wrapped.content.(nestedTraversals); ok {
			ret = append(ret, nested.traversalNodes()...)
		}
	}
	return ret
}

// exprsTraversalNodes returns the nodes of the absolute traversals within
// the given expressions, in source order. Each node in the set must have
// content of type *Expression.
func exprsTraversalNodes(exprs nodeSet) []*node {
	var ret []*node
	for _, n := range exprs.List() {
		ret = append(ret, n.content.(*Expression).traversalNodes()...)
	}
	return ret
}

// RenameVariablePrefix examines each of the absolute traversals in the
// receiving expression to see if they have the given sequence of names as
// a prefix prefix. If so, they are updated in place to have the given
// replacement names instead of that prefix.
//
// This can be used to implement symbol renaming. The calling application can
// visit all relevant expressions in its input and apply the same renaming
// to implement a global symbol rename.
//
// The search and replacement traversals must be the same length, or this
// method will panic. Only attribute access operations can be matched and
// replaced. Index steps never match the prefix.
//
// The traversals examined are the same ones that Variables returns.
func (e *Expression) RenameVariablePrefix(search, replacement []string) {
	if len(search) != len(replacement) {
		panic(fmt.Sprintf("search and replacement length mismatch (%d and %d)", len(search), len(replacement)))
	}
	for _, node := range e.traversalNodes() {
		node.content.(*Traversal).renamePrefix(search, replacement)
	}
}

// AsObjectConsExpr returns the object constructor expression that the
// receiver consists of, or nil if the receiver is any other kind of
// expression.
func (e *Expression) AsObjectConsExpr() *ObjectConsExpr {
	return wrappedAs[*ObjectConsExpr](e)
}

// AsFunctionCallExpr returns the wrapped function call expression, or nil if
// this expression does not parse as a function call expression.
func (e *Expression) AsFunctionCallExpr() *FunctionCallExpr {
	return wrappedAs[*FunctionCallExpr](e)
}

// AsTupleConsExpr returns the wrapped tuple-construct expression, or nil if
// this expression does not parse as a tuple-construct expression.
func (e *Expression) AsTupleConsExpr() *TupleConsExpr {
	return wrappedAs[*TupleConsExpr](e)
}

func (e *Expression) asQuoted() *quoted {
	return wrappedAs[*quoted](e)
}

// wrappedAs returns the structured content that the expression wraps if it
// has type T, or the zero value of T otherwise.
func wrappedAs[T nodeContent](e *Expression) T {
	var zero T
	if e.wrapped == nil {
		return zero
	}
	if c, ok := e.wrapped.content.(T); ok {
		return c
	}
	return zero
}

func (e *Expression) AsQuotedLiteral() Tokens {
	if quoted := e.asQuoted(); quoted == nil {
		return Tokens{}
	} else {
		return quoted.tokens
	}
}

// Traversal represents a sequence of variable, attribute, and/or index
// operations.
type Traversal struct {
	inTree

	steps nodeSet
}

func newTraversal() *Traversal {
	return &Traversal{
		inTree: newInTree(),
		steps:  newNodeSet(),
	}
}

// renamePrefix replaces the leading names of the traversal with the names in
// replacement, if those leading names equal the names in search. Only name
// steps can match. The caller must make sure that search and replacement have
// the same length.
func (t *Traversal) renamePrefix(search, replacement []string) {
	if len(t.steps) < len(search) {
		// If it's shorter then it can't have our prefix
		return
	}

	stepNodes := t.steps.List()
	for i, name := range search {
		step, isName := stepNodes[i].content.(*TraverseName)
		if !isName {
			return // only name nodes can match
		}
		if string(step.name.content.(*identifier).token.Bytes) != name {
			return
		}
	}

	// If we get here then the prefix matched, so now we'll swap in
	// the replacement strings.
	for i, name := range replacement {
		step := stepNodes[i].content.(*TraverseName)
		token := step.name.content.(*identifier).token
		token.Bytes = []byte(name)
	}
}

type TraverseName struct {
	inTree

	name *node
}

func newTraverseName() *TraverseName {
	return &TraverseName{
		inTree: newInTree(),
	}
}

type TraverseIndex struct {
	inTree

	key *node
}

func newTraverseIndex() *TraverseIndex {
	return &TraverseIndex{
		inTree: newInTree(),
	}
}
