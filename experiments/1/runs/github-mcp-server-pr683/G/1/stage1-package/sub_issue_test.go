package github

import (
	"context"
	"encoding/json"
	"net/http"
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

func Test_AddSubIssue(t *testing.T) {
	tool, _ := AddSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	mockSubIssue := &github.SubIssue{
		ID:     github.Ptr(int64(123)),
		Number: github.Ptr(42),
		Title:  github.Ptr("Sub-issue"),
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]any
		expectedErrMsg string
	}{
		{
			name: "successful add",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectRequestBody(t, map[string]any{
						"sub_issue_id": float64(123),
					}).andThen(mockResponse(t, http.StatusCreated, mockSubIssue)),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
			},
		},
		{
			name: "successful add with replace_parent",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectRequestBody(t, map[string]any{
						"sub_issue_id":   float64(123),
						"replace_parent": true,
					}).andThen(mockResponse(t, http.StatusCreated, mockSubIssue)),
				),
			),
			requestArgs: map[string]any{
				"owner":          "owner",
				"repo":           "repo",
				"issue_number":   float64(1),
				"sub_issue_id":   float64(123),
				"replace_parent": true,
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
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
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
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "failed to add sub-issue: unexpected status 200",
		},
		{
			name:         "missing sub_issue_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
			},
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, handler := AddSubIssue(stubGetClientFn(github.NewClient(tc.mockedClient)), translations.NullTranslationHelper)
			result := callSubIssueHandler(t, handler, tc.requestArgs)

			if tc.expectedErrMsg != "" {
				require.True(t, result.IsError)
				assert.Contains(t, getErrorResult(t, result).Text, tc.expectedErrMsg)
				return
			}

			require.False(t, result.IsError)
			var returned github.SubIssue
			require.NoError(t, json.Unmarshal([]byte(getTextResult(t, result).Text), &returned))
			assert.Equal(t, mockSubIssue.ID, returned.ID)
			assert.Equal(t, mockSubIssue.Number, returned.Number)
		})
	}
}

func Test_ListSubIssues(t *testing.T) {
	tool, _ := ListSubIssues(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	mockSubIssues := []*github.SubIssue{
		{ID: github.Ptr(int64(1)), Number: github.Ptr(10), Title: github.Ptr("First")},
		{ID: github.Ptr(int64(2)), Number: github.Ptr(11), Title: github.Ptr("Second")},
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]any
		expectedErrMsg string
	}{
		{
			name: "successful list with default pagination",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectQueryParams(t, map[string]string{
						"page":     "1",
						"per_page": "30",
					}).andThen(mockResponse(t, http.StatusOK, mockSubIssues)),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
			},
		},
		{
			name: "successful list with pagination",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					expectQueryParams(t, map[string]string{
						"page":     "2",
						"per_page": "5",
					}).andThen(mockResponse(t, http.StatusOK, mockSubIssues)),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"page":         float64(2),
				"per_page":     float64(5),
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
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(999),
			},
			expectedErrMsg: "failed to list sub-issues",
		},
		{
			name:         "missing issue_number",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"owner": "owner",
				"repo":  "repo",
			},
			expectedErrMsg: "missing required parameter: issue_number",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, handler := ListSubIssues(stubGetClientFn(github.NewClient(tc.mockedClient)), translations.NullTranslationHelper)
			result := callSubIssueHandler(t, handler, tc.requestArgs)

			if tc.expectedErrMsg != "" {
				require.True(t, result.IsError)
				assert.Contains(t, getErrorResult(t, result).Text, tc.expectedErrMsg)
				return
			}

			require.False(t, result.IsError)
			var returned []*github.SubIssue
			require.NoError(t, json.Unmarshal([]byte(getTextResult(t, result).Text), &returned))
			require.Len(t, returned, len(mockSubIssues))
			for i := range returned {
				assert.Equal(t, mockSubIssues[i].ID, returned[i].ID)
				assert.Equal(t, mockSubIssues[i].Title, returned[i].Title)
			}
		})
	}
}

func Test_RemoveSubIssue(t *testing.T) {
	tool, _ := RemoveSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	mockParent := &github.Issue{
		ID:     github.Ptr(int64(999)),
		Number: github.Ptr(1),
		Title:  github.Ptr("Parent"),
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]any
		expectedErrMsg string
	}{
		{
			name: "successful remove",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					expectPath(t, "/repos/owner/repo/issues/1/sub_issue").andThen(
						expectRequestBody(t, map[string]any{
							"sub_issue_id": float64(123),
						}).andThen(mockResponse(t, http.StatusOK, mockParent)),
					),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
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
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "failed to remove sub-issue",
		},
		{
			name: "unexpected status",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					mockResponse(t, http.StatusNoContent, nil),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "failed to remove sub-issue: unexpected status 204",
		},
		{
			name:         "missing sub_issue_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
			},
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, handler := RemoveSubIssue(stubGetClientFn(github.NewClient(tc.mockedClient)), translations.NullTranslationHelper)
			result := callSubIssueHandler(t, handler, tc.requestArgs)

			if tc.expectedErrMsg != "" {
				require.True(t, result.IsError)
				assert.Contains(t, getErrorResult(t, result).Text, tc.expectedErrMsg)
				return
			}

			require.False(t, result.IsError)
			var returned github.Issue
			require.NoError(t, json.Unmarshal([]byte(getTextResult(t, result).Text), &returned))
			assert.Equal(t, mockParent.ID, returned.ID)
			assert.Equal(t, mockParent.Number, returned.Number)
		})
	}
}

func Test_ReprioritizeSubIssue(t *testing.T) {
	tool, _ := ReprioritizeSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	mockSubIssue := &github.SubIssue{
		ID:     github.Ptr(int64(123)),
		Number: github.Ptr(42),
		Title:  github.Ptr("Sub-issue"),
	}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]any
		expectedErrMsg string
	}{
		{
			name: "successful reprioritize after",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					expectRequestBody(t, map[string]any{
						"sub_issue_id": float64(123),
						"after_id":     float64(456),
					}).andThen(mockResponse(t, http.StatusOK, mockSubIssue)),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
				"after_id":     float64(456),
			},
		},
		{
			name: "successful reprioritize before",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					expectRequestBody(t, map[string]any{
						"sub_issue_id": float64(123),
						"before_id":    float64(789),
					}).andThen(mockResponse(t, http.StatusOK, mockSubIssue)),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
				"before_id":    float64(789),
			},
		},
		{
			name:         "neither after_id nor before_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
			},
			expectedErrMsg: "either after_id or before_id must be specified",
		},
		{
			name:         "both after_id and before_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
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
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
				"after_id":     float64(456),
			},
			expectedErrMsg: "failed to reprioritize sub-issue",
		},
		{
			name:         "missing owner",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"repo":         "repo",
				"issue_number": float64(1),
				"sub_issue_id": float64(123),
				"after_id":     float64(456),
			},
			expectedErrMsg: "missing required parameter: owner",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, handler := ReprioritizeSubIssue(stubGetClientFn(github.NewClient(tc.mockedClient)), translations.NullTranslationHelper)
			result := callSubIssueHandler(t, handler, tc.requestArgs)

			if tc.expectedErrMsg != "" {
				require.True(t, result.IsError)
				assert.Contains(t, getErrorResult(t, result).Text, tc.expectedErrMsg)
				return
			}

			require.False(t, result.IsError)
			var returned github.SubIssue
			require.NoError(t, json.Unmarshal([]byte(getTextResult(t, result).Text), &returned))
			assert.Equal(t, mockSubIssue.ID, returned.ID)
		})
	}
}

// callSubIssueHandler calls a sub-issue tool handler with args, and fails the
// test if it returns a Go error, since every failure must be a tool error.
func callSubIssueHandler(t *testing.T, handler server.ToolHandlerFunc, args map[string]any) *mcp.CallToolResult {
	t.Helper()
	result, err := handler(context.Background(), createMCPRequest(args))
	require.NoError(t, err)
	require.NotNil(t, result)
	return result
}
