# storage: let a trace reader declare which query filters it evaluates

The structured query filter of RFC 0005 is coming to trace search: a boolean expression over level-qualified attributes and built-in fields, from the `github.com/jaegertracing/jaeger-idl/query/expression/v1` package (imported as `expression`; the repository already depends on a jaeger-idl version that has it). The storage interface needs a place to carry a filter and, more importantly, a way for each backend to say how much of one it can evaluate. Backends differ sharply here: ClickHouse and Elasticsearch/OpenSearch can serve a nested boolean expression, while the flat index backends can serve a conjunction and nothing else. Without a declaration, a caller can neither avoid a query that cannot be answered nor interpret a surprising result.

Search-without-service-name already solved this problem with `SearchCapabilities`. Please extend that mechanism instead of inventing a second one.

## Public API

```go
// ErrFilterUnsupported is returned for a well-formed query filter that the
// storage cannot serve: a level it does not index, an operator it has not
// implemented, or a boolean structure a flat index cannot evaluate. The query
// is refused rather than approximated.
var ErrFilterUnsupported = errors.New("this storage backend cannot serve this query filter")

// ErrFilterInvalid is returned for a query filter whose value does not fit the
// field it compares.
var ErrFilterInvalid = errors.New("invalid query filter")

type SearchCapabilities struct {
	WithoutServiceName bool // existing

	// SameSpanConjunction is true when a conjunction (several entries in the
	// Attributes map, or an OpAnd in Filter) is satisfied within a single span.
	// False means the backend may satisfy different conjuncts from different
	// spans of the same trace, as a flat inverted index does. Callers report
	// this looser scoping rather than refusing the query.
	SameSpanConjunction bool

	// Filter is how much of TraceQueryParams.Filter the reader evaluates. A nil
	// Filter means none of it: the reader serves only the legacy fields.
	Filter *FilterCapabilities
}

// FilterCapabilities declares how much of a structured filter a Reader
// evaluates, by naming the levels and operators it serves.
type FilterCapabilities struct {
	// Levels are the levels a reference may name. Empty means the Reader can
	// serve only unqualified references.
	Levels []expression.Level
	// Operators are the operators the Reader evaluates, including the boolean
	// combinators OpAnd, OpOr and OpNot.
	Operators []expression.Operator
}

func (c *FilterCapabilities) IsEmpty() bool
func (c FilterCapabilities) SupportsLevel(level expression.Level) bool
func (c FilterCapabilities) SupportsOperator(op expression.Operator) bool

type TraceQueryParams struct {
	// ... existing fields ...

	// Filter is the structured query filter. It is mutually exclusive with the
	// predicate fields ServiceName, OperationName, Attributes and the duration
	// bounds, so a reader sees one filtering model, not a mix of the two.
	Filter *expression.Call
}
```

`SearchCapabilities` now has exactly these three fields. The field-count test from the previous change must be updated to match.

## Behaviour

- Nothing is implicit. A reader opts in only by naming something, so an empty `FilterCapabilities` says the same as no `FilterCapabilities`. `IsEmpty` returns true for a nil pointer, for a zero value, and for empty (non-nil) slices, and false as soon as one level or one operator is named. Callers ask `IsEmpty` rather than comparing against nil. This keeps the zero value of `SearchCapabilities` the least capable reader.
- `SupportsLevel` returns true for the empty level (an unqualified reference always reaches the reader), and otherwise only for a level listed in `Levels`.
- `SupportsOperator` returns true only for an operator listed in `Operators`. The boolean combinators are declared like any other operator: a reader confined to conjunctive search lists `OpAnd` and omits `OpOr` and `OpNot`. The zero value supports no operator, not even `OpAnd`. Nesting is not declared, because `and` is associative and a caller flattens it before asking.
- No reader populates the declaration yet and nothing reads it; the query service starts using it in a later change.

The rest of the repository needs no change. Do not change code outside the package. Add unit tests for the three methods.

# Scope of the change

Keep this change to what the task needs. Do not refactor existing code: do not move, rename or split existing declarations or files unless the task requires it.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./internal/storage/v2/api/tracestore
```
