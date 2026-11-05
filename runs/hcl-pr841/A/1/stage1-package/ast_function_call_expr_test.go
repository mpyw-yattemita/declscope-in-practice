// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"slices"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// parseAttrExprForTest parses src and returns the expression of its
// attribute "a".
func parseAttrExprForTest(t *testing.T, src string) (*File, *Expression) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	return f, f.Body().GetAttribute("a").Expr()
}

func TestFunctionCallExprName(t *testing.T) {
	tests := map[string]string{
		"a = join(\" \", [\"fedora\", \"fez\"])\n": "join",
		"a = f()\n":               "f",
		"a = provider::p::f(1)\n": "provider::p::f",
		"a = f(\n  1,\n  2,\n)\n": "f",
	}
	for src, want := range tests {
		t.Run(src, func(t *testing.T) {
			_, expr := parseAttrExprForTest(t, src)
			call := expr.AsFunctionCallExpr()
			if call == nil {
				t.Fatalf("attribute a is not a function call expression")
			}
			if got := call.Name(); got != want {
				t.Errorf("wrong name\ngot:  %q\nwant: %q", got, want)
			}
		})
	}
}

func TestExpressionAsFunctionCallAndTupleCons(t *testing.T) {
	tests := map[string]struct {
		src       string
		wantCall  bool
		wantTuple bool
	}{
		"function call":          {src: "a = f(1)\n", wantCall: true},
		"tuple":                  {src: "a = [1]\n", wantTuple: true},
		"empty tuple":            {src: "a = []\n", wantTuple: true},
		"object":                 {src: "a = { b = f(1) }\n"},
		"traversal":              {src: "a = b.c\n"},
		"call in binary op":      {src: "a = f(1) + 1\n"},
		"tuple in parentheses":   {src: "a = ([1])\n"},
		"for expression":         {src: "a = [for x in y : x]\n"},
		"index on function call": {src: "a = f(1)[0]\n"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, expr := parseAttrExprForTest(t, test.src)
			if got := expr.AsFunctionCallExpr() != nil; got != test.wantCall {
				t.Errorf("AsFunctionCallExpr() != nil is %t, want %t", got, test.wantCall)
			}
			if got := expr.AsTupleConsExpr() != nil; got != test.wantTuple {
				t.Errorf("AsTupleConsExpr() != nil is %t, want %t", got, test.wantTuple)
			}
		})
	}
}

func TestFunctionCallAndTupleConsVariables(t *testing.T) {
	tests := map[string]struct {
		src string
		// want is the source of each traversal that Variables returns,
		// in order.
		want []string
		// renamed is the source after renaming the prefix "old" to "new".
		renamed string
	}{
		"function arguments in source order": {
			src:     "a = f(old.x, other.y, old.z)\n",
			want:    []string{"old.x", "other.y", "old.z"},
			renamed: "a = f(new.x, other.y, new.z)\n",
		},
		"tuple elements in source order": {
			src:     "a = [old.x, other.y, old.z]\n",
			want:    []string{"old.x", "other.y", "old.z"},
			renamed: "a = [new.x, other.y, new.z]\n",
		},
		"nested calls and tuples": {
			src:     "a = f([old.x, g(old.y)], old.z)\n",
			want:    []string{"old.x", "old.y", "old.z"},
			renamed: "a = f([new.x, g(new.y)], new.z)\n",
		},
		"object in function argument": {
			src:     "a = f({ k = old.x }, old.y)\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = f({ k = new.x }, new.y)\n",
		},
		"function call in object item": {
			src:     "a = { k = f(old.x) }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = f(new.x) }\n",
		},
		"tuple in object item": {
			src:     "a = { k = [old.x] }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = [new.x] }\n",
		},
		"expanded final argument": {
			src:     "a = f(old.x, old.y...)\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = f(new.x, new.y...)\n",
		},
		"multi-line tuple with comments": {
			src:     "a = [\n  old.x, # c\n  # d\n  old.y,\n]\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = [\n  new.x, # c\n  # d\n  new.y,\n]\n",
		},
		"argument with operators": {
			src:     "a = f(old.x + old.y)\n",
			want:    []string{"old.x", "old.y"},
			renamed: "a = f(new.x + new.y)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Variables must agree with hclsyntax on the same source.
			nf, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags)
			}
			nativeAttr := nf.Body.(*hclsyntax.Body).Attributes["a"]
			if nativeCount := len(nativeAttr.Expr.Variables()); nativeCount != len(test.want) {
				t.Fatalf("test case is wrong: hclsyntax reports %d traversals, but want has %d", nativeCount, len(test.want))
			}

			f, expr := parseAttrExprForTest(t, test.src)

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if !slices.Equal(got, test.want) {
				t.Errorf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

func TestFunctionCallAndTupleConsRoundTrip(t *testing.T) {
	tests := []string{
		"a = join(\" \", [\"fedora\", \"fez\"])\n",
		"a = f()\n",
		"a = f( )\n",
		"a = [ ]\n",
		"a = provider::p::f(1)\n",
		"a = f(1, 2,)\n",
		"a = f(x...)\n",
		"a = f(\n  1, # one\n  /* two */ 2,\n)\n",
		"a = [\n  1,\n  # comment\n  [2, [3]],\n  { b = f(4) },\n]\n",
		"a = f(<<EOT\nhello\nEOT\n)\n",
		"a = [for x in y : f(x)]\n",
		"block \"label\" {\n  a = f([1], { b = [2] })\n}\n",
	}
	for _, src := range tests {
		t.Run(src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			// The tree must hold every token unchanged.
			if got := string(f.BuildTokens(nil).Bytes()); got != src {
				t.Errorf("wrong tokens\ngot:  %q\nwant: %q", got, src)
			}
			// File.Bytes formats its result, so it can only reproduce
			// source that is already formatted.
			if string(Format([]byte(src))) != src {
				return
			}
			if got := string(f.Bytes()); got != src {
				t.Errorf("wrong result\ngot:  %q\nwant: %q", got, src)
			}
		})
	}
}
