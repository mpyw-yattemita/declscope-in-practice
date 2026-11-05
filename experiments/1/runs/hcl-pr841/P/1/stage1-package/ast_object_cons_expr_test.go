// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"fmt"
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
		want string // source of the value found, or "" for none
	}{
		{`a = { hat = 1 }`, "hat", "1"},
		{`a = { "hat" = 1 }`, "hat", "1"},
		{`a = { hat: 1 }`, "hat", "1"},
		{`a = { cat = 1, hat = 2 }`, "hat", "2"},
		{`a = { "$${hat}" = 1 }`, "${hat}", "1"},
		{`a = { "%%{hat}" = 1 }`, "%{hat}", "1"},
		{`a = { "a hat" = 1 }`, "a hat", "1"},
		{`a = { "say \"hi\"" = 1 }`, `say "hi"`, "1"},
		{`a = { "a\\b" = 1 }`, `a\b`, "1"},
		{`a = { "" = 1 }`, "", "1"},
		{`a = { true = 1 }`, "true", "1"},
		{`a = { hat = 1 }`, "cat", ""},
		{`a = { hat = 1 }`, `"hat"`, ""},
		{`a = { (hat) = 1 }`, "hat", ""},
		{`a = { hat.band = 1 }`, "hat", ""},
		{`a = { hat.band = 1 }`, "hat.band", ""},
		{`a = { "${hat}" = 1 }`, "hat", ""},
		{`a = { "${hat}" = 1 }`, "${hat}", ""},
		{`a = { "$${hat}" = 1 }`, "$${hat}", ""},
		{`a = { "x${hat}" = 1 }`, "xhat", ""},
		{`a = { (hat) = 1, hat = 2 }`, "hat", "2"},
		{`a = {}`, "hat", ""},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %q", test.src, test.key), func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src+"\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			item := obj.ItemFor(test.key)
			value := obj.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil || value != nil {
					t.Fatalf("found an item, but want none")
				}
				return
			}
			if item == nil || value == nil {
				t.Fatalf("found no item, but want one")
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a different value than ItemFor")
			}
			got := strings.TrimSpace(string(value.Expr().BuildTokens(nil).Bytes()))
			if got != test.want {
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
	obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	obj.SetItemValue("eel", cty.NumberIntVal(4))

	var got []string
	for _, item := range obj.Items() {
		key := strings.TrimSpace(string(item.KeyExpr().BuildTokens(nil).Bytes()))
		value := strings.TrimSpace(string(item.ValueExpr().Expr().BuildTokens(nil).Bytes()))
		got = append(got, key+"="+value)
	}
	want := []string{"hat=1", "(cat)=2", `"dog"=3`, "eel=4"}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("wrong items\ngot:  %q\nwant: %q", got, want)
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string // how the new key is written
	}{
		{"dog", "dog"},
		{"dog-1", "dog-1"},
		{"dog_1", "dog_1"},
		{"a hat", `"a hat"`},
		{"hat.band", `"hat.band"`},
		{"1hat", `"1hat"`},
		{`say "hi"`, `"say \"hi\""`},
		{`a\b`, `"a\\b"`},
		{"${hat}", `"$${hat}"`},
		{"%{hat}", `"%%{hat}"`},
		{"$hat", `"$hat"`},
		{"a\nb", `"a\nb"`},
		{"", `""`},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			f, diags := ParseConfig([]byte("a = {}\n"), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, value := obj.SetItemValue(test.key, cty.StringVal("bowler"))

			want := fmt.Sprintf("a = { %s = \"bowler\" }\n", test.want)
			if got := string(f.Bytes()); got != want {
				t.Fatalf("wrong result\ngot:  %s\nwant: %s", got, want)
			}
			if got := strings.TrimSpace(string(key.BuildTokens(nil).Bytes())); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			if got := obj.ItemFor(test.key); got == nil || got.ValueExpr() != value {
				t.Errorf("ItemFor does not find the item that was set")
			}

			// The key must evaluate back to the original string.
			nativeFile, diags := hclsyntax.ParseConfig(f.Bytes(), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("result is not valid HCL: %s", diags)
			}
			attrs, _ := nativeFile.Body.JustAttributes()
			val, diags := attrs["a"].Expr.Value(nil)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics evaluating result: %s", diags)
			}
			want2 := cty.ObjectVal(map[string]cty.Value{test.key: cty.StringVal("bowler")})
			if !val.RawEquals(want2) {
				t.Errorf("wrong value\ngot:  %#v\nwant: %#v", val, want2)
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
			"a = { hat = \"derby\", (cat) = \"calico\" }\n",
			"hat",
			"a = { hat = \"bowler\", (cat) = \"calico\" }\n",
		},
		{
			"a = { \"hat\" = \"derby\" }\n",
			"hat",
			"a = { \"hat\" = \"bowler\" }\n",
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
			"a = {\n  hat = {\n    band = \"red\"\n  }\n}\n",
			"hat",
			"a = {\n  hat = \"bowler\"\n}\n",
		},
		{
			// A key from a variable is not the same key, so a new item is added.
			"a = { (hat) = \"derby\" }\n",
			"hat",
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()

			key, value := obj.SetItemRaw(test.key, TokensForValue(cty.StringVal("bowler")))

			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:  %q\nwant: %q", got, test.want)
			}
			item := obj.ItemFor(test.key)
			if item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("SetItemRaw did not return the key and value of the item")
			}
		})
	}
}

func TestObjectConsExprSetItemLayout(t *testing.T) {
	tests := []struct {
		src  string
		keys []string
		want string
	}{
		// Single line
		{
			"a = { hat = \"derby\" }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"a = {}\n",
			[]string{"dog"},
			"a = { dog = \"bowler\" }\n",
		},
		{
			"a = {}\n",
			[]string{"dog", "cat"},
			"a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"a = { hat = \"derby\", }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		{
			"a = { hat = \"derby\", }\n",
			[]string{"dog", "cat"},
			"a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		{
			"a = { hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		{
			"a = { hat = { band = \"red\" } }\n",
			[]string{"dog"},
			"a = { hat = { band = \"red\" }, dog = \"bowler\" }\n",
		},

		// Several lines
		{
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\"\n}\n",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		{
			"a = {\n}\n",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\",\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\", cat = \"calico\"\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\" }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n    hat = \"derby\", (cat) = \"calico\" }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", (cat) = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = { hat = \"derby\",\n  cat = \"calico\" }\n",
			[]string{"dog"},
			"a = { hat = \"derby\",\n  cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\" # derby\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\" # derby\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\", // derby\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\", // derby\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\" /* derby */\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\" /* derby */\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\"\n  # end of a\n}\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  # end of a\n}\n",
		},
		{
			"a = {\n  hat = \"derby\" # derby\n\n  # end of a\n  # really\n}\n",
			[]string{"dog", "cat"},
			"a = {\n  hat = \"derby\" # derby\n  dog = \"bowler\"\n  cat = \"bowler\"\n\n  # end of a\n  # really\n}\n",
		},
		{
			"a = {\n  hat = \"derby\" /* end of a */ }\n",
			[]string{"dog"},
			"a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		{
			"a = { # a\n}\n",
			[]string{"dog"},
			"a = { # a\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  # nothing yet\n}\n",
			[]string{"dog"},
			"a = {\n  dog = \"bowler\"\n  # nothing yet\n}\n",
		},
		{
			"a = {\n  hat = {\n    band = \"red\"\n  }\n}\n",
			[]string{"dog"},
			"a = {\n  hat = {\n    band = \"red\"\n  }\n  dog = \"bowler\"\n}\n",
		},
		{
			"a = {\n  hat = \"derby\"\n  dog = \"husky\"\n}\n",
			[]string{"a hat"},
			"a = {\n  hat     = \"derby\"\n  dog     = \"husky\"\n  \"a hat\" = \"bowler\"\n}\n",
		},
		{
			"b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			[]string{"dog"},
			"b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},

		// Line breaks inside an item do not make the object multi-line.
		{
			"a = { hat = {\n  band = \"red\"\n} }\n",
			[]string{"dog"},
			"a = { hat = {\n  band = \"red\"\n}, dog = \"bowler\" }\n",
		},
		{
			"a = { hat = <<EOT\nderby\nEOT\n}\n",
			[]string{"dog"},
			"a = { hat = <<EOT\nderby\nEOT\n  dog = \"bowler\"\n}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			attr := f.Body().GetAttribute("a")
			if attr == nil {
				attr = f.Body().Blocks()[0].Body().GetAttribute("a")
			}
			obj := attr.Expr().AsObjectConsExpr()
			if obj == nil {
				t.Fatalf("not an object-construct expression")
			}

			for _, key := range test.keys {
				obj.SetItemValue(key, cty.StringVal("bowler"))
				if obj.ItemFor(key) == nil {
					t.Fatalf("ItemFor(%q) does not find the item that was set", key)
				}
			}

			got := string(f.Bytes())
			if got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			if formatted := string(Format([]byte(got))); formatted != got {
				t.Errorf("result is not in canonical form\ngot:\n%s\nformatted:\n%s", got, formatted)
			}
			if _, diags := hclsyntax.ParseConfig([]byte(got), "", hcl.InitialPos); diags.HasErrors() {
				t.Errorf("result is not valid HCL: %s", diags)
			}
		})
	}
}

func TestObjectConsExprSetItemTraversal(t *testing.T) {
	f, diags := ParseConfig([]byte("a = { hat = old.x }\n"), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	expr := f.Body().GetAttribute("a").Expr()
	obj := expr.AsObjectConsExpr()

	obj.SetItemTraversal("old", hcl.Traversal{
		hcl.TraverseRoot{Name: "old"},
		hcl.TraverseAttr{Name: "y"},
	})
	if got, want := string(f.Bytes()), "a = { hat = old.x, old = old.y }\n"; got != want {
		t.Fatalf("wrong result\ngot:  %q\nwant: %q", got, want)
	}

	// The new key is a literal name, so only the new value is a variable.
	if got := len(expr.Variables()); got != 2 {
		t.Errorf("got %d variables, want 2", got)
	}
	expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
	if got, want := string(f.Bytes()), "a = { hat = new.x, old = new.y }\n"; got != want {
		t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, want)
	}
}

func TestObjectConsExprSetItemAttachedPanics(t *testing.T) {
	src := "a = { hat = \"derby\", cat = \"calico\" }\nb = \"bowler\"\n"

	tests := map[string]func(f *File, obj *ObjectConsExpr) *Expression{
		"attribute from parser": func(f *File, obj *ObjectConsExpr) *Expression {
			return f.Body().GetAttribute("b").Expr()
		},
		"item value from parser": func(f *File, obj *ObjectConsExpr) *Expression {
			return obj.ValueExprFor("cat").Expr()
		},
		"value set by SetItem": func(f *File, obj *ObjectConsExpr) *Expression {
			expr := NewExpressionLiteral(cty.StringVal("bowler"))
			obj.SetItem("dog", expr)
			return expr
		},
		"value replaced by SetItem": func(f *File, obj *ObjectConsExpr) *Expression {
			expr := NewExpressionLiteral(cty.StringVal("bowler"))
			obj.SetItem("hat", expr)
			return expr
		},
	}

	for name, getExpr := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
			expr := getExpr(f, obj)
			before := string(f.Bytes())

			func() {
				defer func() {
					if recover() == nil {
						t.Errorf("SetItem did not panic")
					}
				}()
				obj.SetItem("eel", expr)
			}()

			if after := string(f.Bytes()); after != before {
				t.Errorf("SetItem changed the file before it panicked\nbefore: %q\nafter:  %q", before, after)
			}
		})
	}

	t.Run("replaced value is detached", func(t *testing.T) {
		f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
		if diags.HasErrors() {
			t.Fatalf("unexpected diagnostics: %s", diags)
		}
		obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
		old := obj.ValueExprFor("hat").Expr()
		obj.SetItemValue("hat", cty.StringVal("bowler"))
		obj.SetItem("eel", old)

		want := "a = { hat = \"bowler\", cat = \"calico\", eel = \"derby\" }\nb = \"bowler\"\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:  %q\nwant: %q", got, want)
		}
	})
}
