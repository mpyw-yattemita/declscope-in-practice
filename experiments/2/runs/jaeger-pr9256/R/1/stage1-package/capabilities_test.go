// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSearchCapabilities_FieldCount fails when a field is added to
// SearchCapabilities. It is a reminder, not a contract: when it fails, extend
// the forwarding tests of every Reader decorator, which enumerate each field
// (tracestoremetrics.TestReadMetricsDecorator_SearchCapabilities and
// queryinterceptor.TestReader_SearchCapabilities_ForwardsBackendDeclaration),
// then update the count here.
func TestSearchCapabilities_FieldCount(t *testing.T) {
	assert.Equal(t, 3, reflect.TypeFor[SearchCapabilities]().NumField(),
		"SearchCapabilities gained a field; extend the decorators' forwarding tests to cover it")
}

func TestSearchCapabilities_ZeroValueIsLeastCapable(t *testing.T) {
	var caps SearchCapabilities
	assert.False(t, caps.WithoutServiceName)
	assert.False(t, caps.SameSpanConjunction)
	assert.True(t, caps.Filter.IsEmpty())
}
