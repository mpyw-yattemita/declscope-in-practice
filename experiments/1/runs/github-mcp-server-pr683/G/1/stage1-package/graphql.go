package github

// DefaultGraphQLPageSize is the number of items a GraphQL tool requests per page
// when the caller does not set "perPage".
const DefaultGraphQLPageSize = 30

// GraphQLPaginationParams holds the arguments of a GraphQL connection field.
// pagination.go builds it from the pagination parameters of a request.
// https://docs.github.com/en/graphql/reference/objects#connection
type GraphQLPaginationParams struct {
	First *int32
	After *string
}
