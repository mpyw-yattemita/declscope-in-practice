// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// FunctionCallExpr represents a function call expression such as
// join(" ", list). Use Expression.AsFunctionCallExpr to get one from a parsed
// expression.
//
// Each argument is an Expression of its own. The function name, the
// parentheses, the commas and a final ellipsis stay as tokens of the call.
type FunctionCallExpr struct {
	inTree

	args nodeSet
}

// Name returns the name of the called function. A namespaced name, such as
// provider::time::rfc3339_parse, is returned whole.
func (call *FunctionCallExpr) Name() string {
	var b strings.Builder
	for _, tok := range call.BuildTokens(nil) {
		if tok.Type == hclsyntax.TokenOParen {
			break
		}
		if tok.Type == hclsyntax.TokenIdent || tok.Type == hclsyntax.TokenDoubleColon {
			b.Write(tok.Bytes)
		}
	}
	return b.String()
}

// variables returns the absolute traversals in the arguments of the call, in
// source order.
//
func (call *FunctionCallExpr) variables() []*Traversal {
	var ret []*Traversal
	for _, n := range call.args.List() {
		ret = append(ret, n.content.(*Expression).Variables()...)
	}
	return ret
}

// parseFunctionCallExpr parses a function call expression, defined as:
//
//	FunctionCall = Identifier ("::" Identifier)* "(" Arguments ")";
//	Arguments = (
//	    () ||
//	    (Expression ("," Expression)* ("," | "...")?)
//	);
//
// Each argument becomes an Expression of its own. The tokens between them,
// and any comments, stay as unstructured tokens of the call.
//
func parseFunctionCallExpr(nativeExpr *hclsyntax.FunctionCallExpr, from inputTokens) *node {
	call := &FunctionCallExpr{
		inTree: newInTree(),
		args:   newNodeSet(),
	}

	var before, argTokens inputTokens
	for _, nativeArg := range nativeExpr.Args {
		before, argTokens, from = from.Partition(nativeArg.Range())
		call.children.AppendUnstructuredTokens(before.Tokens())
		arg := parseExpression(nativeArg, argTokens)
		call.children.AppendNode(arg)
		call.args.Add(arg)
	}
	call.children.AppendUnstructuredTokens(from.Tokens())

	wrapExpr := newExpression()
	wrapExpr.wrap(call)
	return newNode(wrapExpr)
}
