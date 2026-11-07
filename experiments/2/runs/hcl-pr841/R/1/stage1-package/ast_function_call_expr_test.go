// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/v2"
)

// parseTestExpr parses src and returns the expression of its attribute "a".
func parseTestExpr(t *testing.T, src string) (*File, *Expression) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	return f, f.Body().GetAttribute("a").Expr()
}

// checkRoundTrip checks that the tokens of f are exactly src, and that
// File.Bytes returns src too if src is already formatted.
func checkRoundTrip(t *testing.T, f *File, src string) {
	t.Helper()
	if got := string(f.children.BuildTokens(nil).Bytes()); got != src {
		t.Errorf("wrong tokens\ngot:  %q\nwant: %q", got, src)
	}
	if string(Format([]byte(src))) != src {
		return
	}
	if got := string(f.Bytes()); got != src {
		t.Errorf("wrong File.Bytes result\ngot:  %q\nwant: %q", got, src)
	}
}

// exprsStrings returns the source of each expression in the given set,
// without surrounding spaces.
func exprsStrings(exprs nodeSet) []string {
	var ret []string
	for _, n := range exprs.List() {
		ret = append(ret, tokensString(n.content.(*Expression).BuildTokens(nil)))
	}
	return ret
}

func TestFunctionCallExprParse(t *testing.T) {
	tests := map[string]struct {
		src      string
		wantName string
		wantArgs []string
	}{
		"no arguments":      {`a = f()`, "f", nil},
		"one argument":      {`a = upper("fez")`, "upper", []string{`"fez"`}},
		"nested tuple":      {`a = join(" ", ["fedora", "fez"])`, "join", []string{`" "`, `["fedora", "fez"]`}},
		"trailing comma":    {`a = f(1, 2,)`, "f", []string{"1", "2"}},
		"expand final":      {`a = f(1, list...)`, "f", []string{"1", "list"}},
		"namespaced":        {`a = provider::example::upper("fez")`, "provider::example::upper", []string{`"fez"`}},
		"operation":         {`a = f(1 + 2, x)`, "f", []string{"1 + 2", "x"}},
		"nested object":     {`a = f({ k = 1 })`, "f", []string{"{ k = 1 }"}},
		"multi-line":        {"a = f(\n  1, # one\n  2,\n)", "f", []string{"1", "2"}},
		"unformatted":       {"a = f (  1,2 , )", "f", []string{"1", "2"}},
		"template argument": {`a = f("${x}")`, "f", []string{`"${x}"`}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			src := test.src + "\n"
			f, expr := parseTestExpr(t, src)

			call := expr.AsFunctionCallExpr()
			if call == nil {
				t.Fatalf("expression is not a function call")
			}
			if expr.AsTupleConsExpr() != nil || expr.AsObjectConsExpr() != nil {
				t.Errorf("function call is also reported as another kind of expression")
			}
			if got := call.Name(); got != test.wantName {
				t.Errorf("wrong name %q; want %q", got, test.wantName)
			}
			if diff := cmp.Diff(test.wantArgs, exprsStrings(call.args)); diff != "" {
				t.Errorf("wrong arguments (-want +got):\n%s", diff)
			}
			checkRoundTrip(t, f, src)
		})
	}
}
