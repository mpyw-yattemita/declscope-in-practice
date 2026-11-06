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
// The decorators' forwarding tests enumerate every field, so when this fails, extend
// those tests to cover the new field and then update the count here:
//   - tracestoremetrics.TestReadMetricsDecorator_SearchCapabilities
//   - queryinterceptor.TestReader_SearchCapabilities_ForwardsBackendDeclaration
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
		{name: "nil pointer", caps: nil, want: true},
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
	tests := []struct {
		name  string
		caps  FilterCapabilities
		level expression.Level
		want  bool
	}{
		{name: "unqualified on zero value", caps: FilterCapabilities{}, level: "", want: true},
		{name: "qualified on zero value", caps: FilterCapabilities{}, level: expression.LevelSpan, want: false},
		{name: "unqualified with levels", caps: caps, level: "", want: true},
		{name: "listed level", caps: caps, level: expression.LevelResource, want: true},
		{name: "unlisted level", caps: caps, level: expression.LevelEvent, want: false},
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
		{name: "zero value lacks and", caps: FilterCapabilities{}, op: expression.OpAnd, want: false},
		{name: "zero value lacks eq", caps: FilterCapabilities{}, op: expression.OpEq, want: false},
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
