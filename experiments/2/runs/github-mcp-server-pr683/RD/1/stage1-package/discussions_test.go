package github

import (
	"context"
	"encoding/json"
	"fmt"
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
	discussionsPageInfo = map[string]any{
		"hasNextPage":     true,
		"hasPreviousPage": false,
		"startCursor":     "start-cursor",
		"endCursor":       "end-cursor",
	}
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
				"nodes":      discussionsAll,
				"pageInfo":   discussionsPageInfo,
				"totalCount": 3,
			},
		},
	})
	mockResponseListGeneral = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes":      discussionsGeneral,
				"pageInfo":   discussionsPageInfo,
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
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo"})

	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")

	// The documents the tool must send, with and without a category filter and an ordering
	qDiscussions := "query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsFiltered := "query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsOrdered := "query($after:String$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	qDiscussionsFilteredOrdered := "query($after:String$categoryId:ID!$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"

	// Without pagination arguments, a page of 30 is requested and no cursor is sent
	varsListAll := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(30),
		"after": nil,
	}

	varsListWithPagination := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(5),
		"after": githubv4.String("previous-end-cursor"),
	}

	varsRepoNotFound := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("nonexistent-repo"),
		"first": githubv4.Int(30),
		"after": nil,
	}

	varsDiscussionsFiltered := map[string]interface{}{
		"owner":      githubv4.String("owner"),
		"repo":       githubv4.String("repo"),
		"categoryId": githubv4.ID("DIC_kwDOABC123"),
		"first":      githubv4.Int(30),
		"after":      nil,
	}

	varsOrderByCreatedAsc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"orderByField":     githubv4.DiscussionOrderField("CREATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
		"first":            githubv4.Int(30),
		"after":            nil,
	}

	varsOrderByUpdatedDesc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
		"first":            githubv4.Int(30),
		"after":            nil,
	}

	varsFilteredOrderedWithPagination := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"categoryId":       githubv4.ID("DIC_kwDOABC123"),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
		"first":            githubv4.Int(5),
		"after":            githubv4.String("previous-end-cursor"),
	}

	tests := []struct {
		name          string
		reqParams     map[string]interface{}
		query         string
		vars          map[string]interface{}
		response      githubv4mock.GQLResponse
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
			query:         qDiscussions,
			vars:          varsListAll,
			response:      mockResponseListAll,
			expectedCount: 3, // All discussions
		},
		{
			name: "list discussions with pagination",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(5),
				"after":   "previous-end-cursor",
			},
			query:         qDiscussions,
			vars:          varsListWithPagination,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "filter by category ID",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
			},
			query:         qDiscussionsFiltered,
			vars:          varsDiscussionsFiltered,
			response:      mockResponseListGeneral,
			expectedCount: 2, // Only General discussions (matching the category ID)
		},
		{
			name: "order by creation time ascending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "CREATED_AT",
				"direction": "ASC",
			},
			query:         qDiscussionsOrdered,
			vars:          varsOrderByCreatedAsc,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "order by update time descending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "UPDATED_AT",
				"direction": "DESC",
			},
			query:         qDiscussionsOrdered,
			vars:          varsOrderByUpdatedDesc,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "filter by category ID and order, with pagination",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"category":  "DIC_kwDOABC123",
				"orderBy":   "UPDATED_AT",
				"direction": "ASC",
				"perPage":   float64(5),
				"after":     "previous-end-cursor",
			},
			query:         qDiscussionsFilteredOrdered,
			vars:          varsFilteredOrderedWithPagination,
			response:      mockResponseListGeneral,
			expectedCount: 2,
		},
		{
			name: "orderBy without direction is ignored",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"orderBy": "UPDATED_AT",
			},
			query:         qDiscussions,
			vars:          varsListAll,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "direction without orderBy is ignored",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"direction": "DESC",
			},
			query:         qDiscussions,
			vars:          varsListAll,
			response:      mockResponseListAll,
			expectedCount: 3,
		},
		{
			name: "orderBy without direction keeps the category filter",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
				"orderBy":  "CREATED_AT",
			},
			query:         qDiscussionsFiltered,
			vars:          varsDiscussionsFiltered,
			response:      mockResponseListGeneral,
			expectedCount: 2,
		},
		{
			name: "repository not found error",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "nonexistent-repo",
			},
			query:       qDiscussions,
			vars:        varsRepoNotFound,
			response:    mockErrorRepoNotFound,
			expectError: true,
			errContains: "repository not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matcher := githubv4mock.NewQueryMatcher(tc.query, tc.vars, tc.response)
			gqlClient := githubv4.NewClient(githubv4mock.NewMockedHTTPClient(matcher))
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

			assert.True(t, response.PageInfo.HasNextPage)
			assert.False(t, response.PageInfo.HasPreviousPage)
			assert.Equal(t, "start-cursor", response.PageInfo.StartCursor)
			assert.Equal(t, "end-cursor", response.PageInfo.EndCursor)
			assert.Equal(t, tc.expectedCount, response.TotalCount)

			assert.Len(t, returnedDiscussions, tc.expectedCount, "Expected %d discussions, got %d", tc.expectedCount, len(returnedDiscussions))

			// Every discussion carries its update time and its author
			for _, discussion := range returnedDiscussions {
				require.NotNil(t, discussion.UpdatedAt, "Discussion should have updatedAt")
				assert.False(t, discussion.UpdatedAt.IsZero(), "Discussion should have updatedAt")
				require.NotNil(t, discussion.User, "Discussion should have an author")
				assert.Equal(t, fmt.Sprintf("user%d", discussion.GetNumber()), discussion.User.GetLogin())
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
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo", "discussionNumber"})

	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")

	q := "query($after:String$discussionNumber:Int!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussion(number: $discussionNumber){comments(first: $first, after: $after){nodes{body},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"
	// Without pagination arguments, a page of 30 is requested and no cursor is sent
	vars := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"discussionNumber": githubv4.Int(1),
		"first":            githubv4.Int(30),
		"after":            nil,
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
						"hasNextPage":     false,
						"hasPreviousPage": false,
						"startCursor":     "",
						"endCursor":       "",
					},
					"totalCount": 2,
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
			HasNextPage bool   `json:"hasNextPage"`
			EndCursor   string `json:"endCursor"`
		} `json:"pageInfo"`
		TotalCount int `json:"totalCount"`
	}
	err = json.Unmarshal([]byte(textContent.Text), &response)
	require.NoError(t, err)
	returnedComments := response.Comments
	assert.Len(t, returnedComments, 2)
	assert.False(t, response.PageInfo.HasNextPage)
	assert.Equal(t, 2, response.TotalCount)
	expectedBodies := []string{"This is the first comment", "This is the second comment"}
	for i, comment := range returnedComments {
		assert.Equal(t, expectedBodies[i], *comment.Body)
	}
}

func Test_ListDiscussionCategories(t *testing.T) {
	q := "query($first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussionCategories(first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	// The first 25 categories are always requested
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
