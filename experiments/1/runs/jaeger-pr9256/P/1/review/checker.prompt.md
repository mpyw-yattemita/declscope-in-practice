This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

0. pkg/search_capabilities.go:28: Filter is a *FilterCapabilities, so the type has two ways to say 'no filter support' (nil and empty). The comment has to tell callers to use IsEmpty instead of comparing against nil. The pointer also gains nothing, because the zero value of FilterCapabilities already means 'none', which matches the struct's own rule that the zero value is least capable. Fix: Change the field to `Filter FilterCapabilities` (a value, not a pointer). Remove the sentences about nil from the field comment and from the FilterCapabilities type comment.
1. pkg/filter.go:42: The receivers are mixed: IsEmpty has a pointer receiver and handles nil, but SupportsLevel (line 48) and SupportsOperator (line 54) have value receivers. While SearchCapabilities.Filter is a pointer, `caps.Filter.SupportsOperator(op)` panics when a reader leaves Filter nil, which is the zero value the package recommends. The nil check in IsEmpty exists only because of that pointer. Fix: After making Filter a value, give IsEmpty a value receiver: `func (c FilterCapabilities) IsEmpty() bool { return len(c.Levels) == 0 && len(c.Operators) == 0 }`. Remove the 'nil pointer' case from the comment and from TestFilterCapabilities_IsEmpty in filter_test.go. If the pointer has to stay, use pointer receivers on all three methods and return false for nil in SupportsLevel and SupportsOperator.
2. pkg/reader.go:48: The FindTraces comment still says it is implementation-dependent whether all query filters must match the same span. That is out of date: SearchCapabilities.SameSpanConjunction now states this for each reader. The comment also covers only FindTraces, while the capability covers FindTraceIDs and FindTraceSummaries too. Fix: Replace the 'There's currently an implementation-dependent ambiguity…' paragraph with: 'Whether a conjunction must be satisfied within a single span is reported by SearchCapabilities().SameSpanConjunction.'
3. pkg/reader.go:89: Adding SearchCapabilities to the Reader interface forces every backend and test fake to write the same 'I can't tell' method. FindTraceSummaries handles this case with an embeddable helper (UnsupportedTraceSummaries in summary.go), but nothing like it exists for this method, so every implementation will repeat the same fmt.Errorf wrapping of errors.ErrUnsupported. Fix: Add `type UnsupportedSearchCapabilities struct{}` with `func (UnsupportedSearchCapabilities) SearchCapabilities(context.Context) (SearchCapabilities, error)`. It should return the zero value and an error that wraps errors.ErrUnsupported with %w. Put it in search_capabilities.go next to SearchCapabilities, test it the same way summary_test.go tests UnsupportedTraceSummaries, and mention it in the method comment.

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
