# query: serve a filter query in the shape the backend understands

Readers can now declare, through `SearchCapabilities.Filter`, how much of a structured filter they evaluate. Someone has to act on that declaration: send the filter where it can be served, express it in the older predicate fields where it cannot, and refuse what neither can answer. That decision belongs in one place, so that every caller gets the same answer instead of each backend's own. Please add it to the `tracestore` package, as methods beside the query type and the capability declaration.

## Public API

```go
// EnsureFilterStandsAlone rejects a query that carries both a filter and one
// of the predicate fields the filter replaces.
func (q TraceQueryParams) EnsureFilterStandsAlone() error

// ForCapabilities gives the Reader whichever of the two filtering models it
// declared it can evaluate.
func (q TraceQueryParams) ForCapabilities(caps SearchCapabilities) (TraceQueryParams, error)

// EnsureSupported walks the filter and refuses the first predicate the Reader
// did not declare it can evaluate.
func (c FilterCapabilities) EnsureSupported(filter *expression.Call) error

// ToFilterShape expresses q's scalar predicate fields as filter predicates.
func (q TraceQueryParams) ToFilterShape() TraceQueryParams

// ToLegacyShape expresses q's Filter in the scalar predicate fields every
// backend serves.
func (q TraceQueryParams) ToLegacyShape() (TraceQueryParams, error)
```

Every refusal wraps `ErrFilterUnsupported` or `ErrFilterInvalid` (test with `errors.Is`) and its message includes the texts given below.

## `EnsureFilterStandsAlone`

- A query without a filter, or with a filter and none of the legacy predicate fields, passes. An empty (non-nil) `Attributes` map is not a predicate.
- Otherwise it returns an error wrapping `ErrFilterInvalid` whose message names every conflicting field in this order, in Go's `%v` form of a string slice: `[service_name operation_name duration_min duration_max attributes]`. For example, a filter beside a service name gives a message containing `[service_name]`.

## `ForCapabilities`

- A query without a filter is returned unchanged.
- If `caps.Filter` declares nothing (`IsEmpty`), the result is `q.ToLegacyShape()`.
- Otherwise the query is returned as it is, with the error of `caps.Filter.EnsureSupported(q.Filter)`.

## `EnsureSupported`

- A nil filter is supported.
- It checks the call's operator, then each argument: a nested call is checked recursively (also on the right side of a comparison), and an attribute, field or nested reference is checked for its level. Constants need no support.
- Refusals wrap `ErrFilterUnsupported`, with the message `it does not support the operator "<op>"` or `it does not index the "<level>" level`.

## `ToFilterShape`

- It cannot fail: each legacy field has an exact filter equivalent. A query that already has a `Filter` is returned as it is.
- `ServiceName` becomes `eq(resource field "service", string)`. `OperationName` becomes `eq(span field "name", string)`. Each entry of `Attributes` becomes `eq(unqualified attribute <key>, AnyValue(value as string))`, in map order. `DurationMin` becomes `gte(span field "duration", DurationValue)` and `DurationMax` becomes `lte(...)`. Use the field-name constants of the `expression` package. The predicates come in this order: service, operation, attributes, min, max.
- The converted fields are cleared (`Attributes` becomes a new empty map); everything else, such as the time range and `SearchDepth`, is kept.
- A conjunction appears only when there is something to conjoin: no predicates give a nil `Filter`, one predicate stands on its own, and two or more are the arguments of one `OpAnd` call. The result needs no further resolution: it passes `expression.ValidateFilter`, and `expression.ResolveConstants` returns it unchanged.
- A query whose `Attributes` was left at its zero value (not created with `pcommon.NewMap()`) must convert without panicking.

## `ToLegacyShape`

- A query without a filter is returned unchanged. Otherwise the result has a nil `Filter` and a new `Attributes` map, so the caller's own map is not changed. On error the original query is returned.
- `ToFilterShape` followed by `ToLegacyShape` gives back the original query.
- Only a flat conjunction converts. A nested `and` is flattened. These forms convert: `eq` on an unqualified attribute, or on a span-, resource- or event-level attribute (which widens to a plain tag), against a string or untyped constant; `eq` on the resource field `service` or the span field `name` against a string or untyped constant; and `gte` / `lte` on the span field `duration` against a duration constant, which become the inclusive bounds.
- Everything else is refused with `ErrFilterUnsupported`. The message contains:
  - `does not support the operator "or"` (or `"not"`, `"exists"`) for those operators, also when nested in a conjunction, and for any predicate that does not have exactly two arguments;
  - `conjunction of predicates only` when a conjunct is not a predicate, for example a bare reference;
  - `does not support the operator "<op>" on "<name>"` for a comparison other than the ones above, for example `"ne" on "http.method"`, `"gt" on "duration"`, `"regex" on "name"`, `"ne" on "service"`, `"in" on "http.status_code"`;
  - `does not index the "<level>" level` for an attribute at the link or scope level;
  - `does not support the built-in field "<name>" of the "<level>" level` for any other built-in field, for example `"kind"` of `"span"`;
  - `compares "<name>" against a string constant only` when an attribute, the service or the operation name is compared against a constant of another type (a list, an integer, a boolean) or against another reference;
  - `compares "duration" against a duration such as "2s" only` when a duration bound is not a duration constant;
  - `compares a reference against a constant only` when the left side is not an attribute or field reference, for example a quantifier over the events;
  - `only one predicate on "<name>"` when two predicates write the same attribute, the service, the operation name, or the same duration bound.
- One case is refused with `ErrFilterInvalid` instead: a duration bound whose constant is still untyped, which means it was never resolved. Its message contains `the bound "2s" on "duration" was never read as a duration` (with the untyped constant's text).

The query service already dispatches through this API, behind the alpha feature gate `jaeger.query.structuredFilters`: it calls `EnsureFilterStandsAlone`, then `ForCapabilities` with the reader's `SearchCapabilities`, and maps both errors to a bad request. That code is in place and compiles once the package has these methods. Do not change code outside the package.

Add unit tests for these rules.
