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
		vars    []string
		renamed string
	}{
		{
			src:     "a = { k = old.x }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = new.x }\n",
		},
		{
			src:     "a = { k = { j = old.x } }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = { j = new.x } }\n",
		},
		{
			src:     "a = { k = [old.x] }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = [new.x] }\n",
		},
		{
			src:     "a = { k = f(old.x) }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = f(new.x) }\n",
		},
		{
			src:     "a = { k = \"${old.x}\" }\n",
			vars:    []string{"old.x"},
			renamed: "a = { k = \"${new.x}\" }\n",
		},
		{
			// A naked identifier key is a literal name, not a variable.
			src:     "a = { old = 1 }\n",
			vars:    nil,
			renamed: "a = { old = 1 }\n",
		},
		{
			src:     "a = { old = old.x }\n",
			vars:    []string{"old.x"},
			renamed: "a = { old = new.x }\n",
		},
		{
			// A parenthesized key is an expression.
			src:     "a = { (old.x) = 1 }\n",
			vars:    []string{"old.x"},
			renamed: "a = { (new.x) = 1 }\n",
		},
		{
			src:     "a = { \"k\" = old.x }\n",
			vars:    []string{"old.x"},
			renamed: "a = { \"k\" = new.x }\n",
		},
		{
			src:     "a = { (old.a) = old.b, k = { j = old.c }, l = other.d }\n",
			vars:    []string{"old.a", "old.b", "old.c", "other.d"},
			renamed: "a = { (new.a) = new.b, k = { j = new.c }, l = other.d }\n",
		},
		{
			src:     "a = {}\n",
			vars:    nil,
			renamed: "a = {}\n",
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
			if diff := cmp.Diff(test.vars, got); diff != "" {
				t.Errorf("wrong variables\n%s", diff)
			}

			nativeFile, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags.Error())
			}
			nativeAttrs, _ := nativeFile.Body.JustAttributes()
			if want := len(nativeAttrs["a"].Expr.Variables()); len(got) != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", len(got), want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if diff := cmp.Diff(test.renamed, string(f.Bytes())); diff != "" {
				t.Errorf("wrong result after rename\n%s", diff)
			}
		})
	}
}

func TestExpressionVariablesFunctionCallAndTupleCons(t *testing.T) {
	tests := []struct {
		src     string
		vars    []string
		renamed string
	}{
		{
			src:     "a = f(old.x, old.y)\n",
			vars:    []string{"old.x", "old.y"},
			renamed: "a = f(new.x, new.y)\n",
		},
		{
			src:     "a = [old.x, other.y]\n",
			vars:    []string{"old.x", "other.y"},
			renamed: "a = [new.x, other.y]\n",
		},
		{
			src:     "a = f(old.a, [old.b, g(old.c)], { k = old.d }, old.e...)\n",
			vars:    []string{"old.a", "old.b", "old.c", "old.d", "old.e"},
			renamed: "a = f(new.a, [new.b, g(new.c)], { k = new.d }, new.e...)\n",
		},
		{
			src:     "a = { k = f(old.x, [old.y]), l = [g(old.z)] }\n",
			vars:    []string{"old.x", "old.y", "old.z"},
			renamed: "a = { k = f(new.x, [new.y]), l = [g(new.z)] }\n",
		},
		{
			src:     "a = f(\"${old.x}\", old.y + 1)\n",
			vars:    []string{"old.x", "old.y"},
			renamed: "a = f(\"${new.x}\", new.y + 1)\n",
		},
		{
			src:     "a = f()\n",
			vars:    nil,
			renamed: "a = f()\n",
		},
		{
			src:     "a = []\n",
			vars:    nil,
			renamed: "a = []\n",
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
			if diff := cmp.Diff(test.vars, got); diff != "" {
				t.Errorf("wrong variables\n%s", diff)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if diff := cmp.Diff(test.renamed, string(f.Bytes())); diff != "" {
				t.Errorf("wrong result after rename\n%s", diff)
			}
		})
	}
}

func TestExpressionAsFunctionCallExprAndTupleConsExpr(t *testing.T) {
	tests := []struct {
		src       string
		wantCall  bool
		wantTuple bool
	}{
		{src: "a = f(x)\n", wantCall: true},
		{src: "a = [x]\n", wantTuple: true},
		{src: "a = f(x) + 1\n"},
		{src: "a = [for x in y : x]\n"},
		{src: "a = { k = [x] }\n"},
		{src: "a = x\n"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			expr := f.Body().GetAttribute("a").Expr()
			if got := expr.AsFunctionCallExpr() != nil; got != test.wantCall {
				t.Errorf("AsFunctionCallExpr returned a call: %t, want %t", got, test.wantCall)
			}
			if got := expr.AsTupleConsExpr() != nil; got != test.wantTuple {
				t.Errorf("AsTupleConsExpr returned a tuple: %t, want %t", got, test.wantTuple)
			}
		})
	}
}
