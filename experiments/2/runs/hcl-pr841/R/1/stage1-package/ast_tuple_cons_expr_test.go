// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTupleConsExprParse(t *testing.T) {
	tests := map[string]struct {
		src       string
		wantExprs []string
	}{
		"empty":          {`a = []`, nil},
		"strings":        {`a = ["fedora", "fez"]`, []string{`"fedora"`, `"fez"`}},
		"trailing comma": {`a = [1, 2,]`, []string{"1", "2"}},
		"nested":         {`a = [[1], f(2), { k = 3 }]`, []string{"[1]", "f(2)", "{ k = 3 }"}},
		"multi-line":     {"a = [\n  # first\n  1,\n  2, // two\n]", []string{"1", "2"}},
		"traversals":     {`a = [x.y, z[0]]`, []string{"x.y", "z[0]"}},
		"unformatted":    {"a = [  1,2 , ]", []string{"1", "2"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			src := test.src + "\n"
			f, expr := parseTestExpr(t, src)

			tuple := expr.AsTupleConsExpr()
			if tuple == nil {
				t.Fatalf("expression is not a tuple-construct expression")
			}
			if expr.AsFunctionCallExpr() != nil || expr.AsObjectConsExpr() != nil {
				t.Errorf("tuple is also reported as another kind of expression")
			}
			if diff := cmp.Diff(test.wantExprs, exprsStrings(tuple.exprs)); diff != "" {
				t.Errorf("wrong elements (-want +got):\n%s", diff)
			}
			checkRoundTrip(t, f, src)
		})
	}
}
