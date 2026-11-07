# hclwrite: parse function call and tuple-construct expressions

`hclwrite` builds a structured syntax tree for object-construct expressions, but a function call (`join(" ", ["fedora", "fez"])`) or a tuple-construct expression (`["fedora", "fez"]`) is still kept as a flat run of tokens. Please parse both into a structured tree as well, in the same style as object-construct expressions, so that later work can add read and write APIs for them.

## Public API

```go
// FunctionCallExpr represents a function call expression such as join(" ", list).
type FunctionCallExpr struct { /* ... */ }

// Name returns the name of the called function.
func (call *FunctionCallExpr) Name() string

// TupleConsExpr represents a tuple-construct expression such as ["fedora", "fez"].
type TupleConsExpr struct { /* ... */ }

// AsFunctionCallExpr returns the wrapped function call expression, or nil if
// this expression does not parse as a function call expression.
func (e *Expression) AsFunctionCallExpr() *FunctionCallExpr

// AsTupleConsExpr returns the wrapped tuple-construct expression, or nil if
// this expression does not parse as a tuple-construct expression.
func (e *Expression) AsTupleConsExpr() *TupleConsExpr
```

## Behaviour

- In the tree, each argument of a function call and each element of a tuple is its own `Expression`. The tokens between them (the function name and opening parenthesis, the commas, the brackets and the closing parenthesis) stay as tokens of the enclosing node. Nested calls, tuples and objects are parsed recursively.
- The source must round-trip unchanged: `File.Bytes()` returns exactly the parsed bytes.
- `Expression.Variables` and `Expression.RenameVariablePrefix` must keep working for traversals inside function arguments and tuple elements, including when such a call or tuple is nested in an object item (`a = { k = f(old.x) }`, `a = { k = [old.x] }`). Variables are still reported in source order.

Add parser tests for these expressions.

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
