package github

import (
	"context"
	"encoding/json"
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

// The GraphQL documents the discussion tools must send.
const (
	discussionsListQuery                  = "query($after:String$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	discussionsListByCategoryQuery        = "query($after:String$categoryId:ID!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	discussionsListOrderedQuery           = "query($after:String$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	discussionsListByCategoryOrderedQuery = "query($after:String$categoryId:ID!$first:Int!$orderByDirection:OrderDirection!$orderByField:DiscussionOrderField!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection }){nodes{number,title,createdAt,updatedAt,author{login},category{name},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
	discussionsCommentsQuery              = "query($after:String$discussionNumber:Int!$first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussion(number: $discussionNumber){comments(first: $first, after: $after){nodes{body},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"
	discussionsCategoriesQuery            = "query($first:Int!$owner:String!$repo:String!){repository(owner: $owner, name: $repo){discussionCategories(first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"
)

var (
	discussionsGeneral = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-04-01T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-05-01T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	discussionsAll = []map[string]any{
		{"number": 1, "title": "Discussion 1 title", "createdAt": "2023-01-01T00:00:00Z", "updatedAt": "2023-04-01T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/1", "category": map[string]any{"name": "General"}},
		{"number": 2, "title": "Discussion 2 title", "createdAt": "2023-02-01T00:00:00Z", "updatedAt": "2023-06-01T00:00:00Z", "author": map[string]any{"login": "user2"}, "url": "https://github.com/owner/repo/discussions/2", "category": map[string]any{"name": "Questions"}},
		{"number": 3, "title": "Discussion 3 title", "createdAt": "2023-03-01T00:00:00Z", "updatedAt": "2023-05-01T00:00:00Z", "author": map[string]any{"login": "user1"}, "url": "https://github.com/owner/repo/discussions/3", "category": map[string]any{"name": "General"}},
	}
	discussionsPageInfoLast = map[string]any{
		"hasNextPage":     false,
		"hasPreviousPage": false,
		"startCursor":     "Y3Vyc29yOjE=",
		"endCursor":       "Y3Vyc29yOjM=",
	}
	discussionsPageInfoMore = map[string]any{
		"hasNextPage":     true,
		"hasPreviousPage": true,
		"startCursor":     "Y3Vyc29yOjI=",
		"endCursor":       "Y3Vyc29yOjM=",
	}
	mockResponseListAll = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{"nodes": discussionsAll, "pageInfo": discussionsPageInfoLast, "totalCount": 3},
		},
	})
	mockResponseListGeneral = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{"nodes": discussionsGeneral, "pageInfo": discussionsPageInfoLast, "totalCount": 2},
		},
	})
	mockResponseListDiscussionsPage = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{"nodes": discussionsAll[1:2], "pageInfo": discussionsPageInfoMore, "totalCount": 3},
		},
	})
	// discussionsByUpdatedDesc is discussionsAll ordered by updatedAt, newest first.
	// Discussion 2 has a deleted author, which the API returns as null.
	discussionsByUpdatedDesc = []map[string]any{
		{"number": 2, "title": "Discussion 2 title", "createdAt": "2023-02-01T00:00:00Z", "updatedAt": "2023-06-01T00:00:00Z", "author": nil, "url": "https://github.com/owner/repo/discussions/2", "category": map[string]any{"name": "Questions"}},
		discussionsAll[2],
		discussionsAll[0],
	}
	mockResponseListByUpdatedDesc = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{"nodes": discussionsByUpdatedDesc, "pageInfo": discussionsPageInfoLast, "totalCount": 3},
		},
	})
	mockResponseListGeneralByCreatedAsc = githubv4mock.DataResponse(map[string]any{
		"repository": map[string]any{
			"discussions": map[string]any{"nodes": discussionsGeneral, "pageInfo": discussionsPageInfoMore, "totalCount": 2},
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
	assert.Contains(t, toolDef.InputSchema.Properties, "category")
	assert.Contains(t, toolDef.InputSchema.Properties, "orderBy")
	assert.Contains(t, toolDef.InputSchema.Properties, "direction")
	assert.Contains(t, toolDef.InputSchema.Properties, "perPage")
	assert.Contains(t, toolDef.InputSchema.Properties, "after")
	assert.NotContains(t, toolDef.InputSchema.Properties, "page")
	assert.ElementsMatch(t, toolDef.InputSchema.Required, []string{"owner", "repo"})
	assert.Equal(t, []string{"CREATED_AT", "UPDATED_AT"}, toolDef.InputSchema.Properties["orderBy"].(map[string]any)["enum"])
	assert.Equal(t, []string{"ASC", "DESC"}, toolDef.InputSchema.Properties["direction"].(map[string]any)["enum"])

	varsListAll := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(30),
		"after": (*githubv4.String)(nil),
	}

	varsListPage := map[string]interface{}{
		"owner": githubv4.String("owner"),
		"repo":  githubv4.String("repo"),
		"first": githubv4.Int(1),
		"after": githubv4.String("Y3Vyc29yOjE="),
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

	varsListByUpdatedDesc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"first":            githubv4.Int(30),
		"after":            (*githubv4.String)(nil),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
	}

	varsListGeneralByCreatedAsc := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"categoryId":       githubv4.ID("DIC_kwDOABC123"),
		"first":            githubv4.Int(2),
		"after":            githubv4.String("Y3Vyc29yOjE="),
		"orderByField":     githubv4.DiscussionOrderField("CREATED_AT"),
		"orderByDirection": githubv4.OrderDirection("ASC"),
	}

	tests := []struct {
		name             string
		reqParams        map[string]interface{}
		matcher          *githubv4mock.Matcher
		expectError      bool
		errContains      string
		expectedNumbers  []int
		expectedPageInfo map[string]any
		expectedTotal    int
	}{
		{
			name: "order by updatedAt descending",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"orderBy":   "UPDATED_AT",
				"direction": "DESC",
			},
			matcher:          githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListOrderedQuery, varsListByUpdatedDesc, mockResponseListByUpdatedDesc)),
			expectedNumbers:  []int{2, 3, 1},
			expectedPageInfo: discussionsPageInfoLast,
			expectedTotal:    3,
		},
		{
			name: "order by createdAt ascending with category and pagination",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"category":  "DIC_kwDOABC123",
				"orderBy":   "CREATED_AT",
				"direction": "ASC",
				"perPage":   float64(2),
				"after":     "Y3Vyc29yOjE=",
			},
			matcher:          githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListByCategoryOrderedQuery, varsListGeneralByCreatedAsc, mockResponseListGeneralByCreatedAsc)),
			expectedNumbers:  []int{1, 3},
			expectedPageInfo: discussionsPageInfoMore,
			expectedTotal:    2,
		},
		{
			name: "orderBy without direction is ignored",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"orderBy": "UPDATED_AT",
			},
			matcher:          githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListQuery, varsListAll, mockResponseListAll)),
			expectedNumbers:  []int{1, 2, 3},
			expectedPageInfo: discussionsPageInfoLast,
			expectedTotal:    3,
		},
		{
			name: "direction without orderBy is ignored",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"direction": "DESC",
			},
			matcher:          githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListQuery, varsListAll, mockResponseListAll)),
			expectedNumbers:  []int{1, 2, 3},
			expectedPageInfo: discussionsPageInfoLast,
			expectedTotal:    3,
		},
		{
			name: "direction without orderBy is ignored with category",
			reqParams: map[string]interface{}{
				"owner":     "owner",
				"repo":      "repo",
				"category":  "DIC_kwDOABC123",
				"direction": "ASC",
			},
			matcher:          githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListByCategoryQuery, varsDiscussionsFiltered, mockResponseListGeneral)),
			expectedNumbers:  []int{1, 3},
			expectedPageInfo: discussionsPageInfoLast,
			expectedTotal:    2,
		},
		{
			name: "list all discussions without category filter",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
			},
			matcher:          githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListQuery, varsListAll, mockResponseListAll)),
			expectedNumbers:  []int{1, 2, 3},
			expectedPageInfo: discussionsPageInfoLast,
			expectedTotal:    3,
		},
		{
			name: "filter by category ID",
			reqParams: map[string]interface{}{
				"owner":    "owner",
				"repo":     "repo",
				"category": "DIC_kwDOABC123",
			},
			matcher:          githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListByCategoryQuery, varsDiscussionsFiltered, mockResponseListGeneral)),
			expectedNumbers:  []int{1, 3},
			expectedPageInfo: discussionsPageInfoLast,
			expectedTotal:    2,
		},
		{
			name: "perPage and after are sent as first and after",
			reqParams: map[string]interface{}{
				"owner":   "owner",
				"repo":    "repo",
				"perPage": float64(1),
				"after":   "Y3Vyc29yOjE=",
			},
			matcher:          githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListQuery, varsListPage, mockResponseListDiscussionsPage)),
			expectedNumbers:  []int{2},
			expectedPageInfo: discussionsPageInfoMore,
			expectedTotal:    3,
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
			name: "after of wrong type",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "repo",
				"after": float64(1),
			},
			expectError: true,
			errContains: "parameter after is not of type string",
		},
		{
			name: "repository not found error",
			reqParams: map[string]interface{}{
				"owner": "owner",
				"repo":  "nonexistent-repo",
			},
			matcher:     githubv4mock.Ptr(githubv4mock.NewQueryMatcher(discussionsListQuery, varsRepoNotFound, mockErrorRepoNotFound)),
			expectError: true,
			errContains: "repository not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var httpClient *http.Client
			if tc.matcher != nil {
				httpClient = githubv4mock.NewMockedHTTPClient(*tc.matcher)
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

			var response struct {
				Discussions []*github.Discussion `json:"discussions"`
				PageInfo    map[string]any       `json:"pageInfo"`
				TotalCount  int                  `json:"totalCount"`
			}
			require.NoError(t, json.Unmarshal([]byte(text), &response))

			numbers := make([]int, 0, len(response.Discussions))
			for _, discussion := range response.Discussions {
				numbers = append(numbers, discussion.GetNumber())
				require.NotNil(t, discussion.DiscussionCategory, "Discussion should have category")
				assert.NotEmpty(t, *discussion.DiscussionCategory.Name, "Discussion should have category name")
				assert.False(t, discussion.GetUpdatedAt().IsZero(), "Discussion should have updatedAt")
			}
			assert.Equal(t, tc.expectedNumbers, numbers)
			assert.Equal(t, tc.expectedPageInfo, response.PageInfo)
			assert.Equal(t, tc.expectedTotal, response.TotalCount)
		})
	}
}

func Test_ListDiscussions_UpdatedAtAndAuthor(t *testing.T) {
	vars := map[string]interface{}{
		"owner":            githubv4.String("owner"),
		"repo":             githubv4.String("repo"),
		"first":            githubv4.Int(30),
		"after":            (*githubv4.String)(nil),
		"orderByField":     githubv4.DiscussionOrderField("UPDATED_AT"),
		"orderByDirection": githubv4.OrderDirection("DESC"),
	}
	httpClient := githubv4mock.NewMockedHTTPClient(githubv4mock.NewQueryMatcher(discussionsListOrderedQuery, vars, mockResponseListByUpdatedDesc))
	_, handler := ListDiscussions(stubGetGQLClientFn(githubv4.NewClient(httpClient)), translations.NullTranslationHelper)

	res, err := handler(context.Background(), createMCPRequest(map[string]interface{}{
		"owner":     "owner",
		"repo":      "repo",
		"orderBy":   "UPDATED_AT",
		"direction": "DESC",
	}))
	require.NoError(t, err)
	text := getTextResult(t, res).Text
	require.False(t, res.IsError, text)

	var response struct {
		Discussions []struct {
			Number    int    `json:"number"`
			UpdatedAt string `json:"updated_at"`
			User      *struct {
				Login string `json:"login"`
			} `json:"user"`
		} `json:"discussions"`
	}
	require.NoError(t, json.Unmarshal([]byte(text), &response))
	require.Len(t, response.Discussions, 3)

	// Discussion 2's author was deleted, so it has no user.
	assert.Equal(t, 2, response.Discussions[0].Number)
	assert.Equal(t, "2023-06-01T00:00:00Z", response.Discussions[0].UpdatedAt)
	assert.Nil(t, response.Discussions[0].User)

	assert.Equal(t, 3, response.Discussions[1].Number)
	assert.Equal(t, "2023-05-01T00:00:00Z", response.Discussions[1].UpdatedAt)
	require.NotNil(t, response.Discussions[1].User)
	assert.Equal(t, "user1", response.Discussions[1].User.Login)

	assert.Equal(t, 1, response.Discussions[2].Number)
	assert.Equal(t, "2023-04-01T00:00:00Z", response.Discussions[2].UpdatedAt)
	require.NotNil(t, response.Discussions[2].User)
	assert.Equal(t, "user1", response.Discussions[2].User.Login)
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
					"pageInfo":   discussionsPageInfoMore,
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
			name: "default page",
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
			name: "perPage and after",
			reqParams: map[string]interface{}{
				"owner":            "owner",
				"repo":             "repo",
				"discussionNumber": int32(1),
				"perPage":          float64(2),
				"after":            "Y3Vyc29yOjE=",
			},
			vars: map[string]interface{}{
				"owner":            githubv4.String("owner"),
				"repo":             githubv4.String("repo"),
				"discussionNumber": githubv4.Int(1),
				"first":            githubv4.Int(2),
				"after":            githubv4.String("Y3Vyc29yOjE="),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matcher := githubv4mock.NewQueryMatcher(discussionsCommentsQuery, tc.vars, mockResponse)
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
			assert.Equal(t, discussionsPageInfoMore, response.PageInfo)
			assert.Equal(t, 5, response.TotalCount)
		})
	}
}

func Test_GetDiscussionCommentsInvalidPerPage(t *testing.T) {
	_, handler := GetDiscussionComments(stubGetGQLClientFn(githubv4.NewClient(nil)), translations.NullTranslationHelper)

	result, err := handler(context.Background(), createMCPRequest(map[string]interface{}{
		"owner":            "owner",
		"repo":             "repo",
		"discussionNumber": int32(1),
		"perPage":          float64(-1),
	}))
	require.NoError(t, err)
	require.True(t, result.IsError)
	assert.Contains(t, getTextResult(t, result).Text, "perPage value -1 cannot be negative")
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
				"pageInfo":   discussionsPageInfoLast,
				"totalCount": 2,
			},
		},
	})
	matcher := githubv4mock.NewQueryMatcher(discussionsCategoriesQuery, vars, mockResp)
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
	var response struct {
		Categories []map[string]string `json:"categories"`
		PageInfo   map[string]any      `json:"pageInfo"`
		TotalCount int                 `json:"totalCount"`
	}
	require.NoError(t, json.Unmarshal([]byte(text), &response))
	assert.Len(t, response.Categories, 2)
	assert.Equal(t, "123", response.Categories[0]["id"])
	assert.Equal(t, "CategoryOne", response.Categories[0]["name"])
	assert.Equal(t, "456", response.Categories[1]["id"])
	assert.Equal(t, "CategoryTwo", response.Categories[1]["name"])
	assert.Equal(t, discussionsPageInfoLast, response.PageInfo)
	assert.Equal(t, 2, response.TotalCount)
}
