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
