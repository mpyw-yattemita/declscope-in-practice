package github

import (
	"context"
	"fmt"
	"net/http"

	ghErrors "github.com/github/github-mcp-server/pkg/errors"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Tools for the sub-issue REST endpoints.
// https://docs.github.com/en/rest/issues/sub-issues

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
			// These are not WithPagination: the tool uses "per_page" rather than "perPage".
			mcp.WithNumber("page",
				mcp.Description("Page number for pagination (default: 1)"),
			),
			mcp.WithNumber("per_page",
				mcp.Description("Number of results per page (max 100, default: 30)"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			const failure = "failed to list sub-issues"

			owner, repo, issueNumber, err := subIssueParentParams(request)
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
				return ghErrors.NewGitHubAPIErrorResponse(ctx, failure, resp, err), nil
			}
			if resp.StatusCode != http.StatusOK {
				return unexpectedStatusResult(failure, resp), nil
			}

			return MarshalledTextResult(subIssues), nil
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
			withSubIssueParentParams(),
			mcp.WithNumber("sub_issue_id",
				mcp.Required(),
				mcp.Description("The ID of the sub-issue to add. ID is not the same as issue number"),
			),
			mcp.WithBoolean("replace_parent",
				mcp.Description("When true, replaces the sub-issue's current parent issue"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			const failure = "failed to add sub-issue"

			owner, repo, issueNumber, err := subIssueParentParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := requiredInt64(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			replaceParent, err := OptionalParam[bool](request, "replace_parent")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			body := github.SubIssueRequest{SubIssueID: subIssueID}
			if replaceParent {
				body.ReplaceParent = github.Ptr(true)
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			subIssue, resp, err := client.SubIssue.Add(ctx, owner, repo, int64(issueNumber), body)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, failure, resp, err), nil
			}
			if resp.StatusCode != http.StatusCreated {
				return unexpectedStatusResult(failure, resp), nil
			}

			return MarshalledTextResult(subIssue), nil
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
			withSubIssueParentParams(),
			mcp.WithNumber("sub_issue_id",
				mcp.Required(),
				mcp.Description("The ID of the sub-issue to remove. ID is not the same as issue number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			const failure = "failed to remove sub-issue"

			owner, repo, issueNumber, err := subIssueParentParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := requiredInt64(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			// SubIssueService.Remove sends the request to ".../sub_issues" instead of
			// ".../sub_issue" (google/go-github#3613), so build the request here.
			u := fmt.Sprintf("repos/%s/%s/issues/%d/sub_issue", owner, repo, issueNumber)
			req, err := client.NewRequest(http.MethodDelete, u, github.SubIssueRequest{SubIssueID: subIssueID})
			if err != nil {
				return nil, fmt.Errorf("failed to create request: %w", err)
			}
			// The response body is the parent issue.
			var parent github.Issue
			resp, err := client.Do(ctx, req, &parent)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, failure, resp, err), nil
			}
			if resp.StatusCode != http.StatusOK {
				return unexpectedStatusResult(failure, resp), nil
			}

			return MarshalledTextResult(&parent), nil
		}
}

// ReprioritizeSubIssue creates a tool to move a sub-issue to a different
// position in the parent issue's sub-issue list.
func ReprioritizeSubIssue(getClient GetClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("reprioritize_sub_issue",
			mcp.WithDescription(t("TOOL_REPRIORITIZE_SUB_ISSUE_DESCRIPTION", "Reprioritize a sub-issue to a different position in the parent issue's sub-issue list.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_REPRIORITIZE_SUB_ISSUE_USER_TITLE", "Reprioritize sub-issue"),
				ReadOnlyHint: ToBoolPtr(false),
			}),
			withSubIssueParentParams(),
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
			const failure = "failed to reprioritize sub-issue"

			owner, repo, issueNumber, err := subIssueParentParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := requiredInt64(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			afterID, err := optionalInt64Param(request, "after_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			beforeID, err := optionalInt64Param(request, "before_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			body := github.SubIssueRequest{SubIssueID: subIssueID}
			switch {
			case afterID == 0 && beforeID == 0:
				return mcp.NewToolResultError("either after_id or before_id must be specified"), nil
			case afterID != 0 && beforeID != 0:
				return mcp.NewToolResultError("only one of after_id or before_id should be specified, not both"), nil
			case afterID != 0:
				body.AfterID = github.Ptr(afterID)
			default:
				body.BeforeID = github.Ptr(beforeID)
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			subIssue, resp, err := client.SubIssue.Reprioritize(ctx, owner, repo, int64(issueNumber), body)
			if err != nil {
				return ghErrors.NewGitHubAPIErrorResponse(ctx, failure, resp, err), nil
			}
			if resp.StatusCode != http.StatusOK {
				return unexpectedStatusResult(failure, resp), nil
			}

			return MarshalledTextResult(subIssue), nil
		}
}

// withSubIssueParentParams adds the "owner", "repo" and "issue_number" parameters
// that identify the parent issue for the sub-issue write tools.
func withSubIssueParentParams() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("Repository owner"),
		)(tool)
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("Repository name"),
		)(tool)
		mcp.WithNumber("issue_number",
			mcp.Required(),
			mcp.Description("The number of the parent issue"),
		)(tool)
	}
}

// subIssueParentParams reads the "owner", "repo" and "issue_number" parameters.
func subIssueParentParams(r mcp.CallToolRequest) (owner, repo string, issueNumber int, err error) {
	if owner, err = RequiredParam[string](r, "owner"); err != nil {
		return "", "", 0, err
	}
	if repo, err = RequiredParam[string](r, "repo"); err != nil {
		return "", "", 0, err
	}
	if issueNumber, err = RequiredInt(r, "issue_number"); err != nil {
		return "", "", 0, err
	}
	return owner, repo, issueNumber, nil
}

// unexpectedStatusResult reports a successful (2xx) response whose status is not
// the one the endpoint documents. The client has already read and closed the
// body by this point, so only the status is reported.
func unexpectedStatusResult(message string, resp *github.Response) *mcp.CallToolResult {
	return mcp.NewToolResultError(fmt.Sprintf("%s: unexpected status %d", message, resp.StatusCode))
}
