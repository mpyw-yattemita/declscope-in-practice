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
