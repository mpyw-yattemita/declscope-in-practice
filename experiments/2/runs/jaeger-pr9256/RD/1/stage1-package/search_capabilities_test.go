// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSearchCapabilitiesFields fails when a field is added to
// SearchCapabilities. When it does, extend the forwarding tests of the Reader
// decorators, which enumerate every field, then add the field here:
//   - internal/storage/v2/api/tracestore/tracestoremetrics
//   - cmd/jaeger/internal/extension/jaegerquery/queryinterceptor
func TestSearchCapabilitiesFields(t *testing.T) {
	typ := reflect.TypeFor[SearchCapabilities]()
	fields := make([]string, 0, typ.NumField())
	for i := range typ.NumField() {
		fields = append(fields, typ.Field(i).Name)
	}
	assert.Equal(t, []string{"WithoutServiceName", "SameSpanConjunction", "Filter"}, fields)
}

func TestSearchCapabilitiesZeroValue(t *testing.T) {
	zero := SearchCapabilities{}
	assert.False(t, zero.WithoutServiceName,
		"the zero value must describe the least capable reader")
	assert.False(t, zero.SameSpanConjunction,
		"the zero value must describe the least capable reader")
	assert.True(t, zero.Filter.IsEmpty(),
		"the zero value must describe the least capable reader")
}
