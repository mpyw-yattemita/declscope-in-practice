Two versions of the same Go package are in this directory, as `x/` and `y/`.
Both started from the same code and gained the same features, written by
different people. Compare how well each package is divided into files and types.

Judge the package as a whole, and how the new code fits into it:
- Does each file and each type have one clear job?
- Is closely related code kept together?
- Do the boundaries between files follow the boundaries between features?
- Does the new code fit the structure the package already had?

Behaviour and tests are known to be equivalent. Do not judge style, comments,
or naming conventions on their own, except where a name shows how the code is
divided.

Below is a map of each package: every file with its line count and its
top-level declarations. Files marked NEW were added and files marked CHANGED
were modified. The full text of every NEW and CHANGED file follows the maps.
You may read any other file under x/ or y/ if you need to.

### x/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| empty_test.go | 14 |  | TestMain (func) |
| factory.go | 25 |  | Factory (type), SyncBulkWriteConfig (type) |
| reader.go | 218 | CHANGED | ErrFilterUnsupported (var), ErrFilterInvalid (var), Reader (type), SearchCapabilities (type), FilterCapabilities (type), FilterCapabilities.IsEmpty (method), FilterCapabilities.SupportsLevel (method), FilterCapabilities.SupportsOperator (method), GetTraceParams (type), TraceQueryParams (type), FoundTraceID (type), OperationQueryParams (type), Operation (type) |
| reader_test.go | 81 | CHANGED | TestSearchCapabilities_FieldCount (func), TestFilterCapabilities_IsEmpty (func), TestFilterCapabilities_SupportsLevel (func), TestFilterCapabilities_SupportsOperator (func) |
| summary.go | 60 |  | ServiceSummary (type), TraceSummary (type), UnsupportedTraceSummaries (type), UnsupportedTraceSummaries.FindTraceSummaries (method) |
| summary_test.go | 23 |  | TestUnsupportedTraceSummaries_FindTraceSummaries (func) |
| writer.go | 30 |  | Writer (type) |

### y/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| empty_test.go | 14 |  | TestMain (func) |
| factory.go | 25 |  | Factory (type), SyncBulkWriteConfig (type) |
| filter.go | 57 | NEW | ErrFilterUnsupported (var), ErrFilterInvalid (var), FilterCapabilities (type), FilterCapabilities.IsEmpty (method), FilterCapabilities.SupportsLevel (method), FilterCapabilities.SupportsOperator (method) |
| filter_test.go | 58 | NEW | TestFilterCapabilities_IsEmpty (func), TestFilterCapabilities_SupportsLevel (func), TestFilterCapabilities_SupportsOperator (func) |
| reader.go | 147 | CHANGED | Reader (type), GetTraceParams (type), TraceQueryParams (type), FoundTraceID (type), OperationQueryParams (type), Operation (type) |
| reader_test.go | 4 |  |  |
| search_capabilities.go | 30 | NEW | SearchCapabilities (type) |
| search_capabilities_test.go | 29 | NEW | TestSearchCapabilities_ZeroValueDeclaresNothing (func), TestSearchCapabilities_FieldCount (func) |
| summary.go | 60 |  | ServiceSummary (type), TraceSummary (type), UnsupportedTraceSummaries (type), UnsupportedTraceSummaries.FindTraceSummaries (method) |
| summary_test.go | 23 |  | TestUnsupportedTraceSummaries_FindTraceSummaries (func) |
| writer.go | 30 |  | Writer (type) |

#### x/reader.go

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

// ErrFilterUnsupported is returned for a well-formed query filter that the
// storage cannot serve: a level it does not index, an operator it has not
// implemented, or a boolean structure a flat index cannot evaluate. The query
// is refused rather than approximated.
var ErrFilterUnsupported = errors.New("this storage backend cannot serve this query filter")

// ErrFilterInvalid is returned for a query filter whose value does not fit the
// field it compares.
var ErrFilterInvalid = errors.New("invalid query filter")

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
	// reader declares which in SearchCapabilities.SameSpanConjunction.
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
	// API that cannot be asked, must return an error that wraps errors.ErrUnsupported
	// rather than a guessed value. The zero SearchCapabilities with a nil error means
	// that the reader declares no capabilities, which is a different answer.
	//
	// A Reader that wraps another Reader (a decorator) must forward this call to the
	// wrapped Reader and return its answer, including its error, unchanged.
	SearchCapabilities(ctx context.Context) (SearchCapabilities, error)
}

// SearchCapabilities describes how a Reader's search methods behave where
// backends differ.
//
// Every field is a capability that is false by default, so the zero value
// describes the least capable reader. A capability added later is false for
// every existing implementation until that implementation declares it.
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
	Filter *FilterCapabilities
}

// FilterCapabilities declares how much of a structured filter a Reader
// evaluates, by naming the levels and operators it serves.
//
// Nothing is implicit: a Reader opts in only by naming a level or an operator,
// so an empty FilterCapabilities declares the same as a nil one. Callers ask
// IsEmpty rather than comparing against nil.
type FilterCapabilities struct {
	// Levels are the levels a reference may name. Empty means the Reader can
	// serve only unqualified references.
	Levels []expression.Level
	// Operators are the operators the Reader evaluates, including the boolean
	// combinators OpAnd, OpOr and OpNot. A Reader confined to conjunctive
	// search lists OpAnd and omits OpOr and OpNot. Nesting is not declared:
	// and is associative, so a caller flattens it before asking.
	Operators []expression.Operator
}

// IsEmpty reports whether c declares nothing: it is nil, or names no level and
// no operator.
func (c *FilterCapabilities) IsEmpty() bool {
	return c == nil || (len(c.Levels) == 0 && len(c.Operators) == 0)
}

// SupportsLevel reports whether a reference may name level. The empty level,
// an unqualified reference, is always supported.
func (c FilterCapabilities) SupportsLevel(level expression.Level) bool {
	return level == "" || slices.Contains(c.Levels, level)
}

// SupportsOperator reports whether op is listed in Operators. The zero value
// supports no operator, not even OpAnd.
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

#### x/reader_test.go

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
// The forwarding tests of the Reader decorators enumerate every field, so a new field
// must be added to them too:
//   - TestReadMetricsDecorator_SearchCapabilities in tracestoremetrics
//   - TestReader_SearchCapabilities_ForwardsBackendDeclaration in
//     cmd/jaeger/internal/extension/jaegerquery/queryinterceptor
//
// After extending those tests, update the expected count here.
func TestSearchCapabilities_FieldCount(t *testing.T) {
	assert.Equal(t, 3, reflect.TypeFor[SearchCapabilities]().NumField(),
		"SearchCapabilities changed: extend the decorators' forwarding tests, then update this count")
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
		{
			name: "one level",
			caps: &FilterCapabilities{Levels: []expression.Level{expression.LevelSpan}},
			want: false,
		},
		{
			name: "one operator",
			caps: &FilterCapabilities{Operators: []expression.Operator{expression.OpEq}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.caps.IsEmpty())
		})
	}
}

func TestFilterCapabilities_SupportsLevel(t *testing.T) {
	caps := FilterCapabilities{Levels: []expression.Level{expression.LevelSpan, expression.LevelResource}}
	assert.True(t, caps.SupportsLevel(""), "unqualified reference")
	assert.True(t, caps.SupportsLevel(expression.LevelSpan))
	assert.True(t, caps.SupportsLevel(expression.LevelResource))
	assert.False(t, caps.SupportsLevel(expression.LevelEvent))

	var zero FilterCapabilities
	assert.True(t, zero.SupportsLevel(""), "unqualified reference on zero value")
	assert.False(t, zero.SupportsLevel(expression.LevelSpan))
}

func TestFilterCapabilities_SupportsOperator(t *testing.T) {
	conjunctive := FilterCapabilities{Operators: []expression.Operator{expression.OpAnd, expression.OpEq}}
	assert.True(t, conjunctive.SupportsOperator(expression.OpAnd))
	assert.True(t, conjunctive.SupportsOperator(expression.OpEq))
	assert.False(t, conjunctive.SupportsOperator(expression.OpOr))
	assert.False(t, conjunctive.SupportsOperator(expression.OpNot))
	assert.False(t, conjunctive.SupportsOperator(expression.OpRegex))

	var zero FilterCapabilities
	assert.False(t, zero.SupportsOperator(expression.OpAnd), "zero value supports no operator, not even and")
	assert.False(t, zero.SupportsOperator(expression.OpEq))
}

```

#### y/filter.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"errors"
	"slices"

	"github.com/jaegertracing/jaeger-idl/query/expression/v1"
)

// ErrFilterUnsupported is returned for a well-formed query filter that the
// storage cannot serve: a level it does not index, an operator it has not
// implemented, or a boolean structure a flat index cannot evaluate. The query
// is refused rather than approximated.
var ErrFilterUnsupported = errors.New("this storage backend cannot serve this query filter")

// ErrFilterInvalid is returned for a query filter whose value does not fit the
// field it compares.
var ErrFilterInvalid = errors.New("invalid query filter")

// FilterCapabilities declares how much of a structured filter a Reader
// evaluates, by naming the levels and operators it serves.
//
// Nothing is implicit: a Reader opts in only by naming a level or an operator,
// so an empty FilterCapabilities declares the same as none at all. Callers ask
// IsEmpty rather than comparing against nil.
type FilterCapabilities struct {
	// Levels are the levels a reference may name. Empty means the Reader can
	// serve only unqualified references.
	Levels []expression.Level
	// Operators are the operators the Reader evaluates, including the boolean
	// combinators OpAnd, OpOr and OpNot.
	Operators []expression.Operator
}

// IsEmpty reports whether c declares nothing: c is nil, or it names no level
// and no operator.
func (c *FilterCapabilities) IsEmpty() bool {
	return c == nil || (len(c.Levels) == 0 && len(c.Operators) == 0)
}

// SupportsLevel reports whether a reference may name level. The empty level,
// an unqualified reference, is always supported; any other level only when it
// is listed in Levels.
func (c FilterCapabilities) SupportsLevel(level expression.Level) bool {
	return level == "" || slices.Contains(c.Levels, level)
}

// SupportsOperator reports whether op is listed in Operators. The boolean
// combinators are declared like any other operator, so the zero value does not
// support even OpAnd. Nesting is not declared: OpAnd is associative, and a
// caller flattens it before asking.
func (c FilterCapabilities) SupportsOperator(op expression.Operator) bool {
	return slices.Contains(c.Operators, op)
}

```

#### y/filter_test.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"testing"

	"github.com/jaegertracing/jaeger-idl/query/expression/v1"
	"github.com/stretchr/testify/assert"
)

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
	caps := FilterCapabilities{Levels: []expression.Level{expression.LevelSpan, expression.LevelResource}}
	assert.True(t, caps.SupportsLevel(""), "an unqualified reference is always supported")
	assert.True(t, caps.SupportsLevel(expression.LevelSpan))
	assert.True(t, caps.SupportsLevel(expression.LevelResource))
	assert.False(t, caps.SupportsLevel(expression.LevelEvent))

	var zero FilterCapabilities
	assert.True(t, zero.SupportsLevel(""), "an unqualified reference is always supported")
	assert.False(t, zero.SupportsLevel(expression.LevelSpan))
}

func TestFilterCapabilities_SupportsOperator(t *testing.T) {
	conjunctive := FilterCapabilities{Operators: []expression.Operator{expression.OpAnd, expression.OpEq}}
	assert.True(t, conjunctive.SupportsOperator(expression.OpAnd))
	assert.True(t, conjunctive.SupportsOperator(expression.OpEq))
	assert.False(t, conjunctive.SupportsOperator(expression.OpOr))
	assert.False(t, conjunctive.SupportsOperator(expression.OpNot))

	var zero FilterCapabilities
	assert.False(t, zero.SupportsOperator(expression.OpAnd), "the zero value supports no operator, not even and")
	assert.False(t, zero.SupportsOperator(expression.OpEq))
}

```

#### y/reader.go

```go
// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"context"
	"iter"
	"time"

	"github.com/jaegertracing/jaeger-idl/query/expression/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
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
	// There's currently an implementation-dependent ambiguity whether all query filters
	// (such as multiple tags) must apply to the same span within a trace, or can be satisfied
	// by different spans.
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
	// A Reader that cannot find out, for example because its backend sits behind
	// an API that cannot be asked, returns an error that wraps errors.ErrUnsupported
	// instead of a value that a caller might trust. The context is there because
	// the answer may come over the network.
	//
	// A Reader that wraps another Reader (a decorator) must forward this call to
	// the wrapped Reader, so that wrapping does not change the declaration.
	SearchCapabilities(ctx context.Context) (SearchCapabilities, error)
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

#### y/search_capabilities.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

// SearchCapabilities describes how a Reader's search methods behave where
// backends differ. Reader.SearchCapabilities returns it.
//
// The zero value is the least capable reader: every field is false or nil, and
// a field that is false or nil means the reader does not declare that
// behaviour. A new capability is added as a new field, so every existing
// implementation declares nothing for it until it opts in.
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
	// Callers ask Filter.IsEmpty rather than comparing against nil.
	Filter *FilterCapabilities
}

```

#### y/search_capabilities_test.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCapabilities_ZeroValueDeclaresNothing(t *testing.T) {
	assert.False(t, SearchCapabilities{}.WithoutServiceName)
	assert.False(t, SearchCapabilities{}.SameSpanConjunction)
	assert.True(t, SearchCapabilities{}.Filter.IsEmpty())
}

// TestSearchCapabilities_FieldCount fails when a field is added to
// SearchCapabilities. When it fails, extend the forwarding tests of the
// decorators, which enumerate every field:
//   - tracestoremetrics: TestReadMetricsDecorator_SearchCapabilities
//   - queryinterceptor: TestReader_SearchCapabilities_ForwardsBackendDeclaration
//
// Then update the count below.
func TestSearchCapabilities_FieldCount(t *testing.T) {
	assert.Equal(t, 3, reflect.TypeFor[SearchCapabilities]().NumField(),
		"a field was added to SearchCapabilities: extend the decorators' forwarding tests, then update this count")
}

```

Answer with one JSON object and nothing else:
{"better": "x" | "y" | "tie", "reasons": ["...", "..."]}
Give two to five short reasons, each pointing at concrete files or types.
