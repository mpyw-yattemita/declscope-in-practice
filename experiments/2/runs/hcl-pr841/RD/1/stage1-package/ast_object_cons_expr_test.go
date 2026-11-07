// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0


package hclwrite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

func TestObjectConsExprItemFor(t *testing.T) {
	tests := []struct {
		src   string // the object assigned to a
		key   string
		found bool
	}{
		{`{ hat = 1 }`, "hat", true},
		{`{ "hat" = 1 }`, "hat", true},
		{`{ "$${hat}" = 1 }`, "${hat}", true},
		{`{ "%%{hat}" = 1 }`, "%{hat}", true},
		{`{ "a hat" = 1 }`, "a hat", true},
		{`{ "say \"hi\"" = 1 }`, `say "hi"`, true},
		{`{ "a\\b" = 1 }`, `a\b`, true},
		{`{ "" = 1 }`, "", true},
		{`{ null = 1 }`, "null", true},
		{`{ hat: 1 }`, "hat", true},
		{`{ cat = 1, hat = 1 }`, "hat", true},
		{`{ (hat) = 1 }`, "hat", false},
		{`{ ("hat") = 1 }`, "hat", false},
		{`{ hat.band = 1 }`, "hat", false},
		{`{ hat.band = 1 }`, "hat.band", false},
		{`{ "${hat}" = 1 }`, "hat", false},
		{`{ "${hat}" = 1 }`, "${hat}", false},
		{`{ "$${hat}" = 1 }`, "$${hat}", false},
		{`{ hat = 1 }`, "Hat", false},
		{`{ cat = { hat = 1 } }`, "hat", false},
		{`{}`, "hat", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %q", test.src, test.key), func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = "+test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			item := object.ItemFor(test.key)
			value := object.ValueExprFor(test.key)
			if found := item != nil; found != test.found {
				t.Fatalf("ItemFor found an item: %v, want %v", found, test.found)
			}
			if !test.found {
				if value != nil {
					t.Errorf("ValueExprFor returned %q, want nil", value.BuildTokens(nil).Bytes())
				}
				return
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned another value than ItemFor")
			}
			if got := strings.TrimSpace(string(value.Expr().BuildTokens(nil).Bytes())); got != "1" {
				t.Errorf("wrong value %q", got)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	src := "a = { hat = \"derby\", (cat) = \"calico\", \"dog\" = 3 }\n"
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	var got [][2]string
	for _, item := range object.Items() {
		got = append(got, [2]string{
			strings.TrimSpace(string(item.KeyExpr().BuildTokens(nil).Bytes())),
			strings.TrimSpace(string(item.ValueExpr().Expr().BuildTokens(nil).Bytes())),
		})
	}
	want := [][2]string{
		{"hat", `"derby"`},
		{"(cat)", `"calico"`},
		{`"dog"`, "3"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("wrong items (-want +got):\n%s", diff)
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string // the key as written
	}{
		{"hat", "hat"},
		{"hat_band", "hat_band"},
		{"hat-band", "hat-band"},
		{"null", "null"},
		{"for", `"for"`},
		{"if", "if"},
		{"a hat", `"a hat"`},
		{"hat.band", `"hat.band"`},
		{"1hat", `"1hat"`},
		{`say "hi"`, `"say \"hi\""`},
		{`a\b`, `"a\\b"`},
		{"${hat}", `"$${hat}"`},
		{"%{hat}", `"%%{hat}"`},
		{"$hat", `"$hat"`},
		{"", `""`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = { cat = 1 }\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, value := object.SetItemValue(test.key, cty.NumberIntVal(2))
			if got := strings.TrimSpace(string(key.BuildTokens(nil).Bytes())); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			if item := object.ItemFor(test.key); item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("ItemFor does not find the item just set")
			}

			// The key must evaluate back to the original string.
			native, diags := hclsyntax.ParseConfig(f.Bytes(), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("result is not valid HCL: %s\n%s", diags.Error(), f.Bytes())
			}
			val, diags := native.Body.(*hclsyntax.Body).Attributes["a"].Expr.Value(nil)
			if diags.HasErrors() {
				t.Fatalf("result does not evaluate: %s\n%s", diags.Error(), f.Bytes())
			}
			wantVal := cty.ObjectVal(map[string]cty.Value{
				"cat":    cty.NumberIntVal(1),
				test.key: cty.NumberIntVal(2),
			})
			if !val.RawEquals(wantVal) {
				t.Errorf("wrong value\ngot:  %#v\nwant: %#v", val, wantVal)
			}
		})
	}
}

func TestObjectConsExprSetItemExisting(t *testing.T) {
	tests := []struct {
		src  string
		key  string
		want string
	}{
		{
			"a = { hat = \"derby\" }\n",
			"hat",
			"a = { hat = \"bowler\" }\n",
		},
		{
			"a = { \"hat\" = \"derby\", cat = 1 }\n",
			"hat",
			"a = { \"hat\" = \"bowler\", cat = 1 }\n",
		},
		{
			"a = { \"$${hat}\" = \"derby\" }\n",
			"${hat}",
			"a = { \"$${hat}\" = \"bowler\" }\n",
		},
		{
			"a = {\n  hat = \"derby\" # felt\n  cat = 1\n}\n",
			"hat",
			"a = {\n  hat = \"bowler\" # felt\n  cat = 1\n}\n",
		},
		{
			// A key from a variable is not the one that is set.
			"a = { (hat) = \"derby\" }\n",
			"hat",
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, value := object.SetItemValue(test.key, cty.StringVal("bowler"))
			if item := object.ItemFor(test.key); item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("SetItemValue did not return the item that ItemFor finds")
			}
			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, diags := ParseConfig([]byte("a = { hat = 1 }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	object.SetItemRaw("hat", TokensForValue(cty.StringVal("derby")))
	object.SetItemValue("cat", cty.ListVal([]cty.Value{cty.StringVal("calico")}))
	_, value := object.SetItemTraversal("dog", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "dog"},
	})

	want := "a = { hat = \"derby\", cat = [\"calico\"], dog = var.dog }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	var got []string
	for _, traversal := range f.Body().GetAttribute("a").Expr().Variables() {
		got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
	}
	if diff := cmp.Diff([]string{"var.dog"}, got); diff != "" {
		t.Errorf("wrong traversals (-want +got):\n%s", diff)
	}
	if got := value.Expr().Variables(); len(got) != 1 {
		t.Errorf("the value has %d traversals, want 1", len(got))
	}
}

func TestObjectConsExprSetItemAttached(t *testing.T) {
	tests := map[string]func(object *ObjectConsExpr) *Expression{
		"attached by the parser": func(object *ObjectConsExpr) *Expression {
			return object.ValueExprFor("hat").Expr()
		},
		"attached by SetItem": func(object *ObjectConsExpr) *Expression {
			expr := NewExpressionLiteral(cty.StringVal("bowler"))
			object.SetItem("cat", expr)
			return expr
		},
		"attached to an attribute": func(object *ObjectConsExpr) *Expression {
			f, _ := ParseConfig([]byte("b = 1\n"), "", hcl.InitialPos)
			return f.Body().GetAttribute("b").Expr()
		},
	}

	for name, attached := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = { hat = \"derby\" }\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
			expr := attached(object)
			before := string(f.Bytes())

			for _, key := range []string{"hat", "dog"} {
				func() {
					defer func() {
						if recover() == nil {
							t.Errorf("SetItem(%q) did not panic", key)
						}
					}()
					object.SetItem(key, expr)
				}()
			}
			if got := string(f.Bytes()); got != before {
				t.Errorf("SetItem changed the object before panicking\ngot:  %s\nwant: %s", got, before)
			}
		})
	}

	t.Run("replaced", func(t *testing.T) {
		f, diags := ParseConfig([]byte("a = { hat = \"derby\" }\n"), "", hcl.InitialPos)
		if diags.HasErrors() {
			t.Fatalf("unexpected diagnostics: %s", diags.Error())
		}
		object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

		// An expression that SetItem replaced is no longer in the tree, so it
		// can be used again.
		old := object.ValueExprFor("hat").Expr()
		object.SetItemValue("hat", cty.StringVal("bowler"))
		object.SetItem("dog", old)

		want := "a = { hat = \"bowler\", dog = \"derby\" }\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
		}
	})
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := []struct {
		name string
		src  string
		keys []string // the keys added, each set to "bowler"
		want string
	}{
		{
			"single line",
			`a = { hat = "derby" }`,
			[]string{"dog"},
			`a = { hat = "derby", dog = "bowler" }`,
		},
		{
			"empty",
			`a = {}`,
			[]string{"dog"},
			`a = { dog = "bowler" }`,
		},
		{
			"empty, two items",
			`a = {}`,
			[]string{"dog", "cat"},
			`a = { dog = "bowler", cat = "bowler" }`,
		},
		{
			"single line, two items",
			`a = { hat = "derby" }`,
			[]string{"dog", "cat"},
			`a = { hat = "derby", dog = "bowler", cat = "bowler" }`,
		},
		{
			"single line, trailing comma",
			`a = { hat = "derby", }`,
			[]string{"dog"},
			`a = { hat = "derby", dog = "bowler" }`,
		},
		{
			"single line, colon",
			`a = { hat: "derby" }`,
			[]string{"dog"},
			`a = { hat : "derby", dog = "bowler" }`,
		},
		{
			"single line, inline comment before the brace",
			`a = { hat = "derby" /* end of a */ }`,
			[]string{"dog"},
			`a = { hat = "derby", dog = "bowler" /* end of a */ }`,
		},
		{
			"single line, nested multi-line value",
			"a = { hat = {\n  band = \"silk\"\n} }",
			[]string{"dog"},
			"a = { hat = {\n  band = \"silk\"\n}, dog = \"bowler\" }",
		},
		{
			"multi-line",
			"a = {\n  hat = \"derby\"\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, two items",
			"a = {\n  hat = \"derby\"\n}",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}",
		},
		{
			"multi-line, empty",
			"a = {\n}",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, longer key",
			"a = {\n  hat = \"derby\"\n}",
			[]string{"bowler"},
			"a = {\n  hat    = \"derby\"\n  bowler = \"bowler\"\n}",
		},
		{
			"multi-line, comma",
			"a = {\n  hat = \"derby\",\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, line comment",
			"a = {\n  hat = \"derby\" # felt\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\" # felt\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, comma and line comment",
			"a = {\n  hat = \"derby\", // felt\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", // felt\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, comments before the brace",
			"a = {\n  hat = \"derby\"\n  # more hats\n  /* to come */\n}",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  # more hats\n  /* to come */\n}",
		},
		{
			"multi-line, empty with comment",
			"a = {\n  # no hats\n}",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n  # no hats\n}",
		},
		{
			"multi-line, comment after the brace",
			"a = { # hats\n}",
			[]string{"dog"},
			"a = { # hats\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, brace after the last item",
			"a = {\n  hat = \"derby\" }",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, brace after the last item, two items",
			"a = {\n  hat = \"derby\" }",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}",
		},
		{
			"multi-line, comma and brace after the last item",
			"a = {\n  hat = \"derby\", }",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, first item on the brace's line",
			"a = { hat = \"derby\"\n  cat = \"calico\" }",
			[]string{"dog"},
			"a = { hat = \"derby\"\n  cat = \"calico\"\n  dog = \"bowler\"\n}",
		},
		{
			"multi-line, inline comment before the brace",
			"a = {\n  hat = \"derby\" /* end of a */ }",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }",
		},
		{
			"multi-line, nested",
			"a = {\n  hat = {\n    band = \"silk\"\n  }\n}",
			[]string{"dog"},
			"a = {\n  hat = {\n    band = \"silk\"\n  }\n  dog = \"bowler\"\n}",
		},
		{
			"empty, keyword key",
			`a = {}`,
			[]string{"for"},
			`a = { "for" = "bowler" }`,
		},
		{
			"nested object set",
			"a = { hat = {} }",
			[]string{"hat"},
			`a = { hat = "bowler" }`,
		},
	}

	bowler := TokensForValue(cty.StringVal("bowler"))
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			for _, key := range test.keys {
				object.SetItemRaw(key, bowler)
			}

			got := string(f.Bytes())
			want := test.want + "\n"
			if got != want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
			}
			if formatted := string(Format([]byte(got))); formatted != got {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := hclsyntax.ParseConfig([]byte(got), "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags.Error())
			}
			for _, key := range test.keys {
				if object.ItemFor(key) == nil {
					t.Errorf("ItemFor(%q) does not find the item just set", key)
				}
			}
		})
	}
}
