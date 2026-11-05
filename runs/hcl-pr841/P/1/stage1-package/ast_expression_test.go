// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestExpressionVariablesObjectCons(t *testing.T) {
	tests := []struct {
		src     string
		want    []string // source of each traversal, in order
		renamed string
	}{
		{
			"a = { k = old.x }\n",
			[]string{"old.x"},
			"a = { k = new.x }\n",
		},
		{
			"a = { k = old.x, j = old.y }\n",
			[]string{"old.x", "old.y"},
			"a = { k = new.x, j = new.y }\n",
		},
		{
			"a = { k = { j = old.x } }\n",
			[]string{"old.x"},
			"a = { k = { j = new.x } }\n",
		},
		{
			"a = { k = [old.x] }\n",
			[]string{"old.x"},
			"a = { k = [new.x] }\n",
		},
		{
			"a = { k = f(old.x) }\n",
			[]string{"old.x"},
			"a = { k = f(new.x) }\n",
		},
		{
			"a = { old = 1 }\n",
			nil,
			"a = { old = 1 }\n",
		},
		{
			"a = { old = old.x }\n",
			[]string{"old.x"},
			"a = { old = new.x }\n",
		},
		{
			"a = { (old.x) = 1 }\n",
			[]string{"old.x"},
			"a = { (new.x) = 1 }\n",
		},
		{
			"a = { (old.x) = old.y }\n",
			[]string{"old.x", "old.y"},
			"a = { (new.x) = new.y }\n",
		},
		{
			"a = { \"k\" = old.x }\n",
			[]string{"old.x"},
			"a = { \"k\" = new.x }\n",
		},
		{
			"a = { k = { old = 1 } }\n",
			nil,
			"a = { k = { old = 1 } }\n",
		},
		{
			"a = {\n  k = old.x\n  j = {\n    (old.y) = [f(old.z)]\n  }\n}\n",
			[]string{"old.x", "old.y", "old.z"},
			"a = {\n  k = new.x\n  j = {\n    (new.y) = [f(new.z)]\n  }\n}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if fmt.Sprint(got) != fmt.Sprint(test.want) {
				t.Errorf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
			}

			nativeFile, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags)
			}
			nativeAttrs, _ := nativeFile.Body.JustAttributes()
			if nativeVars := nativeAttrs["a"].Expr.Variables(); len(got) != len(nativeVars) {
				t.Errorf("got %d variables, but hclsyntax reports %d", len(got), len(nativeVars))
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestExpressionVariablesFunctionCallTupleCons(t *testing.T) {
	tests := []struct {
		src     string
		want    []string // source of each traversal, in order
		renamed string
	}{
		{"a = f(old.x)\n", []string{"old.x"}, "a = f(new.x)\n"},
		{"a = f(old.x, 1, old.y)\n", []string{"old.x", "old.y"}, "a = f(new.x, 1, new.y)\n"},
		{"a = f(old.x...)\n", []string{"old.x"}, "a = f(new.x...)\n"},
		{"a = [old.x, old.y]\n", []string{"old.x", "old.y"}, "a = [new.x, new.y]\n"},
		{"a = f([old.x], g(old.y))\n", []string{"old.x", "old.y"}, "a = f([new.x], g(new.y))\n"},
		{"a = [f(old.x), { k = old.y }]\n", []string{"old.x", "old.y"}, "a = [f(new.x), { k = new.y }]\n"},
		{"a = { k = f(old.x) }\n", []string{"old.x"}, "a = { k = f(new.x) }\n"},
		{"a = { k = [old.x] }\n", []string{"old.x"}, "a = { k = [new.x] }\n"},
		{"a = f(old.x[0] + old.y)\n", []string{"old.x[0]", "old.y"}, "a = f(new.x[0] + new.y)\n"},
		{"a = old(1)\n", nil, "a = old(1)\n"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if fmt.Sprint(got) != fmt.Sprint(test.want) {
				t.Errorf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestExpressionAsFunctionCallTupleCons(t *testing.T) {
	tests := []struct {
		src      string
		wantCall bool
		wantTup  bool
	}{
		{`a = f(1)`, true, false},
		{`a = [1]`, false, true},
		{`a = f(1)[0]`, false, false},
		{`a = [for x in xs : x]`, false, false},
		{`a = { k = 1 }`, false, false},
		{`a = x`, false, false},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			expr := f.Body().GetAttribute("a").Expr()
			if got := expr.AsFunctionCallExpr() != nil; got != test.wantCall {
				t.Errorf("AsFunctionCallExpr() != nil is %t; want %t", got, test.wantCall)
			}
			if got := expr.AsTupleConsExpr() != nil; got != test.wantTup {
				t.Errorf("AsTupleConsExpr() != nil is %t; want %t", got, test.wantTup)
			}
		})
	}
}
