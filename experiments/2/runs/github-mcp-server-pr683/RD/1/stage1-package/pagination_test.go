package github

import (
	"testing"

	"github.com/google/go-github/v73/github"
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
				"after": "Y3Vyc29yOnYyOpHOAA==",
			},
			expected: PaginationParams{
				Page:    1,
				PerPage: 30,
				After:   "Y3Vyc29yOnYyOpHOAA==",
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
			expected: CursorPaginationParams{PerPage: DefaultGraphQLPageSize},
		},
		{
			name:     "perPage and after parameters",
			params:   map[string]any{"perPage": float64(50), "after": "cursor"},
			expected: CursorPaginationParams{PerPage: 50, After: "cursor"},
		},
		{
			name:     "page parameter is ignored",
			params:   map[string]any{"page": float64(2)},
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
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestToGraphQLParams(t *testing.T) {
	tests := []struct {
		name        string
		perPage     int
		after       string
		expected    *GraphQLPaginationParams
		errContains string
	}{
		{
			name:     "perPage becomes first, no cursor",
			perPage:  30,
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(30))},
		},
		{
			name:     "after becomes the cursor",
			perPage:  100,
			after:    "cursor",
			expected: &GraphQLPaginationParams{First: github.Ptr(int32(100)), After: github.Ptr("cursor")},
		},
		{
			name:        "perPage above 100",
			perPage:     101,
			errContains: "perPage value 101 exceeds maximum of 100",
		},
		{
			name:        "negative perPage",
			perPage:     -1,
			errContains: "perPage value -1 cannot be negative",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cursorResult, cursorErr := CursorPaginationParams{PerPage: tc.perPage, After: tc.after}.ToGraphQLParams()
			// Page has no GraphQL counterpart, so it must not change the result.
			restResult, restErr := PaginationParams{Page: 3, PerPage: tc.perPage, After: tc.after}.ToGraphQLParams()

			if tc.errContains != "" {
				assert.EqualError(t, cursorErr, tc.errContains)
				assert.EqualError(t, restErr, tc.errContains)
				return
			}
			require.NoError(t, cursorErr)
			require.NoError(t, restErr)
			assert.Equal(t, tc.expected, cursorResult)
			assert.Equal(t, tc.expected, restResult)
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
			names := make([]string, 0, len(props))
			for name := range props {
				names = append(names, name)
			}
			assert.ElementsMatch(t, tc.expected, names)
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
