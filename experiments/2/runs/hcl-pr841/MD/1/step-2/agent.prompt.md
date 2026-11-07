# hclwrite: add a read and write API for `ObjectConsExpr`

`Expression.AsObjectConsExpr()` gives access to an object-construct expression, but there is no way to look up its items by key or to rewrite them. Please add a read and write API, similar to the existing `Body` API for attributes (`GetAttribute`, `SetAttributeRaw`, `SetAttributeValue`, `SetAttributeTraversal`).

Example:

```go
config := []byte(`a = {
    hat = "derby", (cat) = "calico" }` + "\n")

bowler := hclwrite.TokensForValue(cty.StringVal(`bowler`))

file, _ := hclwrite.ParseConfig(config, "<inline>", hcl.Pos{})
attr := file.Body().GetAttribute("a")

expr := attr.Expr().AsObjectConsExpr()
key, value := expr.SetItemRaw("hat", bowler)

key.BuildTokens(nil).WriteTo(os.Stdout)
value.BuildTokens(nil).WriteTo(os.Stdout)
```

## Public API

On `*ObjectConsExpr`:

```go
func (o *ObjectConsExpr) Items() []*ObjectConsItem
func (o *ObjectConsExpr) ItemFor(key string) *ObjectConsItem
func (o *ObjectConsExpr) ValueExprFor(key string) *ObjectConsValue
func (o *ObjectConsExpr) SetItem(key string, expr *Expression) (*ObjectConsKeyExpr, *ObjectConsValue)
func (o *ObjectConsExpr) SetItemRaw(key string, tokens Tokens) (*ObjectConsKeyExpr, *ObjectConsValue)
func (o *ObjectConsExpr) SetItemValue(key string, val cty.Value) (*ObjectConsKeyExpr, *ObjectConsValue)
func (o *ObjectConsExpr) SetItemTraversal(key string, traversal hcl.Traversal) (*ObjectConsKeyExpr, *ObjectConsValue)
```

On `*ObjectConsItem`:

```go
func (item *ObjectConsItem) KeyExpr() *ObjectConsKeyExpr
func (item *ObjectConsItem) ValueExpr() *ObjectConsValue
```

The existing `ObjectConsExpr.ValueFor` is renamed to `ValueExprFor`, and `ObjectConsKeyExpr.String` is removed. Both were added recently and have not been released, so no compatibility shim is needed.

Add doc comments to the exported types `ObjectConsExpr`, `ObjectConsItem`, `ObjectConsKeyExpr` and `ObjectConsValue` and to the new methods.

## Behaviour

Lookup (`ItemFor`, `ValueExprFor`):

- An item is found by key only if its key is an identifier or a quoted string literal. The two forms are equivalent: `ItemFor("hat")` finds both `hat = ...` and `"hat" = ...`.
- A quoted key is compared by its value after escape sequences are decoded: `"hat"` matches `"hat"`, `"$${hat}"` matches `"${hat}"`, and `"a hat"` matches `"a hat"`.
- Keys taken from a variable or another expression are never found: `(hat)`, `hat.band` and `"${hat}"` do not match `"hat"`.
- Both return nil if no item matches. `ValueExprFor` returns the value of the item that `ItemFor` finds.

Writing (`SetItem` and the `SetItem*` helpers):

- If an item with the key exists (by the lookup rules above), its value is replaced and its key is left as written. Otherwise, a new item is added at the end of the object.
- A new key is written as an identifier if it is a valid identifier. Otherwise it is written as a quoted string literal, escaped so that it evaluates back to the original string: `say "hi"` becomes `"say \"hi\""`, `a\b` becomes `"a\\b"`, `${hat}` becomes `"$${hat}"`, `%{hat}` becomes `"%%{hat}"`, and the empty string becomes `""`. Keys such as `a hat`, `hat.band` and `1hat` are quoted. A key that was just set must be found again by `ItemFor`.
- `SetItemRaw` uses the tokens verbatim, like `NewExpressionRaw`. `SetItemValue` writes a literal, like `NewExpressionLiteral`. `SetItemTraversal` writes a traversal, like `NewExpressionAbsTraversal`.
- All of them return the key and the value of the item that was modified or created.
- `SetItem` panics if the given expression is already part of a syntax tree. This covers an expression that the parser attached, and an expression that an earlier `SetItem` call attached. Reusing an attached expression would make two items share one expression, so renaming a variable in one would also change the other.
- The result must be valid HCL, laid out as a person would write it, and already in canonical form, so that `Format` does not change it:
  - In an object written on a single line, a new item goes on the same line after a comma: `a = { hat = "derby" }` becomes `a = { hat = "derby", dog = "bowler" }`. An empty object `a = {}` becomes `a = { dog = "bowler" }`. Adding two items gives `a = { dog = "bowler", cat = "bowler" }`. An existing trailing comma is not doubled.
  - In an object written on several lines, a new item goes on its own line after the last item, before the closing brace. If the closing brace shares a line with the last item (`hat = "derby" }`), the brace moves to its own line.
  - A comma or a line comment that follows the last item on its line stays with that item. Comments on the following lines, before the closing brace, stay with the closing brace, so new items go before them. An inline comment directly before the closing brace (`/* end of a */ }`) stays with the brace.
  - Line breaks inside an item's key or value, such as a nested multi-line object, do not make the outer object count as multi-line.

Add tests for the lookup rules, the key escaping, the panics and the layout cases.

# declscope

This repository runs declscope, a linter that checks where each declaration of a Go package is used and how it is named. The `declscope-authoring` skill explains its rules and how to act on its reports. Read it before you start.

`.declscope.yaml` sets the rules, and `.declscope-baseline.yaml` records the reports the existing code already has, so only new reports are shown. Never edit either file, and never regenerate the baseline.

Before you finish, run these as well, and resolve what they report:

```bash
declscope ./...
declscope shrink ./...
```

# Scope of the change

Keep this change to what the task needs. Do not refactor existing code: do not move, rename or split existing declarations or files unless the task requires it.

Nobody will answer questions while you work on this task. When a decision would normally go to the repository's owner, make the choice you would recommend, carry it out, and say in your final message what you chose and why.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./hclwrite
declscope ./...
declscope shrink ./...
```
