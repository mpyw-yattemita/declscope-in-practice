This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/filter.go:49: FilterCapabilities mixes receiver types. IsEmpty has a nil-safe pointer receiver, but SupportsLevel and SupportsOperator have value receivers. SearchCapabilities.Filter is a *FilterCapabilities and nil is documented as a valid value, so caps.Filter.SupportsLevel(...) or caps.Filter.SupportsOperator(...) panics when a Reader declares no filter support. Only IsEmpty is safe to call on the nil value. Fix: Give all three methods a *FilterCapabilities receiver and make them nil-safe. For a nil c, SupportsLevel returns level == "" and SupportsOperator returns false. Add nil-pointer cases to TestFilterCapabilities_SupportsLevel and TestFilterCapabilities_SupportsOperator. Alternatively, apply the value-type change in the next finding, then switch IsEmpty to a value receiver.
1. pkg/capabilities.go:30: SearchCapabilities.Filter is a pointer, but nil and &FilterCapabilities{} mean the same thing. The docs on both types have to tell callers not to compare against nil. Two encodings of 'no filter support' invite bugs like the nil-receiver panic above, and they make struct equality in decorator forwarding tests depend on which encoding was used. Fix: Change the field to `Filter FilterCapabilities` (a value, not a pointer). The zero value then is the least-capable declaration, which matches the type's documented zero-value design. Change IsEmpty to a value receiver, drop the 'rather than comparing against nil' sentences from both doc comments, and use caps.Filter.IsEmpty() on the value in TestSearchCapabilities_ZeroValueIsLeastCapable.
2. pkg/query.go:30: The ServiceName and Filter contracts contradict each other. ServiceName 'may be empty only when the Reader's SearchCapabilities reports WithoutServiceName'. Filter is 'mutually exclusive with ... ServiceName', so every Filter query has an empty ServiceName. As written, a Reader that declares filter support but not WithoutServiceName can never legally receive a Filter query. The contract does not say whether the service is constrained inside the Filter or whether filter support implies WithoutServiceName. Fix: Reword the ServiceName comment: 'When Filter is nil, ServiceName may be empty only when SearchCapabilities.WithoutServiceName is true; when Filter is set, ServiceName is empty and the service, if any, is constrained in Filter.' In the SearchCapabilities.WithoutServiceName doc (capabilities.go:15), state that the flag applies only to the legacy predicate fields.
3. pkg/capabilities_test.go:19: TestSearchCapabilities_FieldCount uses reflection to count fields and fails on any addition. Its own comment says it is a reminder, not a contract. It also names tests in other packages (tracestoremetrics, queryinterceptor) by function name, so a rename there makes this comment stale silently. It covers for decorator forwarding tests that check fields one by one when they should compare whole values. Fix: Delete TestSearchCapabilities_FieldCount and its reflect import. In the two decorator tests, have the backend mock return a SearchCapabilities with every field set to a non-zero value, then check the forwarded result with assert.Equal(t, want, got) against the whole struct. Fields added later are then covered without a count to keep in sync.

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
