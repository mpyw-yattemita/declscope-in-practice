This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/reader.go:136: SearchCapabilities.Filter is a *FilterCapabilities, but a nil pointer means exactly what the zero value means. The doc admits this and tells callers to use IsEmpty instead of comparing to nil. The pointer adds a second way to say "no filter support" and a nil-dereference trap. IsEmpty has a pointer receiver and is nil-safe, but SupportsLevel and SupportsOperator have value receivers. So a caller who writes caps.Filter.SupportsOperator(op) without an IsEmpty guard panics on every reader that leaves Filter unset, and that is the default for all existing readers. Fix: Change the field to `Filter FilterCapabilities` (a value). Its zero value already means "serves no filter". Give IsEmpty a value receiver like the other two methods and remove the `c == nil` branch. Remove the 'A nil Filter means none of it ... rather than comparing against nil' sentences from the field doc. In reader_test.go, change the IsEmpty table to use FilterCapabilities values and drop the "nil" case. If the pointer has to stay, at least make SupportsLevel and SupportsOperator pointer receivers that are nil-safe (nil supports only the "" level and no operator), so all three methods share one receiver kind and none of them panic.
1. pkg/reader.go:191: The doc says TraceQueryParams.Filter is mutually exclusive with ServiceName, OperationName, Attributes and the duration bounds, but nothing in the package checks this. Every Reader implementation and every decorator would have to write the same check itself, and each could get it slightly different. A Reader that forgets it ends up silently mixing the two filtering models, which is exactly what the doc says must not happen. Fix: Add one exported helper next to the type, for example `func (p TraceQueryParams) HasLegacyPredicates() bool` (true if ServiceName or OperationName is non-empty, Attributes.Len() > 0, or DurationMin or DurationMax is non-zero), or a `Validate() error` that returns an error wrapping ErrFilterInvalid when Filter != nil && HasLegacyPredicates(). Point to it from the Filter field doc so readers and the query service all use the same check. Add a table test for it in reader_test.go.

For each finding, decide whether it is real and worth fixing now. For each
kept finding, give one category:

- naming: a name does not match its role, or one idea has two names
- duplication: the same logic is written twice
- over-abstraction: a shared helper couples callers that are not related
- mixed-job: one function or type does two unrelated jobs
- dead: a definition is unused
- other: none of the above

Answer with one JSON object and nothing else:
{"verdicts": [{"index": 0, "keep": true, "category": "naming", "why": "..."}]}
