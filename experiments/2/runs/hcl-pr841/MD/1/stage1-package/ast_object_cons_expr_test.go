// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/hcl/v2"
)

func TestObjectConsExprItemFor(t *testing.T) {
	src := `a = {
  hat         = "derby"
  "cat"       = "calico"
  "$${dog}"   = "beagle"
  "a hat"     = "fedora"
  "say \"hi\"" = "hello"
  true        = "yes"
  ""          = "empty"
  (band)      = "variable"
  hat.band    = "traversal"
  "${hat}"    = "template"
  "x${y}"     = "mixed"
}
`
	tests := []struct {
		key  string
		want string // the value's source, or "" when no item is found
	}{
		{key: "hat", want: `"derby"`},
		{key: "cat", want: `"calico"`},
		{key: "${dog}", want: `"beagle"`},
		{key: "$${dog}", want: ""},
		{key: "a hat", want: `"fedora"`},
		{key: `say "hi"`, want: `"hello"`},
		{key: "true", want: `"yes"`},
		{key: "", want: `"empty"`},
		{key: "band", want: ""},
		{key: "(band)", want: ""},
		{key: "hat.band", want: ""},
		{key: "${hat}", want: ""},
		{key: "x${y}", want: ""},
		{key: "missing", want: ""},
	}

	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	if got := len(object.Items()); got != 11 {
		t.Fatalf("got %d items, want 11", got)
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			item := object.ItemFor(test.key)
			value := object.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil || value != nil {
					t.Fatalf("found an item for %q, want none", test.key)
				}
				return
			}
			if item == nil || value == nil {
				t.Fatalf("found no item for %q", test.key)
			}
			if value != item.ValueExpr() {
				t.Error("ValueExprFor returned a different value than ItemFor")
			}
			got := strings.TrimSpace(string(value.Expr().BuildTokens(nil).Bytes()))
			if got != test.want {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	f, diags := ParseConfig([]byte("a = { b = 1, (c) = 2, \"d\" = 3 }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	var got []string
	for _, item := range object.Items() {
		got = append(got, strings.TrimSpace(string(item.KeyExpr().BuildTokens(nil).Bytes())))
	}
	want := []string{"b", "(c)", `"d"`}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("wrong keys\n%s", diff)
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{key: "hat", want: `hat`},
		{key: "hat_band-2", want: `hat_band-2`},
		{key: "a hat", want: `"a hat"`},
		{key: "hat.band", want: `"hat.band"`},
		{key: "1hat", want: `"1hat"`},
		{key: `say "hi"`, want: `"say \"hi\""`},
		{key: `a\b`, want: `"a\\b"`},
		{key: "${hat}", want: `"$${hat}"`},
		{key: "%{hat}", want: `"%%{hat}"`},
		{key: "$hat", want: `"$hat"`},
		{key: "line\nbreak", want: `"line\nbreak"`},
		{key: "", want: `""`},
		// A naked "for" as the first key would start a for expression.
		{key: "for", want: `"for"`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = {}\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatal(diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, value := object.SetItemValue(test.key, cty.StringVal("bowler"))
			if got := string(key.BuildTokens(nil).Bytes()); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			item := object.ItemFor(test.key)
			if item == nil {
				t.Fatalf("ItemFor(%q) found nothing after SetItemValue", test.key)
			}
			if item.KeyExpr() != key || item.ValueExpr() != value {
				t.Error("ItemFor found a different item than SetItemValue returned")
			}

			// The result evaluates to an object with the original key.
			out := f.Bytes()
			reparsed, diags := ParseConfig(out, "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("invalid result:\n%s\n%s", out, diags)
			}
			if reparsed.Body().GetAttribute("a").Expr().AsObjectConsExpr().ItemFor(test.key) == nil {
				t.Errorf("ItemFor(%q) found nothing after reparsing:\n%s", test.key, out)
			}
		})
	}
}

func TestObjectConsExprSetItemExisting(t *testing.T) {
	src := "a = { \"hat\" = \"derby\", cat = \"calico\" }\n"
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	oldItem := object.ItemFor("hat")

	key, value := object.SetItemValue("hat", cty.StringVal("bowler"))
	if key != oldItem.KeyExpr() || value != oldItem.ValueExpr() {
		t.Error("SetItemValue did not return the existing item's key and value")
	}
	if got := len(object.Items()); got != 2 {
		t.Errorf("got %d items, want 2", got)
	}

	want := "a = { \"hat\" = \"bowler\", cat = \"calico\" }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestObjectConsExprSetItemKinds(t *testing.T) {
	f, diags := ParseConfig([]byte("a = {}\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	object.SetItemRaw("raw", TokensForFunctionCall("upper", TokensForValue(cty.StringVal("x"))))
	object.SetItemValue("value", cty.NumberIntVal(5))
	_, value := object.SetItemTraversal("traversal", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "hat"},
	})
	object.SetItem("expr", NewExpressionLiteral(cty.True))

	want := "a = { raw = upper(\"x\"), value = 5, traversal = var.hat, expr = true }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
	}

	// The traversal is a variable of the object, so it can be renamed.
	vars := value.Expr().Variables()
	if len(vars) != 1 {
		t.Fatalf("got %d variables in the traversal value, want 1", len(vars))
	}
	f.Body().GetAttribute("a").Expr().RenameVariablePrefix([]string{"var"}, []string{"local"})
	want = "a = { raw = upper(\"x\"), value = 5, traversal = local.hat, expr = true }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result after renaming\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	src := "a = { hat = \"derby\" }\nb = \"bowler\"\n"
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	attached := NewExpressionLiteral(cty.StringVal("bowler"))
	object.SetItem("cat", attached)

	tests := map[string]*Expression{
		"attribute expression from the parser": f.Body().GetAttribute("b").Expr(),
		"item value from the parser":           object.ValueExprFor("hat").Expr(),
		"expression attached by SetItem":       attached,
	}
	for name, expr := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("SetItem did not panic")
				}
			}()
			object.SetItem("dog", expr)
		})
	}

	// A replaced expression is no longer attached, so it can be reused.
	replaced := object.ValueExprFor("cat").Expr()
	object.SetItemValue("cat", cty.StringVal("calico"))
	object.SetItem("dog", replaced)

	want := "a = { hat = \"derby\", cat = \"calico\", dog = \"bowler\" }\nb = \"bowler\"\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := []struct {
		name string
		src  string
		keys []string
		want string
	}{
		{
			name: "single line",
			src:  "a = { hat = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			name: "empty",
			src:  "a = {}\n",
			keys: []string{"dog"},
			want: "a = { dog = \"bowler\" }\n",
		},
		{
			name: "empty, two items",
			src:  "a = {}\n",
			keys: []string{"dog", "cat"},
			want: "a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			name: "single line, trailing comma",
			src:  "a = { hat = \"derby\", }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			name: "single line, inline comment before the brace",
			src:  "a = { hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		{
			name: "single line, multi-line value",
			src:  "a = { hat = {\n  band = \"red\"\n} }\n",
			keys: []string{"dog"},
			want: "a = { hat = {\n  band = \"red\"\n}, dog = \"bowler\" }\n",
		},
		{
			name: "single line, multi-line key",
			src:  "a = { (\n  hat\n) = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = { (\n  hat\n) = \"derby\", dog = \"bowler\" }\n",
		},
		{
			name: "multi-line",
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, empty",
			src:  "a = {\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, brace on the last item's line",
			src:  "a = {\n  hat = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, brace on the first item's line",
			src:  "a = { hat = \"derby\"\n  cat = \"calico\" }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\"\n  cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, comma and line comment after the last item",
			src:  "a = {\n  hat = \"derby\", # a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", # a hat\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, inline comment after the last item",
			src:  "a = {\n  hat = \"derby\" /* a hat */\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" /* a hat */\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, comments before the brace",
			src:  "a = {\n  hat = \"derby\"\n\n  # end of a\n  // really\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n\n  # end of a\n  // really\n}\n",
		},
		{
			name: "multi-line, inline comment directly before the brace",
			src:  "a = {\n  hat = \"derby\"\n/* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			name: "multi-line, line comment after the opening brace",
			src:  "a = { # no items\n}\n",
			keys: []string{"dog"},
			want: "a = { # no items\n  dog = \"bowler\"\n}\n",
		},
		{
			name: "multi-line, aligned with existing items",
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"a hat"},
			want: "a = {\n  hat     = \"derby\"\n  \"a hat\" = \"bowler\"\n}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatal(diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
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
			if _, diags := ParseConfig(got, "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is invalid: %s", diags)
			}
		})
	}
}
