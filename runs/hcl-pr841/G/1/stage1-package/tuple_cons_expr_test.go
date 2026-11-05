// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
)

func TestTupleConsExprAsTupleConsExpr(t *testing.T) {
	tests := []struct {
		src  string
		want bool // whether the expression is a tuple
	}{
		{`a = ["fedora", "fez"]`, true},
		{`a = []`, true},
		{"a = [\n  \"fedora\", # hat\n  [\"fez\"],\n]", true},
		{`a = [for hat in hats : hat]`, false},
		{`a = ["fedora"][0]`, false},
		{`a = join(" ", ["fedora"])`, false},
		{`a = "fez"`, false},
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

			tuple := f.Body().GetAttribute("a").Expr().AsTupleConsExpr()
			if got := tuple != nil; got != test.want {
				t.Errorf("got a tuple: %t, want %t", got, test.want)
			}
		})
	}
}
