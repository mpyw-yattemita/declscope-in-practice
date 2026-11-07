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

func TestExpressionVariablesNested(t *testing.T) {
	tests := []struct {
		src     string
		want    []string // the traversals Variables reports, in order
		renamed string   // the source after renaming old to new
	}{
		{
			src:     "a = { k = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = new.x }\n",
		},
		{
			src:     "a = { k = old.x, j = old.y }\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = { k = new.x, j = new.y }\n",
		},
		{
			src:     "a = { k = { j = old.x } }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = { j = new.x } }\n",
		},
		{
			src:     "a = { k = [old.x] }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = [new.x] }\n",
		},
		{
			src:     "a = { k = f(old.x) }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = f(new.x) }\n",
		},
		{
			// A naked identifier key is a literal name, not a variable.
			src:     "a = { old = 1 }\n",
			want:    nil,
			renamed: "a = { old = 1 }\n",
		},
		{
			src:     "a = { old = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { old = new.x }\n",
		},
		{
			src:     "a = { (old.x) = 1 }\n",
			want:    []string{"old.x"},
			renamed: "a = { (new.x) = 1 }\n",
		},
		{
			src:     "a = { (old.x) = old.y }\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = { (new.x) = new.y }\n",
		},
		{
			src:     "a = { \"k\" = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { \"k\" = new.x }\n",
		},
		{
			src:     "a = { k = { old = 1 } }\n",
			want:    nil,
			renamed: "a = { k = { old = 1 } }\n",
		},
		{
			src:     "a = {}\n",
			want:    nil,
			renamed: "a = {}\n",
		},
		{
			src:     "a = f(old.x)\n",
			want:    []string{"old.x"},
			renamed: "a = f(new.x)\n",
		},
		{
			src:     "a = f(old.x, g(old.y), [old.z])\n",
			want:    []string{"old.x", "old.y", "old.z"},
			renamed: "a = f(new.x, g(new.y), [new.z])\n",
		},
		{
			src:     "a = f(old.x...)\n",
			want:    []string{"old.x"},
			renamed: "a = f(new.x...)\n",
		},
		{
			src:     "a = provider::p::f(old.x)\n",
			want:    []string{"old.x"},
			renamed: "a = provider::p::f(new.x)\n",
		},
		{
			src:     "a = [old.x, { k = old.y }, [old.z]]\n",
			want:    []string{"old.x", "old.y", "old.z"},
			renamed: "a = [new.x, { k = new.y }, [new.z]]\n",
		},
		{
			src:     "a = { k = f(old.x), j = [old.y] }\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = { k = f(new.x), j = [new.y] }\n",
		},
		{
			src:     "a = [f(old.x), old.y.z]\n",
			want:    []string{"old.x", "old.y.z"},
			renamed: "a = [f(new.x), new.y.z]\n",
		},
		{
			src:     "a = f()\n",
			want:    nil,
			renamed: "a = f()\n",
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
				t.Errorf("wrong traversals (-want +got):\n%s", diff)
			}

			native, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags.Error())
			}
			nativeVars := native.Body.(*hclsyntax.Body).Attributes["a"].Expr.Variables()
			if len(got) != len(nativeVars) {
				t.Errorf("got %d traversals, but hclsyntax reports %d", len(got), len(nativeVars))
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after renaming\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestExpressionAsFunctionCallExprAndTupleConsExpr(t *testing.T) {
	tests := []struct {
		src       string // the expression assigned to a
		wantCall  bool
		wantTuple bool
	}{
		{`f(1)`, true, false},
		{`[1, 2]`, false, true},
		{`[f(1)]`, false, true},
		{`(f(1))`, false, false},
		{`f(1)[0]`, false, false},
		{`[for x in xs : x]`, false, false},
		{`{ k = [1] }`, false, false},
		{`x`, false, false},
		{`"f(1)"`, false, false},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = "+test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()
			if got := expr.AsFunctionCallExpr() != nil; got != test.wantCall {
				t.Errorf("AsFunctionCallExpr returned a call: %v, want %v", got, test.wantCall)
			}
			if got := expr.AsTupleConsExpr() != nil; got != test.wantTuple {
				t.Errorf("AsTupleConsExpr returned a tuple: %v, want %v", got, test.wantTuple)
			}
		})
	}
}
