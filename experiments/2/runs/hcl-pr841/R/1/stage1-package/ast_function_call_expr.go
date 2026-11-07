// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Use Expression.AsFunctionCallExpr to get the
// FunctionCallExpr of an expression.
//
// Each argument is its own Expression. The function name, the parentheses,
// the commas between the arguments and any expansion symbol after the last
// argument are tokens of the FunctionCallExpr itself.
type FunctionCallExpr struct {
	inTree

	name string

	// args holds the nodes of the arguments, in source order. Each node has
	// content of type *Expression.
	args nodeSet
}

func newFunctionCallExpr(name string) *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		name:   name,
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function. A namespaced name, such as
// provider::example::upper, is returned in full.
func (call *FunctionCallExpr) Name() string {
	return call.name
}

// traversalNodes returns the nodes of the absolute traversals in the
// arguments of the call, in source order.
func (call *FunctionCallExpr) traversalNodes() []*node {
	return exprsTraversalNodes(call.args)
}
