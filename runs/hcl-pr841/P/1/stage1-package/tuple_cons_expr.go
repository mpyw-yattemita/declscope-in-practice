// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import "github.com/hashicorp/hcl/v2/hclsyntax"

// TupleConsExpr represents a tuple-construct expression such as ["fedora", "fez"].
//
// Each element is its own Expression. The brackets and the commas are tokens
// of the TupleConsExpr itself.
type TupleConsExpr struct {
	inTree

	elems nodeSet
}

func newTupleConsExpr() *TupleConsExpr {
	return &TupleConsExpr{
		inTree: newInTree(),
		elems:  newNodeSet(),
	}
}

// parseTupleConsExpr parses a tuple-construct expression, defined as:
//
//	tuple = "[" (
//	    (Expression (("," | Newline) Expression)* ","?)?
//	) "]";
//
// It is called by parseExpression in parser.go, and calls it back for each
// element.
func parseTupleConsExpr(nativeExpr *hclsyntax.TupleConsExpr, from inputTokens) *node {
	tuple := newTupleConsExpr()
	children := tuple.children

	var before, elemTokens inputTokens
	for _, nativeElem := range nativeExpr.Exprs {
		before, elemTokens, from = from.Partition(nativeElem.Range())
		children.AppendUnstructuredTokens(before.Tokens())
		elemNode := parseExpression(nativeElem, elemTokens)
		children.AppendNode(elemNode)
		tuple.elems.Add(elemNode)
	}
	children.AppendUnstructuredTokens(from.Tokens())

	wrapExpr := newExpression()
	wrapExpr.wrap(tuple)
	return newNode(wrapExpr)
}
