// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// TupleConsExpr represents a tuple-construct expression such as
// ["fedora", "fez"]. Get one from an expression with
// Expression.AsTupleConsExpr.
//
// Each element is an Expression of its own. The brackets and the commas stay
// as tokens of the tuple.
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

// variableNodes returns the traversal nodes in the elements, in source
// order.
//
func (tuple *TupleConsExpr) variableNodes() []*node {
	var ret []*node
	for _, n := range tuple.elems.List() {
		ret = append(ret, n.content.(*Expression).variableNodes()...)
	}
	return ret
}
