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

func TestExpressionVariablesAndRename(t *testing.T) {
	tests := map[string]struct {
		src string
		// want lists the traversals that Variables returns, in order.
		want []string
		// wantRenamed is src after renaming the prefix "old" to "new".
		wantRenamed string
	}{
		"plain traversal": {
			src:         "a = old.x\n",
			want:        []string{"old.x"},
			wantRenamed: "a = new.x\n",
		},
		"object item value": {
			src:         "a = { k = old.x }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { k = new.x }\n",
		},
		"object items in source order": {
			src:         "a = { k = old.x, j = other.y, (old.z) = old.w }\n",
			want:        []string{"old.x", "other.y", "old.z", "old.w"},
			wantRenamed: "a = { k = new.x, j = other.y, (new.z) = new.w }\n",
		},
		"nested object": {
			src:         "a = { k = { j = old.x } }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { k = { j = new.x } }\n",
		},
		"tuple item value": {
			src:         "a = { k = [old.x] }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { k = [new.x] }\n",
		},
		"function call item value": {
			src:         "a = { k = f(old.x) }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { k = f(new.x) }\n",
		},
		"function call arguments": {
			src:         "a = f(old.x, other.y, old.z)\n",
			want:        []string{"old.x", "other.y", "old.z"},
			wantRenamed: "a = f(new.x, other.y, new.z)\n",
		},
		"tuple elements": {
			src:         "a = [old.x, other.y, old.z]\n",
			want:        []string{"old.x", "other.y", "old.z"},
			wantRenamed: "a = [new.x, other.y, new.z]\n",
		},
		"nested calls, tuples and objects in source order": {
			src:         "a = f(old.a, [old.b, g(old.c)], { k = [old.d] }, old.e...)\n",
			want:        []string{"old.a", "old.b", "old.c", "old.d", "old.e"},
			wantRenamed: "a = f(new.a, [new.b, g(new.c)], { k = [new.d] }, new.e...)\n",
		},
		"function call in object key": {
			src:         "a = { (f(old.x)) = [old.y] }\n",
			want:        []string{"old.x", "old.y"},
			wantRenamed: "a = { (f(new.x)) = [new.y] }\n",
		},
		"operation in function argument": {
			src:         "a = f(old.x + old.y)\n",
			want:        []string{"old.x", "old.y"},
			wantRenamed: "a = f(new.x + new.y)\n",
		},
		"naked identifier key": {
			src:         "a = { old = 1 }\n",
			want:        nil,
			wantRenamed: "a = { old = 1 }\n",
		},
		"naked identifier key with traversal value": {
			src:         "a = { old = old.x }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { old = new.x }\n",
		},
		"parenthesized key": {
			src:         "a = { (old.x) = 1 }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { (new.x) = 1 }\n",
		},
		"parenthesized single identifier key": {
			src:         "a = { (old) = 1 }\n",
			want:        []string{"old"},
			wantRenamed: "a = { (new) = 1 }\n",
		},
		"quoted key": {
			src:         "a = { \"k\" = old.x }\n",
			want:        []string{"old.x"},
			wantRenamed: "a = { \"k\" = new.x }\n",
		},
		"quoted key without traversal": {
			src:         "a = { \"old\" = 1 }\n",
			want:        nil,
			wantRenamed: "a = { \"old\" = 1 }\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
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
				t.Errorf("wrong variables (-want +got):\n%s", diff)
			}

			nativeFile, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags.Error())
			}
			nativeAttrs, _ := nativeFile.Body.JustAttributes()
			if got, want := len(got), len(nativeAttrs["a"].Expr.Variables()); got != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", got, want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.wantRenamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.wantRenamed)
			}
		})
	}
}

func TestExpressionAsFunctionCallAndTupleConsExpr(t *testing.T) {
	for _, src := range []string{
		`a = x`,
		`a = "fez"`,
		`a = { k = f(1) }`,
		`a = f(1)[0]`,
		`a = (f(1))`,
		`a = [for x in y : x]`,
		`a = f(1) + 1`,
	} {
		t.Run(src, func(t *testing.T) {
			_, expr := parseTestExpr(t, src+"\n")
			if expr.AsFunctionCallExpr() != nil {
				t.Errorf("AsFunctionCallExpr returned non-nil")
			}
			if expr.AsTupleConsExpr() != nil {
				t.Errorf("AsTupleConsExpr returned non-nil")
			}
		})
	}
	if NewExpressionRaw(Tokens{}).AsFunctionCallExpr() != nil {
		t.Errorf("AsFunctionCallExpr returned non-nil for a raw expression")
	}
}
