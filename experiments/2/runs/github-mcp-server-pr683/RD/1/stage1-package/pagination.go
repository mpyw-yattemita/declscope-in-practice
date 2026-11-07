package github

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/shurcooL/githubv4"
)

// DefaultGraphQLPageSize is the number of items a GraphQL tool requests when the caller sets no "perPage".
//
const DefaultGraphQLPageSize = 30

// PaginationParams holds the REST pagination parameters, and an optional GraphQL cursor.
type PaginationParams struct {
	Page    int
	PerPage int
	After   string
}

// ToGraphQLParams converts REST pagination parameters to GraphQL parameters.
// PerPage becomes First. If After is set, it is used as the cursor. Page has no
// GraphQL counterpart and is not converted.
func (p PaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	return CursorPaginationParams{
		PerPage: p.PerPage,
		After:   p.After,
	}.ToGraphQLParams()
}

// CursorPaginationParams holds cursor-only pagination parameters, with no page number.
type CursorPaginationParams struct {
	PerPage int
	After   string
}

// ToGraphQLParams converts cursor pagination parameters to GraphQL parameters.
// It returns an error when PerPage is negative or above 100.
func (p CursorPaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	if p.PerPage > 100 {
		return nil, fmt.Errorf("perPage value %d exceeds maximum of 100", p.PerPage)
	}
	if p.PerPage < 0 {
		return nil, fmt.Errorf("perPage value %d cannot be negative", p.PerPage)
	}
	first := int32(p.PerPage) // #nosec G115 -- bounded to [0, 100] above

	var after *string
	if p.After != "" {
		after = &p.After
	}

	return &GraphQLPaginationParams{
		First: &first,
		After: after,
	}, nil
}

// GraphQLPaginationParams holds the arguments of a GraphQL connection field.
// See https://graphql.org/learn/pagination/.
type GraphQLPaginationParams struct {
	First *int32
	After *string
}

// queryVariables returns the "first" and "after" variables of a query that pages a connection.
// "after" is sent as null when there is no cursor.
//
func (p *GraphQLPaginationParams) queryVariables() map[string]any {
	vars := map[string]any{
		"after": (*githubv4.String)(nil),
	}
	if p.First != nil {
		vars["first"] = githubv4.Int(*p.First)
	}
	if p.After != nil {
		vars["after"] = githubv4.NewString(githubv4.String(*p.After))
	}
	return vars
}

// paginationPageInfo is the pageInfo of a GraphQL connection.
// It is returned to the caller as is, so its JSON names follow the GraphQL ones.
//
type paginationPageInfo struct {
	HasNextPage     githubv4.Boolean `json:"hasNextPage"`
	HasPreviousPage githubv4.Boolean `json:"hasPreviousPage"`
	StartCursor     githubv4.String  `json:"startCursor"`
	EndCursor       githubv4.String  `json:"endCursor"`
}

// WithPagination returns a ToolOption that adds "page" and "perPage" parameters to the tool.
// The "page" parameter is optional, min 1.
// The "perPage" parameter is optional, min 1, max 100. If unset, defaults to 30.
// https://docs.github.com/en/rest/using-the-rest-api/using-pagination-in-the-rest-api
func WithPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		paginationPageParam()(tool)
		paginationPerPageParam()(tool)
	}
}

// WithCursorPagination returns a ToolOption that adds "perPage" and "after" parameters to the tool,
// for GraphQL tools that page with a cursor only.
// The "perPage" parameter is optional, min 1, max 100. If unset, defaults to 30.
// https://docs.github.com/en/graphql/reference/objects#connection
func WithCursorPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		paginationPerPageParam()(tool)
		paginationAfterParam()(tool)
	}
}

// WithUnifiedPagination returns a ToolOption that adds "page", "perPage" and "after" parameters to the tool,
// for GraphQL tools that accept page/perPage and convert them to a cursor internally.
func WithUnifiedPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		paginationPageParam()(tool)
		paginationPerPageParam()(tool)
		paginationAfterParam()(tool)
	}
}

// OptionalPaginationParams returns the "page", "perPage" and "after" parameters from the request,
// or their default values if not present, "page" default is 1, "perPage" default is 30, and
// "after" default is empty.
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

// OptionalCursorPaginationParams returns the "perPage" and "after" parameters from the request,
// or their default values if not present, "perPage" default is DefaultGraphQLPageSize, and
// "after" default is empty.
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

func paginationPageParam() mcp.ToolOption {
	return mcp.WithNumber("page",
		mcp.Description("Page number for pagination (min 1)"),
		mcp.Min(1),
	)
}

func paginationPerPageParam() mcp.ToolOption {
	return mcp.WithNumber("perPage",
		mcp.Description("Results per page for pagination (min 1, max 100)"),
		mcp.Min(1),
		mcp.Max(100),
	)
}

func paginationAfterParam() mcp.ToolOption {
	return mcp.WithString("after",
		mcp.Description("Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs."),
	)
}
