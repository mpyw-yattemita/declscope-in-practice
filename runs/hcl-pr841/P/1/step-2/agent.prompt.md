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

# Where code goes in this package

This package keeps each file's internals to itself. Follow these rules when you add or change code.

Existing code may not follow them yet. Do not fix existing code beyond what the task needs. Follow the rules in the code you add or change.

## The rule

| Declaration | Who may use it |
| --- | --- |
| Unexported | Only the file it is declared in |
| Exported | Any file of the package |

The same holds for members:

- An unexported field may be read or written only in the file that declares its struct.
- A method belongs to the file it is written in, not to its type's file. A method in another file must not touch the type's unexported fields.
- A test file goes with the source file of the same name. `parse_test.go` may use the unexported code of `parse.go`, and of no other file. So name a test file for the source file whose unexported code it calls, not for the behaviour it tests.

When code in one file needs an unexported declaration of another file, that is a crossing. Avoid creating new ones.

## Names carry their file's name

A file's name, without `.go` and in camelCase, is a word that the names declared in it carry. `user_repository.go` gives `userRepository`.

| Declaration | Its name |
| --- | --- |
| A package-level function, type, variable or constant, exported or not | Contains the file's word, beginning where a word in the name begins: `newUserRepository`, `UserRepositoryOption`, `findUserRepositoryRows` |
| A method written in its type's file, a field, an interface method | Free |
| A method written in another file | Contains that other file's word |
| A helper in a test file | Contains the word of the source file the test file goes with |
| `TestXxx`, `BenchmarkXxx`, `FuzzXxx`, `ExampleXxx` | Free |

This applies once the package has more than one file with declarations.

- Fit the word where it reads naturally: as a modifier (`userProfile`), as an object (`loadUser`), or later in the phrase (`addToCollection`). Read the name aloud.
- Name the thing accurately first, then make sure the word is in it. Do not bend a name only to carry the word.
- When the word reads badly in a name, the declaration is usually in the wrong file. Move it rather than force the name.
- When a name is fixed, for example by the task, put the declaration in a file whose word that name carries.
- Do not rename existing declarations only to follow this.

## Where a new declaration goes

| The declaration is used by | Put it |
| --- | --- |
| Code in one file | That file |
| Several files, as a helper they use in common on purpose | A file named for its concept, not the file of its first caller. Say in its doc comment which files use it and why |
| Most files of the package, as its common model | The file that already holds the package's common types, if there is one |
| Another package | Wherever it belongs, exported |

**Before you reach into another file, ask whether your code is in the wrong file.** Often the function you are writing belongs beside what it uses.

**Do not create a file just to hold one small declaration** that only its old file uses.

## Common changes

| Change | Do this |
| --- | --- |
| Add a method to a type | Put it in the type's file |
| Add a field that another file reads | Keep it only if the other file really needs it. Say in the field's comment which file reads it |
| Move code between files | Move its tests with it |
| Split a file | Split where the code really divides. Calls between the halves become crossings |

## Do not hide a crossing

| Tempting | Why not |
| --- | --- |
| Exporting a name so another file may use it | It hides the crossing and widens the API. In an `internal/` package, do not export something no other package uses |
| Moving everything into one file | It hides the structure instead of fixing it |
| Widening production code for a test | Move the test instead |

## Order declarations by reach

Inside a file, put what other packages use first, then what other files use, then the file's own internals. Keep a type with its constructors and methods. Do not reorder declarations your change does not touch, and do not reorder where the order is observable, such as `iota` or unkeyed literals.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./hclwrite
```
