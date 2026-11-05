package github

// DefaultGraphQLPageSize is the page size requested from GraphQL connections
// when the caller does not give one. OptionalCursorPaginationParams (pagination.go)
// uses it as the default "perPage".
const DefaultGraphQLPageSize = 30
