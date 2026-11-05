# Add tools for the sub-issue endpoints

GitHub issues can have sub-issues (https://docs.github.com/en/rest/issues/sub-issues). Please add four tools for them to the `issues` toolset: `list_sub_issues` as a read tool, and `add_sub_issue`, `remove_sub_issue` and `reprioritize_sub_issue` as write tools. They use the REST API.

The expected tool schema snapshots for the four tools are provided in the workspace, next to the existing snapshots. The tool definitions must match them exactly (names, titles, descriptions, read-only hints, parameters and required parameters).

## Public API

```go
func AddSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc)
func ListSubIssues(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc)
func RemoveSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc)
func ReprioritizeSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc)
```

Use translation keys in the existing style, for example `TOOL_ADD_SUB_ISSUE_DESCRIPTION` and `TOOL_ADD_SUB_ISSUE_USER_TITLE`.

## Tools

- `add_sub_issue`: parameters `owner`, `repo`, `issue_number` (the parent issue), `sub_issue_id` (the sub-issue's ID, which is not its number), all required, and optional boolean `replace_parent`. Calls `POST /repos/{owner}/{repo}/issues/{issue_number}/sub_issues`. Success is 201; returns the sub-issue as JSON.
- `list_sub_issues`: parameters `owner`, `repo`, `issue_number` (required), and optional `page` (default 1) and `per_page` (default 30). Calls `GET /repos/{owner}/{repo}/issues/{issue_number}/sub_issues`. Success is 200; returns the list as JSON.
- `remove_sub_issue`: parameters `owner`, `repo`, `issue_number`, `sub_issue_id`, all required. Calls `DELETE /repos/{owner}/{repo}/issues/{issue_number}/sub_issue` with the JSON body `{"sub_issue_id": ...}`. Success is 200; returns the response body (the parent issue) as JSON. The go-github `SubIssue.Remove` method has a bug (google/go-github#3613), so this may need a direct request through the authenticated client.
- `reprioritize_sub_issue`: parameters `owner`, `repo`, `issue_number`, `sub_issue_id` (required), and `after_id` or `before_id` (the sub-issue to place it after or before). Calls `PATCH /repos/{owner}/{repo}/issues/{issue_number}/sub_issues/priority`. Success is 200; returns the sub-issue as JSON.

Errors (returned as tool error results, not Go errors):

- A missing required parameter gives the existing message, for example `missing required parameter: sub_issue_id`.
- `reprioritize_sub_issue` with neither `after_id` nor `before_id`: `either after_id or before_id must be specified`. With both: `only one of after_id or before_id should be specified, not both`.
- An API error or an unexpected status gives a message that starts with `failed to add sub-issue`, `failed to list sub-issues`, `failed to remove sub-issue` or `failed to reprioritize sub-issue`, using the package's GitHub API error response helper where the client returns an error.

Register the tools, regenerate the tool documentation in the README with the repository's docs generator, and add tests for success and error cases.

The following files are part of the specification and are already in place:
- pkg/github/__toolsnaps__/add_sub_issue.snap
- pkg/github/__toolsnaps__/list_sub_issues.snap
- pkg/github/__toolsnaps__/remove_sub_issue.snap
- pkg/github/__toolsnaps__/reprioritize_sub_issue.snap

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
go test ./pkg/github
```
