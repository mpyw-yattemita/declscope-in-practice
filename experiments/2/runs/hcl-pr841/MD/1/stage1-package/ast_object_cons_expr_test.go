// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

func TestObjectConsExprItemFor(t *testing.T) {
	tests := map[string]struct {
		src       string
		key       string
		wantValue string // the value's source, or "" for no match
	}{
		"identifier":                    {`a = { hat = 1 }`, "hat", "1"},
		"quoted":                        {`a = { "hat" = 1 }`, "hat", "1"},
		"quoted with escaped template":  {`a = { "$${hat}" = 1 }`, "${hat}", "1"},
		"quoted with escaped directive": {`a = { "%%{hat}" = 1 }`, "%{hat}", "1"},
		"quoted with space":             {`a = { "a hat" = 1 }`, "a hat", "1"},
		"quoted with escapes":           {`a = { "say \"hi\"\\" = 1 }`, `say "hi"\`, "1"},
		"quoted empty":                  {`a = { "" = 1 }`, "", "1"},
		"colon":                         {`a = { hat: 1 }`, "hat", "1"},
		"second item":                   {`a = { cat = 1, hat = 2 }`, "hat", "2"},
		"first of duplicates":           {`a = { hat = 1, "hat" = 2 }`, "hat", "1"},
		"other key":                     {`a = { cat = 1 }`, "hat", ""},
		"empty object":                  {`a = {}`, "hat", ""},
		"parenthesized variable":        {`a = { (hat) = 1 }`, "hat", ""},
		"traversal":                     {`a = { hat.band = 1 }`, "hat.band", ""},
		"template":                      {`a = { "${hat}" = 1 }`, "hat", ""},
		"template matched literally":    {`a = { "${hat}" = 1 }`, "${hat}", ""},
		"escaped template not decoded":  {`a = { "$${hat}" = 1 }`, "$${hat}", ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			obj := parseObjectConsExprForTest(t, test.src+"\n")

			item := obj.ItemFor(test.key)
			value := obj.ValueExprFor(test.key)
			if test.wantValue == "" {
				if item != nil || value != nil {
					t.Fatalf("found an item for %q, want none", test.key)
				}
				return
			}
			if item == nil || value == nil {
				t.Fatalf("found no item for %q", test.key)
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor and ItemFor disagree")
			}
			got := strings.TrimSpace(string(value.Expr().BuildTokens(nil).Bytes()))
			if got != test.wantValue {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.wantValue)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	obj := parseObjectConsExprForTest(t, "a = { hat = 1, (cat) = 2, \"dog\" = 3 }\n")

	var got []string
	for _, item := range obj.Items() {
		got = append(got, strings.TrimSpace(string(item.KeyExpr().BuildTokens(nil).Bytes())))
	}
	want := []string{"hat", "(cat)", `"dog"`}
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("wrong keys\ngot:  %q\nwant: %q", got, want)
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := map[string]struct {
		key  string
		want string
	}{
		"identifier":       {"dog", `dog`},
		"identifier dash":  {"dog-1", `dog-1`},
		"space":            {"a hat", `"a hat"`},
		"traversal":        {"hat.band", `"hat.band"`},
		"leading digit":    {"1hat", `"1hat"`},
		"quotes":           {`say "hi"`, `"say \"hi\""`},
		"backslash":        {`a\b`, `"a\\b"`},
		"template":         {"${hat}", `"$${hat}"`},
		"directive":        {"%{hat}", `"%%{hat}"`},
		"lone dollar":      {"$hat", `"$hat"`},
		"empty":            {"", `""`},
		"newline":          {"a\nb", `"a\nb"`},
		"for keyword":      {"for", `"for"`},
		"unicode":          {"日本", `日本`},
		"unicode with gap": {"日 本", `"日 本"`},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseFileObjectConsExprForTest(t, "a = {}\n")

			key, _ := obj.SetItemValue(test.key, cty.True)
			got := strings.TrimSpace(string(key.BuildTokens(nil).Bytes()))
			if got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			if obj.ItemFor(test.key) == nil {
				t.Errorf("ItemFor(%q) does not find the key just set", test.key)
			}

			// The key must evaluate back to the original string.
			src := f.Bytes()
			native, diags := hclsyntax.ParseConfig(src, "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("result is not valid HCL: %s\n%s", diags, src)
			}
			attrs, _ := native.Body.JustAttributes()
			val, diags := attrs["a"].Expr.Value(nil)
			if diags.HasErrors() {
				t.Fatalf("result does not evaluate: %s\n%s", diags, src)
			}
			if !val.Type().IsObjectType() || !val.Type().HasAttribute(test.key) {
				t.Errorf("result has no key %q: %#v\n%s", test.key, val, src)
			}
			assertObjectConsExprFormatted(t, src)
		})
	}
}

func TestObjectConsExprSetItemReplaces(t *testing.T) {
	tests := map[string]struct {
		src  string
		key  string
		want string
	}{
		"identifier": {
			"a = { hat = \"derby\", cat = \"calico\" }\n",
			"hat",
			"a = { hat = \"bowler\", cat = \"calico\" }\n",
		},
		"quoted key stays quoted": {
			"a = { \"hat\" = \"derby\" }\n",
			"hat",
			"a = { \"hat\" = \"bowler\" }\n",
		},
		"escaped key stays as written": {
			"a = { \"$${hat}\" = \"derby\" }\n",
			"${hat}",
			"a = { \"$${hat}\" = \"bowler\" }\n",
		},
		"multi-line value": {
			"a = {\n  hat = {\n    b = 1\n  }\n  cat = 2\n}\n",
			"hat",
			"a = {\n  hat = \"bowler\"\n  cat = 2\n}\n",
		},
		"variable key is not replaced": {
			"a = { (hat) = \"derby\" }\n",
			"hat",
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseFileObjectConsExprForTest(t, test.src)
			before := len(obj.Items())

			_, value := obj.SetItemValue(test.key, cty.StringVal("bowler"))
			if value != obj.ValueExprFor(test.key) {
				t.Errorf("SetItemValue did not return the value that ValueExprFor finds")
			}
			got := string(f.Bytes())
			if got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if !strings.Contains(name, "not replaced") && len(obj.Items()) != before {
				t.Errorf("item count changed from %d to %d", before, len(obj.Items()))
			}
			assertObjectConsExprFormatted(t, f.Bytes())
		})
	}
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := map[string]struct {
		src  string
		keys []string
		want string
	}{
		"single line": {
			src:  "a = { hat = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		"empty": {
			src:  "a = {}\n",
			keys: []string{"dog"},
			want: "a = { dog = \"bowler\" }\n",
		},
		"empty, two items": {
			src:  "a = {}\n",
			keys: []string{"dog", "cat"},
			want: "a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		"single line, two items": {
			src:  "a = { hat = \"derby\" }\n",
			keys: []string{"dog", "cat"},
			want: "a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		"single line, trailing comma": {
			src:  "a = { hat = \"derby\", }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		"single line, inline comment before brace": {
			src:  "a = { hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		"single line, comma and inline comment before brace": {
			src:  "a = { hat = \"derby\", /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		"single line, line comment after object": {
			src:  "a = { hat = \"derby\" } # about a\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" } # about a\n",
		},
		"single line, multi-line value": {
			src:  "a = { hat = {\n  b = 1\n} }\n",
			keys: []string{"dog"},
			want: "a = { hat = {\n  b = 1\n}, dog = \"bowler\" }\n",
		},
		"multi-line": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"multi-line, two items": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi-line empty": {
			src:  "a = {\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  dog = \"bowler\"\n}\n",
		},
		"multi-line, commas": {
			src:  "a = {\n  hat = \"derby\",\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		"multi-line, brace on last line": {
			src:  "a = {\n  hat = \"derby\", cat = \"calico\" }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		"multi-line, brace on last line, two items": {
			src:  "a = {\n  hat = \"derby\" }\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi-line, brace on last line after comma": {
			src:  "a = {\n  hat = \"derby\", }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		"multi-line, inline comment before brace": {
			src:  "a = {\n  hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"multi-line, line comment after last item": {
			src:  "a = {\n  hat = \"derby\" # a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" # a hat\n  dog = \"bowler\"\n}\n",
		},
		"multi-line, comma and line comment after last item": {
			src:  "a = {\n  hat = \"derby\", # a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", # a hat\n  dog = \"bowler\"\n}\n",
		},
		"multi-line, comments before brace": {
			src:  "a = {\n  hat = \"derby\"\n\n  # more to come\n  // and more\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n\n  # more to come\n  // and more\n}\n",
		},
		"multi-line empty, comment": {
			src:  "a = { # nothing yet\n}\n",
			keys: []string{"dog"},
			want: "a = { # nothing yet\n  dog = \"bowler\"\n}\n",
		},
		"multi-line, nested in block": {
			src:  "b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			keys: []string{"dog"},
			want: "b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseFileObjectConsExprForTest(t, test.src)

			for _, key := range test.keys {
				gotKey, gotValue := obj.SetItemValue(key, cty.StringVal("bowler"))
				item := obj.ItemFor(key)
				if item == nil {
					t.Fatalf("ItemFor(%q) does not find the new item", key)
				}
				if gotKey != item.KeyExpr() || gotValue != item.ValueExpr() {
					t.Errorf("SetItemValue(%q) did not return the new item's key and value", key)
				}
			}

			got := string(f.Bytes())
			if got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if _, diags := hclsyntax.ParseConfig(f.Bytes(), "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags)
			}
			assertObjectConsExprFormatted(t, f.Bytes())
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, obj := parseFileObjectConsExprForTest(t, "a = {}\n")

	obj.SetItemRaw("raw", Tokens{{Type: hclsyntax.TokenIdent, Bytes: []byte("whatever")}})
	obj.SetItemValue("value", cty.NumberIntVal(1))
	obj.SetItemTraversal("traversal", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "hat"},
	})
	obj.SetItem("expr", NewExpressionLiteral(cty.StringVal("derby")))

	want := "a = { raw = whatever, value = 1, traversal = var.hat, expr = \"derby\" }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	vars := obj.ValueExprFor("traversal").Expr().Variables()
	if len(vars) != 1 || strings.TrimSpace(string(vars[0].BuildTokens(nil).Bytes())) != "var.hat" {
		t.Errorf("wrong variables in the traversal value: %#v", vars)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	t.Run("expression attached by the parser", func(t *testing.T) {
		f, obj := parseFileObjectConsExprForTest(t, "a = { hat = \"derby\" }\nb = c.d\n")
		expr := f.Body().GetAttribute("b").Expr()
		assertObjectConsExprPanics(t, func() { obj.SetItem("dog", expr) })
	})
	t.Run("item value attached by the parser", func(t *testing.T) {
		_, obj := parseFileObjectConsExprForTest(t, "a = { hat = c.d }\n")
		expr := obj.ValueExprFor("hat").Expr()
		assertObjectConsExprPanics(t, func() { obj.SetItem("dog", expr) })
	})
	t.Run("expression attached by SetItem", func(t *testing.T) {
		_, obj := parseFileObjectConsExprForTest(t, "a = {}\n")
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "c"}})
		obj.SetItem("hat", expr)
		assertObjectConsExprPanics(t, func() { obj.SetItem("dog", expr) })
		assertObjectConsExprPanics(t, func() { obj.SetItem("hat", expr) })
	})
	t.Run("replaced expression may be reused", func(t *testing.T) {
		_, obj := parseFileObjectConsExprForTest(t, "a = {}\n")
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "c"}})
		obj.SetItem("hat", expr)
		obj.SetItemValue("hat", cty.True)
		obj.SetItem("dog", expr)
	})
}

// TestObjectConsExprSetItemExample checks the example from the API's
// proposal: the key is kept as written and the value is replaced.
func TestObjectConsExprSetItemExample(t *testing.T) {
	f, obj := parseFileObjectConsExprForTest(t, "a = {\n    hat = \"derby\", (cat) = \"calico\" }\n")

	key, value := obj.SetItemRaw("hat", TokensForValue(cty.StringVal("bowler")))
	if got := strings.TrimSpace(string(key.BuildTokens(nil).Bytes())); got != "hat" {
		t.Errorf("wrong key: %s", got)
	}
	if got := strings.TrimSpace(string(value.BuildTokens(nil).Bytes())); got != `"bowler"` {
		t.Errorf("wrong value: %s", got)
	}
	want := "a = {\nhat = \"bowler\", (cat) = \"calico\" }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func parseFileObjectConsExprForTest(t *testing.T, src string) (*File, *ObjectConsExpr) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("invalid source: %s", diags)
	}
	body := f.Body()
	if blocks := body.Blocks(); len(blocks) > 0 {
		body = blocks[0].Body()
	}
	obj := body.GetAttribute("a").Expr().AsObjectConsExpr()
	if obj == nil {
		t.Fatalf("a is not an object-construct expression")
	}
	return f, obj
}

func parseObjectConsExprForTest(t *testing.T, src string) *ObjectConsExpr {
	t.Helper()
	_, obj := parseFileObjectConsExprForTest(t, src)
	return obj
}

func assertObjectConsExprFormatted(t *testing.T, src []byte) {
	t.Helper()
	if formatted := Format(src); string(formatted) != string(src) {
		t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", src, formatted)
	}
}

func assertObjectConsExprPanics(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		t.Helper()
		if recover() == nil {
			t.Errorf("did not panic")
		}
	}()
	f()
}
