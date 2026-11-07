// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// TupleConsExpr is a tuple-construct expression, such as ["fedora", "fez"].
// Use Expression.AsTupleConsExpr to get one from an expression.
//
// Each element is its own Expression. The brackets and the commas are
// unstructured tokens of the tuple.
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

// variables returns the absolute traversals in the elements of the tuple,
// in source order.
func (tuple *TupleConsExpr) variables() []*Traversal {
	return variablesOf(tuple.elems)
}
