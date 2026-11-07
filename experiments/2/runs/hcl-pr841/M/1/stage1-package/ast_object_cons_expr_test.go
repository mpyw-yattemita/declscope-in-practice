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

// parseObjectAttr parses src and returns the file and the object-construct
// expression of its attribute "a".
func parseObjectAttr(t *testing.T, src string) (*File, *ObjectConsExpr) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	if object == nil {
		t.Fatalf("attribute a is not an object")
	}
	return f, object
}

// tokensText returns the source text of tokens without surrounding spaces.
func tokensText(tokens Tokens) string {
	return strings.TrimSpace(string(tokens.Bytes()))
}

// evalObjectAttr evaluates attribute "a" of src with hclsyntax.
func evalObjectAttr(t *testing.T, src []byte) cty.Value {
	t.Helper()
	f, diags := hclsyntax.ParseConfig(src, "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("result is not valid HCL: %s\n%s", diags.Error(), src)
	}
	val, diags := f.Body.(*hclsyntax.Body).Attributes["a"].Expr.Value(nil)
	if diags.HasErrors() {
		t.Fatalf("can't evaluate result: %s\n%s", diags.Error(), src)
	}
	return val
}

func TestObjectConsExprItemFor(t *testing.T) {
	tests := []struct {
		src  string
		key  string
		want string // value as source text, or "" if not found
	}{
		{`a = { hat = 1 }`, "hat", "1"},
		{`a = { "hat" = 1 }`, "hat", "1"},
		{`a = { hat: 1 }`, "hat", "1"},
		{`a = { "a hat" = 1 }`, "a hat", "1"},
		{`a = { "$${hat}" = 1 }`, "${hat}", "1"},
		{`a = { "%%{hat}" = 1 }`, "%{hat}", "1"},
		{`a = { "say \"hi\"" = 1 }`, `say "hi"`, "1"},
		{`a = { "a\\b" = 1 }`, `a\b`, "1"},
		{`a = { "" = 1 }`, "", "1"},
		{`a = { true = 1 }`, "true", "1"},
		{`a = { cat = 1, hat = 2 }`, "hat", "2"},
		{`a = { hat = { cat = 1 } }`, "hat", "{ cat = 1 }"},
		{`a = { hat = 1 }`, "cat", ""},
		{`a = { hat = 1 }`, `"hat"`, ""},
		{`a = { "$${hat}" = 1 }`, "$${hat}", ""},
		{`a = { (hat) = 1 }`, "hat", ""},
		{`a = { (hat) = 1 }`, "(hat)", ""},
		{`a = { hat.band = 1 }`, "hat", ""},
		{`a = { hat.band = 1 }`, "hat.band", ""},
		{`a = { "${hat}" = 1 }`, "hat", ""},
		{`a = { "${hat}" = 1 }`, "${hat}", ""},
		{`a = { "${"hat"}" = 1 }`, "hat", ""},
		{`a = {}`, "hat", ""},
	}

	for _, test := range tests {
		t.Run(test.src+" "+test.key, func(t *testing.T) {
			_, object := parseObjectAttr(t, test.src+"\n")

			item := object.ItemFor(test.key)
			value := object.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil {
					t.Errorf("ItemFor found an item, but want nil")
				}
				if value != nil {
					t.Errorf("ValueExprFor found a value, but want nil")
				}
				return
			}

			if item == nil {
				t.Fatalf("ItemFor found no item")
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a different value than ItemFor")
			}
			if got := tokensText(value.Expr().BuildTokens(nil)); got != test.want {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	_, object := parseObjectAttr(t, "a = {\n  hat = 1\n  \"cat\" = 2\n  (dog) = 3\n}\n")

	var keys, values []string
	for _, item := range object.Items() {
		keys = append(keys, tokensText(item.KeyExpr().BuildTokens(nil)))
		values = append(values, tokensText(item.ValueExpr().BuildTokens(nil)))
	}
	if got, want := strings.Join(keys, " "), `hat "cat" (dog)`; got != want {
		t.Errorf("wrong keys\ngot:  %s\nwant: %s", got, want)
	}
	if got, want := strings.Join(values, " "), "1 2 3"; got != want {
		t.Errorf("wrong values\ngot:  %s\nwant: %s", got, want)
	}

	_, empty := parseObjectAttr(t, "a = {}\n")
	if got := empty.Items(); len(got) != 0 {
		t.Errorf("got %d items in an empty object", len(got))
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string // key as source text
	}{
		{"dog", `dog`},
		{"dog_1", `dog_1`},
		{"dog-1", `dog-1`},
		{"true", `true`},
		{"for", `"for"`},
		{"a hat", `"a hat"`},
		{"hat.band", `"hat.band"`},
		{"1hat", `"1hat"`},
		{"", `""`},
		{`say "hi"`, `"say \"hi\""`},
		{`a\b`, `"a\\b"`},
		{"${hat}", `"$${hat}"`},
		{"%{hat}", `"%%{hat}"`},
		{"a\nb", `"a\nb"`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, object := parseObjectAttr(t, "a = { hat = \"derby\" }\n")

			key, value := object.SetItemValue(test.key, cty.StringVal("bowler"))
			if got := tokensText(key.BuildTokens(nil)); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}

			item := object.ItemFor(test.key)
			if item == nil {
				t.Fatalf("ItemFor can't find the key that was just set")
			}
			if item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("ItemFor found a different item than the one that was set")
			}

			got := evalObjectAttr(t, f.Bytes())
			want := cty.ObjectVal(map[string]cty.Value{
				"hat":    cty.StringVal("derby"),
				test.key: cty.StringVal("bowler"),
			})
			if !got.RawEquals(want) {
				t.Errorf("wrong value\ngot:  %#v\nwant: %#v", got, want)
			}

			// Setting the key again replaces the value instead of adding an item.
			key2, value2 := object.SetItemValue(test.key, cty.StringVal("fedora"))
			if key2 != key || value2 != value {
				t.Errorf("setting the key again did not return the same item")
			}
			if got := len(object.Items()); got != 2 {
				t.Errorf("got %d items after setting the key again, want 2", got)
			}
		})
	}
}

func TestObjectConsExprSetItemReplace(t *testing.T) {
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
			"a = { \"hat\" = \"derby\" }\n",
			"hat",
			"a = { \"hat\" = \"bowler\" }\n",
		},
		{
			"a = { hat: \"derby\" }\n",
			"hat",
			"a = { hat : \"bowler\" }\n",
		},
		{
			"a = { \"$${hat}\" = \"derby\" }\n",
			"${hat}",
			"a = { \"$${hat}\" = \"bowler\" }\n",
		},
		{
			"a = {\n  hat = \"derby\" # comment\n  cat = \"calico\"\n}\n",
			"hat",
			"a = {\n  hat = \"bowler\" # comment\n  cat = \"calico\"\n}\n",
		},
		{
			// A key from a variable is not replaced, so a new item is added.
			"a = { (hat) = \"derby\" }\n",
			"hat",
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, object := parseObjectAttr(t, test.src)
			key, value := object.SetItemValue(test.key, cty.StringVal("bowler"))

			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			item := object.ItemFor(test.key)
			if item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("SetItemValue did not return the item that ItemFor finds")
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, object := parseObjectAttr(t, "a = { hat = \"derby\" }\n")

	object.SetItemRaw("hat", Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte("upper")},
		{Type: hclsyntax.TokenOParen, Bytes: []byte("(")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("var")},
		{Type: hclsyntax.TokenDot, Bytes: []byte(".")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("hat")},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(")")},
	})
	object.SetItemTraversal("cat", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "cat"},
	})
	object.SetItemValue("dog", cty.NumberIntVal(2))

	want := "a = { hat = upper(var.hat), cat = var.cat, dog = 2 }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	// A traversal set with SetItemTraversal is a variable of the object.
	var vars []string
	for _, traversal := range f.Body().GetAttribute("a").Expr().Variables() {
		vars = append(vars, tokensText(traversal.BuildTokens(nil)))
	}
	if got, want := strings.Join(vars, " "), "var.cat"; got != want {
		t.Errorf("wrong variables\ngot:  %s\nwant: %s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	assertPanics := func(t *testing.T, f func()) {
		t.Helper()
		defer func() {
			t.Helper()
			if recover() == nil {
				t.Errorf("SetItem did not panic")
			}
		}()
		f()
	}

	src := "a = { hat = \"derby\", cat = { dog = 1 } }\nb = var.b\n"

	t.Run("attribute expression from the parser", func(t *testing.T) {
		f, object := parseObjectAttr(t, src)
		expr := f.Body().GetAttribute("b").Expr()
		assertPanics(t, func() { object.SetItem("hat", expr) })
		assertPanics(t, func() { object.SetItem("new", expr) })
	})

	t.Run("item value from the parser", func(t *testing.T) {
		_, object := parseObjectAttr(t, src)
		expr := object.ValueExprFor("hat").Expr()
		assertPanics(t, func() { object.SetItem("new", expr) })
	})

	t.Run("nested item value from the parser", func(t *testing.T) {
		_, object := parseObjectAttr(t, src)
		nested := object.ValueExprFor("cat").Expr().AsObjectConsExpr()
		expr := nested.ValueExprFor("dog").Expr()
		assertPanics(t, func() { object.SetItem("new", expr) })
	})

	t.Run("expression from an earlier SetItem", func(t *testing.T) {
		_, object := parseObjectAttr(t, src)
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "var"}})
		object.SetItem("hat", expr)
		assertPanics(t, func() { object.SetItem("new", expr) })
		assertPanics(t, func() { object.SetItem("hat", expr) })

		expr = NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "var"}})
		object.SetItem("new", expr)
		assertPanics(t, func() { object.SetItem("other", expr) })
	})

	t.Run("replaced expression can be reused", func(t *testing.T) {
		f, object := parseObjectAttr(t, src)
		old := object.ValueExprFor("hat").Expr()
		object.SetItemValue("hat", cty.StringVal("bowler"))
		object.SetItem("new", old)

		want := "a = { hat = \"bowler\", cat = { dog = 1 }, new = \"derby\" }\nb = var.b\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
		}
	})
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := []struct {
		name string
		src  string
		keys []string
		want string
	}{
		{
			"single line",
			"a = { hat = \"derby\" }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"single line, empty",
			"a = {}\n",
			[]string{"dog"},
			"a = { dog = \"bowler\" }\n",
		},
		{
			"single line, empty, two items",
			"a = {}\n",
			[]string{"dog", "cat"},
			"a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"single line, two items",
			"a = { hat = \"derby\" }\n",
			[]string{"dog", "cat"},
			"a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"single line, trailing comma",
			"a = { hat = \"derby\", }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"single line, inline comment before brace",
			"a = { hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		{
			"single line, nested multi-line value",
			"a = { hat = {\n  b = 1\n} }\n",
			[]string{"dog"},
			"a = { hat = {\n  b = 1\n}, dog = \"bowler\" }\n",
		},
		{
			"single line, multi-line key",
			"a = { (\n  hat\n  ) = 1 }\n",
			[]string{"dog"},
			"a = { (\n  hat\n) = 1, dog = \"bowler\" }\n",
		},
		{
			"multi-line",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, two items",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		{
			"multi-line, empty",
			"a = {\n}\n",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, key alignment",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"a_long_key"},
			"a = {\n  hat        = \"derby\"\n  a_long_key = \"bowler\"\n}\n",
		},
		{
			"multi-line, brace on the last item's line",
			"a = {\n  hat = \"derby\" }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, brace on the last item's line, after a comma",
			"a = {\n  hat = \"derby\", }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, from the example",
			"a = {\n    hat = \"derby\", (cat) = \"calico\" }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", (cat) = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, trailing comma",
			"a = {\n  hat = \"derby\",\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, line comment",
			"a = {\n  hat = \"derby\" # the hat\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\" # the hat\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, comma and line comment",
			"a = {\n  hat = \"derby\", // the hat\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", // the hat\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, comments before the brace",
			"a = {\n  hat = \"derby\"\n  # end of a\n  # really\n}\n",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n  # end of a\n  # really\n}\n",
		},
		{
			"multi-line, inline comment before the brace on its own line",
			"a = {\n  hat = \"derby\"\n  /* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"multi-line, inline comment before the brace on the last item's line",
			"a = {\n  hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"multi-line, nested multi-line value",
			"a = {\n  hat = {\n    b = 1\n  }\n}\n",
			[]string{"dog"},
			"a = {\n  hat = {\n    b = 1\n  }\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, in a block",
			"b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			[]string{"dog"},
			"b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			body := f.Body()
			if blocks := body.Blocks(); len(blocks) > 0 {
				body = blocks[0].Body()
			}
			object := body.GetAttribute("a").Expr().AsObjectConsExpr()

			for _, key := range test.keys {
				object.SetItemValue(key, cty.StringVal("bowler"))
			}

			got := f.Bytes()
			if string(got) != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if formatted := Format(got); string(formatted) != string(got) {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := hclsyntax.ParseConfig(got, "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags.Error())
			}
		})
	}
}
