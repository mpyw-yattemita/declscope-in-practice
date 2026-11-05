// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// FunctionCallExpr represents a function call expression such as join(" ", list).
//
// Each argument is its own Expression. The function name, the parentheses,
// the commas and any ellipsis after the last argument are tokens of the
// FunctionCallExpr itself.
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

// Name returns the name of the called function. A namespaced name, such as
// provider::example::upper, is returned with its namespaces.
func (call *FunctionCallExpr) Name() string {
	// The name and the opening parenthesis are the first tokens of the call.
	var toks Tokens
	if first := call.children.first; first != nil {
		toks, _ = first.content.(Tokens)
	}
	var b strings.Builder
	for _, tok := range toks {
		switch tok.Type {
		case hclsyntax.TokenIdent, hclsyntax.TokenDoubleColon:
			b.Write(tok.Bytes)
		case hclsyntax.TokenOParen:
			return b.String()
		}
	}
	return b.String()
}

// parseFunctionCallExpr parses a function call expression, defined as:
//
//	FunctionCall = Identifier ("::" Identifier)* "(" Arguments ")";
//	Arguments = (
//	    () ||
//	    (Expression ("," Expression)* ("," | "...")?)
//	);
//
// It is called by parseExpression in parser.go, and calls it back for each
// argument.
func parseFunctionCallExpr(nativeExpr *hclsyntax.FunctionCallExpr, from inputTokens) *node {
	call := newFunctionCallExpr()
	children := call.children

	var before, argTokens inputTokens
	for _, nativeArg := range nativeExpr.Args {
		before, argTokens, from = from.Partition(nativeArg.Range())
		children.AppendUnstructuredTokens(before.Tokens())
		argNode := parseExpression(nativeArg, argTokens)
		children.AppendNode(argNode)
		call.args.Add(argNode)
	}
	children.AppendUnstructuredTokens(from.Tokens())

	wrapExpr := newExpression()
	wrapExpr.wrap(call)
	return newNode(wrapExpr)
}
