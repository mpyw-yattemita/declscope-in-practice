// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package hclwrite

import (
	"slices"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

func TestObjectConsExprVariables(t *testing.T) {
	tests := map[string]struct {
		src string
		// want is the source of each traversal that Variables returns,
		// in order.
		want []string
		// renamed is the source after renaming the prefix "old" to "new".
		renamed string
	}{
		"value": {
			src:     "a = { k = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = new.x }\n",
		},
		"multiple items in source order": {
			src:     "a = { k = old.x, j = other.y, (old.z) = old.w }\n",
			want:    []string{"old.x", "other.y", "old.z", "old.w"},
			renamed: "a = { k = new.x, j = other.y, (new.z) = new.w }\n",
		},
		"nested object": {
			src:     "a = { k = { j = old.x } }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = { j = new.x } }\n",
		},
		"tuple value": {
			src:     "a = { k = [old.x] }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = [new.x] }\n",
		},
		"function call value": {
			src:     "a = { k = f(old.x) }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = f(new.x) }\n",
		},
		"object in tuple value": {
			src:     "a = { k = [{ j = old.x }] }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = [{ j = new.x }] }\n",
		},
		"naked identifier key": {
			src:     "a = { old = 1 }\n",
			want:    nil,
			renamed: "a = { old = 1 }\n",
		},
		"naked identifier key in nested object": {
			src:     "a = { k = { old = old.x } }\n",
			want:    []string{"old.x"},
			renamed: "a = { k = { old = new.x } }\n",
		},
		"parenthesized key": {
			src:     "a = { (old.x) = 1 }\n",
			want:    []string{"old.x"},
			renamed: "a = { (new.x) = 1 }\n",
		},
		"parenthesized single name key": {
			src:     "a = { (old) = 1 }\n",
			want:    []string{"old"},
			renamed: "a = { (new) = 1 }\n",
		},
		"naked traversal key": {
			// hclsyntax treats only a single name as a literal key.
			src:     "a = { old.x = 1 }\n",
			want:    []string{"old.x"},
			renamed: "a = { new.x = 1 }\n",
		},
		"quoted key": {
			src:     "a = { \"k\" = old.x }\n",
			want:    []string{"old.x"},
			renamed: "a = { \"k\" = new.x }\n",
		},
		"empty object": {
			src:     "a = {}\n",
			want:    nil,
			renamed: "a = {}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Variables must agree with hclsyntax on the same source.
			nf, diags := hclsyntax.ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected hclsyntax diagnostics: %s", diags)
			}
			nativeAttr := nf.Body.(*hclsyntax.Body).Attributes["a"]
			nativeCount := len(nativeAttr.Expr.Variables())
			if nativeCount != len(test.want) {
				t.Fatalf("test case is wrong: hclsyntax reports %d traversals, but want has %d", nativeCount, len(test.want))
			}

			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			expr := f.Body().GetAttribute("a").Expr()

			var got []string
			for _, traversal := range expr.Variables() {
				got = append(got, strings.TrimSpace(string(traversal.BuildTokens(nil).Bytes())))
			}
			if !slices.Equal(got, test.want) {
				t.Errorf("wrong variables\ngot:  %q\nwant: %q", got, test.want)
			}

			expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
			if got := string(f.Bytes()); got != test.renamed {
				t.Errorf("wrong result after rename\ngot:  %q\nwant: %q", got, test.renamed)
			}
		})
	}
}

// parseObjectConsExprForTest parses src and returns the object-construct expression
// of its attribute "a".
func parseObjectConsExprForTest(t *testing.T, src string) (*File, *ObjectConsExpr) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags)
	}
	obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	if obj == nil {
		t.Fatalf("attribute a is not an object-construct expression")
	}
	return f, obj
}

func tokensSource(tokens Tokens) string {
	return strings.TrimSpace(string(tokens.Bytes()))
}

func TestObjectConsExprItems(t *testing.T) {
	_, obj := parseObjectConsExprForTest(t, "a = { hat = 1, (cat) = 2, \"dog\" = 3 }\n")

	var got []string
	for _, item := range obj.Items() {
		got = append(got, tokensSource(item.KeyExpr().BuildTokens(nil))+"="+tokensSource(item.ValueExpr().BuildTokens(nil)))
	}
	want := []string{"hat=1", "(cat)=2", `"dog"=3`}
	if !slices.Equal(got, want) {
		t.Errorf("wrong items\ngot:  %q\nwant: %q", got, want)
	}

	_, empty := parseObjectConsExprForTest(t, "a = {}\n")
	if items := empty.Items(); len(items) != 0 {
		t.Errorf("empty object has %d items", len(items))
	}
}

func TestObjectConsExprItemFor(t *testing.T) {
	tests := map[string]struct {
		src string
		key string
		// want is the source of the value that is found, or "" if no
		// item is found.
		want string
	}{
		"identifier":                     {`a = { hat = "derby" }`, "hat", `"derby"`},
		"quoted":                         {`a = { "hat" = "derby" }`, "hat", `"derby"`},
		"colon":                          {`a = { hat: "derby" }`, "hat", `"derby"`},
		"quoted with space":              {`a = { "a hat" = "derby" }`, "a hat", `"derby"`},
		"quoted with escaped template":   {`a = { "$${hat}" = "derby" }`, "${hat}", `"derby"`},
		"quoted with escaped quote":      {`a = { "say \"hi\"" = "derby" }`, `say "hi"`, `"derby"`},
		"keyword":                        {`a = { null = "derby" }`, "null", `"derby"`},
		"second item":                    {`a = { cat = "calico", hat = "derby" }`, "hat", `"derby"`},
		"first of duplicates":            {`a = { hat = 1, "hat" = 2 }`, "hat", `1`},
		"after expression key":           {`a = { (hat) = 1, hat = 2 }`, "hat", `2`},
		"missing":                        {`a = { cat = "calico" }`, "hat", ``},
		"empty object":                   {`a = {}`, "hat", ``},
		"parenthesized":                  {`a = { (hat) = "derby" }`, "hat", ``},
		"traversal":                      {`a = { hat.band = "derby" }`, "hat.band", ``},
		"traversal prefix":               {`a = { hat.band = "derby" }`, "hat", ``},
		"template":                       {`a = { "${hat}" = "derby" }`, "hat", ``},
		"template source":                {`a = { "${hat}" = "derby" }`, "${hat}", ``},
		"escaped template is not raw":    {`a = { "$${hat}" = "derby" }`, "$${hat}", ``},
		"identifier is case sensitive":   {`a = { Hat = "derby" }`, "hat", ``},
		"nested object is not searched":  {`a = { cat = { hat = "derby" } }`, "hat", ``},
		"nested object value is found":   {`a = { cat = { hat = "derby" } }`, "cat", `{ hat = "derby" }`},
		"function call key":              {`a = { upper("hat") = "derby" }`, "HAT", ``},
		"number key":                     {`a = { 1 = "derby" }`, "1", ``},
		"quoted key with only a newline": {`a = { "\n" = "derby" }`, "\n", `"derby"`},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, obj := parseObjectConsExprForTest(t, test.src+"\n")

			item := obj.ItemFor(test.key)
			value := obj.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil || value != nil {
					t.Fatalf("found an item, want none")
				}
				return
			}
			if item == nil || value == nil {
				t.Fatalf("found no item")
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a different value than ItemFor")
			}
			if got := tokensSource(value.BuildTokens(nil)); got != test.want {
				t.Errorf("wrong value\ngot:  %s\nwant: %s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := map[string]struct {
		key  string
		want string
	}{
		"identifier":       {"hat", `hat`},
		"identifier dash":  {"top-hat", `top-hat`},
		"keyword":          {"null", `null`},
		"space":            {"a hat", `"a hat"`},
		"dot":              {"hat.band", `"hat.band"`},
		"leading digit":    {"1hat", `"1hat"`},
		"quote":            {`say "hi"`, `"say \"hi\""`},
		"backslash":        {`a\b`, `"a\\b"`},
		"interpolation":    {"${hat}", `"$${hat}"`},
		"directive":        {"%{hat}", `"%%{hat}"`},
		"escaped template": {"$${hat}", `"$$${hat}"`},
		"newline":          {"a\nb", `"a\nb"`},
		"empty":            {"", `""`},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseObjectConsExprForTest(t, "a = {}\n")

			key, value := obj.SetItemValue(test.key, cty.True)
			if got := tokensSource(key.BuildTokens(nil)); got != test.want {
				t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
			}
			if got := obj.ItemFor(test.key); got == nil || got.KeyExpr() != key || got.ValueExpr() != value {
				t.Errorf("ItemFor does not find the item that was set")
			}

			// The key must evaluate back to the original string.
			nf, diags := hclsyntax.ParseConfig(f.Bytes(), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("result is not valid HCL: %s\n%s", diags, f.Bytes())
			}
			val, diags := nf.Body.(*hclsyntax.Body).Attributes["a"].Expr.Value(nil)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			want := cty.ObjectVal(map[string]cty.Value{test.key: cty.True})
			if !val.RawEquals(want) {
				t.Errorf("wrong value\ngot:  %#v\nwant: %#v", val, want)
			}
		})
	}
}

func TestObjectConsExprSetItemReplace(t *testing.T) {
	tests := map[string]struct {
		src  string
		key  string
		want string
		// added is true if SetItemRaw adds a new item instead of
		// replacing a value.
		added bool
	}{
		"identifier": {
			src:  "a = { hat = \"derby\", cat = \"calico\" }\n",
			key:  "hat",
			want: "a = { hat = \"bowler\", cat = \"calico\" }\n",
		},
		"quoted key is left as written": {
			src:  "a = { \"hat\" = \"derby\" }\n",
			key:  "hat",
			want: "a = { \"hat\" = \"bowler\" }\n",
		},
		"escaped key is left as written": {
			src:  "a = { \"$${hat}\" = \"derby\" }\n",
			key:  "${hat}",
			want: "a = { \"$${hat}\" = \"bowler\" }\n",
		},
		"colon is kept": {
			src:  "a = { hat: \"derby\" }\n",
			key:  "hat",
			want: "a = { hat : \"bowler\" }\n",
		},
		"nested object value": {
			src:  "a = {\n  hat = {\n    b = 1\n  }\n  cat = 2\n}\n",
			key:  "hat",
			want: "a = {\n  hat = \"bowler\"\n  cat = 2\n}\n",
		},
		"comments are kept": {
			src:  "a = {\n  # lead\n  hat = \"derby\" # line\n}\n",
			key:  "hat",
			want: "a = {\n  # lead\n  hat = \"bowler\" # line\n}\n",
		},
		"expression key is not replaced": {
			src:   "a = { (hat) = \"derby\" }\n",
			key:   "hat",
			want:  "a = { (hat) = \"derby\", hat = \"bowler\" }\n",
			added: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseObjectConsExprForTest(t, test.src)
			itemsBefore := len(obj.Items())

			key, value := obj.SetItemRaw(test.key, TokensForValue(cty.StringVal("bowler")))
			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
			item := obj.ItemFor(test.key)
			if item == nil || item.KeyExpr() != key || item.ValueExpr() != value {
				t.Errorf("SetItemRaw did not return the item that ItemFor finds")
			}
			wantItems := itemsBefore
			if test.added {
				wantItems++
			}
			if got := len(obj.Items()); got != wantItems {
				t.Errorf("wrong number of items %d, want %d", got, wantItems)
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, obj := parseObjectConsExprForTest(t, "a = { hat = \"derby\" }\n")

	obj.SetItemValue("hat", cty.StringVal("bowler"))
	obj.SetItemValue("size", cty.NumberIntVal(7))
	obj.SetItemTraversal("cat", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "cat"},
	})
	obj.SetItemRaw("dog", Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte("upper")},
		{Type: hclsyntax.TokenOParen, Bytes: []byte("(")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("var")},
		{Type: hclsyntax.TokenDot, Bytes: []byte(".")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("dog")},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(")")},
	})

	want := "a = { hat = \"bowler\", size = 7, cat = var.cat, dog = upper(var.dog) }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	// The traversal is a variable of the object, but raw tokens are not
	// interpreted, as with NewExpressionRaw.
	expr := f.Body().GetAttribute("a").Expr()
	expr.RenameVariablePrefix([]string{"var"}, []string{"local"})
	want = "a = { hat = \"bowler\", size = 7, cat = local.cat, dog = upper(var.dog) }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result after rename\ngot:  %s\nwant: %s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	const src = "a = { hat = \"derby\", cat = { b = 1 } }\nb = 1\n"

	tests := map[string]func(f *File, obj *ObjectConsExpr) *Expression{
		"parsed item value": func(f *File, obj *ObjectConsExpr) *Expression {
			return obj.ValueExprFor("hat").Expr()
		},
		"parsed nested item value": func(f *File, obj *ObjectConsExpr) *Expression {
			return obj.ValueExprFor("cat").Expr().AsObjectConsExpr().ValueExprFor("b").Expr()
		},
		"parsed attribute": func(f *File, obj *ObjectConsExpr) *Expression {
			return f.Body().GetAttribute("b").Expr()
		},
		"object itself": func(f *File, obj *ObjectConsExpr) *Expression {
			return f.Body().GetAttribute("a").Expr()
		},
		"new item value": func(f *File, obj *ObjectConsExpr) *Expression {
			_, value := obj.SetItemValue("dog", cty.True)
			return value.Expr()
		},
		"replaced item value": func(f *File, obj *ObjectConsExpr) *Expression {
			_, value := obj.SetItemValue("hat", cty.True)
			return value.Expr()
		},
		"expression passed to SetItem": func(f *File, obj *ObjectConsExpr) *Expression {
			expr := NewExpressionLiteral(cty.True)
			obj.SetItem("dog", expr)
			return expr
		},
		"attribute set on the body": func(f *File, obj *ObjectConsExpr) *Expression {
			return f.Body().SetAttributeValue("c", cty.True).Expr()
		},
	}

	for name, get := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseObjectConsExprForTest(t, src)
			expr := get(f, obj)
			before := string(f.Bytes())

			func() {
				defer func() {
					if recover() == nil {
						t.Errorf("SetItem did not panic")
					}
				}()
				obj.SetItem("other", expr)
			}()

			if got := string(f.Bytes()); got != before {
				t.Errorf("the tree changed\ngot:\n%s\nwant:\n%s", got, before)
			}
		})
	}

	t.Run("replaced expression can be reused", func(t *testing.T) {
		f, obj := parseObjectConsExprForTest(t, src)
		old := obj.ValueExprFor("hat").Expr()
		obj.SetItemValue("hat", cty.True)
		obj.SetItem("dog", old)

		want := "a = { hat = true, cat = { b = 1 }, dog = \"derby\" }\nb = 1\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, want)
		}
	})
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
		"single line two items": {
			src:  "a = { hat = \"derby\" }\n",
			keys: []string{"dog", "cat"},
			want: "a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		"empty": {
			src:  "a = {}\n",
			keys: []string{"dog"},
			want: "a = { dog = \"bowler\" }\n",
		},
		"empty two items": {
			src:  "a = {}\n",
			keys: []string{"dog", "cat"},
			want: "a = { dog = \"bowler\", cat = \"bowler\" }\n",
		},
		"empty with spaces": {
			src:  "a = {  }\n",
			keys: []string{"dog"},
			want: "a = { dog = \"bowler\" }\n",
		},
		"single line trailing comma": {
			src:  "a = { hat = \"derby\", }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" }\n",
		},
		"single line trailing comma two items": {
			src:  "a = { hat = \"derby\", }\n",
			keys: []string{"dog", "cat"},
			want: "a = { hat = \"derby\", dog = \"bowler\", cat = \"bowler\" }\n",
		},
		"single line inline comment before brace": {
			src:  "a = { hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		"single line comma and inline comment": {
			src:  "a = { hat = \"derby\", /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		"single line comment after attribute": {
			src:  "a = { hat = \"derby\" } # a\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" } # a\n",
		},
		"single line nested multi-line value": {
			src:  "a = { hat = {\n  b = 1\n} }\n",
			keys: []string{"dog"},
			want: "a = { hat = {\n  b = 1\n}, dog = \"bowler\" }\n",
		},
		"heredoc value ends a line": {
			src:  "a = { hat = <<EOT\nderby\nEOT\n}\n",
			keys: []string{"dog"},
			want: "a = { hat = <<EOT\nderby\nEOT\n  dog = \"bowler\"\n}\n",
		},
		"multi line": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"multi line two items": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi line empty": {
			src:  "a = {\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi line empty with comment": {
			src:  "a = {\n  # nothing yet\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  dog = \"bowler\"\n  # nothing yet\n}\n",
		},
		"multi line empty with comment after brace": {
			src:  "a = { # nothing yet\n}\n",
			keys: []string{"dog"},
			want: "a = { # nothing yet\n  dog = \"bowler\"\n}\n",
		},
		"multi line commas": {
			src:  "a = {\n  hat = \"derby\",\n  cat = \"calico\",\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\",\n  cat = \"calico\",\n  dog = \"bowler\"\n}\n",
		},
		"brace on the line of the last item": {
			src:  "a = {\n  hat = \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"brace on the line of the last item with comma": {
			src:  "a = {\n  hat = \"derby\", }\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"items on the line of the opening brace": {
			src:  "a = { hat = \"derby\",\n  cat = \"calico\" }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\",\n  cat = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		"line comment after the last item": {
			src:  "a = {\n  hat = \"derby\" # a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" # a hat\n  dog = \"bowler\"\n}\n",
		},
		"comma and line comment after the last item": {
			src:  "a = {\n  hat = \"derby\", // a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", // a hat\n  dog = \"bowler\"\n}\n",
		},
		"comments before the closing brace": {
			src:  "a = {\n  hat = \"derby\"\n  # more hats\n\n  # later\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n  # more hats\n\n  # later\n}\n",
		},
		"inline comment before the closing brace": {
			src:  "a = {\n  hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"inline comment on the line of the closing brace": {
			src:  "a = {\n  hat = \"derby\"\n/* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"inline comment after the last item": {
			src:  "a = {\n  hat = \"derby\" /* a hat */\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" /* a hat */\n  dog = \"bowler\"\n}\n",
		},
		"multi line nested object": {
			src:  "a = {\n  hat = {\n    b = 1\n  }\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = {\n    b = 1\n  }\n  dog = \"bowler\"\n}\n",
		},
		"nested object": {
			src:  "a = {\n  hat = { b = 1 }\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = { b = 1 }\n  dog = \"bowler\"\n}\n",
		},
		"in a block": {
			src:  "b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			keys: []string{"dog"},
			want: "b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags)
			}
			body := f.Body()
			if blocks := body.Blocks(); len(blocks) > 0 {
				body = blocks[0].Body()
			}
			obj := body.GetAttribute("a").Expr().AsObjectConsExpr()

			for _, key := range test.keys {
				obj.SetItemRaw(key, TokensForValue(cty.StringVal("bowler")))
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
			for _, key := range test.keys {
				if obj.ItemFor(key) == nil {
					t.Errorf("ItemFor(%q) found nothing", key)
				}
			}
		})
	}
}
