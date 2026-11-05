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
