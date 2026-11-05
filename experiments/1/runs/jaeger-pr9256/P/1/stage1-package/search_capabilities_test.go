// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package tracestore

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSearchCapabilities_FieldCount fails when a field is added to
// SearchCapabilities. When it fails, extend the forwarding tests of every
// Reader decorator, which enumerate each field, then update the count here:
//   - tracestoremetrics/reader_metrics_test.go
//   - cmd/jaeger/internal/extension/jaegerquery/queryinterceptor/reader_decorator_test.go
func TestSearchCapabilities_FieldCount(t *testing.T) {
	assert.Equal(t, 3, reflect.TypeFor[SearchCapabilities]().NumField(),
		"SearchCapabilities gained a field: extend the decorators' forwarding tests")
}
