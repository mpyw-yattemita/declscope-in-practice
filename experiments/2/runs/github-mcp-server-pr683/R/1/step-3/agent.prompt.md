# list_discussions: sort by creation or update time, and return updatedAt and author

Callers of `list_discussions` cannot find recently active discussions or see who started a discussion without extra API calls. Please:

1. Add `updatedAt` and the author's login to each discussion that `list_discussions` returns. In the returned JSON (a go-github `Discussion`), set `updated_at` and `user.login` in addition to the fields already returned.
2. Add two optional parameters to `list_discussions`:
   - `orderBy`: string, enum `CREATED_AT`, `UPDATED_AT`, description `Order discussions by field. If provided, the 'direction' also needs to be provided.`
   - `direction`: string, enum `ASC`, `DESC`, description `Order direction.`

Ordering is used only when both `orderBy` and `direction` are given. If only one of them is given, the tool behaves as if neither was given. Category filtering and pagination keep working with and without ordering. The response keeps its current shape (`discussions`, `pageInfo`, `totalCount`).

The `githubv4` client needs the query shape at compile time, so the query with and without ordering, and with and without a category, cannot be built from one string at runtime. The GraphQL documents sent to the API must be exactly these. `orderByField` is a `DiscussionOrderField` and `orderByDirection` an `OrderDirection`; they are sent only when ordering is used:

```
query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}
query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}
query($after:String$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}
query($after:String$categoryId:ID!$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}
```

Filtering by author is out of scope (the discussions connection does not support it). Regenerate the tool documentation in the README with the repository's docs generator, and add tests for the orderings and for the one-parameter cases.

# Scope of the change

As part of this change, refactor package `pkg/github` where its structure makes the change awkward. You may move, rename and split existing declarations and files within that package. Keep the package's behaviour and its exported API unchanged, and keep its tests passing.

Nobody will answer questions while you work on this task. When a decision would normally go to the repository's owner, make the choice you would recommend, carry it out, and say in your final message what you chose and why.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/github
```
