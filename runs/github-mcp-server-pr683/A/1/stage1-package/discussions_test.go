package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/github/github-mcp-server/internal/githubv4mock"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	mockPageInfo = map[string]any{
		"hasNextPage":     true,
		"hasPreviousPage": false,
		"startCursor":     "start-cursor",
		"endCursor":       "end-cursor",
	}
	mockResponseListAll = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{"nodes": discussionsAll, "pageInfo": mockPageInfo, "totalCount": 3},
		},
	})
	mockResponseListGeneral = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{"nodes": discussionsGeneral, "pageInfo": mockPageInfo, "totalCount": 2},
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

	assert.Contains(t, toolDef.InputSchema.Properties, "orderBy")
	assert.Contains(t, toolDef.InputSchema.Properties, "direction")

	qDiscussions := &listDiscussionsQuery{}
	qDiscussionsFiltered := &listDiscussionsByCategoryQuery{}
	qDiscussionsOrdered := &listDiscussionsOrderedQuery{}
	qDiscussionsFilteredOrdered := &listDiscussionsByCategoryOrderedQuery{}

	varsListAll := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(30),
		"after": (*githubv4.String)(nil),
	}

	varsListWithPagination := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(10),
		"after": githubv4.NewString("cursor123"),
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

	varsOrderByCreatedAsc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"orderByField":     githubv4.DiscussionOrderField("CREATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
		"first":            githubv4.Int(30),
		"after":            (*githubv4.String)(nil),
	}

	varsOrderByUpdatedDescWithPagination := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
		"first":            githubv4.Int(10),
		"after":            githubv4.NewString("cursor123"),
	}

	varsCategoryOrderByUpdatedDesc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"categoryId":       githubv4.ID("DIC_kwDOABC123"),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
		"first":            githubv4.Int(30),
		"after":            (*githubv4.String)(nil),
	}

	tests := []struct {
		name               string
		reqParams          map[string]interface{}
		matcher            githubv4mock.Matcher
		expectError        bool
		errContains        string
		expectedCount      int
		expectedTotalCount int
	}{
		{
			name: "list all discussions without category filter",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussions, varsListAll, mockResponseListAll),
			expectedCount:      3, // All discussions
			expectedTotalCount: 3,
		},
		{
			name: "filter by category ID",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussionsFiltered, varsDiscussionsFiltered, mockResponseListGeneral),
			expectedCount:      2, // Only General discussions (matching the category ID)
			expectedTotalCount: 2,
		},
		{
			name: "with perPage and after",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(10),
				"after":   "cursor123",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussions, varsListWithPagination, mockResponseListAll),
			expectedCount:      3,
			expectedTotalCount: 3,
		},
		{
			name: "order by created at ascending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "CREATED_AT",
				"direction": "ASC",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussionsOrdered, varsOrderByCreatedAsc, mockResponseListAll),
			expectedCount:      3,
			expectedTotalCount: 3,
		},
		{
			name: "order by updated at descending with pagination",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "UPDATED_AT",
				"direction": "DESC",
				"perPage":   float64(10),
				"after":     "cursor123",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussionsOrdered, varsOrderByUpdatedDescWithPagination, mockResponseListAll),
			expectedCount:      3,
			expectedTotalCount: 3,
		},
		{
			name: "filter by category and order by updated at descending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"category":  "DIC_kwDOABC123",
				"orderBy":   "UPDATED_AT",
				"direction": "DESC",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussionsFilteredOrdered, varsCategoryOrderByUpdatedDesc, mockResponseListGeneral),
			expectedCount:      2,
			expectedTotalCount: 2,
		},
		{
			name: "orderBy without direction is ignored",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"orderBy": "UPDATED_AT",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussions, varsListAll, mockResponseListAll),
			expectedCount:      3,
			expectedTotalCount: 3,
		},
		{
			name: "direction without orderBy is ignored",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"direction": "DESC",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussions, varsListAll, mockResponseListAll),
			expectedCount:      3,
			expectedTotalCount: 3,
		},
		{
			name: "category with orderBy only ignores ordering",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
				"orderBy":  "CREATED_AT",
			},
			matcher:            githubv4mock.NewQueryMatcher(qDiscussionsFiltered, varsDiscussionsFiltered, mockResponseListGeneral),
			expectedCount:      2,
			expectedTotalCount: 2,
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
			matcher:     githubv4mock.NewQueryMatcher(qDiscussions, varsRepoNotFound, mockErrorRepoNotFound),
			expectError: true,
			errContains: "repository not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var httpClient *http.Client
			if tc.matcher.Request != "" {
				httpClient = githubv4mock.NewMockedHTTPClient(tc.matcher)
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
			require.False(t, res.IsError, text)

			var response struct {
				Discussions []*github.Discussion `json:"discussions"`
				PageInfo    map[string]any       `json:"pageInfo"`
				TotalCount  int                  `json:"totalCount"`
			}
			err = json.Unmarshal([]byte(text), &response)
			require.NoError(t, err)

			assert.Len(t, response.Discussions, tc.expectedCount, "Expected %d discussions, got %d", tc.expectedCount, len(response.Discussions))
			assert.Equal(t, tc.expectedTotalCount, response.TotalCount)
			assert.Equal(t, mockPageInfo, response.PageInfo)

			// Every discussion carries its update time and author
			for _, discussion := range response.Discussions {
				require.NotNil(t, discussion.UpdatedAt, "Discussion should have updatedAt")
				assert.False(t, discussion.UpdatedAt.IsZero(), "Discussion should have a non-zero updatedAt")
				require.NotNil(t, discussion.User, "Discussion should have an author")
				assert.Equal(t, fmt.Sprintf("user%d", discussion.GetNumber()), discussion.GetUser().GetLogin())
			}

			// Verify that all returned discussions have a category if filtered
			if _, hasCategory := tc.reqParams["category"]; hasCategory {
				for _, discussion := range response.Discussions {
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

	var q struct {
		Repository struct {
			Discussion struct {
				Comments struct {
					Nodes []struct {
						Body githubv4.String
					}
					PageInfo   pageInfoFragment
					TotalCount githubv4.Int
				} `graphql:"comments(first: $first, after: $after)"`
			} `graphql:"discussion(number: $discussionNumber)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	mockResponse := githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussion": map[string]any{
				"comments": map[string]any{
					"nodes": []map[string]any{
						{"body": "This is the first comment"},
						{"body": "This is the second comment"},
					},
					"pageInfo":   mockPageInfo,
					"totalCount": 5,
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
				"owner":            githubv4.String("owner"),
				"repo":             githubv4.String("repo"),
				"discussionNumber": githubv4.Int(1),
				"first":            githubv4.Int(30),
				"after":            (*githubv4.String)(nil),
			},
		},
		{
			name: "with perPage and after",
			reqParams: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"discussionNumber": int32(1),
				"perPage":          float64(2),
				"after":            "cursor123",
			},
			vars: map[string]interface{}{
				"owner":            githubv4.String("owner"),
				"repo":             githubv4.String("repo"),
				"discussionNumber": githubv4.Int(1),
				"first":            githubv4.Int(2),
				"after":            githubv4.NewString("cursor123"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matcher := githubv4mock.NewQueryMatcher(q, tc.vars, mockResponse)
			httpClient := githubv4mock.NewMockedHTTPClient(matcher)
			gqlClient := githubv4.NewClient(httpClient)
			_, handler := GetDiscussionComments(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.reqParams))
			require.NoError(t, err)

			textContent := getTextResult(t, result)
			require.False(t, result.IsError, textContent.Text)

			var response struct {
				Comments   []*github.IssueComment `json:"comments"`
				PageInfo   map[string]any         `json:"pageInfo"`
				TotalCount int                    `json:"totalCount"`
			}
			require.NoError(t, json.Unmarshal([]byte(textContent.Text), &response))
			assert.Len(t, response.Comments, 2)
			expectedBodies := []string{"This is the first comment", "This is the second comment"}
			for i, comment := range response.Comments {
				assert.Equal(t, expectedBodies[i], *comment.Body)
			}
			assert.Equal(t, mockPageInfo, response.PageInfo)
			assert.Equal(t, 5, response.TotalCount)
		})
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
				PageInfo   pageInfoFragment
				TotalCount githubv4.Int
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
				"pageInfo":   mockPageInfo,
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
		PageInfo   map[string]any      `json:"pageInfo"`
		TotalCount int                 `json:"totalCount"`
	}
	require.NoError(t, json.Unmarshal([]byte(text), &response))
	categories := response.Categories
	assert.Len(t, categories, 2)
	assert.Equal(t, mockPageInfo, response.PageInfo)
	assert.Equal(t, 2, response.TotalCount)
	assert.Equal(t, "123", categories[0]["id"])
	assert.Equal(t, "CategoryOne", categories[0]["name"])
	assert.Equal(t, "456", categories[1]["id"])
	assert.Equal(t, "CategoryTwo", categories[1]["name"])
}

// queryRecorder is an http.RoundTripper that records GraphQL documents and
// answers every request with an empty data object.
type queryRecorder struct {
	queries []string
}

func (r *queryRecorder) RoundTrip(req *http.Request) (*http.Response, error) {
	var body struct {
		Query string `json:"query"`
	}
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		return nil, err
	}
	r.queries = append(r.queries, body.Query)
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		Request:    req,
	}, nil
}

func Test_DiscussionGraphQLDocuments(t *testing.T) {
	tests := []struct {
		name     string
		tool     func(GetGQLClientFn, translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc)
		args     map[string]interface{}
		expected string
	}{
		{
			name:     "list_discussions",
			tool:     ListDiscussions,
			args:     map[string]interface{}{"owner": "owner", "repo": "repo"},
			expected: "query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
		},
		{
			name:     "list_discussions with category",
			tool:     ListDiscussions,
			args:     map[string]interface{}{"owner": "owner", "repo": "repo", "category": "DIC_kwDOABC123"},
			expected: "query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
		},
		{
			name:     "list_discussions with ordering",
			tool:     ListDiscussions,
			args:     map[string]interface{}{"owner": "owner", "repo": "repo", "orderBy": "UPDATED_AT", "direction": "DESC"},
			expected: "query($after:String$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
		},
		{
			name:     "list_discussions with category and ordering",
			tool:     ListDiscussions,
			args:     map[string]interface{}{"owner": "owner", "repo": "repo", "category": "DIC_kwDOABC123", "orderBy": "CREATED_AT", "direction": "ASC"},
			expected: "query($after:String$categoryId:ID!$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
		},
		{
			name:     "list_discussions with orderBy only",
			tool:     ListDiscussions,
			args:     map[string]interface{}{"owner": "owner", "repo": "repo", "orderBy": "UPDATED_AT"},
			expected: "query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
		},
		{
			name:     "list_discussions with category and direction only",
			tool:     ListDiscussions,
			args:     map[string]interface{}{"owner": "owner", "repo": "repo", "category": "DIC_kwDOABC123", "direction": "ASC"},
			expected: "query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
		},
		{
			name:     "get_discussion_comments",
			tool:     GetDiscussionComments,
			args:     map[string]interface{}{"owner": "owner", "repo": "repo", "discussionNumber": int32(1)},
			expected: "query($after:String$discussionNumber:Int!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussion(number: $discussionNumber){comments(first: $first, after: $after){nodes{body},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
		},
		{
			name:     "list_discussion_categories",
			tool:     ListDiscussionCategories,
			args:     map[string]interface{}{"owner": "owner", "repo": "repo"},
			expected: "query($first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussionCategories(first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := &queryRecorder{}
			gqlClient := githubv4.NewClient(&http.Client{Transport: recorder})
			_, handler := tc.tool(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			_, err := handler(context.Background(), createMCPRequest(tc.args))
			require.NoError(t, err)
			require.Equal(t, []string{tc.expected}, recorder.queries)
		})
	}
}

func Test_ListDiscussions_DeletedAuthor(t *testing.T) {
	// The author is null when the account that started the discussion was deleted.
	response := githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes": []map[string]any{
					{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-15T00:00:00Z", "author": nil, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
				},
				"pageInfo":   mockPageInfo,
				"totalCount": 1,
			},
		},
	})
	vars := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(30),
		"after": (*githubv4.String)(nil),
	}
	httpClient := githubv4mock.NewMockedHTTPClient(githubv4mock.NewQueryMatcher(&listDiscussionsQuery{}, vars, response))
	_, handler := ListDiscussions(stubGetGQLClientFn(githubv4.NewClient(httpClient)), translations.NullTranslationHelper)

	res, err := handler(context.Background(), createMCPRequest(map[string]interface{}{"owner": "owner", "repo": "repo"}))
	require.NoError(t, err)
	text := getTextResult(t, res).Text
	require.False(t, res.IsError, text)

	var out struct {
		Discussions []*github.Discussion `json:"discussions"`
	}
	require.NoError(t, json.Unmarshal([]byte(text), &out))
	require.Len(t, out.Discussions, 1)
	assert.Nil(t, out.Discussions[0].User)
	assert.Equal(t, time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC), out.Discussions[0].GetUpdatedAt().Time)
}
