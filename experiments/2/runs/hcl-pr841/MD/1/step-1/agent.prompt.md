# hclwrite: find traversals inside object-construct expressions

`hclwrite` now parses each item of an object-construct expression (`{ k = v }`) into its own key and value expressions. Since that change, the object's enclosing `Expression` no longer reports the traversals inside its items:

```go
src := "a = { k = old.x }\n"
f, _ := hclwrite.ParseConfig([]byte(src), "", hcl.InitialPos)
expr := f.Body().GetAttribute("a").Expr()
len(expr.Variables())                                   // now: 0; before the change and in hclsyntax: 1
expr.RenameVariablePrefix([]string{"old"}, []string{"new"})
string(f.Bytes())                                       // now: "a = { k = old.x }\n", unchanged
```

Please make `Expression.Variables` and `Expression.RenameVariablePrefix` see the traversals inside an object-construct expression again.

Requirements:

- `Variables` on an object-construct expression returns the traversals found in the items' keys and values, in source order.
- `RenameVariablePrefix` renames those same traversals.
- This also applies to objects nested in objects (`a = { k = { j = old.x } }`), and to traversals inside tuples or function calls that are item values (`a = { k = [old.x] }`, `a = { k = f(old.x) }`).
- A naked identifier key (`{ old = 1 }`) is a literal name, as in `hclsyntax`. It is not a variable: it is neither reported by `Variables` nor changed by `RenameVariablePrefix`. A parenthesized key (`{ (old.x) = 1 }`) is an expression, so its traversal is reported and renamed. A quoted key (`{ "k" = old.x }`) contributes no traversal.
- For each shape above, `Variables` must report the same number of traversals that `hclsyntax` reports for the same source.

Add tests for these shapes.

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

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./hclwrite
declscope ./...
declscope shrink ./...
```
