// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestExpressionVariablesObjectCons(t *testing.T) {
	checkASTExpressionVariables(t, []astExpressionVariablesCase{
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
			"a = { (old) = 1 }\n",
			[]string{"old"},
			"a = { (new) = 1 }\n",
		},
		{
			"a = { \"k\" = old.x }\n",
			[]string{"old.x"},
			"a = { \"k\" = new.x }\n",
		},
		{
			"a = { (old.x) = old.y, k = { old = [f(old.z)] } }\n",
			[]string{"old.x", "old.y", "old.z"},
			"a = { (new.x) = new.y, k = { old = [f(new.z)] } }\n",
		},
	})
}

func TestExpressionVariablesCallAndTuple(t *testing.T) {
	checkASTExpressionVariables(t, []astExpressionVariablesCase{
		{
			"a = f(old.x)\n",
			[]string{"old.x"},
			"a = f(new.x)\n",
		},
		{
			"a = f(old.x, g(old.y), [old.z])\n",
			[]string{"old.x", "old.y", "old.z"},
			"a = f(new.x, g(new.y), [new.z])\n",
		},
		{
			"a = f(old.x...)\n",
			[]string{"old.x"},
			"a = f(new.x...)\n",
		},
		{
			"a = [old.x, [old.y], f(old.z)]\n",
			[]string{"old.x", "old.y", "old.z"},
			"a = [new.x, [new.y], f(new.z)]\n",
		},
		{
			"a = [old.x, { k = old.y }]\n",
			[]string{"old.x", "old.y"},
			"a = [new.x, { k = new.y }]\n",
		},
		{
			"a = [\n  old.x, # old.w\n  f(\n    old.y,\n  ),\n]\n",
			[]string{"old.x", "old.y"},
			"a = [\n  new.x, # old.w\n  f(\n    new.y,\n  ),\n]\n",
		},
		{
			"a = f(old.x + old.y, old.z[0])\n",
			[]string{"old.x", "old.y", "old.z[0]"},
			"a = f(new.x + new.y, new.z[0])\n",
		},
		{
			"a = f()\n",
			nil,
			"a = f()\n",
		},
		{
			"a = []\n",
			nil,
			"a = []\n",
		},
	})
}

// astExpressionVariablesCase is a source whose expression a has its
// variables checked by checkASTExpressionVariables.
type astExpressionVariablesCase struct {
	src     string
	want    []string // the traversals Variables reports, as source text
	renamed string   // the source after renaming old to new
}

// checkASTExpressionVariables checks that Variables reports the traversals of
// each case in source order, as many as hclsyntax does, and that
// RenameVariablePrefix renames them.
func checkASTExpressionVariables(t *testing.T, tests []astExpressionVariablesCase) {
	t.Helper()
	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			if got := string(f.Bytes()); got != test.src {
				t.Errorf("source does not round-trip\ngot:  %q\nwant: %q", got, test.src)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if len(got) != len(test.want) {
				t.Fatalf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Fatalf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
				}
			}

			native, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags)
			}
			nativeExpr := native.Body.(*hclsyntax.Body).Attributes["a"].Expr
			if want := len(nativeExpr.Variables()); len(got) != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", len(got), want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after renaming\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}
