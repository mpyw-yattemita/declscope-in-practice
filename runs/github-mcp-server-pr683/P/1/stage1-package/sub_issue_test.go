package github

import (
	"context"
	"encoding/json"
	"io"
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
	tool, _ := AddSubIssue(subIssueClientFn(nil), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	mockSubIssue := &github.Issue{
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
			name: "successful add with replace_parent",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					subIssueHandler(t, http.StatusCreated, mockSubIssue, map[string]any{
						"sub_issue_id":   float64(123),
						"replace_parent": true,
					}),
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
			name: "successful add without replace_parent",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					subIssueHandler(t, http.StatusCreated, mockSubIssue, map[string]any{
						"sub_issue_id": float64(123),
					}),
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
			name:         "missing sub_issue_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
			},
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					subIssueHandler(t, http.StatusUnprocessableEntity, map[string]string{"message": "Validation Failed"}, nil),
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
					subIssueHandler(t, http.StatusOK, mockSubIssue, nil),
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, handler := AddSubIssue(subIssueClientFn(tc.mockedClient), translations.NullTranslationHelper)
			text := callSubIssueTool(t, handler, tc.requestArgs, tc.expectedErrMsg)
			if tc.expectedErrMsg != "" {
				return
			}

			var got github.Issue
			require.NoError(t, json.Unmarshal([]byte(text), &got))
			assert.Equal(t, *mockSubIssue.ID, *got.ID)
			assert.Equal(t, *mockSubIssue.Number, *got.Number)
			assert.Equal(t, *mockSubIssue.Title, *got.Title)
		})
	}
}

func Test_ListSubIssues(t *testing.T) {
	tool, _ := ListSubIssues(subIssueClientFn(nil), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	mockSubIssues := []*github.Issue{
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
					subIssueQueryHandler(t, map[string]string{"page": "1", "per_page": "30"}, mockSubIssues),
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
					subIssueQueryHandler(t, map[string]string{"page": "2", "per_page": "10"}, mockSubIssues),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"page":         float64(2),
				"per_page":     float64(10),
			},
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
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber,
					subIssueHandler(t, http.StatusNotFound, map[string]string{"message": "Not Found"}, nil),
				),
			),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
			},
			expectedErrMsg: "failed to list sub-issues",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, handler := ListSubIssues(subIssueClientFn(tc.mockedClient), translations.NullTranslationHelper)
			text := callSubIssueTool(t, handler, tc.requestArgs, tc.expectedErrMsg)
			if tc.expectedErrMsg != "" {
				return
			}

			var got []github.Issue
			require.NoError(t, json.Unmarshal([]byte(text), &got))
			require.Len(t, got, len(mockSubIssues))
			for i := range got {
				assert.Equal(t, *mockSubIssues[i].ID, *got[i].ID)
				assert.Equal(t, *mockSubIssues[i].Title, *got[i].Title)
			}
		})
	}
}

func Test_RemoveSubIssue(t *testing.T) {
	tool, _ := RemoveSubIssue(subIssueClientFn(nil), translations.NullTranslationHelper)
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
					subIssueHandler(t, http.StatusOK, mockParent, map[string]any{
						"sub_issue_id": float64(123),
					}),
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
			name:         "missing sub_issue_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
			},
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber,
					subIssueHandler(t, http.StatusNotFound, map[string]string{"message": "Not Found"}, nil),
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
					subIssueHandler(t, http.StatusNoContent, nil, nil),
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, handler := RemoveSubIssue(subIssueClientFn(tc.mockedClient), translations.NullTranslationHelper)
			text := callSubIssueTool(t, handler, tc.requestArgs, tc.expectedErrMsg)
			if tc.expectedErrMsg != "" {
				return
			}

			var got github.Issue
			require.NoError(t, json.Unmarshal([]byte(text), &got))
			assert.Equal(t, *mockParent.ID, *got.ID)
			assert.Equal(t, *mockParent.Title, *got.Title)
		})
	}
}

func Test_ReprioritizeSubIssue(t *testing.T) {
	tool, _ := ReprioritizeSubIssue(subIssueClientFn(nil), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	mockSubIssue := &github.Issue{
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
			name: "successful reprioritize with after_id",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					subIssueHandler(t, http.StatusOK, mockSubIssue, map[string]any{
						"sub_issue_id": float64(123),
						"after_id":     float64(456),
					}),
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
			name: "successful reprioritize with before_id",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					subIssueHandler(t, http.StatusOK, mockSubIssue, map[string]any{
						"sub_issue_id": float64(123),
						"before_id":    float64(789),
					}),
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
			name:         "missing sub_issue_id",
			mockedClient: mock.NewMockedHTTPClient(),
			requestArgs: map[string]any{
				"owner":        "owner",
				"repo":         "repo",
				"issue_number": float64(1),
				"after_id":     float64(456),
			},
			expectedErrMsg: "missing required parameter: sub_issue_id",
		},
		{
			name: "API error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber,
					subIssueHandler(t, http.StatusUnprocessableEntity, map[string]string{"message": "Validation Failed"}, nil),
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, handler := ReprioritizeSubIssue(subIssueClientFn(tc.mockedClient), translations.NullTranslationHelper)
			text := callSubIssueTool(t, handler, tc.requestArgs, tc.expectedErrMsg)
			if tc.expectedErrMsg != "" {
				return
			}

			var got github.Issue
			require.NoError(t, json.Unmarshal([]byte(text), &got))
			assert.Equal(t, *mockSubIssue.ID, *got.ID)
			assert.Equal(t, *mockSubIssue.Title, *got.Title)
		})
	}
}

// subIssueClientFn returns a GetClientFn whose client sends requests through httpClient.
func subIssueClientFn(httpClient *http.Client) GetClientFn {
	client := github.NewClient(httpClient)
	return func(context.Context) (*github.Client, error) {
		return client, nil
	}
}

// subIssueHandler responds with status and body as JSON. A nil body writes no body.
// When wantBody is not nil, it first checks that the request body is wantBody as JSON.
func subIssueHandler(t *testing.T, status int, body any, wantBody map[string]any) http.HandlerFunc {
	t.Helper()
	return func(w http.ResponseWriter, r *http.Request) {
		if wantBody != nil {
			raw, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			var got map[string]any
			require.NoError(t, json.Unmarshal(raw, &got))
			assert.Equal(t, wantBody, got)
		}
		w.WriteHeader(status)
		if body != nil {
			require.NoError(t, json.NewEncoder(w).Encode(body))
		}
	}
}

// subIssueQueryHandler checks that the request has wantQuery and responds 200 with body as JSON.
func subIssueQueryHandler(t *testing.T, wantQuery map[string]string, body any) http.HandlerFunc {
	t.Helper()
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		for k, v := range wantQuery {
			assert.Equal(t, v, query.Get(k), "query parameter %s", k)
		}
		w.WriteHeader(http.StatusOK)
		require.NoError(t, json.NewEncoder(w).Encode(body))
	}
}

// callSubIssueTool calls handler with args and returns the text of its result.
// When wantErrMsg is not empty, it checks that the result is an error that contains wantErrMsg.
// Otherwise it checks that the result is not an error.
func callSubIssueTool(t *testing.T, handler server.ToolHandlerFunc, args map[string]any, wantErrMsg string) string {
	t.Helper()
	request := mcp.CallToolRequest{}
	request.Params.Arguments = args

	result, err := handler(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Content, 1)
	text, ok := result.Content[0].(mcp.TextContent)
	require.True(t, ok, "expected text content, got %T", result.Content[0])

	if wantErrMsg != "" {
		require.True(t, result.IsError, "expected an error result, got %s", text.Text)
		assert.Contains(t, text.Text, wantErrMsg)
	} else {
		require.False(t, result.IsError, "unexpected error result: %s", text.Text)
	}
	return text.Text
}
