# Add pagination support to GraphQL-based tools

The discussion tools use GitHub's GraphQL API, but they always fetch a fixed first page (`first: 100`) and return a bare list, so a caller cannot page through a repository with many discussions. Please add pagination to them, following the GraphQL connection spec (https://docs.github.com/en/graphql/reference/objects#connection, https://graphql.org/learn/pagination/), and expose reusable pagination helpers next to the existing REST ones (`WithPagination`, `OptionalPaginationParams`).

## Public API

```go
// The REST pagination parameters, now exported, plus a cursor.
type PaginationParams struct {
	Page    int
	PerPage int
	After   string
}

// ToGraphQLParams converts REST pagination parameters to GraphQL parameters.
// If After is set, it is used as the cursor.
func (p PaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error)

// Cursor-only pagination parameters (no page number).
type CursorPaginationParams struct {
	PerPage int
	After   string
}

func (p CursorPaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error)

type GraphQLPaginationParams struct {
	First *int32
	After *string
}

// OptionalCursorPaginationParams reads "perPage" (default 30) and "after" from the request.
func OptionalCursorPaginationParams(r mcp.CallToolRequest) (CursorPaginationParams, error)

// WithCursorPagination adds the "perPage" and "after" tool parameters.
func WithCursorPagination() mcp.ToolOption

// WithUnifiedPagination adds the "page", "perPage" and "after" tool parameters,
// for GraphQL tools that accept page/perPage and convert them to a cursor internally.
func WithUnifiedPagination() mcp.ToolOption

const DefaultGraphQLPageSize = 30
```

The existing `PaginationParams` fields become exported (`Page`, `PerPage`) and gain `After`. `OptionalPaginationParams` also reads an optional `after` string argument; its defaults stay page 1 and perPage 30. Update all existing callers.

Conversion to GraphQL parameters: `First` is `PerPage`; `After` is nil when the cursor is empty. A `PerPage` above 100 is an error (`perPage value %d exceeds maximum of 100`), and a negative one is an error (`perPage value %d cannot be negative`).

Tool parameter definitions:

- `perPage`: number, description `Results per page for pagination (min 1, max 100)`, minimum 1, maximum 100.
- `after`: string, description `Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs.`
- `page` (unified only): number, description `Page number for pagination (min 1)`, minimum 1.

## Tool changes

- `list_discussions` and `get_discussion_comments` accept `perPage` and `after`. Without them, a page of 30 is requested and no cursor is sent.
- `list_discussion_categories` drops its `first`, `last`, `after` and `before` parameters (and their validation errors). It always requests the first 25 categories.
- All three tools return an object instead of a bare array: the items under `discussions`, `comments` or `categories`, plus `pageInfo` (`hasNextPage`, `hasPreviousPage`, `startCursor`, `endCursor`) and `totalCount`, taken from the connection.

The GraphQL documents sent to the API must be exactly these (variables: `owner`, `repo`, `first` as Int, `after` as a nullable String that is sent as null when there is no cursor, `categoryId` as ID, `discussionNumber` as Int):

```
query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}
query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}
query($after:String$discussionNumber:Int!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussion(number: $discussionNumber){comments(first: $first, after: $after){nodes{body},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}
query($first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussionCategories(first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}
```

Regenerate the tool documentation in the README with the repository's docs generator. Update the tests.

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
