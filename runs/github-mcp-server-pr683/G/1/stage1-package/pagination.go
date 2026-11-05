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
		paginationPageParam()(tool)
		paginationPerPageParam()(tool)
	}
}

// WithUnifiedPagination returns a ToolOption that adds "page", "perPage" and "after"
// parameters to the tool, for GraphQL tools that accept page/perPage and convert
// them to a cursor internally.
func WithUnifiedPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		paginationPageParam()(tool)
		paginationPerPageParam()(tool)
		paginationAfterParam()(tool)
	}
}

// WithCursorPagination returns a ToolOption that adds "perPage" and "after"
// parameters to the tool, for GraphQL tools that page by cursor only.
// https://docs.github.com/en/graphql/guides/using-pagination-in-the-graphql-api
func WithCursorPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		paginationPerPageParam()(tool)
		paginationAfterParam()(tool)
	}
}

// PaginationParams holds the REST pagination parameters of a request,
// and an optional cursor for tools backed by GraphQL.
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
// If After is set, it is used as the cursor.
func (p PaginationParams) ToGraphQLParams() (*GraphQLPaginationParams, error) {
	return CursorPaginationParams{PerPage: p.PerPage, After: p.After}.ToGraphQLParams()
}

// CursorPaginationParams holds the cursor pagination parameters of a request.
// It has no page number.
type CursorPaginationParams struct {
	PerPage int
	After   string
}

// OptionalCursorPaginationParams returns the "perPage" and "after" parameters from the request,
// or their default values if not present. "perPage" defaults to DefaultGraphQLPageSize.
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
// First is PerPage, and After is nil when the cursor is empty.
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
