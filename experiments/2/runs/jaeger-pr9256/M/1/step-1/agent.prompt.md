# storage: let a trace reader declare that search can omit the service name

Trace search cannot run until a service is selected, so a user who knows a tag but not which service emits it has to search every service in turn (jaegertracing/jaeger#423). The requirement comes from Cassandra, whose indices are all keyed by service name. It is not inherent to the search contract, and backends differ today in a way that nothing documents: Elasticsearch/OpenSearch, ClickHouse and the in-memory store already answer a query without a service name, while Cassandra silently returns no rows.

As a first step, each storage backend must be able to declare what its search methods accept. Please add that declaration to the `tracestore` package's `Reader` interface.

## Public API

```go
type Reader interface {
	// ... existing methods ...

	// SearchCapabilities reports how this reader's search methods behave.
	SearchCapabilities(ctx context.Context) (SearchCapabilities, error)
}

// SearchCapabilities describes how a Reader's search methods behave where
// backends differ.
type SearchCapabilities struct {
	// WithoutServiceName is true when FindTraces, FindTraceIDs and
	// FindTraceSummaries accept a TraceQueryParams whose ServiceName is empty
	// and read it as "any service", rather than as an error or an empty result.
	WithoutServiceName bool
}
```

`SearchCapabilities` has exactly this one field for now.

## Behaviour and design constraints

- It is a method on the interface, not an optional interface that callers type-assert. Readers are routinely wrapped (by the metrics decorator and by query interceptors), and a type assertion against a wrapper silently answers "no capability". With a method, a decorator that forgets to forward the call is a compile error.
- It returns a struct, not one boolean method per capability, so that a later capability is a new field that every implementation inherits as `false`. The zero value of `SearchCapabilities` is therefore the least capable reader. Document this on the type.
- It returns an error because "declares nothing" and "cannot say" are different answers. A reader whose backend sits behind an API that cannot be asked returns an error that wraps `errors.ErrUnsupported`, rather than a value that a caller might trust. It takes a context because the answer may come over the network.
- A Reader that wraps another must forward the call. Document this on the method.

The implementations outside this package are already done: every backend (Elasticsearch/OpenSearch, ClickHouse and memory declare `WithoutServiceName: true`; Cassandra, the v1 adapter and remote gRPC storage declare nothing or report `ErrUnsupported`), the metrics decorator, the query interceptor, the generated mocks, and the query service, which passes the answer to the UI. They already call this API, so they compile once the package has it. Do not change code outside the package.

Add a test that fails when a field is added to `SearchCapabilities`, as a reminder to extend the decorators' forwarding tests, which enumerate every field.

# Scope of the change

Keep this change to what the task needs. Do not refactor existing code: do not move, rename or split existing declarations or files unless the task requires it.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./internal/storage/v2/api/tracestore
```
