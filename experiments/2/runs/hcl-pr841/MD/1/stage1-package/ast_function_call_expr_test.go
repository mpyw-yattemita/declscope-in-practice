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
		{src: "a = join(\" \", [\"fedora\", \"fez\"])\n", want: "join"},
		{src: "a = timestamp()\n", want: "timestamp"},
		{src: "a = provider::example::upper(\"x\")\n", want: "provider::example::upper"},
		{src: "a = f(g(x))\n", want: "f"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			call := f.Body().GetAttribute("a").Expr().AsFunctionCallExpr()
			if call == nil {
				t.Fatal("not parsed as a function call expression")
			}
			if got := call.Name(); got != test.want {
				t.Errorf("wrong name\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}
