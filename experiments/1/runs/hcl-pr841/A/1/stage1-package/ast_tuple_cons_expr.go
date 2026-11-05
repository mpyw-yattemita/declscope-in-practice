// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Use Expression.AsTupleConsExpr to get one.
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

// variables returns the traversals in the elements, in source order.
func (tuple *TupleConsExpr) variables() []*Traversal {
	return expressionsVariables(tuple.elems)
}
