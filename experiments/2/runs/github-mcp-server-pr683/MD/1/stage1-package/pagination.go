package github

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// DefaultGraphQLPageSize is the default number of items requested per page
// from a GraphQL connection.
//
const DefaultGraphQLPageSize = 30

// CursorPaginationParams holds cursor-only pagination parameters, for tools
// that page through a GraphQL connection without page numbers.
type CursorPaginationParams struct {
	PerPage int
	After   string
}

// ToGraphQLParams converts cursor pagination parameters to GraphQL parameters.
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

// GraphQLPaginationParams holds the "first" and "after" arguments of a
// GraphQL connection. After is nil when there is no cursor.
// https://docs.github.com/en/graphql/reference/objects#connection
type GraphQLPaginationParams struct {
	First *int32
	After *string
}

// OptionalCursorPaginationParams returns the "perPage" and "after" parameters
// from the request. "perPage" defaults to DefaultGraphQLPageSize, and "after"
// defaults to an empty cursor.
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

// WithCursorPagination returns a ToolOption that adds "perPage" and "after"
// parameters to the tool, for GraphQL tools that page by cursor only.
func WithCursorPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		withPaginationPerPage()(tool)
		withPaginationAfter()(tool)
	}
}

// WithUnifiedPagination returns a ToolOption that adds "page", "perPage" and
// "after" parameters to the tool, for GraphQL tools that accept page/perPage
// and convert them to a cursor internally.
func WithUnifiedPagination() mcp.ToolOption {
	return func(tool *mcp.Tool) {
		mcp.WithNumber("page",
			mcp.Description("Page number for pagination (min 1)"),
			mcp.Min(1),
		)(tool)
		withPaginationPerPage()(tool)
		withPaginationAfter()(tool)
	}
}

func withPaginationPerPage() mcp.ToolOption {
	return mcp.WithNumber("perPage",
		mcp.Description("Results per page for pagination (min 1, max 100)"),
		mcp.Min(1),
		mcp.Max(100),
	)
}

func withPaginationAfter() mcp.ToolOption {
	return mcp.WithString("after",
		mcp.Description("Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs."),
	)
}
