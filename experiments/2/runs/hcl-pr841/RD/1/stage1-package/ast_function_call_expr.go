// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Use Expression.AsFunctionCallExpr to get one from an
// expression.
//
// Each argument is an Expression of its own. The function name, the
// parentheses, the commas and an expansion symbol (...) after the last
// argument are kept as tokens of the call.
type FunctionCallExpr struct {
	inTree

	args nodeSet
}

func newFunctionCallExpr() *FunctionCallExpr {
	return &FunctionCallExpr{
		inTree: newInTree(),
		args:   newNodeSet(),
	}
}

// Name returns the name of the called function. For a function in a
// namespace, the name includes the namespace, as in provider::aws::arn_parse.
func (call *FunctionCallExpr) Name() string {
	var name []byte
	for _, n := range call.children.List() {
		tokens, ok := n.content.(Tokens)
		if !ok {
			// The name and the opening parenthesis come before any argument.
			break
		}
		for _, tok := range tokens {
			switch tok.Type {
			case hclsyntax.TokenOParen:
				return string(name)
			case hclsyntax.TokenIdent, hclsyntax.TokenDoubleColon:
				name = append(name, tok.Bytes...)
			}
		}
	}
	return string(name)
}

// traversals appends the absolute traversals in the arguments to to, in
// source order, and returns the result.
//
func (call *FunctionCallExpr) traversals(to []*Traversal) []*Traversal {
	for _, n := range call.args.List() {
		to = n.content.(*Expression).traversals(to)
	}
	return to
}
