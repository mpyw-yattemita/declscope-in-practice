Two versions of the same Go package are in this directory, as `x/` and `y/`.
Both started from the same code and gained the same features, written by
different people. Compare how well each package is divided into files and types.

Judge the package as a whole, and how the new code fits into it:
- Does each file and each type have one clear job?
- Is closely related code kept together?
- Do the boundaries between files follow the boundaries between features?
- Does the new code fit the structure the package already had?

Behaviour and tests are known to be equivalent. Do not judge style, comments,
or naming conventions on their own, except where a name shows how the code is
divided.

Below is a map of each package: every file with its line count and its
top-level declarations. Files marked NEW were added and files marked CHANGED
were modified. The full text of every NEW and CHANGED file follows the maps.
You may read any other file under x/ or y/ if you need to.

### x/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| actions.go | 1230 | CHANGED | DescriptionRepositoryOwner (const), DescriptionRepositoryName (const), ListWorkflows (func), ListWorkflowRuns (func), RunWorkflow (func), GetWorkflowRun (func), GetWorkflowRunLogs (func), ListWorkflowJobs (func), GetJobLogs (func), handleFailedJobLogs (func), handleSingleJobLogs (func), getJobLogData (func), downloadLogContent (func), trimContent (func), RerunWorkflowRun (func), RerunFailedJobs (func), CancelWorkflowRun (func), ListWorkflowRunArtifacts (func), DownloadWorkflowRunArtifact (func), DeleteWorkflowRunLogs (func), GetWorkflowRunUsage (func) |
| actions_test.go | 1146 |  | Test_ListWorkflows (func), Test_RunWorkflow (func), Test_RunWorkflow_WithFilename (func), Test_CancelWorkflowRun (func), Test_ListWorkflowRunArtifacts (func), Test_DownloadWorkflowRunArtifact (func), Test_DeleteWorkflowRunLogs (func), Test_GetWorkflowRunUsage (func), Test_GetJobLogs (func), Test_GetJobLogs_WithContentReturn (func), Test_GetJobLogs_WithContentReturnAndTailLines (func) |
| code_scanning.go | 169 |  | GetCodeScanningAlert (func), ListCodeScanningAlerts (func) |
| code_scanning_test.go | 249 |  | Test_GetCodeScanningAlert (func), Test_ListCodeScanningAlerts (func) |
| context_tools.go | 92 |  | UserDetails (type), GetMe (func) |
| context_tools_test.go | 141 |  | Test_GetMe (func) |
| dependabot.go | 161 |  | GetDependabotAlert (func), ListDependabotAlerts (func) |
| dependabot_test.go | 276 |  | Test_GetDependabotAlert (func), Test_ListDependabotAlerts (func) |
| discussions.go | 414 | CHANGED | ListDiscussions (func), discussionsConnection (type), queryDiscussions (func), GetDiscussion (func), GetDiscussionComments (func), ListDiscussionCategories (func) |
| discussions_test.go | 560 | CHANGED | discussionsPageInfo (var), discussionsGeneral (var), discussionsAll (var), mockResponseListAll (var), mockResponseListGeneral (var), mockErrorRepoNotFound (var), Test_ListDiscussions (func), Test_GetDiscussion (func), Test_GetDiscussionComments (func), Test_ListDiscussionCategories (func) |
| dynamic_tools.go | 138 |  | ToolsetEnum (func), EnableToolset (func), ListAvailableToolsets (func), GetToolsetsTools (func) |
| helper_test.go | 277 | CHANGED | expectations (type), expect (func), expectPath (func), expectQueryParams (func), expectRequestBody (func), partialMock (type), partialMock.andThen (method), mockResponse (func), createMCPRequest (func), getTextResult (func), getErrorResult (func), getTextResourceResult (func), getBlobResourceResult (func), TestOptionalParamOK (func) |
| issues.go | 920 | CHANGED | GetIssue (func), AddIssueComment (func), SearchIssues (func), CreateIssue (func), ListIssues (func), UpdateIssue (func), GetIssueComments (func), mvpDescription (type), mvpDescription.String (method), AssignCopilotToIssue (func), ReplaceActorsForAssignableInput (type), parseISOTimestamp (func), AssignCodingAgentPrompt (func) |
| issues_test.go | 1631 |  | Test_GetIssue (func), Test_AddIssueComment (func), Test_SearchIssues (func), Test_CreateIssue (func), Test_ListIssues (func), Test_UpdateIssue (func), Test_ParseISOTimestamp (func), Test_GetIssueComments (func), TestAssignCopilotToIssue (func) |
| notifications.go | 525 | CHANGED | FilterDefault (const), FilterIncludeRead (const), FilterOnlyParticipating (const), ListNotifications (func), DismissNotification (func), MarkAllNotificationsRead (func), GetNotificationDetails (func), NotificationActionIgnore (const), NotificationActionWatch (const), NotificationActionDelete (const), ManageNotificationSubscription (func), RepositorySubscriptionActionWatch (const), RepositorySubscriptionActionIgnore (const), RepositorySubscriptionActionDelete (const), ManageRepositoryNotificationSubscription (func) |
| notifications_test.go | 765 |  | Test_ListNotifications (func), Test_ManageNotificationSubscription (func), Test_ManageRepositoryNotificationSubscription (func), Test_DismissNotification (func), Test_MarkAllNotificationsRead (func), Test_GetNotificationDetails (func) |
| pagination.go | 187 | NEW | DefaultGraphQLPageSize (const), PaginationParams (type), PaginationParams.ToGraphQLParams (method), CursorPaginationParams (type), CursorPaginationParams.ToGraphQLParams (method), GraphQLPaginationParams (type), GraphQLPaginationParams.queryVariables (method), paginationPageInfo (type), WithPagination (func), WithCursorPagination (func), WithUnifiedPagination (func), OptionalPaginationParams (func), OptionalCursorPaginationParams (func), paginationPageParam (func), paginationPerPageParam (func), paginationAfterParam (func) |
| pagination_test.go | 251 | NEW | TestOptionalPaginationParams (func), TestOptionalCursorPaginationParams (func), TestToGraphQLParams (func), TestPaginationToolOptions (func) |
| pullrequests.go | 1741 | CHANGED | GetPullRequest (func), CreatePullRequest (func), UpdatePullRequest (func), ListPullRequests (func), MergePullRequest (func), SearchPullRequests (func), GetPullRequestFiles (func), GetPullRequestStatus (func), UpdatePullRequestBranch (func), GetPullRequestComments (func), GetPullRequestReviews (func), CreateAndSubmitPullRequestReview (func), CreatePendingPullRequestReview (func), AddCommentToPendingReview (func), SubmitPendingPullRequestReview (func), DeletePendingPullRequestReview (func), GetPullRequestDiff (func), RequestCopilotReview (func), newGQLStringlike (func), newGQLStringlikePtr (func), newGQLIntPtr (func) |
| pullrequests_test.go | 2592 |  | Test_GetPullRequest (func), Test_UpdatePullRequest (func), Test_ListPullRequests (func), Test_MergePullRequest (func), Test_SearchPullRequests (func), Test_GetPullRequestFiles (func), Test_GetPullRequestStatus (func), Test_UpdatePullRequestBranch (func), Test_GetPullRequestComments (func), Test_GetPullRequestReviews (func), Test_CreatePullRequest (func), TestCreateAndSubmitPullRequestReview (func), Test_RequestCopilotReview (func), TestCreatePendingPullRequestReview (func), TestAddPullRequestReviewCommentToPendingReview (func), TestSubmitPendingPullRequestReview (func), TestDeletePendingPullRequestReview (func), TestGetPullRequestDiff (func), viewerQuery (func), getLatestPendingReviewQueryReview (type), getLatestPendingReviewQueryParams (type), getLatestPendingReviewQuery (func) |
| repositories.go | 1393 | CHANGED | GetCommit (func), ListCommits (func), ListBranches (func), CreateOrUpdateFile (func), CreateRepository (func), GetFileContents (func), ForkRepository (func), DeleteFile (func), CreateBranch (func), PushFiles (func), ListTags (func), GetTag (func), filterPaths (func), resolveGitReference (func) |
| repositories_test.go | 2281 |  | Test_GetFileContents (func), Test_ForkRepository (func), Test_CreateBranch (func), Test_GetCommit (func), Test_ListCommits (func), Test_CreateOrUpdateFile (func), Test_CreateRepository (func), Test_PushFiles (func), Test_ListBranches (func), Test_DeleteFile (func), Test_ListTags (func), Test_GetTag (func), Test_filterPaths (func), Test_resolveGitReference (func) |
| repository_resource.go | 192 |  | GetRepositoryResourceContent (func), GetRepositoryResourceBranchContent (func), GetRepositoryResourceCommitContent (func), GetRepositoryResourceTagContent (func), GetRepositoryResourcePrContent (func), RepositoryResourceContentsHandler (func) |
| repository_resource_test.go | 280 |  | Test_repositoryResourceContentsHandler (func), Test_GetRepositoryResourceContent (func), Test_GetRepositoryResourceBranchContent (func), Test_GetRepositoryResourceCommitContent (func), Test_GetRepositoryResourceTagContent (func) |
| search.go | 304 | CHANGED | SearchRepositories (func), SearchCode (func), MinimalUser (type), MinimalSearchUsersResult (type), userOrOrgHandler (func), SearchUsers (func), SearchOrgs (func) |
| search_test.go | 600 | CHANGED | Test_SearchRepositories (func), Test_SearchCode (func), Test_SearchUsers (func), Test_SearchOrgs (func) |
| search_utils.go | 88 | CHANGED | searchHandler (func) |
| secret_scanning.go | 163 |  | GetSecretScanningAlert (func), ListSecretScanningAlerts (func) |
| secret_scanning_test.go | 249 |  | Test_GetSecretScanningAlert (func), Test_ListSecretScanningAlerts (func) |
| server.go | 184 | CHANGED | NewServer (func), OptionalParamOK (func), isAcceptedError (func), RequiredParam (func), RequiredInt (func), OptionalParam (func), OptionalIntParam (func), OptionalIntParamWithDefault (func), OptionalStringArrayParam (func), MarshalledTextResult (func) |
| server_test.go | 479 | CHANGED | stubGetClientFn (func), stubGetClientFromHTTPFn (func), stubGetClientFnErr (func), stubGetGQLClientFn (func), stubGetRawClientFn (func), badRequestHandler (func), Test_IsAcceptedError (func), Test_RequiredStringParam (func), Test_OptionalStringParam (func), Test_RequiredInt (func), Test_OptionalIntParam (func), Test_OptionalNumberParamWithDefault (func), Test_OptionalBooleanParam (func), TestOptionalStringArrayParam (func) |
| sub_issues.go | 315 | NEW | ListSubIssues (func), AddSubIssue (func), RemoveSubIssue (func), ReprioritizeSubIssue (func), requiredSubIssueParent (func), subIssueResult (func) |
| sub_issues_test.go | 366 | NEW | subIssueTestCase (type), newSubIssueTool (type), runSubIssueTestCases (func), subIssueParentArgs (func), withoutSubIssueArg (func), subIssueMock (var), subIssueParentMock (var), Test_AddSubIssue (func), Test_ListSubIssues (func), Test_RemoveSubIssue (func), Test_ReprioritizeSubIssue (func) |
| tools.go | 209 | CHANGED | GetClientFn (type), GetGQLClientFn (type), DefaultTools (var), DefaultToolsetGroup (func), InitDynamicToolset (func), ToBoolPtr (func), ToStringPtr (func) |

### y/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| actions.go | 1230 | CHANGED | DescriptionRepositoryOwner (const), DescriptionRepositoryName (const), ListWorkflows (func), ListWorkflowRuns (func), RunWorkflow (func), GetWorkflowRun (func), GetWorkflowRunLogs (func), ListWorkflowJobs (func), GetJobLogs (func), handleFailedJobLogs (func), handleSingleJobLogs (func), getJobLogData (func), downloadLogContent (func), trimContent (func), RerunWorkflowRun (func), RerunFailedJobs (func), CancelWorkflowRun (func), ListWorkflowRunArtifacts (func), DownloadWorkflowRunArtifact (func), DeleteWorkflowRunLogs (func), GetWorkflowRunUsage (func) |
| actions_test.go | 1146 |  | Test_ListWorkflows (func), Test_RunWorkflow (func), Test_RunWorkflow_WithFilename (func), Test_CancelWorkflowRun (func), Test_ListWorkflowRunArtifacts (func), Test_DownloadWorkflowRunArtifact (func), Test_DeleteWorkflowRunLogs (func), Test_GetWorkflowRunUsage (func), Test_GetJobLogs (func), Test_GetJobLogs_WithContentReturn (func), Test_GetJobLogs_WithContentReturnAndTailLines (func) |
| code_scanning.go | 169 |  | GetCodeScanningAlert (func), ListCodeScanningAlerts (func) |
| code_scanning_test.go | 249 |  | Test_GetCodeScanningAlert (func), Test_ListCodeScanningAlerts (func) |
| context_tools.go | 92 |  | UserDetails (type), GetMe (func) |
| context_tools_test.go | 141 |  | Test_GetMe (func) |
| dependabot.go | 161 |  | GetDependabotAlert (func), ListDependabotAlerts (func) |
| dependabot_test.go | 276 |  | Test_GetDependabotAlert (func), Test_ListDependabotAlerts (func) |
| discussions.go | 437 | CHANGED | ListDiscussions (func), GetDiscussion (func), GetDiscussionComments (func), ListDiscussionCategories (func), discussionsPageInfo (type), discussionsConnection (type) |
| discussions_test.go | 641 | CHANGED | discussionsGeneral (var), discussionsAll (var), mockResponseListAll (var), mockResponseListGeneral (var), mockErrorRepoNotFound (var), Test_ListDiscussions (func), Test_GetDiscussion (func), Test_GetDiscussionComments (func), Test_ListDiscussionCategories (func) |
| dynamic_tools.go | 138 |  | ToolsetEnum (func), EnableToolset (func), ListAvailableToolsets (func), GetToolsetsTools (func) |
| helper_test.go | 275 |  | expectations (type), expect (func), expectPath (func), expectQueryParams (func), expectRequestBody (func), partialMock (type), partialMock.andThen (method), mockResponse (func), createMCPRequest (func), getTextResult (func), getErrorResult (func), getTextResourceResult (func), getBlobResourceResult (func), TestOptionalParamOK (func) |
| issues.go | 920 | CHANGED | GetIssue (func), AddIssueComment (func), SearchIssues (func), CreateIssue (func), ListIssues (func), UpdateIssue (func), GetIssueComments (func), mvpDescription (type), mvpDescription.String (method), AssignCopilotToIssue (func), ReplaceActorsForAssignableInput (type), parseISOTimestamp (func), AssignCodingAgentPrompt (func) |
| issues_test.go | 1631 |  | Test_GetIssue (func), Test_AddIssueComment (func), Test_SearchIssues (func), Test_CreateIssue (func), Test_ListIssues (func), Test_UpdateIssue (func), Test_ParseISOTimestamp (func), Test_GetIssueComments (func), TestAssignCopilotToIssue (func) |
| notifications.go | 525 | CHANGED | FilterDefault (const), FilterIncludeRead (const), FilterOnlyParticipating (const), ListNotifications (func), DismissNotification (func), MarkAllNotificationsRead (func), GetNotificationDetails (func), NotificationActionIgnore (const), NotificationActionWatch (const), NotificationActionDelete (const), ManageNotificationSubscription (func), RepositorySubscriptionActionWatch (const), RepositorySubscriptionActionIgnore (const), RepositorySubscriptionActionDelete (const), ManageRepositoryNotificationSubscription (func) |
| notifications_test.go | 765 |  | Test_ListNotifications (func), Test_ManageNotificationSubscription (func), Test_ManageRepositoryNotificationSubscription (func), Test_DismissNotification (func), Test_MarkAllNotificationsRead (func), Test_GetNotificationDetails (func) |
| pagination.go | 103 | NEW | DefaultGraphQLPageSize (const), CursorPaginationParams (type), CursorPaginationParams.ToGraphQLParams (method), GraphQLPaginationParams (type), OptionalCursorPaginationParams (func), WithCursorPagination (func), WithUnifiedPagination (func), withPaginationPerPage (func), withPaginationAfter (func) |
| pagination_test.go | 127 | NEW | TestCursorPaginationParams_ToGraphQLParams (func), TestOptionalCursorPaginationParams (func), TestWithCursorPagination (func), TestWithUnifiedPagination (func) |
| pullrequests.go | 1741 | CHANGED | GetPullRequest (func), CreatePullRequest (func), UpdatePullRequest (func), ListPullRequests (func), MergePullRequest (func), SearchPullRequests (func), GetPullRequestFiles (func), GetPullRequestStatus (func), UpdatePullRequestBranch (func), GetPullRequestComments (func), GetPullRequestReviews (func), CreateAndSubmitPullRequestReview (func), CreatePendingPullRequestReview (func), AddCommentToPendingReview (func), SubmitPendingPullRequestReview (func), DeletePendingPullRequestReview (func), GetPullRequestDiff (func), RequestCopilotReview (func), newGQLStringlike (func), newGQLStringlikePtr (func), newGQLIntPtr (func) |
| pullrequests_test.go | 2592 |  | Test_GetPullRequest (func), Test_UpdatePullRequest (func), Test_ListPullRequests (func), Test_MergePullRequest (func), Test_SearchPullRequests (func), Test_GetPullRequestFiles (func), Test_GetPullRequestStatus (func), Test_UpdatePullRequestBranch (func), Test_GetPullRequestComments (func), Test_GetPullRequestReviews (func), Test_CreatePullRequest (func), TestCreateAndSubmitPullRequestReview (func), Test_RequestCopilotReview (func), TestCreatePendingPullRequestReview (func), TestAddPullRequestReviewCommentToPendingReview (func), TestSubmitPendingPullRequestReview (func), TestDeletePendingPullRequestReview (func), TestGetPullRequestDiff (func), viewerQuery (func), getLatestPendingReviewQueryReview (type), getLatestPendingReviewQueryParams (type), getLatestPendingReviewQuery (func) |
| repositories.go | 1393 | CHANGED | GetCommit (func), ListCommits (func), ListBranches (func), CreateOrUpdateFile (func), CreateRepository (func), GetFileContents (func), ForkRepository (func), DeleteFile (func), CreateBranch (func), PushFiles (func), ListTags (func), GetTag (func), filterPaths (func), resolveGitReference (func) |
| repositories_test.go | 2281 |  | Test_GetFileContents (func), Test_ForkRepository (func), Test_CreateBranch (func), Test_GetCommit (func), Test_ListCommits (func), Test_CreateOrUpdateFile (func), Test_CreateRepository (func), Test_PushFiles (func), Test_ListBranches (func), Test_DeleteFile (func), Test_ListTags (func), Test_GetTag (func), Test_filterPaths (func), Test_resolveGitReference (func) |
| repository_resource.go | 192 |  | GetRepositoryResourceContent (func), GetRepositoryResourceBranchContent (func), GetRepositoryResourceCommitContent (func), GetRepositoryResourceTagContent (func), GetRepositoryResourcePrContent (func), RepositoryResourceContentsHandler (func) |
| repository_resource_test.go | 280 |  | Test_repositoryResourceContentsHandler (func), Test_GetRepositoryResourceContent (func), Test_GetRepositoryResourceBranchContent (func), Test_GetRepositoryResourceCommitContent (func), Test_GetRepositoryResourceTagContent (func) |
| search.go | 304 | CHANGED | SearchRepositories (func), SearchCode (func), MinimalUser (type), MinimalSearchUsersResult (type), userOrOrgHandler (func), SearchUsers (func), SearchOrgs (func) |
| search_test.go | 600 | CHANGED | Test_SearchRepositories (func), Test_SearchCode (func), Test_SearchUsers (func), Test_SearchOrgs (func) |
| search_utils.go | 88 | CHANGED | searchHandler (func) |
| secret_scanning.go | 163 |  | GetSecretScanningAlert (func), ListSecretScanningAlerts (func) |
| secret_scanning_test.go | 249 |  | Test_GetSecretScanningAlert (func), Test_ListSecretScanningAlerts (func) |
| server.go | 246 | CHANGED | NewServer (func), OptionalParamOK (func), isAcceptedError (func), RequiredParam (func), RequiredInt (func), OptionalParam (func), OptionalIntParam (func), OptionalIntParamWithDefault (func), OptionalStringArrayParam (func), WithPagination (func), PaginationParams (type), PaginationParams.ToGraphQLParams (method), OptionalPaginationParams (func), MarshalledTextResult (func) |
| server_test.go | 627 | CHANGED | stubGetClientFn (func), stubGetClientFromHTTPFn (func), stubGetClientFnErr (func), stubGetGQLClientFn (func), stubGetRawClientFn (func), badRequestHandler (func), Test_IsAcceptedError (func), Test_RequiredStringParam (func), Test_OptionalStringParam (func), Test_RequiredInt (func), Test_OptionalIntParam (func), Test_OptionalNumberParamWithDefault (func), Test_OptionalBooleanParam (func), TestOptionalStringArrayParam (func), TestOptionalPaginationParams (func), TestPaginationParams_ToGraphQLParams (func) |
| sub_issue.go | 347 | NEW | AddSubIssue (func), ListSubIssues (func), RemoveSubIssue (func), ReprioritizeSubIssue (func) |
| sub_issue_test.go | 478 | NEW | Test_AddSubIssue (func), Test_ListSubIssues (func), Test_RemoveSubIssue (func), Test_ReprioritizeSubIssue (func) |
| tools.go | 209 | CHANGED | GetClientFn (type), GetGQLClientFn (type), DefaultTools (var), DefaultToolsetGroup (func), InitDynamicToolset (func), ToBoolPtr (func), ToStringPtr (func) |

#### x/actions.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	DescriptionRepositoryOwner = "Repository owner"
	DescriptionRepositoryName  = "Repository name"
)

// ListWorkflows creates a tool to list workflows in a repository
func ListWorkflows(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_workflows",
			mcp.WithDescription(t("TOOL_LIST_WORKFLOWS_DESCRIPTION", "List workflows in a repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_WORKFLOWS_USER_TITLE", "List workflows"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional pagination parameters
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Set up list options
			opts := &github.ListOptions{
				PerPage: pagination.PerPage,
				Page:    pagination.Page,
			}

			workflows, resp, err := client.Actions.ListWorkflows(ctx, owner, repo, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to list workflows: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(workflows)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListWorkflowRuns creates a tool to list workflow runs for a specific workflow
func ListWorkflowRuns(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_workflow_runs",
			mcp.WithDescription(t("TOOL_LIST_WORKFLOW_RUNS_DESCRIPTION", "List workflow runs for a specific workflow")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_WORKFLOW_RUNS_USER_TITLE", "List workflow runs"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithString("workflow_id",
				mcp.Required(),
				mcp.Description("The workflow ID or workflow file name"),
			),
			mcp.WithString("actor",
				mcp.Description("Returns someone's workflow runs. Use the login for the user who created the workflow run."),
			),
			mcp.WithString("branch",
				mcp.Description("Returns workflow runs associated with a branch. Use the name of the branch."),
			),
			mcp.WithString("event",
				mcp.Description("Returns workflow runs for a specific event type"),
				mcp.Enum(
					"branch_protection_rule",
					"check_run",
					"check_suite",
					"create",
					"delete",
					"deployment",
					"deployment_status",
					"discussion",
					"discussion_comment",
					"fork",
					"gollum",
					"issue_comment",
					"issues",
					"label",
					"merge_group",
					"milestone",
					"page_build",
					"public",
					"pull_request",
					"pull_request_review",
					"pull_request_review_comment",
					"pull_request_target",
					"push",
					"registry_package",
					"release",
					"repository_dispatch",
					"schedule",
					"status",
					"watch",
					"workflow_call",
					"workflow_dispatch",
					"workflow_run",
				),
			),
			mcp.WithString("status",
				mcp.Description("Returns workflow runs with the check run status"),
				mcp.Enum("queued", "in_progress", "completed", "requested", "waiting"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			workflowID, err := RequiredParam[string](request, "workflow_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional filtering parameters
			actor, err := OptionalParam[string](request, "actor")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := OptionalParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			event, err := OptionalParam[string](request, "event")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			status, err := OptionalParam[string](request, "status")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional pagination parameters
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Set up list options
			opts := &github.ListWorkflowRunsOptions{
				Actor:  actor,
				Branch: branch,
				Event:  event,
				Status: status,
				ListOptions: github.ListOptions{
					PerPage: pagination.PerPage,
					Page:    pagination.Page,
				},
			}

			workflowRuns, resp, err := client.Actions.ListWorkflowRunsByFileName(ctx, owner, repo, workflowID, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to list workflow runs: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(workflowRuns)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// RunWorkflow creates a tool to run an Actions workflow
func RunWorkflow(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("run_workflow",
			mcp.WithDescription(t("TOOL_RUN_WORKFLOW_DESCRIPTION", "Run an Actions workflow by workflow ID or filename")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_RUN_WORKFLOW_USER_TITLE", "Run workflow"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithString("workflow_id",
				mcp.Required(),
				mcp.Description("The workflow ID (numeric) or workflow file name (e.g., main.yml, ci.yaml)"),
			),
			mcp.WithString("ref",
				mcp.Required(),
				mcp.Description("The git reference for the workflow. The reference can be a branch or tag name."),
			),
			mcp.WithObject("inputs",
				mcp.Description("Inputs the workflow accepts"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			workflowID, err := RequiredParam[string](request, "workflow_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			ref, err := RequiredParam[string](request, "ref")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional inputs parameter
			var inputs map[string]interface{}
			if requestInputs, ok := request.GetArguments()["inputs"]; ok {
				if inputsMap, ok := requestInputs.(map[string]interface{}); ok {
					inputs = inputsMap
				}
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			event := github.CreateWorkflowDispatchEventRequest{
				Ref:    ref,
				Inputs: inputs,
			}

			var resp *github.Response
			var workflowType string

			if workflowIDInt, parseErr := strconv.ParseInt(workflowID, 10, 64); parseErr == nil {
				resp, err = client.Actions.CreateWorkflowDispatchEventByID(ctx, owner, repo, workflowIDInt, event)
				workflowType = "workflow_id"
			} else {
				resp, err = client.Actions.CreateWorkflowDispatchEventByFileName(ctx, owner, repo, workflowID, event)
				workflowType = "workflow_file"
			}

			if err != nil {
				return nil, fmt.Errorf("failed to run workflow: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":       "Workflow run has been queued",
				"workflow_type": workflowType,
				"workflow_id":   workflowID,
				"ref":           ref,
				"inputs":        inputs,
				"status":        resp.Status,
				"status_code":   resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetWorkflowRun creates a tool to get details of a specific workflow run
func GetWorkflowRun(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_workflow_run",
			mcp.WithDescription(t("TOOL_GET_WORKFLOW_RUN_DESCRIPTION", "Get details of a specific workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_WORKFLOW_RUN_USER_TITLE", "Get workflow run"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			workflowRun, resp, err := client.Actions.GetWorkflowRunByID(ctx, owner, repo, runID)
			if err != nil {
				return nil, fmt.Errorf("failed to get workflow run: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(workflowRun)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetWorkflowRunLogs creates a tool to download logs for a specific workflow run
func GetWorkflowRunLogs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_workflow_run_logs",
			mcp.WithDescription(t("TOOL_GET_WORKFLOW_RUN_LOGS_DESCRIPTION", "Download logs for a specific workflow run (EXPENSIVE: downloads ALL logs as ZIP. Consider using get_job_logs with failed_only=true for debugging failed jobs)")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_WORKFLOW_RUN_LOGS_USER_TITLE", "Get workflow run logs"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the download URL for the logs
			url, resp, err := client.Actions.GetWorkflowRunLogs(ctx, owner, repo, runID, 1)
			if err != nil {
				return nil, fmt.Errorf("failed to get workflow run logs: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			// Create response with the logs URL and information
			result := map[string]any{
				"logs_url":         url.String(),
				"message":          "Workflow run logs are available for download",
				"note":             "The logs_url provides a download link for the complete workflow run logs as a ZIP archive. You can download this archive to extract and examine individual job logs.",
				"warning":          "This downloads ALL logs as a ZIP file which can be large and expensive. For debugging failed jobs, consider using get_job_logs with failed_only=true and run_id instead.",
				"optimization_tip": "Use: get_job_logs with parameters {run_id: " + fmt.Sprintf("%d", runID) + ", failed_only: true} for more efficient failed job debugging",
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListWorkflowJobs creates a tool to list jobs for a specific workflow run
func ListWorkflowJobs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_workflow_jobs",
			mcp.WithDescription(t("TOOL_LIST_WORKFLOW_JOBS_DESCRIPTION", "List jobs for a specific workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_WORKFLOW_JOBS_USER_TITLE", "List workflow jobs"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
			mcp.WithString("filter",
				mcp.Description("Filters jobs by their completed_at timestamp"),
				mcp.Enum("latest", "all"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			// Get optional filtering parameters
			filter, err := OptionalParam[string](request, "filter")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional pagination parameters
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Set up list options
			opts := &github.ListWorkflowJobsOptions{
				Filter: filter,
				ListOptions: github.ListOptions{
					PerPage: pagination.PerPage,
					Page:    pagination.Page,
				},
			}

			jobs, resp, err := client.Actions.ListWorkflowJobs(ctx, owner, repo, runID, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to list workflow jobs: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			// Add optimization tip for failed job debugging
			response := map[string]any{
				"jobs":             jobs,
				"optimization_tip": "For debugging failed jobs, consider using get_job_logs with failed_only=true and run_id=" + fmt.Sprintf("%d", runID) + " to get logs directly without needing to list jobs first",
			}

			r, err := json.Marshal(response)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetJobLogs creates a tool to download logs for a specific workflow job or efficiently get all failed job logs for a workflow run
func GetJobLogs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_job_logs",
			mcp.WithDescription(t("TOOL_GET_JOB_LOGS_DESCRIPTION", "Download logs for a specific workflow job or efficiently get all failed job logs for a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_JOB_LOGS_USER_TITLE", "Get job logs"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("job_id",
				mcp.Description("The unique identifier of the workflow job (required for single job logs)"),
			),
			mcp.WithNumber("run_id",
				mcp.Description("Workflow run ID (required when using failed_only)"),
			),
			mcp.WithBoolean("failed_only",
				mcp.Description("When true, gets logs for all failed jobs in run_id"),
			),
			mcp.WithBoolean("return_content",
				mcp.Description("Returns actual log content instead of URLs"),
			),
			mcp.WithNumber("tail_lines",
				mcp.Description("Number of lines to return from the end of the log"),
				mcp.DefaultNumber(500),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional parameters
			jobID, err := OptionalIntParam(request, "job_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID, err := OptionalIntParam(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			failedOnly, err := OptionalParam[bool](request, "failed_only")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			returnContent, err := OptionalParam[bool](request, "return_content")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			tailLines, err := OptionalIntParam(request, "tail_lines")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			// Default to 500 lines if not specified
			if tailLines == 0 {
				tailLines = 500
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Validate parameters
			if failedOnly && runID == 0 {
				return mcp.NewToolResultError("run_id is required when failed_only is true"), nil
			}
			if !failedOnly && jobID == 0 {
				return mcp.NewToolResultError("job_id is required when failed_only is false"), nil
			}

			if failedOnly && runID > 0 {
				// Handle failed-only mode: get logs for all failed jobs in the workflow run
				return handleFailedJobLogs(ctx, client, owner, repo, int64(runID), returnContent, tailLines)
			} else if jobID > 0 {
				// Handle single job mode
				return handleSingleJobLogs(ctx, client, owner, repo, int64(jobID), returnContent, tailLines)
			}

			return mcp.NewToolResultError("Either job_id must be provided for single job logs, or run_id with failed_only=true for failed job logs"), nil
		}
}

// handleFailedJobLogs gets logs for all failed jobs in a workflow run
func handleFailedJobLogs(ctx context.Context, client *github.Client, owner, repo string, runID int64, returnContent bool, tailLines int) (*mcp.CallToolResult, error) {
	// First, get all jobs for the workflow run
	jobs, resp, err := client.Actions.ListWorkflowJobs(ctx, owner, repo, runID, &github.ListWorkflowJobsOptions{
		Filter: "latest",
	})
	if err != nil {
		return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to list workflow jobs", resp, err), nil
	}
	defer func() { _ = resp.Body.Close() }()

	// Filter for failed jobs
	var failedJobs []*github.WorkflowJob
	for _, job := range jobs.Jobs {
		if job.GetConclusion() == "failure" {
			failedJobs = append(failedJobs, job)
		}
	}

	if len(failedJobs) == 0 {
		result := map[string]any{
			"message":     "No failed jobs found in this workflow run",
			"run_id":      runID,
			"total_jobs":  len(jobs.Jobs),
			"failed_jobs": 0,
		}
		r, _ := json.Marshal(result)
		return mcp.NewToolResultText(string(r)), nil
	}

	// Collect logs for all failed jobs
	var logResults []map[string]any
	for _, job := range failedJobs {
		jobResult, resp, err := getJobLogData(ctx, client, owner, repo, job.GetID(), job.GetName(), returnContent, tailLines)
		if err != nil {
			// Continue with other jobs even if one fails
			jobResult = map[string]any{
				"job_id":   job.GetID(),
				"job_name": job.GetName(),
				"error":    err.Error(),
			}
			// Enable reporting of status codes and error causes
			_, _ = ghErrors.NewGitHubAPIErrorToCtx(ctx, "failed to get job logs", resp, err) // Explicitly ignore error for graceful handling
		}

		logResults = append(logResults, jobResult)
	}

	result := map[string]any{
		"message":       fmt.Sprintf("Retrieved logs for %d failed jobs", len(failedJobs)),
		"run_id":        runID,
		"total_jobs":    len(jobs.Jobs),
		"failed_jobs":   len(failedJobs),
		"logs":          logResults,
		"return_format": map[string]bool{"content": returnContent, "urls": !returnContent},
	}

	r, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	return mcp.NewToolResultText(string(r)), nil
}

// handleSingleJobLogs gets logs for a single job
func handleSingleJobLogs(ctx context.Context, client *github.Client, owner, repo string, jobID int64, returnContent bool, tailLines int) (*mcp.CallToolResult, error) {
	jobResult, resp, err := getJobLogData(ctx, client, owner, repo, jobID, "", returnContent, tailLines)
	if err != nil {
		return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to get job logs", resp, err), nil
	}

	r, err := json.Marshal(jobResult)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	return mcp.NewToolResultText(string(r)), nil
}

// getJobLogData retrieves log data for a single job, either as URL or content
func getJobLogData(ctx context.Context, client *github.Client, owner, repo string, jobID int64, jobName string, returnContent bool, tailLines int) (map[string]any, *github.Response, error) {
	// Get the download URL for the job logs
	url, resp, err := client.Actions.GetWorkflowJobLogs(ctx, owner, repo, jobID, 1)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get job logs for job %d: %w", jobID, err)
	}
	defer func() { _ = resp.Body.Close() }()

	result := map[string]any{
		"job_id": jobID,
	}
	if jobName != "" {
		result["job_name"] = jobName
	}

	if returnContent {
		// Download and return the actual log content
		content, originalLength, httpResp, err := downloadLogContent(url.String(), tailLines) //nolint:bodyclose // Response body is closed in downloadLogContent, but we need to return httpResp
		if err != nil {
			// To keep the return value consistent wrap the response as a GitHub Response
			ghRes := &github.Response{
				Response: httpResp,
			}
			return nil, ghRes, fmt.Errorf("failed to download log content for job %d: %w", jobID, err)
		}
		result["logs_content"] = content
		result["message"] = "Job logs content retrieved successfully"
		result["original_length"] = originalLength
	} else {
		// Return just the URL
		result["logs_url"] = url.String()
		result["message"] = "Job logs are available for download"
		result["note"] = "The logs_url provides a download link for the individual job logs in plain text format. Use return_content=true to get the actual log content."
	}

	return result, resp, nil
}

// downloadLogContent downloads the actual log content from a GitHub logs URL
func downloadLogContent(logURL string, tailLines int) (string, int, *http.Response, error) {
	httpResp, err := http.Get(logURL) //nolint:gosec // URLs are provided by GitHub API and are safe
	if err != nil {
		return "", 0, httpResp, fmt.Errorf("failed to download logs: %w", err)
	}
	defer func() { _ = httpResp.Body.Close() }()

	if httpResp.StatusCode != http.StatusOK {
		return "", 0, httpResp, fmt.Errorf("failed to download logs: HTTP %d", httpResp.StatusCode)
	}

	content, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return "", 0, httpResp, fmt.Errorf("failed to read log content: %w", err)
	}

	// Clean up and format the log content for better readability
	logContent := strings.TrimSpace(string(content))

	trimmedContent, lineCount := trimContent(logContent, tailLines)
	return trimmedContent, lineCount, httpResp, nil
}

// trimContent trims the content to a maximum length and returns the trimmed content and an original length
func trimContent(content string, tailLines int) (string, int) {
	// Truncate to tail_lines if specified
	lineCount := 0
	if tailLines > 0 {

		// Count backwards to find the nth newline from the end and a total number of lines
		for i := len(content) - 1; i >= 0 && lineCount < tailLines; i-- {
			if content[i] == '\n' {
				lineCount++
				// If we have reached the tailLines, trim the content
				if lineCount == tailLines {
					content = content[i+1:]
				}
			}
		}
	}
	return content, lineCount
}

// RerunWorkflowRun creates a tool to re-run an entire workflow run
func RerunWorkflowRun(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("rerun_workflow_run",
			mcp.WithDescription(t("TOOL_RERUN_WORKFLOW_RUN_DESCRIPTION", "Re-run an entire workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_RERUN_WORKFLOW_RUN_USER_TITLE", "Rerun workflow run"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			resp, err := client.Actions.RerunWorkflowByID(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to rerun workflow run", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":     "Workflow run has been queued for re-run",
				"run_id":      runID,
				"status":      resp.Status,
				"status_code": resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// RerunFailedJobs creates a tool to re-run only the failed jobs in a workflow run
func RerunFailedJobs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("rerun_failed_jobs",
			mcp.WithDescription(t("TOOL_RERUN_FAILED_JOBS_DESCRIPTION", "Re-run only the failed jobs in a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_RERUN_FAILED_JOBS_USER_TITLE", "Rerun failed jobs"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			resp, err := client.Actions.RerunFailedJobsByID(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to rerun failed jobs", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":     "Failed jobs have been queued for re-run",
				"run_id":      runID,
				"status":      resp.Status,
				"status_code": resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CancelWorkflowRun creates a tool to cancel a workflow run
func CancelWorkflowRun(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("cancel_workflow_run",
			mcp.WithDescription(t("TOOL_CANCEL_WORKFLOW_RUN_DESCRIPTION", "Cancel a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CANCEL_WORKFLOW_RUN_USER_TITLE", "Cancel workflow run"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			resp, err := client.Actions.CancelWorkflowRunByID(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to cancel workflow run", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":     "Workflow run has been cancelled",
				"run_id":      runID,
				"status":      resp.Status,
				"status_code": resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListWorkflowRunArtifacts creates a tool to list artifacts for a workflow run
func ListWorkflowRunArtifacts(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_workflow_run_artifacts",
			mcp.WithDescription(t("TOOL_LIST_WORKFLOW_RUN_ARTIFACTS_DESCRIPTION", "List artifacts for a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_WORKFLOW_RUN_ARTIFACTS_USER_TITLE", "List workflow artifacts"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			// Get optional pagination parameters
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Set up list options
			opts := &github.ListOptions{
				PerPage: pagination.PerPage,
				Page:    pagination.Page,
			}

			artifacts, resp, err := client.Actions.ListWorkflowRunArtifacts(ctx, owner, repo, runID, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to list workflow run artifacts", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(artifacts)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// DownloadWorkflowRunArtifact creates a tool to download a workflow run artifact
func DownloadWorkflowRunArtifact(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("download_workflow_run_artifact",
			mcp.WithDescription(t("TOOL_DOWNLOAD_WORKFLOW_RUN_ARTIFACT_DESCRIPTION", "Get download URL for a workflow run artifact")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_DOWNLOAD_WORKFLOW_RUN_ARTIFACT_USER_TITLE", "Download workflow artifact"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("artifact_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the artifact"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			artifactIDInt, err := RequiredInt(request, "artifact_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			artifactID := int64(artifactIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the download URL for the artifact
			url, resp, err := client.Actions.DownloadArtifact(ctx, owner, repo, artifactID, 1)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to get artifact download URL", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Create response with the download URL and information
			result := map[string]any{
				"download_url": url.String(),
				"message":      "Artifact is available for download",
				"note":         "The download_url provides a download link for the artifact as a ZIP archive. The link is temporary and expires after a short time.",
				"artifact_id":  artifactID,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// DeleteWorkflowRunLogs creates a tool to delete logs for a workflow run
func DeleteWorkflowRunLogs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("delete_workflow_run_logs",
			mcp.WithDescription(t("TOOL_DELETE_WORKFLOW_RUN_LOGS_DESCRIPTION", "Delete logs for a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:           t("TOOL_DELETE_WORKFLOW_RUN_LOGS_USER_TITLE", "Delete workflow logs"),
				ReadOnlyHint:    ToBoolPtr(false),
				DestructiveHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			resp, err := client.Actions.DeleteWorkflowRunLogs(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to delete workflow run logs", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":     "Workflow run logs have been deleted",
				"run_id":      runID,
				"status":      resp.Status,
				"status_code": resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetWorkflowRunUsage creates a tool to get usage metrics for a workflow run
func GetWorkflowRunUsage(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_workflow_run_usage",
			mcp.WithDescription(t("TOOL_GET_WORKFLOW_RUN_USAGE_DESCRIPTION", "Get usage metrics for a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_WORKFLOW_RUN_USAGE_USER_TITLE", "Get workflow usage"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			usage, resp, err := client.Actions.GetWorkflowRunUsageByID(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to get workflow run usage", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(usage)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

```

#### x/discussions.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

func ListDiscussions(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_discussions",
			mcp.WithDescription(t("TOOL_LIST_DISCUSSIONS_DESCRIPTION", "List discussions for a repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_DISCUSSIONS_USER_TITLE", "List discussions"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("category",
				mcp.Description("Optional filter by discussion category ID. If provided, only discussions with this category are listed."),
			),
			mcp.WithString("orderBy",
				mcp.Description("Order discussions by field. If provided, the 'direction' also needs to be provided."),
				mcp.Enum("CREATED_AT", "UPDATED_AT"),
			),
			mcp.WithString("direction",
				mcp.Description("Order direction."),
				mcp.Enum("ASC", "DESC"),
			),
			WithCursorPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Required params
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Optional params
			category, err := OptionalParam[string](request, "category")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			orderBy, err := OptionalParam[string](request, "orderBy")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			direction, err := OptionalParam[string](request, "direction")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			pagination, err := OptionalCursorPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			graphQLParams, err := pagination.ToGraphQLParams()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			vars := graphQLParams.queryVariables()
			vars["owner"] = githubv4.String(owner)
			vars["repo"] = githubv4.String(repo)
			// If a category filter is specified, use it as the category ID for server-side filtering.
			if category != "" {
				vars["categoryId"] = githubv4.ID(category)
			}
			// Ordering needs both the field and the direction. With only one, the default order is kept.
			ordered := orderBy != "" && direction != ""
			if ordered {
				vars["orderByField"] = githubv4.DiscussionOrderField(orderBy)
				vars["orderByDirection"] = githubv4.OrderDirection(direction)
			}

			conn, err := queryDiscussions(ctx, client, vars, category != "", ordered)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Map nodes to GitHub Discussion objects
			discussions := make([]*github.Discussion, 0, len(conn.Nodes))
			for _, n := range conn.Nodes {
				discussion := &github.Discussion{
					Number:    github.Ptr(int(n.Number)),
					Title:     github.Ptr(string(n.Title)),
					HTMLURL:   github.Ptr(string(n.URL)),
					CreatedAt: &github.Timestamp{Time: n.CreatedAt.Time},
					UpdatedAt: &github.Timestamp{Time: n.UpdatedAt.Time},
					DiscussionCategory: &github.DiscussionCategory{
						Name: github.Ptr(string(n.Category.Name)),
					},
				}
				// The author is null when the account was deleted.
				if n.Author.Login != "" {
					discussion.User = &github.User{Login: github.Ptr(string(n.Author.Login))}
				}
				discussions = append(discussions, discussion)
			}

			// Marshal and return
			out, err := json.Marshal(map[string]any{
				"discussions": discussions,
				"pageInfo":    conn.PageInfo,
				"totalCount":  int(conn.TotalCount),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussions: %w", err)
			}
			return mcp.NewToolResultText(string(out)), nil
		}
}

// discussionsConnection is the page of discussions that list_discussions reads.
type discussionsConnection struct {
	Nodes []struct {
		Number    githubv4.Int
		Title     githubv4.String
		CreatedAt githubv4.DateTime
		UpdatedAt githubv4.DateTime
		Author    struct {
			Login githubv4.String
		}
		Category struct {
			Name githubv4.String
		} `graphql:"category"`
		URL githubv4.String `graphql:"url"`
	}
	PageInfo   paginationPageInfo
	TotalCount githubv4.Int
}

// queryDiscussions runs the discussions query that matches the arguments in vars.
// githubv4 builds a document from the type of its query, so each combination of
// a category filter and an ordering needs its own query type.
func queryDiscussions(ctx context.Context, client *githubv4.Client, vars map[string]any, byCategory, ordered bool) (discussionsConnection, error) {
	switch {
	case byCategory && ordered:
		var q struct {
			Repository struct {
				Discussions discussionsConnection `graphql:"discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection })"`
			} `graphql:"repository(owner: $owner, name: $repo)"`
		}
		err := client.Query(ctx, &q, vars)
		return q.Repository.Discussions, err
	case byCategory:
		var q struct {
			Repository struct {
				Discussions discussionsConnection `graphql:"discussions(first: $first, after: $after, categoryId: $categoryId)"`
			} `graphql:"repository(owner: $owner, name: $repo)"`
		}
		err := client.Query(ctx, &q, vars)
		return q.Repository.Discussions, err
	case ordered:
		var q struct {
			Repository struct {
				Discussions discussionsConnection `graphql:"discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection })"`
			} `graphql:"repository(owner: $owner, name: $repo)"`
		}
		err := client.Query(ctx, &q, vars)
		return q.Repository.Discussions, err
	default:
		var q struct {
			Repository struct {
				Discussions discussionsConnection `graphql:"discussions(first: $first, after: $after)"`
			} `graphql:"repository(owner: $owner, name: $repo)"`
		}
		err := client.Query(ctx, &q, vars)
		return q.Repository.Discussions, err
	}
}

func GetDiscussion(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_discussion",
			mcp.WithDescription(t("TOOL_GET_DISCUSSION_DESCRIPTION", "Get a specific discussion by ID")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_DISCUSSION_USER_TITLE", "Get discussion"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("discussionNumber",
				mcp.Required(),
				mcp.Description("Discussion Number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner            string
				Repo             string
				DiscussionNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var q struct {
				Repository struct {
					Discussion struct {
						Number    githubv4.Int
						Body      githubv4.String
						CreatedAt githubv4.DateTime
						URL       githubv4.String `graphql:"url"`
						Category  struct {
							Name githubv4.String
						} `graphql:"category"`
					} `graphql:"discussion(number: $discussionNumber)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			vars := map[string]interface{}{
				"owner":            githubv4.String(params.Owner),
				"repo":             githubv4.String(params.Repo),
				"discussionNumber": githubv4.Int(params.DiscussionNumber),
			}
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			d := q.Repository.Discussion
			discussion := &github.Discussion{
				Number:    github.Ptr(int(d.Number)),
				Body:      github.Ptr(string(d.Body)),
				HTMLURL:   github.Ptr(string(d.URL)),
				CreatedAt: &github.Timestamp{Time: d.CreatedAt.Time},
				DiscussionCategory: &github.DiscussionCategory{
					Name: github.Ptr(string(d.Category.Name)),
				},
			}
			out, err := json.Marshal(discussion)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussion: %w", err)
			}

			return mcp.NewToolResultText(string(out)), nil
		}
}

func GetDiscussionComments(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_discussion_comments",
			mcp.WithDescription(t("TOOL_GET_DISCUSSION_COMMENTS_DESCRIPTION", "Get comments from a discussion")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_DISCUSSION_COMMENTS_USER_TITLE", "Get discussion comments"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner", mcp.Required(), mcp.Description("Repository owner")),
			mcp.WithString("repo", mcp.Required(), mcp.Description("Repository name")),
			mcp.WithNumber("discussionNumber", mcp.Required(), mcp.Description("Discussion Number")),
			WithCursorPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner            string
				Repo             string
				DiscussionNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			pagination, err := OptionalCursorPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			graphQLParams, err := pagination.ToGraphQLParams()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var q struct {
				Repository struct {
					Discussion struct {
						Comments struct {
							Nodes []struct {
								Body githubv4.String
							}
							PageInfo   paginationPageInfo
							TotalCount githubv4.Int
						} `graphql:"comments(first: $first, after: $after)"`
					} `graphql:"discussion(number: $discussionNumber)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			vars := graphQLParams.queryVariables()
			vars["owner"] = githubv4.String(params.Owner)
			vars["repo"] = githubv4.String(params.Repo)
			vars["discussionNumber"] = githubv4.Int(params.DiscussionNumber)
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			conn := q.Repository.Discussion.Comments
			comments := make([]*github.IssueComment, 0, len(conn.Nodes))
			for _, c := range conn.Nodes {
				comments = append(comments, &github.IssueComment{Body: github.Ptr(string(c.Body))})
			}

			out, err := json.Marshal(map[string]any{
				"comments":   comments,
				"pageInfo":   conn.PageInfo,
				"totalCount": int(conn.TotalCount),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal comments: %w", err)
			}

			return mcp.NewToolResultText(string(out)), nil
		}
}

func ListDiscussionCategories(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_discussion_categories",
			mcp.WithDescription(t("TOOL_LIST_DISCUSSION_CATEGORIES_DESCRIPTION", "List discussion categories with their id and name, for a repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_DISCUSSION_CATEGORIES_USER_TITLE", "List discussion categories"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner string
				Repo  string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}
			var q struct {
				Repository struct {
					DiscussionCategories struct {
						Nodes []struct {
							ID   githubv4.ID
							Name githubv4.String
						}
						PageInfo   paginationPageInfo
						TotalCount githubv4.Int
					} `graphql:"discussionCategories(first: $first)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			vars := map[string]interface{}{
				"owner": githubv4.String(params.Owner),
				"repo":  githubv4.String(params.Repo),
				// A repository has at most 25 discussion categories, so one page holds them all.
				"first": githubv4.Int(25),
			}
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			conn := q.Repository.DiscussionCategories
			categories := make([]map[string]string, 0, len(conn.Nodes))
			for _, c := range conn.Nodes {
				categories = append(categories, map[string]string{
					"id":   fmt.Sprint(c.ID),
					"name": string(c.Name),
				})
			}
			out, err := json.Marshal(map[string]any{
				"categories": categories,
				"pageInfo":   conn.PageInfo,
				"totalCount": int(conn.TotalCount),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussion categories: %w", err)
			}
			return mcp.NewToolResultText(string(out)), nil
		}
}

```

#### x/discussions_test.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/github/github-mcp-server/internal/githubv4mock"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	discussionsPageInfo = map[string]any{
		"hasNextPage":     true,
		"hasPreviousPage": false,
		"startCursor":     "start-cursor",
		"endCursor":       "end-cursor",
	}
	discussionsGeneral = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-02T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-03-02T00:00:00Z", "author": map[string]any{"login": "user3"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	discussionsAll = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-02T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 2, "title": "Discussion 2 title", "createdAt": "2023-02-01T00:00:00Z", "updatedAt": "2023-02-02T00:00:00Z", "author": map[string]any{"login": "user2"}, "url": "https://github.com/owner/repo/discussions/2", "category": map[string]any{"name": "Questions"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-03-02T00:00:00Z", "author": map[string]any{"login": "user3"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	mockResponseListAll = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes":      discussionsAll,
				"pageInfo":   discussionsPageInfo,
				"totalCount": 3,
			},
		},
	})
	mockResponseListGeneral = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes":      discussionsGeneral,
				"pageInfo":   discussionsPageInfo,
				"totalCount": 2,
			},
		},
	})
	mockErrorRepoNotFound = githubv4mock.ErrorResponse("repository not found")
)

func Test_ListDiscussions(t *testing.T) {
	mockClient := githubv4.NewClient(nil)
	// Verify tool definition and schema
	toolDef, _ := ListDiscussions(stubGetGQLClientFn(mockClient), translations.NullTranslationHelper)
	assert.Equal(t, "list_discussions", toolDef.Name)
	assert.NotEmpty(t, toolDef.Description)
	assert.Contains(t, toolDef.InputSchema.Properties, "owner")
	assert.Contains(t, toolDef.InputSchema.Properties, "repo")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo"})

	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")

	// The documents the tool must send, with and without a category filter and an ordering
	qDiscussions := "query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsFiltered := "query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsOrdered := "query($after:String$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsFilteredOrdered := "query($after:String$categoryId:ID!$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"

	// Without pagination arguments, a page of 30 is requested and no cursor is sent
	varsListAll := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(30),
		"after": nil,
	}

	varsListWithPagination := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(5),
		"after": githubv4.String("previous-end-cursor"),
	}

	varsRepoNotFound := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("nonexistent-repo"),
		"first": githubv4.Int(30),
		"after": nil,
	}

	varsDiscussionsFiltered := map[string]interface{}{
		"owner":      githubv4.String("owner"),
		"repo":       githubv4.String("repo"),
		"categoryId": githubv4.ID("DIC_kwDOABC123"),
		"first":      githubv4.Int(30),
		"after":      nil,
	}

	varsOrderByCreatedAsc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"orderByField":     githubv4.DiscussionOrderField("CREATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
		"first":            githubv4.Int(30),
		"after":            nil,
	}

	varsOrderByUpdatedDesc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
		"first":            githubv4.Int(30),
		"after":            nil,
	}

	varsFilteredOrderedWithPagination := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"categoryId":       githubv4.ID("DIC_kwDOABC123"),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
		"first":            githubv4.Int(5),
		"after":            githubv4.String("previous-end-cursor"),
	}

	tests := []struct {
		name          string
		reqParams     map[string]interface{}
		query         string
		vars          map[string]interface{}
		response      githubv4mock.GQLResponse
		expectError   bool
		errContains   string
		expectedCount int
	}{
		{
			name: "list all discussions without category filter",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
			},
			query:         qDiscussions,
			vars:          varsListAll,
			response:      mockResponseListAll,
			expectedCount: 3, // All discussions
		},
		{
			name: "list discussions with pagination",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(5),
				"after":   "previous-end-cursor",
			},
			query:         qDiscussions,
			vars:          varsListWithPagination,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "filter by category ID",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
			},
			query:         qDiscussionsFiltered,
			vars:          varsDiscussionsFiltered,
			response:      mockResponseListGeneral,
			expectedCount: 2, // Only General discussions (matching the category ID)
		},
		{
			name: "order by creation time ascending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "CREATED_AT",
				"direction": "ASC",
			},
			query:         qDiscussionsOrdered,
			vars:          varsOrderByCreatedAsc,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "order by update time descending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "UPDATED_AT",
				"direction": "DESC",
			},
			query:         qDiscussionsOrdered,
			vars:          varsOrderByUpdatedDesc,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "filter by category ID and order, with pagination",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"category":  "DIC_kwDOABC123",
				"orderBy":   "UPDATED_AT",
				"direction": "ASC",
				"perPage":   float64(5),
				"after":     "previous-end-cursor",
			},
			query:         qDiscussionsFilteredOrdered,
			vars:          varsFilteredOrderedWithPagination,
			response:      mockResponseListGeneral,
			expectedCount: 2,
		},
		{
			name: "orderBy without direction is ignored",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"orderBy": "UPDATED_AT",
			},
			query:         qDiscussions,
			vars:          varsListAll,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "direction without orderBy is ignored",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"direction": "DESC",
			},
			query:         qDiscussions,
			vars:          varsListAll,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "orderBy without direction keeps the category filter",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
				"orderBy":  "CREATED_AT",
			},
			query:         qDiscussionsFiltered,
			vars:          varsDiscussionsFiltered,
			response:      mockResponseListGeneral,
			expectedCount: 2,
		},
		{
			name: "repository not found error",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "nonexistent-repo",
			},
			query:       qDiscussions,
			vars:        varsRepoNotFound,
			response:    mockErrorRepoNotFound,
			expectError: true,
			errContains: "repository not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matcher := githubv4mock.NewQueryMatcher(tc.query, tc.vars, tc.response)
			gqlClient := githubv4.NewClient(githubv4mock.NewMockedHTTPClient(matcher))
			_, handler := ListDiscussions(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			req := createMCPRequest(tc.reqParams)
			res, err := handler(context.Background(), req)
			text := getTextResult(t, res).Text

			if tc.expectError {
				require.True(t, res.IsError)
				assert.Contains(t, text, tc.errContains)
				return
			}
			require.NoError(t, err)
			require.False(t, res.IsError, text)

			var response struct {
				Discussions []*github.Discussion `json:"discussions"`
				PageInfo    struct {
					HasNextPage     bool   `json:"hasNextPage"`
					HasPreviousPage bool   `json:"hasPreviousPage"`
					StartCursor     string `json:"startCursor"`
					EndCursor       string `json:"endCursor"`
				} `json:"pageInfo"`
				TotalCount int `json:"totalCount"`
			}
			err = json.Unmarshal([]byte(text), &response)
			require.NoError(t, err)
			returnedDiscussions := response.Discussions

			assert.True(t, response.PageInfo.HasNextPage)
			assert.False(t, response.PageInfo.HasPreviousPage)
			assert.Equal(t, "start-cursor", response.PageInfo.StartCursor)
			assert.Equal(t, "end-cursor", response.PageInfo.EndCursor)
			assert.Equal(t, tc.expectedCount, response.TotalCount)

			assert.Len(t, returnedDiscussions, tc.expectedCount, "Expected %d discussions, got %d", tc.expectedCount, len(returnedDiscussions))

			// Every discussion carries its update time and its author
			for _, discussion := range returnedDiscussions {
				require.NotNil(t, discussion.UpdatedAt, "Discussion should have updatedAt")
				assert.False(t, discussion.UpdatedAt.IsZero(), "Discussion should have updatedAt")
				require.NotNil(t, discussion.User, "Discussion should have an author")
				assert.Equal(t, fmt.Sprintf("user%d", discussion.GetNumber()), discussion.User.GetLogin())
			}

			// Verify that all returned discussions have a category if filtered
			if _, hasCategory := tc.reqParams["category"]; hasCategory {
				for _, discussion := range returnedDiscussions {
					require.NotNil(t, discussion.DiscussionCategory, "Discussion should have category")
					assert.NotEmpty(t, *discussion.DiscussionCategory.Name, "Discussion should have category name")
				}
			}
		})
	}
}

func Test_GetDiscussion(t *testing.T) {
	// Verify tool definition and schema
	toolDef, _ := GetDiscussion(nil, translations.NullTranslationHelper)
	assert.Equal(t, "get_discussion", toolDef.Name)
	assert.NotEmpty(t, toolDef.Description)
	assert.Contains(t, toolDef.InputSchema.Properties, "owner")
	assert.Contains(t, toolDef.InputSchema.Properties, "repo")
	assert.Contains(t, toolDef.InputSchema.Properties, "discussionNumber")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo", "discussionNumber"})

	var q struct {
		Repository struct {
			Discussion struct {
				Number    githubv4.Int
				Body      githubv4.String
				CreatedAt githubv4.DateTime
				URL       githubv4.String `graphql:"url"`
				Category  struct {
					Name githubv4.String
				} `graphql:"category"`
			} `graphql:"discussion(number: $discussionNumber)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	vars := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"discussionNumber": githubv4.Int(1),
	}
	tests := []struct {
		name        string
		response    githubv4mock.GQLResponse
		expectError bool
		expected    *github.Discussion
		errContains string
	}{
		{
			name: "successful retrieval",
			response: githubv4mock.DataResponse(map[string]any{
				"repository": map[string]any{"discussion": map[string]any{
					"number":    1,
					"body":      "This is a test discussion",
					"url":       "https://github.com/owner/repo/discussions/1",
					"createdAt": "2025-04-25T12:00:00Z",
					"category":  map[string]any{"name": "General"},
				}},
			}),
			expectError: false,
			expected: &github.Discussion{
				HTMLURL:   github.Ptr("https://github.com/owner/repo/discussions/1"),
				Number:    github.Ptr(1),
				Body:      github.Ptr("This is a test discussion"),
				CreatedAt: &github.Timestamp{Time: time.Date(2025, 4, 25, 12, 0, 0, 0, time.UTC)},
				DiscussionCategory: &github.DiscussionCategory{
					Name: github.Ptr("General"),
				},
			},
		},
		{
			name:        "discussion not found",
			response:    githubv4mock.ErrorResponse("discussion not found"),
			expectError: true,
			errContains: "discussion not found",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matcher := githubv4mock.NewQueryMatcher(q, vars, tc.response)
			httpClient := githubv4mock.NewMockedHTTPClient(matcher)
			gqlClient := githubv4.NewClient(httpClient)
			_, handler := GetDiscussion(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			req := createMCPRequest(map[string]interface{}{"owner": "owner", "repo": "repo", "discussionNumber": int32(1)})
			res, err := handler(context.Background(), req)
			text := getTextResult(t, res).Text

			if tc.expectError {
				require.True(t, res.IsError)
				assert.Contains(t, text, tc.errContains)
				return
			}

			require.NoError(t, err)
			var out github.Discussion
			require.NoError(t, json.Unmarshal([]byte(text), &out))
			assert.Equal(t, *tc.expected.HTMLURL, *out.HTMLURL)
			assert.Equal(t, *tc.expected.Number, *out.Number)
			assert.Equal(t, *tc.expected.Body, *out.Body)
			// Check category label
			assert.Equal(t, *tc.expected.DiscussionCategory.Name, *out.DiscussionCategory.Name)
		})
	}
}

func Test_GetDiscussionComments(t *testing.T) {
	// Verify tool definition and schema
	toolDef, _ := GetDiscussionComments(nil, translations.NullTranslationHelper)
	assert.Equal(t, "get_discussion_comments", toolDef.Name)
	assert.NotEmpty(t, toolDef.Description)
	assert.Contains(t, toolDef.InputSchema.Properties, "owner")
	assert.Contains(t, toolDef.InputSchema.Properties, "repo")
	assert.Contains(t, toolDef.InputSchema.Properties, "discussionNumber")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo", "discussionNumber"})

	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")

	q := "query($after:String$discussionNumber:Int!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussion(number: $discussionNumber){comments(first: $first, after: $after){nodes{body},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"
	// Without pagination arguments, a page of 30 is requested and no cursor is sent
	vars := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"discussionNumber": githubv4.Int(1),
		"first":            githubv4.Int(30),
		"after":            nil,
	}
	mockResponse := githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussion": map[string]any{
				"comments": map[string]any{
					"nodes": []map[string]any{
						{"body": "This is the first comment"},
						{"body": "This is the second comment"},
					},
					"pageInfo": map[string]any{
						"hasNextPage":     false,
						"hasPreviousPage": false,
						"startCursor":     "",
						"endCursor":       "",
					},
					"totalCount": 2,
				},
			},
		},
	})
	matcher := githubv4mock.NewQueryMatcher(q, vars, mockResponse)
	httpClient := githubv4mock.NewMockedHTTPClient(matcher)
	gqlClient := githubv4.NewClient(httpClient)
	_, handler := GetDiscussionComments(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

	request := createMCPRequest(map[string]interface{}{
		"owner":            "owner",
		"repo":             "repo",
		"discussionNumber": int32(1),
	})

	result, err := handler(context.Background(), request)
	require.NoError(t, err)

	textContent := getTextResult(t, result)

	var response struct {
		Comments []*github.IssueComment `json:"comments"`
		PageInfo struct {
			HasNextPage bool   `json:"hasNextPage"`
			EndCursor   string `json:"endCursor"`
		} `json:"pageInfo"`
		TotalCount int `json:"totalCount"`
	}
	err = json.Unmarshal([]byte(textContent.Text), &response)
	require.NoError(t, err)
	returnedComments := response.Comments
	assert.Len(t, returnedComments, 2)
	assert.False(t, response.PageInfo.HasNextPage)
	assert.Equal(t, 2, response.TotalCount)
	expectedBodies := []string{"This is the first comment", "This is the second comment"}
	for i, comment := range returnedComments {
		assert.Equal(t, expectedBodies[i], *comment.Body)
	}
}

func Test_ListDiscussionCategories(t *testing.T) {
	q := "query($first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussionCategories(first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	// The first 25 categories are always requested
	vars := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(25),
	}
	mockResp := githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussionCategories": map[string]any{
				"nodes": []map[string]any{
					{"id": "123", "name": "CategoryOne"},
					{"id": "456", "name": "CategoryTwo"},
				},
				"pageInfo": map[string]any{
					"hasNextPage":     false,
					"hasPreviousPage": false,
					"startCursor":     "",
					"endCursor":       "",
				},
				"totalCount": 2,
			},
		},
	})
	matcher := githubv4mock.NewQueryMatcher(q, vars, mockResp)
	httpClient := githubv4mock.NewMockedHTTPClient(matcher)
	gqlClient := githubv4.NewClient(httpClient)

	tool, handler := ListDiscussionCategories(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)
	assert.Equal(t, "list_discussion_categories", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "owner")
	assert.Contains(t, tool.InputSchema.Properties, "repo")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo"})
	for _, removed := range []string{"first", "last", "after", "before"} {
		assert.NotContains(t, tool.InputSchema.Properties, removed)
	}

	request := createMCPRequest(map[string]interface{}{"owner": "owner", "repo": "repo"})
	result, err := handler(context.Background(), request)
	require.NoError(t, err)

	text := getTextResult(t, result).Text
	var response struct {
		Categories []map[string]string `json:"categories"`
		PageInfo   struct {
			HasNextPage bool `json:"hasNextPage"`
		} `json:"pageInfo"`
		TotalCount int `json:"totalCount"`
	}
	require.NoError(t, json.Unmarshal([]byte(text), &response))
	categories := response.Categories
	assert.Len(t, categories, 2)
	assert.False(t, response.PageInfo.HasNextPage)
	assert.Equal(t, 2, response.TotalCount)
	assert.Equal(t, "123", categories[0]["id"])
	assert.Equal(t, "CategoryOne", categories[0]["name"])
	assert.Equal(t, "456", categories[1]["id"])
	assert.Equal(t, "CategoryTwo", categories[1]["name"])
}

```

#### x/helper_test.go

```go
package github

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// expectations are what a mocked endpoint checks in the request it receives.
//
type expectations struct {
	path        string
	queryParams map[string]string
	requestBody any
}

// expect is a helper function to create a partial mock that expects various
// request behaviors, such as path, query parameters, and request body.
func expect(t *testing.T, e expectations) *partialMock {
	return &partialMock{
		t:                   t,
		expectedPath:        e.path,
		expectedQueryParams: e.queryParams,
		expectedRequestBody: e.requestBody,
	}
}

// expectPath is a helper function to create a partial mock that expects a
// request with the given path, with the ability to chain a response handler.
func expectPath(t *testing.T, expectedPath string) *partialMock {
	return &partialMock{
		t:            t,
		expectedPath: expectedPath,
	}
}

// expectQueryParams is a helper function to create a partial mock that expects a
// request with the given query parameters, with the ability to chain a response handler.
func expectQueryParams(t *testing.T, expectedQueryParams map[string]string) *partialMock {
	return &partialMock{
		t:                   t,
		expectedQueryParams: expectedQueryParams,
	}
}

// expectRequestBody is a helper function to create a partial mock that expects a
// request with the given body, with the ability to chain a response handler.
func expectRequestBody(t *testing.T, expectedRequestBody any) *partialMock {
	return &partialMock{
		t:                   t,
		expectedRequestBody: expectedRequestBody,
	}
}

type partialMock struct {
	t *testing.T

	expectedPath        string
	expectedQueryParams map[string]string
	expectedRequestBody any
}

func (p *partialMock) andThen(responseHandler http.HandlerFunc) http.HandlerFunc {
	p.t.Helper()
	return func(w http.ResponseWriter, r *http.Request) {
		if p.expectedPath != "" {
			require.Equal(p.t, p.expectedPath, r.URL.Path)
		}

		if p.expectedQueryParams != nil {
			require.Equal(p.t, len(p.expectedQueryParams), len(r.URL.Query()))
			for k, v := range p.expectedQueryParams {
				require.Equal(p.t, v, r.URL.Query().Get(k))
			}
		}

		if p.expectedRequestBody != nil {
			var unmarshaledRequestBody any
			err := json.NewDecoder(r.Body).Decode(&unmarshaledRequestBody)
			require.NoError(p.t, err)

			require.Equal(p.t, p.expectedRequestBody, unmarshaledRequestBody)
		}

		responseHandler(w, r)
	}
}

// mockResponse is a helper function to create a mock HTTP response handler
// that returns a specified status code and marshaled body.
func mockResponse(t *testing.T, code int, body interface{}) http.HandlerFunc {
	t.Helper()
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(code)
		// Some tests do not expect to return a JSON object, such as fetching a raw pull request diff,
		// so allow strings to be returned directly.
		s, ok := body.(string)
		if ok {
			_, _ = w.Write([]byte(s))
			return
		}

		b, err := json.Marshal(body)
		require.NoError(t, err)
		_, _ = w.Write(b)
	}
}

// createMCPRequest is a helper function to create a MCP request with the given arguments.
func createMCPRequest(args any) mcp.CallToolRequest {
	return mcp.CallToolRequest{
		Params: struct {
			Name      string    `json:"name"`
			Arguments any       `json:"arguments,omitempty"`
			Meta      *mcp.Meta `json:"_meta,omitempty"`
		}{
			Arguments: args,
		},
	}
}

// getTextResult is a helper function that returns a text result from a tool call.
func getTextResult(t *testing.T, result *mcp.CallToolResult) mcp.TextContent {
	t.Helper()
	assert.NotNil(t, result)
	require.Len(t, result.Content, 1)
	require.IsType(t, mcp.TextContent{}, result.Content[0])
	textContent := result.Content[0].(mcp.TextContent)
	assert.Equal(t, "text", textContent.Type)
	return textContent
}

func getErrorResult(t *testing.T, result *mcp.CallToolResult) mcp.TextContent {
	res := getTextResult(t, result)
	require.True(t, result.IsError, "expected tool call result to be an error")
	return res
}

// getTextResourceResult is a helper function that returns a text result from a tool call.
func getTextResourceResult(t *testing.T, result *mcp.CallToolResult) mcp.TextResourceContents {
	t.Helper()
	assert.NotNil(t, result)
	require.Len(t, result.Content, 2)
	content := result.Content[1]
	require.IsType(t, mcp.EmbeddedResource{}, content)
	resource := content.(mcp.EmbeddedResource)
	require.IsType(t, mcp.TextResourceContents{}, resource.Resource)
	return resource.Resource.(mcp.TextResourceContents)
}

// getBlobResourceResult is a helper function that returns a blob result from a tool call.
func getBlobResourceResult(t *testing.T, result *mcp.CallToolResult) mcp.BlobResourceContents {
	t.Helper()
	assert.NotNil(t, result)
	require.Len(t, result.Content, 2)
	content := result.Content[1]
	require.IsType(t, mcp.EmbeddedResource{}, content)
	resource := content.(mcp.EmbeddedResource)
	require.IsType(t, mcp.BlobResourceContents{}, resource.Resource)
	return resource.Resource.(mcp.BlobResourceContents)
}

func TestOptionalParamOK(t *testing.T) {
	tests := []struct {
		name        string
		args        map[string]interface{}
		paramName   string
		expectedVal interface{}
		expectedOk  bool
		expectError bool
		errorMsg    string
	}{
		{
			name:        "present and correct type (string)",
			args:        map[string]interface{}{"myParam": "hello"},
			paramName:   "myParam",
			expectedVal: "hello",
			expectedOk:  true,
			expectError: false,
		},
		{
			name:        "present and correct type (bool)",
			args:        map[string]interface{}{"myParam": true},
			paramName:   "myParam",
			expectedVal: true,
			expectedOk:  true,
			expectError: false,
		},
		{
			name:        "present and correct type (number)",
			args:        map[string]interface{}{"myParam": float64(123)},
			paramName:   "myParam",
			expectedVal: float64(123),
			expectedOk:  true,
			expectError: false,
		},
		{
			name:        "present but wrong type (string expected, got bool)",
			args:        map[string]interface{}{"myParam": true},
			paramName:   "myParam",
			expectedVal: "",   // Zero value for string
			expectedOk:  true, // ok is true because param exists
			expectError: true,
			errorMsg:    "parameter myParam is not of type string, is bool",
		},
		{
			name:        "present but wrong type (bool expected, got string)",
			args:        map[string]interface{}{"myParam": "true"},
			paramName:   "myParam",
			expectedVal: false, // Zero value for bool
			expectedOk:  true,  // ok is true because param exists
			expectError: true,
			errorMsg:    "parameter myParam is not of type bool, is string",
		},
		{
			name:        "parameter not present",
			args:        map[string]interface{}{"anotherParam": "value"},
			paramName:   "myParam",
			expectedVal: "", // Zero value for string
			expectedOk:  false,
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.args)

			// Test with string type assertion
			if _, isString := tc.expectedVal.(string); isString || tc.errorMsg == "parameter myParam is not of type string, is bool" {
				val, ok, err := OptionalParamOK[string](request, tc.paramName)
				if tc.expectError {
					require.Error(t, err)
					assert.Contains(t, err.Error(), tc.errorMsg)
					assert.Equal(t, tc.expectedOk, ok)   // Check ok even on error
					assert.Equal(t, tc.expectedVal, val) // Check zero value on error
				} else {
					require.NoError(t, err)
					assert.Equal(t, tc.expectedOk, ok)
					assert.Equal(t, tc.expectedVal, val)
				}
			}

			// Test with bool type assertion
			if _, isBool := tc.expectedVal.(bool); isBool || tc.errorMsg == "parameter myParam is not of type bool, is string" {
				val, ok, err := OptionalParamOK[bool](request, tc.paramName)
				if tc.expectError {
					require.Error(t, err)
					assert.Contains(t, err.Error(), tc.errorMsg)
					assert.Equal(t, tc.expectedOk, ok)   // Check ok even on error
					assert.Equal(t, tc.expectedVal, val) // Check zero value on error
				} else {
					require.NoError(t, err)
					assert.Equal(t, tc.expectedOk, ok)
					assert.Equal(t, tc.expectedVal, val)
				}
			}

			// Test with float64 type assertion (for number case)
			if _, isFloat := tc.expectedVal.(float64); isFloat {
				val, ok, err := OptionalParamOK[float64](request, tc.paramName)
				if tc.expectError {
					// This case shouldn't happen for float64 in the defined tests
					require.Fail(t, "Unexpected error case for float64")
				} else {
					require.NoError(t, err)
					assert.Equal(t, tc.expectedOk, ok)
					assert.Equal(t, tc.expectedVal, val)
				}
			}
		})
	}
}

```

#### x/issues.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

// GetIssue creates a tool to get details of a specific issue in a GitHub repository.
func GetIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_issue",
			mcp.WithDescription(t("TOOL_GET_ISSUE_DESCRIPTION", "Get details of a specific issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_ISSUE_USER_TITLE", "Get issue details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("The owner of the repository"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("The name of the repository"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("The number of the issue"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			issue, resp, err := client.Issues.Get(ctx, owner, repo, issueNumber)
			if err != nil {
				return nil, fmt.Errorf("failed to get issue: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get issue: %s", string(body))), nil
			}

			r, err := json.Marshal(issue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal issue: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// AddIssueComment creates a tool to add a comment to an issue.
func AddIssueComment(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("add_issue_comment",
			mcp.WithDescription(t("TOOL_ADD_ISSUE_COMMENT_DESCRIPTION", "Add a comment to a specific issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_ADD_ISSUE_COMMENT_USER_TITLE", "Add comment to issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("Issue number to comment on"),
			),
			mcp.WithString("body",
				mcp.Required(),
				mcp.Description("Comment content"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			body, err := RequiredParam[string](request, "body")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			comment := &github.IssueComment{
				Body: github.Ptr(body),
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			createdComment, resp, err := client.Issues.CreateComment(ctx, owner, repo, issueNumber, comment)
			if err != nil {
				return nil, fmt.Errorf("failed to create comment: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create comment: %s", string(body))), nil
			}

			r, err := json.Marshal(createdComment)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// SearchIssues creates a tool to search for issues.
func SearchIssues(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_issues",
			mcp.WithDescription(t("TOOL_SEARCH_ISSUES_DESCRIPTION", "Search for issues in GitHub repositories using issues search syntax already scoped to is:issue")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SEARCH_ISSUES_USER_TITLE", "Search issues"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("Search query using GitHub issues search syntax"),
			),
			mcp.WithString("owner",
				mcp.Description("Optional repository owner. If provided with repo, only notifications for this repository are listed."),
			),
			mcp.WithString("repo",
				mcp.Description("Optional repository name. If provided with owner, only notifications for this repository are listed."),
			),
			mcp.WithString("sort",
				mcp.Description("Sort field by number of matches of categories, defaults to best match"),
				mcp.Enum(
					"comments",
					"reactions",
					"reactions-+1",
					"reactions--1",
					"reactions-smile",
					"reactions-thinking_face",
					"reactions-heart",
					"reactions-tada",
					"interactions",
					"created",
					"updated",
				),
			),
			mcp.WithString("order",
				mcp.Description("Sort order"),
				mcp.Enum("asc", "desc"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return searchHandler(ctx, getClient, request, "issue", "failed to search issues")
		}
}

// CreateIssue creates a tool to create a new issue in a GitHub repository.
func CreateIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("create_issue",
			mcp.WithDescription(t("TOOL_CREATE_ISSUE_DESCRIPTION", "Create a new issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_ISSUE_USER_TITLE", "Open new issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("title",
				mcp.Required(),
				mcp.Description("Issue title"),
			),
			mcp.WithString("body",
				mcp.Description("Issue body content"),
			),
			mcp.WithArray("assignees",
				mcp.Description("Usernames to assign to this issue"),
				mcp.Items(
					map[string]any{
						"type": "string",
					},
				),
			),
			mcp.WithArray("labels",
				mcp.Description("Labels to apply to this issue"),
				mcp.Items(
					map[string]any{
						"type": "string",
					},
				),
			),
			mcp.WithNumber("milestone",
				mcp.Description("Milestone number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			title, err := RequiredParam[string](request, "title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Optional parameters
			body, err := OptionalParam[string](request, "body")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get assignees
			assignees, err := OptionalStringArrayParam(request, "assignees")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get labels
			labels, err := OptionalStringArrayParam(request, "labels")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional milestone
			milestone, err := OptionalIntParam(request, "milestone")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var milestoneNum *int
			if milestone != 0 {
				milestoneNum = &milestone
			}

			// Create the issue request
			issueRequest := &github.IssueRequest{
				Title:     github.Ptr(title),
				Body:      github.Ptr(body),
				Assignees: &assignees,
				Labels:    &labels,
				Milestone: milestoneNum,
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			issue, resp, err := client.Issues.Create(ctx, owner, repo, issueRequest)
			if err != nil {
				return nil, fmt.Errorf("failed to create issue: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create issue: %s", string(body))), nil
			}

			r, err := json.Marshal(issue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListIssues creates a tool to list and filter repository issues
func ListIssues(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_issues",
			mcp.WithDescription(t("TOOL_LIST_ISSUES_DESCRIPTION", "List issues in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_ISSUES_USER_TITLE", "List issues"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("state",
				mcp.Description("Filter by state"),
				mcp.Enum("open", "closed", "all"),
			),
			mcp.WithArray("labels",
				mcp.Description("Filter by labels"),
				mcp.Items(
					map[string]interface{}{
						"type": "string",
					},
				),
			),
			mcp.WithString("sort",
				mcp.Description("Sort order"),
				mcp.Enum("created", "updated", "comments"),
			),
			mcp.WithString("direction",
				mcp.Description("Sort direction"),
				mcp.Enum("asc", "desc"),
			),
			mcp.WithString("since",
				mcp.Description("Filter by date (ISO 8601 timestamp)"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.IssueListByRepoOptions{}

			// Set optional parameters if provided
			opts.State, err = OptionalParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get labels
			opts.Labels, err = OptionalStringArrayParam(request, "labels")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts.Sort, err = OptionalParam[string](request, "sort")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts.Direction, err = OptionalParam[string](request, "direction")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			since, err := OptionalParam[string](request, "since")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if since != "" {
				timestamp, err := parseISOTimestamp(since)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("failed to list issues: %s", err.Error())), nil
				}
				opts.Since = timestamp
			}

			if page, ok := request.GetArguments()["page"].(float64); ok {
				opts.ListOptions.Page = int(page)
			}

			if perPage, ok := request.GetArguments()["perPage"].(float64); ok {
				opts.ListOptions.PerPage = int(perPage)
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			issues, resp, err := client.Issues.ListByRepo(ctx, owner, repo, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to list issues: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list issues: %s", string(body))), nil
			}

			r, err := json.Marshal(issues)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal issues: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// UpdateIssue creates a tool to update an existing issue in a GitHub repository.
func UpdateIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("update_issue",
			mcp.WithDescription(t("TOOL_UPDATE_ISSUE_DESCRIPTION", "Update an existing issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_UPDATE_ISSUE_USER_TITLE", "Edit issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("Issue number to update"),
			),
			mcp.WithString("title",
				mcp.Description("New title"),
			),
			mcp.WithString("body",
				mcp.Description("New description"),
			),
			mcp.WithString("state",
				mcp.Description("New state"),
				mcp.Enum("open", "closed"),
			),
			mcp.WithArray("labels",
				mcp.Description("New labels"),
				mcp.Items(
					map[string]interface{}{
						"type": "string",
					},
				),
			),
			mcp.WithArray("assignees",
				mcp.Description("New assignees"),
				mcp.Items(
					map[string]interface{}{
						"type": "string",
					},
				),
			),
			mcp.WithNumber("milestone",
				mcp.Description("New milestone number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Create the issue request with only provided fields
			issueRequest := &github.IssueRequest{}

			// Set optional parameters if provided
			title, err := OptionalParam[string](request, "title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if title != "" {
				issueRequest.Title = github.Ptr(title)
			}

			body, err := OptionalParam[string](request, "body")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if body != "" {
				issueRequest.Body = github.Ptr(body)
			}

			state, err := OptionalParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if state != "" {
				issueRequest.State = github.Ptr(state)
			}

			// Get labels
			labels, err := OptionalStringArrayParam(request, "labels")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if len(labels) > 0 {
				issueRequest.Labels = &labels
			}

			// Get assignees
			assignees, err := OptionalStringArrayParam(request, "assignees")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if len(assignees) > 0 {
				issueRequest.Assignees = &assignees
			}

			milestone, err := OptionalIntParam(request, "milestone")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if milestone != 0 {
				milestoneNum := milestone
				issueRequest.Milestone = &milestoneNum
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			updatedIssue, resp, err := client.Issues.Edit(ctx, owner, repo, issueNumber, issueRequest)
			if err != nil {
				return nil, fmt.Errorf("failed to update issue: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to update issue: %s", string(body))), nil
			}

			r, err := json.Marshal(updatedIssue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetIssueComments creates a tool to get comments for a GitHub issue.
func GetIssueComments(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_issue_comments",
			mcp.WithDescription(t("TOOL_GET_ISSUE_COMMENTS_DESCRIPTION", "Get comments for a specific issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_ISSUE_COMMENTS_USER_TITLE", "Get issue comments"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("Issue number"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.IssueListCommentsOptions{
				ListOptions: github.ListOptions{
					Page:    pagination.Page,
					PerPage: pagination.PerPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			comments, resp, err := client.Issues.ListComments(ctx, owner, repo, issueNumber, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to get issue comments: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get issue comments: %s", string(body))), nil
			}

			r, err := json.Marshal(comments)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// mvpDescription is an MVP idea for generating tool descriptions from structured data in a shared format.
// It is not intended for widespread usage and is not a complete implementation.
type mvpDescription struct {
	summary        string
	outcomes       []string
	referenceLinks []string
}

func (d *mvpDescription) String() string {
	var sb strings.Builder
	sb.WriteString(d.summary)
	if len(d.outcomes) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString("This tool can help with the following outcomes:\n")
		for _, outcome := range d.outcomes {
			sb.WriteString(fmt.Sprintf("- %s\n", outcome))
		}
	}

	if len(d.referenceLinks) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString("More information can be found at:\n")
		for _, link := range d.referenceLinks {
			sb.WriteString(fmt.Sprintf("- %s\n", link))
		}
	}

	return sb.String()
}

func AssignCopilotToIssue(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	description := mvpDescription{
		summary: "Assign Copilot to a specific issue in a GitHub repository.",
		outcomes: []string{
			"a Pull Request created with source code changes to resolve the issue",
		},
		referenceLinks: []string{
			"https://docs.github.com/en/copilot/using-github-copilot/using-copilot-coding-agent-to-work-on-tasks/about-assigning-tasks-to-copilot",
		},
	}

	return mcp.NewTool("assign_copilot_to_issue",
			mcp.WithDescription(t("TOOL_ASSIGN_COPILOT_TO_ISSUE_DESCRIPTION", description.String())),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:          t("TOOL_ASSIGN_COPILOT_TO_ISSUE_USER_TITLE", "Assign Copilot to issue"),
				ReadOnlyHint:   ToBoolPtr(false),
				IdempotentHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issueNumber",
				mcp.Required(),
				mcp.Description("Issue number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner       string
				Repo        string
				IssueNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Firstly, we try to find the copilot bot in the suggested actors for the repository.
			// Although as I write this, we would expect copilot to be at the top of the list, in future, maybe
			// it will not be on the first page of responses, thus we will keep paginating until we find it.
			type botAssignee struct {
				ID       githubv4.ID
				Login    string
				TypeName string `graphql:"__typename"`
			}

			type suggestedActorsQuery struct {
				Repository struct {
					SuggestedActors struct {
						Nodes []struct {
							Bot botAssignee `graphql:"... on Bot"`
						}
						PageInfo struct {
							HasNextPage bool
							EndCursor   string
						}
					} `graphql:"suggestedActors(first: 100, after: $endCursor, capabilities: CAN_BE_ASSIGNED)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			variables := map[string]any{
				"owner":     githubv4.String(params.Owner),
				"name":      githubv4.String(params.Repo),
				"endCursor": (*githubv4.String)(nil),
			}

			var copilotAssignee *botAssignee
			for {
				var query suggestedActorsQuery
				err := client.Query(ctx, &query, variables)
				if err != nil {
					return nil, err
				}

				// Iterate all the returned nodes looking for the copilot bot, which is supposed to have the
				// same name on each host. We need this in order to get the ID for later assignment.
				for _, node := range query.Repository.SuggestedActors.Nodes {
					if node.Bot.Login == "copilot-swe-agent" {
						copilotAssignee = &node.Bot
						break
					}
				}

				if !query.Repository.SuggestedActors.PageInfo.HasNextPage {
					break
				}
				variables["endCursor"] = githubv4.String(query.Repository.SuggestedActors.PageInfo.EndCursor)
			}

			// If we didn't find the copilot bot, we can't proceed any further.
			if copilotAssignee == nil {
				// The e2e tests depend upon this specific message to skip the test.
				return mcp.NewToolResultError("copilot isn't available as an assignee for this issue. Please inform the user to visit https://docs.github.com/en/copilot/using-github-copilot/using-copilot-coding-agent-to-work-on-tasks/about-assigning-tasks-to-copilot for more information."), nil
			}

			// Next let's get the GQL Node ID and current assignees for this issue because the only way to
			// assign copilot is to use replaceActorsForAssignable which requires the full list.
			var getIssueQuery struct {
				Repository struct {
					Issue struct {
						ID        githubv4.ID
						Assignees struct {
							Nodes []struct {
								ID githubv4.ID
							}
						} `graphql:"assignees(first: 100)"`
					} `graphql:"issue(number: $number)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			variables = map[string]any{
				"owner":  githubv4.String(params.Owner),
				"name":   githubv4.String(params.Repo),
				"number": githubv4.Int(params.IssueNumber),
			}

			if err := client.Query(ctx, &getIssueQuery, variables); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get issue ID: %v", err)), nil
			}

			// Finally, do the assignment. Just for reference, assigning copilot to an issue that it is already
			// assigned to seems to have no impact (which is a good thing).
			var assignCopilotMutation struct {
				ReplaceActorsForAssignable struct {
					Typename string `graphql:"__typename"` // Not required but we need a selector or GQL errors
				} `graphql:"replaceActorsForAssignable(input: $input)"`
			}

			actorIDs := make([]githubv4.ID, len(getIssueQuery.Repository.Issue.Assignees.Nodes)+1)
			for i, node := range getIssueQuery.Repository.Issue.Assignees.Nodes {
				actorIDs[i] = node.ID
			}
			actorIDs[len(getIssueQuery.Repository.Issue.Assignees.Nodes)] = copilotAssignee.ID

			if err := client.Mutate(
				ctx,
				&assignCopilotMutation,
				ReplaceActorsForAssignableInput{
					AssignableID: getIssueQuery.Repository.Issue.ID,
					ActorIDs:     actorIDs,
				},
				nil,
			); err != nil {
				return nil, fmt.Errorf("failed to replace actors for assignable: %w", err)
			}

			return mcp.NewToolResultText("successfully assigned copilot to issue"), nil
		}
}

type ReplaceActorsForAssignableInput struct {
	AssignableID githubv4.ID   `json:"assignableId"`
	ActorIDs     []githubv4.ID `json:"actorIds"`
}

// parseISOTimestamp parses an ISO 8601 timestamp string into a time.Time object.
// Returns the parsed time or an error if parsing fails.
// Example formats supported: "2023-01-15T14:30:00Z", "2023-01-15"
func parseISOTimestamp(timestamp string) (time.Time, error) {
	if timestamp == "" {
		return time.Time{}, fmt.Errorf("empty timestamp")
	}

	// Try RFC3339 format (standard ISO 8601 with time)
	t, err := time.Parse(time.RFC3339, timestamp)
	if err == nil {
		return t, nil
	}

	// Try simple date format (YYYY-MM-DD)
	t, err = time.Parse("2006-01-02", timestamp)
	if err == nil {
		return t, nil
	}

	// Return error with supported formats
	return time.Time{}, fmt.Errorf("invalid ISO 8601 timestamp: %s (supported formats: YYYY-MM-DDThh:mm:ssZ or YYYY-MM-DD)", timestamp)
}

func AssignCodingAgentPrompt(t translations.TranslationHelperFunc) (tool mcp.Prompt, handler server.PromptHandlerFunc) {
	return mcp.NewPrompt("AssignCodingAgent",
			mcp.WithPromptDescription(t("PROMPT_ASSIGN_CODING_AGENT_DESCRIPTION", "Assign GitHub Coding Agent to multiple tasks in a GitHub repository.")),
			mcp.WithArgument("repo", mcp.ArgumentDescription("The repository to assign tasks in (owner/repo)."), mcp.RequiredArgument()),
		), func(_ context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			repo := request.Params.Arguments["repo"]

			messages := []mcp.PromptMessage{
				{
					Role:    "system",
					Content: mcp.NewTextContent("You are a personal assistant for GitHub the Copilot GitHub Coding Agent. Your task is to help the user assign tasks to the Coding Agent based on their open GitHub issues. You can use `assign_copilot_to_issue` tool to assign the Coding Agent to issues that are suitable for autonomous work, and `search_issues` tool to find issues that match the user's criteria. You can also use `list_issues` to get a list of issues in the repository."),
				},
				{
					Role:    "user",
					Content: mcp.NewTextContent(fmt.Sprintf("Please go and get a list of the most recent 10 issues from the %s GitHub repository", repo)),
				},
				{
					Role:    "assistant",
					Content: mcp.NewTextContent(fmt.Sprintf("Sure! I will get a list of the 10 most recent issues for the repo %s.", repo)),
				},
				{
					Role:    "user",
					Content: mcp.NewTextContent("For each issue, please check if it is a clearly defined coding task with acceptance criteria and a low to medium complexity to identify issues that are suitable for an AI Coding Agent to work on. Then assign each of the identified issues to Copilot."),
				},
				{
					Role:    "assistant",
					Content: mcp.NewTextContent("Certainly! Let me carefully check which ones are clearly scoped issues that are good to assign to the coding agent, and I will summarize and assign them now."),
				},
				{
					Role:    "user",
					Content: mcp.NewTextContent("Great, if you are unsure if an issue is good to assign, ask me first, rather than assigning copilot. If you are certain the issue is clear and suitable you can assign it to Copilot without asking."),
				},
			}
			return &mcp.GetPromptResult{
				Messages: messages,
			}, nil
		}
}

```

#### x/notifications.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	FilterDefault           = "default"
	FilterIncludeRead       = "include_read_notifications"
	FilterOnlyParticipating = "only_participating"
)

// ListNotifications creates a tool to list notifications for the current user.
func ListNotifications(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_notifications",
			mcp.WithDescription(t("TOOL_LIST_NOTIFICATIONS_DESCRIPTION", "Lists all GitHub notifications for the authenticated user, including unread notifications, mentions, review requests, assignments, and updates on issues or pull requests. Use this tool whenever the user asks what to work on next, requests a summary of their GitHub activity, wants to see pending reviews, or needs to check for new updates or tasks. This tool is the primary way to discover actionable items, reminders, and outstanding work on GitHub. Always call this tool when asked what to work on next, what is pending, or what needs attention in GitHub.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_NOTIFICATIONS_USER_TITLE", "List notifications"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("filter",
				mcp.Description("Filter notifications to, use default unless specified. Read notifications are ones that have already been acknowledged by the user. Participating notifications are those that the user is directly involved in, such as issues or pull requests they have commented on or created."),
				mcp.Enum(FilterDefault, FilterIncludeRead, FilterOnlyParticipating),
			),
			mcp.WithString("since",
				mcp.Description("Only show notifications updated after the given time (ISO 8601 format)"),
			),
			mcp.WithString("before",
				mcp.Description("Only show notifications updated before the given time (ISO 8601 format)"),
			),
			mcp.WithString("owner",
				mcp.Description("Optional repository owner. If provided with repo, only notifications for this repository are listed."),
			),
			mcp.WithString("repo",
				mcp.Description("Optional repository name. If provided with owner, only notifications for this repository are listed."),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			filter, err := OptionalParam[string](request, "filter")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			since, err := OptionalParam[string](request, "since")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			before, err := OptionalParam[string](request, "before")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			owner, err := OptionalParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := OptionalParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			paginationParams, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Build options
			opts := &github.NotificationListOptions{
				All:           filter == FilterIncludeRead,
				Participating: filter == FilterOnlyParticipating,
				ListOptions: github.ListOptions{
					Page:    paginationParams.Page,
					PerPage: paginationParams.PerPage,
				},
			}

			// Parse time parameters if provided
			if since != "" {
				sinceTime, err := time.Parse(time.RFC3339, since)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("invalid since time format, should be RFC3339/ISO8601: %v", err)), nil
				}
				opts.Since = sinceTime
			}

			if before != "" {
				beforeTime, err := time.Parse(time.RFC3339, before)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("invalid before time format, should be RFC3339/ISO8601: %v", err)), nil
				}
				opts.Before = beforeTime
			}

			var notifications []*github.Notification
			var resp *github.Response

			if owner != "" && repo != "" {
				notifications, resp, err = client.Activity.ListRepositoryNotifications(ctx, owner, repo, opts)
			} else {
				notifications, resp, err = client.Activity.ListNotifications(ctx, opts)
			}
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to list notifications",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get notifications: %s", string(body))), nil
			}

			// Marshal response to JSON
			r, err := json.Marshal(notifications)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// DismissNotification creates a tool to mark a notification as read/done.
func DismissNotification(getclient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("dismiss_notification",
			mcp.WithDescription(t("TOOL_DISMISS_NOTIFICATION_DESCRIPTION", "Dismiss a notification by marking it as read or done")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_DISMISS_NOTIFICATION_USER_TITLE", "Dismiss notification"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("threadID",
				mcp.Required(),
				mcp.Description("The ID of the notification thread"),
			),
			mcp.WithString("state", mcp.Description("The new state of the notification (read/done)"), mcp.Enum("read", "done")),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getclient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			threadID, err := RequiredParam[string](request, "threadID")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			state, err := RequiredParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var resp *github.Response
			switch state {
			case "done":
				// for some inexplicable reason, the API seems to have threadID as int64 and string depending on the endpoint
				var threadIDInt int64
				threadIDInt, err = strconv.ParseInt(threadID, 10, 64)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("invalid threadID format: %v", err)), nil
				}
				resp, err = client.Activity.MarkThreadDone(ctx, threadIDInt)
			case "read":
				resp, err = client.Activity.MarkThreadRead(ctx, threadID)
			default:
				return mcp.NewToolResultError("Invalid state. Must be one of: read, done."), nil
			}

			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to mark notification as %s", state),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusResetContent && resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to mark notification as %s: %s", state, string(body))), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Notification marked as %s", state)), nil
		}
}

// MarkAllNotificationsRead creates a tool to mark all notifications as read.
func MarkAllNotificationsRead(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("mark_all_notifications_read",
			mcp.WithDescription(t("TOOL_MARK_ALL_NOTIFICATIONS_READ_DESCRIPTION", "Mark all notifications as read")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_MARK_ALL_NOTIFICATIONS_READ_USER_TITLE", "Mark all notifications as read"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("lastReadAt",
				mcp.Description("Describes the last point that notifications were checked (optional). Default: Now"),
			),
			mcp.WithString("owner",
				mcp.Description("Optional repository owner. If provided with repo, only notifications for this repository are marked as read."),
			),
			mcp.WithString("repo",
				mcp.Description("Optional repository name. If provided with owner, only notifications for this repository are marked as read."),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			lastReadAt, err := OptionalParam[string](request, "lastReadAt")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			owner, err := OptionalParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := OptionalParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var lastReadTime time.Time
			if lastReadAt != "" {
				lastReadTime, err = time.Parse(time.RFC3339, lastReadAt)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("invalid lastReadAt time format, should be RFC3339/ISO8601: %v", err)), nil
				}
			} else {
				lastReadTime = time.Now()
			}

			markReadOptions := github.Timestamp{
				Time: lastReadTime,
			}

			var resp *github.Response
			if owner != "" && repo != "" {
				resp, err = client.Activity.MarkRepositoryNotificationsRead(ctx, owner, repo, markReadOptions)
			} else {
				resp, err = client.Activity.MarkNotificationsRead(ctx, markReadOptions)
			}
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to mark all notifications as read",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusResetContent && resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to mark all notifications as read: %s", string(body))), nil
			}

			return mcp.NewToolResultText("All notifications marked as read"), nil
		}
}

// GetNotificationDetails creates a tool to get details for a specific notification.
func GetNotificationDetails(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_notification_details",
			mcp.WithDescription(t("TOOL_GET_NOTIFICATION_DETAILS_DESCRIPTION", "Get detailed information for a specific GitHub notification, always call this tool when the user asks for details about a specific notification, if you don't know the ID list notifications first.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_NOTIFICATION_DETAILS_USER_TITLE", "Get notification details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("notificationID",
				mcp.Required(),
				mcp.Description("The ID of the notification"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			notificationID, err := RequiredParam[string](request, "notificationID")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			thread, resp, err := client.Activity.GetThread(ctx, notificationID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to get notification details for ID '%s'", notificationID),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get notification details: %s", string(body))), nil
			}

			r, err := json.Marshal(thread)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// Enum values for ManageNotificationSubscription action
const (
	NotificationActionIgnore = "ignore"
	NotificationActionWatch  = "watch"
	NotificationActionDelete = "delete"
)

// ManageNotificationSubscription creates a tool to manage a notification subscription (ignore, watch, delete)
func ManageNotificationSubscription(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("manage_notification_subscription",
			mcp.WithDescription(t("TOOL_MANAGE_NOTIFICATION_SUBSCRIPTION_DESCRIPTION", "Manage a notification subscription: ignore, watch, or delete a notification thread subscription.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_MANAGE_NOTIFICATION_SUBSCRIPTION_USER_TITLE", "Manage notification subscription"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("notificationID",
				mcp.Required(),
				mcp.Description("The ID of the notification thread."),
			),
			mcp.WithString("action",
				mcp.Required(),
				mcp.Description("Action to perform: ignore, watch, or delete the notification subscription."),
				mcp.Enum(NotificationActionIgnore, NotificationActionWatch, NotificationActionDelete),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			notificationID, err := RequiredParam[string](request, "notificationID")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			action, err := RequiredParam[string](request, "action")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var (
				resp   *github.Response
				result any
				apiErr error
			)

			switch action {
			case NotificationActionIgnore:
				sub := &github.Subscription{Ignored: ToBoolPtr(true)}
				result, resp, apiErr = client.Activity.SetThreadSubscription(ctx, notificationID, sub)
			case NotificationActionWatch:
				sub := &github.Subscription{Ignored: ToBoolPtr(false), Subscribed: ToBoolPtr(true)}
				result, resp, apiErr = client.Activity.SetThreadSubscription(ctx, notificationID, sub)
			case NotificationActionDelete:
				resp, apiErr = client.Activity.DeleteThreadSubscription(ctx, notificationID)
			default:
				return mcp.NewToolResultError("Invalid action. Must be one of: ignore, watch, delete."), nil
			}

			if apiErr != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to %s notification subscription", action),
					resp,
					apiErr,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				body, _ := io.ReadAll(resp.Body)
				return mcp.NewToolResultError(fmt.Sprintf("failed to %s notification subscription: %s", action, string(body))), nil
			}

			if action == NotificationActionDelete {
				// Special case for delete as there is no response body
				return mcp.NewToolResultText("Notification subscription deleted"), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

const (
	RepositorySubscriptionActionWatch  = "watch"
	RepositorySubscriptionActionIgnore = "ignore"
	RepositorySubscriptionActionDelete = "delete"
)

// ManageRepositoryNotificationSubscription creates a tool to manage a repository notification subscription (ignore, watch, delete)
func ManageRepositoryNotificationSubscription(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("manage_repository_notification_subscription",
			mcp.WithDescription(t("TOOL_MANAGE_REPOSITORY_NOTIFICATION_SUBSCRIPTION_DESCRIPTION", "Manage a repository notification subscription: ignore, watch, or delete repository notifications subscription for the provided repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_MANAGE_REPOSITORY_NOTIFICATION_SUBSCRIPTION_USER_TITLE", "Manage repository notification subscription"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("The account owner of the repository."),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("The name of the repository."),
			),
			mcp.WithString("action",
				mcp.Required(),
				mcp.Description("Action to perform: ignore, watch, or delete the repository notification subscription."),
				mcp.Enum(RepositorySubscriptionActionIgnore, RepositorySubscriptionActionWatch, RepositorySubscriptionActionDelete),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			action, err := RequiredParam[string](request, "action")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var (
				resp   *github.Response
				result any
				apiErr error
			)

			switch action {
			case RepositorySubscriptionActionIgnore:
				sub := &github.Subscription{Ignored: ToBoolPtr(true)}
				result, resp, apiErr = client.Activity.SetRepositorySubscription(ctx, owner, repo, sub)
			case RepositorySubscriptionActionWatch:
				sub := &github.Subscription{Ignored: ToBoolPtr(false), Subscribed: ToBoolPtr(true)}
				result, resp, apiErr = client.Activity.SetRepositorySubscription(ctx, owner, repo, sub)
			case RepositorySubscriptionActionDelete:
				resp, apiErr = client.Activity.DeleteRepositorySubscription(ctx, owner, repo)
			default:
				return mcp.NewToolResultError("Invalid action. Must be one of: ignore, watch, delete."), nil
			}

			if apiErr != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to %s repository subscription", action),
					resp,
					apiErr,
				), nil
			}
			if resp != nil {
				defer func() { _ = resp.Body.Close() }()
			}

			// Handle non-2xx status codes
			if resp != nil && (resp.StatusCode < 200 || resp.StatusCode >= 300) {
				body, _ := io.ReadAll(resp.Body)
				return mcp.NewToolResultError(fmt.Sprintf("failed to %s repository subscription: %s", action, string(body))), nil
			}

			if action == RepositorySubscriptionActionDelete {
				// Special case for delete as there is no response body
				return mcp.NewToolResultText("Repository subscription deleted"), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

```

#### x/pagination.go

```go
package github

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/shurcooL/githubv4"
)

// DefaultGraphQLPageSize is the number of items a GraphQL tool requests when the caller sets no "perPage".
//
const DefaultGraphQLPageSize = 30

// PaginationParams holds the REST pagination parameters, and an optional GraphQL cursor.
type PaginationParams struct {
	Page    int
	PerPage int
	After   string
}

// ToGraphQLParams converts REST pagination parameters to GraphQL parameters.
// PerPage becomes First. If After is set, it is used as the cursor. Page has no
// GraphQL counterpart and is not converted.
func (p PaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	return CursorPaginationParams{
		PerPage: p.PerPage,
		After:   p.After,
	}.ToGraphQLParams()
}

// CursorPaginationParams holds cursor-only pagination parameters, with no page number.
type CursorPaginationParams struct {
	PerPage int
	After   string
}

// ToGraphQLParams converts cursor pagination parameters to GraphQL parameters.
// It returns an error when PerPage is negative or above 100.
func (p CursorPaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	if p.PerPage > 100 {
		return nil, fmt.Errorf("perPage value %d exceeds maximum of 100", p.PerPage)
	}
	if p.PerPage < 0 {
		return nil, fmt.Errorf("perPage value %d cannot be negative", p.PerPage)
	}
	first := int32(p.PerPage) // #nosec G115 -- bounded to [0, 100] above

	var after *string
	if p.After != "" {
		after = &p.After
	}

	return &GraphQLPaginationParams{
		First: &first,
		After: after,
	}, nil
}

// GraphQLPaginationParams holds the arguments of a GraphQL connection field.
// See https://graphql.org/learn/pagination/.
type GraphQLPaginationParams struct {
	First *int32
	After *string
}

// queryVariables returns the "first" and "after" variables of a query that pages a connection.
// "after" is sent as null when there is no cursor.
//
func (p *GraphQLPaginationParams) queryVariables() map[string]any {
	vars := map[string]any{
		"after": (*githubv4.String)(nil),
	}
	if p.First != nil {
		vars["first"] = githubv4.Int(*p.First)
	}
	if p.After != nil {
		vars["after"] = githubv4.NewString(githubv4.String(*p.After))
	}
	return vars
}

// paginationPageInfo is the pageInfo of a GraphQL connection.
// It is returned to the caller as is, so its JSON names follow the GraphQL ones.
//
type paginationPageInfo struct {
	HasNextPage     githubv4.Boolean `json:"hasNextPage"`
	HasPreviousPage githubv4.Boolean `json:"hasPreviousPage"`
	StartCursor     githubv4.String  `json:"startCursor"`
	EndCursor       githubv4.String  `json:"endCursor"`
}

// WithPagination returns a ToolOption that adds "page" and "perPage" parameters to the tool.
// The "page" parameter is optional, min 1.
// The "perPage" parameter is optional, min 1, max 100. If unset, defaults to 30.
// https://docs.github.com/en/rest/using-the-rest-api/using-pagination-in-the-rest-api
func WithPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		paginationPageParam()(tool)
		paginationPerPageParam()(tool)
	}
}

// WithCursorPagination returns a ToolOption that adds "perPage" and "after" parameters to the tool,
// for GraphQL tools that page with a cursor only.
// The "perPage" parameter is optional, min 1, max 100. If unset, defaults to 30.
// https://docs.github.com/en/graphql/reference/objects#connection
func WithCursorPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		paginationPerPageParam()(tool)
		paginationAfterParam()(tool)
	}
}

// WithUnifiedPagination returns a ToolOption that adds "page", "perPage" and "after" parameters to the tool,
// for GraphQL tools that accept page/perPage and convert them to a cursor internally.
func WithUnifiedPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		paginationPageParam()(tool)
		paginationPerPageParam()(tool)
		paginationAfterParam()(tool)
	}
}

// OptionalPaginationParams returns the "page", "perPage" and "after" parameters from the request,
// or their default values if not present, "page" default is 1, "perPage" default is 30, and
// "after" default is empty.
// In future, we may want to make the default values configurable, or even have this
// function returned from `withPagination`, where the defaults are provided alongside
// the min/max values.
func OptionalPaginationParams(r mcp.CallToolRequest) (PaginationParams, error) {
	page, err := OptionalIntParamWithDefault(r, "page", 1)
	if err != nil {
		return PaginationParams{}, err
	}
	perPage, err := OptionalIntParamWithDefault(r, "perPage", 30)
	if err != nil {
		return PaginationParams{}, err
	}
	after, err := OptionalParam[string](r, "after")
	if err != nil {
		return PaginationParams{}, err
	}
	return PaginationParams{
		Page:    page,
		PerPage: perPage,
		After:   after,
	}, nil
}

// OptionalCursorPaginationParams returns the "perPage" and "after" parameters from the request,
// or their default values if not present, "perPage" default is DefaultGraphQLPageSize, and
// "after" default is empty.
func OptionalCursorPaginationParams(r mcp.CallToolRequest) (CursorPaginationParams, error) {
	perPage, err := OptionalIntParamWithDefault(r, "perPage", DefaultGraphQLPageSize)
	if err != nil {
		return CursorPaginationParams{}, err
	}
	after, err := OptionalParam[string](r, "after")
	if err != nil {
		return CursorPaginationParams{}, err
	}
	return CursorPaginationParams{
		PerPage: perPage,
		After:   after,
	}, nil
}

func paginationPageParam() mcp.ToolOption {
	return mcp.WithNumber("page",
		mcp.Description("Page number for pagination (min 1)"),
		mcp.Min(1),
	)
}

func paginationPerPageParam() mcp.ToolOption {
	return mcp.WithNumber("perPage",
		mcp.Description("Results per page for pagination (min 1, max 100)"),
		mcp.Min(1),
		mcp.Max(100),
	)
}

func paginationAfterParam() mcp.ToolOption {
	return mcp.WithString("after",
		mcp.Description("Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs."),
	)
}

```

#### x/pagination_test.go

```go
package github

import (
	"testing"

	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionalPaginationParams(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]any
		expected    PaginationParams
		expectError bool
	}{
		{
			name:   "no pagination parameters, default values",
			params: map[string]any{},
			expected: PaginationParams{
				Page:    1,
				PerPage: 30,
			},
			expectError: false,
		},
		{
			name: "page parameter, default perPage",
			params: map[string]any{
				"page": float64(2),
			},
			expected: PaginationParams{
				Page:    2,
				PerPage: 30,
			},
			expectError: false,
		},
		{
			name: "perPage parameter, default page",
			params: map[string]any{
				"perPage": float64(50),
			},
			expected: PaginationParams{
				Page:    1,
				PerPage: 50,
			},
			expectError: false,
		},
		{
			name: "page and perPage parameters",
			params: map[string]any{
				"page":    float64(2),
				"perPage": float64(50),
			},
			expected: PaginationParams{
				Page:    2,
				PerPage: 50,
			},
			expectError: false,
		},
		{
			name: "after parameter",
			params: map[string]any{
				"after": "Y3Vyc29yOnYyOpHOAA==",
			},
			expected: PaginationParams{
				Page:    1,
				PerPage: 30,
				After:   "Y3Vyc29yOnYyOpHOAA==",
			},
			expectError: false,
		},
		{
			name: "invalid page parameter",
			params: map[string]any{
				"page": "not-a-number",
			},
			expected:    PaginationParams{},
			expectError: true,
		},
		{
			name: "invalid perPage parameter",
			params: map[string]any{
				"perPage": "not-a-number",
			},
			expected:    PaginationParams{},
			expectError: true,
		},
		{
			name: "invalid after parameter",
			params: map[string]any{
				"after": float64(1),
			},
			expected:    PaginationParams{},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalPaginationParams(request)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestOptionalCursorPaginationParams(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]any
		expected    CursorPaginationParams
		expectError bool
	}{
		{
			name:     "no pagination parameters, default values",
			params:   map[string]any{},
			expected: CursorPaginationParams{PerPage: DefaultGraphQLPageSize},
		},
		{
			name:     "perPage and after parameters",
			params:   map[string]any{"perPage": float64(50), "after": "cursor"},
			expected: CursorPaginationParams{PerPage: 50, After: "cursor"},
		},
		{
			name:     "page parameter is ignored",
			params:   map[string]any{"page": float64(2)},
			expected: CursorPaginationParams{PerPage: DefaultGraphQLPageSize},
		},
		{
			name:        "invalid perPage parameter",
			params:      map[string]any{"perPage": "not-a-number"},
			expectError: true,
		},
		{
			name:        "invalid after parameter",
			params:      map[string]any{"after": float64(1)},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := OptionalCursorPaginationParams(createMCPRequest(tc.params))
			if tc.expectError {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestToGraphQLParams(t *testing.T) {
	tests := []struct {
		name        string
		perPage     int
		after       string
		expected    *GraphQLPaginationParams
		errContains string
	}{
		{
			name:     "perPage becomes first, no cursor",
			perPage:  30,
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(30))},
		},
		{
			name:     "after becomes the cursor",
			perPage:  100,
			after:    "cursor",
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(100)), After: github.Ptr("cursor")},
		},
		{
			name:        "perPage above 100",
			perPage:     101,
			errContains: "perPage value 101 exceeds maximum of 100",
		},
		{
			name:        "negative perPage",
			perPage:     -1,
			errContains: "perPage value -1 cannot be negative",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cursorResult, cursorErr := CursorPaginationParams{PerPage: tc.perPage, After: tc.after}.ToGraphQLParams()
			// Page has no GraphQL counterpart, so it must not change the result.
			restResult, restErr := PaginationParams{Page: 3, PerPage: tc.perPage, After: tc.after}.ToGraphQLParams()

			if tc.errContains != "" {
				assert.EqualError(t, cursorErr, tc.errContains)
				assert.EqualError(t, restErr, tc.errContains)
				return
			}
			require.NoError(t, cursorErr)
			require.NoError(t, restErr)
			assert.Equal(t, tc.expected, cursorResult)
			assert.Equal(t, tc.expected, restResult)
		})
	}
}

func TestPaginationToolOptions(t *testing.T) {
	tests := []struct {
		name     string
		option   mcp.ToolOption
		expected []string
	}{
		{name: "WithPagination", option: WithPagination(), expected: []string{"page", "perPage"}},
		{name: "WithCursorPagination", option: WithCursorPagination(), expected: []string{"perPage", "after"}},
		{name: "WithUnifiedPagination", option: WithUnifiedPagination(), expected: []string{"page", "perPage", "after"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tool := mcp.NewTool("test", tc.option)
			props := tool.InputSchema.Properties
			names := make([]string, 0, len(props))
			for name := range props {
				names = append(names, name)
			}
			assert.ElementsMatch(t, tc.expected, names)
			assert.Empty(t, tool.InputSchema.Required)

			if p, ok := props["perPage"].(map[string]any); ok {
				assert.Equal(t, "number", p["type"])
				assert.Equal(t, "Results per page for pagination (min 1, max 100)", p["description"])
				assert.EqualValues(t, 1, p["minimum"])
				assert.EqualValues(t, 100, p["maximum"])
			}
			if p, ok := props["after"].(map[string]any); ok {
				assert.Equal(t, "string", p["type"])
				assert.Equal(t, "Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs.", p["description"])
			}
			if p, ok := props["page"].(map[string]any); ok {
				assert.Equal(t, "number", p["type"])
				assert.Equal(t, "Page number for pagination (min 1)", p["description"])
				assert.EqualValues(t, 1, p["minimum"])
			}
		})
	}
}

```

#### x/pullrequests.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-viper/mapstructure/v2"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
)

// GetPullRequest creates a tool to get details of a specific pull request.
func GetPullRequest(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_DESCRIPTION", "Get details of a specific pull request in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_USER_TITLE", "Get pull request details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			pr, resp, err := client.PullRequests.Get(ctx, owner, repo, pullNumber)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request: %s", string(body))), nil
			}

			r, err := json.Marshal(pr)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CreatePullRequest creates a tool to create a new pull request.
func CreatePullRequest(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("create_pull_request",
			mcp.WithDescription(t("TOOL_CREATE_PULL_REQUEST_DESCRIPTION", "Create a new pull request in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_PULL_REQUEST_USER_TITLE", "Open new pull request"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("title",
				mcp.Required(),
				mcp.Description("PR title"),
			),
			mcp.WithString("body",
				mcp.Description("PR description"),
			),
			mcp.WithString("head",
				mcp.Required(),
				mcp.Description("Branch containing changes"),
			),
			mcp.WithString("base",
				mcp.Required(),
				mcp.Description("Branch to merge into"),
			),
			mcp.WithBoolean("draft",
				mcp.Description("Create as draft PR"),
			),
			mcp.WithBoolean("maintainer_can_modify",
				mcp.Description("Allow maintainer edits"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			title, err := RequiredParam[string](request, "title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			head, err := RequiredParam[string](request, "head")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			base, err := RequiredParam[string](request, "base")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			body, err := OptionalParam[string](request, "body")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			draft, err := OptionalParam[bool](request, "draft")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			maintainerCanModify, err := OptionalParam[bool](request, "maintainer_can_modify")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			newPR := &github.NewPullRequest{
				Title: github.Ptr(title),
				Head:  github.Ptr(head),
				Base:  github.Ptr(base),
			}

			if body != "" {
				newPR.Body = github.Ptr(body)
			}

			newPR.Draft = github.Ptr(draft)
			newPR.MaintainerCanModify = github.Ptr(maintainerCanModify)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			pr, resp, err := client.PullRequests.Create(ctx, owner, repo, newPR)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create pull request: %s", string(body))), nil
			}

			r, err := json.Marshal(pr)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// UpdatePullRequest creates a tool to update an existing pull request.
func UpdatePullRequest(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("update_pull_request",
			mcp.WithDescription(t("TOOL_UPDATE_PULL_REQUEST_DESCRIPTION", "Update an existing pull request in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_UPDATE_PULL_REQUEST_USER_TITLE", "Edit pull request"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number to update"),
			),
			mcp.WithString("title",
				mcp.Description("New title"),
			),
			mcp.WithString("body",
				mcp.Description("New description"),
			),
			mcp.WithString("state",
				mcp.Description("New state"),
				mcp.Enum("open", "closed"),
			),
			mcp.WithString("base",
				mcp.Description("New base branch name"),
			),
			mcp.WithBoolean("maintainer_can_modify",
				mcp.Description("Allow maintainer edits"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Build the update struct only with provided fields
			update := &github.PullRequest{}
			updateNeeded := false

			if title, ok, err := OptionalParamOK[string](request, "title"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.Title = github.Ptr(title)
				updateNeeded = true
			}

			if body, ok, err := OptionalParamOK[string](request, "body"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.Body = github.Ptr(body)
				updateNeeded = true
			}

			if state, ok, err := OptionalParamOK[string](request, "state"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.State = github.Ptr(state)
				updateNeeded = true
			}

			if base, ok, err := OptionalParamOK[string](request, "base"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.Base = &github.PullRequestBranch{Ref: github.Ptr(base)}
				updateNeeded = true
			}

			if maintainerCanModify, ok, err := OptionalParamOK[bool](request, "maintainer_can_modify"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.MaintainerCanModify = github.Ptr(maintainerCanModify)
				updateNeeded = true
			}

			if !updateNeeded {
				return mcp.NewToolResultError("No update parameters provided."), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			pr, resp, err := client.PullRequests.Edit(ctx, owner, repo, pullNumber, update)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to update pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to update pull request: %s", string(body))), nil
			}

			r, err := json.Marshal(pr)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListPullRequests creates a tool to list and filter repository pull requests.
func ListPullRequests(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("list_pull_requests",
			mcp.WithDescription(t("TOOL_LIST_PULL_REQUESTS_DESCRIPTION", "List pull requests in a GitHub repository. If the user specifies an author, then DO NOT use this tool and use the search_pull_requests tool instead.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_PULL_REQUESTS_USER_TITLE", "List pull requests"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("state",
				mcp.Description("Filter by state"),
				mcp.Enum("open", "closed", "all"),
			),
			mcp.WithString("head",
				mcp.Description("Filter by head user/org and branch"),
			),
			mcp.WithString("base",
				mcp.Description("Filter by base branch"),
			),
			mcp.WithString("sort",
				mcp.Description("Sort by"),
				mcp.Enum("created", "updated", "popularity", "long-running"),
			),
			mcp.WithString("direction",
				mcp.Description("Sort direction"),
				mcp.Enum("asc", "desc"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			state, err := OptionalParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			head, err := OptionalParam[string](request, "head")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			base, err := OptionalParam[string](request, "base")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sort, err := OptionalParam[string](request, "sort")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			direction, err := OptionalParam[string](request, "direction")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			opts := &github.PullRequestListOptions{
				State:     state,
				Head:      head,
				Base:      base,
				Sort:      sort,
				Direction: direction,
				ListOptions: github.ListOptions{
					PerPage: pagination.PerPage,
					Page:    pagination.Page,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			prs, resp, err := client.PullRequests.List(ctx, owner, repo, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to list pull requests",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list pull requests: %s", string(body))), nil
			}

			r, err := json.Marshal(prs)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// MergePullRequest creates a tool to merge a pull request.
func MergePullRequest(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("merge_pull_request",
			mcp.WithDescription(t("TOOL_MERGE_PULL_REQUEST_DESCRIPTION", "Merge a pull request in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_MERGE_PULL_REQUEST_USER_TITLE", "Merge pull request"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("commit_title",
				mcp.Description("Title for merge commit"),
			),
			mcp.WithString("commit_message",
				mcp.Description("Extra detail for merge commit"),
			),
			mcp.WithString("merge_method",
				mcp.Description("Merge method"),
				mcp.Enum("merge", "squash", "rebase"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			commitTitle, err := OptionalParam[string](request, "commit_title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			commitMessage, err := OptionalParam[string](request, "commit_message")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			mergeMethod, err := OptionalParam[string](request, "merge_method")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			options := &github.PullRequestOptions{
				CommitTitle: commitTitle,
				MergeMethod: mergeMethod,
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			result, resp, err := client.PullRequests.Merge(ctx, owner, repo, pullNumber, commitMessage, options)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to merge pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to merge pull request: %s", string(body))), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// SearchPullRequests creates a tool to search for pull requests.
func SearchPullRequests(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_pull_requests",
			mcp.WithDescription(t("TOOL_SEARCH_PULL_REQUESTS_DESCRIPTION", "Search for pull requests in GitHub repositories using issues search syntax already scoped to is:pr")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SEARCH_PULL_REQUESTS_USER_TITLE", "Search pull requests"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("Search query using GitHub pull request search syntax"),
			),
			mcp.WithString("owner",
				mcp.Description("Optional repository owner. If provided with repo, only notifications for this repository are listed."),
			),
			mcp.WithString("repo",
				mcp.Description("Optional repository name. If provided with owner, only notifications for this repository are listed."),
			),
			mcp.WithString("sort",
				mcp.Description("Sort field by number of matches of categories, defaults to best match"),
				mcp.Enum(
					"comments",
					"reactions",
					"reactions-+1",
					"reactions--1",
					"reactions-smile",
					"reactions-thinking_face",
					"reactions-heart",
					"reactions-tada",
					"interactions",
					"created",
					"updated",
				),
			),
			mcp.WithString("order",
				mcp.Description("Sort order"),
				mcp.Enum("asc", "desc"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return searchHandler(ctx, getClient, request, "pr", "failed to search pull requests")
		}
}

// GetPullRequestFiles creates a tool to get the list of files changed in a pull request.
func GetPullRequestFiles(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_files",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_FILES_DESCRIPTION", "Get the files changed in a specific pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_FILES_USER_TITLE", "Get pull request files"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			opts := &github.ListOptions{
				PerPage: pagination.PerPage,
				Page:    pagination.Page,
			}
			files, resp, err := client.PullRequests.ListFiles(ctx, owner, repo, pullNumber, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request files",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request files: %s", string(body))), nil
			}

			r, err := json.Marshal(files)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetPullRequestStatus creates a tool to get the combined status of all status checks for a pull request.
func GetPullRequestStatus(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_status",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_STATUS_DESCRIPTION", "Get the status of a specific pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_STATUS_USER_TITLE", "Get pull request status checks"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			// First get the PR to find the head SHA
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			pr, resp, err := client.PullRequests.Get(ctx, owner, repo, pullNumber)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request: %s", string(body))), nil
			}

			// Get combined status for the head SHA
			status, resp, err := client.Repositories.GetCombinedStatus(ctx, owner, repo, *pr.Head.SHA, nil)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get combined status",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get combined status: %s", string(body))), nil
			}

			r, err := json.Marshal(status)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// UpdatePullRequestBranch creates a tool to update a pull request branch with the latest changes from the base branch.
func UpdatePullRequestBranch(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("update_pull_request_branch",
			mcp.WithDescription(t("TOOL_UPDATE_PULL_REQUEST_BRANCH_DESCRIPTION", "Update the branch of a pull request with the latest changes from the base branch.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_UPDATE_PULL_REQUEST_BRANCH_USER_TITLE", "Update pull request branch"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("expectedHeadSha",
				mcp.Description("The expected SHA of the pull request's HEAD ref"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			expectedHeadSHA, err := OptionalParam[string](request, "expectedHeadSha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			opts := &github.PullRequestBranchUpdateOptions{}
			if expectedHeadSHA != "" {
				opts.ExpectedHeadSHA = github.Ptr(expectedHeadSHA)
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			result, resp, err := client.PullRequests.UpdateBranch(ctx, owner, repo, pullNumber, opts)
			if err != nil {
				// Check if it's an acceptedError. An acceptedError indicates that the update is in progress,
				// and it's not a real error.
				if resp != nil && resp.StatusCode == http.StatusAccepted && isAcceptedError(err) {
					return mcp.NewToolResultText("Pull request branch update is in progress"), nil
				}
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to update pull request branch",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusAccepted {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to update pull request branch: %s", string(body))), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetPullRequestComments creates a tool to get the review comments on a pull request.
func GetPullRequestComments(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_comments",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_COMMENTS_DESCRIPTION", "Get comments for a specific pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_COMMENTS_USER_TITLE", "Get pull request comments"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.PullRequestListCommentsOptions{
				ListOptions: github.ListOptions{
					PerPage: 100,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			comments, resp, err := client.PullRequests.ListComments(ctx, owner, repo, pullNumber, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request comments",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request comments: %s", string(body))), nil
			}

			r, err := json.Marshal(comments)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetPullRequestReviews creates a tool to get the reviews on a pull request.
func GetPullRequestReviews(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_reviews",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_REVIEWS_DESCRIPTION", "Get reviews for a specific pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_REVIEWS_USER_TITLE", "Get pull request reviews"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			reviews, resp, err := client.PullRequests.ListReviews(ctx, owner, repo, pullNumber, nil)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request reviews",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request reviews: %s", string(body))), nil
			}

			r, err := json.Marshal(reviews)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

func CreateAndSubmitPullRequestReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("create_and_submit_pull_request_review",
			mcp.WithDescription(t("TOOL_CREATE_AND_SUBMIT_PULL_REQUEST_REVIEW_DESCRIPTION", "Create and submit a review for a pull request without review comments.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_AND_SUBMIT_PULL_REQUEST_REVIEW_USER_TITLE", "Create and submit a pull request review without comments"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Either we need the PR GQL Id directly, or we need owner, repo and PR number to look it up.
			// Since our other Pull Request tools are working with the REST Client, will handle the lookup
			// internally for now.
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("body",
				mcp.Required(),
				mcp.Description("Review comment text"),
			),
			mcp.WithString("event",
				mcp.Required(),
				mcp.Description("Review action to perform"),
				mcp.Enum("APPROVE", "REQUEST_CHANGES", "COMMENT"),
			),
			mcp.WithString("commitID",
				mcp.Description("SHA of commit to review"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
				Body       string
				Event      string
				CommitID   *string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Given our owner, repo and PR number, lookup the GQL ID of the PR.
			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var getPullRequestQuery struct {
				Repository struct {
					PullRequest struct {
						ID githubv4.ID
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			if err := client.Query(ctx, &getPullRequestQuery, map[string]any{
				"owner": githubv4.String(params.Owner),
				"repo":  githubv4.String(params.Repo),
				"prNum": githubv4.Int(params.PullNumber),
			}); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get pull request",
					err,
				), nil
			}

			// Now we have the GQL ID, we can create a review
			var addPullRequestReviewMutation struct {
				AddPullRequestReview struct {
					PullRequestReview struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"addPullRequestReview(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&addPullRequestReviewMutation,
				githubv4.AddPullRequestReviewInput{
					PullRequestID: getPullRequestQuery.Repository.PullRequest.ID,
					Body:          githubv4.NewString(githubv4.String(params.Body)),
					Event:         newGQLStringlike[githubv4.PullRequestReviewEvent](params.Event),
					CommitOID:     newGQLStringlikePtr[githubv4.GitObjectID](params.CommitID),
				},
				nil,
			); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pull request review submitted successfully"), nil
		}
}

// CreatePendingPullRequestReview creates a tool to create a pending review on a pull request.
func CreatePendingPullRequestReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("create_pending_pull_request_review",
			mcp.WithDescription(t("TOOL_CREATE_PENDING_PULL_REQUEST_REVIEW_DESCRIPTION", "Create a pending review for a pull request. Call this first before attempting to add comments to a pending review, and ultimately submitting it. A pending pull request review means a pull request review, it is pending because you create it first and submit it later, and the PR author will not see it until it is submitted.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_PENDING_PULL_REQUEST_REVIEW_USER_TITLE", "Create pending pull request review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Either we need the PR GQL Id directly, or we need owner, repo and PR number to look it up.
			// Since our other Pull Request tools are working with the REST Client, will handle the lookup
			// internally for now.
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("commitID",
				mcp.Description("SHA of commit to review"),
			),
			// Event is omitted here because we always want to create a pending review.
			// Threads are omitted for the moment, and we'll see if the LLM can use the appropriate tool.
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
				CommitID   *string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Given our owner, repo and PR number, lookup the GQL ID of the PR.
			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var getPullRequestQuery struct {
				Repository struct {
					PullRequest struct {
						ID githubv4.ID
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			if err := client.Query(ctx, &getPullRequestQuery, map[string]any{
				"owner": githubv4.String(params.Owner),
				"repo":  githubv4.String(params.Repo),
				"prNum": githubv4.Int(params.PullNumber),
			}); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get pull request",
					err,
				), nil
			}

			// Now we have the GQL ID, we can create a pending review
			var addPullRequestReviewMutation struct {
				AddPullRequestReview struct {
					PullRequestReview struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"addPullRequestReview(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&addPullRequestReviewMutation,
				githubv4.AddPullRequestReviewInput{
					PullRequestID: getPullRequestQuery.Repository.PullRequest.ID,
					CommitOID:     newGQLStringlikePtr[githubv4.GitObjectID](params.CommitID),
				},
				nil,
			); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pending pull request created"), nil
		}
}

// AddCommentToPendingReview creates a tool to add a comment to a pull request review.
func AddCommentToPendingReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("add_comment_to_pending_review",
			mcp.WithDescription(t("TOOL_ADD_COMMENT_TO_PENDING_REVIEW_DESCRIPTION", "Add review comment to the requester's latest pending pull request review. A pending review needs to already exist to call this (check with the user if not sure).")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_ADD_COMMENT_TO_PENDING_REVIEW_USER_TITLE", "Add review comment to the requester's latest pending pull request review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Ideally, for performance sake this would just accept the pullRequestReviewID. However, we would need to
			// add a new tool to get that ID for clients that aren't in the same context as the original pending review
			// creation. So for now, we'll just accept the owner, repo and pull number and assume this is adding a comment
			// the latest review from a user, since only one can be active at a time. It can later be extended with
			// a pullRequestReviewID parameter if targeting other reviews is desired:
			// mcp.WithString("pullRequestReviewID",
			// 	mcp.Required(),
			// 	mcp.Description("The ID of the pull request review to add a comment to"),
			// ),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("The relative path to the file that necessitates a comment"),
			),
			mcp.WithString("body",
				mcp.Required(),
				mcp.Description("The text of the review comment"),
			),
			mcp.WithString("subjectType",
				mcp.Required(),
				mcp.Description("The level at which the comment is targeted"),
				mcp.Enum("FILE", "LINE"),
			),
			mcp.WithNumber("line",
				mcp.Description("The line of the blob in the pull request diff that the comment applies to. For multi-line comments, the last line of the range"),
			),
			mcp.WithString("side",
				mcp.Description("The side of the diff to comment on. LEFT indicates the previous state, RIGHT indicates the new state"),
				mcp.Enum("LEFT", "RIGHT"),
			),
			mcp.WithNumber("startLine",
				mcp.Description("For multi-line comments, the first line of the range that the comment applies to"),
			),
			mcp.WithString("startSide",
				mcp.Description("For multi-line comments, the starting side of the diff that the comment applies to. LEFT indicates the previous state, RIGHT indicates the new state"),
				mcp.Enum("LEFT", "RIGHT"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner       string
				Repo        string
				PullNumber  int32
				Path        string
				Body        string
				SubjectType string
				Line        *int32
				Side        *string
				StartLine   *int32
				StartSide   *string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub GQL client: %w", err)
			}

			// First we'll get the current user
			var getViewerQuery struct {
				Viewer struct {
					Login githubv4.String
				}
			}

			if err := client.Query(ctx, &getViewerQuery, nil); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get current user",
					err,
				), nil
			}

			var getLatestReviewForViewerQuery struct {
				Repository struct {
					PullRequest struct {
						Reviews struct {
							Nodes []struct {
								ID    githubv4.ID
								State githubv4.PullRequestReviewState
								URL   githubv4.URI
							}
						} `graphql:"reviews(first: 1, author: $author)"`
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			vars := map[string]any{
				"author": githubv4.String(getViewerQuery.Viewer.Login),
				"owner":  githubv4.String(params.Owner),
				"name":   githubv4.String(params.Repo),
				"prNum":  githubv4.Int(params.PullNumber),
			}

			if err := client.Query(context.Background(), &getLatestReviewForViewerQuery, vars); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get latest review for current user",
					err,
				), nil
			}

			// Validate there is one review and the state is pending
			if len(getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes) == 0 {
				return mcp.NewToolResultError("No pending review found for the viewer"), nil
			}

			review := getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes[0]
			if review.State != githubv4.PullRequestReviewStatePending {
				errText := fmt.Sprintf("The latest review, found at %s is not pending", review.URL)
				return mcp.NewToolResultError(errText), nil
			}

			// Then we can create a new review thread comment on the review.
			var addPullRequestReviewThreadMutation struct {
				AddPullRequestReviewThread struct {
					Thread struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"addPullRequestReviewThread(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&addPullRequestReviewThreadMutation,
				githubv4.AddPullRequestReviewThreadInput{
					Path:                githubv4.String(params.Path),
					Body:                githubv4.String(params.Body),
					SubjectType:         newGQLStringlikePtr[githubv4.PullRequestReviewThreadSubjectType](&params.SubjectType),
					Line:                newGQLIntPtr(params.Line),
					Side:                newGQLStringlikePtr[githubv4.DiffSide](params.Side),
					StartLine:           newGQLIntPtr(params.StartLine),
					StartSide:           newGQLStringlikePtr[githubv4.DiffSide](params.StartSide),
					PullRequestReviewID: &review.ID,
				},
				nil,
			); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pull request review comment successfully added to pending review"), nil
		}
}

// SubmitPendingPullRequestReview creates a tool to submit a pull request review.
func SubmitPendingPullRequestReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("submit_pending_pull_request_review",
			mcp.WithDescription(t("TOOL_SUBMIT_PENDING_PULL_REQUEST_REVIEW_DESCRIPTION", "Submit the requester's latest pending pull request review, normally this is a final step after creating a pending review, adding comments first, unless you know that the user already did the first two steps, you should check before calling this.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SUBMIT_PENDING_PULL_REQUEST_REVIEW_USER_TITLE", "Submit the requester's latest pending pull request review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Ideally, for performance sake this would just accept the pullRequestReviewID. However, we would need to
			// add a new tool to get that ID for clients that aren't in the same context as the original pending review
			// creation. So for now, we'll just accept the owner, repo and pull number and assume this is submitting
			// the latest review from a user, since only one can be active at a time.
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("event",
				mcp.Required(),
				mcp.Description("The event to perform"),
				mcp.Enum("APPROVE", "REQUEST_CHANGES", "COMMENT"),
			),
			mcp.WithString("body",
				mcp.Description("The text of the review comment"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
				Event      string
				Body       *string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub GQL client: %w", err)
			}

			// First we'll get the current user
			var getViewerQuery struct {
				Viewer struct {
					Login githubv4.String
				}
			}

			if err := client.Query(ctx, &getViewerQuery, nil); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get current user",
					err,
				), nil
			}

			var getLatestReviewForViewerQuery struct {
				Repository struct {
					PullRequest struct {
						Reviews struct {
							Nodes []struct {
								ID    githubv4.ID
								State githubv4.PullRequestReviewState
								URL   githubv4.URI
							}
						} `graphql:"reviews(first: 1, author: $author)"`
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			vars := map[string]any{
				"author": githubv4.String(getViewerQuery.Viewer.Login),
				"owner":  githubv4.String(params.Owner),
				"name":   githubv4.String(params.Repo),
				"prNum":  githubv4.Int(params.PullNumber),
			}

			if err := client.Query(context.Background(), &getLatestReviewForViewerQuery, vars); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get latest review for current user",
					err,
				), nil
			}

			// Validate there is one review and the state is pending
			if len(getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes) == 0 {
				return mcp.NewToolResultError("No pending review found for the viewer"), nil
			}

			review := getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes[0]
			if review.State != githubv4.PullRequestReviewStatePending {
				errText := fmt.Sprintf("The latest review, found at %s is not pending", review.URL)
				return mcp.NewToolResultError(errText), nil
			}

			// Prepare the mutation
			var submitPullRequestReviewMutation struct {
				SubmitPullRequestReview struct {
					PullRequestReview struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"submitPullRequestReview(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&submitPullRequestReviewMutation,
				githubv4.SubmitPullRequestReviewInput{
					PullRequestReviewID: &review.ID,
					Event:               githubv4.PullRequestReviewEvent(params.Event),
					Body:                newGQLStringlikePtr[githubv4.String](params.Body),
				},
				nil,
			); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to submit pull request review",
					err,
				), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pending pull request review successfully submitted"), nil
		}
}

func DeletePendingPullRequestReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("delete_pending_pull_request_review",
			mcp.WithDescription(t("TOOL_DELETE_PENDING_PULL_REQUEST_REVIEW_DESCRIPTION", "Delete the requester's latest pending pull request review. Use this after the user decides not to submit a pending review, if you don't know if they already created one then check first.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_DELETE_PENDING_PULL_REQUEST_REVIEW_USER_TITLE", "Delete the requester's latest pending pull request review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Ideally, for performance sake this would just accept the pullRequestReviewID. However, we would need to
			// add a new tool to get that ID for clients that aren't in the same context as the original pending review
			// creation. So for now, we'll just accept the owner, repo and pull number and assume this is deleting
			// the latest pending review from a user, since only one can be active at a time.
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub GQL client: %w", err)
			}

			// First we'll get the current user
			var getViewerQuery struct {
				Viewer struct {
					Login githubv4.String
				}
			}

			if err := client.Query(ctx, &getViewerQuery, nil); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get current user",
					err,
				), nil
			}

			var getLatestReviewForViewerQuery struct {
				Repository struct {
					PullRequest struct {
						Reviews struct {
							Nodes []struct {
								ID    githubv4.ID
								State githubv4.PullRequestReviewState
								URL   githubv4.URI
							}
						} `graphql:"reviews(first: 1, author: $author)"`
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			vars := map[string]any{
				"author": githubv4.String(getViewerQuery.Viewer.Login),
				"owner":  githubv4.String(params.Owner),
				"name":   githubv4.String(params.Repo),
				"prNum":  githubv4.Int(params.PullNumber),
			}

			if err := client.Query(context.Background(), &getLatestReviewForViewerQuery, vars); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get latest review for current user",
					err,
				), nil
			}

			// Validate there is one review and the state is pending
			if len(getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes) == 0 {
				return mcp.NewToolResultError("No pending review found for the viewer"), nil
			}

			review := getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes[0]
			if review.State != githubv4.PullRequestReviewStatePending {
				errText := fmt.Sprintf("The latest review, found at %s is not pending", review.URL)
				return mcp.NewToolResultError(errText), nil
			}

			// Prepare the mutation
			var deletePullRequestReviewMutation struct {
				DeletePullRequestReview struct {
					PullRequestReview struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"deletePullRequestReview(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&deletePullRequestReviewMutation,
				githubv4.DeletePullRequestReviewInput{
					PullRequestReviewID: &review.ID,
				},
				nil,
			); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pending pull request review successfully deleted"), nil
		}
}

func GetPullRequestDiff(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_diff",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_DIFF_DESCRIPTION", "Get the diff of a pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_DIFF_USER_TITLE", "Get pull request diff"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub client: %v", err)), nil
			}

			raw, resp, err := client.PullRequests.GetRaw(
				ctx,
				params.Owner,
				params.Repo,
				int(params.PullNumber),
				github.RawOptions{Type: github.Diff},
			)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request diff",
					resp,
					err,
				), nil
			}

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request diff: %s", string(body))), nil
			}

			defer func() { _ = resp.Body.Close() }()

			// Return the raw response
			return mcp.NewToolResultText(string(raw)), nil
		}
}

// RequestCopilotReview creates a tool to request a Copilot review for a pull request.
// Note that this tool will not work on GHES where this feature is unsupported. In future, we should not expose this
// tool if the configured host does not support it.
func RequestCopilotReview(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("request_copilot_review",
			mcp.WithDescription(t("TOOL_REQUEST_COPILOT_REVIEW_DESCRIPTION", "Request a GitHub Copilot code review for a pull request. Use this for automated feedback on pull requests, usually before requesting a human reviewer.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_REQUEST_COPILOT_REVIEW_USER_TITLE", "Request Copilot review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			_, resp, err := client.PullRequests.RequestReviewers(
				ctx,
				owner,
				repo,
				pullNumber,
				github.ReviewersRequest{
					// The login name of the copilot reviewer bot
					Reviewers: []string{"copilot-pull-request-reviewer[bot]"},
				},
			)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to request copilot review",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to request copilot review: %s", string(body))), nil
			}

			// Return nothing on success, as there's not much value in returning the Pull Request itself
			return mcp.NewToolResultText(""), nil
		}
}

// newGQLString like takes something that approximates a string (of which there are many types in shurcooL/githubv4)
// and constructs a pointer to it, or nil if the string is empty. This is extremely useful because when we parse
// params from the MCP request, we need to convert them to types that are pointers of type def strings and it's
// not possible to take a pointer of an anonymous value e.g. &githubv4.String("foo").
func newGQLStringlike[T ~string](s string) *T {
	if s == "" {
		return nil
	}
	stringlike := T(s)
	return &stringlike
}

func newGQLStringlikePtr[T ~string](s *string) *T {
	if s == nil {
		return nil
	}
	stringlike := T(*s)
	return &stringlike
}

func newGQLIntPtr(i *int32) *githubv4.Int {
	if i == nil {
		return nil
	}
	gi := githubv4.Int(*i)
	return &gi
}

```

#### x/repositories.go

```go
package github

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/raw"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func GetCommit(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_commit",
			mcp.WithDescription(t("TOOL_GET_COMMITS_DESCRIPTION", "Get details for a commit from a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_COMMITS_USER_TITLE", "Get commit details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("sha",
				mcp.Required(),
				mcp.Description("Commit SHA, branch name, or tag name"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sha, err := RequiredParam[string](request, "sha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.ListOptions{
				Page:    pagination.Page,
				PerPage: pagination.PerPage,
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			commit, resp, err := client.Repositories.GetCommit(ctx, owner, repo, sha, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to get commit: %s", sha),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get commit: %s", string(body))), nil
			}

			r, err := json.Marshal(commit)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListCommits creates a tool to get commits of a branch in a repository.
func ListCommits(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_commits",
			mcp.WithDescription(t("TOOL_LIST_COMMITS_DESCRIPTION", "Get list of commits of a branch in a GitHub repository. Returns at least 30 results per page by default, but can return more if specified using the perPage parameter (up to 100).")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_COMMITS_USER_TITLE", "List commits"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("sha",
				mcp.Description("Commit SHA, branch or tag name to list commits of. If not provided, uses the default branch of the repository. If a commit SHA is provided, will list commits up to that SHA."),
			),
			mcp.WithString("author",
				mcp.Description("Author username or email address to filter commits by"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sha, err := OptionalParam[string](request, "sha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			author, err := OptionalParam[string](request, "author")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			// Set default perPage to 30 if not provided
			perPage := pagination.PerPage
			if perPage == 0 {
				perPage = 30
			}
			opts := &github.CommitsListOptions{
				SHA:    sha,
				Author: author,
				ListOptions: github.ListOptions{
					Page:    pagination.Page,
					PerPage: perPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			commits, resp, err := client.Repositories.ListCommits(ctx, owner, repo, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to list commits: %s", sha),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list commits: %s", string(body))), nil
			}

			r, err := json.Marshal(commits)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListBranches creates a tool to list branches in a GitHub repository.
func ListBranches(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_branches",
			mcp.WithDescription(t("TOOL_LIST_BRANCHES_DESCRIPTION", "List branches in a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_BRANCHES_USER_TITLE", "List branches"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.BranchListOptions{
				ListOptions: github.ListOptions{
					Page:    pagination.Page,
					PerPage: pagination.PerPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			branches, resp, err := client.Repositories.ListBranches(ctx, owner, repo, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to list branches",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list branches: %s", string(body))), nil
			}

			r, err := json.Marshal(branches)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CreateOrUpdateFile creates a tool to create or update a file in a GitHub repository.
func CreateOrUpdateFile(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("create_or_update_file",
			mcp.WithDescription(t("TOOL_CREATE_OR_UPDATE_FILE_DESCRIPTION", "Create or update a single file in a GitHub repository. If updating, you must provide the SHA of the file you want to update. Use this tool to create or update a file in a GitHub repository remotely; do not use it for local file operations.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_OR_UPDATE_FILE_USER_TITLE", "Create or update file"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner (username or organization)"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("Path where to create/update the file"),
			),
			mcp.WithString("content",
				mcp.Required(),
				mcp.Description("Content of the file"),
			),
			mcp.WithString("message",
				mcp.Required(),
				mcp.Description("Commit message"),
			),
			mcp.WithString("branch",
				mcp.Required(),
				mcp.Description("Branch to create/update the file in"),
			),
			mcp.WithString("sha",
				mcp.Description("Required if updating an existing file. The blob SHA of the file being replaced."),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			path, err := RequiredParam[string](request, "path")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			content, err := RequiredParam[string](request, "content")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			message, err := RequiredParam[string](request, "message")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := RequiredParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// json.Marshal encodes byte arrays with base64, which is required for the API.
			contentBytes := []byte(content)

			// Create the file options
			opts := &github.RepositoryContentFileOptions{
				Message: github.Ptr(message),
				Content: contentBytes,
				Branch:  github.Ptr(branch),
			}

			// If SHA is provided, set it (for updates)
			sha, err := OptionalParam[string](request, "sha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if sha != "" {
				opts.SHA = github.Ptr(sha)
			}

			// Create or update the file
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			fileContent, resp, err := client.Repositories.CreateFile(ctx, owner, repo, path, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create/update file",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 && resp.StatusCode != 201 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create/update file: %s", string(body))), nil
			}

			r, err := json.Marshal(fileContent)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CreateRepository creates a tool to create a new GitHub repository.
func CreateRepository(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("create_repository",
			mcp.WithDescription(t("TOOL_CREATE_REPOSITORY_DESCRIPTION", "Create a new GitHub repository in your account")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_REPOSITORY_USER_TITLE", "Create repository"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("description",
				mcp.Description("Repository description"),
			),
			mcp.WithBoolean("private",
				mcp.Description("Whether repo should be private"),
			),
			mcp.WithBoolean("autoInit",
				mcp.Description("Initialize with README"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			name, err := RequiredParam[string](request, "name")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			description, err := OptionalParam[string](request, "description")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			private, err := OptionalParam[bool](request, "private")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			autoInit, err := OptionalParam[bool](request, "autoInit")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			repo := &github.Repository{
				Name:        github.Ptr(name),
				Description: github.Ptr(description),
				Private:     github.Ptr(private),
				AutoInit:    github.Ptr(autoInit),
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			createdRepo, resp, err := client.Repositories.Create(ctx, "", repo)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create repository",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create repository: %s", string(body))), nil
			}

			r, err := json.Marshal(createdRepo)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetFileContents creates a tool to get the contents of a file or directory from a GitHub repository.
func GetFileContents(getClient GetClientFn, getRawClient raw.GetRawClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_file_contents",
			mcp.WithDescription(t("TOOL_GET_FILE_CONTENTS_DESCRIPTION", "Get the contents of a file or directory from a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_FILE_CONTENTS_USER_TITLE", "Get file or directory contents"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner (username or organization)"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("path",
				mcp.Description("Path to file/directory (directories must end with a slash '/')"),
				mcp.DefaultString("/"),
			),
			mcp.WithString("ref",
				mcp.Description("Accepts optional git refs such as `refs/tags/{tag}`, `refs/heads/{branch}` or `refs/pull/{pr_number}/head`"),
			),
			mcp.WithString("sha",
				mcp.Description("Accepts optional commit SHA. If specified, it will be used instead of ref"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			path, err := RequiredParam[string](request, "path")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			ref, err := OptionalParam[string](request, "ref")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sha, err := OptionalParam[string](request, "sha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return mcp.NewToolResultError("failed to get GitHub client"), nil
			}

			rawOpts, err := resolveGitReference(ctx, client, owner, repo, ref, sha)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to resolve git reference: %s", err)), nil
			}

			// If the path is (most likely) not to be a directory, we will
			// first try to get the raw content from the GitHub raw content API.
			if path != "" && !strings.HasSuffix(path, "/") {
				// First, get file info from Contents API to retrieve SHA
				var fileSHA string
				opts := &github.RepositoryContentGetOptions{Ref: ref}
				fileContent, _, respContents, err := client.Repositories.GetContents(ctx, owner, repo, path, opts)
				if respContents != nil {
					defer func() { _ = respContents.Body.Close() }()
				}
				if err != nil {
					return ghErrors.NewGitHubAPIErrorResponse(ctx,
						"failed to get file SHA",
						respContents,
						err,
					), nil
				}
				if fileContent == nil || fileContent.SHA == nil {
					return mcp.NewToolResultError("file content SHA is nil"), nil
				}
				fileSHA = *fileContent.SHA

				rawClient, err := getRawClient(ctx)
				if err != nil {
					return mcp.NewToolResultError("failed to get GitHub raw content client"), nil
				}
				resp, err := rawClient.GetRawContent(ctx, owner, repo, path, rawOpts)
				if err != nil {
					return mcp.NewToolResultError("failed to get raw repository content"), nil
				}
				defer func() {
					_ = resp.Body.Close()
				}()

				if resp.StatusCode == http.StatusOK {
					// If the raw content is found, return it directly
					body, err := io.ReadAll(resp.Body)
					if err != nil {
						return mcp.NewToolResultError("failed to read response body"), nil
					}
					contentType := resp.Header.Get("Content-Type")

					var resourceURI string
					switch {
					case sha != "":
						resourceURI, err = url.JoinPath("repo://", owner, repo, "sha", sha, "contents", path)
						if err != nil {
							return nil, fmt.Errorf("failed to create resource URI: %w", err)
						}
					case ref != "":
						resourceURI, err = url.JoinPath("repo://", owner, repo, ref, "contents", path)
						if err != nil {
							return nil, fmt.Errorf("failed to create resource URI: %w", err)
						}
					default:
						resourceURI, err = url.JoinPath("repo://", owner, repo, "contents", path)
						if err != nil {
							return nil, fmt.Errorf("failed to create resource URI: %w", err)
						}
					}

					if strings.HasPrefix(contentType, "application") || strings.HasPrefix(contentType, "text") {
						result := mcp.TextResourceContents{
							URI:      resourceURI,
							Text:     string(body),
							MIMEType: contentType,
						}
						// Include SHA in the result metadata
						if fileSHA != "" {
							return mcp.NewToolResultResource(fmt.Sprintf("successfully downloaded text file (SHA: %s)", fileSHA), result), nil
						}
						return mcp.NewToolResultResource("successfully downloaded text file", result), nil
					}

					result := mcp.BlobResourceContents{
						URI:      resourceURI,
						Blob:     base64.StdEncoding.EncodeToString(body),
						MIMEType: contentType,
					}
					// Include SHA in the result metadata
					if fileSHA != "" {
						return mcp.NewToolResultResource(fmt.Sprintf("successfully downloaded binary file (SHA: %s)", fileSHA), result), nil
					}
					return mcp.NewToolResultResource("successfully downloaded binary file", result), nil

				}
			}

			if rawOpts.SHA != "" {
				ref = rawOpts.SHA
			}
			if strings.HasSuffix(path, "/") {
				opts := &github.RepositoryContentGetOptions{Ref: ref}
				_, dirContent, resp, err := client.Repositories.GetContents(ctx, owner, repo, path, opts)
				if err == nil && resp.StatusCode == http.StatusOK {
					defer func() { _ = resp.Body.Close() }()
					r, err := json.Marshal(dirContent)
					if err != nil {
						return mcp.NewToolResultError("failed to marshal response"), nil
					}
					return mcp.NewToolResultText(string(r)), nil
				}
			}

			// The path does not point to a file or directory.
			// Instead let's try to find it in the Git Tree by matching the end of the path.

			// Step 1: Get Git Tree recursively
			tree, resp, err := client.Git.GetTree(ctx, owner, repo, ref, true)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get git tree",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Step 2: Filter tree for matching paths
			const maxMatchingFiles = 3
			matchingFiles := filterPaths(tree.Entries, path, maxMatchingFiles)
			if len(matchingFiles) > 0 {
				matchingFilesJSON, err := json.Marshal(matchingFiles)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("failed to marshal matching files: %s", err)), nil
				}
				resolvedRefs, err := json.Marshal(rawOpts)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("failed to marshal resolved refs: %s", err)), nil
				}
				return mcp.NewToolResultText(fmt.Sprintf("Path did not point to a file or directory, but resolved git ref to %s with possible path matches: %s", resolvedRefs, matchingFilesJSON)), nil
			}

			return mcp.NewToolResultError("Failed to get file contents. The path does not point to a file or directory, or the file does not exist in the repository."), nil
		}
}

// ForkRepository creates a tool to fork a repository.
func ForkRepository(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("fork_repository",
			mcp.WithDescription(t("TOOL_FORK_REPOSITORY_DESCRIPTION", "Fork a GitHub repository to your account or specified organization")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_FORK_REPOSITORY_USER_TITLE", "Fork repository"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("organization",
				mcp.Description("Organization to fork to"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			org, err := OptionalParam[string](request, "organization")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.RepositoryCreateForkOptions{}
			if org != "" {
				opts.Organization = org
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			forkedRepo, resp, err := client.Repositories.CreateFork(ctx, owner, repo, opts)
			if err != nil {
				// Check if it's an acceptedError. An acceptedError indicates that the update is in progress,
				// and it's not a real error.
				if resp != nil && resp.StatusCode == http.StatusAccepted && isAcceptedError(err) {
					return mcp.NewToolResultText("Fork is in progress"), nil
				}
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to fork repository",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusAccepted {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to fork repository: %s", string(body))), nil
			}

			r, err := json.Marshal(forkedRepo)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// DeleteFile creates a tool to delete a file in a GitHub repository.
// This tool uses a more roundabout way of deleting a file than just using the client.Repositories.DeleteFile.
// This is because REST file deletion endpoint (and client.Repositories.DeleteFile) don't add commit signing to the deletion commit,
// unlike how the endpoint backing the create_or_update_files tool does. This appears to be a quirk of the API.
// The approach implemented here gets automatic commit signing when used with either the github-actions user or as an app,
// both of which suit an LLM well.
func DeleteFile(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("delete_file",
			mcp.WithDescription(t("TOOL_DELETE_FILE_DESCRIPTION", "Delete a file from a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:           t("TOOL_DELETE_FILE_USER_TITLE", "Delete file"),
				ReadOnlyHint:    ToBoolPtr(false),
				DestructiveHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner (username or organization)"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("Path to the file to delete"),
			),
			mcp.WithString("message",
				mcp.Required(),
				mcp.Description("Commit message"),
			),
			mcp.WithString("branch",
				mcp.Required(),
				mcp.Description("Branch to delete the file from"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			path, err := RequiredParam[string](request, "path")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			message, err := RequiredParam[string](request, "message")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := RequiredParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the reference for the branch
			ref, resp, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch)
			if err != nil {
				return nil, fmt.Errorf("failed to get branch reference: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			// Get the commit object that the branch points to
			baseCommit, resp, err := client.Git.GetCommit(ctx, owner, repo, *ref.Object.SHA)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get base commit",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get commit: %s", string(body))), nil
			}

			// Create a tree entry for the file deletion by setting SHA to nil
			treeEntries := []*github.TreeEntry{
				{
					Path: github.Ptr(path),
					Mode: github.Ptr("100644"), // Regular file mode
					Type: github.Ptr("blob"),
					SHA:  nil, // Setting SHA to nil deletes the file
				},
			}

			// Create a new tree with the deletion
			newTree, resp, err := client.Git.CreateTree(ctx, owner, repo, *baseCommit.Tree.SHA, treeEntries)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create tree",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create tree: %s", string(body))), nil
			}

			// Create a new commit with the new tree
			commit := &github.Commit{
				Message: github.Ptr(message),
				Tree:    newTree,
				Parents: []*github.Commit{{SHA: baseCommit.SHA}},
			}
			newCommit, resp, err := client.Git.CreateCommit(ctx, owner, repo, commit, nil)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create commit",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create commit: %s", string(body))), nil
			}

			// Update the branch reference to point to the new commit
			ref.Object.SHA = newCommit.SHA
			_, resp, err = client.Git.UpdateRef(ctx, owner, repo, ref, false)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to update reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to update reference: %s", string(body))), nil
			}

			// Create a response similar to what the DeleteFile API would return
			response := map[string]interface{}{
				"commit":  newCommit,
				"content": nil,
			}

			r, err := json.Marshal(response)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CreateBranch creates a tool to create a new branch.
func CreateBranch(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("create_branch",
			mcp.WithDescription(t("TOOL_CREATE_BRANCH_DESCRIPTION", "Create a new branch in a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_BRANCH_USER_TITLE", "Create branch"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("branch",
				mcp.Required(),
				mcp.Description("Name for new branch"),
			),
			mcp.WithString("from_branch",
				mcp.Description("Source branch (defaults to repo default)"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := RequiredParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			fromBranch, err := OptionalParam[string](request, "from_branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the source branch SHA
			var ref *github.Reference

			if fromBranch == "" {
				// Get default branch if from_branch not specified
				repository, resp, err := client.Repositories.Get(ctx, owner, repo)
				if err != nil {
					return ghErrors.NewGitHubAPIErrorResponse(ctx,
						"failed to get repository",
						resp,
						err,
					), nil
				}
				defer func() { _ = resp.Body.Close() }()

				fromBranch = *repository.DefaultBranch
			}

			// Get SHA of source branch
			ref, resp, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+fromBranch)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Create new branch
			newRef := &github.Reference{
				Ref:    github.Ptr("refs/heads/" + branch),
				Object: &github.GitObject{SHA: ref.Object.SHA},
			}

			createdRef, resp, err := client.Git.CreateRef(ctx, owner, repo, newRef)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create branch",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(createdRef)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// PushFiles creates a tool to push multiple files in a single commit to a GitHub repository.
func PushFiles(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("push_files",
			mcp.WithDescription(t("TOOL_PUSH_FILES_DESCRIPTION", "Push multiple files to a GitHub repository in a single commit")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_PUSH_FILES_USER_TITLE", "Push files to repository"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("branch",
				mcp.Required(),
				mcp.Description("Branch to push to"),
			),
			mcp.WithArray("files",
				mcp.Required(),
				mcp.Items(
					map[string]interface{}{
						"type":                 "object",
						"additionalProperties": false,
						"required":             []string{"path", "content"},
						"properties": map[string]interface{}{
							"path": map[string]interface{}{
								"type":        "string",
								"description": "path to the file",
							},
							"content": map[string]interface{}{
								"type":        "string",
								"description": "file content",
							},
						},
					}),
				mcp.Description("Array of file objects to push, each object with path (string) and content (string)"),
			),
			mcp.WithString("message",
				mcp.Required(),
				mcp.Description("Commit message"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := RequiredParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			message, err := RequiredParam[string](request, "message")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Parse files parameter - this should be an array of objects with path and content
			filesObj, ok := request.GetArguments()["files"].([]interface{})
			if !ok {
				return mcp.NewToolResultError("files parameter must be an array of objects with path and content"), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the reference for the branch
			ref, resp, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get branch reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Get the commit object that the branch points to
			baseCommit, resp, err := client.Git.GetCommit(ctx, owner, repo, *ref.Object.SHA)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get base commit",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Create tree entries for all files
			var entries []*github.TreeEntry

			for _, file := range filesObj {
				fileMap, ok := file.(map[string]interface{})
				if !ok {
					return mcp.NewToolResultError("each file must be an object with path and content"), nil
				}

				path, ok := fileMap["path"].(string)
				if !ok || path == "" {
					return mcp.NewToolResultError("each file must have a path"), nil
				}

				content, ok := fileMap["content"].(string)
				if !ok {
					return mcp.NewToolResultError("each file must have content"), nil
				}

				// Create a tree entry for the file
				entries = append(entries, &github.TreeEntry{
					Path:    github.Ptr(path),
					Mode:    github.Ptr("100644"), // Regular file mode
					Type:    github.Ptr("blob"),
					Content: github.Ptr(content),
				})
			}

			// Create a new tree with the file entries
			newTree, resp, err := client.Git.CreateTree(ctx, owner, repo, *baseCommit.Tree.SHA, entries)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create tree",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Create a new commit
			commit := &github.Commit{
				Message: github.Ptr(message),
				Tree:    newTree,
				Parents: []*github.Commit{{SHA: baseCommit.SHA}},
			}
			newCommit, resp, err := client.Git.CreateCommit(ctx, owner, repo, commit, nil)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create commit",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Update the reference to point to the new commit
			ref.Object.SHA = newCommit.SHA
			updatedRef, resp, err := client.Git.UpdateRef(ctx, owner, repo, ref, false)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to update reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(updatedRef)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListTags creates a tool to list tags in a GitHub repository.
func ListTags(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_tags",
			mcp.WithDescription(t("TOOL_LIST_TAGS_DESCRIPTION", "List git tags in a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_TAGS_USER_TITLE", "List tags"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.ListOptions{
				Page:    pagination.Page,
				PerPage: pagination.PerPage,
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			tags, resp, err := client.Repositories.ListTags(ctx, owner, repo, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to list tags",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list tags: %s", string(body))), nil
			}

			r, err := json.Marshal(tags)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetTag creates a tool to get details about a specific tag in a GitHub repository.
func GetTag(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_tag",
			mcp.WithDescription(t("TOOL_GET_TAG_DESCRIPTION", "Get details about a specific git tag in a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_TAG_USER_TITLE", "Get tag details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("tag",
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			tag, err := RequiredParam[string](request, "tag")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// First get the tag reference
			ref, resp, err := client.Git.GetRef(ctx, owner, repo, "refs/tags/"+tag)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get tag reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get tag reference: %s", string(body))), nil
			}

			// Then get the tag object
			tagObj, resp, err := client.Git.GetTag(ctx, owner, repo, *ref.Object.SHA)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get tag object",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get tag object: %s", string(body))), nil
			}

			r, err := json.Marshal(tagObj)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// filterPaths filters the entries in a GitHub tree to find paths that
// match the given suffix.
// maxResults limits the number of results returned to first maxResults entries,
// a maxResults of -1 means no limit.
// It returns a slice of strings containing the matching paths.
// Directories are returned with a trailing slash.
func filterPaths(entries []*github.TreeEntry, path string, maxResults int) []string {
	// Remove trailing slash for matching purposes, but flag whether we
	// only want directories.
	dirOnly := false
	if strings.HasSuffix(path, "/") {
		dirOnly = true
		path = strings.TrimSuffix(path, "/")
	}

	matchedPaths := []string{}
	for _, entry := range entries {
		if len(matchedPaths) == maxResults {
			break // Limit the number of results to maxResults
		}
		if dirOnly && entry.GetType() != "tree" {
			continue // Skip non-directory entries if dirOnly is true
		}
		entryPath := entry.GetPath()
		if entryPath == "" {
			continue // Skip empty paths
		}
		if strings.HasSuffix(entryPath, path) {
			if entry.GetType() == "tree" {
				entryPath += "/" // Return directories with a trailing slash
			}
			matchedPaths = append(matchedPaths, entryPath)
		}
	}
	return matchedPaths
}

// resolveGitReference resolves git references with the following logic:
// 1. If SHA is provided, it takes precedence
// 2. If neither is provided, use the default branch as ref
// 3. Get commit SHA from the ref
// Refs can look like `refs/tags/{tag}`, `refs/heads/{branch}` or `refs/pull/{pr_number}/head`
// The function returns the resolved ref, commit SHA and any error.
func resolveGitReference(ctx context.Context, githubClient *github.Client, owner, repo, ref, sha string) (*raw.ContentOpts, error) {
	// 1. If SHA is provided, use it directly
	if sha != "" {
		return &raw.ContentOpts{Ref: "", SHA: sha}, nil
	}

	// 2. If neither provided, use the default branch as ref
	if ref == "" {
		repoInfo, resp, err := githubClient.Repositories.Get(ctx, owner, repo)
		if err != nil {
			_, _ = ghErrors.NewGitHubAPIErrorToCtx(ctx, "failed to get repository info", resp, err)
			return nil, fmt.Errorf("failed to get repository info: %w", err)
		}
		ref = fmt.Sprintf("refs/heads/%s", repoInfo.GetDefaultBranch())
	}

	// 3. Get the SHA from the ref
	reference, resp, err := githubClient.Git.GetRef(ctx, owner, repo, ref)
	if err != nil {
		_, _ = ghErrors.NewGitHubAPIErrorToCtx(ctx, "failed to get reference", resp, err)
		return nil, fmt.Errorf("failed to get reference: %w", err)
	}
	sha = reference.GetObject().GetSHA()

	// Use provided ref, or it will be empty which defaults to the default branch
	return &raw.ContentOpts{Ref: ref, SHA: sha}, nil
}

```

#### x/search.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// SearchRepositories creates a tool to search for GitHub repositories.
func SearchRepositories(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_repositories",
			mcp.WithDescription(t("TOOL_SEARCH_REPOSITORIES_DESCRIPTION", "Search for GitHub repositories")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SEARCH_REPOSITORIES_USER_TITLE", "Search repositories"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("Search query"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			query, err := RequiredParam[string](request, "query")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.SearchOptions{
				ListOptions: github.ListOptions{
					Page:    pagination.Page,
					PerPage: pagination.PerPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			result, resp, err := client.Search.Repositories(ctx, query, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to search repositories with query '%s'", query),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to search repositories: %s", string(body))), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// SearchCode creates a tool to search for code across GitHub repositories.
func SearchCode(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_code",
			mcp.WithDescription(t("TOOL_SEARCH_CODE_DESCRIPTION", "Search for code across GitHub repositories")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SEARCH_CODE_USER_TITLE", "Search code"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("Search query using GitHub code search syntax"),
			),
			mcp.WithString("sort",
				mcp.Description("Sort field ('indexed' only)"),
			),
			mcp.WithString("order",
				mcp.Description("Sort order"),
				mcp.Enum("asc", "desc"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			query, err := RequiredParam[string](request, "query")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sort, err := OptionalParam[string](request, "sort")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			order, err := OptionalParam[string](request, "order")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.SearchOptions{
				Sort:  sort,
				Order: order,
				ListOptions: github.ListOptions{
					PerPage: pagination.PerPage,
					Page:    pagination.Page,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			result, resp, err := client.Search.Code(ctx, query, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to search code with query '%s'", query),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to search code: %s", string(body))), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// MinimalUser is the output type for user and organization search results.
type MinimalUser struct {
	Login      string       `json:"login"`
	ID         int64        `json:"id,omitempty"`
	ProfileURL string       `json:"profile_url,omitempty"`
	AvatarURL  string       `json:"avatar_url,omitempty"`
	Details    *UserDetails `json:"details,omitempty"` // Optional field for additional user details
}

type MinimalSearchUsersResult struct {
	TotalCount        int           `json:"total_count"`
	IncompleteResults bool          `json:"incomplete_results"`
	Items             []MinimalUser `json:"items"`
}

func userOrOrgHandler(accountType string, getClient GetClientFn) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query, err := RequiredParam[string](request, "query")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		sort, err := OptionalParam[string](request, "sort")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		order, err := OptionalParam[string](request, "order")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		pagination, err := OptionalPaginationParams(request)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		opts := &github.SearchOptions{
			Sort:  sort,
			Order: order,
			ListOptions: github.ListOptions{
				PerPage: pagination.PerPage,
				Page:    pagination.Page,
			},
		}

		client, err := getClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get GitHub client: %w", err)
		}

		searchQuery := "type:" + accountType + " " + query
		result, resp, err := client.Search.Users(ctx, searchQuery, opts)
		if err != nil {
			return ghErrors.NewGitHubAPIErrorResponse(ctx,
				fmt.Sprintf("failed to search %ss with query '%s'", accountType, query),
				resp,
				err,
			), nil
		}
		defer func() { _ = resp.Body.Close() }()

		if resp.StatusCode != 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to read response body: %w", err)
			}
			return mcp.NewToolResultError(fmt.Sprintf("failed to search %ss: %s", accountType, string(body))), nil
		}

		minimalUsers := make([]MinimalUser, 0, len(result.Users))

		for _, user := range result.Users {
			if user.Login != nil {
				mu := MinimalUser{
					Login:      user.GetLogin(),
					ID:         user.GetID(),
					ProfileURL: user.GetHTMLURL(),
					AvatarURL:  user.GetAvatarURL(),
				}
				minimalUsers = append(minimalUsers, mu)
			}
		}
		minimalResp := &MinimalSearchUsersResult{
			TotalCount:        result.GetTotal(),
			IncompleteResults: result.GetIncompleteResults(),
			Items:             minimalUsers,
		}
		if result.Total != nil {
			minimalResp.TotalCount = *result.Total
		}
		if result.IncompleteResults != nil {
			minimalResp.IncompleteResults = *result.IncompleteResults
		}

		r, err := json.Marshal(minimalResp)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return mcp.NewToolResultText(string(r)), nil
	}
}

// SearchUsers creates a tool to search for GitHub users.
func SearchUsers(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_users",
		mcp.WithDescription(t("TOOL_SEARCH_USERS_DESCRIPTION", "Search for GitHub users exclusively")),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        t("TOOL_SEARCH_USERS_USER_TITLE", "Search users"),
			ReadOnlyHint: ToBoolPtr(true),
		}),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search query using GitHub users search syntax scoped to type:user"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort field by category"),
			mcp.Enum("followers", "repositories", "joined"),
		),
		mcp.WithString("order",
			mcp.Description("Sort order"),
			mcp.Enum("asc", "desc"),
		),
		WithPagination(),
	), userOrOrgHandler("user", getClient)
}

// SearchOrgs creates a tool to search for GitHub organizations.
func SearchOrgs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_orgs",
		mcp.WithDescription(t("TOOL_SEARCH_ORGS_DESCRIPTION", "Search for GitHub organizations exclusively")),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        t("TOOL_SEARCH_ORGS_USER_TITLE", "Search organizations"),
			ReadOnlyHint: ToBoolPtr(true),
		}),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search query using GitHub organizations search syntax scoped to type:org"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort field by category"),
			mcp.Enum("followers", "repositories", "joined"),
		),
		mcp.WithString("order",
			mcp.Description("Sort order"),
			mcp.Enum("asc", "desc"),
		),
		WithPagination(),
	), userOrOrgHandler("org", getClient)
}

```

#### x/search_test.go

```go
package github

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/github/github-mcp-server/internal/toolsnaps"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SearchRepositories(t *testing.T) {
	// Verify tool definition once
	mockClient := github.NewClient(nil)
	tool, _ := SearchRepositories(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "search_repositories", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "query")
	assert.Contains(t, tool.InputSchema.Properties, "page")
	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"query"})

	// Setup mock search results
	mockSearchResult := &github.RepositoriesSearchResult{
		Total:             github.Ptr(2),
		IncompleteResults: github.Ptr(false),
		Repositories: []*github.Repository{
			{
				ID:              github.Ptr(int64(12345)),
				Name:            github.Ptr("repo-1"),
				FullName:        github.Ptr("owner/repo-1"),
				HTMLURL:         github.Ptr("https://github.com/owner/repo-1"),
				Description:     github.Ptr("Test repository 1"),
				StargazersCount: github.Ptr(100),
			},
			{
				ID:              github.Ptr(int64(67890)),
				Name:            github.Ptr("repo-2"),
				FullName:        github.Ptr("owner/repo-2"),
				HTMLURL:         github.Ptr("https://github.com/owner/repo-2"),
				Description:     github.Ptr("Test repository 2"),
				StargazersCount: github.Ptr(50),
			},
		},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult *github.RepositoriesSearchResult
		expectedErrMsg string
	}{
		{
			name: "successful repository search",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchRepositories,
					expectQueryParams(t, map[string]string{
						"q":        "golang test",
						"page":     "2",
						"per_page": "10",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query":   "golang test",
				"page":    float64(2),
				"perPage": float64(10),
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "repository search with default pagination",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchRepositories,
					expectQueryParams(t, map[string]string{
						"q":        "golang test",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "golang test",
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "search fails",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchRepositories,
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message": "Invalid query"}`))
					}),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "invalid:query",
			},
			expectError:    true,
			expectedErrMsg: "failed to search repositories",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup client with mock
			client := github.NewClient(tc.mockedClient)
			_, handler := SearchRepositories(stubGetClientFn(client), translations.NullTranslationHelper)

			// Create call request
			request := createMCPRequest(tc.requestArgs)

			// Call handler
			result, err := handler(context.Background(), request)

			// Verify results
			if tc.expectError {
				require.NoError(t, err)
				require.True(t, result.IsError)
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.False(t, result.IsError)

			// Parse the result and get the text content if no error
			textContent := getTextResult(t, result)

			// Unmarshal and verify the result
			var returnedResult github.RepositoriesSearchResult
			err = json.Unmarshal([]byte(textContent.Text), &returnedResult)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedResult.Total, *returnedResult.Total)
			assert.Equal(t, *tc.expectedResult.IncompleteResults, *returnedResult.IncompleteResults)
			assert.Len(t, returnedResult.Repositories, len(tc.expectedResult.Repositories))
			for i, repo := range returnedResult.Repositories {
				assert.Equal(t, *tc.expectedResult.Repositories[i].ID, *repo.ID)
				assert.Equal(t, *tc.expectedResult.Repositories[i].Name, *repo.Name)
				assert.Equal(t, *tc.expectedResult.Repositories[i].FullName, *repo.FullName)
				assert.Equal(t, *tc.expectedResult.Repositories[i].HTMLURL, *repo.HTMLURL)
			}

		})
	}
}

func Test_SearchCode(t *testing.T) {
	// Verify tool definition once
	mockClient := github.NewClient(nil)
	tool, _ := SearchCode(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "search_code", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "query")
	assert.Contains(t, tool.InputSchema.Properties, "sort")
	assert.Contains(t, tool.InputSchema.Properties, "order")
	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.Contains(t, tool.InputSchema.Properties, "page")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"query"})

	// Setup mock search results
	mockSearchResult := &github.CodeSearchResult{
		Total:             github.Ptr(2),
		IncompleteResults: github.Ptr(false),
		CodeResults: []*github.CodeResult{
			{
				Name:       github.Ptr("file1.go"),
				Path:       github.Ptr("path/to/file1.go"),
				SHA:        github.Ptr("abc123def456"),
				HTMLURL:    github.Ptr("https://github.com/owner/repo/blob/main/path/to/file1.go"),
				Repository: &github.Repository{Name: github.Ptr("repo"), FullName: github.Ptr("owner/repo")},
			},
			{
				Name:       github.Ptr("file2.go"),
				Path:       github.Ptr("path/to/file2.go"),
				SHA:        github.Ptr("def456abc123"),
				HTMLURL:    github.Ptr("https://github.com/owner/repo/blob/main/path/to/file2.go"),
				Repository: &github.Repository{Name: github.Ptr("repo"), FullName: github.Ptr("owner/repo")},
			},
		},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult *github.CodeSearchResult
		expectedErrMsg string
	}{
		{
			name: "successful code search with all parameters",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchCode,
					expectQueryParams(t, map[string]string{
						"q":        "fmt.Println language:go",
						"sort":     "indexed",
						"order":    "desc",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query":   "fmt.Println language:go",
				"sort":    "indexed",
				"order":   "desc",
				"page":    float64(1),
				"perPage": float64(30),
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "code search with minimal parameters",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchCode,
					expectQueryParams(t, map[string]string{
						"q":        "fmt.Println language:go",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "fmt.Println language:go",
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "search code fails",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchCode,
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message": "Validation Failed"}`))
					}),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "invalid:query",
			},
			expectError:    true,
			expectedErrMsg: "failed to search code",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup client with mock
			client := github.NewClient(tc.mockedClient)
			_, handler := SearchCode(stubGetClientFn(client), translations.NullTranslationHelper)

			// Create call request
			request := createMCPRequest(tc.requestArgs)

			// Call handler
			result, err := handler(context.Background(), request)

			// Verify results
			if tc.expectError {
				require.NoError(t, err)
				require.True(t, result.IsError)
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.False(t, result.IsError)

			// Parse the result and get the text content if no error
			textContent := getTextResult(t, result)

			// Unmarshal and verify the result
			var returnedResult github.CodeSearchResult
			err = json.Unmarshal([]byte(textContent.Text), &returnedResult)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedResult.Total, *returnedResult.Total)
			assert.Equal(t, *tc.expectedResult.IncompleteResults, *returnedResult.IncompleteResults)
			assert.Len(t, returnedResult.CodeResults, len(tc.expectedResult.CodeResults))
			for i, code := range returnedResult.CodeResults {
				assert.Equal(t, *tc.expectedResult.CodeResults[i].Name, *code.Name)
				assert.Equal(t, *tc.expectedResult.CodeResults[i].Path, *code.Path)
				assert.Equal(t, *tc.expectedResult.CodeResults[i].SHA, *code.SHA)
				assert.Equal(t, *tc.expectedResult.CodeResults[i].HTMLURL, *code.HTMLURL)
				assert.Equal(t, *tc.expectedResult.CodeResults[i].Repository.FullName, *code.Repository.FullName)
			}
		})
	}
}

func Test_SearchUsers(t *testing.T) {
	// Verify tool definition once
	mockClient := github.NewClient(nil)
	tool, _ := SearchUsers(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "search_users", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "query")
	assert.Contains(t, tool.InputSchema.Properties, "sort")
	assert.Contains(t, tool.InputSchema.Properties, "order")
	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.Contains(t, tool.InputSchema.Properties, "page")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"query"})

	// Setup mock search results
	mockSearchResult := &github.UsersSearchResult{
		Total:             github.Ptr(2),
		IncompleteResults: github.Ptr(false),
		Users: []*github.User{
			{
				Login:     github.Ptr("user1"),
				ID:        github.Ptr(int64(1001)),
				HTMLURL:   github.Ptr("https://github.com/user1"),
				AvatarURL: github.Ptr("https://avatars.githubusercontent.com/u/1001"),
			},
			{
				Login:     github.Ptr("user2"),
				ID:        github.Ptr(int64(1002)),
				HTMLURL:   github.Ptr("https://github.com/user2"),
				AvatarURL: github.Ptr("https://avatars.githubusercontent.com/u/1002"),
				Type:      github.Ptr("User"),
			},
		},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult *github.UsersSearchResult
		expectedErrMsg string
	}{
		{
			name: "successful users search with all parameters",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					expectQueryParams(t, map[string]string{
						"q":        "type:user location:finland language:go",
						"sort":     "followers",
						"order":    "desc",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query":   "location:finland language:go",
				"sort":    "followers",
				"order":   "desc",
				"page":    float64(1),
				"perPage": float64(30),
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "users search with minimal parameters",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					expectQueryParams(t, map[string]string{
						"q":        "type:user location:finland language:go",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "location:finland language:go",
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "search users fails",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message": "Validation Failed"}`))
					}),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "invalid:query",
			},
			expectError:    true,
			expectedErrMsg: "failed to search users",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup client with mock
			client := github.NewClient(tc.mockedClient)
			_, handler := SearchUsers(stubGetClientFn(client), translations.NullTranslationHelper)

			// Create call request
			request := createMCPRequest(tc.requestArgs)

			// Call handler
			result, err := handler(context.Background(), request)

			// Verify results
			if tc.expectError {
				require.NoError(t, err)
				require.True(t, result.IsError)
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.False(t, result.IsError)

			// Parse the result and get the text content if no error
			require.NotNil(t, result)

			textContent := getTextResult(t, result)

			// Unmarshal and verify the result
			var returnedResult MinimalSearchUsersResult
			err = json.Unmarshal([]byte(textContent.Text), &returnedResult)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedResult.Total, returnedResult.TotalCount)
			assert.Equal(t, *tc.expectedResult.IncompleteResults, returnedResult.IncompleteResults)
			assert.Len(t, returnedResult.Items, len(tc.expectedResult.Users))
			for i, user := range returnedResult.Items {
				assert.Equal(t, *tc.expectedResult.Users[i].Login, user.Login)
				assert.Equal(t, *tc.expectedResult.Users[i].ID, user.ID)
				assert.Equal(t, *tc.expectedResult.Users[i].HTMLURL, user.ProfileURL)
				assert.Equal(t, *tc.expectedResult.Users[i].AvatarURL, user.AvatarURL)
			}
		})
	}
}

func Test_SearchOrgs(t *testing.T) {
	// Verify tool definition once
	mockClient := github.NewClient(nil)
	tool, _ := SearchOrgs(stubGetClientFn(mockClient), translations.NullTranslationHelper)

	assert.Equal(t, "search_orgs", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "query")
	assert.Contains(t, tool.InputSchema.Properties, "sort")
	assert.Contains(t, tool.InputSchema.Properties, "order")
	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.Contains(t, tool.InputSchema.Properties, "page")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"query"})

	// Setup mock search results
	mockSearchResult := &github.UsersSearchResult{
		Total:             github.Ptr(int(2)),
		IncompleteResults: github.Ptr(false),
		Users: []*github.User{
			{
				Login:     github.Ptr("org-1"),
				ID:        github.Ptr(int64(111)),
				HTMLURL:   github.Ptr("https://github.com/org-1"),
				AvatarURL: github.Ptr("https://avatars.githubusercontent.com/u/111?v=4"),
			},
			{
				Login:     github.Ptr("org-2"),
				ID:        github.Ptr(int64(222)),
				HTMLURL:   github.Ptr("https://github.com/org-2"),
				AvatarURL: github.Ptr("https://avatars.githubusercontent.com/u/222?v=4"),
			},
		},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult *github.UsersSearchResult
		expectedErrMsg string
	}{
		{
			name: "successful org search",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					expectQueryParams(t, map[string]string{
						"q":        "type:org github",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "github",
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "org search fails",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message": "Validation Failed"}`))
					}),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "invalid:query",
			},
			expectError:    true,
			expectedErrMsg: "failed to search orgs",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup client with mock
			client := github.NewClient(tc.mockedClient)
			_, handler := SearchOrgs(stubGetClientFn(client), translations.NullTranslationHelper)

			// Create call request
			request := createMCPRequest(tc.requestArgs)

			// Call handler
			result, err := handler(context.Background(), request)

			// Verify results
			if tc.expectError {
				require.NoError(t, err)
				require.True(t, result.IsError)
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)

			textContent := getTextResult(t, result)

			// Unmarshal and verify the result
			var returnedResult MinimalSearchUsersResult
			err = json.Unmarshal([]byte(textContent.Text), &returnedResult)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedResult.Total, returnedResult.TotalCount)
			assert.Equal(t, *tc.expectedResult.IncompleteResults, returnedResult.IncompleteResults)
			assert.Len(t, returnedResult.Items, len(tc.expectedResult.Users))
			for i, org := range returnedResult.Items {
				assert.Equal(t, *tc.expectedResult.Users[i].Login, org.Login)
				assert.Equal(t, *tc.expectedResult.Users[i].ID, org.ID)
				assert.Equal(t, *tc.expectedResult.Users[i].HTMLURL, org.ProfileURL)
				assert.Equal(t, *tc.expectedResult.Users[i].AvatarURL, org.AvatarURL)
			}
		})
	}
}

```

#### x/search_utils.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
)

func searchHandler(
	ctx context.Context,
	getClient GetClientFn,
	request mcp.CallToolRequest,
	searchType string,
	errorPrefix string,
) (*mcp.CallToolResult, error) {
	query, err := RequiredParam[string](request, "query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	query = fmt.Sprintf("is:%s %s", searchType, query)

	owner, err := OptionalParam[string](request, "owner")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	repo, err := OptionalParam[string](request, "repo")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if owner != "" && repo != "" {
		query = fmt.Sprintf("repo:%s/%s %s", owner, repo, query)
	}

	sort, err := OptionalParam[string](request, "sort")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	order, err := OptionalParam[string](request, "order")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	pagination, err := OptionalPaginationParams(request)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	opts := &github.SearchOptions{
		// Default to "created" if no sort is provided, as it's a common use case.
		Sort:  sort,
		Order: order,
		ListOptions: github.ListOptions{
			Page:    pagination.Page,
			PerPage: pagination.PerPage,
		},
	}

	client, err := getClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get GitHub client: %w", errorPrefix, err)
	}
	result, resp, err := client.Search.Issues(ctx, query, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errorPrefix, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to read response body: %w", errorPrefix, err)
		}
		return mcp.NewToolResultError(fmt.Sprintf("%s: %s", errorPrefix, string(body))), nil
	}

	r, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to marshal response: %w", errorPrefix, err)
	}

	return mcp.NewToolResultText(string(r)), nil
}

```

#### x/server.go

```go
package github

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NewServer creates a new GitHub MCP server with the specified GH client and logger.

func NewServer(version string, opts ...server.ServerOption) *server.MCPServer {
	// Add default options
	defaultOpts := []server.ServerOption{
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	}
	opts = append(defaultOpts, opts...)

	// Create a new MCP server
	s := server.NewMCPServer(
		"github-mcp-server",
		version,
		opts...,
	)
	return s
}

// OptionalParamOK is a helper function that can be used to fetch a requested parameter from the request.
// It returns the value, a boolean indicating if the parameter was present, and an error if the type is wrong.
func OptionalParamOK[T any](r mcp.CallToolRequest, p string) (value T, ok bool, err error) {
	// Check if the parameter is present in the request
	val, exists := r.GetArguments()[p]
	if !exists {
		// Not present, return zero value, false, no error
		return
	}

	// Check if the parameter is of the expected type
	value, ok = val.(T)
	if !ok {
		// Present but wrong type
		err = fmt.Errorf("parameter %s is not of type %T, is %T", p, value, val)
		ok = true // Set ok to true because the parameter *was* present, even if wrong type
		return
	}

	// Present and correct type
	ok = true
	return
}

// isAcceptedError checks if the error is an accepted error.
func isAcceptedError(err error) bool {
	var acceptedError *github.AcceptedError
	return errors.As(err, &acceptedError)
}

// RequiredParam is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request.
// 2. Checks if the parameter is of the expected type.
// 3. Checks if the parameter is not empty, i.e: non-zero value
func RequiredParam[T comparable](r mcp.CallToolRequest, p string) (T, error) {
	var zero T

	// Check if the parameter is present in the request
	if _, ok := r.GetArguments()[p]; !ok {
		return zero, fmt.Errorf("missing required parameter: %s", p)
	}

	// Check if the parameter is of the expected type
	val, ok := r.GetArguments()[p].(T)
	if !ok {
		return zero, fmt.Errorf("parameter %s is not of type %T", p, zero)
	}

	if val == zero {
		return zero, fmt.Errorf("missing required parameter: %s", p)
	}

	return val, nil
}

// RequiredInt is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request.
// 2. Checks if the parameter is of the expected type.
// 3. Checks if the parameter is not empty, i.e: non-zero value
func RequiredInt(r mcp.CallToolRequest, p string) (int, error) {
	v, err := RequiredParam[float64](r, p)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

// OptionalParam is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request, if not, it returns its zero-value
// 2. If it is present, it checks if the parameter is of the expected type and returns it
func OptionalParam[T any](r mcp.CallToolRequest, p string) (T, error) {
	var zero T

	// Check if the parameter is present in the request
	if _, ok := r.GetArguments()[p]; !ok {
		return zero, nil
	}

	// Check if the parameter is of the expected type
	if _, ok := r.GetArguments()[p].(T); !ok {
		return zero, fmt.Errorf("parameter %s is not of type %T, is %T", p, zero, r.GetArguments()[p])
	}

	return r.GetArguments()[p].(T), nil
}

// OptionalIntParam is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request, if not, it returns its zero-value
// 2. If it is present, it checks if the parameter is of the expected type and returns it
func OptionalIntParam(r mcp.CallToolRequest, p string) (int, error) {
	v, err := OptionalParam[float64](r, p)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

// OptionalIntParamWithDefault is a helper function that can be used to fetch a requested parameter from the request
// similar to optionalIntParam, but it also takes a default value.
func OptionalIntParamWithDefault(r mcp.CallToolRequest, p string, d int) (int, error) {
	v, err := OptionalIntParam(r, p)
	if err != nil {
		return 0, err
	}
	if v == 0 {
		return d, nil
	}
	return v, nil
}

// OptionalStringArrayParam is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request, if not, it returns its zero-value
// 2. If it is present, iterates the elements and checks each is a string
func OptionalStringArrayParam(r mcp.CallToolRequest, p string) ([]string, error) {
	// Check if the parameter is present in the request
	if _, ok := r.GetArguments()[p]; !ok {
		return []string{}, nil
	}

	switch v := r.GetArguments()[p].(type) {
	case nil:
		return []string{}, nil
	case []string:
		return v, nil
	case []any:
		strSlice := make([]string, len(v))
		for i, v := range v {
			s, ok := v.(string)
			if !ok {
				return []string{}, fmt.Errorf("parameter %s is not of type string, is %T", p, v)
			}
			strSlice[i] = s
		}
		return strSlice, nil
	default:
		return []string{}, fmt.Errorf("parameter %s could not be coerced to []string, is %T", p, r.GetArguments()[p])
	}
}

func MarshalledTextResult(v any) *mcp.CallToolResult {
	data, err := json.Marshal(v)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("failed to marshal text result to json", err)
	}

	return mcp.NewToolResultText(string(data))
}

```

#### x/server_test.go

```go
package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/github/github-mcp-server/pkg/raw"
	"github.com/google/go-github/v73/github"
	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
)

func stubGetClientFn(client *github.Client) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return client, nil
	}
}

func stubGetClientFromHTTPFn(client *http.Client) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return github.NewClient(client), nil
	}
}

func stubGetClientFnErr(err string) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return nil, errors.New(err)
	}
}

func stubGetGQLClientFn(client *githubv4.Client) GetGQLClientFn {
	return func(_ context.Context) (*githubv4.Client, error) {
		return client, nil
	}
}

func stubGetRawClientFn(client *raw.Client) raw.GetRawClientFn {
	return func(_ context.Context) (*raw.Client, error) {
		return client, nil
	}
}

func badRequestHandler(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		structuredErrorResponse := github.ErrorResponse{
			Message: msg,
		}

		b, err := json.Marshal(structuredErrorResponse)
		if err != nil {
			http.Error(w, "failed to marshal error response", http.StatusInternalServerError)
		}

		http.Error(w, string(b), http.StatusBadRequest)
	}
}

func Test_IsAcceptedError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectAccepted bool
	}{
		{
			name:           "github AcceptedError",
			err:            &github.AcceptedError{},
			expectAccepted: true,
		},
		{
			name:           "regular error",
			err:            fmt.Errorf("some other error"),
			expectAccepted: false,
		},
		{
			name:           "nil error",
			err:            nil,
			expectAccepted: false,
		},
		{
			name:           "wrapped AcceptedError",
			err:            fmt.Errorf("wrapped: %w", &github.AcceptedError{}),
			expectAccepted: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isAcceptedError(tc.err)
			assert.Equal(t, tc.expectAccepted, result)
		})
	}
}

func Test_RequiredStringParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    string
		expectError bool
	}{
		{
			name:        "valid string parameter",
			params:      map[string]interface{}{"name": "test-value"},
			paramName:   "name",
			expected:    "test-value",
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "name",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty string parameter",
			params:      map[string]interface{}{"name": ""},
			paramName:   "name",
			expected:    "",
			expectError: true,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"name": 123},
			paramName:   "name",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := RequiredParam[string](request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_OptionalStringParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    string
		expectError bool
	}{
		{
			name:        "valid string parameter",
			params:      map[string]interface{}{"name": "test-value"},
			paramName:   "name",
			expected:    "test-value",
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "name",
			expected:    "",
			expectError: false,
		},
		{
			name:        "empty string parameter",
			params:      map[string]interface{}{"name": ""},
			paramName:   "name",
			expected:    "",
			expectError: false,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"name": 123},
			paramName:   "name",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalParam[string](request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_RequiredInt(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    int
		expectError bool
	}{
		{
			name:        "valid number parameter",
			params:      map[string]interface{}{"count": float64(42)},
			paramName:   "count",
			expected:    42,
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "count",
			expected:    0,
			expectError: true,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"count": "not-a-number"},
			paramName:   "count",
			expected:    0,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := RequiredInt(request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
func Test_OptionalIntParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    int
		expectError bool
	}{
		{
			name:        "valid number parameter",
			params:      map[string]interface{}{"count": float64(42)},
			paramName:   "count",
			expected:    42,
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "count",
			expected:    0,
			expectError: false,
		},
		{
			name:        "zero value",
			params:      map[string]interface{}{"count": float64(0)},
			paramName:   "count",
			expected:    0,
			expectError: false,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"count": "not-a-number"},
			paramName:   "count",
			expected:    0,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalIntParam(request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_OptionalNumberParamWithDefault(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		defaultVal  int
		expected    int
		expectError bool
	}{
		{
			name:        "valid number parameter",
			params:      map[string]interface{}{"count": float64(42)},
			paramName:   "count",
			defaultVal:  10,
			expected:    42,
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "count",
			defaultVal:  10,
			expected:    10,
			expectError: false,
		},
		{
			name:        "zero value",
			params:      map[string]interface{}{"count": float64(0)},
			paramName:   "count",
			defaultVal:  10,
			expected:    10,
			expectError: false,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"count": "not-a-number"},
			paramName:   "count",
			defaultVal:  10,
			expected:    0,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalIntParamWithDefault(request, tc.paramName, tc.defaultVal)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_OptionalBooleanParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    bool
		expectError bool
	}{
		{
			name:        "true value",
			params:      map[string]interface{}{"flag": true},
			paramName:   "flag",
			expected:    true,
			expectError: false,
		},
		{
			name:        "false value",
			params:      map[string]interface{}{"flag": false},
			paramName:   "flag",
			expected:    false,
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "flag",
			expected:    false,
			expectError: false,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"flag": "not-a-boolean"},
			paramName:   "flag",
			expected:    false,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalParam[bool](request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestOptionalStringArrayParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    []string
		expectError bool
	}{
		{
			name:        "parameter not in request",
			params:      map[string]any{},
			paramName:   "flag",
			expected:    []string{},
			expectError: false,
		},
		{
			name: "valid any array parameter",
			params: map[string]any{
				"flag": []any{"v1", "v2"},
			},
			paramName:   "flag",
			expected:    []string{"v1", "v2"},
			expectError: false,
		},
		{
			name: "valid string array parameter",
			params: map[string]any{
				"flag": []string{"v1", "v2"},
			},
			paramName:   "flag",
			expected:    []string{"v1", "v2"},
			expectError: false,
		},
		{
			name: "wrong type parameter",
			params: map[string]any{
				"flag": 1,
			},
			paramName:   "flag",
			expected:    []string{},
			expectError: true,
		},
		{
			name: "wrong slice type parameter",
			params: map[string]any{
				"flag": []any{"foo", 2},
			},
			paramName:   "flag",
			expected:    []string{},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalStringArrayParam(request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

```

#### x/sub_issues.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ListSubIssues creates a tool to list the sub-issues of an issue.
func ListSubIssues(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_sub_issues",
			mcp.WithDescription(t("TOOL_LIST_SUB_ISSUES_DESCRIPTION", "List sub-issues for a specific issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_SUB_ISSUES_USER_TITLE", "List sub-issues"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("Issue number"),
			),
			mcp.WithNumber("page",
				mcp.Description("Page number for pagination (default: 1)"),
			),
			mcp.WithNumber("per_page",
				mcp.Description("Number of results per page (max 100, default: 30)"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, repo, issueNumber, err := requiredSubIssueParent(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			page, err := OptionalIntParamWithDefault(request, "page", 1)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			perPage, err := OptionalIntParamWithDefault(request, "per_page", 30)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			opts := &github.IssueListOptions{
				ListOptions: github.ListOptions{Page: page, PerPage: perPage},
			}
			subIssues, resp, err := client.SubIssue.ListByIssue(ctx, owner, repo, int64(issueNumber), opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to list sub-issues", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			return subIssueResult(resp, http.StatusOK, "failed to list sub-issues", subIssues)
		}
}

// AddSubIssue creates a tool to add a sub-issue to a parent issue.
func AddSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("add_sub_issue",
			mcp.WithDescription(t("TOOL_ADD_SUB_ISSUE_DESCRIPTION", "Add a sub-issue to a parent issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_ADD_SUB_ISSUE_USER_TITLE", "Add sub-issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("The number of the parent issue"),
			),
			mcp.WithNumber("sub_issue_id",
				mcp.Required(),
				mcp.Description("The ID of the sub-issue to add. ID is not the same as issue number"),
			),
			mcp.WithBoolean("replace_parent",
				mcp.Description("When true, replaces the sub-issue's current parent issue"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, repo, issueNumber, err := requiredSubIssueParent(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := RequiredInt(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			replaceParent, err := OptionalParam[bool](request, "replace_parent")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			subIssueRequest := github.SubIssueRequest{SubIssueID: int64(subIssueID)}
			if replaceParent {
				subIssueRequest.ReplaceParent = github.Ptr(true)
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			subIssue, resp, err := client.SubIssue.Add(ctx, owner, repo, int64(issueNumber), subIssueRequest)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to add sub-issue", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			return subIssueResult(resp, http.StatusCreated, "failed to add sub-issue", subIssue)
		}
}

// RemoveSubIssue creates a tool to remove a sub-issue from a parent issue.
func RemoveSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("remove_sub_issue",
			mcp.WithDescription(t("TOOL_REMOVE_SUB_ISSUE_DESCRIPTION", "Remove a sub-issue from a parent issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_REMOVE_SUB_ISSUE_USER_TITLE", "Remove sub-issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("The number of the parent issue"),
			),
			mcp.WithNumber("sub_issue_id",
				mcp.Required(),
				mcp.Description("The ID of the sub-issue to remove. ID is not the same as issue number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, repo, issueNumber, err := requiredSubIssueParent(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := RequiredInt(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// client.SubIssue.Remove sends the request to ".../sub_issues" instead of
			// ".../sub_issue" (google/go-github#3613), so the request is built here.
			u := fmt.Sprintf("repos/%v/%v/issues/%v/sub_issue", owner, repo, issueNumber)
			req, err := client.NewRequest(http.MethodDelete, u, github.SubIssueRequest{SubIssueID: int64(subIssueID)})
			if err != nil {
				return nil, fmt.Errorf("failed to create request: %w", err)
			}
			resp, err := client.BareDo(ctx, req)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to remove sub-issue", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to read response body: %w", err)
			}
			if resp.StatusCode != http.StatusOK {
				return mcp.NewToolResultError(fmt.Sprintf("failed to remove sub-issue: %s", string(body))), nil
			}

			var parent github.Issue
			if err := json.Unmarshal(body, &parent); err != nil {
				return nil, fmt.Errorf("failed to unmarshal response: %w", err)
			}
			r, err := json.Marshal(&parent)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

// ReprioritizeSubIssue creates a tool to move a sub-issue to another position in its parent's list.
func ReprioritizeSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("reprioritize_sub_issue",
			mcp.WithDescription(t("TOOL_REPRIORITIZE_SUB_ISSUE_DESCRIPTION", "Reprioritize a sub-issue to a different position in the parent issue's sub-issue list.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_REPRIORITIZE_SUB_ISSUE_USER_TITLE", "Reprioritize sub-issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("The number of the parent issue"),
			),
			mcp.WithNumber("sub_issue_id",
				mcp.Required(),
				mcp.Description("The ID of the sub-issue to reprioritize. ID is not the same as issue number"),
			),
			mcp.WithNumber("after_id",
				mcp.Description("The ID of the sub-issue to be prioritized after (either after_id OR before_id should be specified)"),
			),
			mcp.WithNumber("before_id",
				mcp.Description("The ID of the sub-issue to be prioritized before (either after_id OR before_id should be specified)"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, repo, issueNumber, err := requiredSubIssueParent(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := RequiredInt(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			afterID, err := OptionalIntParam(request, "after_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			beforeID, err := OptionalIntParam(request, "before_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			subIssueRequest := github.SubIssueRequest{SubIssueID: int64(subIssueID)}
			switch {
			case afterID == 0 && beforeID == 0:
				return mcp.NewToolResultError("either after_id or before_id must be specified"), nil
			case afterID != 0 && beforeID != 0:
				return mcp.NewToolResultError("only one of after_id or before_id should be specified, not both"), nil
			case afterID != 0:
				subIssueRequest.AfterID = github.Ptr(int64(afterID))
			default:
				subIssueRequest.BeforeID = github.Ptr(int64(beforeID))
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			subIssue, resp, err := client.SubIssue.Reprioritize(ctx, owner, repo, int64(issueNumber), subIssueRequest)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to reprioritize sub-issue", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			return subIssueResult(resp, http.StatusOK, "failed to reprioritize sub-issue", subIssue)
		}
}

// requiredSubIssueParent returns the "owner", "repo" and "issue_number" parameters,
// which name the parent issue in every sub-issue tool.
func requiredSubIssueParent(request mcp.CallToolRequest) (owner, repo string, issueNumber int, err error) {
	owner, err = RequiredParam[string](request, "owner")
	if err != nil {
		return "", "", 0, err
	}
	repo, err = RequiredParam[string](request, "repo")
	if err != nil {
		return "", "", 0, err
	}
	issueNumber, err = RequiredInt(request, "issue_number")
	if err != nil {
		return "", "", 0, err
	}
	return owner, repo, issueNumber, nil
}

// subIssueResult returns v as JSON text when the response has the status the endpoint
// documents for success. Otherwise it returns a tool error that starts with failMessage.
// go-github has already read and closed the body, so the error gives the status only.
func subIssueResult(resp *github.Response, wantStatus int, failMessage string, v any) (*mcp.CallToolResult, error) {
	if resp.StatusCode != wantStatus {
		return mcp.NewToolResultError(fmt.Sprintf("%s: unexpected status %d", failMessage, resp.StatusCode)), nil
	}

	r, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(r)), nil
}

```

#### x/sub_issues_test.go

```go
package github

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/github/github-mcp-server/internal/toolsnaps"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// subIssueTestCase is one call of a sub-issue tool against a mocked API.
type subIssueTestCase struct {
	name         string
	mockedClient *http.Client
	requestArgs  map[string]any
	// expectedErrMsg is the start of the tool error. It is empty when the call succeeds.
	expectedErrMsg string
	// expectedIssues are the issues the result holds when the call succeeds.
	// A tool that returns one issue is expected to return the first.
	expectedIssues []*github.Issue
}

// newSubIssueTool is the signature that every sub-issue tool constructor has.
type newSubIssueTool func(GetClientFn, translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc)

// runSubIssueTestCases calls the tool's handler for each case and checks its result.
// listResult says whether the tool returns a list of issues instead of one.
func runSubIssueTestCases(t *testing.T, newTool newSubIssueTool, listResult bool, tests []subIssueTestCase) {
	t.Helper()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := github.NewClient(tc.mockedClient)
			_, handler := newTool(stubGetClientFn(client), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.requestArgs))
			require.NoError(t, err)

			if tc.expectedErrMsg != "" {
				errorContent := getErrorResult(t, result)
				assert.True(t, strings.HasPrefix(errorContent.Text, tc.expectedErrMsg), "error %q does not start with %q", errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.False(t, result.IsError, "unexpected tool error: %v", result.Content)
			textContent := getTextResult(t, result)
			var returned []*github.Issue
			if listResult {
				require.NoError(t, json.Unmarshal([]byte(textContent.Text), &returned))
			} else {
				var issue github.Issue
				require.NoError(t, json.Unmarshal([]byte(textContent.Text), &issue))
				returned = []*github.Issue{&issue}
			}
			require.Len(t, returned, len(tc.expectedIssues))
			for i, want := range tc.expectedIssues {
				assert.Equal(t, want.GetID(), returned[i].GetID())
				assert.Equal(t, want.GetNumber(), returned[i].GetNumber())
				assert.Equal(t, want.GetTitle(), returned[i].GetTitle())
			}
		})
	}
}

// subIssueParentArgs returns the arguments that name the parent issue, with extra added.
func subIssueParentArgs(extra map[string]any) map[string]any {
	args := map[string]any{
		"owner":        "owner",
		"repo":         "repo",
		"issue_number": float64(42),
	}
	for k, v := range extra {
		args[k] = v
	}
	return args
}

// withoutSubIssueArg returns args without the named argument.
func withoutSubIssueArg(args map[string]any, name string) map[string]any {
	delete(args, name)
	return args
}

var subIssueMock = &github.Issue{
	ID:     github.Ptr(int64(123)),
	Number: github.Ptr(43),
	Title:  github.Ptr("Sub-issue"),
}

var subIssueParentMock = &github.Issue{
	ID:     github.Ptr(int64(100)),
	Number: github.Ptr(42),
	Title:  github.Ptr("Parent issue"),
}

func Test_AddSubIssue(t *testing.T) {
	tool, _ := AddSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))
	assert.Equal(t, "add_sub_issue", tool.Name)
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo", "issue_number", "sub_issue_id"})

	runSubIssueTestCases(t, AddSubIssue, false, []subIssueTestCase{
		{
			name: "adds the sub-issue",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expect(t, expectations{
						path:        "/repos/owner/repo/issues/42/sub_issues",
						requestBody: map[string]any{"sub_issue_id": float64(123)},
					}).andThen(mockResponse(t, http.StatusCreated, subIssueMock)),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123)}),
			expectedIssues: []*github.Issue{subIssueMock},
		},
		{
			name: "sends replace_parent",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectRequestBody(t, map[string]any{"sub_issue_id": float64(123), "replace_parent": true}).
						andThen(mockResponse(t, http.StatusCreated, subIssueMock)),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123), "replace_parent": true}),
			expectedIssues: []*github.Issue{subIssueMock},
		},
		{
			name:           "missing sub_issue_id",
			mockedClient:   mock.NewMockedHTTPClient(),
			requestArgs:    subIssueParentArgs(nil),
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
		{
			name:           "missing owner",
			mockedClient:   mock.NewMockedHTTPClient(),
			requestArgs:    withoutSubIssueArg(subIssueParentArgs(map[string]any{"sub_issue_id": float64(123)}), "owner"),
			expectedErrMsg: "missing required parameter: owner",
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusUnprocessableEntity, `{"message": "Validation Failed"}`),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123)}),
			expectedErrMsg: "failed to add sub-issue",
		},
		{
			name: "unexpected status",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusOK, subIssueMock),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123)}),
			expectedErrMsg: "failed to add sub-issue",
		},
	})
}

func Test_ListSubIssues(t *testing.T) {
	tool, _ := ListSubIssues(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))
	assert.Equal(t, "list_sub_issues", tool.Name)
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo", "issue_number"})

	otherSubIssue := &github.Issue{ID: github.Ptr(int64(124)), Number: github.Ptr(44), Title: github.Ptr("Another")}

	runSubIssueTestCases(t, ListSubIssues, true, []subIssueTestCase{
		{
			name: "lists with the default page",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expect(t, expectations{
						path:        "/repos/owner/repo/issues/42/sub_issues",
						queryParams: map[string]string{"page": "1", "per_page": "30"},
					}).andThen(mockResponse(t, http.StatusOK, []*github.Issue{subIssueMock, otherSubIssue})),
				),
			),
			requestArgs:    subIssueParentArgs(nil),
			expectedIssues: []*github.Issue{subIssueMock, otherSubIssue},
		},
		{
			name: "lists the requested page",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectQueryParams(t, map[string]string{"page": "2", "per_page": "10"}).
						andThen(mockResponse(t, http.StatusOK, []*github.Issue{otherSubIssue})),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"page": float64(2), "per_page": float64(10)}),
			expectedIssues: []*github.Issue{otherSubIssue},
		},
		{
			name:           "missing issue_number",
			mockedClient:   mock.NewMockedHTTPClient(),
			requestArgs:    withoutSubIssueArg(subIssueParentArgs(nil), "issue_number"),
			expectedErrMsg: "missing required parameter: issue_number",
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusNotFound, `{"message": "Not Found"}`),
				),
			),
			requestArgs:    subIssueParentArgs(nil),
			expectedErrMsg: "failed to list sub-issues",
		},
		{
			name: "unexpected status",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusNonAuthoritativeInfo, []*github.Issue{}),
				),
			),
			requestArgs:    subIssueParentArgs(nil),
			expectedErrMsg: "failed to list sub-issues",
		},
	})
}

func Test_RemoveSubIssue(t *testing.T) {
	tool, _ := RemoveSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))
	assert.Equal(t, "remove_sub_issue", tool.Name)
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo", "issue_number", "sub_issue_id"})

	runSubIssueTestCases(t, RemoveSubIssue, false, []subIssueTestCase{
		{
			name: "removes the sub-issue",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					expect(t, expectations{
						path:        "/repos/owner/repo/issues/42/sub_issue",
						requestBody: map[string]any{"sub_issue_id": float64(123)},
					}).andThen(mockResponse(t, http.StatusOK, subIssueParentMock)),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123)}),
			expectedIssues: []*github.Issue{subIssueParentMock},
		},
		{
			name:           "missing sub_issue_id",
			mockedClient:   mock.NewMockedHTTPClient(),
			requestArgs:    subIssueParentArgs(nil),
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusNotFound, `{"message": "Not Found"}`),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123)}),
			expectedErrMsg: "failed to remove sub-issue",
		},
		{
			name: "unexpected status",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusNoContent, ""),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123)}),
			expectedErrMsg: "failed to remove sub-issue",
		},
	})
}

func Test_ReprioritizeSubIssue(t *testing.T) {
	tool, _ := ReprioritizeSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))
	assert.Equal(t, "reprioritize_sub_issue", tool.Name)
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo", "issue_number", "sub_issue_id"})

	runSubIssueTestCases(t, ReprioritizeSubIssue, false, []subIssueTestCase{
		{
			name: "places the sub-issue after another",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					expect(t, expectations{
						path:        "/repos/owner/repo/issues/42/sub_issues/priority",
						requestBody: map[string]any{"sub_issue_id": float64(123), "after_id": float64(456)},
					}).andThen(mockResponse(t, http.StatusOK, subIssueMock)),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123), "after_id": float64(456)}),
			expectedIssues: []*github.Issue{subIssueMock},
		},
		{
			name: "places the sub-issue before another",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					expectRequestBody(t, map[string]any{"sub_issue_id": float64(123), "before_id": float64(789)}).
						andThen(mockResponse(t, http.StatusOK, subIssueMock)),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123), "before_id": float64(789)}),
			expectedIssues: []*github.Issue{subIssueMock},
		},
		{
			name:           "neither after_id nor before_id",
			mockedClient:   mock.NewMockedHTTPClient(),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123)}),
			expectedErrMsg: "either after_id or before_id must be specified",
		},
		{
			name:           "both after_id and before_id",
			mockedClient:   mock.NewMockedHTTPClient(),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123), "after_id": float64(456), "before_id": float64(789)}),
			expectedErrMsg: "only one of after_id or before_id should be specified, not both",
		},
		{
			name:           "missing sub_issue_id",
			mockedClient:   mock.NewMockedHTTPClient(),
			requestArgs:    subIssueParentArgs(map[string]any{"after_id": float64(456)}),
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusUnprocessableEntity, `{"message": "Validation Failed"}`),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123), "after_id": float64(456)}),
			expectedErrMsg: "failed to reprioritize sub-issue",
		},
		{
			name: "unexpected status",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusNonAuthoritativeInfo, subIssueMock),
				),
			),
			requestArgs:    subIssueParentArgs(map[string]any{"sub_issue_id": float64(123), "after_id": float64(456)}),
			expectedErrMsg: "failed to reprioritize sub-issue",
		},
	})
}

```

#### x/tools.go

```go
package github

import (
	"context"

	"github.com/github/github-mcp-server/pkg/raw"
	"github.com/github/github-mcp-server/pkg/toolsets"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

type GetClientFn func(context.Context) (*github.Client, error)
type GetGQLClientFn func(context.Context) (*githubv4.Client, error)

var DefaultTools = []string{"all"}

func DefaultToolsetGroup(readOnly bool, getClient GetClientFn, getGQLClient GetGQLClientFn, getRawClient raw.GetRawClientFn, t translations.TranslationHelperFunc) *toolsets.ToolsetGroup {
	tsg := toolsets.NewToolsetGroup(readOnly)

	// Define all available features with their default state (disabled)
	// Create toolsets
	repos := toolsets.NewToolset("repos", "GitHub Repository related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchRepositories(getClient, t)),
			toolsets.NewServerTool(GetFileContents(getClient, getRawClient, t)),
			toolsets.NewServerTool(ListCommits(getClient, t)),
			toolsets.NewServerTool(SearchCode(getClient, t)),
			toolsets.NewServerTool(GetCommit(getClient, t)),
			toolsets.NewServerTool(ListBranches(getClient, t)),
			toolsets.NewServerTool(ListTags(getClient, t)),
			toolsets.NewServerTool(GetTag(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(CreateOrUpdateFile(getClient, t)),
			toolsets.NewServerTool(CreateRepository(getClient, t)),
			toolsets.NewServerTool(ForkRepository(getClient, t)),
			toolsets.NewServerTool(CreateBranch(getClient, t)),
			toolsets.NewServerTool(PushFiles(getClient, t)),
			toolsets.NewServerTool(DeleteFile(getClient, t)),
		).
		AddResourceTemplates(
			toolsets.NewServerResourceTemplate(GetRepositoryResourceContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceBranchContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceCommitContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceTagContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourcePrContent(getClient, getRawClient, t)),
		)
	issues := toolsets.NewToolset("issues", "GitHub Issues related tools").
		AddReadTools(
			toolsets.NewServerTool(GetIssue(getClient, t)),
			toolsets.NewServerTool(SearchIssues(getClient, t)),
			toolsets.NewServerTool(ListIssues(getClient, t)),
			toolsets.NewServerTool(GetIssueComments(getClient, t)),
			toolsets.NewServerTool(ListSubIssues(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(CreateIssue(getClient, t)),
			toolsets.NewServerTool(AddIssueComment(getClient, t)),
			toolsets.NewServerTool(UpdateIssue(getClient, t)),
			toolsets.NewServerTool(AssignCopilotToIssue(getGQLClient, t)),
			toolsets.NewServerTool(AddSubIssue(getClient, t)),
			toolsets.NewServerTool(RemoveSubIssue(getClient, t)),
			toolsets.NewServerTool(ReprioritizeSubIssue(getClient, t)),
		).AddPrompts(toolsets.NewServerPrompt(AssignCodingAgentPrompt(t)))
	users := toolsets.NewToolset("users", "GitHub User related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchUsers(getClient, t)),
		)
	orgs := toolsets.NewToolset("orgs", "GitHub Organization related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchOrgs(getClient, t)),
		)
	pullRequests := toolsets.NewToolset("pull_requests", "GitHub Pull Request related tools").
		AddReadTools(
			toolsets.NewServerTool(GetPullRequest(getClient, t)),
			toolsets.NewServerTool(ListPullRequests(getClient, t)),
			toolsets.NewServerTool(GetPullRequestFiles(getClient, t)),
			toolsets.NewServerTool(SearchPullRequests(getClient, t)),
			toolsets.NewServerTool(GetPullRequestStatus(getClient, t)),
			toolsets.NewServerTool(GetPullRequestComments(getClient, t)),
			toolsets.NewServerTool(GetPullRequestReviews(getClient, t)),
			toolsets.NewServerTool(GetPullRequestDiff(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(MergePullRequest(getClient, t)),
			toolsets.NewServerTool(UpdatePullRequestBranch(getClient, t)),
			toolsets.NewServerTool(CreatePullRequest(getClient, t)),
			toolsets.NewServerTool(UpdatePullRequest(getClient, t)),
			toolsets.NewServerTool(RequestCopilotReview(getClient, t)),

			// Reviews
			toolsets.NewServerTool(CreateAndSubmitPullRequestReview(getGQLClient, t)),
			toolsets.NewServerTool(CreatePendingPullRequestReview(getGQLClient, t)),
			toolsets.NewServerTool(AddCommentToPendingReview(getGQLClient, t)),
			toolsets.NewServerTool(SubmitPendingPullRequestReview(getGQLClient, t)),
			toolsets.NewServerTool(DeletePendingPullRequestReview(getGQLClient, t)),
		)
	codeSecurity := toolsets.NewToolset("code_security", "Code security related tools, such as GitHub Code Scanning").
		AddReadTools(
			toolsets.NewServerTool(GetCodeScanningAlert(getClient, t)),
			toolsets.NewServerTool(ListCodeScanningAlerts(getClient, t)),
		)
	secretProtection := toolsets.NewToolset("secret_protection", "Secret protection related tools, such as GitHub Secret Scanning").
		AddReadTools(
			toolsets.NewServerTool(GetSecretScanningAlert(getClient, t)),
			toolsets.NewServerTool(ListSecretScanningAlerts(getClient, t)),
		)
	dependabot := toolsets.NewToolset("dependabot", "Dependabot tools").
		AddReadTools(
			toolsets.NewServerTool(GetDependabotAlert(getClient, t)),
			toolsets.NewServerTool(ListDependabotAlerts(getClient, t)),
		)

	notifications := toolsets.NewToolset("notifications", "GitHub Notifications related tools").
		AddReadTools(
			toolsets.NewServerTool(ListNotifications(getClient, t)),
			toolsets.NewServerTool(GetNotificationDetails(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(DismissNotification(getClient, t)),
			toolsets.NewServerTool(MarkAllNotificationsRead(getClient, t)),
			toolsets.NewServerTool(ManageNotificationSubscription(getClient, t)),
			toolsets.NewServerTool(ManageRepositoryNotificationSubscription(getClient, t)),
		)

	discussions := toolsets.NewToolset("discussions", "GitHub Discussions related tools").
		AddReadTools(
			toolsets.NewServerTool(ListDiscussions(getGQLClient, t)),
			toolsets.NewServerTool(GetDiscussion(getGQLClient, t)),
			toolsets.NewServerTool(GetDiscussionComments(getGQLClient, t)),
			toolsets.NewServerTool(ListDiscussionCategories(getGQLClient, t)),
		)

	actions := toolsets.NewToolset("actions", "GitHub Actions workflows and CI/CD operations").
		AddReadTools(
			toolsets.NewServerTool(ListWorkflows(getClient, t)),
			toolsets.NewServerTool(ListWorkflowRuns(getClient, t)),
			toolsets.NewServerTool(GetWorkflowRun(getClient, t)),
			toolsets.NewServerTool(GetWorkflowRunLogs(getClient, t)),
			toolsets.NewServerTool(ListWorkflowJobs(getClient, t)),
			toolsets.NewServerTool(GetJobLogs(getClient, t)),
			toolsets.NewServerTool(ListWorkflowRunArtifacts(getClient, t)),
			toolsets.NewServerTool(DownloadWorkflowRunArtifact(getClient, t)),
			toolsets.NewServerTool(GetWorkflowRunUsage(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(RunWorkflow(getClient, t)),
			toolsets.NewServerTool(RerunWorkflowRun(getClient, t)),
			toolsets.NewServerTool(RerunFailedJobs(getClient, t)),
			toolsets.NewServerTool(CancelWorkflowRun(getClient, t)),
			toolsets.NewServerTool(DeleteWorkflowRunLogs(getClient, t)),
		)

	// Keep experiments alive so the system doesn't error out when it's always enabled
	experiments := toolsets.NewToolset("experiments", "Experimental features that are not considered stable yet")

	contextTools := toolsets.NewToolset("context", "Tools that provide context about the current user and GitHub context you are operating in").
		AddReadTools(
			toolsets.NewServerTool(GetMe(getClient, t)),
		)

	// Add toolsets to the group
	tsg.AddToolset(contextTools)
	tsg.AddToolset(repos)
	tsg.AddToolset(issues)
	tsg.AddToolset(orgs)
	tsg.AddToolset(users)
	tsg.AddToolset(pullRequests)
	tsg.AddToolset(actions)
	tsg.AddToolset(codeSecurity)
	tsg.AddToolset(secretProtection)
	tsg.AddToolset(dependabot)
	tsg.AddToolset(notifications)
	tsg.AddToolset(experiments)
	tsg.AddToolset(discussions)

	return tsg
}

// InitDynamicToolset creates a dynamic toolset that can be used to enable other toolsets, and so requires the server and toolset group as arguments
func InitDynamicToolset(s *server.MCPServer, tsg *toolsets.ToolsetGroup, t translations.TranslationHelperFunc) *toolsets.Toolset {
	// Create a new dynamic toolset
	// Need to add the dynamic toolset last so it can be used to enable other toolsets
	dynamicToolSelection := toolsets.NewToolset("dynamic", "Discover GitHub MCP tools that can help achieve tasks by enabling additional sets of tools, you can control the enablement of any toolset to access its tools when this toolset is enabled.").
		AddReadTools(
			toolsets.NewServerTool(ListAvailableToolsets(tsg, t)),
			toolsets.NewServerTool(GetToolsetsTools(tsg, t)),
			toolsets.NewServerTool(EnableToolset(s, tsg, t)),
		)

	dynamicToolSelection.Enabled = true
	return dynamicToolSelection
}

// ToBoolPtr converts a bool to a *bool pointer.
func ToBoolPtr(b bool) *bool {
	return &b
}

// ToStringPtr converts a string to a *string pointer.
// Returns nil if the string is empty.
func ToStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

```

#### y/actions.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	DescriptionRepositoryOwner = "Repository owner"
	DescriptionRepositoryName  = "Repository name"
)

// ListWorkflows creates a tool to list workflows in a repository
func ListWorkflows(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_workflows",
			mcp.WithDescription(t("TOOL_LIST_WORKFLOWS_DESCRIPTION", "List workflows in a repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_WORKFLOWS_USER_TITLE", "List workflows"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional pagination parameters
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Set up list options
			opts := &github.ListOptions{
				PerPage: pagination.PerPage,
				Page:    pagination.Page,
			}

			workflows, resp, err := client.Actions.ListWorkflows(ctx, owner, repo, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to list workflows: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(workflows)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListWorkflowRuns creates a tool to list workflow runs for a specific workflow
func ListWorkflowRuns(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_workflow_runs",
			mcp.WithDescription(t("TOOL_LIST_WORKFLOW_RUNS_DESCRIPTION", "List workflow runs for a specific workflow")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_WORKFLOW_RUNS_USER_TITLE", "List workflow runs"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithString("workflow_id",
				mcp.Required(),
				mcp.Description("The workflow ID or workflow file name"),
			),
			mcp.WithString("actor",
				mcp.Description("Returns someone's workflow runs. Use the login for the user who created the workflow run."),
			),
			mcp.WithString("branch",
				mcp.Description("Returns workflow runs associated with a branch. Use the name of the branch."),
			),
			mcp.WithString("event",
				mcp.Description("Returns workflow runs for a specific event type"),
				mcp.Enum(
					"branch_protection_rule",
					"check_run",
					"check_suite",
					"create",
					"delete",
					"deployment",
					"deployment_status",
					"discussion",
					"discussion_comment",
					"fork",
					"gollum",
					"issue_comment",
					"issues",
					"label",
					"merge_group",
					"milestone",
					"page_build",
					"public",
					"pull_request",
					"pull_request_review",
					"pull_request_review_comment",
					"pull_request_target",
					"push",
					"registry_package",
					"release",
					"repository_dispatch",
					"schedule",
					"status",
					"watch",
					"workflow_call",
					"workflow_dispatch",
					"workflow_run",
				),
			),
			mcp.WithString("status",
				mcp.Description("Returns workflow runs with the check run status"),
				mcp.Enum("queued", "in_progress", "completed", "requested", "waiting"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			workflowID, err := RequiredParam[string](request, "workflow_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional filtering parameters
			actor, err := OptionalParam[string](request, "actor")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := OptionalParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			event, err := OptionalParam[string](request, "event")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			status, err := OptionalParam[string](request, "status")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional pagination parameters
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Set up list options
			opts := &github.ListWorkflowRunsOptions{
				Actor:  actor,
				Branch: branch,
				Event:  event,
				Status: status,
				ListOptions: github.ListOptions{
					PerPage: pagination.PerPage,
					Page:    pagination.Page,
				},
			}

			workflowRuns, resp, err := client.Actions.ListWorkflowRunsByFileName(ctx, owner, repo, workflowID, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to list workflow runs: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(workflowRuns)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// RunWorkflow creates a tool to run an Actions workflow
func RunWorkflow(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("run_workflow",
			mcp.WithDescription(t("TOOL_RUN_WORKFLOW_DESCRIPTION", "Run an Actions workflow by workflow ID or filename")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_RUN_WORKFLOW_USER_TITLE", "Run workflow"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithString("workflow_id",
				mcp.Required(),
				mcp.Description("The workflow ID (numeric) or workflow file name (e.g., main.yml, ci.yaml)"),
			),
			mcp.WithString("ref",
				mcp.Required(),
				mcp.Description("The git reference for the workflow. The reference can be a branch or tag name."),
			),
			mcp.WithObject("inputs",
				mcp.Description("Inputs the workflow accepts"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			workflowID, err := RequiredParam[string](request, "workflow_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			ref, err := RequiredParam[string](request, "ref")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional inputs parameter
			var inputs map[string]interface{}
			if requestInputs, ok := request.GetArguments()["inputs"]; ok {
				if inputsMap, ok := requestInputs.(map[string]interface{}); ok {
					inputs = inputsMap
				}
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			event := github.CreateWorkflowDispatchEventRequest{
				Ref:    ref,
				Inputs: inputs,
			}

			var resp *github.Response
			var workflowType string

			if workflowIDInt, parseErr := strconv.ParseInt(workflowID, 10, 64); parseErr == nil {
				resp, err = client.Actions.CreateWorkflowDispatchEventByID(ctx, owner, repo, workflowIDInt, event)
				workflowType = "workflow_id"
			} else {
				resp, err = client.Actions.CreateWorkflowDispatchEventByFileName(ctx, owner, repo, workflowID, event)
				workflowType = "workflow_file"
			}

			if err != nil {
				return nil, fmt.Errorf("failed to run workflow: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":       "Workflow run has been queued",
				"workflow_type": workflowType,
				"workflow_id":   workflowID,
				"ref":           ref,
				"inputs":        inputs,
				"status":        resp.Status,
				"status_code":   resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetWorkflowRun creates a tool to get details of a specific workflow run
func GetWorkflowRun(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_workflow_run",
			mcp.WithDescription(t("TOOL_GET_WORKFLOW_RUN_DESCRIPTION", "Get details of a specific workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_WORKFLOW_RUN_USER_TITLE", "Get workflow run"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			workflowRun, resp, err := client.Actions.GetWorkflowRunByID(ctx, owner, repo, runID)
			if err != nil {
				return nil, fmt.Errorf("failed to get workflow run: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(workflowRun)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetWorkflowRunLogs creates a tool to download logs for a specific workflow run
func GetWorkflowRunLogs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_workflow_run_logs",
			mcp.WithDescription(t("TOOL_GET_WORKFLOW_RUN_LOGS_DESCRIPTION", "Download logs for a specific workflow run (EXPENSIVE: downloads ALL logs as ZIP. Consider using get_job_logs with failed_only=true for debugging failed jobs)")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_WORKFLOW_RUN_LOGS_USER_TITLE", "Get workflow run logs"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the download URL for the logs
			url, resp, err := client.Actions.GetWorkflowRunLogs(ctx, owner, repo, runID, 1)
			if err != nil {
				return nil, fmt.Errorf("failed to get workflow run logs: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			// Create response with the logs URL and information
			result := map[string]any{
				"logs_url":         url.String(),
				"message":          "Workflow run logs are available for download",
				"note":             "The logs_url provides a download link for the complete workflow run logs as a ZIP archive. You can download this archive to extract and examine individual job logs.",
				"warning":          "This downloads ALL logs as a ZIP file which can be large and expensive. For debugging failed jobs, consider using get_job_logs with failed_only=true and run_id instead.",
				"optimization_tip": "Use: get_job_logs with parameters {run_id: " + fmt.Sprintf("%d", runID) + ", failed_only: true} for more efficient failed job debugging",
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListWorkflowJobs creates a tool to list jobs for a specific workflow run
func ListWorkflowJobs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_workflow_jobs",
			mcp.WithDescription(t("TOOL_LIST_WORKFLOW_JOBS_DESCRIPTION", "List jobs for a specific workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_WORKFLOW_JOBS_USER_TITLE", "List workflow jobs"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
			mcp.WithString("filter",
				mcp.Description("Filters jobs by their completed_at timestamp"),
				mcp.Enum("latest", "all"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			// Get optional filtering parameters
			filter, err := OptionalParam[string](request, "filter")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional pagination parameters
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Set up list options
			opts := &github.ListWorkflowJobsOptions{
				Filter: filter,
				ListOptions: github.ListOptions{
					PerPage: pagination.PerPage,
					Page:    pagination.Page,
				},
			}

			jobs, resp, err := client.Actions.ListWorkflowJobs(ctx, owner, repo, runID, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to list workflow jobs: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			// Add optimization tip for failed job debugging
			response := map[string]any{
				"jobs":             jobs,
				"optimization_tip": "For debugging failed jobs, consider using get_job_logs with failed_only=true and run_id=" + fmt.Sprintf("%d", runID) + " to get logs directly without needing to list jobs first",
			}

			r, err := json.Marshal(response)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetJobLogs creates a tool to download logs for a specific workflow job or efficiently get all failed job logs for a workflow run
func GetJobLogs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_job_logs",
			mcp.WithDescription(t("TOOL_GET_JOB_LOGS_DESCRIPTION", "Download logs for a specific workflow job or efficiently get all failed job logs for a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_JOB_LOGS_USER_TITLE", "Get job logs"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("job_id",
				mcp.Description("The unique identifier of the workflow job (required for single job logs)"),
			),
			mcp.WithNumber("run_id",
				mcp.Description("Workflow run ID (required when using failed_only)"),
			),
			mcp.WithBoolean("failed_only",
				mcp.Description("When true, gets logs for all failed jobs in run_id"),
			),
			mcp.WithBoolean("return_content",
				mcp.Description("Returns actual log content instead of URLs"),
			),
			mcp.WithNumber("tail_lines",
				mcp.Description("Number of lines to return from the end of the log"),
				mcp.DefaultNumber(500),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional parameters
			jobID, err := OptionalIntParam(request, "job_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID, err := OptionalIntParam(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			failedOnly, err := OptionalParam[bool](request, "failed_only")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			returnContent, err := OptionalParam[bool](request, "return_content")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			tailLines, err := OptionalIntParam(request, "tail_lines")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			// Default to 500 lines if not specified
			if tailLines == 0 {
				tailLines = 500
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Validate parameters
			if failedOnly && runID == 0 {
				return mcp.NewToolResultError("run_id is required when failed_only is true"), nil
			}
			if !failedOnly && jobID == 0 {
				return mcp.NewToolResultError("job_id is required when failed_only is false"), nil
			}

			if failedOnly && runID > 0 {
				// Handle failed-only mode: get logs for all failed jobs in the workflow run
				return handleFailedJobLogs(ctx, client, owner, repo, int64(runID), returnContent, tailLines)
			} else if jobID > 0 {
				// Handle single job mode
				return handleSingleJobLogs(ctx, client, owner, repo, int64(jobID), returnContent, tailLines)
			}

			return mcp.NewToolResultError("Either job_id must be provided for single job logs, or run_id with failed_only=true for failed job logs"), nil
		}
}

// handleFailedJobLogs gets logs for all failed jobs in a workflow run
func handleFailedJobLogs(ctx context.Context, client *github.Client, owner, repo string, runID int64, returnContent bool, tailLines int) (*mcp.CallToolResult, error) {
	// First, get all jobs for the workflow run
	jobs, resp, err := client.Actions.ListWorkflowJobs(ctx, owner, repo, runID, &github.ListWorkflowJobsOptions{
		Filter: "latest",
	})
	if err != nil {
		return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to list workflow jobs", resp, err), nil
	}
	defer func() { _ = resp.Body.Close() }()

	// Filter for failed jobs
	var failedJobs []*github.WorkflowJob
	for _, job := range jobs.Jobs {
		if job.GetConclusion() == "failure" {
			failedJobs = append(failedJobs, job)
		}
	}

	if len(failedJobs) == 0 {
		result := map[string]any{
			"message":     "No failed jobs found in this workflow run",
			"run_id":      runID,
			"total_jobs":  len(jobs.Jobs),
			"failed_jobs": 0,
		}
		r, _ := json.Marshal(result)
		return mcp.NewToolResultText(string(r)), nil
	}

	// Collect logs for all failed jobs
	var logResults []map[string]any
	for _, job := range failedJobs {
		jobResult, resp, err := getJobLogData(ctx, client, owner, repo, job.GetID(), job.GetName(), returnContent, tailLines)
		if err != nil {
			// Continue with other jobs even if one fails
			jobResult = map[string]any{
				"job_id":   job.GetID(),
				"job_name": job.GetName(),
				"error":    err.Error(),
			}
			// Enable reporting of status codes and error causes
			_, _ = ghErrors.NewGitHubAPIErrorToCtx(ctx, "failed to get job logs", resp, err) // Explicitly ignore error for graceful handling
		}

		logResults = append(logResults, jobResult)
	}

	result := map[string]any{
		"message":       fmt.Sprintf("Retrieved logs for %d failed jobs", len(failedJobs)),
		"run_id":        runID,
		"total_jobs":    len(jobs.Jobs),
		"failed_jobs":   len(failedJobs),
		"logs":          logResults,
		"return_format": map[string]bool{"content": returnContent, "urls": !returnContent},
	}

	r, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	return mcp.NewToolResultText(string(r)), nil
}

// handleSingleJobLogs gets logs for a single job
func handleSingleJobLogs(ctx context.Context, client *github.Client, owner, repo string, jobID int64, returnContent bool, tailLines int) (*mcp.CallToolResult, error) {
	jobResult, resp, err := getJobLogData(ctx, client, owner, repo, jobID, "", returnContent, tailLines)
	if err != nil {
		return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to get job logs", resp, err), nil
	}

	r, err := json.Marshal(jobResult)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	return mcp.NewToolResultText(string(r)), nil
}

// getJobLogData retrieves log data for a single job, either as URL or content
func getJobLogData(ctx context.Context, client *github.Client, owner, repo string, jobID int64, jobName string, returnContent bool, tailLines int) (map[string]any, *github.Response, error) {
	// Get the download URL for the job logs
	url, resp, err := client.Actions.GetWorkflowJobLogs(ctx, owner, repo, jobID, 1)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get job logs for job %d: %w", jobID, err)
	}
	defer func() { _ = resp.Body.Close() }()

	result := map[string]any{
		"job_id": jobID,
	}
	if jobName != "" {
		result["job_name"] = jobName
	}

	if returnContent {
		// Download and return the actual log content
		content, originalLength, httpResp, err := downloadLogContent(url.String(), tailLines) //nolint:bodyclose // Response body is closed in downloadLogContent, but we need to return httpResp
		if err != nil {
			// To keep the return value consistent wrap the response as a GitHub Response
			ghRes := &github.Response{
				Response: httpResp,
			}
			return nil, ghRes, fmt.Errorf("failed to download log content for job %d: %w", jobID, err)
		}
		result["logs_content"] = content
		result["message"] = "Job logs content retrieved successfully"
		result["original_length"] = originalLength
	} else {
		// Return just the URL
		result["logs_url"] = url.String()
		result["message"] = "Job logs are available for download"
		result["note"] = "The logs_url provides a download link for the individual job logs in plain text format. Use return_content=true to get the actual log content."
	}

	return result, resp, nil
}

// downloadLogContent downloads the actual log content from a GitHub logs URL
func downloadLogContent(logURL string, tailLines int) (string, int, *http.Response, error) {
	httpResp, err := http.Get(logURL) //nolint:gosec // URLs are provided by GitHub API and are safe
	if err != nil {
		return "", 0, httpResp, fmt.Errorf("failed to download logs: %w", err)
	}
	defer func() { _ = httpResp.Body.Close() }()

	if httpResp.StatusCode != http.StatusOK {
		return "", 0, httpResp, fmt.Errorf("failed to download logs: HTTP %d", httpResp.StatusCode)
	}

	content, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return "", 0, httpResp, fmt.Errorf("failed to read log content: %w", err)
	}

	// Clean up and format the log content for better readability
	logContent := strings.TrimSpace(string(content))

	trimmedContent, lineCount := trimContent(logContent, tailLines)
	return trimmedContent, lineCount, httpResp, nil
}

// trimContent trims the content to a maximum length and returns the trimmed content and an original length
func trimContent(content string, tailLines int) (string, int) {
	// Truncate to tail_lines if specified
	lineCount := 0
	if tailLines > 0 {

		// Count backwards to find the nth newline from the end and a total number of lines
		for i := len(content) - 1; i >= 0 && lineCount < tailLines; i-- {
			if content[i] == '\n' {
				lineCount++
				// If we have reached the tailLines, trim the content
				if lineCount == tailLines {
					content = content[i+1:]
				}
			}
		}
	}
	return content, lineCount
}

// RerunWorkflowRun creates a tool to re-run an entire workflow run
func RerunWorkflowRun(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("rerun_workflow_run",
			mcp.WithDescription(t("TOOL_RERUN_WORKFLOW_RUN_DESCRIPTION", "Re-run an entire workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_RERUN_WORKFLOW_RUN_USER_TITLE", "Rerun workflow run"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			resp, err := client.Actions.RerunWorkflowByID(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to rerun workflow run", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":     "Workflow run has been queued for re-run",
				"run_id":      runID,
				"status":      resp.Status,
				"status_code": resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// RerunFailedJobs creates a tool to re-run only the failed jobs in a workflow run
func RerunFailedJobs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("rerun_failed_jobs",
			mcp.WithDescription(t("TOOL_RERUN_FAILED_JOBS_DESCRIPTION", "Re-run only the failed jobs in a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_RERUN_FAILED_JOBS_USER_TITLE", "Rerun failed jobs"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			resp, err := client.Actions.RerunFailedJobsByID(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to rerun failed jobs", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":     "Failed jobs have been queued for re-run",
				"run_id":      runID,
				"status":      resp.Status,
				"status_code": resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CancelWorkflowRun creates a tool to cancel a workflow run
func CancelWorkflowRun(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("cancel_workflow_run",
			mcp.WithDescription(t("TOOL_CANCEL_WORKFLOW_RUN_DESCRIPTION", "Cancel a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CANCEL_WORKFLOW_RUN_USER_TITLE", "Cancel workflow run"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			resp, err := client.Actions.CancelWorkflowRunByID(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to cancel workflow run", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":     "Workflow run has been cancelled",
				"run_id":      runID,
				"status":      resp.Status,
				"status_code": resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListWorkflowRunArtifacts creates a tool to list artifacts for a workflow run
func ListWorkflowRunArtifacts(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_workflow_run_artifacts",
			mcp.WithDescription(t("TOOL_LIST_WORKFLOW_RUN_ARTIFACTS_DESCRIPTION", "List artifacts for a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_WORKFLOW_RUN_ARTIFACTS_USER_TITLE", "List workflow artifacts"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			// Get optional pagination parameters
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Set up list options
			opts := &github.ListOptions{
				PerPage: pagination.PerPage,
				Page:    pagination.Page,
			}

			artifacts, resp, err := client.Actions.ListWorkflowRunArtifacts(ctx, owner, repo, runID, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to list workflow run artifacts", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(artifacts)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// DownloadWorkflowRunArtifact creates a tool to download a workflow run artifact
func DownloadWorkflowRunArtifact(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("download_workflow_run_artifact",
			mcp.WithDescription(t("TOOL_DOWNLOAD_WORKFLOW_RUN_ARTIFACT_DESCRIPTION", "Get download URL for a workflow run artifact")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_DOWNLOAD_WORKFLOW_RUN_ARTIFACT_USER_TITLE", "Download workflow artifact"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("artifact_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the artifact"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			artifactIDInt, err := RequiredInt(request, "artifact_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			artifactID := int64(artifactIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the download URL for the artifact
			url, resp, err := client.Actions.DownloadArtifact(ctx, owner, repo, artifactID, 1)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to get artifact download URL", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Create response with the download URL and information
			result := map[string]any{
				"download_url": url.String(),
				"message":      "Artifact is available for download",
				"note":         "The download_url provides a download link for the artifact as a ZIP archive. The link is temporary and expires after a short time.",
				"artifact_id":  artifactID,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// DeleteWorkflowRunLogs creates a tool to delete logs for a workflow run
func DeleteWorkflowRunLogs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("delete_workflow_run_logs",
			mcp.WithDescription(t("TOOL_DELETE_WORKFLOW_RUN_LOGS_DESCRIPTION", "Delete logs for a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:           t("TOOL_DELETE_WORKFLOW_RUN_LOGS_USER_TITLE", "Delete workflow logs"),
				ReadOnlyHint:    ToBoolPtr(false),
				DestructiveHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			resp, err := client.Actions.DeleteWorkflowRunLogs(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to delete workflow run logs", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			result := map[string]any{
				"message":     "Workflow run logs have been deleted",
				"run_id":      runID,
				"status":      resp.Status,
				"status_code": resp.StatusCode,
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetWorkflowRunUsage creates a tool to get usage metrics for a workflow run
func GetWorkflowRunUsage(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_workflow_run_usage",
			mcp.WithDescription(t("TOOL_GET_WORKFLOW_RUN_USAGE_DESCRIPTION", "Get usage metrics for a workflow run")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_WORKFLOW_RUN_USAGE_USER_TITLE", "Get workflow usage"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryOwner),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description(DescriptionRepositoryName),
			),
			mcp.WithNumber("run_id",
				mcp.Required(),
				mcp.Description("The unique identifier of the workflow run"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runIDInt, err := RequiredInt(request, "run_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			runID := int64(runIDInt)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			usage, resp, err := client.Actions.GetWorkflowRunUsageByID(ctx, owner, repo, runID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to get workflow run usage", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(usage)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

```

#### y/discussions.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

func ListDiscussions(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_discussions",
			mcp.WithDescription(t("TOOL_LIST_DISCUSSIONS_DESCRIPTION", "List discussions for a repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_DISCUSSIONS_USER_TITLE", "List discussions"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("category",
				mcp.Description("Optional filter by discussion category ID. If provided, only discussions with this category are listed."),
			),
			mcp.WithString("orderBy",
				mcp.Description("Order discussions by field. If provided, the 'direction' also needs to be provided."),
				mcp.Enum("CREATED_AT", "UPDATED_AT"),
			),
			mcp.WithString("direction",
				mcp.Description("Order direction."),
				mcp.Enum("ASC", "DESC"),
			),
			WithCursorPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Required params
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Optional params
			category, err := OptionalParam[string](request, "category")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			orderBy, err := OptionalParam[string](request, "orderBy")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			direction, err := OptionalParam[string](request, "direction")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			pagination, err := OptionalCursorPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			paginationParams, err := pagination.ToGraphQLParams()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			vars := map[string]interface{}{
				"owner": githubv4.String(owner),
				"repo":  githubv4.String(repo),
				"first": githubv4.Int(*paginationParams.First),
				"after": (*githubv4.String)(paginationParams.After),
			}
			// If category filter is specified, use it as the category ID for server-side filtering
			filtered := category != ""
			if filtered {
				vars["categoryId"] = githubv4.ID(category)
			}
			// Ordering needs both the field and the direction; with only one
			// of them, the discussions are listed in the default order.
			ordered := orderBy != "" && direction != ""
			if ordered {
				vars["orderByField"] = githubv4.DiscussionOrderField(orderBy)
				vars["orderByDirection"] = githubv4.OrderDirection(direction)
			}

			// githubv4 builds the query from the struct type, so each
			// combination of arguments needs its own type.
			var conn discussionsConnection
			switch {
			case filtered && ordered:
				var query struct {
					Repository struct {
						Discussions discussionsConnection `graphql:"discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection })"`
					} `graphql:"repository(owner: $owner, name: $repo)"`
				}
				if err := client.Query(ctx, &query, vars); err != nil {
					return mcp.NewToolResultError(err.Error()), nil
				}
				conn = query.Repository.Discussions
			case filtered:
				var query struct {
					Repository struct {
						Discussions discussionsConnection `graphql:"discussions(first: $first, after: $after, categoryId: $categoryId)"`
					} `graphql:"repository(owner: $owner, name: $repo)"`
				}
				if err := client.Query(ctx, &query, vars); err != nil {
					return mcp.NewToolResultError(err.Error()), nil
				}
				conn = query.Repository.Discussions
			case ordered:
				var query struct {
					Repository struct {
						Discussions discussionsConnection `graphql:"discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection })"`
					} `graphql:"repository(owner: $owner, name: $repo)"`
				}
				if err := client.Query(ctx, &query, vars); err != nil {
					return mcp.NewToolResultError(err.Error()), nil
				}
				conn = query.Repository.Discussions
			default:
				var query struct {
					Repository struct {
						Discussions discussionsConnection `graphql:"discussions(first: $first, after: $after)"`
					} `graphql:"repository(owner: $owner, name: $repo)"`
				}
				if err := client.Query(ctx, &query, vars); err != nil {
					return mcp.NewToolResultError(err.Error()), nil
				}
				conn = query.Repository.Discussions
			}

			// Map nodes to GitHub Discussion objects
			var discussions []*github.Discussion
			for _, n := range conn.Nodes {
				di := &github.Discussion{
					Number:    github.Ptr(int(n.Number)),
					Title:     github.Ptr(string(n.Title)),
					HTMLURL:   github.Ptr(string(n.URL)),
					CreatedAt: &github.Timestamp{Time: n.CreatedAt.Time},
					UpdatedAt: &github.Timestamp{Time: n.UpdatedAt.Time},
					DiscussionCategory: &github.DiscussionCategory{
						Name: github.Ptr(string(n.Category.Name)),
					},
				}
				// The author is null when the account was deleted.
				if n.Author.Login != "" {
					di.User = &github.User{Login: github.Ptr(string(n.Author.Login))}
				}
				discussions = append(discussions, di)
			}

			// Create response with pagination info
			response := map[string]interface{}{
				"discussions": discussions,
				"pageInfo":    conn.PageInfo,
				"totalCount":  int(conn.TotalCount),
			}

			// Marshal and return
			out, err := json.Marshal(response)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussions: %w", err)
			}
			return mcp.NewToolResultText(string(out)), nil
		}
}

func GetDiscussion(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_discussion",
			mcp.WithDescription(t("TOOL_GET_DISCUSSION_DESCRIPTION", "Get a specific discussion by ID")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_DISCUSSION_USER_TITLE", "Get discussion"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("discussionNumber",
				mcp.Required(),
				mcp.Description("Discussion Number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner            string
				Repo             string
				DiscussionNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var q struct {
				Repository struct {
					Discussion struct {
						Number    githubv4.Int
						Body      githubv4.String
						CreatedAt githubv4.DateTime
						URL       githubv4.String `graphql:"url"`
						Category  struct {
							Name githubv4.String
						} `graphql:"category"`
					} `graphql:"discussion(number: $discussionNumber)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			vars := map[string]interface{}{
				"owner":            githubv4.String(params.Owner),
				"repo":             githubv4.String(params.Repo),
				"discussionNumber": githubv4.Int(params.DiscussionNumber),
			}
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			d := q.Repository.Discussion
			discussion := &github.Discussion{
				Number:    github.Ptr(int(d.Number)),
				Body:      github.Ptr(string(d.Body)),
				HTMLURL:   github.Ptr(string(d.URL)),
				CreatedAt: &github.Timestamp{Time: d.CreatedAt.Time},
				DiscussionCategory: &github.DiscussionCategory{
					Name: github.Ptr(string(d.Category.Name)),
				},
			}
			out, err := json.Marshal(discussion)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussion: %w", err)
			}

			return mcp.NewToolResultText(string(out)), nil
		}
}

func GetDiscussionComments(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_discussion_comments",
			mcp.WithDescription(t("TOOL_GET_DISCUSSION_COMMENTS_DESCRIPTION", "Get comments from a discussion")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_DISCUSSION_COMMENTS_USER_TITLE", "Get discussion comments"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner", mcp.Required(), mcp.Description("Repository owner")),
			mcp.WithString("repo", mcp.Required(), mcp.Description("Repository name")),
			mcp.WithNumber("discussionNumber", mcp.Required(), mcp.Description("Discussion Number")),
			WithCursorPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner            string
				Repo             string
				DiscussionNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			pagination, err := OptionalCursorPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			paginationParams, err := pagination.ToGraphQLParams()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var q struct {
				Repository struct {
					Discussion struct {
						Comments struct {
							Nodes []struct {
								Body githubv4.String
							}
							PageInfo   discussionsPageInfo
							TotalCount githubv4.Int
						} `graphql:"comments(first: $first, after: $after)"`
					} `graphql:"discussion(number: $discussionNumber)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			vars := map[string]interface{}{
				"owner":            githubv4.String(params.Owner),
				"repo":             githubv4.String(params.Repo),
				"discussionNumber": githubv4.Int(params.DiscussionNumber),
				"first":            githubv4.Int(*paginationParams.First),
				"after":            (*githubv4.String)(paginationParams.After),
			}
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			var comments []*github.IssueComment
			for _, c := range q.Repository.Discussion.Comments.Nodes {
				comments = append(comments, &github.IssueComment{Body: github.Ptr(string(c.Body))})
			}

			response := map[string]interface{}{
				"comments":   comments,
				"pageInfo":   q.Repository.Discussion.Comments.PageInfo,
				"totalCount": int(q.Repository.Discussion.Comments.TotalCount),
			}

			out, err := json.Marshal(response)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal comments: %w", err)
			}

			return mcp.NewToolResultText(string(out)), nil
		}
}

func ListDiscussionCategories(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_discussion_categories",
			mcp.WithDescription(t("TOOL_LIST_DISCUSSION_CATEGORIES_DESCRIPTION", "List discussion categories with their id and name, for a repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_DISCUSSION_CATEGORIES_USER_TITLE", "List discussion categories"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner string
				Repo  string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}
			var q struct {
				Repository struct {
					DiscussionCategories struct {
						Nodes []struct {
							ID   githubv4.ID
							Name githubv4.String
						}
						PageInfo   discussionsPageInfo
						TotalCount githubv4.Int
					} `graphql:"discussionCategories(first: $first)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			vars := map[string]interface{}{
				"owner": githubv4.String(params.Owner),
				"repo":  githubv4.String(params.Repo),
				"first": githubv4.Int(25),
			}
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			var categories []map[string]string
			for _, c := range q.Repository.DiscussionCategories.Nodes {
				categories = append(categories, map[string]string{
					"id":   fmt.Sprint(c.ID),
					"name": string(c.Name),
				})
			}
			response := map[string]interface{}{
				"categories": categories,
				"pageInfo":   q.Repository.DiscussionCategories.PageInfo,
				"totalCount": int(q.Repository.DiscussionCategories.TotalCount),
			}

			out, err := json.Marshal(response)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussion categories: %w", err)
			}
			return mcp.NewToolResultText(string(out)), nil
		}
}

// discussionsPageInfo is the pageInfo of a GraphQL connection, as the
// discussion tools query it and return it.
type discussionsPageInfo struct {
	HasNextPage     githubv4.Boolean `json:"hasNextPage"`
	HasPreviousPage githubv4.Boolean `json:"hasPreviousPage"`
	StartCursor     githubv4.String  `json:"startCursor"`
	EndCursor       githubv4.String  `json:"endCursor"`
}

// discussionsConnection is the discussions connection that list_discussions
// queries, with the fields it returns for each discussion.
type discussionsConnection struct {
	Nodes []struct {
		Number    githubv4.Int
		Title     githubv4.String
		CreatedAt githubv4.DateTime
		UpdatedAt githubv4.DateTime
		Author    struct {
			Login githubv4.String
		}
		Category struct {
			Name githubv4.String
		} `graphql:"category"`
		URL githubv4.String `graphql:"url"`
	}
	PageInfo   discussionsPageInfo
	TotalCount githubv4.Int
}

```

#### y/discussions_test.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/github/github-mcp-server/internal/githubv4mock"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	discussionsGeneral = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-02T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-03-02T00:00:00Z", "author": map[string]any{"login": "user3"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	discussionsAll = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-02T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 2, "title": "Discussion 2 title", "createdAt": "2023-02-01T00:00:00Z", "updatedAt": "2023-02-02T00:00:00Z", "author": map[string]any{"login": "user2"}, "url": "https://github.com/owner/repo/discussions/2", "category": map[string]any{"name": "Questions"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-03-02T00:00:00Z", "author": map[string]any{"login": "user3"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	mockResponseListAll = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes": discussionsAll,
				"pageInfo": map[string]any{
					"hasNextPage":     false,
					"hasPreviousPage": false,
					"startCursor":     "",
					"endCursor":       "",
				},
				"totalCount": 3,
			},
		},
	})
	mockResponseListGeneral = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes": discussionsGeneral,
				"pageInfo": map[string]any{
					"hasNextPage":     false,
					"hasPreviousPage": false,
					"startCursor":     "",
					"endCursor":       "",
				},
				"totalCount": 2,
			},
		},
	})
	mockErrorRepoNotFound = githubv4mock.ErrorResponse("repository not found")
)

func Test_ListDiscussions(t *testing.T) {
	mockClient := githubv4.NewClient(nil)
	// Verify tool definition and schema
	toolDef, _ := ListDiscussions(stubGetGQLClientFn(mockClient), translations.NullTranslationHelper)
	assert.Equal(t, "list_discussions", toolDef.Name)
	assert.NotEmpty(t, toolDef.Description)
	assert.Contains(t, toolDef.InputSchema.Properties, "owner")
	assert.Contains(t, toolDef.InputSchema.Properties, "repo")
	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo"})

	// The GraphQL documents list_discussions sends, one for each combination
	// of the category filter and the ordering.
	qDiscussions := "query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsFiltered := "query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsOrdered := "query($after:String$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsFilteredOrdered := "query($after:String$categoryId:ID!$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"

	varsListAll := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(30),
		"after": (*githubv4.String)(nil),
	}

	varsRepoNotFound := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("nonexistent-repo"),
		"first": githubv4.Int(30),
		"after": (*githubv4.String)(nil),
	}

	varsDiscussionsFiltered := map[string]interface{}{
		"owner":      githubv4.String("owner"),
		"repo":       githubv4.String("repo"),
		"categoryId": githubv4.ID("DIC_kwDOABC123"),
		"first":      githubv4.Int(30),
		"after":      (*githubv4.String)(nil),
	}

	varsListWithPerPage := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(10),
		"after": (*githubv4.String)(nil),
	}

	varsOrderByCreatedAtDesc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"first":            githubv4.Int(30),
		"after":            (*githubv4.String)(nil),
		"orderByField":     githubv4.DiscussionOrderField("CREATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
	}

	varsOrderByUpdatedAtAsc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"first":            githubv4.Int(30),
		"after":            (*githubv4.String)(nil),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
	}

	varsFilteredOrdered := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"categoryId":       githubv4.ID("DIC_kwDOABC123"),
		"first":            githubv4.Int(30),
		"after":            (*githubv4.String)(nil),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
	}

	varsOrderedWithCursor := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"first":            githubv4.Int(10),
		"after":            githubv4.String("cursor123"),
		"orderByField":     githubv4.DiscussionOrderField("CREATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
	}

	tests := []struct {
		name          string
		reqParams     map[string]interface{}
		expectError   bool
		errContains   string
		expectedCount int
	}{
		{
			name: "list all discussions without category filter",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
			},
			expectError:   false,
			expectedCount: 3, // All discussions
		},
		{
			name: "filter by category ID",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
			},
			expectError:   false,
			expectedCount: 2, // Only General discussions (matching the category ID)
		},
		{
			name: "list discussions with perPage",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(10),
			},
			expectError:   false,
			expectedCount: 3,
		},
		{
			name: "order by created at descending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "CREATED_AT",
				"direction": "DESC",
			},
			expectError:   false,
			expectedCount: 3,
		},
		{
			name: "order by updated at ascending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "UPDATED_AT",
				"direction": "ASC",
			},
			expectError:   false,
			expectedCount: 3,
		},
		{
			name: "filter by category ID with ordering",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"category":  "DIC_kwDOABC123",
				"orderBy":   "UPDATED_AT",
				"direction": "DESC",
			},
			expectError:   false,
			expectedCount: 2,
		},
		{
			name: "ordering with pagination",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "CREATED_AT",
				"direction": "ASC",
				"perPage":   float64(10),
				"after":     "cursor123",
			},
			expectError:   false,
			expectedCount: 3,
		},
		{
			name: "orderBy without direction is ignored",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"orderBy": "UPDATED_AT",
			},
			expectError:   false,
			expectedCount: 3,
		},
		{
			name: "direction without orderBy is ignored",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"direction": "DESC",
			},
			expectError:   false,
			expectedCount: 3,
		},
		{
			name: "orderBy without direction with category filter is ignored",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
				"orderBy":  "CREATED_AT",
			},
			expectError:   false,
			expectedCount: 2,
		},
		{
			name: "perPage above maximum",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(101),
			},
			expectError: true,
			errContains: "perPage value 101 exceeds maximum of 100",
		},
		{
			name: "repository not found error",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "nonexistent-repo",
			},
			expectError: true,
			errContains: "repository not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var httpClient *http.Client

			switch tc.name {
			case "list all discussions without category filter":
				// Simple case - no category filter
				matcher := githubv4mock.NewQueryMatcher(qDiscussions, varsListAll, mockResponseListAll)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "filter by category ID":
				// Simple case - category filter using category ID directly
				matcher := githubv4mock.NewQueryMatcher(qDiscussionsFiltered, varsDiscussionsFiltered, mockResponseListGeneral)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "list discussions with perPage":
				matcher := githubv4mock.NewQueryMatcher(qDiscussions, varsListWithPerPage, mockResponseListAll)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "order by created at descending":
				matcher := githubv4mock.NewQueryMatcher(qDiscussionsOrdered, varsOrderByCreatedAtDesc, mockResponseListAll)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "order by updated at ascending":
				matcher := githubv4mock.NewQueryMatcher(qDiscussionsOrdered, varsOrderByUpdatedAtAsc, mockResponseListAll)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "filter by category ID with ordering":
				matcher := githubv4mock.NewQueryMatcher(qDiscussionsFilteredOrdered, varsFilteredOrdered, mockResponseListGeneral)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "ordering with pagination":
				matcher := githubv4mock.NewQueryMatcher(qDiscussionsOrdered, varsOrderedWithCursor, mockResponseListAll)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "orderBy without direction is ignored", "direction without orderBy is ignored":
				// With only one of the two, the unordered query is sent.
				matcher := githubv4mock.NewQueryMatcher(qDiscussions, varsListAll, mockResponseListAll)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "orderBy without direction with category filter is ignored":
				matcher := githubv4mock.NewQueryMatcher(qDiscussionsFiltered, varsDiscussionsFiltered, mockResponseListGeneral)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			case "perPage above maximum":
				// The request is rejected before any query is sent.
				httpClient = githubv4mock.NewMockedHTTPClient()
			case "repository not found error":
				matcher := githubv4mock.NewQueryMatcher(qDiscussions, varsRepoNotFound, mockErrorRepoNotFound)
				httpClient = githubv4mock.NewMockedHTTPClient(matcher)
			}

			gqlClient := githubv4.NewClient(httpClient)
			_, handler := ListDiscussions(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			req := createMCPRequest(tc.reqParams)
			res, err := handler(context.Background(), req)
			text := getTextResult(t, res).Text

			if tc.expectError {
				require.True(t, res.IsError)
				assert.Contains(t, text, tc.errContains)
				return
			}
			require.NoError(t, err)

			var response struct {
				Discussions []*github.Discussion `json:"discussions"`
				PageInfo    struct {
					HasNextPage     bool   `json:"hasNextPage"`
					HasPreviousPage bool   `json:"hasPreviousPage"`
					StartCursor     string `json:"startCursor"`
					EndCursor       string `json:"endCursor"`
				} `json:"pageInfo"`
				TotalCount int `json:"totalCount"`
			}
			err = json.Unmarshal([]byte(text), &response)
			require.NoError(t, err)
			returnedDiscussions := response.Discussions

			assert.Len(t, returnedDiscussions, tc.expectedCount, "Expected %d discussions, got %d", tc.expectedCount, len(returnedDiscussions))

			assert.Equal(t, tc.expectedCount, response.TotalCount)
			assert.False(t, response.PageInfo.HasNextPage)

			// Every discussion carries its update time and its author
			for _, discussion := range returnedDiscussions {
				require.NotNil(t, discussion.UpdatedAt, "Discussion should have updatedAt")
				assert.False(t, discussion.UpdatedAt.IsZero(), "Discussion should have a non-zero updatedAt")
				require.NotNil(t, discussion.User, "Discussion should have an author")
				assert.Equal(t, fmt.Sprintf("user%d", discussion.GetNumber()), discussion.GetUser().GetLogin())
			}

			// Verify that all returned discussions have a category if filtered
			if _, hasCategory := tc.reqParams["category"]; hasCategory {
				for _, discussion := range returnedDiscussions {
					require.NotNil(t, discussion.DiscussionCategory, "Discussion should have category")
					assert.NotEmpty(t, *discussion.DiscussionCategory.Name, "Discussion should have category name")
				}
			}
		})
	}
}

func Test_GetDiscussion(t *testing.T) {
	// Verify tool definition and schema
	toolDef, _ := GetDiscussion(nil, translations.NullTranslationHelper)
	assert.Equal(t, "get_discussion", toolDef.Name)
	assert.NotEmpty(t, toolDef.Description)
	assert.Contains(t, toolDef.InputSchema.Properties, "owner")
	assert.Contains(t, toolDef.InputSchema.Properties, "repo")
	assert.Contains(t, toolDef.InputSchema.Properties, "discussionNumber")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo", "discussionNumber"})

	var q struct {
		Repository struct {
			Discussion struct {
				Number    githubv4.Int
				Body      githubv4.String
				CreatedAt githubv4.DateTime
				URL       githubv4.String `graphql:"url"`
				Category  struct {
					Name githubv4.String
				} `graphql:"category"`
			} `graphql:"discussion(number: $discussionNumber)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	vars := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"discussionNumber": githubv4.Int(1),
	}
	tests := []struct {
		name        string
		response    githubv4mock.GQLResponse
		expectError bool
		expected    *github.Discussion
		errContains string
	}{
		{
			name: "successful retrieval",
			response: githubv4mock.DataResponse(map[string]any{
				"repository": map[string]any{"discussion": map[string]any{
					"number":    1,
					"body":      "This is a test discussion",
					"url":       "https://github.com/owner/repo/discussions/1",
					"createdAt": "2025-04-25T12:00:00Z",
					"category":  map[string]any{"name": "General"},
				}},
			}),
			expectError: false,
			expected: &github.Discussion{
				HTMLURL:   github.Ptr("https://github.com/owner/repo/discussions/1"),
				Number:    github.Ptr(1),
				Body:      github.Ptr("This is a test discussion"),
				CreatedAt: &github.Timestamp{Time: time.Date(2025, 4, 25, 12, 0, 0, 0, time.UTC)},
				DiscussionCategory: &github.DiscussionCategory{
					Name: github.Ptr("General"),
				},
			},
		},
		{
			name:        "discussion not found",
			response:    githubv4mock.ErrorResponse("discussion not found"),
			expectError: true,
			errContains: "discussion not found",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matcher := githubv4mock.NewQueryMatcher(q, vars, tc.response)
			httpClient := githubv4mock.NewMockedHTTPClient(matcher)
			gqlClient := githubv4.NewClient(httpClient)
			_, handler := GetDiscussion(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			req := createMCPRequest(map[string]interface{}{"owner": "owner", "repo": "repo", "discussionNumber": int32(1)})
			res, err := handler(context.Background(), req)
			text := getTextResult(t, res).Text

			if tc.expectError {
				require.True(t, res.IsError)
				assert.Contains(t, text, tc.errContains)
				return
			}

			require.NoError(t, err)
			var out github.Discussion
			require.NoError(t, json.Unmarshal([]byte(text), &out))
			assert.Equal(t, *tc.expected.HTMLURL, *out.HTMLURL)
			assert.Equal(t, *tc.expected.Number, *out.Number)
			assert.Equal(t, *tc.expected.Body, *out.Body)
			// Check category label
			assert.Equal(t, *tc.expected.DiscussionCategory.Name, *out.DiscussionCategory.Name)
		})
	}
}

func Test_GetDiscussionComments(t *testing.T) {
	// Verify tool definition and schema
	toolDef, _ := GetDiscussionComments(nil, translations.NullTranslationHelper)
	assert.Equal(t, "get_discussion_comments", toolDef.Name)
	assert.NotEmpty(t, toolDef.Description)
	assert.Contains(t, toolDef.InputSchema.Properties, "owner")
	assert.Contains(t, toolDef.InputSchema.Properties, "repo")
	assert.Contains(t, toolDef.InputSchema.Properties, "discussionNumber")
	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo", "discussionNumber"})

	var q struct {
		Repository struct {
			Discussion struct {
				Comments struct {
					Nodes []struct {
						Body githubv4.String
					}
					PageInfo struct {
						HasNextPage     bool
						HasPreviousPage bool
						StartCursor     string
						EndCursor       string
					}
					TotalCount int
				} `graphql:"comments(first: $first, after: $after)"`
			} `graphql:"discussion(number: $discussionNumber)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	vars := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"discussionNumber": githubv4.Int(1),
		"first":            githubv4.Int(30),
		"after":            (*githubv4.String)(nil),
	}
	mockResponse := githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussion": map[string]any{
				"comments": map[string]any{
					"nodes": []map[string]any{
						{"body": "This is the first comment"},
						{"body": "This is the second comment"},
					},
					"pageInfo": map[string]any{
						"hasNextPage":     true,
						"hasPreviousPage": false,
						"startCursor":     "start",
						"endCursor":       "end",
					},
					"totalCount": 5,
				},
			},
		},
	})
	matcher := githubv4mock.NewQueryMatcher(q, vars, mockResponse)
	httpClient := githubv4mock.NewMockedHTTPClient(matcher)
	gqlClient := githubv4.NewClient(httpClient)
	_, handler := GetDiscussionComments(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

	request := createMCPRequest(map[string]interface{}{
		"owner":            "owner",
		"repo":             "repo",
		"discussionNumber": int32(1),
	})

	result, err := handler(context.Background(), request)
	require.NoError(t, err)

	textContent := getTextResult(t, result)

	var response struct {
		Comments []*github.IssueComment `json:"comments"`
		PageInfo struct {
			HasNextPage     bool   `json:"hasNextPage"`
			HasPreviousPage bool   `json:"hasPreviousPage"`
			StartCursor     string `json:"startCursor"`
			EndCursor       string `json:"endCursor"`
		} `json:"pageInfo"`
		TotalCount int `json:"totalCount"`
	}
	err = json.Unmarshal([]byte(textContent.Text), &response)
	require.NoError(t, err)
	returnedComments := response.Comments
	assert.Len(t, returnedComments, 2)
	assert.True(t, response.PageInfo.HasNextPage)
	assert.False(t, response.PageInfo.HasPreviousPage)
	assert.Equal(t, "start", response.PageInfo.StartCursor)
	assert.Equal(t, "end", response.PageInfo.EndCursor)
	assert.Equal(t, 5, response.TotalCount)
	expectedBodies := []string{"This is the first comment", "This is the second comment"}
	for i, comment := range returnedComments {
		assert.Equal(t, expectedBodies[i], *comment.Body)
	}
}

func Test_ListDiscussionCategories(t *testing.T) {
	var q struct {
		Repository struct {
			DiscussionCategories struct {
				Nodes []struct {
					ID   githubv4.ID
					Name githubv4.String
				}
				PageInfo struct {
					HasNextPage     bool
					HasPreviousPage bool
					StartCursor     string
					EndCursor       string
				}
				TotalCount int
			} `graphql:"discussionCategories(first: $first)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	vars := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(25),
	}
	mockResp := githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussionCategories": map[string]any{
				"nodes": []map[string]any{
					{"id": "123", "name": "CategoryOne"},
					{"id": "456", "name": "CategoryTwo"},
				},
				"pageInfo": map[string]any{
					"hasNextPage":     false,
					"hasPreviousPage": false,
					"startCursor":     "",
					"endCursor":       "",
				},
				"totalCount": 2,
			},
		},
	})
	matcher := githubv4mock.NewQueryMatcher(q, vars, mockResp)
	httpClient := githubv4mock.NewMockedHTTPClient(matcher)
	gqlClient := githubv4.NewClient(httpClient)

	tool, handler := ListDiscussionCategories(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)
	assert.Equal(t, "list_discussion_categories", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "owner")
	assert.Contains(t, tool.InputSchema.Properties, "repo")
	assert.NotContains(t, tool.InputSchema.Properties, "first")
	assert.NotContains(t, tool.InputSchema.Properties, "last")
	assert.NotContains(t, tool.InputSchema.Properties, "after")
	assert.NotContains(t, tool.InputSchema.Properties, "before")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo"})

	request := createMCPRequest(map[string]interface{}{"owner": "owner", "repo": "repo"})
	result, err := handler(context.Background(), request)
	require.NoError(t, err)

	text := getTextResult(t, result).Text
	var response struct {
		Categories []map[string]string `json:"categories"`
		PageInfo   struct {
			HasNextPage bool `json:"hasNextPage"`
		} `json:"pageInfo"`
		TotalCount int `json:"totalCount"`
	}
	require.NoError(t, json.Unmarshal([]byte(text), &response))
	categories := response.Categories
	assert.Len(t, categories, 2)
	assert.False(t, response.PageInfo.HasNextPage)
	assert.Equal(t, 2, response.TotalCount)
	assert.Equal(t, "123", categories[0]["id"])
	assert.Equal(t, "CategoryOne", categories[0]["name"])
	assert.Equal(t, "456", categories[1]["id"])
	assert.Equal(t, "CategoryTwo", categories[1]["name"])
}

```

#### y/issues.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

// GetIssue creates a tool to get details of a specific issue in a GitHub repository.
func GetIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_issue",
			mcp.WithDescription(t("TOOL_GET_ISSUE_DESCRIPTION", "Get details of a specific issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_ISSUE_USER_TITLE", "Get issue details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("The owner of the repository"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("The name of the repository"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("The number of the issue"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			issue, resp, err := client.Issues.Get(ctx, owner, repo, issueNumber)
			if err != nil {
				return nil, fmt.Errorf("failed to get issue: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get issue: %s", string(body))), nil
			}

			r, err := json.Marshal(issue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal issue: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// AddIssueComment creates a tool to add a comment to an issue.
func AddIssueComment(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("add_issue_comment",
			mcp.WithDescription(t("TOOL_ADD_ISSUE_COMMENT_DESCRIPTION", "Add a comment to a specific issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_ADD_ISSUE_COMMENT_USER_TITLE", "Add comment to issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("Issue number to comment on"),
			),
			mcp.WithString("body",
				mcp.Required(),
				mcp.Description("Comment content"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			body, err := RequiredParam[string](request, "body")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			comment := &github.IssueComment{
				Body: github.Ptr(body),
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			createdComment, resp, err := client.Issues.CreateComment(ctx, owner, repo, issueNumber, comment)
			if err != nil {
				return nil, fmt.Errorf("failed to create comment: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create comment: %s", string(body))), nil
			}

			r, err := json.Marshal(createdComment)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// SearchIssues creates a tool to search for issues.
func SearchIssues(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_issues",
			mcp.WithDescription(t("TOOL_SEARCH_ISSUES_DESCRIPTION", "Search for issues in GitHub repositories using issues search syntax already scoped to is:issue")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SEARCH_ISSUES_USER_TITLE", "Search issues"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("Search query using GitHub issues search syntax"),
			),
			mcp.WithString("owner",
				mcp.Description("Optional repository owner. If provided with repo, only notifications for this repository are listed."),
			),
			mcp.WithString("repo",
				mcp.Description("Optional repository name. If provided with owner, only notifications for this repository are listed."),
			),
			mcp.WithString("sort",
				mcp.Description("Sort field by number of matches of categories, defaults to best match"),
				mcp.Enum(
					"comments",
					"reactions",
					"reactions-+1",
					"reactions--1",
					"reactions-smile",
					"reactions-thinking_face",
					"reactions-heart",
					"reactions-tada",
					"interactions",
					"created",
					"updated",
				),
			),
			mcp.WithString("order",
				mcp.Description("Sort order"),
				mcp.Enum("asc", "desc"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return searchHandler(ctx, getClient, request, "issue", "failed to search issues")
		}
}

// CreateIssue creates a tool to create a new issue in a GitHub repository.
func CreateIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("create_issue",
			mcp.WithDescription(t("TOOL_CREATE_ISSUE_DESCRIPTION", "Create a new issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_ISSUE_USER_TITLE", "Open new issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("title",
				mcp.Required(),
				mcp.Description("Issue title"),
			),
			mcp.WithString("body",
				mcp.Description("Issue body content"),
			),
			mcp.WithArray("assignees",
				mcp.Description("Usernames to assign to this issue"),
				mcp.Items(
					map[string]any{
						"type": "string",
					},
				),
			),
			mcp.WithArray("labels",
				mcp.Description("Labels to apply to this issue"),
				mcp.Items(
					map[string]any{
						"type": "string",
					},
				),
			),
			mcp.WithNumber("milestone",
				mcp.Description("Milestone number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			title, err := RequiredParam[string](request, "title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Optional parameters
			body, err := OptionalParam[string](request, "body")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get assignees
			assignees, err := OptionalStringArrayParam(request, "assignees")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get labels
			labels, err := OptionalStringArrayParam(request, "labels")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get optional milestone
			milestone, err := OptionalIntParam(request, "milestone")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var milestoneNum *int
			if milestone != 0 {
				milestoneNum = &milestone
			}

			// Create the issue request
			issueRequest := &github.IssueRequest{
				Title:     github.Ptr(title),
				Body:      github.Ptr(body),
				Assignees: &assignees,
				Labels:    &labels,
				Milestone: milestoneNum,
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			issue, resp, err := client.Issues.Create(ctx, owner, repo, issueRequest)
			if err != nil {
				return nil, fmt.Errorf("failed to create issue: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create issue: %s", string(body))), nil
			}

			r, err := json.Marshal(issue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListIssues creates a tool to list and filter repository issues
func ListIssues(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_issues",
			mcp.WithDescription(t("TOOL_LIST_ISSUES_DESCRIPTION", "List issues in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_ISSUES_USER_TITLE", "List issues"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("state",
				mcp.Description("Filter by state"),
				mcp.Enum("open", "closed", "all"),
			),
			mcp.WithArray("labels",
				mcp.Description("Filter by labels"),
				mcp.Items(
					map[string]interface{}{
						"type": "string",
					},
				),
			),
			mcp.WithString("sort",
				mcp.Description("Sort order"),
				mcp.Enum("created", "updated", "comments"),
			),
			mcp.WithString("direction",
				mcp.Description("Sort direction"),
				mcp.Enum("asc", "desc"),
			),
			mcp.WithString("since",
				mcp.Description("Filter by date (ISO 8601 timestamp)"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.IssueListByRepoOptions{}

			// Set optional parameters if provided
			opts.State, err = OptionalParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Get labels
			opts.Labels, err = OptionalStringArrayParam(request, "labels")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts.Sort, err = OptionalParam[string](request, "sort")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts.Direction, err = OptionalParam[string](request, "direction")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			since, err := OptionalParam[string](request, "since")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if since != "" {
				timestamp, err := parseISOTimestamp(since)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("failed to list issues: %s", err.Error())), nil
				}
				opts.Since = timestamp
			}

			if page, ok := request.GetArguments()["page"].(float64); ok {
				opts.ListOptions.Page = int(page)
			}

			if perPage, ok := request.GetArguments()["perPage"].(float64); ok {
				opts.ListOptions.PerPage = int(perPage)
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			issues, resp, err := client.Issues.ListByRepo(ctx, owner, repo, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to list issues: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list issues: %s", string(body))), nil
			}

			r, err := json.Marshal(issues)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal issues: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// UpdateIssue creates a tool to update an existing issue in a GitHub repository.
func UpdateIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("update_issue",
			mcp.WithDescription(t("TOOL_UPDATE_ISSUE_DESCRIPTION", "Update an existing issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_UPDATE_ISSUE_USER_TITLE", "Edit issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("Issue number to update"),
			),
			mcp.WithString("title",
				mcp.Description("New title"),
			),
			mcp.WithString("body",
				mcp.Description("New description"),
			),
			mcp.WithString("state",
				mcp.Description("New state"),
				mcp.Enum("open", "closed"),
			),
			mcp.WithArray("labels",
				mcp.Description("New labels"),
				mcp.Items(
					map[string]interface{}{
						"type": "string",
					},
				),
			),
			mcp.WithArray("assignees",
				mcp.Description("New assignees"),
				mcp.Items(
					map[string]interface{}{
						"type": "string",
					},
				),
			),
			mcp.WithNumber("milestone",
				mcp.Description("New milestone number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Create the issue request with only provided fields
			issueRequest := &github.IssueRequest{}

			// Set optional parameters if provided
			title, err := OptionalParam[string](request, "title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if title != "" {
				issueRequest.Title = github.Ptr(title)
			}

			body, err := OptionalParam[string](request, "body")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if body != "" {
				issueRequest.Body = github.Ptr(body)
			}

			state, err := OptionalParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if state != "" {
				issueRequest.State = github.Ptr(state)
			}

			// Get labels
			labels, err := OptionalStringArrayParam(request, "labels")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if len(labels) > 0 {
				issueRequest.Labels = &labels
			}

			// Get assignees
			assignees, err := OptionalStringArrayParam(request, "assignees")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if len(assignees) > 0 {
				issueRequest.Assignees = &assignees
			}

			milestone, err := OptionalIntParam(request, "milestone")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if milestone != 0 {
				milestoneNum := milestone
				issueRequest.Milestone = &milestoneNum
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			updatedIssue, resp, err := client.Issues.Edit(ctx, owner, repo, issueNumber, issueRequest)
			if err != nil {
				return nil, fmt.Errorf("failed to update issue: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to update issue: %s", string(body))), nil
			}

			r, err := json.Marshal(updatedIssue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetIssueComments creates a tool to get comments for a GitHub issue.
func GetIssueComments(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_issue_comments",
			mcp.WithDescription(t("TOOL_GET_ISSUE_COMMENTS_DESCRIPTION", "Get comments for a specific issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_ISSUE_COMMENTS_USER_TITLE", "Get issue comments"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("Issue number"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.IssueListCommentsOptions{
				ListOptions: github.ListOptions{
					Page:    pagination.Page,
					PerPage: pagination.PerPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			comments, resp, err := client.Issues.ListComments(ctx, owner, repo, issueNumber, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to get issue comments: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get issue comments: %s", string(body))), nil
			}

			r, err := json.Marshal(comments)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// mvpDescription is an MVP idea for generating tool descriptions from structured data in a shared format.
// It is not intended for widespread usage and is not a complete implementation.
type mvpDescription struct {
	summary        string
	outcomes       []string
	referenceLinks []string
}

func (d *mvpDescription) String() string {
	var sb strings.Builder
	sb.WriteString(d.summary)
	if len(d.outcomes) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString("This tool can help with the following outcomes:\n")
		for _, outcome := range d.outcomes {
			sb.WriteString(fmt.Sprintf("- %s\n", outcome))
		}
	}

	if len(d.referenceLinks) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString("More information can be found at:\n")
		for _, link := range d.referenceLinks {
			sb.WriteString(fmt.Sprintf("- %s\n", link))
		}
	}

	return sb.String()
}

func AssignCopilotToIssue(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	description := mvpDescription{
		summary: "Assign Copilot to a specific issue in a GitHub repository.",
		outcomes: []string{
			"a Pull Request created with source code changes to resolve the issue",
		},
		referenceLinks: []string{
			"https://docs.github.com/en/copilot/using-github-copilot/using-copilot-coding-agent-to-work-on-tasks/about-assigning-tasks-to-copilot",
		},
	}

	return mcp.NewTool("assign_copilot_to_issue",
			mcp.WithDescription(t("TOOL_ASSIGN_COPILOT_TO_ISSUE_DESCRIPTION", description.String())),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:          t("TOOL_ASSIGN_COPILOT_TO_ISSUE_USER_TITLE", "Assign Copilot to issue"),
				ReadOnlyHint:   ToBoolPtr(false),
				IdempotentHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issueNumber",
				mcp.Required(),
				mcp.Description("Issue number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner       string
				Repo        string
				IssueNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Firstly, we try to find the copilot bot in the suggested actors for the repository.
			// Although as I write this, we would expect copilot to be at the top of the list, in future, maybe
			// it will not be on the first page of responses, thus we will keep paginating until we find it.
			type botAssignee struct {
				ID       githubv4.ID
				Login    string
				TypeName string `graphql:"__typename"`
			}

			type suggestedActorsQuery struct {
				Repository struct {
					SuggestedActors struct {
						Nodes []struct {
							Bot botAssignee `graphql:"... on Bot"`
						}
						PageInfo struct {
							HasNextPage bool
							EndCursor   string
						}
					} `graphql:"suggestedActors(first: 100, after: $endCursor, capabilities: CAN_BE_ASSIGNED)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			variables := map[string]any{
				"owner":     githubv4.String(params.Owner),
				"name":      githubv4.String(params.Repo),
				"endCursor": (*githubv4.String)(nil),
			}

			var copilotAssignee *botAssignee
			for {
				var query suggestedActorsQuery
				err := client.Query(ctx, &query, variables)
				if err != nil {
					return nil, err
				}

				// Iterate all the returned nodes looking for the copilot bot, which is supposed to have the
				// same name on each host. We need this in order to get the ID for later assignment.
				for _, node := range query.Repository.SuggestedActors.Nodes {
					if node.Bot.Login == "copilot-swe-agent" {
						copilotAssignee = &node.Bot
						break
					}
				}

				if !query.Repository.SuggestedActors.PageInfo.HasNextPage {
					break
				}
				variables["endCursor"] = githubv4.String(query.Repository.SuggestedActors.PageInfo.EndCursor)
			}

			// If we didn't find the copilot bot, we can't proceed any further.
			if copilotAssignee == nil {
				// The e2e tests depend upon this specific message to skip the test.
				return mcp.NewToolResultError("copilot isn't available as an assignee for this issue. Please inform the user to visit https://docs.github.com/en/copilot/using-github-copilot/using-copilot-coding-agent-to-work-on-tasks/about-assigning-tasks-to-copilot for more information."), nil
			}

			// Next let's get the GQL Node ID and current assignees for this issue because the only way to
			// assign copilot is to use replaceActorsForAssignable which requires the full list.
			var getIssueQuery struct {
				Repository struct {
					Issue struct {
						ID        githubv4.ID
						Assignees struct {
							Nodes []struct {
								ID githubv4.ID
							}
						} `graphql:"assignees(first: 100)"`
					} `graphql:"issue(number: $number)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			variables = map[string]any{
				"owner":  githubv4.String(params.Owner),
				"name":   githubv4.String(params.Repo),
				"number": githubv4.Int(params.IssueNumber),
			}

			if err := client.Query(ctx, &getIssueQuery, variables); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get issue ID: %v", err)), nil
			}

			// Finally, do the assignment. Just for reference, assigning copilot to an issue that it is already
			// assigned to seems to have no impact (which is a good thing).
			var assignCopilotMutation struct {
				ReplaceActorsForAssignable struct {
					Typename string `graphql:"__typename"` // Not required but we need a selector or GQL errors
				} `graphql:"replaceActorsForAssignable(input: $input)"`
			}

			actorIDs := make([]githubv4.ID, len(getIssueQuery.Repository.Issue.Assignees.Nodes)+1)
			for i, node := range getIssueQuery.Repository.Issue.Assignees.Nodes {
				actorIDs[i] = node.ID
			}
			actorIDs[len(getIssueQuery.Repository.Issue.Assignees.Nodes)] = copilotAssignee.ID

			if err := client.Mutate(
				ctx,
				&assignCopilotMutation,
				ReplaceActorsForAssignableInput{
					AssignableID: getIssueQuery.Repository.Issue.ID,
					ActorIDs:     actorIDs,
				},
				nil,
			); err != nil {
				return nil, fmt.Errorf("failed to replace actors for assignable: %w", err)
			}

			return mcp.NewToolResultText("successfully assigned copilot to issue"), nil
		}
}

type ReplaceActorsForAssignableInput struct {
	AssignableID githubv4.ID   `json:"assignableId"`
	ActorIDs     []githubv4.ID `json:"actorIds"`
}

// parseISOTimestamp parses an ISO 8601 timestamp string into a time.Time object.
// Returns the parsed time or an error if parsing fails.
// Example formats supported: "2023-01-15T14:30:00Z", "2023-01-15"
func parseISOTimestamp(timestamp string) (time.Time, error) {
	if timestamp == "" {
		return time.Time{}, fmt.Errorf("empty timestamp")
	}

	// Try RFC3339 format (standard ISO 8601 with time)
	t, err := time.Parse(time.RFC3339, timestamp)
	if err == nil {
		return t, nil
	}

	// Try simple date format (YYYY-MM-DD)
	t, err = time.Parse("2006-01-02", timestamp)
	if err == nil {
		return t, nil
	}

	// Return error with supported formats
	return time.Time{}, fmt.Errorf("invalid ISO 8601 timestamp: %s (supported formats: YYYY-MM-DDThh:mm:ssZ or YYYY-MM-DD)", timestamp)
}

func AssignCodingAgentPrompt(t translations.TranslationHelperFunc) (tool mcp.Prompt, handler server.PromptHandlerFunc) {
	return mcp.NewPrompt("AssignCodingAgent",
			mcp.WithPromptDescription(t("PROMPT_ASSIGN_CODING_AGENT_DESCRIPTION", "Assign GitHub Coding Agent to multiple tasks in a GitHub repository.")),
			mcp.WithArgument("repo", mcp.ArgumentDescription("The repository to assign tasks in (owner/repo)."), mcp.RequiredArgument()),
		), func(_ context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			repo := request.Params.Arguments["repo"]

			messages := []mcp.PromptMessage{
				{
					Role:    "system",
					Content: mcp.NewTextContent("You are a personal assistant for GitHub the Copilot GitHub Coding Agent. Your task is to help the user assign tasks to the Coding Agent based on their open GitHub issues. You can use `assign_copilot_to_issue` tool to assign the Coding Agent to issues that are suitable for autonomous work, and `search_issues` tool to find issues that match the user's criteria. You can also use `list_issues` to get a list of issues in the repository."),
				},
				{
					Role:    "user",
					Content: mcp.NewTextContent(fmt.Sprintf("Please go and get a list of the most recent 10 issues from the %s GitHub repository", repo)),
				},
				{
					Role:    "assistant",
					Content: mcp.NewTextContent(fmt.Sprintf("Sure! I will get a list of the 10 most recent issues for the repo %s.", repo)),
				},
				{
					Role:    "user",
					Content: mcp.NewTextContent("For each issue, please check if it is a clearly defined coding task with acceptance criteria and a low to medium complexity to identify issues that are suitable for an AI Coding Agent to work on. Then assign each of the identified issues to Copilot."),
				},
				{
					Role:    "assistant",
					Content: mcp.NewTextContent("Certainly! Let me carefully check which ones are clearly scoped issues that are good to assign to the coding agent, and I will summarize and assign them now."),
				},
				{
					Role:    "user",
					Content: mcp.NewTextContent("Great, if you are unsure if an issue is good to assign, ask me first, rather than assigning copilot. If you are certain the issue is clear and suitable you can assign it to Copilot without asking."),
				},
			}
			return &mcp.GetPromptResult{
				Messages: messages,
			}, nil
		}
}

```

#### y/notifications.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	FilterDefault           = "default"
	FilterIncludeRead       = "include_read_notifications"
	FilterOnlyParticipating = "only_participating"
)

// ListNotifications creates a tool to list notifications for the current user.
func ListNotifications(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_notifications",
			mcp.WithDescription(t("TOOL_LIST_NOTIFICATIONS_DESCRIPTION", "Lists all GitHub notifications for the authenticated user, including unread notifications, mentions, review requests, assignments, and updates on issues or pull requests. Use this tool whenever the user asks what to work on next, requests a summary of their GitHub activity, wants to see pending reviews, or needs to check for new updates or tasks. This tool is the primary way to discover actionable items, reminders, and outstanding work on GitHub. Always call this tool when asked what to work on next, what is pending, or what needs attention in GitHub.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_NOTIFICATIONS_USER_TITLE", "List notifications"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("filter",
				mcp.Description("Filter notifications to, use default unless specified. Read notifications are ones that have already been acknowledged by the user. Participating notifications are those that the user is directly involved in, such as issues or pull requests they have commented on or created."),
				mcp.Enum(FilterDefault, FilterIncludeRead, FilterOnlyParticipating),
			),
			mcp.WithString("since",
				mcp.Description("Only show notifications updated after the given time (ISO 8601 format)"),
			),
			mcp.WithString("before",
				mcp.Description("Only show notifications updated before the given time (ISO 8601 format)"),
			),
			mcp.WithString("owner",
				mcp.Description("Optional repository owner. If provided with repo, only notifications for this repository are listed."),
			),
			mcp.WithString("repo",
				mcp.Description("Optional repository name. If provided with owner, only notifications for this repository are listed."),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			filter, err := OptionalParam[string](request, "filter")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			since, err := OptionalParam[string](request, "since")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			before, err := OptionalParam[string](request, "before")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			owner, err := OptionalParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := OptionalParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			paginationParams, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Build options
			opts := &github.NotificationListOptions{
				All:           filter == FilterIncludeRead,
				Participating: filter == FilterOnlyParticipating,
				ListOptions: github.ListOptions{
					Page:    paginationParams.Page,
					PerPage: paginationParams.PerPage,
				},
			}

			// Parse time parameters if provided
			if since != "" {
				sinceTime, err := time.Parse(time.RFC3339, since)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("invalid since time format, should be RFC3339/ISO8601: %v", err)), nil
				}
				opts.Since = sinceTime
			}

			if before != "" {
				beforeTime, err := time.Parse(time.RFC3339, before)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("invalid before time format, should be RFC3339/ISO8601: %v", err)), nil
				}
				opts.Before = beforeTime
			}

			var notifications []*github.Notification
			var resp *github.Response

			if owner != "" && repo != "" {
				notifications, resp, err = client.Activity.ListRepositoryNotifications(ctx, owner, repo, opts)
			} else {
				notifications, resp, err = client.Activity.ListNotifications(ctx, opts)
			}
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to list notifications",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get notifications: %s", string(body))), nil
			}

			// Marshal response to JSON
			r, err := json.Marshal(notifications)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// DismissNotification creates a tool to mark a notification as read/done.
func DismissNotification(getclient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("dismiss_notification",
			mcp.WithDescription(t("TOOL_DISMISS_NOTIFICATION_DESCRIPTION", "Dismiss a notification by marking it as read or done")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_DISMISS_NOTIFICATION_USER_TITLE", "Dismiss notification"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("threadID",
				mcp.Required(),
				mcp.Description("The ID of the notification thread"),
			),
			mcp.WithString("state", mcp.Description("The new state of the notification (read/done)"), mcp.Enum("read", "done")),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getclient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			threadID, err := RequiredParam[string](request, "threadID")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			state, err := RequiredParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var resp *github.Response
			switch state {
			case "done":
				// for some inexplicable reason, the API seems to have threadID as int64 and string depending on the endpoint
				var threadIDInt int64
				threadIDInt, err = strconv.ParseInt(threadID, 10, 64)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("invalid threadID format: %v", err)), nil
				}
				resp, err = client.Activity.MarkThreadDone(ctx, threadIDInt)
			case "read":
				resp, err = client.Activity.MarkThreadRead(ctx, threadID)
			default:
				return mcp.NewToolResultError("Invalid state. Must be one of: read, done."), nil
			}

			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to mark notification as %s", state),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusResetContent && resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to mark notification as %s: %s", state, string(body))), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Notification marked as %s", state)), nil
		}
}

// MarkAllNotificationsRead creates a tool to mark all notifications as read.
func MarkAllNotificationsRead(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("mark_all_notifications_read",
			mcp.WithDescription(t("TOOL_MARK_ALL_NOTIFICATIONS_READ_DESCRIPTION", "Mark all notifications as read")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_MARK_ALL_NOTIFICATIONS_READ_USER_TITLE", "Mark all notifications as read"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("lastReadAt",
				mcp.Description("Describes the last point that notifications were checked (optional). Default: Now"),
			),
			mcp.WithString("owner",
				mcp.Description("Optional repository owner. If provided with repo, only notifications for this repository are marked as read."),
			),
			mcp.WithString("repo",
				mcp.Description("Optional repository name. If provided with owner, only notifications for this repository are marked as read."),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			lastReadAt, err := OptionalParam[string](request, "lastReadAt")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			owner, err := OptionalParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := OptionalParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var lastReadTime time.Time
			if lastReadAt != "" {
				lastReadTime, err = time.Parse(time.RFC3339, lastReadAt)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("invalid lastReadAt time format, should be RFC3339/ISO8601: %v", err)), nil
				}
			} else {
				lastReadTime = time.Now()
			}

			markReadOptions := github.Timestamp{
				Time: lastReadTime,
			}

			var resp *github.Response
			if owner != "" && repo != "" {
				resp, err = client.Activity.MarkRepositoryNotificationsRead(ctx, owner, repo, markReadOptions)
			} else {
				resp, err = client.Activity.MarkNotificationsRead(ctx, markReadOptions)
			}
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to mark all notifications as read",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusResetContent && resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to mark all notifications as read: %s", string(body))), nil
			}

			return mcp.NewToolResultText("All notifications marked as read"), nil
		}
}

// GetNotificationDetails creates a tool to get details for a specific notification.
func GetNotificationDetails(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_notification_details",
			mcp.WithDescription(t("TOOL_GET_NOTIFICATION_DETAILS_DESCRIPTION", "Get detailed information for a specific GitHub notification, always call this tool when the user asks for details about a specific notification, if you don't know the ID list notifications first.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_NOTIFICATION_DETAILS_USER_TITLE", "Get notification details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("notificationID",
				mcp.Required(),
				mcp.Description("The ID of the notification"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			notificationID, err := RequiredParam[string](request, "notificationID")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			thread, resp, err := client.Activity.GetThread(ctx, notificationID)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to get notification details for ID '%s'", notificationID),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get notification details: %s", string(body))), nil
			}

			r, err := json.Marshal(thread)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// Enum values for ManageNotificationSubscription action
const (
	NotificationActionIgnore = "ignore"
	NotificationActionWatch  = "watch"
	NotificationActionDelete = "delete"
)

// ManageNotificationSubscription creates a tool to manage a notification subscription (ignore, watch, delete)
func ManageNotificationSubscription(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("manage_notification_subscription",
			mcp.WithDescription(t("TOOL_MANAGE_NOTIFICATION_SUBSCRIPTION_DESCRIPTION", "Manage a notification subscription: ignore, watch, or delete a notification thread subscription.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_MANAGE_NOTIFICATION_SUBSCRIPTION_USER_TITLE", "Manage notification subscription"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("notificationID",
				mcp.Required(),
				mcp.Description("The ID of the notification thread."),
			),
			mcp.WithString("action",
				mcp.Required(),
				mcp.Description("Action to perform: ignore, watch, or delete the notification subscription."),
				mcp.Enum(NotificationActionIgnore, NotificationActionWatch, NotificationActionDelete),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			notificationID, err := RequiredParam[string](request, "notificationID")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			action, err := RequiredParam[string](request, "action")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var (
				resp   *github.Response
				result any
				apiErr error
			)

			switch action {
			case NotificationActionIgnore:
				sub := &github.Subscription{Ignored: ToBoolPtr(true)}
				result, resp, apiErr = client.Activity.SetThreadSubscription(ctx, notificationID, sub)
			case NotificationActionWatch:
				sub := &github.Subscription{Ignored: ToBoolPtr(false), Subscribed: ToBoolPtr(true)}
				result, resp, apiErr = client.Activity.SetThreadSubscription(ctx, notificationID, sub)
			case NotificationActionDelete:
				resp, apiErr = client.Activity.DeleteThreadSubscription(ctx, notificationID)
			default:
				return mcp.NewToolResultError("Invalid action. Must be one of: ignore, watch, delete."), nil
			}

			if apiErr != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to %s notification subscription", action),
					resp,
					apiErr,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				body, _ := io.ReadAll(resp.Body)
				return mcp.NewToolResultError(fmt.Sprintf("failed to %s notification subscription: %s", action, string(body))), nil
			}

			if action == NotificationActionDelete {
				// Special case for delete as there is no response body
				return mcp.NewToolResultText("Notification subscription deleted"), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

const (
	RepositorySubscriptionActionWatch  = "watch"
	RepositorySubscriptionActionIgnore = "ignore"
	RepositorySubscriptionActionDelete = "delete"
)

// ManageRepositoryNotificationSubscription creates a tool to manage a repository notification subscription (ignore, watch, delete)
func ManageRepositoryNotificationSubscription(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("manage_repository_notification_subscription",
			mcp.WithDescription(t("TOOL_MANAGE_REPOSITORY_NOTIFICATION_SUBSCRIPTION_DESCRIPTION", "Manage a repository notification subscription: ignore, watch, or delete repository notifications subscription for the provided repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_MANAGE_REPOSITORY_NOTIFICATION_SUBSCRIPTION_USER_TITLE", "Manage repository notification subscription"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("The account owner of the repository."),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("The name of the repository."),
			),
			mcp.WithString("action",
				mcp.Required(),
				mcp.Description("Action to perform: ignore, watch, or delete the repository notification subscription."),
				mcp.Enum(RepositorySubscriptionActionIgnore, RepositorySubscriptionActionWatch, RepositorySubscriptionActionDelete),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			action, err := RequiredParam[string](request, "action")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var (
				resp   *github.Response
				result any
				apiErr error
			)

			switch action {
			case RepositorySubscriptionActionIgnore:
				sub := &github.Subscription{Ignored: ToBoolPtr(true)}
				result, resp, apiErr = client.Activity.SetRepositorySubscription(ctx, owner, repo, sub)
			case RepositorySubscriptionActionWatch:
				sub := &github.Subscription{Ignored: ToBoolPtr(false), Subscribed: ToBoolPtr(true)}
				result, resp, apiErr = client.Activity.SetRepositorySubscription(ctx, owner, repo, sub)
			case RepositorySubscriptionActionDelete:
				resp, apiErr = client.Activity.DeleteRepositorySubscription(ctx, owner, repo)
			default:
				return mcp.NewToolResultError("Invalid action. Must be one of: ignore, watch, delete."), nil
			}

			if apiErr != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to %s repository subscription", action),
					resp,
					apiErr,
				), nil
			}
			if resp != nil {
				defer func() { _ = resp.Body.Close() }()
			}

			// Handle non-2xx status codes
			if resp != nil && (resp.StatusCode < 200 || resp.StatusCode >= 300) {
				body, _ := io.ReadAll(resp.Body)
				return mcp.NewToolResultError(fmt.Sprintf("failed to %s repository subscription: %s", action, string(body))), nil
			}

			if action == RepositorySubscriptionActionDelete {
				// Special case for delete as there is no response body
				return mcp.NewToolResultText("Repository subscription deleted"), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

```

#### y/pagination.go

```go
package github

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// DefaultGraphQLPageSize is the default number of items requested per page
// from a GraphQL connection.
//
const DefaultGraphQLPageSize = 30

// CursorPaginationParams holds cursor-only pagination parameters, for tools
// that page through a GraphQL connection without page numbers.
type CursorPaginationParams struct {
	PerPage int
	After   string
}

// ToGraphQLParams converts cursor pagination parameters to GraphQL parameters.
func (p CursorPaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	if p.PerPage > 100 {
		return nil, fmt.Errorf("perPage value %d exceeds maximum of 100", p.PerPage)
	}
	if p.PerPage < 0 {
		return nil, fmt.Errorf("perPage value %d cannot be negative", p.PerPage)
	}
	first := int32(p.PerPage)

	var after *string
	if p.After != "" {
		after = &p.After
	}

	return &GraphQLPaginationParams{
		First: &first,
		After: after,
	}, nil
}

// GraphQLPaginationParams holds the "first" and "after" arguments of a
// GraphQL connection. After is nil when there is no cursor.
// https://docs.github.com/en/graphql/reference/objects#connection
type GraphQLPaginationParams struct {
	First *int32
	After *string
}

// OptionalCursorPaginationParams returns the "perPage" and "after" parameters
// from the request. "perPage" defaults to DefaultGraphQLPageSize, and "after"
// defaults to an empty cursor.
func OptionalCursorPaginationParams(r mcp.CallToolRequest) (CursorPaginationParams, error) {
	perPage, err := OptionalIntParamWithDefault(r, "perPage", DefaultGraphQLPageSize)
	if err != nil {
		return CursorPaginationParams{}, err
	}
	after, err := OptionalParam[string](r, "after")
	if err != nil {
		return CursorPaginationParams{}, err
	}
	return CursorPaginationParams{
		PerPage: perPage,
		After:   after,
	}, nil
}

// WithCursorPagination returns a ToolOption that adds "perPage" and "after"
// parameters to the tool, for GraphQL tools that page by cursor only.
func WithCursorPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		withPaginationPerPage()(tool)
		withPaginationAfter()(tool)
	}
}

// WithUnifiedPagination returns a ToolOption that adds "page", "perPage" and
// "after" parameters to the tool, for GraphQL tools that accept page/perPage
// and convert them to a cursor internally.
func WithUnifiedPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		mcp.WithNumber("page",
			mcp.Description("Page number for pagination (min 1)"),
			mcp.Min(1),
		)(tool)
		withPaginationPerPage()(tool)
		withPaginationAfter()(tool)
	}
}

func withPaginationPerPage() mcp.ToolOption {
	return mcp.WithNumber("perPage",
		mcp.Description("Results per page for pagination (min 1, max 100)"),
		mcp.Min(1),
		mcp.Max(100),
	)
}

func withPaginationAfter() mcp.ToolOption {
	return mcp.WithString("after",
		mcp.Description("Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs."),
	)
}

```

#### y/pagination_test.go

```go
package github

import (
	"testing"

	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCursorPaginationParams_ToGraphQLParams(t *testing.T) {
	tests := []struct {
		name        string
		params      CursorPaginationParams
		expected    *GraphQLPaginationParams
		errContains string
	}{
		{
			name:     "perPage without cursor",
			params:   CursorPaginationParams{PerPage: DefaultGraphQLPageSize},
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(30))},
		},
		{
			name:     "perPage with cursor",
			params:   CursorPaginationParams{PerPage: 100, After: "cursor-abc"},
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(100)), After: github.Ptr("cursor-abc")},
		},
		{
			name:        "perPage above maximum",
			params:      CursorPaginationParams{PerPage: 101},
			errContains: "perPage value 101 exceeds maximum of 100",
		},
		{
			name:        "negative perPage",
			params:      CursorPaginationParams{PerPage: -5},
			errContains: "perPage value -5 cannot be negative",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.params.ToGraphQLParams()
			if tc.errContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errContains)
				assert.Nil(t, result)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestOptionalCursorPaginationParams(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]any
		expected    CursorPaginationParams
		expectError bool
	}{
		{
			name:     "no pagination parameters, default values",
			params:   map[string]any{},
			expected: CursorPaginationParams{PerPage: 30},
		},
		{
			name:     "perPage and after parameters",
			params:   map[string]any{"perPage": float64(50), "after": "cursor-abc"},
			expected: CursorPaginationParams{PerPage: 50, After: "cursor-abc"},
		},
		{
			name:        "invalid perPage parameter",
			params:      map[string]any{"perPage": "not-a-number"},
			expectError: true,
		},
		{
			name:        "invalid after parameter",
			params:      map[string]any{"after": float64(1)},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := OptionalCursorPaginationParams(createMCPRequest(tc.params))
			if tc.expectError {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestWithCursorPagination(t *testing.T) {
	tool := mcp.NewTool("test_tool", WithCursorPagination())

	assert.NotContains(t, tool.InputSchema.Properties, "page")

	perPage, ok := tool.InputSchema.Properties["perPage"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "number", perPage["type"])
	assert.Equal(t, "Results per page for pagination (min 1, max 100)", perPage["description"])
	assert.Equal(t, float64(1), perPage["minimum"])
	assert.Equal(t, float64(100), perPage["maximum"])

	after, ok := tool.InputSchema.Properties["after"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "string", after["type"])
	assert.Equal(t, "Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs.", after["description"])
}

func TestWithUnifiedPagination(t *testing.T) {
	tool := mcp.NewTool("test_tool", WithUnifiedPagination())

	page, ok := tool.InputSchema.Properties["page"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "number", page["type"])
	assert.Equal(t, "Page number for pagination (min 1)", page["description"])
	assert.Equal(t, float64(1), page["minimum"])

	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.Contains(t, tool.InputSchema.Properties, "after")
}

```

#### y/pullrequests.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-viper/mapstructure/v2"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
)

// GetPullRequest creates a tool to get details of a specific pull request.
func GetPullRequest(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_DESCRIPTION", "Get details of a specific pull request in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_USER_TITLE", "Get pull request details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			pr, resp, err := client.PullRequests.Get(ctx, owner, repo, pullNumber)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request: %s", string(body))), nil
			}

			r, err := json.Marshal(pr)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CreatePullRequest creates a tool to create a new pull request.
func CreatePullRequest(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("create_pull_request",
			mcp.WithDescription(t("TOOL_CREATE_PULL_REQUEST_DESCRIPTION", "Create a new pull request in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_PULL_REQUEST_USER_TITLE", "Open new pull request"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("title",
				mcp.Required(),
				mcp.Description("PR title"),
			),
			mcp.WithString("body",
				mcp.Description("PR description"),
			),
			mcp.WithString("head",
				mcp.Required(),
				mcp.Description("Branch containing changes"),
			),
			mcp.WithString("base",
				mcp.Required(),
				mcp.Description("Branch to merge into"),
			),
			mcp.WithBoolean("draft",
				mcp.Description("Create as draft PR"),
			),
			mcp.WithBoolean("maintainer_can_modify",
				mcp.Description("Allow maintainer edits"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			title, err := RequiredParam[string](request, "title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			head, err := RequiredParam[string](request, "head")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			base, err := RequiredParam[string](request, "base")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			body, err := OptionalParam[string](request, "body")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			draft, err := OptionalParam[bool](request, "draft")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			maintainerCanModify, err := OptionalParam[bool](request, "maintainer_can_modify")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			newPR := &github.NewPullRequest{
				Title: github.Ptr(title),
				Head:  github.Ptr(head),
				Base:  github.Ptr(base),
			}

			if body != "" {
				newPR.Body = github.Ptr(body)
			}

			newPR.Draft = github.Ptr(draft)
			newPR.MaintainerCanModify = github.Ptr(maintainerCanModify)

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			pr, resp, err := client.PullRequests.Create(ctx, owner, repo, newPR)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create pull request: %s", string(body))), nil
			}

			r, err := json.Marshal(pr)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// UpdatePullRequest creates a tool to update an existing pull request.
func UpdatePullRequest(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("update_pull_request",
			mcp.WithDescription(t("TOOL_UPDATE_PULL_REQUEST_DESCRIPTION", "Update an existing pull request in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_UPDATE_PULL_REQUEST_USER_TITLE", "Edit pull request"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number to update"),
			),
			mcp.WithString("title",
				mcp.Description("New title"),
			),
			mcp.WithString("body",
				mcp.Description("New description"),
			),
			mcp.WithString("state",
				mcp.Description("New state"),
				mcp.Enum("open", "closed"),
			),
			mcp.WithString("base",
				mcp.Description("New base branch name"),
			),
			mcp.WithBoolean("maintainer_can_modify",
				mcp.Description("Allow maintainer edits"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Build the update struct only with provided fields
			update := &github.PullRequest{}
			updateNeeded := false

			if title, ok, err := OptionalParamOK[string](request, "title"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.Title = github.Ptr(title)
				updateNeeded = true
			}

			if body, ok, err := OptionalParamOK[string](request, "body"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.Body = github.Ptr(body)
				updateNeeded = true
			}

			if state, ok, err := OptionalParamOK[string](request, "state"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.State = github.Ptr(state)
				updateNeeded = true
			}

			if base, ok, err := OptionalParamOK[string](request, "base"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.Base = &github.PullRequestBranch{Ref: github.Ptr(base)}
				updateNeeded = true
			}

			if maintainerCanModify, ok, err := OptionalParamOK[bool](request, "maintainer_can_modify"); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			} else if ok {
				update.MaintainerCanModify = github.Ptr(maintainerCanModify)
				updateNeeded = true
			}

			if !updateNeeded {
				return mcp.NewToolResultError("No update parameters provided."), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			pr, resp, err := client.PullRequests.Edit(ctx, owner, repo, pullNumber, update)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to update pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to update pull request: %s", string(body))), nil
			}

			r, err := json.Marshal(pr)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListPullRequests creates a tool to list and filter repository pull requests.
func ListPullRequests(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("list_pull_requests",
			mcp.WithDescription(t("TOOL_LIST_PULL_REQUESTS_DESCRIPTION", "List pull requests in a GitHub repository. If the user specifies an author, then DO NOT use this tool and use the search_pull_requests tool instead.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_PULL_REQUESTS_USER_TITLE", "List pull requests"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("state",
				mcp.Description("Filter by state"),
				mcp.Enum("open", "closed", "all"),
			),
			mcp.WithString("head",
				mcp.Description("Filter by head user/org and branch"),
			),
			mcp.WithString("base",
				mcp.Description("Filter by base branch"),
			),
			mcp.WithString("sort",
				mcp.Description("Sort by"),
				mcp.Enum("created", "updated", "popularity", "long-running"),
			),
			mcp.WithString("direction",
				mcp.Description("Sort direction"),
				mcp.Enum("asc", "desc"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			state, err := OptionalParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			head, err := OptionalParam[string](request, "head")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			base, err := OptionalParam[string](request, "base")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sort, err := OptionalParam[string](request, "sort")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			direction, err := OptionalParam[string](request, "direction")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			opts := &github.PullRequestListOptions{
				State:     state,
				Head:      head,
				Base:      base,
				Sort:      sort,
				Direction: direction,
				ListOptions: github.ListOptions{
					PerPage: pagination.PerPage,
					Page:    pagination.Page,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			prs, resp, err := client.PullRequests.List(ctx, owner, repo, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to list pull requests",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list pull requests: %s", string(body))), nil
			}

			r, err := json.Marshal(prs)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// MergePullRequest creates a tool to merge a pull request.
func MergePullRequest(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("merge_pull_request",
			mcp.WithDescription(t("TOOL_MERGE_PULL_REQUEST_DESCRIPTION", "Merge a pull request in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_MERGE_PULL_REQUEST_USER_TITLE", "Merge pull request"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("commit_title",
				mcp.Description("Title for merge commit"),
			),
			mcp.WithString("commit_message",
				mcp.Description("Extra detail for merge commit"),
			),
			mcp.WithString("merge_method",
				mcp.Description("Merge method"),
				mcp.Enum("merge", "squash", "rebase"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			commitTitle, err := OptionalParam[string](request, "commit_title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			commitMessage, err := OptionalParam[string](request, "commit_message")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			mergeMethod, err := OptionalParam[string](request, "merge_method")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			options := &github.PullRequestOptions{
				CommitTitle: commitTitle,
				MergeMethod: mergeMethod,
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			result, resp, err := client.PullRequests.Merge(ctx, owner, repo, pullNumber, commitMessage, options)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to merge pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to merge pull request: %s", string(body))), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// SearchPullRequests creates a tool to search for pull requests.
func SearchPullRequests(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_pull_requests",
			mcp.WithDescription(t("TOOL_SEARCH_PULL_REQUESTS_DESCRIPTION", "Search for pull requests in GitHub repositories using issues search syntax already scoped to is:pr")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SEARCH_PULL_REQUESTS_USER_TITLE", "Search pull requests"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("Search query using GitHub pull request search syntax"),
			),
			mcp.WithString("owner",
				mcp.Description("Optional repository owner. If provided with repo, only notifications for this repository are listed."),
			),
			mcp.WithString("repo",
				mcp.Description("Optional repository name. If provided with owner, only notifications for this repository are listed."),
			),
			mcp.WithString("sort",
				mcp.Description("Sort field by number of matches of categories, defaults to best match"),
				mcp.Enum(
					"comments",
					"reactions",
					"reactions-+1",
					"reactions--1",
					"reactions-smile",
					"reactions-thinking_face",
					"reactions-heart",
					"reactions-tada",
					"interactions",
					"created",
					"updated",
				),
			),
			mcp.WithString("order",
				mcp.Description("Sort order"),
				mcp.Enum("asc", "desc"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return searchHandler(ctx, getClient, request, "pr", "failed to search pull requests")
		}
}

// GetPullRequestFiles creates a tool to get the list of files changed in a pull request.
func GetPullRequestFiles(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_files",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_FILES_DESCRIPTION", "Get the files changed in a specific pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_FILES_USER_TITLE", "Get pull request files"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			opts := &github.ListOptions{
				PerPage: pagination.PerPage,
				Page:    pagination.Page,
			}
			files, resp, err := client.PullRequests.ListFiles(ctx, owner, repo, pullNumber, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request files",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request files: %s", string(body))), nil
			}

			r, err := json.Marshal(files)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetPullRequestStatus creates a tool to get the combined status of all status checks for a pull request.
func GetPullRequestStatus(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_status",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_STATUS_DESCRIPTION", "Get the status of a specific pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_STATUS_USER_TITLE", "Get pull request status checks"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			// First get the PR to find the head SHA
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			pr, resp, err := client.PullRequests.Get(ctx, owner, repo, pullNumber)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request: %s", string(body))), nil
			}

			// Get combined status for the head SHA
			status, resp, err := client.Repositories.GetCombinedStatus(ctx, owner, repo, *pr.Head.SHA, nil)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get combined status",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get combined status: %s", string(body))), nil
			}

			r, err := json.Marshal(status)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// UpdatePullRequestBranch creates a tool to update a pull request branch with the latest changes from the base branch.
func UpdatePullRequestBranch(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("update_pull_request_branch",
			mcp.WithDescription(t("TOOL_UPDATE_PULL_REQUEST_BRANCH_DESCRIPTION", "Update the branch of a pull request with the latest changes from the base branch.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_UPDATE_PULL_REQUEST_BRANCH_USER_TITLE", "Update pull request branch"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("expectedHeadSha",
				mcp.Description("The expected SHA of the pull request's HEAD ref"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			expectedHeadSHA, err := OptionalParam[string](request, "expectedHeadSha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			opts := &github.PullRequestBranchUpdateOptions{}
			if expectedHeadSHA != "" {
				opts.ExpectedHeadSHA = github.Ptr(expectedHeadSHA)
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			result, resp, err := client.PullRequests.UpdateBranch(ctx, owner, repo, pullNumber, opts)
			if err != nil {
				// Check if it's an acceptedError. An acceptedError indicates that the update is in progress,
				// and it's not a real error.
				if resp != nil && resp.StatusCode == http.StatusAccepted && isAcceptedError(err) {
					return mcp.NewToolResultText("Pull request branch update is in progress"), nil
				}
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to update pull request branch",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusAccepted {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to update pull request branch: %s", string(body))), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetPullRequestComments creates a tool to get the review comments on a pull request.
func GetPullRequestComments(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_comments",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_COMMENTS_DESCRIPTION", "Get comments for a specific pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_COMMENTS_USER_TITLE", "Get pull request comments"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.PullRequestListCommentsOptions{
				ListOptions: github.ListOptions{
					PerPage: 100,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			comments, resp, err := client.PullRequests.ListComments(ctx, owner, repo, pullNumber, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request comments",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request comments: %s", string(body))), nil
			}

			r, err := json.Marshal(comments)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetPullRequestReviews creates a tool to get the reviews on a pull request.
func GetPullRequestReviews(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_reviews",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_REVIEWS_DESCRIPTION", "Get reviews for a specific pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_REVIEWS_USER_TITLE", "Get pull request reviews"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			reviews, resp, err := client.PullRequests.ListReviews(ctx, owner, repo, pullNumber, nil)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request reviews",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request reviews: %s", string(body))), nil
			}

			r, err := json.Marshal(reviews)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

func CreateAndSubmitPullRequestReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("create_and_submit_pull_request_review",
			mcp.WithDescription(t("TOOL_CREATE_AND_SUBMIT_PULL_REQUEST_REVIEW_DESCRIPTION", "Create and submit a review for a pull request without review comments.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_AND_SUBMIT_PULL_REQUEST_REVIEW_USER_TITLE", "Create and submit a pull request review without comments"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Either we need the PR GQL Id directly, or we need owner, repo and PR number to look it up.
			// Since our other Pull Request tools are working with the REST Client, will handle the lookup
			// internally for now.
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("body",
				mcp.Required(),
				mcp.Description("Review comment text"),
			),
			mcp.WithString("event",
				mcp.Required(),
				mcp.Description("Review action to perform"),
				mcp.Enum("APPROVE", "REQUEST_CHANGES", "COMMENT"),
			),
			mcp.WithString("commitID",
				mcp.Description("SHA of commit to review"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
				Body       string
				Event      string
				CommitID   *string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Given our owner, repo and PR number, lookup the GQL ID of the PR.
			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var getPullRequestQuery struct {
				Repository struct {
					PullRequest struct {
						ID githubv4.ID
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			if err := client.Query(ctx, &getPullRequestQuery, map[string]any{
				"owner": githubv4.String(params.Owner),
				"repo":  githubv4.String(params.Repo),
				"prNum": githubv4.Int(params.PullNumber),
			}); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get pull request",
					err,
				), nil
			}

			// Now we have the GQL ID, we can create a review
			var addPullRequestReviewMutation struct {
				AddPullRequestReview struct {
					PullRequestReview struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"addPullRequestReview(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&addPullRequestReviewMutation,
				githubv4.AddPullRequestReviewInput{
					PullRequestID: getPullRequestQuery.Repository.PullRequest.ID,
					Body:          githubv4.NewString(githubv4.String(params.Body)),
					Event:         newGQLStringlike[githubv4.PullRequestReviewEvent](params.Event),
					CommitOID:     newGQLStringlikePtr[githubv4.GitObjectID](params.CommitID),
				},
				nil,
			); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pull request review submitted successfully"), nil
		}
}

// CreatePendingPullRequestReview creates a tool to create a pending review on a pull request.
func CreatePendingPullRequestReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("create_pending_pull_request_review",
			mcp.WithDescription(t("TOOL_CREATE_PENDING_PULL_REQUEST_REVIEW_DESCRIPTION", "Create a pending review for a pull request. Call this first before attempting to add comments to a pending review, and ultimately submitting it. A pending pull request review means a pull request review, it is pending because you create it first and submit it later, and the PR author will not see it until it is submitted.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_PENDING_PULL_REQUEST_REVIEW_USER_TITLE", "Create pending pull request review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Either we need the PR GQL Id directly, or we need owner, repo and PR number to look it up.
			// Since our other Pull Request tools are working with the REST Client, will handle the lookup
			// internally for now.
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("commitID",
				mcp.Description("SHA of commit to review"),
			),
			// Event is omitted here because we always want to create a pending review.
			// Threads are omitted for the moment, and we'll see if the LLM can use the appropriate tool.
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
				CommitID   *string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Given our owner, repo and PR number, lookup the GQL ID of the PR.
			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var getPullRequestQuery struct {
				Repository struct {
					PullRequest struct {
						ID githubv4.ID
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			if err := client.Query(ctx, &getPullRequestQuery, map[string]any{
				"owner": githubv4.String(params.Owner),
				"repo":  githubv4.String(params.Repo),
				"prNum": githubv4.Int(params.PullNumber),
			}); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get pull request",
					err,
				), nil
			}

			// Now we have the GQL ID, we can create a pending review
			var addPullRequestReviewMutation struct {
				AddPullRequestReview struct {
					PullRequestReview struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"addPullRequestReview(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&addPullRequestReviewMutation,
				githubv4.AddPullRequestReviewInput{
					PullRequestID: getPullRequestQuery.Repository.PullRequest.ID,
					CommitOID:     newGQLStringlikePtr[githubv4.GitObjectID](params.CommitID),
				},
				nil,
			); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pending pull request created"), nil
		}
}

// AddCommentToPendingReview creates a tool to add a comment to a pull request review.
func AddCommentToPendingReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("add_comment_to_pending_review",
			mcp.WithDescription(t("TOOL_ADD_COMMENT_TO_PENDING_REVIEW_DESCRIPTION", "Add review comment to the requester's latest pending pull request review. A pending review needs to already exist to call this (check with the user if not sure).")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_ADD_COMMENT_TO_PENDING_REVIEW_USER_TITLE", "Add review comment to the requester's latest pending pull request review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Ideally, for performance sake this would just accept the pullRequestReviewID. However, we would need to
			// add a new tool to get that ID for clients that aren't in the same context as the original pending review
			// creation. So for now, we'll just accept the owner, repo and pull number and assume this is adding a comment
			// the latest review from a user, since only one can be active at a time. It can later be extended with
			// a pullRequestReviewID parameter if targeting other reviews is desired:
			// mcp.WithString("pullRequestReviewID",
			// 	mcp.Required(),
			// 	mcp.Description("The ID of the pull request review to add a comment to"),
			// ),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("The relative path to the file that necessitates a comment"),
			),
			mcp.WithString("body",
				mcp.Required(),
				mcp.Description("The text of the review comment"),
			),
			mcp.WithString("subjectType",
				mcp.Required(),
				mcp.Description("The level at which the comment is targeted"),
				mcp.Enum("FILE", "LINE"),
			),
			mcp.WithNumber("line",
				mcp.Description("The line of the blob in the pull request diff that the comment applies to. For multi-line comments, the last line of the range"),
			),
			mcp.WithString("side",
				mcp.Description("The side of the diff to comment on. LEFT indicates the previous state, RIGHT indicates the new state"),
				mcp.Enum("LEFT", "RIGHT"),
			),
			mcp.WithNumber("startLine",
				mcp.Description("For multi-line comments, the first line of the range that the comment applies to"),
			),
			mcp.WithString("startSide",
				mcp.Description("For multi-line comments, the starting side of the diff that the comment applies to. LEFT indicates the previous state, RIGHT indicates the new state"),
				mcp.Enum("LEFT", "RIGHT"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner       string
				Repo        string
				PullNumber  int32
				Path        string
				Body        string
				SubjectType string
				Line        *int32
				Side        *string
				StartLine   *int32
				StartSide   *string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub GQL client: %w", err)
			}

			// First we'll get the current user
			var getViewerQuery struct {
				Viewer struct {
					Login githubv4.String
				}
			}

			if err := client.Query(ctx, &getViewerQuery, nil); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get current user",
					err,
				), nil
			}

			var getLatestReviewForViewerQuery struct {
				Repository struct {
					PullRequest struct {
						Reviews struct {
							Nodes []struct {
								ID    githubv4.ID
								State githubv4.PullRequestReviewState
								URL   githubv4.URI
							}
						} `graphql:"reviews(first: 1, author: $author)"`
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			vars := map[string]any{
				"author": githubv4.String(getViewerQuery.Viewer.Login),
				"owner":  githubv4.String(params.Owner),
				"name":   githubv4.String(params.Repo),
				"prNum":  githubv4.Int(params.PullNumber),
			}

			if err := client.Query(context.Background(), &getLatestReviewForViewerQuery, vars); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get latest review for current user",
					err,
				), nil
			}

			// Validate there is one review and the state is pending
			if len(getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes) == 0 {
				return mcp.NewToolResultError("No pending review found for the viewer"), nil
			}

			review := getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes[0]
			if review.State != githubv4.PullRequestReviewStatePending {
				errText := fmt.Sprintf("The latest review, found at %s is not pending", review.URL)
				return mcp.NewToolResultError(errText), nil
			}

			// Then we can create a new review thread comment on the review.
			var addPullRequestReviewThreadMutation struct {
				AddPullRequestReviewThread struct {
					Thread struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"addPullRequestReviewThread(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&addPullRequestReviewThreadMutation,
				githubv4.AddPullRequestReviewThreadInput{
					Path:                githubv4.String(params.Path),
					Body:                githubv4.String(params.Body),
					SubjectType:         newGQLStringlikePtr[githubv4.PullRequestReviewThreadSubjectType](&params.SubjectType),
					Line:                newGQLIntPtr(params.Line),
					Side:                newGQLStringlikePtr[githubv4.DiffSide](params.Side),
					StartLine:           newGQLIntPtr(params.StartLine),
					StartSide:           newGQLStringlikePtr[githubv4.DiffSide](params.StartSide),
					PullRequestReviewID: &review.ID,
				},
				nil,
			); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pull request review comment successfully added to pending review"), nil
		}
}

// SubmitPendingPullRequestReview creates a tool to submit a pull request review.
func SubmitPendingPullRequestReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("submit_pending_pull_request_review",
			mcp.WithDescription(t("TOOL_SUBMIT_PENDING_PULL_REQUEST_REVIEW_DESCRIPTION", "Submit the requester's latest pending pull request review, normally this is a final step after creating a pending review, adding comments first, unless you know that the user already did the first two steps, you should check before calling this.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SUBMIT_PENDING_PULL_REQUEST_REVIEW_USER_TITLE", "Submit the requester's latest pending pull request review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Ideally, for performance sake this would just accept the pullRequestReviewID. However, we would need to
			// add a new tool to get that ID for clients that aren't in the same context as the original pending review
			// creation. So for now, we'll just accept the owner, repo and pull number and assume this is submitting
			// the latest review from a user, since only one can be active at a time.
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
			mcp.WithString("event",
				mcp.Required(),
				mcp.Description("The event to perform"),
				mcp.Enum("APPROVE", "REQUEST_CHANGES", "COMMENT"),
			),
			mcp.WithString("body",
				mcp.Description("The text of the review comment"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
				Event      string
				Body       *string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub GQL client: %w", err)
			}

			// First we'll get the current user
			var getViewerQuery struct {
				Viewer struct {
					Login githubv4.String
				}
			}

			if err := client.Query(ctx, &getViewerQuery, nil); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get current user",
					err,
				), nil
			}

			var getLatestReviewForViewerQuery struct {
				Repository struct {
					PullRequest struct {
						Reviews struct {
							Nodes []struct {
								ID    githubv4.ID
								State githubv4.PullRequestReviewState
								URL   githubv4.URI
							}
						} `graphql:"reviews(first: 1, author: $author)"`
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			vars := map[string]any{
				"author": githubv4.String(getViewerQuery.Viewer.Login),
				"owner":  githubv4.String(params.Owner),
				"name":   githubv4.String(params.Repo),
				"prNum":  githubv4.Int(params.PullNumber),
			}

			if err := client.Query(context.Background(), &getLatestReviewForViewerQuery, vars); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get latest review for current user",
					err,
				), nil
			}

			// Validate there is one review and the state is pending
			if len(getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes) == 0 {
				return mcp.NewToolResultError("No pending review found for the viewer"), nil
			}

			review := getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes[0]
			if review.State != githubv4.PullRequestReviewStatePending {
				errText := fmt.Sprintf("The latest review, found at %s is not pending", review.URL)
				return mcp.NewToolResultError(errText), nil
			}

			// Prepare the mutation
			var submitPullRequestReviewMutation struct {
				SubmitPullRequestReview struct {
					PullRequestReview struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"submitPullRequestReview(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&submitPullRequestReviewMutation,
				githubv4.SubmitPullRequestReviewInput{
					PullRequestReviewID: &review.ID,
					Event:               githubv4.PullRequestReviewEvent(params.Event),
					Body:                newGQLStringlikePtr[githubv4.String](params.Body),
				},
				nil,
			); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to submit pull request review",
					err,
				), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pending pull request review successfully submitted"), nil
		}
}

func DeletePendingPullRequestReview(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("delete_pending_pull_request_review",
			mcp.WithDescription(t("TOOL_DELETE_PENDING_PULL_REQUEST_REVIEW_DESCRIPTION", "Delete the requester's latest pending pull request review. Use this after the user decides not to submit a pending review, if you don't know if they already created one then check first.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_DELETE_PENDING_PULL_REQUEST_REVIEW_USER_TITLE", "Delete the requester's latest pending pull request review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			// Ideally, for performance sake this would just accept the pullRequestReviewID. However, we would need to
			// add a new tool to get that ID for clients that aren't in the same context as the original pending review
			// creation. So for now, we'll just accept the owner, repo and pull number and assume this is deleting
			// the latest pending review from a user, since only one can be active at a time.
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub GQL client: %w", err)
			}

			// First we'll get the current user
			var getViewerQuery struct {
				Viewer struct {
					Login githubv4.String
				}
			}

			if err := client.Query(ctx, &getViewerQuery, nil); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get current user",
					err,
				), nil
			}

			var getLatestReviewForViewerQuery struct {
				Repository struct {
					PullRequest struct {
						Reviews struct {
							Nodes []struct {
								ID    githubv4.ID
								State githubv4.PullRequestReviewState
								URL   githubv4.URI
							}
						} `graphql:"reviews(first: 1, author: $author)"`
					} `graphql:"pullRequest(number: $prNum)"`
				} `graphql:"repository(owner: $owner, name: $name)"`
			}

			vars := map[string]any{
				"author": githubv4.String(getViewerQuery.Viewer.Login),
				"owner":  githubv4.String(params.Owner),
				"name":   githubv4.String(params.Repo),
				"prNum":  githubv4.Int(params.PullNumber),
			}

			if err := client.Query(context.Background(), &getLatestReviewForViewerQuery, vars); err != nil {
				return ghErrors.NewGitHubGraphQLErrorResponse(ctx,
					"failed to get latest review for current user",
					err,
				), nil
			}

			// Validate there is one review and the state is pending
			if len(getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes) == 0 {
				return mcp.NewToolResultError("No pending review found for the viewer"), nil
			}

			review := getLatestReviewForViewerQuery.Repository.PullRequest.Reviews.Nodes[0]
			if review.State != githubv4.PullRequestReviewStatePending {
				errText := fmt.Sprintf("The latest review, found at %s is not pending", review.URL)
				return mcp.NewToolResultError(errText), nil
			}

			// Prepare the mutation
			var deletePullRequestReviewMutation struct {
				DeletePullRequestReview struct {
					PullRequestReview struct {
						ID githubv4.ID // We don't need this, but a selector is required or GQL complains.
					}
				} `graphql:"deletePullRequestReview(input: $input)"`
			}

			if err := client.Mutate(
				ctx,
				&deletePullRequestReviewMutation,
				githubv4.DeletePullRequestReviewInput{
					PullRequestReviewID: &review.ID,
				},
				nil,
			); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Return nothing interesting, just indicate success for the time being.
			// In future, we may want to return the review ID, but for the moment, we're not leaking
			// API implementation details to the LLM.
			return mcp.NewToolResultText("pending pull request review successfully deleted"), nil
		}
}

func GetPullRequestDiff(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("get_pull_request_diff",
			mcp.WithDescription(t("TOOL_GET_PULL_REQUEST_DIFF_DESCRIPTION", "Get the diff of a pull request.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PULL_REQUEST_DIFF_USER_TITLE", "Get pull request diff"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var params struct {
				Owner      string
				Repo       string
				PullNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub client: %v", err)), nil
			}

			raw, resp, err := client.PullRequests.GetRaw(
				ctx,
				params.Owner,
				params.Repo,
				int(params.PullNumber),
				github.RawOptions{Type: github.Diff},
			)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get pull request diff",
					resp,
					err,
				), nil
			}

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get pull request diff: %s", string(body))), nil
			}

			defer func() { _ = resp.Body.Close() }()

			// Return the raw response
			return mcp.NewToolResultText(string(raw)), nil
		}
}

// RequestCopilotReview creates a tool to request a Copilot review for a pull request.
// Note that this tool will not work on GHES where this feature is unsupported. In future, we should not expose this
// tool if the configured host does not support it.
func RequestCopilotReview(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool("request_copilot_review",
			mcp.WithDescription(t("TOOL_REQUEST_COPILOT_REVIEW_DESCRIPTION", "Request a GitHub Copilot code review for a pull request. Use this for automated feedback on pull requests, usually before requesting a human reviewer.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_REQUEST_COPILOT_REVIEW_USER_TITLE", "Request Copilot review"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("pullNumber",
				mcp.Required(),
				mcp.Description("Pull request number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			pullNumber, err := RequiredInt(request, "pullNumber")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			_, resp, err := client.PullRequests.RequestReviewers(
				ctx,
				owner,
				repo,
				pullNumber,
				github.ReviewersRequest{
					// The login name of the copilot reviewer bot
					Reviewers: []string{"copilot-pull-request-reviewer[bot]"},
				},
			)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to request copilot review",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to request copilot review: %s", string(body))), nil
			}

			// Return nothing on success, as there's not much value in returning the Pull Request itself
			return mcp.NewToolResultText(""), nil
		}
}

// newGQLString like takes something that approximates a string (of which there are many types in shurcooL/githubv4)
// and constructs a pointer to it, or nil if the string is empty. This is extremely useful because when we parse
// params from the MCP request, we need to convert them to types that are pointers of type def strings and it's
// not possible to take a pointer of an anonymous value e.g. &githubv4.String("foo").
func newGQLStringlike[T ~string](s string) *T {
	if s == "" {
		return nil
	}
	stringlike := T(s)
	return &stringlike
}

func newGQLStringlikePtr[T ~string](s *string) *T {
	if s == nil {
		return nil
	}
	stringlike := T(*s)
	return &stringlike
}

func newGQLIntPtr(i *int32) *githubv4.Int {
	if i == nil {
		return nil
	}
	gi := githubv4.Int(*i)
	return &gi
}

```

#### y/repositories.go

```go
package github

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/raw"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func GetCommit(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_commit",
			mcp.WithDescription(t("TOOL_GET_COMMITS_DESCRIPTION", "Get details for a commit from a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_COMMITS_USER_TITLE", "Get commit details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("sha",
				mcp.Required(),
				mcp.Description("Commit SHA, branch name, or tag name"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sha, err := RequiredParam[string](request, "sha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.ListOptions{
				Page:    pagination.Page,
				PerPage: pagination.PerPage,
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			commit, resp, err := client.Repositories.GetCommit(ctx, owner, repo, sha, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to get commit: %s", sha),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get commit: %s", string(body))), nil
			}

			r, err := json.Marshal(commit)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListCommits creates a tool to get commits of a branch in a repository.
func ListCommits(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_commits",
			mcp.WithDescription(t("TOOL_LIST_COMMITS_DESCRIPTION", "Get list of commits of a branch in a GitHub repository. Returns at least 30 results per page by default, but can return more if specified using the perPage parameter (up to 100).")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_COMMITS_USER_TITLE", "List commits"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("sha",
				mcp.Description("Commit SHA, branch or tag name to list commits of. If not provided, uses the default branch of the repository. If a commit SHA is provided, will list commits up to that SHA."),
			),
			mcp.WithString("author",
				mcp.Description("Author username or email address to filter commits by"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sha, err := OptionalParam[string](request, "sha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			author, err := OptionalParam[string](request, "author")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			// Set default perPage to 30 if not provided
			perPage := pagination.PerPage
			if perPage == 0 {
				perPage = 30
			}
			opts := &github.CommitsListOptions{
				SHA:    sha,
				Author: author,
				ListOptions: github.ListOptions{
					Page:    pagination.Page,
					PerPage: perPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			commits, resp, err := client.Repositories.ListCommits(ctx, owner, repo, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to list commits: %s", sha),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list commits: %s", string(body))), nil
			}

			r, err := json.Marshal(commits)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListBranches creates a tool to list branches in a GitHub repository.
func ListBranches(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_branches",
			mcp.WithDescription(t("TOOL_LIST_BRANCHES_DESCRIPTION", "List branches in a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_BRANCHES_USER_TITLE", "List branches"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.BranchListOptions{
				ListOptions: github.ListOptions{
					Page:    pagination.Page,
					PerPage: pagination.PerPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			branches, resp, err := client.Repositories.ListBranches(ctx, owner, repo, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to list branches",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list branches: %s", string(body))), nil
			}

			r, err := json.Marshal(branches)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CreateOrUpdateFile creates a tool to create or update a file in a GitHub repository.
func CreateOrUpdateFile(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("create_or_update_file",
			mcp.WithDescription(t("TOOL_CREATE_OR_UPDATE_FILE_DESCRIPTION", "Create or update a single file in a GitHub repository. If updating, you must provide the SHA of the file you want to update. Use this tool to create or update a file in a GitHub repository remotely; do not use it for local file operations.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_OR_UPDATE_FILE_USER_TITLE", "Create or update file"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner (username or organization)"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("Path where to create/update the file"),
			),
			mcp.WithString("content",
				mcp.Required(),
				mcp.Description("Content of the file"),
			),
			mcp.WithString("message",
				mcp.Required(),
				mcp.Description("Commit message"),
			),
			mcp.WithString("branch",
				mcp.Required(),
				mcp.Description("Branch to create/update the file in"),
			),
			mcp.WithString("sha",
				mcp.Description("Required if updating an existing file. The blob SHA of the file being replaced."),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			path, err := RequiredParam[string](request, "path")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			content, err := RequiredParam[string](request, "content")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			message, err := RequiredParam[string](request, "message")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := RequiredParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// json.Marshal encodes byte arrays with base64, which is required for the API.
			contentBytes := []byte(content)

			// Create the file options
			opts := &github.RepositoryContentFileOptions{
				Message: github.Ptr(message),
				Content: contentBytes,
				Branch:  github.Ptr(branch),
			}

			// If SHA is provided, set it (for updates)
			sha, err := OptionalParam[string](request, "sha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if sha != "" {
				opts.SHA = github.Ptr(sha)
			}

			// Create or update the file
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			fileContent, resp, err := client.Repositories.CreateFile(ctx, owner, repo, path, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create/update file",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 && resp.StatusCode != 201 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create/update file: %s", string(body))), nil
			}

			r, err := json.Marshal(fileContent)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CreateRepository creates a tool to create a new GitHub repository.
func CreateRepository(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("create_repository",
			mcp.WithDescription(t("TOOL_CREATE_REPOSITORY_DESCRIPTION", "Create a new GitHub repository in your account")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_REPOSITORY_USER_TITLE", "Create repository"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("description",
				mcp.Description("Repository description"),
			),
			mcp.WithBoolean("private",
				mcp.Description("Whether repo should be private"),
			),
			mcp.WithBoolean("autoInit",
				mcp.Description("Initialize with README"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			name, err := RequiredParam[string](request, "name")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			description, err := OptionalParam[string](request, "description")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			private, err := OptionalParam[bool](request, "private")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			autoInit, err := OptionalParam[bool](request, "autoInit")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			repo := &github.Repository{
				Name:        github.Ptr(name),
				Description: github.Ptr(description),
				Private:     github.Ptr(private),
				AutoInit:    github.Ptr(autoInit),
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			createdRepo, resp, err := client.Repositories.Create(ctx, "", repo)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create repository",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create repository: %s", string(body))), nil
			}

			r, err := json.Marshal(createdRepo)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetFileContents creates a tool to get the contents of a file or directory from a GitHub repository.
func GetFileContents(getClient GetClientFn, getRawClient raw.GetRawClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_file_contents",
			mcp.WithDescription(t("TOOL_GET_FILE_CONTENTS_DESCRIPTION", "Get the contents of a file or directory from a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_FILE_CONTENTS_USER_TITLE", "Get file or directory contents"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner (username or organization)"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("path",
				mcp.Description("Path to file/directory (directories must end with a slash '/')"),
				mcp.DefaultString("/"),
			),
			mcp.WithString("ref",
				mcp.Description("Accepts optional git refs such as `refs/tags/{tag}`, `refs/heads/{branch}` or `refs/pull/{pr_number}/head`"),
			),
			mcp.WithString("sha",
				mcp.Description("Accepts optional commit SHA. If specified, it will be used instead of ref"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			path, err := RequiredParam[string](request, "path")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			ref, err := OptionalParam[string](request, "ref")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sha, err := OptionalParam[string](request, "sha")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return mcp.NewToolResultError("failed to get GitHub client"), nil
			}

			rawOpts, err := resolveGitReference(ctx, client, owner, repo, ref, sha)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to resolve git reference: %s", err)), nil
			}

			// If the path is (most likely) not to be a directory, we will
			// first try to get the raw content from the GitHub raw content API.
			if path != "" && !strings.HasSuffix(path, "/") {
				// First, get file info from Contents API to retrieve SHA
				var fileSHA string
				opts := &github.RepositoryContentGetOptions{Ref: ref}
				fileContent, _, respContents, err := client.Repositories.GetContents(ctx, owner, repo, path, opts)
				if respContents != nil {
					defer func() { _ = respContents.Body.Close() }()
				}
				if err != nil {
					return ghErrors.NewGitHubAPIErrorResponse(ctx,
						"failed to get file SHA",
						respContents,
						err,
					), nil
				}
				if fileContent == nil || fileContent.SHA == nil {
					return mcp.NewToolResultError("file content SHA is nil"), nil
				}
				fileSHA = *fileContent.SHA

				rawClient, err := getRawClient(ctx)
				if err != nil {
					return mcp.NewToolResultError("failed to get GitHub raw content client"), nil
				}
				resp, err := rawClient.GetRawContent(ctx, owner, repo, path, rawOpts)
				if err != nil {
					return mcp.NewToolResultError("failed to get raw repository content"), nil
				}
				defer func() {
					_ = resp.Body.Close()
				}()

				if resp.StatusCode == http.StatusOK {
					// If the raw content is found, return it directly
					body, err := io.ReadAll(resp.Body)
					if err != nil {
						return mcp.NewToolResultError("failed to read response body"), nil
					}
					contentType := resp.Header.Get("Content-Type")

					var resourceURI string
					switch {
					case sha != "":
						resourceURI, err = url.JoinPath("repo://", owner, repo, "sha", sha, "contents", path)
						if err != nil {
							return nil, fmt.Errorf("failed to create resource URI: %w", err)
						}
					case ref != "":
						resourceURI, err = url.JoinPath("repo://", owner, repo, ref, "contents", path)
						if err != nil {
							return nil, fmt.Errorf("failed to create resource URI: %w", err)
						}
					default:
						resourceURI, err = url.JoinPath("repo://", owner, repo, "contents", path)
						if err != nil {
							return nil, fmt.Errorf("failed to create resource URI: %w", err)
						}
					}

					if strings.HasPrefix(contentType, "application") || strings.HasPrefix(contentType, "text") {
						result := mcp.TextResourceContents{
							URI:      resourceURI,
							Text:     string(body),
							MIMEType: contentType,
						}
						// Include SHA in the result metadata
						if fileSHA != "" {
							return mcp.NewToolResultResource(fmt.Sprintf("successfully downloaded text file (SHA: %s)", fileSHA), result), nil
						}
						return mcp.NewToolResultResource("successfully downloaded text file", result), nil
					}

					result := mcp.BlobResourceContents{
						URI:      resourceURI,
						Blob:     base64.StdEncoding.EncodeToString(body),
						MIMEType: contentType,
					}
					// Include SHA in the result metadata
					if fileSHA != "" {
						return mcp.NewToolResultResource(fmt.Sprintf("successfully downloaded binary file (SHA: %s)", fileSHA), result), nil
					}
					return mcp.NewToolResultResource("successfully downloaded binary file", result), nil

				}
			}

			if rawOpts.SHA != "" {
				ref = rawOpts.SHA
			}
			if strings.HasSuffix(path, "/") {
				opts := &github.RepositoryContentGetOptions{Ref: ref}
				_, dirContent, resp, err := client.Repositories.GetContents(ctx, owner, repo, path, opts)
				if err == nil && resp.StatusCode == http.StatusOK {
					defer func() { _ = resp.Body.Close() }()
					r, err := json.Marshal(dirContent)
					if err != nil {
						return mcp.NewToolResultError("failed to marshal response"), nil
					}
					return mcp.NewToolResultText(string(r)), nil
				}
			}

			// The path does not point to a file or directory.
			// Instead let's try to find it in the Git Tree by matching the end of the path.

			// Step 1: Get Git Tree recursively
			tree, resp, err := client.Git.GetTree(ctx, owner, repo, ref, true)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get git tree",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Step 2: Filter tree for matching paths
			const maxMatchingFiles = 3
			matchingFiles := filterPaths(tree.Entries, path, maxMatchingFiles)
			if len(matchingFiles) > 0 {
				matchingFilesJSON, err := json.Marshal(matchingFiles)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("failed to marshal matching files: %s", err)), nil
				}
				resolvedRefs, err := json.Marshal(rawOpts)
				if err != nil {
					return mcp.NewToolResultError(fmt.Sprintf("failed to marshal resolved refs: %s", err)), nil
				}
				return mcp.NewToolResultText(fmt.Sprintf("Path did not point to a file or directory, but resolved git ref to %s with possible path matches: %s", resolvedRefs, matchingFilesJSON)), nil
			}

			return mcp.NewToolResultError("Failed to get file contents. The path does not point to a file or directory, or the file does not exist in the repository."), nil
		}
}

// ForkRepository creates a tool to fork a repository.
func ForkRepository(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("fork_repository",
			mcp.WithDescription(t("TOOL_FORK_REPOSITORY_DESCRIPTION", "Fork a GitHub repository to your account or specified organization")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_FORK_REPOSITORY_USER_TITLE", "Fork repository"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("organization",
				mcp.Description("Organization to fork to"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			org, err := OptionalParam[string](request, "organization")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.RepositoryCreateForkOptions{}
			if org != "" {
				opts.Organization = org
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			forkedRepo, resp, err := client.Repositories.CreateFork(ctx, owner, repo, opts)
			if err != nil {
				// Check if it's an acceptedError. An acceptedError indicates that the update is in progress,
				// and it's not a real error.
				if resp != nil && resp.StatusCode == http.StatusAccepted && isAcceptedError(err) {
					return mcp.NewToolResultText("Fork is in progress"), nil
				}
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to fork repository",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusAccepted {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to fork repository: %s", string(body))), nil
			}

			r, err := json.Marshal(forkedRepo)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// DeleteFile creates a tool to delete a file in a GitHub repository.
// This tool uses a more roundabout way of deleting a file than just using the client.Repositories.DeleteFile.
// This is because REST file deletion endpoint (and client.Repositories.DeleteFile) don't add commit signing to the deletion commit,
// unlike how the endpoint backing the create_or_update_files tool does. This appears to be a quirk of the API.
// The approach implemented here gets automatic commit signing when used with either the github-actions user or as an app,
// both of which suit an LLM well.
func DeleteFile(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("delete_file",
			mcp.WithDescription(t("TOOL_DELETE_FILE_DESCRIPTION", "Delete a file from a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:           t("TOOL_DELETE_FILE_USER_TITLE", "Delete file"),
				ReadOnlyHint:    ToBoolPtr(false),
				DestructiveHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner (username or organization)"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("Path to the file to delete"),
			),
			mcp.WithString("message",
				mcp.Required(),
				mcp.Description("Commit message"),
			),
			mcp.WithString("branch",
				mcp.Required(),
				mcp.Description("Branch to delete the file from"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			path, err := RequiredParam[string](request, "path")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			message, err := RequiredParam[string](request, "message")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := RequiredParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the reference for the branch
			ref, resp, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch)
			if err != nil {
				return nil, fmt.Errorf("failed to get branch reference: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			// Get the commit object that the branch points to
			baseCommit, resp, err := client.Git.GetCommit(ctx, owner, repo, *ref.Object.SHA)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get base commit",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get commit: %s", string(body))), nil
			}

			// Create a tree entry for the file deletion by setting SHA to nil
			treeEntries := []*github.TreeEntry{
				{
					Path: github.Ptr(path),
					Mode: github.Ptr("100644"), // Regular file mode
					Type: github.Ptr("blob"),
					SHA:  nil, // Setting SHA to nil deletes the file
				},
			}

			// Create a new tree with the deletion
			newTree, resp, err := client.Git.CreateTree(ctx, owner, repo, *baseCommit.Tree.SHA, treeEntries)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create tree",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create tree: %s", string(body))), nil
			}

			// Create a new commit with the new tree
			commit := &github.Commit{
				Message: github.Ptr(message),
				Tree:    newTree,
				Parents: []*github.Commit{{SHA: baseCommit.SHA}},
			}
			newCommit, resp, err := client.Git.CreateCommit(ctx, owner, repo, commit, nil)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create commit",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to create commit: %s", string(body))), nil
			}

			// Update the branch reference to point to the new commit
			ref.Object.SHA = newCommit.SHA
			_, resp, err = client.Git.UpdateRef(ctx, owner, repo, ref, false)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to update reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to update reference: %s", string(body))), nil
			}

			// Create a response similar to what the DeleteFile API would return
			response := map[string]interface{}{
				"commit":  newCommit,
				"content": nil,
			}

			r, err := json.Marshal(response)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// CreateBranch creates a tool to create a new branch.
func CreateBranch(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("create_branch",
			mcp.WithDescription(t("TOOL_CREATE_BRANCH_DESCRIPTION", "Create a new branch in a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_CREATE_BRANCH_USER_TITLE", "Create branch"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("branch",
				mcp.Required(),
				mcp.Description("Name for new branch"),
			),
			mcp.WithString("from_branch",
				mcp.Description("Source branch (defaults to repo default)"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := RequiredParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			fromBranch, err := OptionalParam[string](request, "from_branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the source branch SHA
			var ref *github.Reference

			if fromBranch == "" {
				// Get default branch if from_branch not specified
				repository, resp, err := client.Repositories.Get(ctx, owner, repo)
				if err != nil {
					return ghErrors.NewGitHubAPIErrorResponse(ctx,
						"failed to get repository",
						resp,
						err,
					), nil
				}
				defer func() { _ = resp.Body.Close() }()

				fromBranch = *repository.DefaultBranch
			}

			// Get SHA of source branch
			ref, resp, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+fromBranch)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Create new branch
			newRef := &github.Reference{
				Ref:    github.Ptr("refs/heads/" + branch),
				Object: &github.GitObject{SHA: ref.Object.SHA},
			}

			createdRef, resp, err := client.Git.CreateRef(ctx, owner, repo, newRef)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create branch",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(createdRef)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// PushFiles creates a tool to push multiple files in a single commit to a GitHub repository.
func PushFiles(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("push_files",
			mcp.WithDescription(t("TOOL_PUSH_FILES_DESCRIPTION", "Push multiple files to a GitHub repository in a single commit")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_PUSH_FILES_USER_TITLE", "Push files to repository"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("branch",
				mcp.Required(),
				mcp.Description("Branch to push to"),
			),
			mcp.WithArray("files",
				mcp.Required(),
				mcp.Items(
					map[string]interface{}{
						"type":                 "object",
						"additionalProperties": false,
						"required":             []string{"path", "content"},
						"properties": map[string]interface{}{
							"path": map[string]interface{}{
								"type":        "string",
								"description": "path to the file",
							},
							"content": map[string]interface{}{
								"type":        "string",
								"description": "file content",
							},
						},
					}),
				mcp.Description("Array of file objects to push, each object with path (string) and content (string)"),
			),
			mcp.WithString("message",
				mcp.Required(),
				mcp.Description("Commit message"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			branch, err := RequiredParam[string](request, "branch")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			message, err := RequiredParam[string](request, "message")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Parse files parameter - this should be an array of objects with path and content
			filesObj, ok := request.GetArguments()["files"].([]interface{})
			if !ok {
				return mcp.NewToolResultError("files parameter must be an array of objects with path and content"), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// Get the reference for the branch
			ref, resp, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get branch reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Get the commit object that the branch points to
			baseCommit, resp, err := client.Git.GetCommit(ctx, owner, repo, *ref.Object.SHA)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get base commit",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Create tree entries for all files
			var entries []*github.TreeEntry

			for _, file := range filesObj {
				fileMap, ok := file.(map[string]interface{})
				if !ok {
					return mcp.NewToolResultError("each file must be an object with path and content"), nil
				}

				path, ok := fileMap["path"].(string)
				if !ok || path == "" {
					return mcp.NewToolResultError("each file must have a path"), nil
				}

				content, ok := fileMap["content"].(string)
				if !ok {
					return mcp.NewToolResultError("each file must have content"), nil
				}

				// Create a tree entry for the file
				entries = append(entries, &github.TreeEntry{
					Path:    github.Ptr(path),
					Mode:    github.Ptr("100644"), // Regular file mode
					Type:    github.Ptr("blob"),
					Content: github.Ptr(content),
				})
			}

			// Create a new tree with the file entries
			newTree, resp, err := client.Git.CreateTree(ctx, owner, repo, *baseCommit.Tree.SHA, entries)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create tree",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Create a new commit
			commit := &github.Commit{
				Message: github.Ptr(message),
				Tree:    newTree,
				Parents: []*github.Commit{{SHA: baseCommit.SHA}},
			}
			newCommit, resp, err := client.Git.CreateCommit(ctx, owner, repo, commit, nil)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to create commit",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			// Update the reference to point to the new commit
			ref.Object.SHA = newCommit.SHA
			updatedRef, resp, err := client.Git.UpdateRef(ctx, owner, repo, ref, false)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to update reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			r, err := json.Marshal(updatedRef)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListTags creates a tool to list tags in a GitHub repository.
func ListTags(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_tags",
			mcp.WithDescription(t("TOOL_LIST_TAGS_DESCRIPTION", "List git tags in a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_TAGS_USER_TITLE", "List tags"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.ListOptions{
				Page:    pagination.Page,
				PerPage: pagination.PerPage,
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			tags, resp, err := client.Repositories.ListTags(ctx, owner, repo, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to list tags",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list tags: %s", string(body))), nil
			}

			r, err := json.Marshal(tags)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetTag creates a tool to get details about a specific tag in a GitHub repository.
func GetTag(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_tag",
			mcp.WithDescription(t("TOOL_GET_TAG_DESCRIPTION", "Get details about a specific git tag in a GitHub repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_TAG_USER_TITLE", "Get tag details"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("tag",
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			tag, err := RequiredParam[string](request, "tag")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// First get the tag reference
			ref, resp, err := client.Git.GetRef(ctx, owner, repo, "refs/tags/"+tag)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get tag reference",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get tag reference: %s", string(body))), nil
			}

			// Then get the tag object
			tagObj, resp, err := client.Git.GetTag(ctx, owner, repo, *ref.Object.SHA)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					"failed to get tag object",
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get tag object: %s", string(body))), nil
			}

			r, err := json.Marshal(tagObj)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// filterPaths filters the entries in a GitHub tree to find paths that
// match the given suffix.
// maxResults limits the number of results returned to first maxResults entries,
// a maxResults of -1 means no limit.
// It returns a slice of strings containing the matching paths.
// Directories are returned with a trailing slash.
func filterPaths(entries []*github.TreeEntry, path string, maxResults int) []string {
	// Remove trailing slash for matching purposes, but flag whether we
	// only want directories.
	dirOnly := false
	if strings.HasSuffix(path, "/") {
		dirOnly = true
		path = strings.TrimSuffix(path, "/")
	}

	matchedPaths := []string{}
	for _, entry := range entries {
		if len(matchedPaths) == maxResults {
			break // Limit the number of results to maxResults
		}
		if dirOnly && entry.GetType() != "tree" {
			continue // Skip non-directory entries if dirOnly is true
		}
		entryPath := entry.GetPath()
		if entryPath == "" {
			continue // Skip empty paths
		}
		if strings.HasSuffix(entryPath, path) {
			if entry.GetType() == "tree" {
				entryPath += "/" // Return directories with a trailing slash
			}
			matchedPaths = append(matchedPaths, entryPath)
		}
	}
	return matchedPaths
}

// resolveGitReference resolves git references with the following logic:
// 1. If SHA is provided, it takes precedence
// 2. If neither is provided, use the default branch as ref
// 3. Get commit SHA from the ref
// Refs can look like `refs/tags/{tag}`, `refs/heads/{branch}` or `refs/pull/{pr_number}/head`
// The function returns the resolved ref, commit SHA and any error.
func resolveGitReference(ctx context.Context, githubClient *github.Client, owner, repo, ref, sha string) (*raw.ContentOpts, error) {
	// 1. If SHA is provided, use it directly
	if sha != "" {
		return &raw.ContentOpts{Ref: "", SHA: sha}, nil
	}

	// 2. If neither provided, use the default branch as ref
	if ref == "" {
		repoInfo, resp, err := githubClient.Repositories.Get(ctx, owner, repo)
		if err != nil {
			_, _ = ghErrors.NewGitHubAPIErrorToCtx(ctx, "failed to get repository info", resp, err)
			return nil, fmt.Errorf("failed to get repository info: %w", err)
		}
		ref = fmt.Sprintf("refs/heads/%s", repoInfo.GetDefaultBranch())
	}

	// 3. Get the SHA from the ref
	reference, resp, err := githubClient.Git.GetRef(ctx, owner, repo, ref)
	if err != nil {
		_, _ = ghErrors.NewGitHubAPIErrorToCtx(ctx, "failed to get reference", resp, err)
		return nil, fmt.Errorf("failed to get reference: %w", err)
	}
	sha = reference.GetObject().GetSHA()

	// Use provided ref, or it will be empty which defaults to the default branch
	return &raw.ContentOpts{Ref: ref, SHA: sha}, nil
}

```

#### y/search.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// SearchRepositories creates a tool to search for GitHub repositories.
func SearchRepositories(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_repositories",
			mcp.WithDescription(t("TOOL_SEARCH_REPOSITORIES_DESCRIPTION", "Search for GitHub repositories")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SEARCH_REPOSITORIES_USER_TITLE", "Search repositories"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("Search query"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			query, err := RequiredParam[string](request, "query")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.SearchOptions{
				ListOptions: github.ListOptions{
					Page:    pagination.Page,
					PerPage: pagination.PerPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			result, resp, err := client.Search.Repositories(ctx, query, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to search repositories with query '%s'", query),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to search repositories: %s", string(body))), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// SearchCode creates a tool to search for code across GitHub repositories.
func SearchCode(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_code",
			mcp.WithDescription(t("TOOL_SEARCH_CODE_DESCRIPTION", "Search for code across GitHub repositories")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_SEARCH_CODE_USER_TITLE", "Search code"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("query",
				mcp.Required(),
				mcp.Description("Search query using GitHub code search syntax"),
			),
			mcp.WithString("sort",
				mcp.Description("Sort field ('indexed' only)"),
			),
			mcp.WithString("order",
				mcp.Description("Sort order"),
				mcp.Enum("asc", "desc"),
			),
			WithPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			query, err := RequiredParam[string](request, "query")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			sort, err := OptionalParam[string](request, "sort")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			order, err := OptionalParam[string](request, "order")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.SearchOptions{
				Sort:  sort,
				Order: order,
				ListOptions: github.ListOptions{
					PerPage: pagination.PerPage,
					Page:    pagination.Page,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			result, resp, err := client.Search.Code(ctx, query, opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx,
					fmt.Sprintf("failed to search code with query '%s'", query),
					resp,
					err,
				), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to search code: %s", string(body))), nil
			}

			r, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// MinimalUser is the output type for user and organization search results.
type MinimalUser struct {
	Login      string       `json:"login"`
	ID         int64        `json:"id,omitempty"`
	ProfileURL string       `json:"profile_url,omitempty"`
	AvatarURL  string       `json:"avatar_url,omitempty"`
	Details    *UserDetails `json:"details,omitempty"` // Optional field for additional user details
}

type MinimalSearchUsersResult struct {
	TotalCount        int           `json:"total_count"`
	IncompleteResults bool          `json:"incomplete_results"`
	Items             []MinimalUser `json:"items"`
}

func userOrOrgHandler(accountType string, getClient GetClientFn) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query, err := RequiredParam[string](request, "query")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		sort, err := OptionalParam[string](request, "sort")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		order, err := OptionalParam[string](request, "order")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		pagination, err := OptionalPaginationParams(request)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		opts := &github.SearchOptions{
			Sort:  sort,
			Order: order,
			ListOptions: github.ListOptions{
				PerPage: pagination.PerPage,
				Page:    pagination.Page,
			},
		}

		client, err := getClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get GitHub client: %w", err)
		}

		searchQuery := "type:" + accountType + " " + query
		result, resp, err := client.Search.Users(ctx, searchQuery, opts)
		if err != nil {
			return ghErrors.NewGitHubAPIErrorResponse(ctx,
				fmt.Sprintf("failed to search %ss with query '%s'", accountType, query),
				resp,
				err,
			), nil
		}
		defer func() { _ = resp.Body.Close() }()

		if resp.StatusCode != 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to read response body: %w", err)
			}
			return mcp.NewToolResultError(fmt.Sprintf("failed to search %ss: %s", accountType, string(body))), nil
		}

		minimalUsers := make([]MinimalUser, 0, len(result.Users))

		for _, user := range result.Users {
			if user.Login != nil {
				mu := MinimalUser{
					Login:      user.GetLogin(),
					ID:         user.GetID(),
					ProfileURL: user.GetHTMLURL(),
					AvatarURL:  user.GetAvatarURL(),
				}
				minimalUsers = append(minimalUsers, mu)
			}
		}
		minimalResp := &MinimalSearchUsersResult{
			TotalCount:        result.GetTotal(),
			IncompleteResults: result.GetIncompleteResults(),
			Items:             minimalUsers,
		}
		if result.Total != nil {
			minimalResp.TotalCount = *result.Total
		}
		if result.IncompleteResults != nil {
			minimalResp.IncompleteResults = *result.IncompleteResults
		}

		r, err := json.Marshal(minimalResp)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return mcp.NewToolResultText(string(r)), nil
	}
}

// SearchUsers creates a tool to search for GitHub users.
func SearchUsers(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_users",
		mcp.WithDescription(t("TOOL_SEARCH_USERS_DESCRIPTION", "Search for GitHub users exclusively")),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        t("TOOL_SEARCH_USERS_USER_TITLE", "Search users"),
			ReadOnlyHint: ToBoolPtr(true),
		}),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search query using GitHub users search syntax scoped to type:user"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort field by category"),
			mcp.Enum("followers", "repositories", "joined"),
		),
		mcp.WithString("order",
			mcp.Description("Sort order"),
			mcp.Enum("asc", "desc"),
		),
		WithPagination(),
	), userOrOrgHandler("user", getClient)
}

// SearchOrgs creates a tool to search for GitHub organizations.
func SearchOrgs(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("search_orgs",
		mcp.WithDescription(t("TOOL_SEARCH_ORGS_DESCRIPTION", "Search for GitHub organizations exclusively")),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        t("TOOL_SEARCH_ORGS_USER_TITLE", "Search organizations"),
			ReadOnlyHint: ToBoolPtr(true),
		}),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search query using GitHub organizations search syntax scoped to type:org"),
		),
		mcp.WithString("sort",
			mcp.Description("Sort field by category"),
			mcp.Enum("followers", "repositories", "joined"),
		),
		mcp.WithString("order",
			mcp.Description("Sort order"),
			mcp.Enum("asc", "desc"),
		),
		WithPagination(),
	), userOrOrgHandler("org", getClient)
}

```

#### y/search_test.go

```go
package github

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/github/github-mcp-server/internal/toolsnaps"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SearchRepositories(t *testing.T) {
	// Verify tool definition once
	mockClient := github.NewClient(nil)
	tool, _ := SearchRepositories(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "search_repositories", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "query")
	assert.Contains(t, tool.InputSchema.Properties, "page")
	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"query"})

	// Setup mock search results
	mockSearchResult := &github.RepositoriesSearchResult{
		Total:             github.Ptr(2),
		IncompleteResults: github.Ptr(false),
		Repositories: []*github.Repository{
			{
				ID:              github.Ptr(int64(12345)),
				Name:            github.Ptr("repo-1"),
				FullName:        github.Ptr("owner/repo-1"),
				HTMLURL:         github.Ptr("https://github.com/owner/repo-1"),
				Description:     github.Ptr("Test repository 1"),
				StargazersCount: github.Ptr(100),
			},
			{
				ID:              github.Ptr(int64(67890)),
				Name:            github.Ptr("repo-2"),
				FullName:        github.Ptr("owner/repo-2"),
				HTMLURL:         github.Ptr("https://github.com/owner/repo-2"),
				Description:     github.Ptr("Test repository 2"),
				StargazersCount: github.Ptr(50),
			},
		},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult *github.RepositoriesSearchResult
		expectedErrMsg string
	}{
		{
			name: "successful repository search",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchRepositories,
					expectQueryParams(t, map[string]string{
						"q":        "golang test",
						"page":     "2",
						"per_page": "10",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query":   "golang test",
				"page":    float64(2),
				"perPage": float64(10),
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "repository search with default pagination",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchRepositories,
					expectQueryParams(t, map[string]string{
						"q":        "golang test",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "golang test",
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "search fails",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchRepositories,
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message": "Invalid query"}`))
					}),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "invalid:query",
			},
			expectError:    true,
			expectedErrMsg: "failed to search repositories",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup client with mock
			client := github.NewClient(tc.mockedClient)
			_, handler := SearchRepositories(stubGetClientFn(client), translations.NullTranslationHelper)

			// Create call request
			request := createMCPRequest(tc.requestArgs)

			// Call handler
			result, err := handler(context.Background(), request)

			// Verify results
			if tc.expectError {
				require.NoError(t, err)
				require.True(t, result.IsError)
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.False(t, result.IsError)

			// Parse the result and get the text content if no error
			textContent := getTextResult(t, result)

			// Unmarshal and verify the result
			var returnedResult github.RepositoriesSearchResult
			err = json.Unmarshal([]byte(textContent.Text), &returnedResult)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedResult.Total, *returnedResult.Total)
			assert.Equal(t, *tc.expectedResult.IncompleteResults, *returnedResult.IncompleteResults)
			assert.Len(t, returnedResult.Repositories, len(tc.expectedResult.Repositories))
			for i, repo := range returnedResult.Repositories {
				assert.Equal(t, *tc.expectedResult.Repositories[i].ID, *repo.ID)
				assert.Equal(t, *tc.expectedResult.Repositories[i].Name, *repo.Name)
				assert.Equal(t, *tc.expectedResult.Repositories[i].FullName, *repo.FullName)
				assert.Equal(t, *tc.expectedResult.Repositories[i].HTMLURL, *repo.HTMLURL)
			}

		})
	}
}

func Test_SearchCode(t *testing.T) {
	// Verify tool definition once
	mockClient := github.NewClient(nil)
	tool, _ := SearchCode(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "search_code", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "query")
	assert.Contains(t, tool.InputSchema.Properties, "sort")
	assert.Contains(t, tool.InputSchema.Properties, "order")
	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.Contains(t, tool.InputSchema.Properties, "page")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"query"})

	// Setup mock search results
	mockSearchResult := &github.CodeSearchResult{
		Total:             github.Ptr(2),
		IncompleteResults: github.Ptr(false),
		CodeResults: []*github.CodeResult{
			{
				Name:       github.Ptr("file1.go"),
				Path:       github.Ptr("path/to/file1.go"),
				SHA:        github.Ptr("abc123def456"),
				HTMLURL:    github.Ptr("https://github.com/owner/repo/blob/main/path/to/file1.go"),
				Repository: &github.Repository{Name: github.Ptr("repo"), FullName: github.Ptr("owner/repo")},
			},
			{
				Name:       github.Ptr("file2.go"),
				Path:       github.Ptr("path/to/file2.go"),
				SHA:        github.Ptr("def456abc123"),
				HTMLURL:    github.Ptr("https://github.com/owner/repo/blob/main/path/to/file2.go"),
				Repository: &github.Repository{Name: github.Ptr("repo"), FullName: github.Ptr("owner/repo")},
			},
		},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult *github.CodeSearchResult
		expectedErrMsg string
	}{
		{
			name: "successful code search with all parameters",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchCode,
					expectQueryParams(t, map[string]string{
						"q":        "fmt.Println language:go",
						"sort":     "indexed",
						"order":    "desc",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query":   "fmt.Println language:go",
				"sort":    "indexed",
				"order":   "desc",
				"page":    float64(1),
				"perPage": float64(30),
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "code search with minimal parameters",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchCode,
					expectQueryParams(t, map[string]string{
						"q":        "fmt.Println language:go",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "fmt.Println language:go",
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "search code fails",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchCode,
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message": "Validation Failed"}`))
					}),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "invalid:query",
			},
			expectError:    true,
			expectedErrMsg: "failed to search code",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup client with mock
			client := github.NewClient(tc.mockedClient)
			_, handler := SearchCode(stubGetClientFn(client), translations.NullTranslationHelper)

			// Create call request
			request := createMCPRequest(tc.requestArgs)

			// Call handler
			result, err := handler(context.Background(), request)

			// Verify results
			if tc.expectError {
				require.NoError(t, err)
				require.True(t, result.IsError)
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.False(t, result.IsError)

			// Parse the result and get the text content if no error
			textContent := getTextResult(t, result)

			// Unmarshal and verify the result
			var returnedResult github.CodeSearchResult
			err = json.Unmarshal([]byte(textContent.Text), &returnedResult)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedResult.Total, *returnedResult.Total)
			assert.Equal(t, *tc.expectedResult.IncompleteResults, *returnedResult.IncompleteResults)
			assert.Len(t, returnedResult.CodeResults, len(tc.expectedResult.CodeResults))
			for i, code := range returnedResult.CodeResults {
				assert.Equal(t, *tc.expectedResult.CodeResults[i].Name, *code.Name)
				assert.Equal(t, *tc.expectedResult.CodeResults[i].Path, *code.Path)
				assert.Equal(t, *tc.expectedResult.CodeResults[i].SHA, *code.SHA)
				assert.Equal(t, *tc.expectedResult.CodeResults[i].HTMLURL, *code.HTMLURL)
				assert.Equal(t, *tc.expectedResult.CodeResults[i].Repository.FullName, *code.Repository.FullName)
			}
		})
	}
}

func Test_SearchUsers(t *testing.T) {
	// Verify tool definition once
	mockClient := github.NewClient(nil)
	tool, _ := SearchUsers(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "search_users", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "query")
	assert.Contains(t, tool.InputSchema.Properties, "sort")
	assert.Contains(t, tool.InputSchema.Properties, "order")
	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.Contains(t, tool.InputSchema.Properties, "page")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"query"})

	// Setup mock search results
	mockSearchResult := &github.UsersSearchResult{
		Total:             github.Ptr(2),
		IncompleteResults: github.Ptr(false),
		Users: []*github.User{
			{
				Login:     github.Ptr("user1"),
				ID:        github.Ptr(int64(1001)),
				HTMLURL:   github.Ptr("https://github.com/user1"),
				AvatarURL: github.Ptr("https://avatars.githubusercontent.com/u/1001"),
			},
			{
				Login:     github.Ptr("user2"),
				ID:        github.Ptr(int64(1002)),
				HTMLURL:   github.Ptr("https://github.com/user2"),
				AvatarURL: github.Ptr("https://avatars.githubusercontent.com/u/1002"),
				Type:      github.Ptr("User"),
			},
		},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult *github.UsersSearchResult
		expectedErrMsg string
	}{
		{
			name: "successful users search with all parameters",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					expectQueryParams(t, map[string]string{
						"q":        "type:user location:finland language:go",
						"sort":     "followers",
						"order":    "desc",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query":   "location:finland language:go",
				"sort":    "followers",
				"order":   "desc",
				"page":    float64(1),
				"perPage": float64(30),
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "users search with minimal parameters",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					expectQueryParams(t, map[string]string{
						"q":        "type:user location:finland language:go",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "location:finland language:go",
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "search users fails",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message": "Validation Failed"}`))
					}),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "invalid:query",
			},
			expectError:    true,
			expectedErrMsg: "failed to search users",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup client with mock
			client := github.NewClient(tc.mockedClient)
			_, handler := SearchUsers(stubGetClientFn(client), translations.NullTranslationHelper)

			// Create call request
			request := createMCPRequest(tc.requestArgs)

			// Call handler
			result, err := handler(context.Background(), request)

			// Verify results
			if tc.expectError {
				require.NoError(t, err)
				require.True(t, result.IsError)
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.False(t, result.IsError)

			// Parse the result and get the text content if no error
			require.NotNil(t, result)

			textContent := getTextResult(t, result)

			// Unmarshal and verify the result
			var returnedResult MinimalSearchUsersResult
			err = json.Unmarshal([]byte(textContent.Text), &returnedResult)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedResult.Total, returnedResult.TotalCount)
			assert.Equal(t, *tc.expectedResult.IncompleteResults, returnedResult.IncompleteResults)
			assert.Len(t, returnedResult.Items, len(tc.expectedResult.Users))
			for i, user := range returnedResult.Items {
				assert.Equal(t, *tc.expectedResult.Users[i].Login, user.Login)
				assert.Equal(t, *tc.expectedResult.Users[i].ID, user.ID)
				assert.Equal(t, *tc.expectedResult.Users[i].HTMLURL, user.ProfileURL)
				assert.Equal(t, *tc.expectedResult.Users[i].AvatarURL, user.AvatarURL)
			}
		})
	}
}

func Test_SearchOrgs(t *testing.T) {
	// Verify tool definition once
	mockClient := github.NewClient(nil)
	tool, _ := SearchOrgs(stubGetClientFn(mockClient), translations.NullTranslationHelper)

	assert.Equal(t, "search_orgs", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "query")
	assert.Contains(t, tool.InputSchema.Properties, "sort")
	assert.Contains(t, tool.InputSchema.Properties, "order")
	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.Contains(t, tool.InputSchema.Properties, "page")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"query"})

	// Setup mock search results
	mockSearchResult := &github.UsersSearchResult{
		Total:             github.Ptr(int(2)),
		IncompleteResults: github.Ptr(false),
		Users: []*github.User{
			{
				Login:     github.Ptr("org-1"),
				ID:        github.Ptr(int64(111)),
				HTMLURL:   github.Ptr("https://github.com/org-1"),
				AvatarURL: github.Ptr("https://avatars.githubusercontent.com/u/111?v=4"),
			},
			{
				Login:     github.Ptr("org-2"),
				ID:        github.Ptr(int64(222)),
				HTMLURL:   github.Ptr("https://github.com/org-2"),
				AvatarURL: github.Ptr("https://avatars.githubusercontent.com/u/222?v=4"),
			},
		},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult *github.UsersSearchResult
		expectedErrMsg string
	}{
		{
			name: "successful org search",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					expectQueryParams(t, map[string]string{
						"q":        "type:org github",
						"page":     "1",
						"per_page": "30",
					}).andThen(
						mockResponse(t, http.StatusOK, mockSearchResult),
					),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "github",
			},
			expectError:    false,
			expectedResult: mockSearchResult,
		},
		{
			name: "org search fails",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetSearchUsers,
					http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message": "Validation Failed"}`))
					}),
				),
			),
			requestArgs: map[string]interface{}{
				"query": "invalid:query",
			},
			expectError:    true,
			expectedErrMsg: "failed to search orgs",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup client with mock
			client := github.NewClient(tc.mockedClient)
			_, handler := SearchOrgs(stubGetClientFn(client), translations.NullTranslationHelper)

			// Create call request
			request := createMCPRequest(tc.requestArgs)

			// Call handler
			result, err := handler(context.Background(), request)

			// Verify results
			if tc.expectError {
				require.NoError(t, err)
				require.True(t, result.IsError)
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)

			textContent := getTextResult(t, result)

			// Unmarshal and verify the result
			var returnedResult MinimalSearchUsersResult
			err = json.Unmarshal([]byte(textContent.Text), &returnedResult)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedResult.Total, returnedResult.TotalCount)
			assert.Equal(t, *tc.expectedResult.IncompleteResults, returnedResult.IncompleteResults)
			assert.Len(t, returnedResult.Items, len(tc.expectedResult.Users))
			for i, org := range returnedResult.Items {
				assert.Equal(t, *tc.expectedResult.Users[i].Login, org.Login)
				assert.Equal(t, *tc.expectedResult.Users[i].ID, org.ID)
				assert.Equal(t, *tc.expectedResult.Users[i].HTMLURL, org.ProfileURL)
				assert.Equal(t, *tc.expectedResult.Users[i].AvatarURL, org.AvatarURL)
			}
		})
	}
}

```

#### y/search_utils.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
)

func searchHandler(
	ctx context.Context,
	getClient GetClientFn,
	request mcp.CallToolRequest,
	searchType string,
	errorPrefix string,
) (*mcp.CallToolResult, error) {
	query, err := RequiredParam[string](request, "query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	query = fmt.Sprintf("is:%s %s", searchType, query)

	owner, err := OptionalParam[string](request, "owner")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	repo, err := OptionalParam[string](request, "repo")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if owner != "" && repo != "" {
		query = fmt.Sprintf("repo:%s/%s %s", owner, repo, query)
	}

	sort, err := OptionalParam[string](request, "sort")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	order, err := OptionalParam[string](request, "order")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	pagination, err := OptionalPaginationParams(request)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	opts := &github.SearchOptions{
		// Default to "created" if no sort is provided, as it's a common use case.
		Sort:  sort,
		Order: order,
		ListOptions: github.ListOptions{
			Page:    pagination.Page,
			PerPage: pagination.PerPage,
		},
	}

	client, err := getClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get GitHub client: %w", errorPrefix, err)
	}
	result, resp, err := client.Search.Issues(ctx, query, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errorPrefix, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to read response body: %w", errorPrefix, err)
		}
		return mcp.NewToolResultError(fmt.Sprintf("%s: %s", errorPrefix, string(body))), nil
	}

	r, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to marshal response: %w", errorPrefix, err)
	}

	return mcp.NewToolResultText(string(r)), nil
}

```

#### y/server.go

```go
package github

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NewServer creates a new GitHub MCP server with the specified GH client and logger.

func NewServer(version string, opts ...server.ServerOption) *server.MCPServer {
	// Add default options
	defaultOpts := []server.ServerOption{
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	}
	opts = append(defaultOpts, opts...)

	// Create a new MCP server
	s := server.NewMCPServer(
		"github-mcp-server",
		version,
		opts...,
	)
	return s
}

// OptionalParamOK is a helper function that can be used to fetch a requested parameter from the request.
// It returns the value, a boolean indicating if the parameter was present, and an error if the type is wrong.
func OptionalParamOK[T any](r mcp.CallToolRequest, p string) (value T, ok bool, err error) {
	// Check if the parameter is present in the request
	val, exists := r.GetArguments()[p]
	if !exists {
		// Not present, return zero value, false, no error
		return
	}

	// Check if the parameter is of the expected type
	value, ok = val.(T)
	if !ok {
		// Present but wrong type
		err = fmt.Errorf("parameter %s is not of type %T, is %T", p, value, val)
		ok = true // Set ok to true because the parameter *was* present, even if wrong type
		return
	}

	// Present and correct type
	ok = true
	return
}

// isAcceptedError checks if the error is an accepted error.
func isAcceptedError(err error) bool {
	var acceptedError *github.AcceptedError
	return errors.As(err, &acceptedError)
}

// RequiredParam is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request.
// 2. Checks if the parameter is of the expected type.
// 3. Checks if the parameter is not empty, i.e: non-zero value
func RequiredParam[T comparable](r mcp.CallToolRequest, p string) (T, error) {
	var zero T

	// Check if the parameter is present in the request
	if _, ok := r.GetArguments()[p]; !ok {
		return zero, fmt.Errorf("missing required parameter: %s", p)
	}

	// Check if the parameter is of the expected type
	val, ok := r.GetArguments()[p].(T)
	if !ok {
		return zero, fmt.Errorf("parameter %s is not of type %T", p, zero)
	}

	if val == zero {
		return zero, fmt.Errorf("missing required parameter: %s", p)
	}

	return val, nil
}

// RequiredInt is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request.
// 2. Checks if the parameter is of the expected type.
// 3. Checks if the parameter is not empty, i.e: non-zero value
func RequiredInt(r mcp.CallToolRequest, p string) (int, error) {
	v, err := RequiredParam[float64](r, p)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

// OptionalParam is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request, if not, it returns its zero-value
// 2. If it is present, it checks if the parameter is of the expected type and returns it
func OptionalParam[T any](r mcp.CallToolRequest, p string) (T, error) {
	var zero T

	// Check if the parameter is present in the request
	if _, ok := r.GetArguments()[p]; !ok {
		return zero, nil
	}

	// Check if the parameter is of the expected type
	if _, ok := r.GetArguments()[p].(T); !ok {
		return zero, fmt.Errorf("parameter %s is not of type %T, is %T", p, zero, r.GetArguments()[p])
	}

	return r.GetArguments()[p].(T), nil
}

// OptionalIntParam is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request, if not, it returns its zero-value
// 2. If it is present, it checks if the parameter is of the expected type and returns it
func OptionalIntParam(r mcp.CallToolRequest, p string) (int, error) {
	v, err := OptionalParam[float64](r, p)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

// OptionalIntParamWithDefault is a helper function that can be used to fetch a requested parameter from the request
// similar to optionalIntParam, but it also takes a default value.
func OptionalIntParamWithDefault(r mcp.CallToolRequest, p string, d int) (int, error) {
	v, err := OptionalIntParam(r, p)
	if err != nil {
		return 0, err
	}
	if v == 0 {
		return d, nil
	}
	return v, nil
}

// OptionalStringArrayParam is a helper function that can be used to fetch a requested parameter from the request.
// It does the following checks:
// 1. Checks if the parameter is present in the request, if not, it returns its zero-value
// 2. If it is present, iterates the elements and checks each is a string
func OptionalStringArrayParam(r mcp.CallToolRequest, p string) ([]string, error) {
	// Check if the parameter is present in the request
	if _, ok := r.GetArguments()[p]; !ok {
		return []string{}, nil
	}

	switch v := r.GetArguments()[p].(type) {
	case nil:
		return []string{}, nil
	case []string:
		return v, nil
	case []any:
		strSlice := make([]string, len(v))
		for i, v := range v {
			s, ok := v.(string)
			if !ok {
				return []string{}, fmt.Errorf("parameter %s is not of type string, is %T", p, v)
			}
			strSlice[i] = s
		}
		return strSlice, nil
	default:
		return []string{}, fmt.Errorf("parameter %s could not be coerced to []string, is %T", p, r.GetArguments()[p])
	}
}

// WithPagination returns a ToolOption that adds "page" and "perPage" parameters to the tool.
// The "page" parameter is optional, min 1.
// The "perPage" parameter is optional, min 1, max 100. If unset, defaults to 30.
// https://docs.github.com/en/rest/using-the-rest-api/using-pagination-in-the-rest-api
func WithPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		mcp.WithNumber("page",
			mcp.Description("Page number for pagination (min 1)"),
			mcp.Min(1),
		)(tool)

		mcp.WithNumber("perPage",
			mcp.Description("Results per page for pagination (min 1, max 100)"),
			mcp.Min(1),
			mcp.Max(100),
		)(tool)
	}
}

// PaginationParams holds the REST pagination parameters, plus an optional
// cursor for tools backed by GraphQL.
type PaginationParams struct {
	Page    int
	PerPage int
	After   string
}

// ToGraphQLParams converts REST pagination parameters to GraphQL parameters.
// If After is set, it is used as the cursor.
func (p PaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	return CursorPaginationParams{
		PerPage: p.PerPage,
		After:   p.After,
	}.ToGraphQLParams()
}

// OptionalPaginationParams returns the "page", "perPage" and "after" parameters from the request,
// or their default values if not present, "page" default is 1, "perPage" default is 30,
// "after" default is empty.
// In future, we may want to make the default values configurable, or even have this
// function returned from `withPagination`, where the defaults are provided alongside
// the min/max values.
func OptionalPaginationParams(r mcp.CallToolRequest) (PaginationParams, error) {
	page, err := OptionalIntParamWithDefault(r, "page", 1)
	if err != nil {
		return PaginationParams{}, err
	}
	perPage, err := OptionalIntParamWithDefault(r, "perPage", 30)
	if err != nil {
		return PaginationParams{}, err
	}
	after, err := OptionalParam[string](r, "after")
	if err != nil {
		return PaginationParams{}, err
	}
	return PaginationParams{
		Page:    page,
		PerPage: perPage,
		After:   after,
	}, nil
}

func MarshalledTextResult(v any) *mcp.CallToolResult {
	data, err := json.Marshal(v)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("failed to marshal text result to json", err)
	}

	return mcp.NewToolResultText(string(data))
}

```

#### y/server_test.go

```go
package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/github/github-mcp-server/pkg/raw"
	"github.com/google/go-github/v73/github"
	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func stubGetClientFn(client *github.Client) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return client, nil
	}
}

func stubGetClientFromHTTPFn(client *http.Client) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return github.NewClient(client), nil
	}
}

func stubGetClientFnErr(err string) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return nil, errors.New(err)
	}
}

func stubGetGQLClientFn(client *githubv4.Client) GetGQLClientFn {
	return func(_ context.Context) (*githubv4.Client, error) {
		return client, nil
	}
}

func stubGetRawClientFn(client *raw.Client) raw.GetRawClientFn {
	return func(_ context.Context) (*raw.Client, error) {
		return client, nil
	}
}

func badRequestHandler(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		structuredErrorResponse := github.ErrorResponse{
			Message: msg,
		}

		b, err := json.Marshal(structuredErrorResponse)
		if err != nil {
			http.Error(w, "failed to marshal error response", http.StatusInternalServerError)
		}

		http.Error(w, string(b), http.StatusBadRequest)
	}
}

func Test_IsAcceptedError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectAccepted bool
	}{
		{
			name:           "github AcceptedError",
			err:            &github.AcceptedError{},
			expectAccepted: true,
		},
		{
			name:           "regular error",
			err:            fmt.Errorf("some other error"),
			expectAccepted: false,
		},
		{
			name:           "nil error",
			err:            nil,
			expectAccepted: false,
		},
		{
			name:           "wrapped AcceptedError",
			err:            fmt.Errorf("wrapped: %w", &github.AcceptedError{}),
			expectAccepted: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isAcceptedError(tc.err)
			assert.Equal(t, tc.expectAccepted, result)
		})
	}
}

func Test_RequiredStringParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    string
		expectError bool
	}{
		{
			name:        "valid string parameter",
			params:      map[string]interface{}{"name": "test-value"},
			paramName:   "name",
			expected:    "test-value",
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "name",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty string parameter",
			params:      map[string]interface{}{"name": ""},
			paramName:   "name",
			expected:    "",
			expectError: true,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"name": 123},
			paramName:   "name",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := RequiredParam[string](request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_OptionalStringParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    string
		expectError bool
	}{
		{
			name:        "valid string parameter",
			params:      map[string]interface{}{"name": "test-value"},
			paramName:   "name",
			expected:    "test-value",
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "name",
			expected:    "",
			expectError: false,
		},
		{
			name:        "empty string parameter",
			params:      map[string]interface{}{"name": ""},
			paramName:   "name",
			expected:    "",
			expectError: false,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"name": 123},
			paramName:   "name",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalParam[string](request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_RequiredInt(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    int
		expectError bool
	}{
		{
			name:        "valid number parameter",
			params:      map[string]interface{}{"count": float64(42)},
			paramName:   "count",
			expected:    42,
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "count",
			expected:    0,
			expectError: true,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"count": "not-a-number"},
			paramName:   "count",
			expected:    0,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := RequiredInt(request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
func Test_OptionalIntParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    int
		expectError bool
	}{
		{
			name:        "valid number parameter",
			params:      map[string]interface{}{"count": float64(42)},
			paramName:   "count",
			expected:    42,
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "count",
			expected:    0,
			expectError: false,
		},
		{
			name:        "zero value",
			params:      map[string]interface{}{"count": float64(0)},
			paramName:   "count",
			expected:    0,
			expectError: false,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"count": "not-a-number"},
			paramName:   "count",
			expected:    0,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalIntParam(request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_OptionalNumberParamWithDefault(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		defaultVal  int
		expected    int
		expectError bool
	}{
		{
			name:        "valid number parameter",
			params:      map[string]interface{}{"count": float64(42)},
			paramName:   "count",
			defaultVal:  10,
			expected:    42,
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "count",
			defaultVal:  10,
			expected:    10,
			expectError: false,
		},
		{
			name:        "zero value",
			params:      map[string]interface{}{"count": float64(0)},
			paramName:   "count",
			defaultVal:  10,
			expected:    10,
			expectError: false,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"count": "not-a-number"},
			paramName:   "count",
			defaultVal:  10,
			expected:    0,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalIntParamWithDefault(request, tc.paramName, tc.defaultVal)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_OptionalBooleanParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    bool
		expectError bool
	}{
		{
			name:        "true value",
			params:      map[string]interface{}{"flag": true},
			paramName:   "flag",
			expected:    true,
			expectError: false,
		},
		{
			name:        "false value",
			params:      map[string]interface{}{"flag": false},
			paramName:   "flag",
			expected:    false,
			expectError: false,
		},
		{
			name:        "missing parameter",
			params:      map[string]interface{}{},
			paramName:   "flag",
			expected:    false,
			expectError: false,
		},
		{
			name:        "wrong type parameter",
			params:      map[string]interface{}{"flag": "not-a-boolean"},
			paramName:   "flag",
			expected:    false,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalParam[bool](request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestOptionalStringArrayParam(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]interface{}
		paramName   string
		expected    []string
		expectError bool
	}{
		{
			name:        "parameter not in request",
			params:      map[string]any{},
			paramName:   "flag",
			expected:    []string{},
			expectError: false,
		},
		{
			name: "valid any array parameter",
			params: map[string]any{
				"flag": []any{"v1", "v2"},
			},
			paramName:   "flag",
			expected:    []string{"v1", "v2"},
			expectError: false,
		},
		{
			name: "valid string array parameter",
			params: map[string]any{
				"flag": []string{"v1", "v2"},
			},
			paramName:   "flag",
			expected:    []string{"v1", "v2"},
			expectError: false,
		},
		{
			name: "wrong type parameter",
			params: map[string]any{
				"flag": 1,
			},
			paramName:   "flag",
			expected:    []string{},
			expectError: true,
		},
		{
			name: "wrong slice type parameter",
			params: map[string]any{
				"flag": []any{"foo", 2},
			},
			paramName:   "flag",
			expected:    []string{},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalStringArrayParam(request, tc.paramName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestOptionalPaginationParams(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]any
		expected    PaginationParams
		expectError bool
	}{
		{
			name:   "no pagination parameters, default values",
			params: map[string]any{},
			expected: PaginationParams{
				Page:    1,
				PerPage: 30,
			},
			expectError: false,
		},
		{
			name: "page parameter, default perPage",
			params: map[string]any{
				"page": float64(2),
			},
			expected: PaginationParams{
				Page:    2,
				PerPage: 30,
			},
			expectError: false,
		},
		{
			name: "perPage parameter, default page",
			params: map[string]any{
				"perPage": float64(50),
			},
			expected: PaginationParams{
				Page:    1,
				PerPage: 50,
			},
			expectError: false,
		},
		{
			name: "page and perPage parameters",
			params: map[string]any{
				"page":    float64(2),
				"perPage": float64(50),
			},
			expected: PaginationParams{
				Page:    2,
				PerPage: 50,
			},
			expectError: false,
		},
		{
			name: "after parameter, default page and perPage",
			params: map[string]any{
				"after": "cursor-abc",
			},
			expected: PaginationParams{
				Page:    1,
				PerPage: 30,
				After:   "cursor-abc",
			},
			expectError: false,
		},
		{
			name: "invalid page parameter",
			params: map[string]any{
				"page": "not-a-number",
			},
			expected:    PaginationParams{},
			expectError: true,
		},
		{
			name: "invalid perPage parameter",
			params: map[string]any{
				"perPage": "not-a-number",
			},
			expected:    PaginationParams{},
			expectError: true,
		},
		{
			name: "invalid after parameter",
			params: map[string]any{
				"after": float64(1),
			},
			expected:    PaginationParams{},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalPaginationParams(request)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestPaginationParams_ToGraphQLParams(t *testing.T) {
	tests := []struct {
		name        string
		params      PaginationParams
		expected    *GraphQLPaginationParams
		errContains string
	}{
		{
			name:     "perPage without cursor",
			params:   PaginationParams{Page: 2, PerPage: 30},
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(30))},
		},
		{
			name:     "after is used as the cursor",
			params:   PaginationParams{Page: 1, PerPage: 10, After: "cursor-abc"},
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(10)), After: github.Ptr("cursor-abc")},
		},
		{
			name:        "perPage above maximum",
			params:      PaginationParams{Page: 1, PerPage: 101},
			errContains: "perPage value 101 exceeds maximum of 100",
		},
		{
			name:        "negative perPage",
			params:      PaginationParams{Page: 1, PerPage: -1},
			errContains: "perPage value -1 cannot be negative",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.params.ToGraphQLParams()
			if tc.errContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errContains)
				assert.Nil(t, result)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

```

#### y/sub_issue.go

```go
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AddSubIssue creates a tool to add a sub-issue to a parent issue.
func AddSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("add_sub_issue",
			mcp.WithDescription(t("TOOL_ADD_SUB_ISSUE_DESCRIPTION", "Add a sub-issue to a parent issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_ADD_SUB_ISSUE_USER_TITLE", "Add sub-issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("The number of the parent issue"),
			),
			mcp.WithNumber("sub_issue_id",
				mcp.Required(),
				mcp.Description("The ID of the sub-issue to add. ID is not the same as issue number"),
			),
			mcp.WithBoolean("replace_parent",
				mcp.Description("When true, replaces the sub-issue's current parent issue"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := RequiredInt(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			replaceParent, err := OptionalParam[bool](request, "replace_parent")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			subIssueRequest := github.SubIssueRequest{
				SubIssueID: int64(subIssueID),
			}
			if replaceParent {
				subIssueRequest.ReplaceParent = github.Ptr(true)
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			subIssue, resp, err := client.SubIssue.Add(ctx, owner, repo, int64(issueNumber), subIssueRequest)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to add sub-issue", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusCreated {
				return mcp.NewToolResultError(fmt.Sprintf("failed to add sub-issue: unexpected status %d", resp.StatusCode)), nil
			}

			r, err := json.Marshal(subIssue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListSubIssues creates a tool to list the sub-issues of an issue.
func ListSubIssues(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_sub_issues",
			mcp.WithDescription(t("TOOL_LIST_SUB_ISSUES_DESCRIPTION", "List sub-issues for a specific issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_SUB_ISSUES_USER_TITLE", "List sub-issues"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("Issue number"),
			),
			mcp.WithNumber("page",
				mcp.Description("Page number for pagination (default: 1)"),
			),
			mcp.WithNumber("per_page",
				mcp.Description("Number of results per page (max 100, default: 30)"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			page, err := OptionalIntParamWithDefault(request, "page", 1)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			perPage, err := OptionalIntParamWithDefault(request, "per_page", 30)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			opts := &github.IssueListOptions{
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: perPage,
				},
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			subIssues, resp, err := client.SubIssue.ListByIssue(ctx, owner, repo, int64(issueNumber), opts)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to list sub-issues", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				return mcp.NewToolResultError(fmt.Sprintf("failed to list sub-issues: unexpected status %d", resp.StatusCode)), nil
			}

			r, err := json.Marshal(subIssues)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// RemoveSubIssue creates a tool to remove a sub-issue from a parent issue.
func RemoveSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("remove_sub_issue",
			mcp.WithDescription(t("TOOL_REMOVE_SUB_ISSUE_DESCRIPTION", "Remove a sub-issue from a parent issue in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_REMOVE_SUB_ISSUE_USER_TITLE", "Remove sub-issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("The number of the parent issue"),
			),
			mcp.WithNumber("sub_issue_id",
				mcp.Required(),
				mcp.Description("The ID of the sub-issue to remove. ID is not the same as issue number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := RequiredInt(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// SubIssueService.Remove calls .../sub_issues instead of .../sub_issue
			// (google/go-github#3613), so build the request here.
			u := fmt.Sprintf("repos/%v/%v/issues/%v/sub_issue", owner, repo, issueNumber)
			req, err := client.NewRequest(http.MethodDelete, u, github.SubIssueRequest{SubIssueID: int64(subIssueID)})
			if err != nil {
				return nil, fmt.Errorf("failed to create request: %w", err)
			}
			parent := new(github.Issue)
			resp, err := client.Do(ctx, req, parent)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to remove sub-issue", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				return mcp.NewToolResultError(fmt.Sprintf("failed to remove sub-issue: unexpected status %d", resp.StatusCode)), nil
			}

			r, err := json.Marshal(parent)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ReprioritizeSubIssue creates a tool to move a sub-issue to another position
// in its parent issue's sub-issue list.
func ReprioritizeSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("reprioritize_sub_issue",
			mcp.WithDescription(t("TOOL_REPRIORITIZE_SUB_ISSUE_DESCRIPTION", "Reprioritize a sub-issue to a different position in the parent issue's sub-issue list.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_REPRIORITIZE_SUB_ISSUE_USER_TITLE", "Reprioritize sub-issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("issue_number",
				mcp.Required(),
				mcp.Description("The number of the parent issue"),
			),
			mcp.WithNumber("sub_issue_id",
				mcp.Required(),
				mcp.Description("The ID of the sub-issue to reprioritize. ID is not the same as issue number"),
			),
			mcp.WithNumber("after_id",
				mcp.Description("The ID of the sub-issue to be prioritized after (either after_id OR before_id should be specified)"),
			),
			mcp.WithNumber("before_id",
				mcp.Description("The ID of the sub-issue to be prioritized before (either after_id OR before_id should be specified)"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			issueNumber, err := RequiredInt(request, "issue_number")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := RequiredInt(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			afterID, err := OptionalIntParam(request, "after_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			beforeID, err := OptionalIntParam(request, "before_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// An ID of 0 means the parameter was not given; no issue has ID 0.
			if afterID == 0 && beforeID == 0 {
				return mcp.NewToolResultError("either after_id or before_id must be specified"), nil
			}
			if afterID != 0 && beforeID != 0 {
				return mcp.NewToolResultError("only one of after_id or before_id should be specified, not both"), nil
			}

			subIssueRequest := github.SubIssueRequest{
				SubIssueID: int64(subIssueID),
			}
			if afterID != 0 {
				subIssueRequest.AfterID = github.Ptr(int64(afterID))
			}
			if beforeID != 0 {
				subIssueRequest.BeforeID = github.Ptr(int64(beforeID))
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			subIssue, resp, err := client.SubIssue.Reprioritize(ctx, owner, repo, int64(issueNumber), subIssueRequest)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, "failed to reprioritize sub-issue", resp, err), nil
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				return mcp.NewToolResultError(fmt.Sprintf("failed to reprioritize sub-issue: unexpected status %d", resp.StatusCode)), nil
			}

			r, err := json.Marshal(subIssue)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

```

#### y/sub_issue_test.go

```go
package github

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/github/github-mcp-server/internal/toolsnaps"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_AddSubIssue(t *testing.T) {
	mockClient := github.NewClient(nil)
	tool, _ := AddSubIssue(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "add_sub_issue", tool.Name)
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo", "issue_number", "sub_issue_id"})

	mockSubIssue := &github.Issue{
		ID:     github.Ptr(int64(123)),
		Number: github.Ptr(43),
		Title:  github.Ptr("Sub-issue"),
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectedErrMsg string
	}{
		{
			name: "successful add with replace_parent",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expect(t, expectations{
						path:        "/repos/owner/repo/issues/42/sub_issues",
						requestBody: map[string]any{"sub_issue_id": float64(123), "replace_parent": true},
					}).andThen(mockResponse(t, http.StatusCreated, mockSubIssue)),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":          "owner",
				"repo":           "repo",
				"issue_number":   float64(42),
				"sub_issue_id":   float64(123),
				"replace_parent": true,
			},
		},
		{
			name: "successful add without replace_parent",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectRequestBody(t, map[string]any{"sub_issue_id": float64(123)}).
						andThen(mockResponse(t, http.StatusCreated, mockSubIssue)),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
			},
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusUnprocessableEntity, `{"message": "Validation Failed"}`),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "failed to add sub-issue",
		},
		{
			name: "unexpected status",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusOK, mockSubIssue),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "failed to add sub-issue",
		},
		{
			name:         "missing sub_issue_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
			},
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := github.NewClient(tc.mockedClient)
			_, handler := AddSubIssue(stubGetClientFn(client), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.requestArgs))
			require.NoError(t, err)

			if tc.expectedErrMsg != "" {
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			textContent := getTextResult(t, result)
			var returned github.Issue
			require.NoError(t, json.Unmarshal([]byte(textContent.Text), &returned))
			assert.Equal(t, *mockSubIssue.ID, *returned.ID)
			assert.Equal(t, *mockSubIssue.Title, *returned.Title)
		})
	}
}

func Test_ListSubIssues(t *testing.T) {
	mockClient := github.NewClient(nil)
	tool, _ := ListSubIssues(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "list_sub_issues", tool.Name)
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo", "issue_number"})

	mockSubIssues := []*github.Issue{
		{ID: github.Ptr(int64(1)), Number: github.Ptr(2), Title: github.Ptr("First")},
		{ID: github.Ptr(int64(3)), Number: github.Ptr(4), Title: github.Ptr("Second")},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectedErrMsg string
	}{
		{
			name: "successful list with defaults",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectQueryParams(t, map[string]string{"page": "1", "per_page": "30"}).
						andThen(mockResponse(t, http.StatusOK, mockSubIssues)),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
			},
		},
		{
			name: "successful list with pagination",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectQueryParams(t, map[string]string{"page": "2", "per_page": "10"}).
						andThen(mockResponse(t, http.StatusOK, mockSubIssues)),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"page":         float64(2),
				"per_page":     float64(10),
			},
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusNotFound, `{"message": "Not Found"}`),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
			},
			expectedErrMsg: "failed to list sub-issues",
		},
		{
			name:         "missing issue_number",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
			},
			expectedErrMsg: "missing required parameter: issue_number",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := github.NewClient(tc.mockedClient)
			_, handler := ListSubIssues(stubGetClientFn(client), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.requestArgs))
			require.NoError(t, err)

			if tc.expectedErrMsg != "" {
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			textContent := getTextResult(t, result)
			var returned []*github.Issue
			require.NoError(t, json.Unmarshal([]byte(textContent.Text), &returned))
			require.Len(t, returned, len(mockSubIssues))
			for i := range returned {
				assert.Equal(t, *mockSubIssues[i].ID, *returned[i].ID)
				assert.Equal(t, *mockSubIssues[i].Title, *returned[i].Title)
			}
		})
	}
}

func Test_RemoveSubIssue(t *testing.T) {
	mockClient := github.NewClient(nil)
	tool, _ := RemoveSubIssue(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "remove_sub_issue", tool.Name)
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo", "issue_number", "sub_issue_id"})

	mockParent := &github.Issue{
		ID:     github.Ptr(int64(999)),
		Number: github.Ptr(42),
		Title:  github.Ptr("Parent"),
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectedErrMsg string
	}{
		{
			name: "successful remove",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					expect(t, expectations{
						path:        "/repos/owner/repo/issues/42/sub_issue",
						requestBody: map[string]any{"sub_issue_id": float64(123)},
					}).andThen(mockResponse(t, http.StatusOK, mockParent)),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
			},
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusNotFound, `{"message": "Not Found"}`),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "failed to remove sub-issue",
		},
		{
			name: "unexpected status",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusAccepted, mockParent),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "failed to remove sub-issue",
		},
		{
			name:         "missing sub_issue_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
			},
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := github.NewClient(tc.mockedClient)
			_, handler := RemoveSubIssue(stubGetClientFn(client), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.requestArgs))
			require.NoError(t, err)

			if tc.expectedErrMsg != "" {
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			textContent := getTextResult(t, result)
			var returned github.Issue
			require.NoError(t, json.Unmarshal([]byte(textContent.Text), &returned))
			assert.Equal(t, *mockParent.ID, *returned.ID)
			assert.Equal(t, *mockParent.Number, *returned.Number)
		})
	}
}

func Test_ReprioritizeSubIssue(t *testing.T) {
	mockClient := github.NewClient(nil)
	tool, _ := ReprioritizeSubIssue(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	assert.Equal(t, "reprioritize_sub_issue", tool.Name)
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo", "issue_number", "sub_issue_id"})

	mockSubIssue := &github.Issue{
		ID:     github.Ptr(int64(123)),
		Number: github.Ptr(43),
		Title:  github.Ptr("Sub-issue"),
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectedErrMsg string
	}{
		{
			name: "successful reprioritize with after_id",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					expect(t, expectations{
						path:        "/repos/owner/repo/issues/42/sub_issues/priority",
						requestBody: map[string]any{"sub_issue_id": float64(123), "after_id": float64(456)},
					}).andThen(mockResponse(t, http.StatusOK, mockSubIssue)),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
				"after_id":     float64(456),
			},
		},
		{
			name: "successful reprioritize with before_id",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					expectRequestBody(t, map[string]any{"sub_issue_id": float64(123), "before_id": float64(789)}).
						andThen(mockResponse(t, http.StatusOK, mockSubIssue)),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
				"before_id":    float64(789),
			},
		},
		{
			name:         "neither after_id nor before_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "either after_id or before_id must be specified",
		},
		{
			name:         "both after_id and before_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
				"after_id":     float64(456),
				"before_id":    float64(789),
			},
			expectedErrMsg: "only one of after_id or before_id should be specified, not both",
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusUnprocessableEntity, `{"message": "Validation Failed"}`),
				),
			),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"sub_issue_id": float64(123),
				"after_id":     float64(456),
			},
			expectedErrMsg: "failed to reprioritize sub-issue",
		},
		{
			name:         "missing sub_issue_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]interface{}{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(42),
				"after_id":     float64(456),
			},
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := github.NewClient(tc.mockedClient)
			_, handler := ReprioritizeSubIssue(stubGetClientFn(client), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.requestArgs))
			require.NoError(t, err)

			if tc.expectedErrMsg != "" {
				errorContent := getErrorResult(t, result)
				assert.Contains(t, errorContent.Text, tc.expectedErrMsg)
				return
			}

			textContent := getTextResult(t, result)
			var returned github.Issue
			require.NoError(t, json.Unmarshal([]byte(textContent.Text), &returned))
			assert.Equal(t, *mockSubIssue.ID, *returned.ID)
			assert.Equal(t, *mockSubIssue.Title, *returned.Title)
		})
	}
}

```

#### y/tools.go

```go
package github

import (
	"context"

	"github.com/github/github-mcp-server/pkg/raw"
	"github.com/github/github-mcp-server/pkg/toolsets"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

type GetClientFn func(context.Context) (*github.Client, error)
type GetGQLClientFn func(context.Context) (*githubv4.Client, error)

var DefaultTools = []string{"all"}

func DefaultToolsetGroup(readOnly bool, getClient GetClientFn, getGQLClient GetGQLClientFn, getRawClient raw.GetRawClientFn, t translations.TranslationHelperFunc) *toolsets.ToolsetGroup {
	tsg := toolsets.NewToolsetGroup(readOnly)

	// Define all available features with their default state (disabled)
	// Create toolsets
	repos := toolsets.NewToolset("repos", "GitHub Repository related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchRepositories(getClient, t)),
			toolsets.NewServerTool(GetFileContents(getClient, getRawClient, t)),
			toolsets.NewServerTool(ListCommits(getClient, t)),
			toolsets.NewServerTool(SearchCode(getClient, t)),
			toolsets.NewServerTool(GetCommit(getClient, t)),
			toolsets.NewServerTool(ListBranches(getClient, t)),
			toolsets.NewServerTool(ListTags(getClient, t)),
			toolsets.NewServerTool(GetTag(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(CreateOrUpdateFile(getClient, t)),
			toolsets.NewServerTool(CreateRepository(getClient, t)),
			toolsets.NewServerTool(ForkRepository(getClient, t)),
			toolsets.NewServerTool(CreateBranch(getClient, t)),
			toolsets.NewServerTool(PushFiles(getClient, t)),
			toolsets.NewServerTool(DeleteFile(getClient, t)),
		).
		AddResourceTemplates(
			toolsets.NewServerResourceTemplate(GetRepositoryResourceContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceBranchContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceCommitContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourceTagContent(getClient, getRawClient, t)),
			toolsets.NewServerResourceTemplate(GetRepositoryResourcePrContent(getClient, getRawClient, t)),
		)
	issues := toolsets.NewToolset("issues", "GitHub Issues related tools").
		AddReadTools(
			toolsets.NewServerTool(GetIssue(getClient, t)),
			toolsets.NewServerTool(SearchIssues(getClient, t)),
			toolsets.NewServerTool(ListIssues(getClient, t)),
			toolsets.NewServerTool(GetIssueComments(getClient, t)),
			toolsets.NewServerTool(ListSubIssues(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(CreateIssue(getClient, t)),
			toolsets.NewServerTool(AddIssueComment(getClient, t)),
			toolsets.NewServerTool(UpdateIssue(getClient, t)),
			toolsets.NewServerTool(AssignCopilotToIssue(getGQLClient, t)),
			toolsets.NewServerTool(AddSubIssue(getClient, t)),
			toolsets.NewServerTool(RemoveSubIssue(getClient, t)),
			toolsets.NewServerTool(ReprioritizeSubIssue(getClient, t)),
		).AddPrompts(toolsets.NewServerPrompt(AssignCodingAgentPrompt(t)))
	users := toolsets.NewToolset("users", "GitHub User related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchUsers(getClient, t)),
		)
	orgs := toolsets.NewToolset("orgs", "GitHub Organization related tools").
		AddReadTools(
			toolsets.NewServerTool(SearchOrgs(getClient, t)),
		)
	pullRequests := toolsets.NewToolset("pull_requests", "GitHub Pull Request related tools").
		AddReadTools(
			toolsets.NewServerTool(GetPullRequest(getClient, t)),
			toolsets.NewServerTool(ListPullRequests(getClient, t)),
			toolsets.NewServerTool(GetPullRequestFiles(getClient, t)),
			toolsets.NewServerTool(SearchPullRequests(getClient, t)),
			toolsets.NewServerTool(GetPullRequestStatus(getClient, t)),
			toolsets.NewServerTool(GetPullRequestComments(getClient, t)),
			toolsets.NewServerTool(GetPullRequestReviews(getClient, t)),
			toolsets.NewServerTool(GetPullRequestDiff(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(MergePullRequest(getClient, t)),
			toolsets.NewServerTool(UpdatePullRequestBranch(getClient, t)),
			toolsets.NewServerTool(CreatePullRequest(getClient, t)),
			toolsets.NewServerTool(UpdatePullRequest(getClient, t)),
			toolsets.NewServerTool(RequestCopilotReview(getClient, t)),

			// Reviews
			toolsets.NewServerTool(CreateAndSubmitPullRequestReview(getGQLClient, t)),
			toolsets.NewServerTool(CreatePendingPullRequestReview(getGQLClient, t)),
			toolsets.NewServerTool(AddCommentToPendingReview(getGQLClient, t)),
			toolsets.NewServerTool(SubmitPendingPullRequestReview(getGQLClient, t)),
			toolsets.NewServerTool(DeletePendingPullRequestReview(getGQLClient, t)),
		)
	codeSecurity := toolsets.NewToolset("code_security", "Code security related tools, such as GitHub Code Scanning").
		AddReadTools(
			toolsets.NewServerTool(GetCodeScanningAlert(getClient, t)),
			toolsets.NewServerTool(ListCodeScanningAlerts(getClient, t)),
		)
	secretProtection := toolsets.NewToolset("secret_protection", "Secret protection related tools, such as GitHub Secret Scanning").
		AddReadTools(
			toolsets.NewServerTool(GetSecretScanningAlert(getClient, t)),
			toolsets.NewServerTool(ListSecretScanningAlerts(getClient, t)),
		)
	dependabot := toolsets.NewToolset("dependabot", "Dependabot tools").
		AddReadTools(
			toolsets.NewServerTool(GetDependabotAlert(getClient, t)),
			toolsets.NewServerTool(ListDependabotAlerts(getClient, t)),
		)

	notifications := toolsets.NewToolset("notifications", "GitHub Notifications related tools").
		AddReadTools(
			toolsets.NewServerTool(ListNotifications(getClient, t)),
			toolsets.NewServerTool(GetNotificationDetails(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(DismissNotification(getClient, t)),
			toolsets.NewServerTool(MarkAllNotificationsRead(getClient, t)),
			toolsets.NewServerTool(ManageNotificationSubscription(getClient, t)),
			toolsets.NewServerTool(ManageRepositoryNotificationSubscription(getClient, t)),
		)

	discussions := toolsets.NewToolset("discussions", "GitHub Discussions related tools").
		AddReadTools(
			toolsets.NewServerTool(ListDiscussions(getGQLClient, t)),
			toolsets.NewServerTool(GetDiscussion(getGQLClient, t)),
			toolsets.NewServerTool(GetDiscussionComments(getGQLClient, t)),
			toolsets.NewServerTool(ListDiscussionCategories(getGQLClient, t)),
		)

	actions := toolsets.NewToolset("actions", "GitHub Actions workflows and CI/CD operations").
		AddReadTools(
			toolsets.NewServerTool(ListWorkflows(getClient, t)),
			toolsets.NewServerTool(ListWorkflowRuns(getClient, t)),
			toolsets.NewServerTool(GetWorkflowRun(getClient, t)),
			toolsets.NewServerTool(GetWorkflowRunLogs(getClient, t)),
			toolsets.NewServerTool(ListWorkflowJobs(getClient, t)),
			toolsets.NewServerTool(GetJobLogs(getClient, t)),
			toolsets.NewServerTool(ListWorkflowRunArtifacts(getClient, t)),
			toolsets.NewServerTool(DownloadWorkflowRunArtifact(getClient, t)),
			toolsets.NewServerTool(GetWorkflowRunUsage(getClient, t)),
		).
		AddWriteTools(
			toolsets.NewServerTool(RunWorkflow(getClient, t)),
			toolsets.NewServerTool(RerunWorkflowRun(getClient, t)),
			toolsets.NewServerTool(RerunFailedJobs(getClient, t)),
			toolsets.NewServerTool(CancelWorkflowRun(getClient, t)),
			toolsets.NewServerTool(DeleteWorkflowRunLogs(getClient, t)),
		)

	// Keep experiments alive so the system doesn't error out when it's always enabled
	experiments := toolsets.NewToolset("experiments", "Experimental features that are not considered stable yet")

	contextTools := toolsets.NewToolset("context", "Tools that provide context about the current user and GitHub context you are operating in").
		AddReadTools(
			toolsets.NewServerTool(GetMe(getClient, t)),
		)

	// Add toolsets to the group
	tsg.AddToolset(contextTools)
	tsg.AddToolset(repos)
	tsg.AddToolset(issues)
	tsg.AddToolset(orgs)
	tsg.AddToolset(users)
	tsg.AddToolset(pullRequests)
	tsg.AddToolset(actions)
	tsg.AddToolset(codeSecurity)
	tsg.AddToolset(secretProtection)
	tsg.AddToolset(dependabot)
	tsg.AddToolset(notifications)
	tsg.AddToolset(experiments)
	tsg.AddToolset(discussions)

	return tsg
}

// InitDynamicToolset creates a dynamic toolset that can be used to enable other toolsets, and so requires the server and toolset group as arguments
func InitDynamicToolset(s *server.MCPServer, tsg *toolsets.ToolsetGroup, t translations.TranslationHelperFunc) *toolsets.Toolset {
	// Create a new dynamic toolset
	// Need to add the dynamic toolset last so it can be used to enable other toolsets
	dynamicToolSelection := toolsets.NewToolset("dynamic", "Discover GitHub MCP tools that can help achieve tasks by enabling additional sets of tools, you can control the enablement of any toolset to access its tools when this toolset is enabled.").
		AddReadTools(
			toolsets.NewServerTool(ListAvailableToolsets(tsg, t)),
			toolsets.NewServerTool(GetToolsetsTools(tsg, t)),
			toolsets.NewServerTool(EnableToolset(s, tsg, t)),
		)

	dynamicToolSelection.Enabled = true
	return dynamicToolSelection
}

// ToBoolPtr converts a bool to a *bool pointer.
func ToBoolPtr(b bool) *bool {
	return &b
}

// ToStringPtr converts a string to a *string pointer.
// Returns nil if the string is empty.
func ToStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

```

Answer with one JSON object and nothing else:
{"better": "x" | "y" | "tie", "reasons": ["...", "..."]}
Give two to five short reasons, each pointing at concrete files or types.
