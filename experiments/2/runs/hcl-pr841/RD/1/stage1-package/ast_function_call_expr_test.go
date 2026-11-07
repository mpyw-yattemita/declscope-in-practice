// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
)

func TestFunctionCallExprName(t *testing.T) {
	tests := []struct {
		src  string // the function call assigned to a
		want string
	}{
		{`join(" ", ["fedora", "fez"])`, "join"},
		{`f()`, "f"},
		{`f(xs...)`, "f"},
		{`provider::aws::arn_parse(x)`, "provider::aws::arn_parse"},
		{"f(\n  1,\n)", "f"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = "+test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			call := f.Body().GetAttribute("a").Expr().AsFunctionCallExpr()
			if call == nil {
				t.Fatalf("not parsed as a function call expression")
			}
			if got := call.Name(); got != test.want {
				t.Errorf("wrong name %q, want %q", got, test.want)
			}
		})
	}
}
