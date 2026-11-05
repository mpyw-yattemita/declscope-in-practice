This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/sub_issue.go:92: ListSubIssues declares its own `page` and `per_page` parameters, with no min/max, and reads them with `OptionalIntParamWithDefault(..., 30)`. Every other REST list tool uses `WithPagination()`, which exposes `perPage` with min 1 and max 100. So this new tool uses a different argument name (`per_page` instead of `perPage`) and accepts any number. Fix: Replace the two `mcp.WithNumber` options with `WithPagination()`. In the handler, use `pagination, err := OptionalPaginationParams(request)` and set `ListOptions{Page: pagination.Page, PerPage: pagination.PerPage}`. Change the test arguments from `per_page` to `perPage` and update the toolsnap.
1. pkg/pagination.go:20: Nothing outside the tests calls `WithUnifiedPagination`, `PaginationParams.After` or `PaginationParams.ToGraphQLParams`. The only GraphQL tools that page (discussions) use `WithCursorPagination` and `CursorPaginationParams`. Because of this unused path, every REST tool built with `OptionalPaginationParams` also reads an `after` argument that its schema never declares. Fix: Delete `WithUnifiedPagination`, the `After` field of `PaginationParams`, the `after` read in `OptionalPaginationParams`, and `PaginationParams.ToGraphQLParams`. Remove the matching cases from `pagination_test.go`: the after cases in TestOptionalPaginationParams, the REST loop in TestToGraphQLParams, and the WithUnifiedPagination row in TestPaginationToolOptions. Add them back when a tool actually needs them.
2. pkg/graphql.go:1: This new 13-line file holds only `DefaultGraphQLPageSize` and `GraphQLPaginationParams`. Its own comment says pagination.go builds the type. That splits one concept across two files, and `OptionalPaginationParams` still hard-codes its own `30` (`pagination.go:59`) next to the new constant. Fix: Move both declarations into pagination.go and delete graphql.go. Add one `DefaultPerPage = 30` constant and use it in both `OptionalPaginationParams` and `OptionalCursorPaginationParams`. Keep a separate GraphQL default only if the two defaults are meant to differ.
3. pkg/issues.go:410: ListIssues declares `WithPagination()` but reads `page` and `perPage` straight from `GetArguments()` with type assertions. That skips the shared helper, so values of the wrong type are silently ignored and the default page and perPage are not applied. Fix: Replace the two `if ... .(float64)` blocks with `pagination, err := OptionalPaginationParams(request)`, return a tool error if it fails, then set `opts.ListOptions = github.ListOptions{Page: pagination.Page, PerPage: pagination.PerPage}`.
4. pkg/repositories.go:143: This `if perPage == 0 { perPage = 30 }` block can never run, because `OptionalPaginationParams` already replaces 0 with 30. Fix: Delete the `perPage` local variable and the block, and use `PerPage: pagination.PerPage` directly.
5. pkg/actions.go:597: GetJobLogs sets the 500 default twice: once in the schema with `mcp.DefaultNumber(500)` and again by hand with `OptionalIntParam` and `if tailLines == 0`. The two copies can drift apart. Fix: Use `tailLines, err := OptionalIntParamWithDefault(request, "tail_lines", 500)` and remove the manual check. Better still, put 500 in a named constant that both the schema and the handler use.
6. pkg/actions.go:19: `DescriptionRepositoryOwner` and `DescriptionRepositoryName` are package-wide constants, but they are declared in actions.go and only actions.go uses them. The new and changed files (sub_issue.go:26,82,151,205 and discussions.go:25,186,259) still hard-code "Repository owner" and "Repository name". Fix: Move the constants to a shared file, such as server.go or a new params file. Use them for the `owner` and `repo` descriptions in sub_issue.go and discussions.go.

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
