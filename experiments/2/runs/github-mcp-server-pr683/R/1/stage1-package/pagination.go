package github

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/shurcooL/githubv4"
)

const (
	// DefaultGraphQLPageSize is the page size used for GraphQL connections
	// when the caller does not set "perPage".
	DefaultGraphQLPageSize = 30

	// maxGraphQLPageSize is the largest "first" value GitHub's GraphQL API accepts.
	maxGraphQLPageSize = 100
)

func withPageParam() mcp.ToolOption {
	return mcp.WithNumber("page",
		mcp.Description("Page number for pagination (min 1)"),
		mcp.Min(1),
	)
}

func withPerPageParam() mcp.ToolOption {
	return mcp.WithNumber("perPage",
		mcp.Description("Results per page for pagination (min 1, max 100)"),
		mcp.Min(1),
		mcp.Max(100),
	)
}

func withAfterParam() mcp.ToolOption {
	return mcp.WithString("after",
		mcp.Description("Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs."),
	)
}

// WithPagination returns a ToolOption that adds "page" and "perPage" parameters to the tool.
// The "page" parameter is optional, min 1.
// The "perPage" parameter is optional, min 1, max 100. If unset, defaults to 30.
// https://docs.github.com/en/rest/using-the-rest-api/using-pagination-in-the-rest-api
func WithPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		withPageParam()(tool)
		withPerPageParam()(tool)
	}
}

// WithUnifiedPagination returns a ToolOption that adds "page", "perPage" and "after"
// parameters to the tool. It is meant for GraphQL tools that accept page/perPage and
// convert them to a cursor internally.
func WithUnifiedPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		withPageParam()(tool)
		withPerPageParam()(tool)
		withAfterParam()(tool)
	}
}

// WithCursorPagination returns a ToolOption that adds "perPage" and "after" parameters
// to the tool, for cursor-based (GraphQL connection) pagination.
// https://docs.github.com/en/graphql/guides/using-pagination-in-the-graphql-api
func WithCursorPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		withPerPageParam()(tool)
		withAfterParam()(tool)
	}
}

// PaginationParams holds the "page", "perPage" and "after" parameters of a request.
type PaginationParams struct {
	Page    int
	PerPage int
	After   string
}

// OptionalPaginationParams returns the "page", "perPage" and "after" parameters from the request,
// or their default values if not present, "page" default is 1, "perPage" default is 30.
// In future, we may want to make the default values configurable, or even have this
// function returned from `withPagination`, where the defaults are provided alongside
// the min/max values.
func OptionalPaginationParams(r mcp.CallToolRequest) (PaginationParams, error) {
	page, err := OptionalIntParamWithDefault(r, "page", 1)
	if err != nil {
		return PaginationParams{}, err
	}
	perPage, err := OptionalIntParamWithDefault(r, "perPage", 30)
	if err != nil {
		return PaginationParams{}, err
	}
	after, err := OptionalParam[string](r, "after")
	if err != nil {
		return PaginationParams{}, err
	}
	return PaginationParams{
		Page:    page,
		PerPage: perPage,
		After:   after,
	}, nil
}

// ToGraphQLParams converts REST pagination parameters to GraphQL parameters.
// If After is set, it is used as the cursor. Page is not translated into a cursor,
// because GraphQL connections cannot jump to an arbitrary page.
func (p PaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	return CursorPaginationParams{PerPage: p.PerPage, After: p.After}.ToGraphQLParams()
}

// CursorPaginationParams holds the "perPage" and "after" parameters of a request.
type CursorPaginationParams struct {
	PerPage int
	After   string
}

// OptionalCursorPaginationParams returns the "perPage" and "after" parameters from the request,
// or their default values if not present, "perPage" default is DefaultGraphQLPageSize.
func OptionalCursorPaginationParams(r mcp.CallToolRequest) (CursorPaginationParams, error) {
	perPage, err := OptionalIntParamWithDefault(r, "perPage", DefaultGraphQLPageSize)
	if err != nil {
		return CursorPaginationParams{}, err
	}
	after, err := OptionalParam[string](r, "after")
	if err != nil {
		return CursorPaginationParams{}, err
	}
	return CursorPaginationParams{
		PerPage: perPage,
		After:   after,
	}, nil
}

// ToGraphQLParams converts cursor pagination parameters to GraphQL parameters.
func (p CursorPaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	if p.PerPage > maxGraphQLPageSize {
		return nil, fmt.Errorf("perPage value %d exceeds maximum of %d", p.PerPage, maxGraphQLPageSize)
	}
	if p.PerPage < 0 {
		return nil, fmt.Errorf("perPage value %d cannot be negative", p.PerPage)
	}
	first := int32(p.PerPage) // #nosec G115 - bounded to [0, 100] above
	var after *string
	if p.After != "" {
		after = &p.After
	}
	return &GraphQLPaginationParams{
		First: &first,
		After: after,
	}, nil
}

// GraphQLPaginationParams holds the "first" and "after" arguments of a GraphQL connection.
type GraphQLPaginationParams struct {
	First *int32
	After *string
}

// addTo sets the "first" and "after" variables of a GraphQL query.
// "after" is always declared, and is sent as null when there is no cursor.
func (p *GraphQLPaginationParams) addTo(vars map[string]any) {
	if p.First != nil {
		vars["first"] = githubv4.Int(*p.First)
	}
	var after *githubv4.String
	if p.After != nil {
		after = githubv4.NewString(githubv4.String(*p.After))
	}
	vars["after"] = after
}

// pageInfoFragment is the PageInfo object of a GraphQL connection.
// https://docs.github.com/en/graphql/reference/objects#pageinfo
type pageInfoFragment struct {
	HasNextPage     githubv4.Boolean
	HasPreviousPage githubv4.Boolean
	StartCursor     githubv4.String
	EndCursor       githubv4.String
}

// pageInfoResult is the JSON form of a connection's PageInfo returned by GraphQL tools.
type pageInfoResult struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

func (p pageInfoFragment) toResult() pageInfoResult {
	return pageInfoResult{
		HasNextPage:     bool(p.HasNextPage),
		HasPreviousPage: bool(p.HasPreviousPage),
		StartCursor:     string(p.StartCursor),
		EndCursor:       string(p.EndCursor),
	}
}
