// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Use Expression.AsTupleConsExpr to get the TupleConsExpr
// of an expression.
//
// Each element is its own Expression. The brackets and the commas between
// the elements are tokens of the TupleConsExpr itself.
type TupleConsExpr struct {
	inTree

	// exprs holds the nodes of the elements, in source order. Each node has
	// content of type *Expression.
	exprs nodeSet
}

func newTupleConsExpr() *TupleConsExpr {
	return &TupleConsExpr{
		inTree: newInTree(),
		exprs:  newNodeSet(),
	}
}

// traversalNodes returns the nodes of the absolute traversals in the
// elements of the tuple, in source order.
func (tuple *TupleConsExpr) traversalNodes() []*node {
	return exprsTraversalNodes(tuple.exprs)
}
