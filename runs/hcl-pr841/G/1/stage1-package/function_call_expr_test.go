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
		want string // the name, or "" when the expression is not a call
	}{
		{`a = join(" ", ["fedora", "fez"])`, "join"},
		{`a = f()`, "f"},
		{`a = f(x...)`, "f"},
		{`a = provider::time::rfc3339_parse("2026-11-05T00:00:00Z")`, "provider::time::rfc3339_parse"},
		{"a = f(\n  1,\n  2,\n)", "f"},
		{`a = f(1)[0]`, ""},
		{`a = (f(1))`, ""},
		{`a = ["fedora"]`, ""},
		{`a = "fez"`, ""},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			src := test.src + "\n"
			f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			if got := string(f.Bytes()); got != src {
				t.Errorf("source does not round-trip\ngot:  %q\nwant: %q", got, src)
			}

			call := f.Body().GetAttribute("a").Expr().AsFunctionCallExpr()
			if test.want == "" {
				if call != nil {
					t.Fatalf("got a call named %q, but want none", call.Name())
				}
				return
			}
			if call == nil {
				t.Fatalf("got no call")
			}
			if got := call.Name(); got != test.want {
				t.Errorf("wrong name\ngot:  %q\nwant: %q", got, test.want)
			}
		})
	}
}
