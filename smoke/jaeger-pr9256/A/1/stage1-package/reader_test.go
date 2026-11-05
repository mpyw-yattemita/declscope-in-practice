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
