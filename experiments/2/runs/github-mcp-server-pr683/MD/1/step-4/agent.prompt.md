# search_code: rename the `q` parameter to `query`

`search_code` is the last tool that takes its search string as `q`; every other search tool uses `query`. Please rename the parameter to `query` (still required, same description: `Search query using GitHub code search syntax`). The search sent to the API is unchanged.

The updated expected tool schema snapshot for `search_code` is provided in the workspace. Regenerate the tool documentation in the README with the repository's docs generator, and update the tests.

The following files are part of the specification and are already in place:
- pkg/github/__toolsnaps__/search_code.snap

# declscope

This repository runs declscope, a linter that checks where each declaration of a Go package is used and how it is named. The `declscope-authoring` skill explains its rules and how to act on its reports. Read it before you start.

`.declscope.yaml` sets the rules, and `.declscope-baseline.yaml` records the reports the existing code already has, so only new reports are shown. Never edit either file, and never regenerate the baseline.

Before you finish, run these as well, and resolve what they report:

```bash
declscope ./...
declscope shrink ./...
```

# Scope of the change

Keep this change to what the task needs. Do not refactor existing code: do not move, rename or split existing declarations or files unless the task requires it.

Nobody will answer questions while you work on this task. When a decision would normally go to the repository's owner, make the choice you would recommend, carry it out, and say in your final message what you chose and why.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/github
declscope ./...
declscope shrink ./...
```
