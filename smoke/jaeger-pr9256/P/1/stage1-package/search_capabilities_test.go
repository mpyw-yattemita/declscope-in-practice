// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCapabilities_ZeroValueDeclaresNothing(t *testing.T) {
	assert.False(t, SearchCapabilities{}.WithoutServiceName)
	assert.False(t, SearchCapabilities{}.SameSpanConjunction)
	assert.True(t, SearchCapabilities{}.Filter.IsEmpty())
}

// TestSearchCapabilities_FieldCount fails when a field is added to
// SearchCapabilities. When it fails, extend the forwarding tests of the
// decorators, which enumerate every field:
//   - tracestoremetrics: TestReadMetricsDecorator_SearchCapabilities
//   - queryinterceptor: TestReader_SearchCapabilities_ForwardsBackendDeclaration
//
// Then update the count below.
func TestSearchCapabilities_FieldCount(t *testing.T) {
	assert.Equal(t, 3, reflect.TypeFor[SearchCapabilities]().NumField(),
		"a field was added to SearchCapabilities: extend the decorators' forwarding tests, then update this count")
}
