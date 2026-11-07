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

// The GraphQL documents the discussion tools send. They are written out in full
// so that the tests catch any change to what is sent to the API.
const (
	qListDiscussions                  = "query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qListDiscussionsByCategory        = "query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qListDiscussionsOrdered           = "query($after:String$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qListDiscussionsByCategoryOrdered = "query($after:String$categoryId:ID!$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qGetDiscussionComments            = "query($after:String$discussionNumber:Int!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussion(number: $discussionNumber){comments(first: $first, after: $after){nodes{body},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"
	qListDiscussionCategories         = "query($first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussionCategories(first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
)

var (
	discussionsGeneral = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-15T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-03-15T00:00:00Z", "author": map[string]any{"login": "user3"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	discussionsAll = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-15T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 2, "title": "Discussion 2 title", "createdAt": "2023-02-01T00:00:00Z", "updatedAt": "2023-02-15T00:00:00Z", "author": map[string]any{"login": "user2"}, "url": "https://github.com/owner/repo/discussions/2", "category": map[string]any{"name": "Questions"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-03-15T00:00:00Z", "author": map[string]any{"login": "user3"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	mockErrorRepoNotFound = githubv4mock.ErrorResponse("repository not found")
)

func mockPageInfo(hasNext bool, endCursor string) map[string]any {
	return map[string]any{
		"hasNextPage":     hasNext,
		"hasPreviousPage": false,
		"startCursor":     "start",
		"endCursor":       endCursor,
	}
}

func discussionsResponse(nodes []map[string]any, hasNext bool, endCursor string, total int) githubv4mock.GQLResponse {
	return githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes":      nodes,
				"pageInfo":   mockPageInfo(hasNext, endCursor),
				"totalCount": total,
			},
		},
	})
}

type pageInfoOutput struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

func Test_ListDiscussions(t *testing.T) {
	mockClient := githubv4.NewClient(nil)
	// Verify tool definition and schema
	toolDef, _ := ListDiscussions(stubGetGQLClientFn(mockClient), translations.NullTranslationHelper)
	assert.Equal(t, "list_discussions", toolDef.Name)
	assert.NotEmpty(t, toolDef.Description)
	assert.Contains(t, toolDef.InputSchema.Properties, "owner")
	assert.Contains(t, toolDef.InputSchema.Properties, "repo")
	assert.Contains(t, toolDef.InputSchema.Properties, "category")
	assert.Contains(t, toolDef.InputSchema.Properties, "orderBy")
	assert.Contains(t, toolDef.InputSchema.Properties, "direction")
	assert.ElementsMatch(t, []any{"CREATED_AT", "UPDATED_AT"}, toolDef.InputSchema.Properties["orderBy"].(map[string]any)["enum"])
	assert.ElementsMatch(t, []any{"ASC", "DESC"}, toolDef.InputSchema.Properties["direction"].(map[string]any)["enum"])
	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")
	assert.NotContains(t, toolDef.InputSchema.Properties, "page")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo"})

	tests := []struct {
		name              string
		reqParams         map[string]interface{}
		query             string
		vars              map[string]interface{}
		response          githubv4mock.GQLResponse
		expectError       bool
		errContains       string
		expectedCount     int
		expectedPageInfo  pageInfoOutput
		expectedTotal     int
		expectCategoryCol bool
	}{
		{
			name:      "list all discussions with default pagination",
			reqParams: map[string]interface{}{"owner": "owner", "repo": "repo"},
			query:     qListDiscussions,
			vars: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
				"first": float64(30),
				"after": nil,
			},
			response:         discussionsResponse(discussionsAll, true, "cursor-3", 5),
			expectedCount:    3,
			expectedPageInfo: pageInfoOutput{HasNextPage: true, StartCursor: "start", EndCursor: "cursor-3"},
			expectedTotal:    5,
		},
		{
			name: "list discussions with perPage and after",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(2),
				"after":   "cursor-3",
			},
			query: qListDiscussions,
			vars: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
				"first": float64(2),
				"after": "cursor-3",
			},
			response:         discussionsResponse(discussionsGeneral, false, "cursor-5", 5),
			expectedCount:    2,
			expectedPageInfo: pageInfoOutput{HasNextPage: false, StartCursor: "start", EndCursor: "cursor-5"},
			expectedTotal:    5,
		},
		{
			name: "filter by category ID",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
			},
			query: qListDiscussionsByCategory,
			vars: map[string]interface{}{
				"owner":      "owner",
				"repo":       "repo",
				"categoryId": "DIC_kwDOABC123",
				"first":      float64(30),
				"after":      nil,
			},
			response:          discussionsResponse(discussionsGeneral, false, "cursor-2", 2),
			expectedCount:     2,
			expectedPageInfo:  pageInfoOutput{StartCursor: "start", EndCursor: "cursor-2"},
			expectedTotal:     2,
			expectCategoryCol: true,
		},
		{
			name: "order by updated time descending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "UPDATED_AT",
				"direction": "DESC",
			},
			query: qListDiscussionsOrdered,
			vars: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"orderByField":     "UPDATED_AT",
				"orderByDirection": "DESC",
				"first":            float64(30),
				"after":            nil,
			},
			response:         discussionsResponse(discussionsAll, true, "cursor-3", 5),
			expectedCount:    3,
			expectedPageInfo: pageInfoOutput{HasNextPage: true, StartCursor: "start", EndCursor: "cursor-3"},
			expectedTotal:    5,
		},
		{
			name: "order by creation time ascending with category and pagination",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"category":  "DIC_kwDOABC123",
				"orderBy":   "CREATED_AT",
				"direction": "ASC",
				"perPage":   float64(2),
				"after":     "cursor-1",
			},
			query: qListDiscussionsByCategoryOrdered,
			vars: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"categoryId":       "DIC_kwDOABC123",
				"orderByField":     "CREATED_AT",
				"orderByDirection": "ASC",
				"first":            float64(2),
				"after":            "cursor-1",
			},
			response:          discussionsResponse(discussionsGeneral, false, "cursor-2", 2),
			expectedCount:     2,
			expectedPageInfo:  pageInfoOutput{StartCursor: "start", EndCursor: "cursor-2"},
			expectedTotal:     2,
			expectCategoryCol: true,
		},
		{
			name: "orderBy without direction uses default order",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"orderBy": "UPDATED_AT",
			},
			query: qListDiscussions,
			vars: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
				"first": float64(30),
				"after": nil,
			},
			response:         discussionsResponse(discussionsAll, false, "cursor-3", 3),
			expectedCount:    3,
			expectedPageInfo: pageInfoOutput{StartCursor: "start", EndCursor: "cursor-3"},
			expectedTotal:    3,
		},
		{
			name: "direction without orderBy uses default order",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"direction": "DESC",
			},
			query: qListDiscussions,
			vars: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
				"first": float64(30),
				"after": nil,
			},
			response:         discussionsResponse(discussionsAll, false, "cursor-3", 3),
			expectedCount:    3,
			expectedPageInfo: pageInfoOutput{StartCursor: "start", EndCursor: "cursor-3"},
			expectedTotal:    3,
		},
		{
			name: "orderBy without direction keeps category filter",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
				"orderBy":  "CREATED_AT",
			},
			query: qListDiscussionsByCategory,
			vars: map[string]interface{}{
				"owner":      "owner",
				"repo":       "repo",
				"categoryId": "DIC_kwDOABC123",
				"first":      float64(30),
				"after":      nil,
			},
			response:          discussionsResponse(discussionsGeneral, false, "cursor-2", 2),
			expectedCount:     2,
			expectedPageInfo:  pageInfoOutput{StartCursor: "start", EndCursor: "cursor-2"},
			expectedTotal:     2,
			expectCategoryCol: true,
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
			name:      "repository not found error",
			reqParams: map[string]interface{}{"owner": "owner", "repo": "nonexistent-repo"},
			query:     qListDiscussions,
			vars: map[string]interface{}{
				"owner": "owner",
				"repo":  "nonexistent-repo",
				"first": float64(30),
				"after": nil,
			},
			response:    mockErrorRepoNotFound,
			expectError: true,
			errContains: "repository not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var httpClient *http.Client
			if tc.query != "" {
				httpClient = githubv4mock.NewMockedHTTPClient(githubv4mock.NewQueryMatcher(tc.query, tc.vars, tc.response))
			} else {
				httpClient = githubv4mock.NewMockedHTTPClient()
			}

			gqlClient := githubv4.NewClient(httpClient)
			_, handler := ListDiscussions(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			req := createMCPRequest(tc.reqParams)
			res, err := handler(context.Background(), req)
			require.NoError(t, err)
			text := getTextResult(t, res).Text

			if tc.expectError {
				require.True(t, res.IsError)
				assert.Contains(t, text, tc.errContains)
				return
			}
			require.False(t, res.IsError, text)

			var out struct {
				Discussions []*github.Discussion `json:"discussions"`
				PageInfo    pageInfoOutput       `json:"pageInfo"`
				TotalCount  int                  `json:"totalCount"`
			}
			require.NoError(t, json.Unmarshal([]byte(text), &out))

			assert.Len(t, out.Discussions, tc.expectedCount)
			assert.Equal(t, tc.expectedPageInfo, out.PageInfo)
			assert.Equal(t, tc.expectedTotal, out.TotalCount)

			// The mock data gives discussion N the author "userN" and the update time 2023-0N-15.
			for _, discussion := range out.Discussions {
				n := discussion.GetNumber()
				require.NotNil(t, discussion.User)
				assert.Equal(t, fmt.Sprintf("user%d", n), discussion.User.GetLogin())
				require.NotNil(t, discussion.UpdatedAt)
				assert.Equal(t, time.Date(2023, time.Month(n), 15, 0, 0, 0, 0, time.UTC), discussion.UpdatedAt.UTC())
			}

			if tc.expectCategoryCol {
				for _, discussion := range out.Discussions {
					require.NotNil(t, discussion.DiscussionCategory, "Discussion should have category")
					assert.Equal(t, "General", *discussion.DiscussionCategory.Name)
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

	mockResponse := githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussion": map[string]any{
				"comments": map[string]any{
					"nodes": []map[string]any{
						{"body": "This is the first comment"},
						{"body": "This is the second comment"},
					},
					"pageInfo":   mockPageInfo(true, "comment-cursor"),
					"totalCount": 7,
				},
			},
		},
	})

	tests := []struct {
		name      string
		reqParams map[string]interface{}
		vars      map[string]interface{}
	}{
		{
			name: "default pagination",
			reqParams: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"discussionNumber": int32(1),
			},
			vars: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"discussionNumber": float64(1),
				"first":            float64(30),
				"after":            nil,
			},
		},
		{
			name: "perPage and after",
			reqParams: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"discussionNumber": int32(1),
				"perPage":          float64(2),
				"after":            "prev-cursor",
			},
			vars: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"discussionNumber": float64(1),
				"first":            float64(2),
				"after":            "prev-cursor",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matcher := githubv4mock.NewQueryMatcher(qGetDiscussionComments, tc.vars, mockResponse)
			httpClient := githubv4mock.NewMockedHTTPClient(matcher)
			gqlClient := githubv4.NewClient(httpClient)
			_, handler := GetDiscussionComments(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.reqParams))
			require.NoError(t, err)

			textContent := getTextResult(t, result)
			require.False(t, result.IsError, textContent.Text)

			var out struct {
				Comments   []*github.IssueComment `json:"comments"`
				PageInfo   pageInfoOutput         `json:"pageInfo"`
				TotalCount int                    `json:"totalCount"`
			}
			require.NoError(t, json.Unmarshal([]byte(textContent.Text), &out))
			require.Len(t, out.Comments, 2)
			expectedBodies := []string{"This is the first comment", "This is the second comment"}
			for i, comment := range out.Comments {
				assert.Equal(t, expectedBodies[i], *comment.Body)
			}
			assert.Equal(t, pageInfoOutput{HasNextPage: true, StartCursor: "start", EndCursor: "comment-cursor"}, out.PageInfo)
			assert.Equal(t, 7, out.TotalCount)
		})
	}
}

func Test_ListDiscussionCategories(t *testing.T) {
	vars := map[string]interface{}{
		"owner": "owner",
		"repo":  "repo",
		"first": float64(25),
	}
	mockResp := githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussionCategories": map[string]any{
				"nodes": []map[string]any{
					{"id": "123", "name": "CategoryOne"},
					{"id": "456", "name": "CategoryTwo"},
				},
				"pageInfo":   mockPageInfo(false, "end"),
				"totalCount": 2,
			},
		},
	})
	matcher := githubv4mock.NewQueryMatcher(qListDiscussionCategories, vars, mockResp)
	httpClient := githubv4mock.NewMockedHTTPClient(matcher)
	gqlClient := githubv4.NewClient(httpClient)

	tool, handler := ListDiscussionCategories(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)
	assert.Equal(t, "list_discussion_categories", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "owner")
	assert.Contains(t, tool.InputSchema.Properties, "repo")
	for _, removed := range []string{"first", "last", "after", "before"} {
		assert.NotContains(t, tool.InputSchema.Properties, removed)
	}
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo"})

	request := createMCPRequest(map[string]interface{}{"owner": "owner", "repo": "repo"})
	result, err := handler(context.Background(), request)
	require.NoError(t, err)

	text := getTextResult(t, result).Text
	require.False(t, result.IsError, text)

	var out struct {
		Categories []map[string]string `json:"categories"`
		PageInfo   pageInfoOutput      `json:"pageInfo"`
		TotalCount int                 `json:"totalCount"`
	}
	require.NoError(t, json.Unmarshal([]byte(text), &out))
	require.Len(t, out.Categories, 2)
	assert.Equal(t, "123", out.Categories[0]["id"])
	assert.Equal(t, "CategoryOne", out.Categories[0]["name"])
	assert.Equal(t, "456", out.Categories[1]["id"])
	assert.Equal(t, "CategoryTwo", out.Categories[1]["name"])
	assert.Equal(t, pageInfoOutput{StartCursor: "start", EndCursor: "end"}, out.PageInfo)
	assert.Equal(t, 2, out.TotalCount)
}
