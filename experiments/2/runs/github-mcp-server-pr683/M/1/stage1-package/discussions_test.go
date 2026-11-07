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
	discussionsGeneral = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-10T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-03-10T00:00:00Z", "author": map[string]any{"login": "user3"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	discussionsAll = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-01-10T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 2, "title": "Discussion 2 title", "createdAt": "2023-02-01T00:00:00Z", "updatedAt": "2023-02-10T00:00:00Z", "author": map[string]any{"login": "user2"}, "url": "https://github.com/owner/repo/discussions/2", "category": map[string]any{"name": "Questions"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-03-10T00:00:00Z", "author": map[string]any{"login": "user3"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	pageInfoLastPage = map[string]any{
		"hasNextPage":     false,
		"hasPreviousPage": false,
		"startCursor":     "",
		"endCursor":       "",
	}
	mockResponseListAll = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes":      discussionsAll,
				"pageInfo":   pageInfoLastPage,
				"totalCount": 3,
			},
		},
	})
	mockResponseListGeneral = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes":      discussionsGeneral,
				"pageInfo":   pageInfoLastPage,
				"totalCount": 2,
			},
		},
	})
	mockResponseListNextPage = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{
				"nodes": discussionsAll[:1],
				"pageInfo": map[string]any{
					"hasNextPage":     true,
					"hasPreviousPage": true,
					"startCursor":     "Y3Vyc29yOjI=",
					"endCursor":       "Y3Vyc29yOjI=",
				},
				"totalCount": 3,
			},
		},
	})
	mockErrorRepoNotFound = githubv4mock.ErrorResponse("repository not found")
)

// GraphQL documents the discussion tools are expected to send.
const (
	queryListDiscussions                  = "query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	queryListDiscussionsByCategory        = "query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	queryListDiscussionsOrdered           = "query($after:String$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	queryListDiscussionsByCategoryOrdered = "query($after:String$categoryId:ID!$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	queryGetDiscussionComments            = "query($after:String$discussionNumber:Int!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussion(number: $discussionNumber){comments(first: $first, after: $after){nodes{body},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"
	queryListDiscussionCategories         = "query($first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussionCategories(first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
)

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
	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo"})

	varsListAll := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(DefaultGraphQLPageSize),
		"after": (*githubv4.String)(nil),
	}

	varsListWithPagination := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(1),
		"after": githubv4.String("Y3Vyc29yOjE="),
	}

	varsRepoNotFound := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("nonexistent-repo"),
		"first": githubv4.Int(DefaultGraphQLPageSize),
		"after": (*githubv4.String)(nil),
	}

	varsDiscussionsFiltered := map[string]interface{}{
		"owner":      githubv4.String("owner"),
		"repo":       githubv4.String("repo"),
		"categoryId": githubv4.ID("DIC_kwDOABC123"),
		"first":      githubv4.Int(DefaultGraphQLPageSize),
		"after":      (*githubv4.String)(nil),
	}

	varsOrderedByUpdatedDesc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"first":            githubv4.Int(DefaultGraphQLPageSize),
		"after":            (*githubv4.String)(nil),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
	}

	varsOrderedByCreatedAscWithPagination := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"first":            githubv4.Int(1),
		"after":            githubv4.String("Y3Vyc29yOjE="),
		"orderByField":     githubv4.DiscussionOrderField("CREATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
	}

	varsFilteredOrdered := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"categoryId":       githubv4.ID("DIC_kwDOABC123"),
		"first":            githubv4.Int(DefaultGraphQLPageSize),
		"after":            (*githubv4.String)(nil),
		"orderByField":     githubv4.DiscussionOrderField("CREATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
	}

	tests := []struct {
		name               string
		reqParams          map[string]interface{}
		matcher            githubv4mock.Matcher
		expectError        bool
		errContains        string
		expectedCount      int
		expectedPageInfo   pageInfoOutput
		expectedTotalCount int
	}{
		{
			name: "list all discussions without category filter",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
			},
			matcher:            githubv4mock.NewQueryMatcher(queryListDiscussions, varsListAll, mockResponseListAll),
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
			matcher:            githubv4mock.NewQueryMatcher(queryListDiscussionsByCategory, varsDiscussionsFiltered, mockResponseListGeneral),
			expectedCount:      2, // Only General discussions (matching the category ID)
			expectedTotalCount: 2,
		},
		{
			name: "order by UPDATED_AT DESC",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "UPDATED_AT",
				"direction": "DESC",
			},
			matcher:            githubv4mock.NewQueryMatcher(queryListDiscussionsOrdered, varsOrderedByUpdatedDesc, mockResponseListAll),
			expectedCount:      3,
			expectedTotalCount: 3,
		},
		{
			name: "order by CREATED_AT ASC with pagination",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "CREATED_AT",
				"direction": "ASC",
				"perPage":   float64(1),
				"after":     "Y3Vyc29yOjE=",
			},
			matcher:       githubv4mock.NewQueryMatcher(queryListDiscussionsOrdered, varsOrderedByCreatedAscWithPagination, mockResponseListNextPage),
			expectedCount: 1,
			expectedPageInfo: pageInfoOutput{
				HasNextPage:     true,
				HasPreviousPage: true,
				StartCursor:     "Y3Vyc29yOjI=",
				EndCursor:       "Y3Vyc29yOjI=",
			},
			expectedTotalCount: 3,
		},
		{
			name: "filter by category ID and order by CREATED_AT DESC",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"category":  "DIC_kwDOABC123",
				"orderBy":   "CREATED_AT",
				"direction": "DESC",
			},
			matcher:            githubv4mock.NewQueryMatcher(queryListDiscussionsByCategoryOrdered, varsFilteredOrdered, mockResponseListGeneral),
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
			matcher:            githubv4mock.NewQueryMatcher(queryListDiscussions, varsListAll, mockResponseListAll),
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
			matcher:            githubv4mock.NewQueryMatcher(queryListDiscussions, varsListAll, mockResponseListAll),
			expectedCount:      3,
			expectedTotalCount: 3,
		},
		{
			name: "orderBy without direction is ignored with category filter",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
				"orderBy":  "CREATED_AT",
			},
			matcher:            githubv4mock.NewQueryMatcher(queryListDiscussionsByCategory, varsDiscussionsFiltered, mockResponseListGeneral),
			expectedCount:      2,
			expectedTotalCount: 2,
		},
		{
			name: "perPage and after are passed as first and after",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(1),
				"after":   "Y3Vyc29yOjE=",
			},
			matcher:       githubv4mock.NewQueryMatcher(queryListDiscussions, varsListWithPagination, mockResponseListNextPage),
			expectedCount: 1,
			expectedPageInfo: pageInfoOutput{
				HasNextPage:     true,
				HasPreviousPage: true,
				StartCursor:     "Y3Vyc29yOjI=",
				EndCursor:       "Y3Vyc29yOjI=",
			},
			expectedTotalCount: 3,
		},
		{
			name: "perPage above maximum",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(101),
			},
			matcher:     githubv4mock.NewQueryMatcher(queryListDiscussions, varsListAll, mockResponseListAll),
			expectError: true,
			errContains: "perPage value 101 exceeds maximum of 100",
		},
		{
			name: "repository not found error",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "nonexistent-repo",
			},
			matcher:     githubv4mock.NewQueryMatcher(queryListDiscussions, varsRepoNotFound, mockErrorRepoNotFound),
			expectError: true,
			errContains: "repository not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			httpClient := githubv4mock.NewMockedHTTPClient(tc.matcher)
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

			var response struct {
				Discussions []*github.Discussion `json:"discussions"`
				PageInfo    pageInfoOutput       `json:"pageInfo"`
				TotalCount  int                  `json:"totalCount"`
			}
			require.NoError(t, json.Unmarshal([]byte(text), &response))

			assert.Len(t, response.Discussions, tc.expectedCount, "Expected %d discussions, got %d", tc.expectedCount, len(response.Discussions))
			assert.Equal(t, tc.expectedPageInfo, response.PageInfo)
			assert.Equal(t, tc.expectedTotalCount, response.TotalCount)

			// Verify that updatedAt and the author login are returned
			for _, discussion := range response.Discussions {
				require.NotNil(t, discussion.UpdatedAt, "Discussion should have updatedAt")
				assert.False(t, discussion.UpdatedAt.IsZero(), "Discussion should have a non-zero updatedAt")
				require.NotNil(t, discussion.User, "Discussion should have a user")
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
					"pageInfo": map[string]any{
						"hasNextPage":     true,
						"hasPreviousPage": false,
						"startCursor":     "Y3Vyc29yOjE=",
						"endCursor":       "Y3Vyc29yOjI=",
					},
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
				"first":            githubv4.Int(DefaultGraphQLPageSize),
				"after":            (*githubv4.String)(nil),
			},
		},
		{
			name: "perPage and after",
			reqParams: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"discussionNumber": int32(1),
				"perPage":          float64(2),
				"after":            "Y3Vyc29yOjA=",
			},
			vars: map[string]interface{}{
				"owner":            githubv4.String("owner"),
				"repo":             githubv4.String("repo"),
				"discussionNumber": githubv4.Int(1),
				"first":            githubv4.Int(2),
				"after":            githubv4.String("Y3Vyc29yOjA="),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matcher := githubv4mock.NewQueryMatcher(queryGetDiscussionComments, tc.vars, mockResponse)
			httpClient := githubv4mock.NewMockedHTTPClient(matcher)
			gqlClient := githubv4.NewClient(httpClient)
			_, handler := GetDiscussionComments(stubGetGQLClientFn(gqlClient), translations.NullTranslationHelper)

			result, err := handler(context.Background(), createMCPRequest(tc.reqParams))
			require.NoError(t, err)

			textContent := getTextResult(t, result)
			require.False(t, result.IsError, textContent.Text)

			var response struct {
				Comments   []*github.IssueComment `json:"comments"`
				PageInfo   pageInfoOutput         `json:"pageInfo"`
				TotalCount int                    `json:"totalCount"`
			}
			require.NoError(t, json.Unmarshal([]byte(textContent.Text), &response))
			assert.Len(t, response.Comments, 2)
			expectedBodies := []string{"This is the first comment", "This is the second comment"}
			for i, comment := range response.Comments {
				assert.Equal(t, expectedBodies[i], *comment.Body)
			}
			assert.Equal(t, pageInfoOutput{
				HasNextPage:     true,
				HasPreviousPage: false,
				StartCursor:     "Y3Vyc29yOjE=",
				EndCursor:       "Y3Vyc29yOjI=",
			}, response.PageInfo)
			assert.Equal(t, 5, response.TotalCount)
		})
	}
}

func Test_ListDiscussionCategories(t *testing.T) {
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
					"startCursor":     "Y3Vyc29yOjE=",
					"endCursor":       "Y3Vyc29yOjI=",
				},
				"totalCount": 2,
			},
		},
	})
	matcher := githubv4mock.NewQueryMatcher(queryListDiscussionCategories, vars, mockResp)
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
	require.False(t, result.IsError, text)

	var response struct {
		Categories []map[string]string `json:"categories"`
		PageInfo   pageInfoOutput      `json:"pageInfo"`
		TotalCount int                 `json:"totalCount"`
	}
	require.NoError(t, json.Unmarshal([]byte(text), &response))
	assert.Len(t, response.Categories, 2)
	assert.Equal(t, "123", response.Categories[0]["id"])
	assert.Equal(t, "CategoryOne", response.Categories[0]["name"])
	assert.Equal(t, "456", response.Categories[1]["id"])
	assert.Equal(t, "CategoryTwo", response.Categories[1]["name"])
	assert.Equal(t, pageInfoOutput{
		StartCursor: "Y3Vyc29yOjE=",
		EndCursor:   "Y3Vyc29yOjI=",
	}, response.PageInfo)
	assert.Equal(t, 2, response.TotalCount)
}
