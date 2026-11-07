// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// FunctionCallExpr is a function call expression, such as join(" ", list).
// Use Expression.AsFunctionCallExpr to get one from an expression.
//
// Each argument is its own Expression. The function name, the parentheses,
// the commas and a final "..." are unstructured tokens of the call.
type FunctionCallExpr struct {
	inTree

	name string
	args nodeSet
}

func newFunctionCallExpr(name string) *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		name:   name,
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function. A namespaced function name
// includes its namespaces, such as provider::local::func.
func (call *FunctionCallExpr) Name() string {
	return call.name
}

// variables returns the absolute traversals in the arguments of the call,
// in source order.
func (call *FunctionCallExpr) variables() []*Traversal {
	return variablesOf(call.args)
}
