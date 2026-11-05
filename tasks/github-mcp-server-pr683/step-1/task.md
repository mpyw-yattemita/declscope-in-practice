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
