package github

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/go-viper/mapstructure/v2"
	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shurcooL/githubv4"
)

func ListDiscussions(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_discussions",
			mcp.WithDescription(t("TOOL_LIST_DISCUSSIONS_DESCRIPTION", "List discussions for a repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_DISCUSSIONS_USER_TITLE", "List discussions"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("category",
				mcp.Description("Optional filter by discussion category ID. If provided, only discussions with this category are listed."),
			),
			mcp.WithString("orderBy",
				mcp.Description("Order discussions by field. If provided, the 'direction' also needs to be provided."),
				mcp.Enum("CREATED_AT", "UPDATED_AT"),
			),
			mcp.WithString("direction",
				mcp.Description("Order direction."),
				mcp.Enum("ASC", "DESC"),
			),
			WithCursorPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Required params
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Optional params
			category, err := OptionalParam[string](request, "category")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := OptionalParam[string](request, "orderBy")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			direction, err := OptionalParam[string](request, "direction")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			pagination, err := OptionalCursorPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			paginationParams, err := pagination.ToGraphQLParams()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			vars := map[string]interface{}{
				"owner": githubv4.String(owner),
				"repo":  githubv4.String(repo),
				"first": githubv4.Int(*paginationParams.First),
				"after": (*githubv4.String)(paginationParams.After),
			}
			if category != "" {
				vars["categoryId"] = githubv4.ID(category)
			}
			// Ordering is used only when both orderBy and direction are given.
			ordered := orderBy != "" && direction != ""
			if ordered {
				vars["orderByField"] = githubv4.DiscussionOrderField(orderBy)
				vars["orderByDirection"] = githubv4.OrderDirection(direction)
			}

			// githubv4 builds the query from the struct type, so each combination
			// of category filter and ordering needs its own query type.
			var discussions discussionsListConnection
			switch {
			case category != "" && ordered:
				var query struct {
					Repository struct {
						Discussions discussionsListConnection `graphql:"discussions(first: $first, after: $after, categoryId: $categoryId, orderBy: { field: $orderByField, direction: $orderByDirection })"`
					} `graphql:"repository(owner: $owner, name: $repo)"`
				}
				if err := client.Query(ctx, &query, vars); err != nil {
					return mcp.NewToolResultError(err.Error()), nil
				}
				discussions = query.Repository.Discussions
			case category != "":
				var query struct {
					Repository struct {
						Discussions discussionsListConnection `graphql:"discussions(first: $first, after: $after, categoryId: $categoryId)"`
					} `graphql:"repository(owner: $owner, name: $repo)"`
				}
				if err := client.Query(ctx, &query, vars); err != nil {
					return mcp.NewToolResultError(err.Error()), nil
				}
				discussions = query.Repository.Discussions
			case ordered:
				var query struct {
					Repository struct {
						Discussions discussionsListConnection `graphql:"discussions(first: $first, after: $after, orderBy: { field: $orderByField, direction: $orderByDirection })"`
					} `graphql:"repository(owner: $owner, name: $repo)"`
				}
				if err := client.Query(ctx, &query, vars); err != nil {
					return mcp.NewToolResultError(err.Error()), nil
				}
				discussions = query.Repository.Discussions
			default:
				var query struct {
					Repository struct {
						Discussions discussionsListConnection `graphql:"discussions(first: $first, after: $after)"`
					} `graphql:"repository(owner: $owner, name: $repo)"`
				}
				if err := client.Query(ctx, &query, vars); err != nil {
					return mcp.NewToolResultError(err.Error()), nil
				}
				discussions = query.Repository.Discussions
			}

			// Map nodes to GitHub Discussion objects
			discussionList := []*github.Discussion{}
			for _, n := range discussions.Nodes {
				discussionList = append(discussionList, &github.Discussion{
					Number:    github.Ptr(int(n.Number)),
					Title:     github.Ptr(string(n.Title)),
					HTMLURL:   github.Ptr(string(n.URL)),
					CreatedAt: &github.Timestamp{Time: n.CreatedAt.Time},
					UpdatedAt: &github.Timestamp{Time: n.UpdatedAt.Time},
					User: &github.User{
						Login: github.Ptr(string(n.Author.Login)),
					},
					DiscussionCategory: &github.DiscussionCategory{
						Name: github.Ptr(string(n.Category.Name)),
					},
				})
			}

			// Marshal and return
			out, err := json.Marshal(map[string]any{
				"discussions": discussionList,
				"pageInfo":    discussions.PageInfo.toMap(),
				"totalCount":  discussions.TotalCount,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussions: %w", err)
			}
			return mcp.NewToolResultText(string(out)), nil
		}
}

func GetDiscussion(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_discussion",
			mcp.WithDescription(t("TOOL_GET_DISCUSSION_DESCRIPTION", "Get a specific discussion by ID")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_DISCUSSION_USER_TITLE", "Get discussion"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithNumber("discussionNumber",
				mcp.Required(),
				mcp.Description("Discussion Number"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner            string
				Repo             string
				DiscussionNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

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
				"owner":            githubv4.String(params.Owner),
				"repo":             githubv4.String(params.Repo),
				"discussionNumber": githubv4.Int(params.DiscussionNumber),
			}
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			d := q.Repository.Discussion
			discussion := &github.Discussion{
				Number:    github.Ptr(int(d.Number)),
				Body:      github.Ptr(string(d.Body)),
				HTMLURL:   github.Ptr(string(d.URL)),
				CreatedAt: &github.Timestamp{Time: d.CreatedAt.Time},
				DiscussionCategory: &github.DiscussionCategory{
					Name: github.Ptr(string(d.Category.Name)),
				},
			}
			out, err := json.Marshal(discussion)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussion: %w", err)
			}

			return mcp.NewToolResultText(string(out)), nil
		}
}

func GetDiscussionComments(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_discussion_comments",
			mcp.WithDescription(t("TOOL_GET_DISCUSSION_COMMENTS_DESCRIPTION", "Get comments from a discussion")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_DISCUSSION_COMMENTS_USER_TITLE", "Get discussion comments"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner", mcp.Required(), mcp.Description("Repository owner")),
			mcp.WithString("repo", mcp.Required(), mcp.Description("Repository name")),
			mcp.WithNumber("discussionNumber", mcp.Required(), mcp.Description("Discussion Number")),
			WithCursorPagination(),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner            string
				Repo             string
				DiscussionNumber int32
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			pagination, err := OptionalCursorPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			paginationParams, err := pagination.ToGraphQLParams()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}

			var q struct {
				Repository struct {
					Discussion struct {
						Comments struct {
							Nodes []struct {
								Body githubv4.String
							}
							PageInfo   discussionsPageInfo
							TotalCount githubv4.Int
						} `graphql:"comments(first: $first, after: $after)"`
					} `graphql:"discussion(number: $discussionNumber)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			vars := map[string]interface{}{
				"owner":            githubv4.String(params.Owner),
				"repo":             githubv4.String(params.Repo),
				"discussionNumber": githubv4.Int(params.DiscussionNumber),
				"first":            githubv4.Int(*paginationParams.First),
				"after":            (*githubv4.String)(paginationParams.After),
			}
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			commentsConnection := q.Repository.Discussion.Comments
			comments := []*github.IssueComment{}
			for _, c := range commentsConnection.Nodes {
				comments = append(comments, &github.IssueComment{Body: github.Ptr(string(c.Body))})
			}

			out, err := json.Marshal(map[string]any{
				"comments":   comments,
				"pageInfo":   commentsConnection.PageInfo.toMap(),
				"totalCount": commentsConnection.TotalCount,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal comments: %w", err)
			}

			return mcp.NewToolResultText(string(out)), nil
		}
}

func ListDiscussionCategories(getGQLClient GetGQLClientFn, t translations.TranslationHelperFunc) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("list_discussion_categories",
			mcp.WithDescription(t("TOOL_LIST_DISCUSSION_CATEGORIES_DESCRIPTION", "List discussion categories with their id and name, for a repository")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_DISCUSSION_CATEGORIES_USER_TITLE", "List discussion categories"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// Decode params
			var params struct {
				Owner string
				Repo  string
			}
			if err := mapstructure.Decode(request.Params.Arguments, &params); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getGQLClient(ctx)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to get GitHub GQL client: %v", err)), nil
			}
			var q struct {
				Repository struct {
					DiscussionCategories struct {
						Nodes []struct {
							ID   githubv4.ID
							Name githubv4.String
						}
						PageInfo   discussionsPageInfo
						TotalCount githubv4.Int
					} `graphql:"discussionCategories(first: $first)"`
				} `graphql:"repository(owner: $owner, name: $repo)"`
			}
			vars := map[string]interface{}{
				"owner": githubv4.String(params.Owner),
				"repo":  githubv4.String(params.Repo),
				"first": githubv4.Int(discussionsCategoryPageSize),
			}
			if err := client.Query(ctx, &q, vars); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			categoriesConnection := q.Repository.DiscussionCategories
			categories := []map[string]string{}
			for _, c := range categoriesConnection.Nodes {
				categories = append(categories, map[string]string{
					"id":   fmt.Sprint(c.ID),
					"name": string(c.Name),
				})
			}
			out, err := json.Marshal(map[string]any{
				"categories": categories,
				"pageInfo":   categoriesConnection.PageInfo.toMap(),
				"totalCount": categoriesConnection.TotalCount,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to marshal discussion categories: %w", err)
			}
			return mcp.NewToolResultText(string(out)), nil
		}
}

// discussionsCategoryPageSize is the number of categories list_discussion_categories requests.
const discussionsCategoryPageSize = 25

// discussionsListConnection is the discussions connection that list_discussions queries.
type discussionsListConnection struct {
	Nodes []struct {
		Number    githubv4.Int
		Title     githubv4.String
		CreatedAt githubv4.DateTime
		UpdatedAt githubv4.DateTime
		Author    struct {
			Login githubv4.String
		}
		Category struct {
			Name githubv4.String
		} `graphql:"category"`
		URL githubv4.String `graphql:"url"`
	}
	PageInfo   discussionsPageInfo
	TotalCount githubv4.Int
}

// discussionsPageInfo is the pageInfo of a connection that the discussion tools query.
type discussionsPageInfo struct {
	HasNextPage     githubv4.Boolean
	HasPreviousPage githubv4.Boolean
	StartCursor     githubv4.String
	EndCursor       githubv4.String
}

// toMap returns the page info as the tools return it.
func (p discussionsPageInfo) toMap() map[string]any {
	return map[string]any{
		"hasNextPage":     p.HasNextPage,
		"hasPreviousPage": p.HasPreviousPage,
		"startCursor":     p.StartCursor,
		"endCursor":       p.EndCursor,
	}
}
