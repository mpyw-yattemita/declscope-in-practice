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

# Scope of the change

Keep this change to what the task needs. Do not refactor existing code: do not move, rename or split existing declarations or files unless the task requires it.

Nobody will answer questions while you work on this task. When a decision would normally go to the repository's owner, make the choice you would recommend, carry it out, and say in your final message what you chose and why.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/github
```
