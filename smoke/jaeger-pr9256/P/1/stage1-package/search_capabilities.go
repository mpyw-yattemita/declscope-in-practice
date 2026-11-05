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
