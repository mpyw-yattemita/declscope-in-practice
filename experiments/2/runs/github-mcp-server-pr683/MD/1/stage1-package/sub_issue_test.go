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
