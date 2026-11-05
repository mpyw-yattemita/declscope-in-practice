# search_code: rename the `q` parameter to `query`

`search_code` is the last tool that takes its search string as `q`; every other search tool uses `query`. Please rename the parameter to `query` (still required, same description: `Search query using GitHub code search syntax`). The search sent to the API is unchanged.

The updated expected tool schema snapshot for `search_code` is provided in the workspace. Regenerate the tool documentation in the README with the repository's docs generator, and update the tests.

The following files are part of the specification and are already in place:
- pkg/github/__toolsnaps__/search_code.snap

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/github
```
