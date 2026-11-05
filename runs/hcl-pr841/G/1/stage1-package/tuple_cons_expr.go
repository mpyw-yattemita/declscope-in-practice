// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Use Expression.AsTupleConsExpr to get one from a parsed
// expression.
//
// Each element is an Expression of its own. The brackets and the commas stay
// as tokens of the tuple.
type TupleConsExpr struct {
	inTree

	elems nodeSet
}

// variables returns the absolute traversals in the elements of the tuple, in
// source order.
//
func (tuple *TupleConsExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, n := range tuple.elems.List() {
		ret = append(ret, n.content.(*Expression).Variables()...)
	}
	return ret
}

// parseTupleConsExpr parses a tuple-construct expression, defined as:
//
//	tuple = "[" (
//	    (Expression (("," | Newline) Expression)* ","?)?
//	) "]";
//
// Each element becomes an Expression of its own. The tokens between them,
// and any comments, stay as unstructured tokens of the tuple.
//
func parseTupleConsExpr(nativeExpr *hclsyntax.TupleConsExpr, from inputTokens) *node {
	tuple := &TupleConsExpr{
		inTree: newInTree(),
		elems:  newNodeSet(),
	}

	var before, elemTokens inputTokens
	for _, nativeElem := range nativeExpr.Exprs {
		before, elemTokens, from = from.Partition(nativeElem.Range())
		tuple.children.AppendUnstructuredTokens(before.Tokens())
		elem := parseExpression(nativeElem, elemTokens)
		tuple.children.AppendNode(elem)
		tuple.elems.Add(elem)
	}
	tuple.children.AppendUnstructuredTokens(from.Tokens())

	wrapExpr := newExpression()
	wrapExpr.wrap(tuple)
	return newNode(wrapExpr)
}
