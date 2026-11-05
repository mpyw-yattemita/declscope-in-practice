package github

import (
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionalPaginationParams(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]any
		expected    PaginationParams
		expectError bool
	}{
		{
			name:     "no pagination parameters, default values",
			params:   map[string]any{},
			expected: PaginationParams{Page: 1, PerPage: 30},
		},
		{
			name:     "page parameter, default perPage",
			params:   map[string]any{"page": float64(2)},
			expected: PaginationParams{Page: 2, PerPage: 30},
		},
		{
			name:     "perPage parameter, default page",
			params:   map[string]any{"perPage": float64(50)},
			expected: PaginationParams{Page: 1, PerPage: 50},
		},
		{
			name:     "page and perPage parameters",
			params:   map[string]any{"page": float64(2), "perPage": float64(50)},
			expected: PaginationParams{Page: 2, PerPage: 50},
		},
		{
			name:     "after parameter",
			params:   map[string]any{"after": "cursor"},
			expected: PaginationParams{Page: 1, PerPage: 30, After: "cursor"},
		},
		{
			name:        "invalid page parameter",
			params:      map[string]any{"page": "not-a-number"},
			expectError: true,
		},
		{
			name:        "invalid perPage parameter",
			params:      map[string]any{"perPage": "not-a-number"},
			expectError: true,
		},
		{
			name:        "invalid after parameter",
			params:      map[string]any{"after": float64(1)},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := OptionalPaginationParams(createMCPRequest(tc.params))

			if tc.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestOptionalCursorPaginationParams(t *testing.T) {
	tests := []struct {
		name        string
		params      map[string]any
		expected    CursorPaginationParams
		expectError bool
	}{
		{
			name:     "no pagination parameters, default values",
			params:   map[string]any{},
			expected: CursorPaginationParams{PerPage: DefaultGraphQLPageSize},
		},
		{
			name:     "perPage and after parameters",
			params:   map[string]any{"perPage": float64(10), "after": "cursor"},
			expected: CursorPaginationParams{PerPage: 10, After: "cursor"},
		},
		{
			name:     "page parameter is ignored",
			params:   map[string]any{"page": float64(3)},
			expected: CursorPaginationParams{PerPage: DefaultGraphQLPageSize},
		},
		{
			name:        "invalid perPage parameter",
			params:      map[string]any{"perPage": "not-a-number"},
			expectError: true,
		},
		{
			name:        "invalid after parameter",
			params:      map[string]any{"after": float64(1)},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := OptionalCursorPaginationParams(createMCPRequest(tc.params))

			if tc.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestToGraphQLParams(t *testing.T) {
	int32Ptr := func(v int32) *int32 { return &v }
	stringPtr := func(v string) *string { return &v }

	tests := []struct {
		name     string
		cursor   CursorPaginationParams
		expected *GraphQLPaginationParams
		errMsg   string
	}{
		{
			name:     "no cursor",
			cursor:   CursorPaginationParams{PerPage: 30},
			expected: &GraphQLPaginationParams{First: int32Ptr(30)},
		},
		{
			name:     "with cursor",
			cursor:   CursorPaginationParams{PerPage: 10, After: "cursor"},
			expected: &GraphQLPaginationParams{First: int32Ptr(10), After: stringPtr("cursor")},
		},
		{
			name:     "maximum perPage",
			cursor:   CursorPaginationParams{PerPage: 100},
			expected: &GraphQLPaginationParams{First: int32Ptr(100)},
		},
		{
			name:   "perPage above maximum",
			cursor: CursorPaginationParams{PerPage: 101},
			errMsg: "perPage value 101 exceeds maximum of 100",
		},
		{
			name:   "negative perPage",
			cursor: CursorPaginationParams{PerPage: -1},
			errMsg: "perPage value -1 cannot be negative",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			check := func(t *testing.T, result *GraphQLPaginationParams, err error) {
				if tc.errMsg != "" {
					assert.EqualError(t, err, tc.errMsg)
					assert.Nil(t, result)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}

			t.Run("CursorPaginationParams", func(t *testing.T) {
				result, err := tc.cursor.ToGraphQLParams()
				check(t, result, err)
			})
			t.Run("PaginationParams", func(t *testing.T) {
				result, err := PaginationParams{Page: 2, PerPage: tc.cursor.PerPage, After: tc.cursor.After}.ToGraphQLParams()
				check(t, result, err)
			})
		})
	}
}

func TestPaginationToolOptions(t *testing.T) {
	perPage := map[string]any{
		"type":        "number",
		"description": "Results per page for pagination (min 1, max 100)",
		"minimum":     float64(1),
		"maximum":     float64(100),
	}
	page := map[string]any{
		"type":        "number",
		"description": "Page number for pagination (min 1)",
		"minimum":     float64(1),
	}
	after := map[string]any{
		"type":        "string",
		"description": "Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs.",
	}

	tests := []struct {
		name     string
		option   mcp.ToolOption
		expected map[string]any
	}{
		{
			name:     "WithPagination",
			option:   WithPagination(),
			expected: map[string]any{"page": page, "perPage": perPage},
		},
		{
			name:     "WithCursorPagination",
			option:   WithCursorPagination(),
			expected: map[string]any{"perPage": perPage, "after": after},
		},
		{
			name:     "WithUnifiedPagination",
			option:   WithUnifiedPagination(),
			expected: map[string]any{"page": page, "perPage": perPage, "after": after},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tool := mcp.NewTool("test", tc.option)
			assert.Equal(t, tc.expected, tool.InputSchema.Properties)
			assert.Empty(t, tool.InputSchema.Required)
		})
	}
}
