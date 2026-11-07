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
	tests := map[string]struct {
		src         string
		wantVars    []string
		wantRenamed string
	}{
		"value": {
			src:         "a = { k = old.x }\n",
			wantVars:    []string{"old.x"},
			wantRenamed: "a = { k = new.x }\n",
		},
		"several items in source order": {
			src:         "a = { k = old.x, j = other.y, l = old.z }\n",
			wantVars:    []string{"old.x", "other.y", "old.z"},
			wantRenamed: "a = { k = new.x, j = other.y, l = new.z }\n",
		},
		"nested object": {
			src:         "a = { k = { j = old.x } }\n",
			wantVars:    []string{"old.x"},
			wantRenamed: "a = { k = { j = new.x } }\n",
		},
		"tuple value": {
			src:         "a = { k = [old.x] }\n",
			wantVars:    []string{"old.x"},
			wantRenamed: "a = { k = [new.x] }\n",
		},
		"function call value": {
			src:         "a = { k = f(old.x) }\n",
			wantVars:    []string{"old.x"},
			wantRenamed: "a = { k = f(new.x) }\n",
		},
		"naked identifier key": {
			src:         "a = { old = 1 }\n",
			wantVars:    nil,
			wantRenamed: "a = { old = 1 }\n",
		},
		"naked identifier key with value": {
			src:         "a = { old = old.x }\n",
			wantVars:    []string{"old.x"},
			wantRenamed: "a = { old = new.x }\n",
		},
		"parenthesized key": {
			src:         "a = { (old.x) = 1 }\n",
			wantVars:    []string{"old.x"},
			wantRenamed: "a = { (new.x) = 1 }\n",
		},
		"parenthesized key and value": {
			src:         "a = { (old.x) = old.y }\n",
			wantVars:    []string{"old.x", "old.y"},
			wantRenamed: "a = { (new.x) = new.y }\n",
		},
		"quoted key": {
			src:         "a = { \"k\" = old.x }\n",
			wantVars:    []string{"old.x"},
			wantRenamed: "a = { \"k\" = new.x }\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var gotVars []string
			for _, traversal := range expr.Variables() {
				gotVars = append(gotVars, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if diff := cmp.Diff(test.wantVars, gotVars); diff != "" {
				t.Errorf("wrong variables\n%s", diff)
			}

			nativeFile, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags)
			}
			nativeExpr := nativeFile.Body.(*hclsyntax.Body).Attributes["a"].Expr
			if got, want := len(gotVars), len(nativeExpr.Variables()); got != want {
				t.Errorf("got %d variables, but hclsyntax reports %d", got, want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if diff := cmp.Diff(test.wantRenamed, string(f.Bytes())); diff != "" {
				t.Errorf("wrong result after rename\n%s", diff)
			}
		})
	}
}
