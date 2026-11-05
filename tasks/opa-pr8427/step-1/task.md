# ast: add a `not` AST node type

Add a new AST node, `ast.Not`, that holds a body of expressions. When the compiler rewrites a negated expression, it must put the expressions it generates inside this body instead of in front of the negated expression. The `not` succeeds if the body does not evaluate (if any expression in it fails).

Illustration (`not { ... }` is not accepted Rego syntax; the braces only show the structure):

```rego
deny if {
	not {
		__local1__ = input.user
		data.example.is_admin(__local1__)
	}
}
```

This is the first part of a larger piece of work, so the behaviour is opt-in: it is enabled by a new future keyword, `future.keywords.not`, which is **not** in the default capabilities yet. Users must add it to their capabilities to use it. Evaluation of `Not` expressions in `topdown` has already been written and will be in the workspace; this task is the `ast` package side. Planner support and Rego syntax for explicit `not { ... }` bodies are out of scope.

## Public API

```go
// Not is both a Term value and a Node.
type Not struct {
	Body     Body      `json:"body"`
	Location *Location `json:"location,omitempty"`
}

func NotTerm(exprs ...*Expr) *Term // a Term whose value is &Not{Body: NewBody(exprs...)}
func NotExpr(exprs ...*Expr) *Expr // NewExpr(&Not{Body: NewBody(exprs...)})

func (n *Not) Copy() *Not // deep copy
func (n *Not) Loc() *Location
func (n *Not) SetLoc(l *Location)
func (n *Not) Equal(other Value) bool
func (n *Not) Compare(other Value) int
func (n *Not) Find(path Ref) (Value, error) // the Not itself for an empty path, otherwise not found
func (n *Not) Hash() int
func (n *Not) IsGround() bool               // true if the body is ground
func (n *Not) String() string               // "not {" + body + "}"
func (n *Not) StringLength() int            // length of String()
func (n *Not) AppendText(buf []byte) ([]byte, error) // appends the same text as String()

func (expr *Expr) IsNot() bool     // the expression's terms are a *Not
func (expr *Expr) IsNegated() bool // Negated is set, or IsNot()

func (c *Capabilities) ContainsFutureKeyword(kw string) bool
```

`Not` has no other fields. Two `Not` values compare by their bodies. `NewExpr` accepts a `*Not`, and `Copy`, `Compare`, sorting and the expression comparison handle it: a `*Not` sorts after `*With` and before `*Head`, and an expression whose terms are a `*Not` sorts after one whose terms are an `*Every`. The generic walkers and visitors descend into a `Not`'s body. `WalkClosures` treats a `Not` as a closure, and a `VarVisitor` with `SkipClosures` does not descend into it.

## The future keyword

- `"not"` becomes a future keyword. The new parsing is switched on by `import future.keywords.not`, by `"not"` in `ParserOptions.FutureKeywords`, or by `ParserOptions.AllFutureKeywords`, in each case only if the capabilities allow the keyword. `not` is already a keyword token, so enabling it changes only how negated expressions are parsed. The wildcard `import future.keywords` does not switch on the new parsing.
- `CapabilitiesForThisVersion()` leaves `"not"` out of `FutureKeywords` for every Rego version for now. `import future.keywords.not` is only accepted when the capabilities list it. With such capabilities, the error for an unknown keyword lists it: `unexpected keyword, must be one of [contains every if in not]`.
- For Rego v1 modules, the parser currently enables all allowed future keywords implicitly. From now on it enables only the keywords that v0 needed to import (`contains`, `every`, `if`, `in`), so v1 modules must also import `future.keywords.not` (or use `AllFutureKeywords`) to get the new parsing.
- The required capabilities that the compiler computes for a module include the `not` future keyword when the module imports the `future.keywords` wildcard, for both v0 and v1 modules.

## Parsing

With the keyword enabled, a negated expression `not <expr>` (including its `with` modifiers) is parsed into an expression whose terms are `&Not{Body: NewBody(<expr>), Location: ...}`. Without it, parsing is unchanged: the expression has `Negated: true`.

Example: with `import future.keywords.not`, the body of `p if { not 1 + 1 == 3 }` is `NewBody(NewExpr(&Not{Body: NewBody(Equal.Expr(CallTerm(NewTerm(Plus.Ref()), NumberTerm("1"), NumberTerm("1")), NumberTerm("3")))}))`.

## Compiling

All compiler stages that rewrite bodies must handle `Not` bodies, so that the expressions they generate for a negated expression stay inside its `Not` body: expanding refs and calls into local variables, rewriting dynamic terms, rewriting `print` calls (and erasing them when print is disabled), rewriting `rego.metadata` calls, resolving refs, and reordering for safety. `not` expressions inside comprehensions and `every` bodies are handled the same way. Assignments are allowed inside a `Not` body (the compiler generates them), so the declared-variable rewriting does not reject them.

Example: with `import future.keywords.not`, `p if { not input.x + input.y == input.z }` compiles to a single `Not` expression whose body is

```
__local1__ = input.x
__local2__ = input.y
plus(__local1__, __local2__, __local0__)
__local3__ = input.z
equal(__local0__, __local3__)
```

Without the import, the same rule compiles as before: the generated expressions come first and the last expression is negated.

Safety: a variable inside a `Not` body is safe only if it is safe outside the `Not`, or if it is assigned inside the body (by unification or as a call output) and also used by another expression of the body. If the body has a single expression, all its variables must be safe outside. Unsafe variables are reported with the existing message, for example `var a is unsafe` for `not a := 1`, `var x is unsafe` for `p := [x | not x := "foo"]`, and similarly for `not plus(1, 2, a)` and `not x = input.y`.

## Debugging aid

Also add an internal (unexported) debugging helper that renders a parsed or compiled `*Module` as a Mermaid flowchart and returns it as a string. Its output starts with `flowchart TD\n`. It has a node for the module labelled `"Module"`, a node for the package labelled like `"Package: data.example"`, and a node per import labelled like `"Import: rego.v1"` (with ` as <alias>` if aliased). Rules, heads, expressions (including `Not`, `every` and `with`) and terms become further nodes. The module node is linked to the package, imports and rules by edges labelled `package`, `import` and `rule` (written `-->|package|` and so on). Use it to log the compiled module in the `not` compiler tests.

Add parser and compiler tests for the cases above.
