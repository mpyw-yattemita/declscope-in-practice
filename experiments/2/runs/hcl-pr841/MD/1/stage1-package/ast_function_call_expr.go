// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Get one from an expression with
// Expression.AsFunctionCallExpr.
//
// Each argument is an Expression of its own. The function name, the
// parentheses, the commas and a final "..." stay as tokens of the call.
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

// Name returns the name of the called function. A namespaced name keeps its
// namespaces, as in "provider::example::upper".
func (call *FunctionCallExpr) Name() string {
	// The parser puts the name and the opening parenthesis in the first run
	// of tokens.
	var name []byte
	done := false
	call.walkChildNodes(func(n *node) {
		tokens, ok := n.content.(Tokens)
		if done || !ok {
			done = true
			return
		}
		for _, tok := range tokens {
			switch tok.Type {
			case hclsyntax.TokenOParen:
				done = true
				return
			case hclsyntax.TokenIdent, hclsyntax.TokenDoubleColon:
				name = append(name, tok.Bytes...)
			}
		}
	})
	return string(name)
}

// variableNodes returns the traversal nodes in the arguments, in source
// order.
//
func (call *FunctionCallExpr) variableNodes() []*node {
	var ret []*node
	for _, n := range call.args.List() {
		ret = append(ret, n.content.(*Expression).variableNodes()...)
	}
	return ret
}
