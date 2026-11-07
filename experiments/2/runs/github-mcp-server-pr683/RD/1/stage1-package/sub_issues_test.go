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
