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
