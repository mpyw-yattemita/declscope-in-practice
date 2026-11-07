This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/reader.go:125: SearchCapabilities.Filter is a *FilterCapabilities, so "no filter support" has two forms: nil and an empty struct. The type doc has to tell callers to call IsEmpty instead of checking for nil, and IsEmpty needs a nil-pointer branch just to cover the second form. The pointer gives no extra information, because the doc defines an empty value to mean the same as nil. Fix: Change the field to the value type `Filter FilterCapabilities`. Its zero value then means "serves only the legacy fields". Change IsEmpty to a value receiver, drop the nil check, and remove the sentence telling callers not to compare against nil. In reader_test.go, replace the "nil pointer" case with the zero-value case.
1. pkg/reader.go:153: The receivers do not match: IsEmpty uses a pointer receiver and handles nil, but SupportsLevel and SupportsOperator (line 161) use value receivers. Callers always reach these through the pointer field SearchCapabilities.Filter. So `caps.Filter.SupportsOperator(op)` panics when Filter is nil, even though `caps.Filter.IsEmpty()` on the same value is documented as safe. Fix: Use the same receiver kind for all three methods. If Filter becomes a value (previous finding), make all three value receivers. If the pointer stays, give SupportsLevel and SupportsOperator pointer receivers that return `level == ""` and `false` for a nil c, and add nil-pointer cases to their tests.
2. pkg/reader.go:60: The FindTraces doc still says it is "implementation-dependent" whether all query filters must match the same span. SearchCapabilities.SameSpanConjunction now reports exactly this per reader, so the comment is stale and does not point readers to the capability. Fix: Replace that paragraph with: "Whether a conjunction of filters must be satisfied by a single span is reported by SearchCapabilities.SameSpanConjunction." Add the same reference to FindTraceIDs and FindTraceSummaries, or say once that they follow the same rule.
3. pkg/reader.go:192: The doc says Filter cannot be combined with ServiceName, OperationName, Attributes and the duration bounds, but the package gives no way to check this. Each backend and decorator has to write its own check, and nothing says which error a mixed query returns. The two new errors cover unsupported filters and badly typed values, not a mixed query. Fix: Add a method `func (p TraceQueryParams) ValidateFilterExclusivity() error` (or `Validate`) that returns an error wrapping ErrFilterInvalid when Filter is non-nil and any predicate field is set (`Attributes.Len() > 0` for Attributes). Mention it in the Filter field doc and the ErrFilterInvalid doc, and add a table test for it in reader_test.go.

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
