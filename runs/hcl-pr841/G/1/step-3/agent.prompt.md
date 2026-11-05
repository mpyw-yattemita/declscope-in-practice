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

# Checking the rule above

This repository runs declscope, a linter that checks the rule above. The `declscope-authoring` skill explains how to act on its reports. Read it before you start.

`.declscope.yaml` sets the rules, and `.declscope-baseline.yaml` records the reports the existing code already has, so only what you add is reported. Never edit either file, and never regenerate the baseline.

Before you finish, run these as well, and resolve what they report:

```bash
declscope ./...
declscope shrink ./...
```

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./hclwrite
declscope ./...
declscope shrink ./...
```
