# update_pull_request: mark a pull request as draft or ready for review

`create_pull_request` can create a draft, but no tool can mark a draft as ready for review, so an agent cannot merge a draft pull request without a person (issue #501). Please add a boolean `draft` parameter to `update_pull_request`, description `Mark pull request as draft (true) or ready for review (false)`.

In GraphQL, being a draft is separate from the pull request's state (a pull request can be closed as a draft), so this is a separate parameter rather than a new `state` value. Changing it needs the GraphQL API, so the tool now needs a GraphQL client as well:

```go
func UpdatePullRequest(getClient GetClientFn, getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc)
```

Update its registration in the toolset. The updated expected tool schema snapshot for `update_pull_request` is provided in the workspace.

## Behaviour

- The existing REST fields (`title`, `body`, `state`, `base`, `maintainer_can_modify`) are still sent with one REST edit request, and only if at least one of them is given. The error `No update parameters provided.` is returned only if none of them and no `draft` is given.
- If `draft` is given, the tool reads the pull request's node ID and current draft state, and changes it only if it differs:
  - query: `repository(owner: $owner, name: $repo) { pullRequest(number: $prNum) { id isDraft } }` with variables `owner`, `repo` (String) and `prNum` (Int);
  - to make it a draft: mutation `convertPullRequestToDraft(input: $input) { pullRequest { id isDraft } }` with `ConvertPullRequestToDraftInput{PullRequestID}`;
  - to mark it ready: mutation `markPullRequestReadyForReview(input: $input) { pullRequest { id isDraft } }` with `MarkPullRequestReadyForReviewInput{PullRequestID}`.
  GraphQL failures are returned as tool errors using the package's GraphQL error response helper, with the messages `Failed to find pull request`, `Failed to convert pull request to draft` and `Failed to mark pull request ready for review`.
- After the updates, the tool fetches the pull request again with a REST GET and returns that as JSON, so the result reflects both kinds of change. A REST edit failure still gives `failed to update pull request`; a failure of the final GET gives `Failed to get pull request`.

Regenerate the tool documentation in the README with the repository's docs generator, and add tests for the draft and ready-for-review cases, and for an update that combines REST fields with a final fetch.

The following files are part of the specification and are already in place:
- pkg/github/__toolsnaps__/update_pull_request.snap

Nobody will answer questions while you work on this task. When a decision would normally go to the repository's owner, make the choice you would recommend, carry it out, and say in your final message what you chose and why.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/github
```
