This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/parser.go:443: parseTemplateExpr sets `expr.wrapped = expr.children.Append(quoted)` directly. Every other wrapping site now calls the new `Expression.wrap`, which also guards against wrapping twice. This site skips that guard and keeps a second way of setting `wrapped`. Fix: Replace the two lines with `expr.wrap(quoted)`, the same as parseObjectConsExpr, parseFunctionCallExpr and parseTupleConsExpr.
1. pkg/ast_expression.go:172: AsObjectConsExpr, AsFunctionCallExpr, AsTupleConsExpr and asQuoted repeat the same nil check and type assertion on `e.wrapped.content`. The new code copies this block twice more, and AsObjectConsExpr is still the only exported one with no doc comment. Fix: Add one generic helper, `func wrappedAs[T nodeContent](e *Expression) T`, that returns the zero value when `wrapped` is nil or holds a different type. Make all four methods one-line calls to it, and give AsObjectConsExpr a doc comment like the other two.
2. pkg/ast_expression.go:222: Expression.traversals type-switches over *ObjectConsExpr, *FunctionCallExpr and *TupleConsExpr only to call the same `traversals(to)` method on each. Every new wrapped expression kind needs another case here. Fix: Declare an unexported interface `type traversalHolder interface{ traversals([]*Traversal) []*Traversal }`. Replace the switch with `if h, ok := e.wrapped.content.(traversalHolder); ok { to = h.traversals(to) }`.
3. pkg/ast_tuple_cons_expr.go:29: TupleConsExpr is FunctionCallExpr with only the field renamed (elems instead of args). The constructor and the traversals method are copies, and parseTupleConsExpr and parseFunctionCallExpr are copies too. They already share parseExpressionList. Fix: Pull the shared part into an embedded type, e.g. `type exprList struct { inTree; items nodeSet }`, with `newExprList()` and `traversals(to)`. Embed it in both structs and delete the duplicated traversals methods. parseExpressionList can then take the `*exprList`.
4. pkg/ast_expression.go:91: NewExpressionAbsTraversal builds identifier tokens by hand with `&Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)}` (root and attribute steps). The package already has newIdentToken for this, and ast_attribute.go, ast_block.go and generate.go use it. Fix: Use `newIdentifier(newIdentToken(rootName))` and `newIdentifier(newIdentToken(ts.Name))`.
5. pkg/ast_object_cons_expr.go:162: appendItem writes the newline token literal `&Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}` twice (lines 162 and 164). objectConsBreaksLine only loops tokenIsNewline over a slice, which is a general Tokens helper hidden under an object-specific name. Fix: Add `newNewlineToken()` next to newIdentToken in tokens.go and use it at both places. Rename objectConsBreaksLine to a Tokens method, e.g. `func (ts Tokens) hasNewline() bool`, in tokens.go.
6. pkg/partition.go:133: partitionableTokens.Types() has no callers in the package, test files included. It came over in the move from parser.go and is dead code. Fix: Delete the Types method. While editing this file, also fix the doc comment at line ~72, which says `PartitionIncludeComments` but the method is `PartitionIncludingComments`.
7. pkg/ast_object_cons_expr.go:210: `gofmt -l` lists 11 changed or new files: ast_expression.go, ast_function_call_expr.go, ast_function_call_expr_test.go, ast_object_cons_expr.go, ast_object_cons_expr_test.go, ast_tuple_cons_expr.go, format.go, parser.go, parser_test.go, partition.go, tokens.go. The causes are the unaligned `key *node` / `value *node` fields in ObjectConsItem, the double blank line after the license header in the new files, and a trailing empty `//` line on doc comments (tokenIsNewline, writerTokens, lexConfigTokens, parse, the traversals methods, the partitionableTokens fields). Fix: Run `gofmt -w` on the package. Then delete the leftover trailing `//` lines by hand so the doc comments end at the last sentence.

For each finding, decide whether it is real and worth fixing now. For each
kept finding, give one category:

- naming: a name does not match its role, or one idea has two names
- duplication: the same logic is written twice
- over-abstraction: a shared helper couples callers that are not related
- mixed-job: one function or type does two unrelated jobs
- dead: a definition is unused
- other: none of the above

Answer with one JSON object and nothing else:
{"verdicts": [{"index": 0, "keep": true, "category": "naming", "why": "..."}]}
