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
| filter.go | 58 | NEW | ErrFilterUnsupported (var), ErrFilterInvalid (var), FilterCapabilities (type), FilterCapabilities.IsEmpty (method), FilterCapabilities.SupportsLevel (method), FilterCapabilities.SupportsOperator (method) |
| filter_test.go | 63 | NEW | TestFilterCapabilities_IsEmpty (func), TestFilterCapabilities_SupportsLevel (func), TestFilterCapabilities_SupportsOperator (func) |
| reader.go | 146 | CHANGED | Reader (type), GetTraceParams (type), TraceQueryParams (type), FoundTraceID (type), OperationQueryParams (type), Operation (type) |
| reader_test.go | 4 |  |  |
| search_capabilities.go | 29 | NEW | SearchCapabilities (type) |
| search_capabilities_test.go | 20 | NEW | TestSearchCapabilities_FieldCount (func) |
| summary.go | 60 |  | ServiceSummary (type), TraceSummary (type), UnsupportedTraceSummaries (type), UnsupportedTraceSummaries.FindTraceSummaries (method) |
| summary_test.go | 23 |  | TestUnsupportedTraceSummaries_FindTraceSummaries (func) |
| writer.go | 30 |  | Writer (type) |

#### pkg/filter.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"errors"
	"slices"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
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

// FilterCapabilities declares how much of a structured filter a Reader
// evaluates, by naming the levels and operators it serves.
//
// Nothing is implicit: a Reader opts in only by naming something, so an empty
// FilterCapabilities says the same as no FilterCapabilities. Callers ask
// IsEmpty rather than comparing against nil.
type FilterCapabilities struct {
	// Levels are the levels a reference may name. Empty means the Reader can
	// serve only unqualified references.
	Levels []expression.Level
	// Operators are the operators the Reader evaluates, including the boolean
	// combinators OpAnd, OpOr and OpNot.
	Operators []expression.Operator
}

// IsEmpty reports whether c names no level and no operator, which is the same
// as declaring no filter support at all. It is true for a nil receiver.
func (c *FilterCapabilities) IsEmpty() bool {
	return c == nil || (len(c.Levels) == 0 && len(c.Operators) == 0)
}

// SupportsLevel reports whether a reference may name level. The empty level, an
// unqualified reference, always reaches the Reader.
func (c FilterCapabilities) SupportsLevel(level expression.Level) bool {
	return level == "" || slices.Contains(c.Levels, level)
}

// SupportsOperator reports whether the Reader evaluates op. Only the listed
// operators are supported, the boolean combinators included: a Reader confined
// to conjunctive search lists OpAnd and omits OpOr and OpNot. Nesting is not
// declared, because and is associative and a caller flattens it before asking.
func (c FilterCapabilities) SupportsOperator(op expression.Operator) bool {
	return slices.Contains(c.Operators, op)
}

```

#### pkg/filter_test.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"testing"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
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

	assert.True(t, caps.SupportsLevel(""), "an unqualified reference always reaches the reader")
	assert.True(t, caps.SupportsLevel(expression.LevelSpan))
	assert.True(t, caps.SupportsLevel(expression.LevelResource))
	assert.False(t, caps.SupportsLevel(expression.LevelScope))
	assert.False(t, caps.SupportsLevel(expression.LevelEvent))

	var zero FilterCapabilities
	assert.True(t, zero.SupportsLevel(""), "the zero value still serves unqualified references")
	assert.False(t, zero.SupportsLevel(expression.LevelSpan))
}

func TestFilterCapabilities_SupportsOperator(t *testing.T) {
	conjunctive := FilterCapabilities{Operators: []expression.Operator{expression.OpAnd, expression.OpEq}}

	assert.True(t, conjunctive.SupportsOperator(expression.OpAnd))
	assert.True(t, conjunctive.SupportsOperator(expression.OpEq))
	assert.False(t, conjunctive.SupportsOperator(expression.OpOr), "OpOr is supported only when listed")
	assert.False(t, conjunctive.SupportsOperator(expression.OpNot), "OpNot is supported only when listed")
	assert.False(t, conjunctive.SupportsOperator(expression.OpRegex))

	var zero FilterCapabilities
	for _, op := range []expression.Operator{expression.OpAnd, expression.OpOr, expression.OpNot, expression.OpEq} {
		assert.False(t, zero.SupportsOperator(op), "the zero value supports no operator, not even %q", op)
	}
}

```

#### pkg/reader.go

```go
// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"context"
	"iter"
	"time"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
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
	// It returns the zero value when the reader declares no capability. A reader
	// whose backend sits behind an API that cannot be asked returns an error that
	// wraps errors.ErrUnsupported instead, so that callers do not trust a guess.
	//
	// A Reader that wraps another (a decorator, an interceptor) must forward this
	// call to the wrapped reader rather than answer for itself.
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

#### pkg/search_capabilities.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

// SearchCapabilities describes how a Reader's search methods behave where
// backends differ.
//
// Each capability is a field that defaults to false, so the zero value describes
// the least capable reader. A capability added later is therefore inherited as
// false by every implementation that does not declare it.
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
	// Filter means none of it: the reader serves only the legacy fields. An
	// empty FilterCapabilities means the same; ask Filter.IsEmpty().
	Filter *FilterCapabilities
}

```

#### pkg/search_capabilities_test.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSearchCapabilities_FieldCount fails when a field is added to
// SearchCapabilities. The forwarding tests of the Reader decorators
// (tracestoremetrics and queryinterceptor) enumerate every field; extend them
// for the new field, then update the count here.
func TestSearchCapabilities_FieldCount(t *testing.T) {
	assert.Equal(t, 3, reflect.TypeFor[SearchCapabilities]().NumField(),
		"SearchCapabilities gained a field: extend the decorators' forwarding tests to cover it")
}

```

Answer with one JSON object and nothing else:
{"findings": [{"file": "...", "line": 0, "problem": "...", "fix": "..."}]}
Give each finding a concrete file, line and fix. Leave out vague remarks.
An empty list is a valid answer.
