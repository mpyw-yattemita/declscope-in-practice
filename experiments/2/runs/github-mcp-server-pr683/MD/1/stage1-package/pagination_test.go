package github

import (
	"testing"

	"github.com/google/go-github/v73/github"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCursorPaginationParams_ToGraphQLParams(t *testing.T) {
	tests := []struct {
		name        string
		params      CursorPaginationParams
		expected    *GraphQLPaginationParams
		errContains string
	}{
		{
			name:     "perPage without cursor",
			params:   CursorPaginationParams{PerPage: DefaultGraphQLPageSize},
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(30))},
		},
		{
			name:     "perPage with cursor",
			params:   CursorPaginationParams{PerPage: 100, After: "cursor-abc"},
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(100)), After: github.Ptr("cursor-abc")},
		},
		{
			name:        "perPage above maximum",
			params:      CursorPaginationParams{PerPage: 101},
			errContains: "perPage value 101 exceeds maximum of 100",
		},
		{
			name:        "negative perPage",
			params:      CursorPaginationParams{PerPage: -5},
			errContains: "perPage value -5 cannot be negative",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.params.ToGraphQLParams()
			if tc.errContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errContains)
				assert.Nil(t, result)
				return
			}
			require.NoError(t, err)
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
			expected: CursorPaginationParams{PerPage: 30},
		},
		{
			name:     "perPage and after parameters",
			params:   map[string]any{"perPage": float64(50), "after": "cursor-abc"},
			expected: CursorPaginationParams{PerPage: 50, After: "cursor-abc"},
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

func TestWithCursorPagination(t *testing.T) {
	tool := mcp.NewTool("test_tool", WithCursorPagination())

	assert.NotContains(t, tool.InputSchema.Properties, "page")

	perPage, ok := tool.InputSchema.Properties["perPage"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "number", perPage["type"])
	assert.Equal(t, "Results per page for pagination (min 1, max 100)", perPage["description"])
	assert.Equal(t, float64(1), perPage["minimum"])
	assert.Equal(t, float64(100), perPage["maximum"])

	after, ok := tool.InputSchema.Properties["after"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "string", after["type"])
	assert.Equal(t, "Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs.", after["description"])
}

func TestWithUnifiedPagination(t *testing.T) {
	tool := mcp.NewTool("test_tool", WithUnifiedPagination())

	page, ok := tool.InputSchema.Properties["page"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "number", page["type"])
	assert.Equal(t, "Page number for pagination (min 1)", page["description"])
	assert.Equal(t, float64(1), page["minimum"])

	assert.Contains(t, tool.InputSchema.Properties, "perPage")
	assert.Contains(t, tool.InputSchema.Properties, "after")
}
