# hclwrite: parse a standalone expression, and set an attribute from an `Expression`

`hclwrite` can now parse expressions into a structured tree, which callers can cast to specific expression types (for example with `AsObjectConsExpr`) and change. But there is no way to parse an expression on its own, and no way to set an attribute to an `Expression` value that keeps this structure. `SetAttributeRaw`, `SetAttributeValue` and `SetAttributeTraversal` each build the expression themselves.

To avoid breaking existing callers, please add two new public functions and leave the existing ones as they are.

## Public API

```go
// ParseExpression attempts to parse the given buffer as a HCL expression,
// returning the resulting syntax element if successful.
func ParseExpression(src []byte, filename string, start hcl.Pos) (*Expression, hcl.Diagnostics)

// SetAttribute either replaces the expression of an existing attribute of the
// given name or adds a new attribute definition to the end of the block,
// using the given expression. The return value is the attribute that was
// either modified in-place or created.
func (b *Body) SetAttribute(name string, expr *Expression) *Attribute
```

## Behaviour

- `ParseExpression` works like `ParseConfig`, but for a single expression. It returns the same structured tree that `ParseConfig` would build for that expression as an attribute value: object-construct, function call and tuple-construct expressions are structured, and everything else is kept as tokens.
- If the source is not a valid expression, `ParseExpression` returns a nil expression and the error diagnostics.
- Comments and blank lines before or after the expression are not part of the result. For example, parsing `"# foo\n\n1\n# bar"` gives an expression whose tokens are just `1`.
- `SetAttribute` keeps the expression it is given, so `Attribute.Expr()` on the result returns that same structured expression. For example, after `body.SetAttribute("a", expr)` with an `expr` from `ParseExpression("{ hat = \"derby\" }", ...)`, `body.GetAttribute("a").Expr().AsObjectConsExpr()` is not nil.
- `SetAttributeRaw`, `SetAttributeValue` and `SetAttributeTraversal` keep their current behaviour and output.

Add tests that parse a variety of expressions with `ParseExpression`: literals, templates, traversals, object-construct expressions with the different kinds of keys, nested objects, function calls and tuples.

Nobody will answer questions while you work on this task. When a decision would normally go to the repository's owner, make the choice you would recommend, carry it out, and say in your final message what you chose and why.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./hclwrite
```
