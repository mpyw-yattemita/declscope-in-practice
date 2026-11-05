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
				"after": "cursor-1",
			},
			expected: PaginationParams{
				Page:    1,
				PerPage: 30,
				After:   "cursor-1",
			},
			expectError: false,
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
		{
			name: "invalid after parameter",
			params: map[string]any{
				"after": float64(1),
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
			expected: CursorPaginationParams{PerPage: 30},
		},
		{
			name:     "perPage and after",
			params:   map[string]any{"perPage": float64(10), "after": "cursor-1"},
			expected: CursorPaginationParams{PerPage: 10, After: "cursor-1"},
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
	tests := []struct {
		name          string
		params        CursorPaginationParams
		expectedFirst int32
		expectedAfter *string
		errContains   string
	}{
		{
			name:          "no cursor",
			params:        CursorPaginationParams{PerPage: 30},
			expectedFirst: 30,
		},
		{
			name:          "with cursor",
			params:        CursorPaginationParams{PerPage: 100, After: "cursor-1"},
			expectedFirst: 100,
			expectedAfter: ToStringPtr("cursor-1"),
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
		t.Run(tc.name, func(t *testing.T) {
			// The REST parameters convert the same way, ignoring the page number.
			rest := PaginationParams{Page: 3, PerPage: tc.params.PerPage, After: tc.params.After}
			for _, convert := range []func() (*GraphQLPaginationParams, error){
				tc.params.ToGraphQLParams,
				rest.ToGraphQLParams,
			} {
				result, err := convert()
				if tc.errContains != "" {
					assert.EqualError(t, err, tc.errContains)
					assert.Nil(t, result)
					continue
				}
				require.NoError(t, err)
				require.NotNil(t, result.First)
				assert.Equal(t, tc.expectedFirst, *result.First)
				assert.Equal(t, tc.expectedAfter, result.After)
			}
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
			var names []string
			for name := range tool.InputSchema.Properties {
				names = append(names, name)
			}
			assert.ElementsMatch(t, tc.expected, names)
			assert.Empty(t, tool.InputSchema.Required)

			if perPage, ok := tool.InputSchema.Properties["perPage"].(map[string]any); assert.True(t, ok) {
				assert.Equal(t, "number", perPage["type"])
				assert.Equal(t, "Results per page for pagination (min 1, max 100)", perPage["description"])
				assert.Equal(t, float64(1), perPage["minimum"])
				assert.Equal(t, float64(100), perPage["maximum"])
			}
			if after, ok := tool.InputSchema.Properties["after"].(map[string]any); ok {
				assert.Equal(t, "string", after["type"])
				assert.Equal(t, "Cursor for pagination. Use the endCursor from the previous page's PageInfo for GraphQL APIs.", after["description"])
			}
			if page, ok := tool.InputSchema.Properties["page"].(map[string]any); ok {
				assert.Equal(t, "number", page["type"])
				assert.Equal(t, "Page number for pagination (min 1)", page["description"])
				assert.Equal(t, float64(1), page["minimum"])
			}
		})
	}
}
