# ast: reject functions in the dynamic extent of a ref-head rule

Consider this policy:

```rego
package example

a[x].c := i if some i, x in ["one", "two"]

a.b.d(x) := x
```

The function `a.b.d` lies within the dynamic extent of the rule `a[x].c`, that is, under the part of its ref that is only known at evaluation time. The compiler accepts this today, but evaluation cannot handle it: querying `data.example` fails with `var requires evaluation`, and querying the function directly panics with a nil pointer dereference.

There is no valid scenario in which evaluation supports a function within a rule's dynamic extent, so the compiler should report an error when it builds the rule tree. (This may be loosened in the future.)

Report it as a rule conflict, in the same way as the existing conflicts between a rule and the rules below it:

- `a[x].c := i if ...` with `a.b.d(x) := x` in the same package: `rego_type_error: rule data.pkg.a[x].c conflicts with [data.pkg.a.b.d]` (for package `pkg`).
- The same when the function is defined in another package that lies in the extent: `package pkg` with `a[x].c := ...`, and `package pkg.a.b` with `d(x) := x`, give the same error.
- Deeper nesting: `p[x].q if x := "a"` with `p.r.s.t(x) := x` gives `rego_type_error: rule data.pkg.p[x].q conflicts with [data.pkg.p.r.s.t]`.
- Only functions (rules with arguments) conflict. `p[x].q if x := "a"` with `p.r.s := 1` is not an error.
- Only rules whose ref is not ground are affected. If another conflict is already reported for the node, report that one.

Add compiler tests for these cases.
