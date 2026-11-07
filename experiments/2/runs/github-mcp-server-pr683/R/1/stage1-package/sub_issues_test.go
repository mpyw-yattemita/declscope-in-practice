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
	name        string
	handlers    []mock.MockBackendOption
	args        map[string]any
	expectedErr string // prefix of the tool error text; empty means success
	expected    any    // value the successful result must equal when decoded
}

func runSubIssueTests(t *testing.T, newTool func(GetClientFn, translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc), tests []subIssueTestCase) {
	t.Helper()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := github.NewClient(mock.NewMockedHTTPClient(tc.handlers...))
			_, handler := newTool(stubGetClientFn(client), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.args))
			require.NoError(t, err)

			if tc.expectedErr != "" {
				text := getErrorResult(t, result).Text
				assert.Truef(t, strings.HasPrefix(text, tc.expectedErr), "error %q does not start with %q", text, tc.expectedErr)
				return
			}

			require.False(t, result.IsError, getTextResult(t, result).Text)
			expected, err := json.Marshal(tc.expected)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected), getTextResult(t, result).Text)
		})
	}
}

// noRequest fails the test if the tool calls the API.
func noRequest(t *testing.T) http.HandlerFunc {
	return func(_ http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
	}
}

var (
	mockSubIssue = &github.SubIssue{
		ID:     github.Ptr(int64(123)),
		Number: github.Ptr(43),
		Title:  github.Ptr("Sub-issue"),
	}
	mockParentIssue = &github.Issue{
		ID:     github.Ptr(int64(100)),
		Number: github.Ptr(42),
		Title:  github.Ptr("Parent issue"),
	}
	notFoundBody = map[string]string{"message": "Not Found"}
)

func Test_ListSubIssues(t *testing.T) {
	tool, _ := ListSubIssues(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	endpoint := mock.GetReposIssuesSubIssuesByOwnerByRepoByIssueNumber
	path := "/repos/owner/repo/issues/42/sub_issues"
	subIssues := []*github.SubIssue{mockSubIssue}

	runSubIssueTests(t, ListSubIssues, []subIssueTestCase{
		{
			name: "default pagination",
			handlers: []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint,
				expect(t, expectations{path: path, queryParams: map[string]string{"page": "1", "per_page": "30"}}).
					andThen(mockResponse(t, http.StatusOK, subIssues)))},
			args:     map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42)},
			expected: subIssues,
		},
		{
			name: "explicit pagination",
			handlers: []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint,
				expect(t, expectations{path: path, queryParams: map[string]string{"page": "2", "per_page": "10"}}).
					andThen(mockResponse(t, http.StatusOK, subIssues)))},
			args:     map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42), "page": float64(2), "per_page": float64(10)},
			expected: subIssues,
		},
		{
			name:        "missing issue_number",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, noRequest(t))},
			args:        map[string]any{"owner": "owner", "repo": "repo"},
			expectedErr: "missing required parameter: issue_number",
		},
		{
			name:        "API error",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, mockResponse(t, http.StatusNotFound, notFoundBody))},
			args:        map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42)},
			expectedErr: "failed to list sub-issues",
		},
		{
			name:        "unexpected status",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, mockResponse(t, http.StatusNonAuthoritativeInfo, subIssues))},
			args:        map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42)},
			expectedErr: "failed to list sub-issues: unexpected status 203",
		},
	})
}

func Test_AddSubIssue(t *testing.T) {
	tool, _ := AddSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	endpoint := mock.PostReposIssuesSubIssuesByOwnerByRepoByIssueNumber
	path := "/repos/owner/repo/issues/42/sub_issues"

	runSubIssueTests(t, AddSubIssue, []subIssueTestCase{
		{
			name: "add",
			handlers: []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint,
				expect(t, expectations{path: path, requestBody: map[string]any{"sub_issue_id": float64(123)}}).
					andThen(mockResponse(t, http.StatusCreated, mockSubIssue)))},
			args:     map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42), "sub_issue_id": float64(123)},
			expected: mockSubIssue,
		},
		{
			name: "add with replace_parent",
			handlers: []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint,
				expect(t, expectations{path: path, requestBody: map[string]any{"sub_issue_id": float64(123), "replace_parent": true}}).
					andThen(mockResponse(t, http.StatusCreated, mockSubIssue)))},
			args:     map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42), "sub_issue_id": float64(123), "replace_parent": true},
			expected: mockSubIssue,
		},
		{
			name:        "missing sub_issue_id",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, noRequest(t))},
			args:        map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42)},
			expectedErr: "missing required parameter: sub_issue_id",
		},
		{
			name:        "missing owner",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, noRequest(t))},
			args:        map[string]any{"repo": "repo", "issue_number": float64(42), "sub_issue_id": float64(123)},
			expectedErr: "missing required parameter: owner",
		},
		{
			name: "API error",
			handlers: []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint,
				mockResponse(t, http.StatusUnprocessableEntity, map[string]string{"message": "Validation Failed"}))},
			args:        map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42), "sub_issue_id": float64(123)},
			expectedErr: "failed to add sub-issue",
		},
		{
			name:        "unexpected status",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, mockResponse(t, http.StatusOK, mockSubIssue))},
			args:        map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42), "sub_issue_id": float64(123)},
			expectedErr: "failed to add sub-issue: unexpected status 200",
		},
	})
}

func Test_RemoveSubIssue(t *testing.T) {
	tool, _ := RemoveSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	endpoint := mock.DeleteReposIssuesSubIssueByOwnerByRepoByIssueNumber
	args := map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42), "sub_issue_id": float64(123)}

	runSubIssueTests(t, RemoveSubIssue, []subIssueTestCase{
		{
			// The path is singular "sub_issue"; go-github's Remove gets this wrong.
			name: "remove",
			handlers: []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint,
				expect(t, expectations{path: "/repos/owner/repo/issues/42/sub_issue", requestBody: map[string]any{"sub_issue_id": float64(123)}}).
					andThen(mockResponse(t, http.StatusOK, mockParentIssue)))},
			args:     args,
			expected: mockParentIssue,
		},
		{
			name:        "missing sub_issue_id",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, noRequest(t))},
			args:        map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42)},
			expectedErr: "missing required parameter: sub_issue_id",
		},
		{
			name:        "API error",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, mockResponse(t, http.StatusNotFound, notFoundBody))},
			args:        args,
			expectedErr: "failed to remove sub-issue",
		},
		{
			name:        "unexpected status",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, mockResponse(t, http.StatusNoContent, ""))},
			args:        args,
			expectedErr: "failed to remove sub-issue: unexpected status 204",
		},
	})
}

func Test_ReprioritizeSubIssue(t *testing.T) {
	tool, _ := ReprioritizeSubIssue(stubGetClientFn(github.NewClient(nil)), translations.NullTranslationHelper)
	require.NoError(t, toolsnaps.Test(tool.Name, tool))

	endpoint := mock.PatchReposIssuesSubIssuesPriorityByOwnerByRepoByIssueNumber
	path := "/repos/owner/repo/issues/42/sub_issues/priority"
	args := func(extra map[string]any) map[string]any {
		a := map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42), "sub_issue_id": float64(123)}
		for k, v := range extra {
			a[k] = v
		}
		return a
	}

	runSubIssueTests(t, ReprioritizeSubIssue, []subIssueTestCase{
		{
			name: "after_id",
			handlers: []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint,
				expect(t, expectations{path: path, requestBody: map[string]any{"sub_issue_id": float64(123), "after_id": float64(456)}}).
					andThen(mockResponse(t, http.StatusOK, mockSubIssue)))},
			args:     args(map[string]any{"after_id": float64(456)}),
			expected: mockSubIssue,
		},
		{
			name: "before_id",
			handlers: []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint,
				expect(t, expectations{path: path, requestBody: map[string]any{"sub_issue_id": float64(123), "before_id": float64(789)}}).
					andThen(mockResponse(t, http.StatusOK, mockSubIssue)))},
			args:     args(map[string]any{"before_id": float64(789)}),
			expected: mockSubIssue,
		},
		{
			name:        "neither after_id nor before_id",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, noRequest(t))},
			args:        args(nil),
			expectedErr: "either after_id or before_id must be specified",
		},
		{
			name:        "both after_id and before_id",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, noRequest(t))},
			args:        args(map[string]any{"after_id": float64(456), "before_id": float64(789)}),
			expectedErr: "only one of after_id or before_id should be specified, not both",
		},
		{
			name:        "missing sub_issue_id",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, noRequest(t))},
			args:        map[string]any{"owner": "owner", "repo": "repo", "issue_number": float64(42), "after_id": float64(456)},
			expectedErr: "missing required parameter: sub_issue_id",
		},
		{
			name:        "API error",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, mockResponse(t, http.StatusNotFound, notFoundBody))},
			args:        args(map[string]any{"after_id": float64(456)}),
			expectedErr: "failed to reprioritize sub-issue",
		},
		{
			name:        "unexpected status",
			handlers:    []mock.MockBackendOption{mock.WithRequestMatchHandler(endpoint, mockResponse(t, http.StatusCreated, mockSubIssue))},
			args:        args(map[string]any{"after_id": float64(456)}),
			expectedErr: "failed to reprioritize sub-issue: unexpected status 201",
		},
	})
}
