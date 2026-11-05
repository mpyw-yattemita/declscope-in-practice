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
			target, err := requiredSubIssueTarget(request)
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
			subIssue, resp, err := client.SubIssue.Add(ctx, target.owner, target.repo, int64(target.issueNumber), subIssueRequest)
			return subIssueResult(ctx, "failed to add sub-issue", http.StatusCreated, subIssue, resp, err)
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
			target, err := requiredSubIssueTarget(request)
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
			subIssues, resp, err := client.SubIssue.ListByIssue(ctx, target.owner, target.repo, int64(target.issueNumber), opts)
			return subIssueResult(ctx, "failed to list sub-issues", http.StatusOK, subIssues, resp, err)
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
			target, err := requiredSubIssueTarget(request)
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

			// go-github's SubIssue.Remove sends the request to ".../sub_issues" instead of
			// ".../sub_issue" (google/go-github#3613), so build the request here.
			url := fmt.Sprintf("repos/%s/%s/issues/%d/sub_issue", target.owner, target.repo, target.issueNumber)
			req, err := client.NewRequest(http.MethodDelete, url, github.SubIssueRequest{
				SubIssueID: int64(subIssueID),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create request: %w", err)
			}
			var parent github.Issue
			resp, err := client.Do(ctx, req, &parent)
			return subIssueResult(ctx, "failed to remove sub-issue", http.StatusOK, &parent, resp, err)
		}
}

// ReprioritizeSubIssue creates a tool to move a sub-issue to a different position in its parent's list.
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
			target, err := requiredSubIssueTarget(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			subIssueID, err := RequiredInt(request, "sub_issue_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			afterID, hasAfter, err := OptionalParamOK[float64](request, "after_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			beforeID, hasBefore, err := OptionalParamOK[float64](request, "before_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			switch {
			case !hasAfter && !hasBefore:
				return mcp.NewToolResultError("either after_id or before_id must be specified"), nil
			case hasAfter && hasBefore:
				return mcp.NewToolResultError("only one of after_id or before_id should be specified, not both"), nil
			}

			subIssueRequest := github.SubIssueRequest{
				SubIssueID: int64(subIssueID),
			}
			if hasAfter {
				subIssueRequest.AfterID = github.Ptr(int64(afterID))
			} else {
				subIssueRequest.BeforeID = github.Ptr(int64(beforeID))
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			subIssue, resp, err := client.SubIssue.Reprioritize(ctx, target.owner, target.repo, int64(target.issueNumber), subIssueRequest)
			return subIssueResult(ctx, "failed to reprioritize sub-issue", http.StatusOK, subIssue, resp, err)
		}
}

// subIssueTarget is the parent issue that a sub-issue tool acts on.
type subIssueTarget struct {
	owner       string
	repo        string
	issueNumber int
}

// requiredSubIssueTarget reads the required "owner", "repo" and "issue_number" parameters.
func requiredSubIssueTarget(request mcp.CallToolRequest) (subIssueTarget, error) {
	owner, err := RequiredParam[string](request, "owner")
	if err != nil {
		return subIssueTarget{}, err
	}
	repo, err := RequiredParam[string](request, "repo")
	if err != nil {
		return subIssueTarget{}, err
	}
	issueNumber, err := RequiredInt(request, "issue_number")
	if err != nil {
		return subIssueTarget{}, err
	}
	return subIssueTarget{owner: owner, repo: repo, issueNumber: issueNumber}, nil
}

// subIssueResult turns the outcome of a sub-issue API call into a tool result.
// An API error, or a success status other than wantStatus, gives an error result that starts with message.
// Otherwise the result is v as JSON.
func subIssueResult(ctx context.Context, message string, wantStatus int, v any, resp *github.Response, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return ghErrors.NewGitHubAPIErrorResponse(ctx, message, resp, err), nil
	}
	// The client has already read and closed the body, so only the status is left to report.
	if resp.StatusCode != wantStatus {
		return mcp.NewToolResultError(fmt.Sprintf("%s: unexpected status %d %s", message, resp.StatusCode, http.StatusText(resp.StatusCode))), nil
	}

	return MarshalledTextResult(v), nil
}
