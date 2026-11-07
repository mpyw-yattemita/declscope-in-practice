This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/actions.go:356: Nine handlers read run_id or artifact_id with RequiredInt and then convert with int64(runIDInt). params.go now has requiredInt64 for GitHub database IDs, which can be larger than 32 bits. Going through int adds a step that is not needed. Fix: Change each `runIDInt, err := RequiredInt(request, "run_id")` / `runID := int64(runIDInt)` pair to `runID, err := requiredInt64(request, "run_id")`. The lines are 356, 412, 483, 823, 886, 949, 1013, 1147 and 1210. Do the same for artifact_id at line 1081.
1. pkg/actions.go:576: GetJobLogs reads job_id and run_id with OptionalIntParam and then casts them with int64(jobID) and int64(runID) when it calls handleFailedJobLogs and handleSingleJobLogs. optionalInt64Param already does this. Fix: Use optionalInt64Param for job_id and run_id, and remove the int64(...) casts at the call sites.
2. pkg/pagination.go:54: Only pagination_test.go uses WithUnifiedPagination and PaginationParams.ToGraphQLParams (line 107). All GraphQL tools use WithCursorPagination with CursorPaginationParams. These functions add public API that no code calls. Fix: Delete WithUnifiedPagination and PaginationParams.ToGraphQLParams, along with their cases in TestPaginationToolOptions and the "rest/" subtests in TestToGraphQLParams. Add them back when a tool needs them.
3. pkg/pagination.go:93: OptionalPaginationParams now also reads "after". It is used only by REST tools, and WithPagination does not declare "after" for them. So PaginationParams.After is always empty in real calls, and a stray "after" argument can still make a REST tool fail with a type error. Fix: Remove the After field and the "after" lookup from PaginationParams and OptionalPaginationParams. Cursor tools already use OptionalCursorPaginationParams.
4. pkg/pagination.go:89: The REST default page size is the literal 30 here and again in sub_issues.go:57. The GraphQL default has a named constant, DefaultGraphQLPageSize. Fix: Add `DefaultRESTPageSize = 30` next to DefaultGraphQLPageSize. Use it in OptionalPaginationParams and in ListSubIssues.
5. pkg/repositories.go:143: ListCommits has a dead fallback, `if perPage == 0 { perPage = 30 }`. OptionalPaginationParams already returns 30 when perPage is 0 or missing. Fix: Remove the perPage local and the if block, and pass pagination.PerPage directly.
6. pkg/issues.go:410: ListIssues reads page and perPage with raw type assertions on request.GetArguments(). A wrong type is silently ignored, and the defaults differ from every other WithPagination tool. Fix: Call OptionalPaginationParams(request), return a tool error on failure, and set opts.ListOptions = github.ListOptions{Page: pagination.Page, PerPage: pagination.PerPage}.
7. pkg/discussions_test.go:64: The pageInfoOutput test type copies pageInfoResult from pagination.go field for field, including the JSON tags. Fix: Delete pageInfoOutput and use pageInfoResult in the test table and in the unmarshal targets.
8. pkg/discussions.go:130: discussions.go writes the owner and repo descriptions as "Repository owner" and "Repository name" string literals. actions.go defines the DescriptionRepositoryOwner and DescriptionRepositoryName constants for these, but no other file uses them. Fix: Move the constants to a shared file such as params.go, and use them in discussions.go and sub_issues.go (withSubIssueParentParams). If they will not be shared, remove them.
9. pkg/sub_issues.go:34: ListSubIssues declares owner, repo and issue_number inline, while the other three sub-issue tools use withSubIssueParentParams. The only difference is the issue_number description. Fix: Use withSubIssueParentParams() in ListSubIssues as well, and update its toolsnap. If the different description is needed, give the helper a description argument.
10. pkg/search.go:243: userOrOrgHandler sets TotalCount and IncompleteResults from result.GetTotal() and result.GetIncompleteResults(), then sets them again from the pointers in two if blocks. The second assignment changes nothing. Fix: Delete the `if result.Total != nil` and `if result.IncompleteResults != nil` blocks.
11. pkg/search_utils.go:55: The comment says the sort defaults to "created", but the code passes the user's sort value unchanged, and GitHub's default is best match. Fix: Delete the comment.
12. pkg/helper_test.go:135: getErrorResult does not call t.Helper(), unlike the helpers around it. A failed assertion is reported at the helper's line, not at the test that called it. Fix: Add t.Helper() as the first statement of getErrorResult.

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
