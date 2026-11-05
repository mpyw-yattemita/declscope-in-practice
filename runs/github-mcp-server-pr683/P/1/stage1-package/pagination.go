package github

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// WithPagination returns a ToolOption that adds "page" and "perPage" parameters to the tool.
// The "page" parameter is optional, min 1.
// The "perPage" parameter is optional, min 1, max 100. If unset, defaults to 30.
// https://docs.github.com/en/rest/using-the-rest-api/using-pagination-in-the-rest-api
func WithPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		mcp.WithNumber("page",
			mcp.Description("Page number for pagination (min 1)"),
			mcp.Min(1),
		)(tool)

		mcp.WithNumber("perPage",
			mcp.Description("Results per page for pagination (min 1, max 100)"),
			mcp.Min(1),
			mcp.Max(100),
		)(tool)
	}
}

// WithUnifiedPagination returns a ToolOption that adds "page", "perPage" and "after" parameters to the tool.
// It is for GraphQL tools that accept page/perPage and convert them to a cursor internally.
// The "after" parameter is optional and takes the endCursor of the previous page.
func WithUnifiedPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		mcp.WithNumber("page",
			mcp.Description("Page number for pagination (min 1)"),
			mcp.Min(1),
		)(tool)

		mcp.WithNumber("perPage",
			mcp.Description("Results per page for pagination (min 1, max 100)"),
			mcp.Min(1),
			mcp.Max(100),
		)(tool)

		mcp.WithString("after",
			mcp.Description("Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs."),
		)(tool)
	}
}

// WithCursorPagination returns a ToolOption that adds "perPage" and "after" parameters to the tool.
// It is for GraphQL tools that page with cursors only.
// https://docs.github.com/en/graphql/reference/objects#connection
func WithCursorPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		mcp.WithNumber("perPage",
			mcp.Description("Results per page for pagination (min 1, max 100)"),
			mcp.Min(1),
			mcp.Max(100),
		)(tool)

		mcp.WithString("after",
			mcp.Description("Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs."),
		)(tool)
	}
}

// PaginationParams holds REST pagination parameters, plus an optional GraphQL cursor.
type PaginationParams struct {
	Page    int
	PerPage int
	After   string
}

// ToGraphQLParams converts REST pagination parameters to GraphQL parameters.
// If After is set, it is used as the cursor. Page is not used.
func (p PaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	return CursorPaginationParams{
		PerPage: p.PerPage,
		After:   p.After,
	}.ToGraphQLParams()
}

// OptionalPaginationParams returns the "page", "perPage" and "after" parameters from the request,
// or their default values if not present, "page" default is 1, "perPage" default is 30,
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

// CursorPaginationParams holds cursor-only pagination parameters, with no page number.
type CursorPaginationParams struct {
	PerPage int
	After   string
}

// ToGraphQLParams converts cursor pagination parameters to GraphQL parameters.
// First is PerPage, and After is nil when the cursor is empty.
// It returns an error when PerPage is negative or above 100.
func (p CursorPaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	if p.PerPage > 100 {
		return nil, fmt.Errorf("perPage value %d exceeds maximum of 100", p.PerPage)
	}
	if p.PerPage < 0 {
		return nil, fmt.Errorf("perPage value %d cannot be negative", p.PerPage)
	}

	first := int32(p.PerPage)
	var after *string
	if p.After != "" {
		after = &p.After
	}
	return &GraphQLPaginationParams{
		First: &first,
		After: after,
	}, nil
}

// OptionalCursorPaginationParams returns the "perPage" and "after" parameters from the request,
// or their default values if not present, "perPage" default is DefaultGraphQLPageSize,
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

// GraphQLPaginationParams holds the "first" and "after" arguments of a GraphQL connection.
// https://graphql.org/learn/pagination/
type GraphQLPaginationParams struct {
	First *int32
	After *string
}
