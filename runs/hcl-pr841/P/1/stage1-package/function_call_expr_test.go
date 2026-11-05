// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
)

func TestFunctionCallExprName(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{`a = f()`, "f"},
		{`a = join(" ", ["fedora", "fez"])`, "join"},
		{`a = provider::p::upper("fez")`, "provider::p::upper"},
		{"a = f(\n  1,\n)", "f"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			call := f.Body().GetAttribute("a").Expr().AsFunctionCallExpr()
			if call == nil {
				t.Fatalf("not a function call expression")
			}
			if got := call.Name(); got != test.want {
				t.Errorf("wrong name %q; want %q", got, test.want)
			}
		})
	}
}
