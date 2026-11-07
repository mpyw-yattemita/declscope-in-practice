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
			name:   "no pagination parameters, default values",
			params: map[string]any{},
			expected: PaginationParams{
				Page:    1,
				PerPage: 30,
			},
			expectError: false,
		},
		{
			name: "page parameter, default perPage",
			params: map[string]any{
				"page": float64(2),
			},
			expected: PaginationParams{
				Page:    2,
				PerPage: 30,
			},
			expectError: false,
		},
		{
			name: "perPage parameter, default page",
			params: map[string]any{
				"perPage": float64(50),
			},
			expected: PaginationParams{
				Page:    1,
				PerPage: 50,
			},
			expectError: false,
		},
		{
			name: "page and perPage parameters",
			params: map[string]any{
				"page":    float64(2),
				"perPage": float64(50),
			},
			expected: PaginationParams{
				Page:    2,
				PerPage: 50,
			},
			expectError: false,
		},
		{
			name: "after parameter",
			params: map[string]any{
				"after": "cursor123",
			},
			expected: PaginationParams{
				Page:    1,
				PerPage: 30,
				After:   "cursor123",
			},
			expectError: false,
		},
		{
			name: "invalid after parameter",
			params: map[string]any{
				"after": float64(1),
			},
			expected:    PaginationParams{},
			expectError: true,
		},
		{
			name: "invalid page parameter",
			params: map[string]any{
				"page": "not-a-number",
			},
			expected:    PaginationParams{},
			expectError: true,
		},
		{
			name: "invalid perPage parameter",
			params: map[string]any{
				"perPage": "not-a-number",
			},
			expected:    PaginationParams{},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := createMCPRequest(tc.params)
			result, err := OptionalPaginationParams(request)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
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
			params:   map[string]any{"perPage": float64(50), "after": "cursor123"},
			expected: CursorPaginationParams{PerPage: 50, After: "cursor123"},
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
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestToGraphQLParams(t *testing.T) {
	ptr32 := func(v int32) *int32 { return &v }
	ptrStr := func(v string) *string { return &v }

	tests := []struct {
		name        string
		params      CursorPaginationParams
		expected    *GraphQLPaginationParams
		errContains string
	}{
		{
			name:     "no cursor",
			params:   CursorPaginationParams{PerPage: 30},
			expected: &GraphQLPaginationParams{First: ptr32(30)},
		},
		{
			name:     "with cursor",
			params:   CursorPaginationParams{PerPage: 10, After: "cursor123"},
			expected: &GraphQLPaginationParams{First: ptr32(10), After: ptrStr("cursor123")},
		},
		{
			name:     "maximum perPage",
			params:   CursorPaginationParams{PerPage: 100},
			expected: &GraphQLPaginationParams{First: ptr32(100)},
		},
		{
			name:        "perPage above maximum",
			params:      CursorPaginationParams{PerPage: 101},
			errContains: "perPage value 101 exceeds maximum of 100",
		},
		{
			name:        "negative perPage",
			params:      CursorPaginationParams{PerPage: -1},
			errContains: "perPage value -1 cannot be negative",
		},
	}

	for _, tc := range tests {
		t.Run("cursor/"+tc.name, func(t *testing.T) {
			result, err := tc.params.ToGraphQLParams()
			if tc.errContains != "" {
				assert.EqualError(t, err, tc.errContains)
				assert.Nil(t, result)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})

		t.Run("rest/"+tc.name, func(t *testing.T) {
			rest := PaginationParams{Page: 3, PerPage: tc.params.PerPage, After: tc.params.After}
			result, err := rest.ToGraphQLParams()
			if tc.errContains != "" {
				assert.EqualError(t, err, tc.errContains)
				assert.Nil(t, result)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestPaginationToolOptions(t *testing.T) {
	tests := []struct {
		name     string
		option   mcp.ToolOption
		expected []string
	}{
		{name: "WithPagination", option: WithPagination(), expected: []string{"page", "perPage"}},
		{name: "WithCursorPagination", option: WithCursorPagination(), expected: []string{"perPage", "after"}},
		{name: "WithUnifiedPagination", option: WithUnifiedPagination(), expected: []string{"page", "perPage", "after"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tool := mcp.NewTool("test", tc.option)
			props := tool.InputSchema.Properties
			keys := make([]string, 0, len(props))
			for k := range props {
				keys = append(keys, k)
			}
			assert.ElementsMatch(t, tc.expected, keys)
			assert.Empty(t, tool.InputSchema.Required)

			if p, ok := props["perPage"].(map[string]any); ok {
				assert.Equal(t, "number", p["type"])
				assert.Equal(t, "Results per page for pagination (min 1, max 100)", p["description"])
				assert.EqualValues(t, 1, p["minimum"])
				assert.EqualValues(t, 100, p["maximum"])
			}
			if p, ok := props["after"].(map[string]any); ok {
				assert.Equal(t, "string", p["type"])
				assert.Equal(t, "Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs.", p["description"])
			}
			if p, ok := props["page"].(map[string]any); ok {
				assert.Equal(t, "number", p["type"])
				assert.Equal(t, "Page number for pagination (min 1)", p["description"])
				assert.EqualValues(t, 1, p["minimum"])
			}
		})
	}
}
