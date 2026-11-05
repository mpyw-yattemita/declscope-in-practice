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
	tests := []struct {
		src  string
		key  string
		want string // the value found, as source text, or "" for none
	}{
		{`a = { hat = "derby" }`, "hat", `"derby"`},
		{`a = { "hat" = "derby" }`, "hat", `"derby"`},
		{`a = { "$${hat}" = "derby" }`, "${hat}", `"derby"`},
		{`a = { "%%{hat}" = "derby" }`, "%{hat}", `"derby"`},
		{`a = { "a hat" = "derby" }`, "a hat", `"derby"`},
		{`a = { "say \"hi\"" = "derby" }`, `say "hi"`, `"derby"`},
		{`a = { "a\\b" = "derby" }`, `a\b`, `"derby"`},
		{`a = { "" = "derby" }`, "", `"derby"`},
		{`a = { true = "derby" }`, "true", `"derby"`},
		{`a = { hat: "derby" }`, "hat", `"derby"`},
		{`a = { cat = "calico", hat = "derby" }`, "hat", `"derby"`},
		{`a = { (hat) = "derby" }`, "hat", ""},
		{`a = { hat.band = "derby" }`, "hat.band", ""},
		{`a = { hat.band = "derby" }`, "hat", ""},
		{`a = { "${hat}" = "derby" }`, "hat", ""},
		{`a = { "${hat}" = "derby" }`, "${hat}", ""},
		{`a = { "$${hat}" = "derby" }`, "$${hat}", ""},
		{`a = { hat = "derby" }`, "cat", ""},
		{`a = { hat = "derby" }`, "Hat", ""},
		{`a = {}`, "hat", ""},
	}

	for _, test := range tests {
		t.Run(test.src+" "+test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			item := object.ItemFor(test.key)
			value := object.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil || value != nil {
					t.Fatalf("found an item, but want none")
				}
				return
			}
			if item == nil || value == nil {
				t.Fatalf("found no item")
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a value other than that of ItemFor")
			}
			if got := strings.TrimSpace(string(value.BuildTokens(nil).Bytes())); got != test.want {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	src := "a = { hat = 1, (cat) = 2, \"dog\" = 3 }\n"
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	items := f.Body().GetAttribute("a").Expr().AsObjectConsExpr().Items()

	var got []string
	for _, item := range items {
		key := strings.TrimSpace(string(item.KeyExpr().BuildTokens(nil).Bytes()))
		value := strings.TrimSpace(string(item.ValueExpr().BuildTokens(nil).Bytes()))
		got = append(got, key+"="+value)
	}
	want := []string{"hat=1", "(cat)=2", `"dog"=3`}
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("wrong items\ngot:  %q\nwant: %q", got, want)
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string // the key as written
	}{
		{"dog", `dog`},
		{"dog-1", `dog-1`},
		{"true", `true`},
		{"for", `"for"`},
		{`say "hi"`, `"say \"hi\""`},
		{`a\b`, `"a\\b"`},
		{"${hat}", `"$${hat}"`},
		{"%{hat}", `"%%{hat}"`},
		{"$hat", `"$hat"`},
		{"", `""`},
		{"a hat", `"a hat"`},
		{"hat.band", `"hat.band"`},
		{"1hat", `"1hat"`},
		{"line\nbreak", `"line\nbreak"`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = {}\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, _ := object.SetItemValue(test.key, cty.True)
			if got := strings.TrimSpace(string(key.BuildTokens(nil).Bytes())); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			if item := object.ItemFor(test.key); item == nil || item.KeyExpr() != key {
				t.Errorf("ItemFor does not find the key that was set")
			}

			// The key must evaluate back to the original string.
			src := f.Bytes()
			native, diags := hclsyntax.ParseConfig(src, "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("result is not valid HCL: %s\n%s", diags, src)
			}
			val, diags := native.Body.(*hclsyntax.Body).Attributes["a"].Expr.Value(nil)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			if !val.Type().IsObjectType() || !val.Type().HasAttribute(test.key) {
				t.Errorf("result has no key %q: %#v", test.key, val)
			}
		})
	}
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	bowler := TokensForValue(cty.StringVal("bowler"))

	tests := []struct {
		name string
		src  string
		keys []string // the keys set, in order, each to "bowler"
		want string
	}{
		{
			"single line",
			"a = { hat = \"derby\" }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"empty",
			"a = {}\n",
			[]string{"dog"},
			"a = { dog = \"bowler\" }\n",
		},
		{
			"empty, two items",
			"a = {}\n",
			[]string{"dog", "cat"},
			"a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"single line with trailing comma",
			"a = { hat = \"derby\", }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"single line with trailing comma, two items",
			"a = { hat = \"derby\", }\n",
			[]string{"dog", "cat"},
			"a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"single line with comment before the brace",
			"a = { hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		{
			"single line, replace",
			"a = { hat = \"derby\", cat = \"calico\" }\n",
			[]string{"hat"},
			"a = { hat = \"bowler\", cat = \"calico\" }\n",
		},
		{
			"single line, replace quoted",
			"a = { \"hat\" = \"derby\" }\n",
			[]string{"hat"},
			"a = { \"hat\" = \"bowler\" }\n",
		},
		{
			"single line with nested multi-line value",
			"a = { hat = {\n  band = \"red\"\n} }\n",
			[]string{"dog"},
			"a = { hat = {\n  band = \"red\"\n}, dog = \"bowler\" }\n",
		},
		{
			"multi-line",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line, two items, aligned",
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog", "color"},
			"a = {\n  hat   = \"derby\"\n  dog   = \"bowler\"\n  color = \"bowler\"\n}\n",
		},
		{
			"multi-line, replace",
			"a = {\n  hat = \"derby\"\n  cat = \"calico\"\n}\n",
			[]string{"hat"},
			"a = {\n  hat = \"bowler\"\n  cat = \"calico\"\n}\n",
		},
		{
			"multi-line empty",
			"a = {\n}\n",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n}\n",
		},
		{
			"multi-line with commas",
			"a = {\n  hat = \"derby\",\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"brace on the line of the last item",
			"a = {\n  hat = \"derby\" }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"brace on the line of the last item, two items",
			"a = {\n  hat = \"derby\" }\n",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		{
			"brace on the line of the last item, after a comma",
			"a = {\n  hat = \"derby\", }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"brace on the line of the first item",
			"a = { hat = \"derby\",\n  cat = \"calico\" }\n",
			[]string{"dog"},
			"a = { hat = \"derby\",\n  cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"line comment after the last item",
			"a = {\n  hat = \"derby\" # felt\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\" # felt\n  dog = \"bowler\"\n}\n",
		},
		{
			"comma and line comment after the last item",
			"a = {\n  hat = \"derby\", # felt\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", # felt\n  dog = \"bowler\"\n}\n",
		},
		{
			"comments before the brace",
			"a = {\n  hat = \"derby\"\n\n  # more hats\n  // later\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n\n  # more hats\n  // later\n}\n",
		},
		{
			"comment before the brace in an empty object",
			"a = {\n  # no hats\n}\n",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n  # no hats\n}\n",
		},
		{
			"inline comment directly before the brace",
			"a = {\n  hat = \"derby\"\n/* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"inline comment directly before the brace, on the line of the last item",
			"a = {\n  hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"nested multi-line value",
			"a = {\n  hat = {\n    band = \"red\"\n  }\n}\n",
			[]string{"dog"},
			"a = {\n  hat = {\n    band = \"red\"\n  }\n  dog = \"bowler\"\n}\n",
		},
		{
			"nested multi-line value with the brace on its line",
			"a = {\n  hat = {\n    band = \"red\"\n  } }\n",
			[]string{"dog"},
			"a = {\n  hat = {\n    band = \"red\"\n  }\n  dog = \"bowler\"\n}\n",
		},
		{
			"inside a block",
			"b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			[]string{"dog"},
			"b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
		{
			"key from an expression",
			"a = { (hat) = \"derby\" }\n",
			[]string{"hat"},
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			body := f.Body()
			if body.GetAttribute("a") == nil {
				body = body.Blocks()[0].Body()
			}
			object := body.GetAttribute("a").Expr().AsObjectConsExpr()

			for _, key := range test.keys {
				gotKey, gotValue := object.SetItemRaw(key, bowler)
				item := object.ItemFor(key)
				if item == nil || item.KeyExpr() != gotKey || item.ValueExpr() != gotValue {
					t.Fatalf("SetItemRaw(%q) returned other than the item ItemFor finds", key)
				}
			}

			got := f.Bytes()
			if string(got) != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if formatted := Format(got); string(formatted) != string(got) {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := hclsyntax.ParseConfig(got, "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags)
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, diags := ParseConfig([]byte("a = { hat = \"derby\" }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	expr := f.Body().GetAttribute("a").Expr()
	object := expr.AsObjectConsExpr()

	object.SetItemValue("hat", cty.StringVal("bowler"))
	object.SetItemValue("size", cty.NumberIntVal(7))
	object.SetItemTraversal("owner", hcl.Traversal{
		hcl.TraverseRoot{Name: "old"},
		hcl.TraverseAttr{Name: "name"},
	})
	object.SetItem("cat", NewExpressionLiteral(cty.True))

	want := "a = { hat = \"bowler\", size = 7, owner = old.name, cat = true }\n"
	if got := string(f.Bytes()); got != want {
		t.Fatalf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	// A traversal that was set is a variable of the object.
	expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
	want = "a = { hat = \"bowler\", size = 7, owner = new.name, cat = true }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result after renaming\ngot:  %s\nwant: %s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	src := "a = { hat = \"derby\" }\nb = \"bowler\"\n"

	tests := []struct {
		name   string
		attach func(f *File, object *ObjectConsExpr) *Expression
	}{
		{
			"expression of an attribute",
			func(f *File, object *ObjectConsExpr) *Expression {
				return f.Body().GetAttribute("b").Expr()
			},
		},
		{
			"value of an item",
			func(f *File, object *ObjectConsExpr) *Expression {
				return object.ValueExprFor("hat").Expr()
			},
		},
		{
			"value set by an earlier SetItem",
			func(f *File, object *ObjectConsExpr) *Expression {
				expr := NewExpressionLiteral(cty.True)
				object.SetItem("dog", expr)
				return expr
			},
		},
		{
			"value replaced by an earlier SetItem",
			func(f *File, object *ObjectConsExpr) *Expression {
				expr := NewExpressionLiteral(cty.True)
				object.SetItem("hat", expr)
				return expr
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
			expr := test.attach(f, object)

			defer func() {
				if recover() == nil {
					t.Errorf("SetItem did not panic")
				}
			}()
			object.SetItem("cat", expr)
		})
	}
}

func TestObjectConsExprSetItemReplacedValue(t *testing.T) {
	// The value an earlier SetItem replaced is no longer in the tree, so it
	// may be set again.
	f, diags := ParseConfig([]byte("a = { hat = \"derby\" }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	object := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

	old := object.ValueExprFor("hat").Expr()
	object.SetItemValue("hat", cty.StringVal("bowler"))
	object.SetItem("cat", old)

	want := "a = { hat = \"bowler\", cat = \"derby\" }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}
}
