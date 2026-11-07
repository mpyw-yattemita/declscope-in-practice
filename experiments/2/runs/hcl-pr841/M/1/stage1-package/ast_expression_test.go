// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestExpressionVariablesObjectCons(t *testing.T) {
	tests := []struct {
		src     string
		want    []string // Variables, in order, as source text
		renamed string   // source after renaming "old" to "new"
	}{
		{
			"a = { k = old.x }\n",
			[]string{"old.x"},
			"a = { k = new.x }\n",
		},
		{
			"a = { k = old.x, j = other.y, l = old.z }\n",
			[]string{"old.x", "other.y", "old.z"},
			"a = { k = new.x, j = other.y, l = new.z }\n",
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
			"a = {\n  k = old.x\n  j = {\n    l = [f(old.y)]\n  }\n}\n",
			[]string{"old.x", "old.y"},
			"a = {\n  k = new.x\n  j = {\n    l = [f(new.y)]\n  }\n}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("wrong variables\n%s", diff)
			}

			nativeFile, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags.Error())
			}
			nativeExpr := nativeFile.Body.(*hclsyntax.Body).Attributes["a"].Expr
			if got, want := len(expr.Variables()), len(nativeExpr.Variables()); got != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", got, want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestExpressionVariablesFunctionCallAndTupleCons(t *testing.T) {
	tests := []struct {
		src     string
		want    []string // Variables, in order, as source text
		renamed string   // source after renaming "old" to "new"
	}{
		{
			"a = f(old.x)\n",
			[]string{"old.x"},
			"a = f(new.x)\n",
		},
		{
			"a = f(old.x, other.y, old.z...)\n",
			[]string{"old.x", "other.y", "old.z"},
			"a = f(new.x, other.y, new.z...)\n",
		},
		{
			"a = [old.x, other.y, old.z]\n",
			[]string{"old.x", "other.y", "old.z"},
			"a = [new.x, other.y, new.z]\n",
		},
		{
			"a = f(old.w, [old.x, g(old.y)], { k = old.z })\n",
			[]string{"old.w", "old.x", "old.y", "old.z"},
			"a = f(new.w, [new.x, g(new.y)], { k = new.z })\n",
		},
		{
			"a = [\n  old.x,\n  f(old.y + 1),\n]\n",
			[]string{"old.x", "old.y"},
			"a = [\n  new.x,\n  f(new.y + 1),\n]\n",
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
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			if got := string(f.Bytes()); got != test.src {
				t.Errorf("wrong result before rename\ngot:  %q\nwant: %q", got, test.src)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("wrong variables\n%s", diff)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestExpressionAsFunctionCallAndTupleCons(t *testing.T) {
	tests := []struct {
		src      string
		wantCall string // name of the function, or "" if not a call
		wantArgs int    // number of arguments or elements
		isTuple  bool
	}{
		{`join(" ", ["fedora", "fez"])`, "join", 2, false},
		{`f()`, "f", 0, false},
		{`provider::local::upper("hat")`, "provider::local::upper", 1, false},
		{`["fedora", "fez"]`, "", 2, true},
		{`[]`, "", 0, true},
		{`f(1) + 1`, "", 0, false},
		{`[1][0]`, "", 0, false},
		{`[for x in y : x]`, "", 0, false},
		{`{ k = [1] }`, "", 0, false},
		{`"fedora"`, "", 0, false},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = "+test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()

			call := expr.AsFunctionCallExpr()
			switch {
			case test.wantCall == "" && call != nil:
				t.Errorf("unexpected function call %q", call.Name())
			case test.wantCall != "" && call == nil:
				t.Errorf("not a function call")
			case call != nil:
				if got := call.Name(); got != test.wantCall {
					t.Errorf("wrong name %q; want %q", got, test.wantCall)
				}
				if got := len(call.args); got != test.wantArgs {
					t.Errorf("wrong number of arguments %d; want %d", got, test.wantArgs)
				}
			}

			tuple := expr.AsTupleConsExpr()
			switch {
			case !test.isTuple && tuple != nil:
				t.Errorf("unexpected tuple")
			case test.isTuple && tuple == nil:
				t.Errorf("not a tuple")
			case tuple != nil:
				if got := len(tuple.elems); got != test.wantArgs {
					t.Errorf("wrong number of elements %d; want %d", got, test.wantArgs)
				}
			}
		})
	}
}
