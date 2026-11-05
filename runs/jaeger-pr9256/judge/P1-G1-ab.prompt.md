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
| filter.go | 56 | NEW | ErrFilterUnsupported (var), ErrFilterInvalid (var), FilterCapabilities (type), FilterCapabilities.IsEmpty (method), FilterCapabilities.SupportsLevel (method), FilterCapabilities.SupportsOperator (method) |
| filter_test.go | 67 | NEW | TestFilterCapabilities_IsEmpty (func), TestFilterCapabilities_SupportsLevel (func), TestFilterCapabilities_SupportsOperator (func) |
| reader.go | 147 | CHANGED | Reader (type), GetTraceParams (type), TraceQueryParams (type), FoundTraceID (type), OperationQueryParams (type), Operation (type) |
| reader_test.go | 4 |  |  |
| search_capabilities.go | 29 | NEW | SearchCapabilities (type) |
| search_capabilities_test.go | 21 | NEW | TestSearchCapabilities_FieldCount (func) |
| summary.go | 60 |  | ServiceSummary (type), TraceSummary (type), UnsupportedTraceSummaries (type), UnsupportedTraceSummaries.FindTraceSummaries (method) |
| summary_test.go | 23 |  | TestUnsupportedTraceSummaries_FindTraceSummaries (func) |
| writer.go | 30 |  | Writer (type) |

### y/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| empty_test.go | 14 |  | TestMain (func) |
| factory.go | 25 |  | Factory (type), SyncBulkWriteConfig (type) |
| filter.go | 57 | NEW | ErrFilterUnsupported (var), ErrFilterInvalid (var), FilterCapabilities (type), FilterCapabilities.IsEmpty (method), FilterCapabilities.SupportsLevel (method), FilterCapabilities.SupportsOperator (method) |
| filter_test.go | 69 | NEW | TestFilterCapabilitiesIsEmpty (func), TestFilterCapabilitiesSupportsLevel (func), TestFilterCapabilitiesSupportsOperator (func) |
| reader.go | 148 | CHANGED | Reader (type), GetTraceParams (type), TraceQueryParams (type), FoundTraceID (type), OperationQueryParams (type), Operation (type) |
| reader_test.go | 4 |  |  |
| search_capabilities.go | 30 | NEW | SearchCapabilities (type) |
| search_capabilities_test.go | 33 | NEW | TestSearchCapabilitiesZeroValueDeclaresNothing (func), TestSearchCapabilitiesFields (func) |
| summary.go | 60 |  | ServiceSummary (type), TraceSummary (type), UnsupportedTraceSummaries (type), UnsupportedTraceSummaries.FindTraceSummaries (method) |
| summary_test.go | 23 |  | TestUnsupportedTraceSummaries_FindTraceSummaries (func) |
| writer.go | 30 |  | Writer (type) |

#### x/filter.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"errors"
	"slices"

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
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
// Nothing is implicit: a Reader opts in only by naming something, so an empty
// FilterCapabilities says the same as none. Ask IsEmpty rather than comparing
// against nil.
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

// IsEmpty reports whether c names no level and no operator. It is true for a
// nil pointer, for the zero value and for empty slices.
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

```

#### x/filter_test.go

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
		{name: "nil pointer", caps: nil, want: true},
		{name: "zero value", caps: &FilterCapabilities{}, want: true},
		{
			name: "empty slices",
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
	var zero FilterCapabilities
	assert.True(t, zero.SupportsLevel(""), "an unqualified reference always reaches the reader")
	assert.False(t, zero.SupportsLevel(expression.LevelSpan))

	caps := FilterCapabilities{Levels: []expression.Level{expression.LevelSpan, expression.LevelResource}}
	assert.True(t, caps.SupportsLevel(""))
	assert.True(t, caps.SupportsLevel(expression.LevelSpan))
	assert.True(t, caps.SupportsLevel(expression.LevelResource))
	assert.False(t, caps.SupportsLevel(expression.LevelEvent))
}

func TestFilterCapabilities_SupportsOperator(t *testing.T) {
	var zero FilterCapabilities
	assert.False(t, zero.SupportsOperator(expression.OpAnd), "the zero value supports no operator, not even and")
	assert.False(t, zero.SupportsOperator(expression.OpEq))

	conjunctive := FilterCapabilities{Operators: []expression.Operator{expression.OpAnd, expression.OpEq}}
	assert.True(t, conjunctive.SupportsOperator(expression.OpAnd))
	assert.True(t, conjunctive.SupportsOperator(expression.OpEq))
	assert.False(t, conjunctive.SupportsOperator(expression.OpOr))
	assert.False(t, conjunctive.SupportsOperator(expression.OpNot))
	assert.False(t, conjunctive.SupportsOperator(expression.OpRegex))
}

```

#### x/reader.go

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
	// A reader that cannot tell, for example because its backend sits behind an
	// API that cannot be asked, returns an error that wraps errors.ErrUnsupported
	// rather than a value that a caller might trust. The context is there because
	// the answer may come over the network.
	//
	// A Reader that wraps another Reader must forward this call to it, and must
	// not answer for it.
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

#### x/search_capabilities.go

```go
// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

// SearchCapabilities describes how a Reader's search methods behave where
// backends differ. Reader.SearchCapabilities returns it.
//
// The zero value is the least capable reader: every bool is false and Filter
// is nil. A field added later is therefore unset for every implementation that
// does not set it, so a reader never claims a capability it was not written for.
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
	// empty FilterCapabilities means the same; ask Filter.IsEmpty.
	Filter *FilterCapabilities
}

```

#### x/search_capabilities_test.go

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
// SearchCapabilities. When it fails, extend the forwarding tests of every
// Reader decorator, which enumerate each field, then update the count here:
//   - tracestoremetrics/reader_metrics_test.go
//   - cmd/jaeger/internal/extension/jaegerquery/queryinterceptor/reader_decorator_test.go
func TestSearchCapabilities_FieldCount(t *testing.T) {
	assert.Equal(t, 3, reflect.TypeFor[SearchCapabilities]().NumField(),
		"SearchCapabilities gained a field: extend the decorators' forwarding tests")
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

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
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
// Nothing is implicit: a Reader opts in only by naming something, so an empty
// FilterCapabilities declares the same as none.
type FilterCapabilities struct {
	// Levels are the levels a reference may name. Empty means the Reader can
	// serve only unqualified references.
	Levels []expression.Level
	// Operators are the operators the Reader evaluates, including the boolean
	// combinators OpAnd, OpOr and OpNot.
	Operators []expression.Operator
}

// IsEmpty reports whether c names no level and no operator. It is true for a
// nil pointer, for the zero value, and for empty slices. Callers ask IsEmpty
// rather than comparing against nil.
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
// combinators are declared like any other operator, so the zero value supports
// none of them, not even OpAnd. Nesting is not declared: and is associative,
// so a caller flattens it before asking.
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

	expression "github.com/jaegertracing/jaeger-idl/query/expression/v1"
	"github.com/stretchr/testify/assert"
)

func TestFilterCapabilitiesIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		caps *FilterCapabilities
		want bool
	}{
		{name: "nil", caps: nil, want: true},
		{name: "zero value", caps: &FilterCapabilities{}, want: true},
		{
			name: "empty slices",
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

func TestFilterCapabilitiesSupportsLevel(t *testing.T) {
	caps := FilterCapabilities{Levels: []expression.Level{expression.LevelSpan, expression.LevelResource}}
	assert.True(t, caps.SupportsLevel(""), "an unqualified reference is always supported")
	assert.True(t, caps.SupportsLevel(expression.LevelSpan))
	assert.True(t, caps.SupportsLevel(expression.LevelResource))
	assert.False(t, caps.SupportsLevel(expression.LevelScope))
	assert.False(t, caps.SupportsLevel(expression.LevelEvent))

	var zero FilterCapabilities
	assert.True(t, zero.SupportsLevel(""))
	assert.False(t, zero.SupportsLevel(expression.LevelSpan))
}

func TestFilterCapabilitiesSupportsOperator(t *testing.T) {
	conjunctive := FilterCapabilities{Operators: []expression.Operator{expression.OpAnd, expression.OpEq}}
	assert.True(t, conjunctive.SupportsOperator(expression.OpAnd))
	assert.True(t, conjunctive.SupportsOperator(expression.OpEq))
	assert.False(t, conjunctive.SupportsOperator(expression.OpOr))
	assert.False(t, conjunctive.SupportsOperator(expression.OpNot))
	assert.False(t, conjunctive.SupportsOperator(expression.OpRegex))

	var zero FilterCapabilities
	for _, op := range []expression.Operator{expression.OpAnd, expression.OpOr, expression.OpNot, expression.OpEq} {
		assert.False(t, zero.SupportsOperator(op), "zero value must not support %q", op)
	}
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
	// A reader whose backend cannot be asked (for example, one behind an API
	// with no capability discovery) returns an error that wraps
	// errors.ErrUnsupported, rather than a value a caller might trust.
	// A zero SearchCapabilities with a nil error means the reader declares
	// no capability.
	//
	// A Reader that wraps another Reader, such as a decorator or an
	// interceptor, must forward this call to the wrapped Reader.
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
// backends differ.
//
// The zero value is the least capable reader: every flag is false and Filter
// is nil. A new capability is added as a new field, so every existing
// implementation declares it as absent until it opts in.
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
	// empty FilterCapabilities means the same, so callers ask
	// FilterCapabilities.IsEmpty rather than comparing against nil.
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

func TestSearchCapabilitiesZeroValueDeclaresNothing(t *testing.T) {
	caps := SearchCapabilities{}
	assert.False(t, caps.WithoutServiceName)
	assert.False(t, caps.SameSpanConjunction)
	assert.True(t, caps.Filter.IsEmpty())
}

// TestSearchCapabilitiesFields fails when a field is added to SearchCapabilities.
// When it does, extend the forwarding tests of every Reader decorator, which
// enumerate each field, then add the new field to the list below:
//   - internal/storage/v2/api/tracestore/tracestoremetrics
//   - cmd/jaeger/internal/extension/jaegerquery/queryinterceptor
func TestSearchCapabilitiesFields(t *testing.T) {
	typ := reflect.TypeFor[SearchCapabilities]()
	fields := make([]string, 0, typ.NumField())
	for i := range typ.NumField() {
		fields = append(fields, typ.Field(i).Name)
	}
	assert.Equal(t, []string{"WithoutServiceName", "SameSpanConjunction", "Filter"}, fields,
		"SearchCapabilities changed: extend the decorators' forwarding tests to cover every field")
}

```

Answer with one JSON object and nothing else:
{"better": "x" | "y" | "tie", "reasons": ["...", "..."]}
Give two to five short reasons, each pointing at concrete files or types.
