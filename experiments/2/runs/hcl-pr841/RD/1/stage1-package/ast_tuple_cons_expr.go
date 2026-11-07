// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Use Expression.AsTupleConsExpr to get one from an
// expression.
//
// Each element is an Expression of its own. The brackets and the commas are
// kept as tokens of the tuple.
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

// traversals appends the absolute traversals in the elements to to, in
// source order, and returns the result.
//
func (tuple *TupleConsExpr) traversals(to []*Traversal) []*Traversal {
	for _, n := range tuple.elems.List() {
		to = n.content.(*Expression).traversals(to)
	}
	return to
}
