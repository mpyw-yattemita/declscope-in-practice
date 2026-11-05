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
