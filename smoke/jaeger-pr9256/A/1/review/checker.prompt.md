This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/reader.go:153: The FilterCapabilities methods use different receiver types. IsEmpty has a pointer receiver and handles nil. SupportsLevel (line 153) and SupportsOperator (line 160) have value receivers. Callers reach these methods through SearchCapabilities.Filter, which is a *FilterCapabilities and is documented to be nil when the reader serves no filter. So caps.Filter.SupportsLevel(...) or caps.Filter.SupportsOperator(...) panics with a nil dereference in exactly the case the type documents as normal. The type doc says an empty FilterCapabilities means the same as a nil one, but only IsEmpty keeps that promise. Fix: Change SupportsLevel and SupportsOperator to pointer receivers (c *FilterCapabilities) and make them nil-safe. SupportsLevel should return `level == "" || (c != nil && slices.Contains(c.Levels, level))`. SupportsOperator should return `c != nil && slices.Contains(c.Operators, op)`. All three methods then share one receiver type, and a nil *FilterCapabilities answers like the zero value.
1. pkg/reader_test.go:58: TestFilterCapabilities_SupportsLevel and TestFilterCapabilities_SupportsOperator check only a populated value and the zero value. They never call the methods on a nil *FilterCapabilities, which is the case SearchCapabilities.Filter produces by default. TestFilterCapabilities_IsEmpty does cover nil, so the tests do not hold all three methods to the same nil-means-empty rule. Fix: After the receiver change, add a `var none *FilterCapabilities` case to both tests. Assert that none.SupportsLevel("") is true, none.SupportsLevel(expression.LevelSpan) is false, and none.SupportsOperator(expression.OpAnd) is false, so that nil and the zero value are shown to answer the same.

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
