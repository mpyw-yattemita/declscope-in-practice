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
