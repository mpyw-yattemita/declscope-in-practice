// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSearchCapabilitiesFields fails when a field is added to
// SearchCapabilities. When it does, extend the forwarding tests of every
// Reader decorator (tracestoremetrics and queryinterceptor), which enumerate
// each field, and then update the list below.
func TestSearchCapabilitiesFields(t *testing.T) {
	typ := reflect.TypeFor[SearchCapabilities]()
	fields := make([]string, typ.NumField())
	for i := range fields {
		fields[i] = typ.Field(i).Name
	}
	assert.Equal(t, []string{"WithoutServiceName", "SameSpanConjunction", "Filter"}, fields,
		"SearchCapabilities changed: extend the Reader decorators' forwarding tests to cover every field")
}

func TestSearchCapabilitiesZeroValue(t *testing.T) {
	assert.False(t, SearchCapabilities{}.WithoutServiceName)
	assert.False(t, SearchCapabilities{}.SameSpanConjunction)
	assert.True(t, SearchCapabilities{}.Filter.IsEmpty())
}
