This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/search_capabilities.go:28: `Filter` is a `*FilterCapabilities`, so 'supports no filter' can be written two ways: nil, or a pointer to an empty value. The docs then have to warn callers not to compare it against nil, and `IsEmpty` needs a nil-safe pointer receiver to cover both. The pointer adds nothing, because the zero `FilterCapabilities` already means 'none', just like the zero values of the other fields in `SearchCapabilities`. Fix: Make the field a value: `Filter FilterCapabilities`. Remove the 'A nil Filter means…' and 'rather than comparing it against nil' sentences. Update `TestSearchCapabilitiesZeroValue` to call `zero.Filter.IsEmpty()` on the value.
1. pkg/filter.go:43: The receivers are mixed. `IsEmpty` uses a pointer receiver and accepts nil, but `SupportsLevel` and `SupportsOperator` use value receivers. Since `SearchCapabilities.Filter` is a pointer, `caps.Filter.SupportsOperator(op)` compiles and then panics with a nil dereference when the reader sets no Filter, which is exactly the zero-value case the type was designed for. Fix: Once `Filter` is a value (finding above), change `IsEmpty` to a value receiver: `func (c FilterCapabilities) IsEmpty() bool { return len(c.Levels) == 0 && len(c.Operators) == 0 }`. Drop the 'nil c' sentences from its doc and the `nil` case from `TestFilterCapabilitiesIsEmpty`. If the pointer has to stay, switch all three methods to nil-safe pointer receivers instead.
2. pkg/reader.go:50: The package adds the `ErrFilterUnsupported` and `ErrFilterInvalid` sentinels, but the `Reader` contract never says that `FindTraces`, `FindTraceIDs` or `FindTraceSummaries` return them. It also doesn't say they are wrapped with `%w`, which callers need to know to use `errors.Is`. Nor does it say what a reader with an empty `SearchCapabilities.Filter` must do when it receives a non-nil `TraceQueryParams.Filter`. Implementers have to guess the error behavior. Fix: In the `FindTraces` doc comment, add a paragraph saying: when `query.Filter` uses a level, operator or structure outside `SearchCapabilities.Filter`, the iterator yields an error wrapping `ErrFilterUnsupported` as its first error, before any batch; when a value doesn't fit its field, it yields an error wrapping `ErrFilterInvalid`. The existing 'behaves identically to FindTraces' wording already carries this to `FindTraceIDs`. Add an explicit sentence to the `FindTraceSummaries` doc, because its fallback contract relies on `errors.ErrUnsupported` and the two must not be confused.
3. pkg/reader.go:116: The `Filter` doc says it is mutually exclusive with `ServiceName`, `OperationName`, `Attributes` and the duration bounds. It doesn't say who enforces that rule or which error a mixed query gets, so every reader will either check it with its own code or silently ignore one side. Fix: State the outcome in the field doc: a query that sets `Filter` together with any predicate field is rejected with an error wrapping `ErrFilterInvalid`. Add one helper next to the sentinels in `filter.go`, such as `func (p TraceQueryParams) validateFilterExclusive() error` or an exported equivalent, that the readers and the query service call, so the check exists in one place.

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
