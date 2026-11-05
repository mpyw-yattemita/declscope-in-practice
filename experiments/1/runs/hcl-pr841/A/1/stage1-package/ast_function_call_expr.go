// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Use Expression.AsFunctionCallExpr to get one.
//
// Each argument is its own Expression. The function name, the parentheses,
// the commas and an expansion symbol after the last argument are tokens of
// the FunctionCallExpr itself.
type FunctionCallExpr struct {
	inTree

	// name holds the tokens of the function name, which can be namespaced
	// such as provider::example::func.
	name *node
	args nodeSet
}

func newFunctionCallExpr() *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function.
func (call *FunctionCallExpr) Name() string {
	var name []byte
	for _, tok := range call.name.content.(Tokens) {
		name = append(name, tok.Bytes...)
	}
	return string(name)
}

// variables returns the traversals in the arguments, in source order.
func (call *FunctionCallExpr) variables() []*Traversal {
	return expressionsVariables(call.args)
}

// expressionsVariables returns the traversals in each of the given
// expression nodes, in source order.
func expressionsVariables(exprs nodeSet) []*Traversal {
	var ret []*Traversal
	for _, n := range exprs.List() {
		ret = append(ret, n.content.(*Expression).Variables()...)
	}
	return ret
}
