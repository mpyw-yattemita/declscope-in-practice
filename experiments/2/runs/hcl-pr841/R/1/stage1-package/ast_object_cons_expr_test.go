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

// parseTestObject parses src and returns the object constructor that is
// the value of its attribute "a".
func parseTestObject(t *testing.T, src string) (*File, *ObjectConsExpr) {
	t.Helper()
	f, diags := ParseConfig([]byte(src), "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	obj := f.Body().GetAttribute("a").Expr().AsObjectConsExpr()
	if obj == nil {
		t.Fatalf("attribute a is not an object constructor")
	}
	return f, obj
}

// evalAttribute parses src with hclsyntax and evaluates its attribute "a".
func evalAttribute(t *testing.T, src []byte) cty.Value {
	t.Helper()
	f, diags := hclsyntax.ParseConfig(src, "", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("result is not valid HCL: %s\n%s", diags.Error(), src)
	}
	attrs, diags := f.Body.JustAttributes()
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics: %s", diags.Error())
	}
	val, diags := attrs["a"].Expr.Value(nil)
	if diags.HasErrors() {
		t.Fatalf("unexpected diagnostics evaluating a: %s\n%s", diags.Error(), src)
	}
	return val
}

func tokensString(toks Tokens) string {
	return strings.TrimSpace(string(toks.Bytes()))
}

func TestObjectConsExprItemFor(t *testing.T) {
	tests := map[string]struct {
		src string
		key string
		// want is the value of the item found, or "" if none is found.
		want string
	}{
		"identifier":                    {`a = { hat = 1 }`, "hat", "1"},
		"quoted":                        {`a = { "hat" = 1 }`, "hat", "1"},
		"colon":                         {`a = { hat: 1 }`, "hat", "1"},
		"second item":                   {`a = { cat = 1, hat = 2 }`, "hat", "2"},
		"first of duplicates":           {`a = { hat = 1, "hat" = 2 }`, "hat", "1"},
		"quoted with space":             {`a = { "a hat" = 1 }`, "a hat", "1"},
		"quoted with escaped template":  {`a = { "$${hat}" = 1 }`, "${hat}", "1"},
		"quoted with escaped directive": {`a = { "%%{hat}" = 1 }`, "%{hat}", "1"},
		"quoted with escaped quote":     {`a = { "say \"hi\"" = 1 }`, `say "hi"`, "1"},
		"quoted with escaped backslash": {`a = { "a\\b" = 1 }`, `a\b`, "1"},
		"quoted with unicode escape":    {`a = { "h\u0061t" = 1 }`, "hat", "1"},
		"quoted empty":                  {`a = { "" = 1 }`, "", "1"},
		"keyword":                       {`a = { null = 1 }`, "null", "1"},
		"multi-line":                    {"a = {\n  cat = 1\n  hat = 2\n}", "hat", "2"},
		"nested object value":           {`a = { hat = { b = 1 } }`, "hat", "{ b = 1 }"},

		"missing":                   {`a = { cat = 1 }`, "hat", ""},
		"empty object":              {`a = {}`, "hat", ""},
		"escaped template raw":      {`a = { "$${hat}" = 1 }`, "$${hat}", ""},
		"parenthesized":             {`a = { (hat) = 1 }`, "hat", ""},
		"parenthesized quoted":      {`a = { ("hat") = 1 }`, "hat", ""},
		"traversal":                 {`a = { hat.band = 1 }`, "hat.band", ""},
		"traversal root":            {`a = { hat.band = 1 }`, "hat", ""},
		"template":                  {`a = { "${hat}" = 1 }`, "hat", ""},
		"template source":           {`a = { "${hat}" = 1 }`, "${hat}", ""},
		"function call":             {`a = { (upper("hat")) = 1 }`, "HAT", ""},
		"key of nested object":      {`a = { cat = { hat = 1 } }`, "hat", ""},
		"case sensitive":            {`a = { Hat = 1 }`, "hat", ""},
		"quoted without decoding":   {`a = { "a\\b" = 1 }`, `a\\b`, ""},
		"value is not a key":        {`a = { cat = "hat" }`, "hat", ""},
		"identifier is not a value": {`a = { cat = hat }`, "hat", ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, obj := parseTestObject(t, test.src+"\n")

			item := obj.ItemFor(test.key)
			value := obj.ValueExprFor(test.key)
			if test.want == "" {
				if item != nil {
					t.Errorf("ItemFor found an item with key %q", tokensString(item.KeyExpr().BuildTokens(nil)))
				}
				if value != nil {
					t.Errorf("ValueExprFor found value %q", tokensString(value.BuildTokens(nil)))
				}
				return
			}

			if item == nil {
				t.Fatalf("ItemFor found no item")
			}
			if got := tokensString(item.ValueExpr().Expr().BuildTokens(nil)); got != test.want {
				t.Errorf("ItemFor found wrong item with value %q; want %q", got, test.want)
			}
			if value != item.ValueExpr() {
				t.Errorf("ValueExprFor returned a different value than ItemFor")
			}
		})
	}
}

func TestObjectConsExprItems(t *testing.T) {
	_, obj := parseTestObject(t, "a = { hat = 1, (cat) = 2, \"dog\" = 3 }\n")

	var got []string
	for _, item := range obj.Items() {
		got = append(got, tokensString(item.KeyExpr().BuildTokens(nil))+"="+tokensString(item.ValueExpr().BuildTokens(nil)))
	}
	want := []string{"hat=1", "(cat)=2", `"dog"=3`}
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("wrong items\ngot:  %q\nwant: %q", got, want)
	}

	_, empty := parseTestObject(t, "a = {}\n")
	if got := empty.Items(); len(got) != 0 {
		t.Errorf("got %d items for an empty object", len(got))
	}
}

func TestObjectConsExprSetItemKey(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{"hat", "hat"},
		{"hat-band", "hat-band"},
		{"hat_band", "hat_band"},
		{"null", "null"},
		{"true", "true"},
		{"if", "if"},
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
		{"a\nb", `"a\nb"`},
		{"hat=", `"hat="`},
		{"=", `"="`},
		{"(hat)", `"(hat)"`},
		{"#", `"#"`},
		{"ハット", "ハット"},
	}

	// Each key is set on an empty object both on a single line and on
	// several lines, and in the first and in a later position.
	srcs := []string{
		"a = {}\n",
		"a = {\n}\n",
		"a = { cat = 1 }\n",
		"a = {\n  cat = 1\n}\n",
	}

	for _, test := range tests {
		for _, src := range srcs {
			t.Run(test.key+"/"+src, func(t *testing.T) {
				f, obj := parseTestObject(t, src)

				key, value := obj.SetItemValue(test.key, cty.StringVal("x"))
				if got := tokensString(key.BuildTokens(nil)); got != test.want {
					t.Errorf("wrong key\ngot:  %s\nwant: %s", got, test.want)
				}
				if got := tokensString(value.BuildTokens(nil)); got != `"x"` {
					t.Errorf("wrong value %s", got)
				}

				item := obj.ItemFor(test.key)
				if item == nil {
					t.Fatalf("ItemFor did not find the new item")
				}
				if item.KeyExpr() != key || item.ValueExpr() != value {
					t.Errorf("ItemFor found a different item than the one returned by SetItemValue")
				}

				result := f.Bytes()
				val := evalAttribute(t, result)
				if !val.Type().HasAttribute(test.key) {
					t.Fatalf("result has no key %q:\n%s", test.key, result)
				}
				if got := val.GetAttr(test.key); !got.RawEquals(cty.StringVal("x")) {
					t.Errorf("wrong value for key %q: %#v", test.key, got)
				}

				// The key must also be found after a round trip.
				_, reparsed := parseTestObject(t, string(result))
				if reparsed.ItemFor(test.key) == nil {
					t.Errorf("ItemFor did not find the key after parsing the result:\n%s", result)
				}
			})
		}
	}
}

func TestObjectConsExprSetItemExisting(t *testing.T) {
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
		"quoted key left as written": {
			"a = { \"hat\" = \"derby\" }\n",
			"hat",
			"a = { \"hat\" = \"bowler\" }\n",
		},
		"escaped key left as written": {
			"a = { \"$${hat}\" = \"derby\" }\n",
			"${hat}",
			"a = { \"$${hat}\" = \"bowler\" }\n",
		},
		"last item": {
			"a = { cat = \"calico\", hat = \"derby\" }\n",
			"hat",
			"a = { cat = \"calico\", hat = \"bowler\" }\n",
		},
		"multi-line with comments": {
			"a = {\n  # lead\n  hat = \"derby\" # line\n  cat = \"calico\"\n}\n",
			"hat",
			"a = {\n  # lead\n  hat = \"bowler\" # line\n  cat = \"calico\"\n}\n",
		},
		"nested object value": {
			"a = {\n  hat = {\n    style = \"derby\"\n  }\n}\n",
			"hat",
			"a = {\n  hat = \"bowler\"\n}\n",
		},
		"expression key not replaced": {
			"a = { (hat) = \"derby\" }\n",
			"hat",
			"a = { (hat) = \"derby\", hat = \"bowler\" }\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, obj := parseTestObject(t, test.src)
			before := len(obj.Items())
			existing := obj.ItemFor(test.key)

			key, value := obj.SetItemValue(test.key, cty.StringVal("bowler"))

			if existing != nil {
				if key != existing.KeyExpr() || value != existing.ValueExpr() {
					t.Errorf("SetItemValue did not return the existing item")
				}
				if got := len(obj.Items()); got != before {
					t.Errorf("got %d items; want %d", got, before)
				}
			}
			if got := tokensString(value.Expr().BuildTokens(nil)); got != `"bowler"` {
				t.Errorf("wrong value %s", got)
			}
			if got := string(f.Bytes()); got != test.want {
				t.Errorf("wrong result\ngot:\n%s\nwant:\n%s", got, test.want)
			}
		})
	}
}

func TestObjectConsExprSetItemHelpers(t *testing.T) {
	f, obj := parseTestObject(t, "a = { hat = 1 }\n")

	obj.SetItemRaw("raw", Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte("upper")},
		{Type: hclsyntax.TokenOParen, Bytes: []byte("(")},
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("x")},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(")")},
	})
	obj.SetItemValue("value", cty.ListVal([]cty.Value{cty.NumberIntVal(1)}))
	_, traversal := obj.SetItemTraversal("traversal", hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: "hat"},
	})
	obj.SetItemTraversal("hat", hcl.Traversal{hcl.TraverseRoot{Name: "local"}})

	want := "a = { hat = local, raw = upper(\"x\"), value = [1], traversal = var.hat }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
	}

	// Traversals set by SetItemTraversal are variables of the object.
	expr := f.Body().GetAttribute("a").Expr()
	var vars []string
	for _, v := range expr.Variables() {
		vars = append(vars, tokensString(v.BuildTokens(nil)))
	}
	if got, want := strings.Join(vars, " "), "local var.hat"; got != want {
		t.Errorf("wrong variables %q; want %q", got, want)
	}
	if got := len(traversal.Expr().Variables()); got != 1 {
		t.Errorf("got %d variables in the new value; want 1", got)
	}

	expr.RenameVariablePrefix([]string{"var"}, []string{"input"})
	want = "a = { hat = local, raw = upper(\"x\"), value = [1], traversal = input.hat }\n"
	if got := string(f.Bytes()); got != want {
		t.Errorf("wrong result after rename\ngot:  %s\nwant: %s", got, want)
	}
}

func TestObjectConsExprSetItemPanics(t *testing.T) {
	assertPanics := func(t *testing.T, f func()) {
		t.Helper()
		defer func() {
			t.Helper()
			if recover() == nil {
				t.Errorf("did not panic")
			}
		}()
		f()
	}

	const src = "a = { hat = \"derby\" }\nb = \"calico\"\n"

	t.Run("parsed attribute expression", func(t *testing.T) {
		f, obj := parseTestObject(t, src)
		expr := f.Body().GetAttribute("b").Expr()
		assertPanics(t, func() { obj.SetItem("cat", expr) })
		assertPanics(t, func() { obj.SetItem("hat", expr) })
		if got, want := string(f.Bytes()), src; got != want {
			t.Errorf("result changed after panic\ngot:  %s\nwant: %s", got, want)
		}
	})

	t.Run("parsed item value", func(t *testing.T) {
		f, obj := parseTestObject(t, src)
		expr := obj.ValueExprFor("hat").Expr()
		assertPanics(t, func() { obj.SetItem("cat", expr) })
		assertPanics(t, func() { obj.SetItem("hat", expr) })
		if got, want := string(f.Bytes()), src; got != want {
			t.Errorf("result changed after panic\ngot:  %s\nwant: %s", got, want)
		}
	})

	t.Run("the object itself", func(t *testing.T) {
		f, obj := parseTestObject(t, src)
		expr := f.Body().GetAttribute("a").Expr()
		assertPanics(t, func() { obj.SetItem("cat", expr) })
	})

	t.Run("expression attached by an earlier call", func(t *testing.T) {
		_, obj := parseTestObject(t, src)
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "hat"}})
		obj.SetItem("cat", expr)
		assertPanics(t, func() { obj.SetItem("dog", expr) })
		assertPanics(t, func() { obj.SetItem("hat", expr) })
		assertPanics(t, func() { obj.SetItem("cat", expr) })
	})

	t.Run("expression attached by an earlier replace", func(t *testing.T) {
		_, obj := parseTestObject(t, src)
		expr := NewExpressionAbsTraversal(hcl.Traversal{hcl.TraverseRoot{Name: "hat"}})
		obj.SetItem("hat", expr)
		assertPanics(t, func() { obj.SetItem("dog", expr) })
	})

	t.Run("replaced expression can be reused", func(t *testing.T) {
		f, obj := parseTestObject(t, src)
		old := obj.ValueExprFor("hat").Expr()
		obj.SetItemValue("hat", cty.StringVal("bowler"))
		obj.SetItem("cat", old)
		want := "a = { hat = \"bowler\", cat = \"derby\" }\nb = \"calico\"\n"
		if got := string(f.Bytes()); got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
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
			src:  "a = { }\n",
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
		"single line colon": {
			src:  "a = { hat: \"derby\" }\n",
			keys: []string{"dog"},
			want: "a = { hat : \"derby\", dog = \"bowler\" }\n",
		},
		"single line inline comment before brace": {
			src:  "a = { hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\", dog = \"bowler\" /* end of a */ }\n",
		},
		"single line with nested multi-line object": {
			src:  "a = { hat = {\n  style = \"derby\"\n} }\n",
			keys: []string{"dog"},
			want: "a = { hat = {\n  style = \"derby\"\n}, dog = \"bowler\" }\n",
		},
		"single line with nested multi-line object last": {
			src:  "a = { cat = \"calico\", hat = {\n  style = \"derby\"\n  } }\n",
			keys: []string{"dog"},
			want: "a = { cat = \"calico\", hat = {\n  style = \"derby\"\n}, dog = \"bowler\" }\n",
		},
		"single line with multi-line tuple": {
			src:  "a = { hat = [\n  \"derby\",\n] }\n",
			keys: []string{"dog"},
			want: "a = { hat = [\n  \"derby\",\n], dog = \"bowler\" }\n",
		},
		"single line in block": {
			src:  "b {\n  a = { hat = \"derby\" }\n}\n",
			keys: []string{"dog"},
			want: "b {\n  a = { hat = \"derby\", dog = \"bowler\" }\n}\n",
		},
		"multi-line": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"multi-line two items": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi-line empty": {
			src:  "a = {\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi-line empty with comment": {
			src:  "a = {\n  # nothing yet\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  dog = \"bowler\"\n  # nothing yet\n}\n",
		},
		"multi-line brace on line of last item": {
			src:  "a = {\n    hat = \"derby\", (cat) = \"calico\" }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", (cat) = \"calico\"\n  dog = \"bowler\"\n}\n",
		},
		"multi-line brace on line of last item two items": {
			src:  "a = {\n  hat = \"derby\" }\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n}\n",
		},
		"multi-line first item on line of open brace": {
			src:  "a = { hat = \"derby\"\n}\n",
			keys: []string{"dog"},
			want: "a = { hat = \"derby\"\n  dog = \"bowler\"\n}\n",
		},
		"multi-line trailing comma": {
			src:  "a = {\n  hat = \"derby\",\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		"multi-line trailing comma brace on line of last item": {
			src:  "a = {\n  hat = \"derby\", }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\",\n  dog = \"bowler\"\n}\n",
		},
		"multi-line line comment": {
			src:  "a = {\n  hat = \"derby\" # a fancy hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" # a fancy hat\n  dog = \"bowler\"\n}\n",
		},
		"multi-line comma and line comment": {
			src:  "a = {\n  hat = \"derby\", // a fancy hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\", // a fancy hat\n  dog = \"bowler\"\n}\n",
		},
		"multi-line inline and line comment": {
			src:  "a = {\n  hat = \"derby\" /* fancy */ # a hat\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\" /* fancy */ # a hat\n  dog = \"bowler\"\n}\n",
		},
		"multi-line comments before brace": {
			src:  "a = {\n  hat = \"derby\"\n  # more hats\n  # later\n}\n",
			keys: []string{"dog", "cat"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n  cat = \"bowler\"\n  # more hats\n  # later\n}\n",
		},
		"multi-line blank line and comment before brace": {
			src:  "a = {\n  hat = \"derby\"\n\n  # more hats\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n\n  # more hats\n}\n",
		},
		"multi-line inline comment before brace": {
			src:  "a = {\n  hat = \"derby\"\n  /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"multi-line inline comment before brace on line of last item": {
			src:  "a = {\n  hat = \"derby\" /* end of a */ }\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n/* end of a */ }\n",
		},
		"multi-line with nested multi-line object last": {
			src:  "a = {\n  hat = {\n    style = \"derby\"\n  }\n}\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = {\n    style = \"derby\"\n  }\n  dog = \"bowler\"\n}\n",
		},
		"multi-line in block": {
			src:  "b {\n  a = {\n    hat = \"derby\"\n  }\n}\n",
			keys: []string{"dog"},
			want: "b {\n  a = {\n    hat = \"derby\"\n    dog = \"bowler\"\n  }\n}\n",
		},
		"multi-line followed by attribute": {
			src:  "a = {\n  hat = \"derby\"\n}\nb = 1\n",
			keys: []string{"dog"},
			want: "a = {\n  hat = \"derby\"\n  dog = \"bowler\"\n}\nb = 1\n",
		},
		"existing and new": {
			src:  "a = {\n  hat = \"derby\"\n}\n",
			keys: []string{"hat", "dog"},
			want: "a = {\n  hat = \"bowler\"\n  dog = \"bowler\"\n}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, diags := ParseConfig([]byte(test.src), "", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("unexpected diagnostics: %s", diags.Error())
			}
			var obj *ObjectConsExpr
			if attr := f.Body().GetAttribute("a"); attr != nil {
				obj = attr.Expr().AsObjectConsExpr()
			} else {
				obj = f.Body().Blocks()[0].Body().GetAttribute("a").Expr().AsObjectConsExpr()
			}

			for _, key := range test.keys {
				obj.SetItemValue(key, cty.StringVal("bowler"))
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
			for _, key := range test.keys {
				if obj.ItemFor(key) == nil {
					t.Errorf("ItemFor(%q) found nothing", key)
				}
			}
		})
	}
}
