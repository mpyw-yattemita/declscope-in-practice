This directory holds one Go package, `pkg/`. Some of its files were just added
or changed; they are marked in the map below, and their full text follows it.
You may read any other file under pkg/.

Review the package as a code reviewer would. List the refactoring that should
be done now. Focus on the new and changed code and how it fits the package.

### pkg/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| empty_test.go | 14 |  | TestMain (func) |
| factory.go | 25 |  | Factory (type), SyncBulkWriteConfig (type) |
| reader.go | 226 | CHANGED | ErrFilterUnsupported (var), ErrFilterInvalid (var), Reader (type), SearchCapabilities (type), FilterCapabilities (type), FilterCapabilities.IsEmpty (method), FilterCapabilities.SupportsLevel (method), FilterCapabilities.SupportsOperator (method), GetTraceParams (type), TraceQueryParams (type), FoundTraceID (type), OperationQueryParams (type), Operation (type) |
| reader_test.go | 92 | CHANGED | TestSearchCapabilities_FieldCount (func), TestFilterCapabilities_IsEmpty (func), TestFilterCapabilities_SupportsLevel (func), TestFilterCapabilities_SupportsOperator (func) |
| summary.go | 60 |  | ServiceSummary (type), TraceSummary (type), UnsupportedTraceSummaries (type), UnsupportedTraceSummaries.FindTraceSummaries (method) |
| summary_test.go | 23 |  | TestUnsupportedTraceSummaries_FindTraceSummaries (func) |
| writer.go | 30 |  | Writer (type) |

#### pkg/reader.go

```go
// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"context"
	"errors"
	"iter"
	"slices"
	"time"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	// ErrFilterUnsupported is returned for a well-formed query filter that the
	// storage cannot serve: a level it does not index, an operator it has not
	// implemented, or a boolean structure a flat index cannot evaluate. The query
	// is refused rather than approximated.
	ErrFilterUnsupported = errors.New("this storage backend cannot serve this query filter")

	// ErrFilterInvalid is returned for a query filter whose value does not fit the
	// field it compares.
	ErrFilterInvalid = errors.New("invalid query filter")
)

// Reader finds and loads traces and other data from storage.
type Reader interface {
	// GetTraces returns an iterator that retrieves all traces with given IDs.
	// The iterator is single-use: once consumed, it cannot be used again.
	//
	// Chunking requirements:
	// - A single ptrace.Traces chunk MUST NOT contain spans from multiple traces.
	// - Large traces MAY be split across multiple, *consecutive* ptrace.Traces chunks.
	// - Each returned ptrace.Traces object MUST NOT be empty.
	//
	// Edge cases:
	// - If no spans are found for any given trace ID, the ID is ignored.
	// - If none of the trace IDs are found in the storage, an empty iterator is returned.
	// - If an error is encountered, the iterator returns the error and stops.
	GetTraces(ctx context.Context, traceIDs ...GetTraceParams) iter.Seq2[[]ptrace.Traces, error]

	// GetServices returns all service names known to the backend from spans
	// within its retention period.
	GetServices(ctx context.Context) ([]string, error)

	// GetOperations returns all operation names for a given service
	// known to the backend from spans within its retention period.
	GetOperations(ctx context.Context, query OperationQueryParams) ([]Operation, error)

	// FindTraces returns an iterator that retrieves traces matching query parameters.
	// The iterator is single-use: once consumed, it cannot be used again.
	//
	// The chunking rules is the same as for GetTraces.
	//
	// If no matching traces are found, the function returns an empty iterator.
	// If an error is encountered, the iterator returns the error and stops.
	//
	// Whether all query filters (such as multiple tags) must apply to the same span within
	// a trace, or can be satisfied by different spans, is implementation-dependent; the
	// Reader declares which in SearchCapabilities.SameSpanConjunction.
	FindTraces(ctx context.Context, query TraceQueryParams) iter.Seq2[[]ptrace.Traces, error]

	// FindTraceIDs returns an iterator that retrieves IDs of traces matching query parameters.
	// The iterator is single-use: once consumed, it cannot be used again.
	//
	// If no matching traces are found, the function returns an empty iterator.
	// If an error is encountered, the iterator returns the error and stops.
	//
	// This function behaves identically to FindTraces, except that it returns only the list
	// of matching trace IDs. This is useful in some contexts, such as batch jobs, where a
	// large list of trace IDs may be queried first and then the full traces are loaded
	// in batches.
	FindTraceIDs(ctx context.Context, query TraceQueryParams) iter.Seq2[[]FoundTraceID, error]

	// FindTraceSummaries returns an iterator over lightweight summaries of the traces
	// matching the query parameters (the metadata shown in search-result lists). The
	// iterator is single-use: once consumed, it cannot be used again.
	//
	// Backends that can compute summaries natively (e.g. via a storage-side aggregation)
	// should do so. Backends that cannot must yield errors.ErrUnsupported (wrapped with
	// %w) as the first error, before any batch; the caller then falls back to FindTraces
	// plus client-side aggregation. Such backends can embed UnsupportedTraceSummaries to
	// get this behavior for free.
	//
	// The iterator streams result batches; each yielded batch may contain one or more
	// summaries, and implementations may yield incrementally rather than buffering all
	// results first.
	FindTraceSummaries(ctx context.Context, query TraceQueryParams) iter.Seq2[[]TraceSummary, error]

	// SearchCapabilities reports how this reader's search methods behave.
	//
	// A reader that cannot find out, for example because its backend sits behind an
	// API that has no way to ask, must return an error that wraps errors.ErrUnsupported
	// rather than a SearchCapabilities value that a caller might trust. A returned error
	// means "cannot say", which is different from the zero value, which means "declares
	// no capabilities".
	//
	// The context is used because the answer may come over the network.
	//
	// A Reader that wraps another Reader (a decorator) must forward this call to the
	// wrapped Reader, including its error, and must not answer for itself.
	SearchCapabilities(ctx context.Context) (SearchCapabilities, error)
}

// SearchCapabilities describes how a Reader's search methods behave where
// backends differ.
//
// The zero value describes the least capable reader. Each capability is a field that
// is false unless a reader sets it, so a capability added later is false for every
// implementation that does not opt in.
type SearchCapabilities struct {
	// WithoutServiceName is true when FindTraces, FindTraceIDs and
	// FindTraceSummaries accept a TraceQueryParams whose ServiceName is empty
	// and read it as "any service", rather than as an error or an empty result.
	WithoutServiceName bool

	// SameSpanConjunction is true when a conjunction (several entries in the
	// Attributes map, or an OpAnd in Filter) is satisfied within a single span.
	// False means the backend may satisfy different conjuncts from different
	// spans of the same trace, as a flat inverted index does. Callers report
	// this looser scoping rather than refusing the query.
	SameSpanConjunction bool

	// Filter is how much of TraceQueryParams.Filter the reader evaluates. A nil
	// Filter means none of it: the reader serves only the legacy fields.
	// A non-nil Filter that names nothing means the same; use
	// FilterCapabilities.IsEmpty rather than comparing against nil.
	Filter *FilterCapabilities
}

// FilterCapabilities declares how much of a structured filter a Reader
// evaluates, by naming the levels and operators it serves.
//
// Nothing is implicit: a Reader opts in only by naming a level or an operator,
// so the zero value serves no filter at all.
type FilterCapabilities struct {
	// Levels are the levels a reference may name. Empty means the Reader can
	// serve only unqualified references.
	Levels []expression.Level
	// Operators are the operators the Reader evaluates, including the boolean
	// combinators OpAnd, OpOr and OpNot. A Reader confined to conjunctive search
	// lists OpAnd and omits OpOr and OpNot. Nesting is not declared: OpAnd is
	// associative, so a caller flattens nested conjunctions before asking.
	Operators []expression.Operator
}

// IsEmpty reports whether c names no level and no operator, which declares
// the same as no FilterCapabilities at all. It is safe to call on a nil pointer.
func (c *FilterCapabilities) IsEmpty() bool {
	return c == nil || (len(c.Levels) == 0 && len(c.Operators) == 0)
}

// SupportsLevel reports whether a reference may name level. The empty level
// (an unqualified reference) is always supported; any other level must be
// listed in Levels.
func (c FilterCapabilities) SupportsLevel(level expression.Level) bool {
	return level == "" || slices.Contains(c.Levels, level)
}

// SupportsOperator reports whether op is listed in Operators. The boolean
// combinators are not implied: the zero value supports no operator, not even
// OpAnd.
func (c FilterCapabilities) SupportsOperator(op expression.Operator) bool {
	return slices.Contains(c.Operators, op)
}

// GetTraceParams contains single-trace parameters for a GetTraces request.
// Some storage backends (e.g. Tempo) perform GetTraces much more efficiently
// if they know the approximate time range of the trace.
type GetTraceParams struct {
	// TraceID is the ID of the trace to retrieve. Required.
	TraceID pcommon.TraceID
	// Start of the time interval to search for trace ID. Optional.
	Start time.Time
	// End of the time interval to search for trace ID. Optional.
	End time.Time
}

// TraceQueryParams contains query parameters to find traces. For a detailed
// definition of each field in this message, refer to `TraceQueryParameters` in `jaeger.api_v3`
// (https://github.com/jaegertracing/jaeger-idl/blob/main/proto/api_v3/query_service.proto).
type TraceQueryParams struct {
	ServiceName   string
	OperationName string
	// Attributes must initialized with pcommon.NewMap() before use.
	Attributes   pcommon.Map
	StartTimeMin time.Time
	StartTimeMax time.Time
	DurationMin  time.Duration
	DurationMax  time.Duration
	SearchDepth  int

	// Filter is the structured query filter. It is mutually exclusive with the
	// predicate fields ServiceName, OperationName, Attributes and the duration
	// bounds, so a reader sees one filtering model, not a mix of the two.
	Filter *expression.Call
}

// FoundTraceID is a wrapper around trace ID returned from FindTraceIDs
// with an optional time range that may be used in GetTraces calls.
//
// The time range is provided as an optimization hint for some storage backends
// that can perform more efficient queries when they know the approximate time range.
// The value should not be used for precise time-based filtering or assumptions.
// It is meant as a rough boundary and may not be populated in all cases.
type FoundTraceID struct {
	TraceID pcommon.TraceID
	Start   time.Time
	End     time.Time
}

// OperationQueryParams contains parameters of query operations, empty spanKind means get operations for all kinds of span.
type OperationQueryParams struct {
	ServiceName string
	SpanKind    string
}

// Operation contains operation name and span kind
type Operation struct {
	Name     string
	SpanKind string
}

```

#### pkg/reader_test.go

```go
// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
	"github.com/stretchr/testify/assert"
)

// TestSearchCapabilities_FieldCount fails when a field is added to SearchCapabilities.
// Decorators must forward every field, and their forwarding tests enumerate every
// field. When this test fails, extend those tests to cover the new field, then
// update the count here:
//   - TestReadMetricsDecorator_SearchCapabilities (tracestoremetrics)
//   - TestReader_SearchCapabilities_ForwardsBackendDeclaration
//     (cmd/jaeger/internal/extension/jaegerquery/queryinterceptor)
func TestSearchCapabilities_FieldCount(t *testing.T) {
	assert.Equal(t, 3, reflect.TypeFor[SearchCapabilities]().NumField(),
		"SearchCapabilities changed: extend the decorators' forwarding tests to cover every field")
}

func TestFilterCapabilities_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		caps *FilterCapabilities
		want bool
	}{
		{name: "nil", caps: nil, want: true},
		{name: "zero value", caps: &FilterCapabilities{}, want: true},
		{
			name: "empty non-nil slices",
			caps: &FilterCapabilities{Levels: []expression.Level{}, Operators: []expression.Operator{}},
			want: true,
		},
		{name: "one level", caps: &FilterCapabilities{Levels: []expression.Level{expression.LevelSpan}}, want: false},
		{name: "one operator", caps: &FilterCapabilities{Operators: []expression.Operator{expression.OpEq}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.caps.IsEmpty())
		})
	}
}

func TestFilterCapabilities_SupportsLevel(t *testing.T) {
	spanOnly := FilterCapabilities{Levels: []expression.Level{expression.LevelSpan}}
	tests := []struct {
		name  string
		caps  FilterCapabilities
		level expression.Level
		want  bool
	}{
		{name: "zero value, unqualified", caps: FilterCapabilities{}, level: "", want: true},
		{name: "zero value, span", caps: FilterCapabilities{}, level: expression.LevelSpan, want: false},
		{name: "listed level", caps: spanOnly, level: expression.LevelSpan, want: true},
		{name: "unlisted level", caps: spanOnly, level: expression.LevelResource, want: false},
		{name: "unqualified with levels listed", caps: spanOnly, level: "", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.caps.SupportsLevel(tt.level))
		})
	}
}

func TestFilterCapabilities_SupportsOperator(t *testing.T) {
	conjunctive := FilterCapabilities{Operators: []expression.Operator{expression.OpAnd, expression.OpEq}}
	tests := []struct {
		name string
		caps FilterCapabilities
		op   expression.Operator
		want bool
	}{
		{name: "zero value, and", caps: FilterCapabilities{}, op: expression.OpAnd, want: false},
		{name: "zero value, eq", caps: FilterCapabilities{}, op: expression.OpEq, want: false},
		{name: "zero value, empty operator", caps: FilterCapabilities{}, op: "", want: false},
		{name: "listed combinator", caps: conjunctive, op: expression.OpAnd, want: true},
		{name: "listed comparison", caps: conjunctive, op: expression.OpEq, want: true},
		{name: "unlisted or", caps: conjunctive, op: expression.OpOr, want: false},
		{name: "unlisted not", caps: conjunctive, op: expression.OpNot, want: false},
		{name: "unlisted comparison", caps: conjunctive, op: expression.OpRegex, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.caps.SupportsOperator(tt.op))
		})
	}
}

```

Answer with one JSON object and nothing else:
{"findings": [{"file": "...", "line": 0, "problem": "...", "fix": "..."}]}
Give each finding a concrete file, line and fix. Leave out vague remarks.
An empty list is a valid answer.
