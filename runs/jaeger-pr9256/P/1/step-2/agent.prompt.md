# storage: let a trace reader declare which query filters it evaluates

The structured query filter of RFC 0005 is coming to trace search: a boolean expression over level-qualified attributes and built-in fields, from the `github.com/jaegertracing/jaeger-idl/query/expression/v1` package (imported as `expression`; the repository already depends on a jaeger-idl version that has it). The storage interface needs a place to carry a filter and, more importantly, a way for each backend to say how much of one it can evaluate. Backends differ sharply here: ClickHouse and Elasticsearch/OpenSearch can serve a nested boolean expression, while the flat index backends can serve a conjunction and nothing else. Without a declaration, a caller can neither avoid a query that cannot be answered nor interpret a surprising result.

Search-without-service-name already solved this problem with `SearchCapabilities`. Please extend that mechanism instead of inventing a second one.

## Public API

```go
// ErrFilterUnsupported is returned for a well-formed query filter that the
// storage cannot serve: a level it does not index, an operator it has not
// implemented, or a boolean structure a flat index cannot evaluate. The query
// is refused rather than approximated.
var ErrFilterUnsupported = errors.New("this storage backend cannot serve this query filter")

// ErrFilterInvalid is returned for a query filter whose value does not fit the
// field it compares.
var ErrFilterInvalid = errors.New("invalid query filter")

type SearchCapabilities struct {
	WithoutServiceName bool // existing

	// SameSpanConjunction is true when a conjunction (several entries in the
	// Attributes map, or an OpAnd in Filter) is satisfied within a single span.
	// False means the backend may satisfy different conjuncts from different
	// spans of the same trace, as a flat inverted index does. Callers report
	// this looser scoping rather than refusing the query.
	SameSpanConjunction bool

	// Filter is how much of TraceQueryParams.Filter the reader evaluates. A nil
	// Filter means none of it: the reader serves only the legacy fields.
	Filter *FilterCapabilities
}

// FilterCapabilities declares how much of a structured filter a Reader
// evaluates, by naming the levels and operators it serves.
type FilterCapabilities struct {
	// Levels are the levels a reference may name. Empty means the Reader can
	// serve only unqualified references.
	Levels []expression.Level
	// Operators are the operators the Reader evaluates, including the boolean
	// combinators OpAnd, OpOr and OpNot.
	Operators []expression.Operator
}

func (c *FilterCapabilities) IsEmpty() bool
func (c FilterCapabilities) SupportsLevel(level expression.Level) bool
func (c FilterCapabilities) SupportsOperator(op expression.Operator) bool

type TraceQueryParams struct {
	// ... existing fields ...

	// Filter is the structured query filter. It is mutually exclusive with the
	// predicate fields ServiceName, OperationName, Attributes and the duration
	// bounds, so a reader sees one filtering model, not a mix of the two.
	Filter *expression.Call
}
```

`SearchCapabilities` now has exactly these three fields. The field-count test from the previous change must be updated to match.

## Behaviour

- Nothing is implicit. A reader opts in only by naming something, so an empty `FilterCapabilities` says the same as no `FilterCapabilities`. `IsEmpty` returns true for a nil pointer, for a zero value, and for empty (non-nil) slices, and false as soon as one level or one operator is named. Callers ask `IsEmpty` rather than comparing against nil. This keeps the zero value of `SearchCapabilities` the least capable reader.
- `SupportsLevel` returns true for the empty level (an unqualified reference always reaches the reader), and otherwise only for a level listed in `Levels`.
- `SupportsOperator` returns true only for an operator listed in `Operators`. The boolean combinators are declared like any other operator: a reader confined to conjunctive search lists `OpAnd` and omits `OpOr` and `OpNot`. The zero value supports no operator, not even `OpAnd`. Nesting is not declared, because `and` is associative and a caller flattens it before asking.
- No reader populates the declaration yet and nothing reads it; the query service starts using it in a later change.

The rest of the repository needs no change. Do not change code outside the package. Add unit tests for the three methods.

# Where code goes in this package

This package keeps each file's internals to itself. Follow these rules when you add or change code.

Existing code may not follow them yet. Do not fix existing code beyond what the task needs. Follow the rules in the code you add or change.

## The rule

| Declaration | Who may use it |
| --- | --- |
| Unexported | Only the file it is declared in |
| Exported | Any file of the package |

The same holds for members:

- An unexported field may be read or written only in the file that declares its struct.
- A method belongs to the file it is written in, not to its type's file. A method in another file must not touch the type's unexported fields.
- A test file goes with the source file of the same name. `parse_test.go` may use the unexported code of `parse.go`, and of no other file. So name a test file for the source file whose unexported code it calls, not for the behaviour it tests.

When code in one file needs an unexported declaration of another file, that is a crossing. Avoid creating new ones.

## Names carry their file's name

A file's name, without `.go` and in camelCase, is a word that the names declared in it carry. `user_repository.go` gives `userRepository`.

| Declaration | Its name |
| --- | --- |
| A package-level function, type, variable or constant, exported or not | Contains the file's word, beginning where a word in the name begins: `newUserRepository`, `UserRepositoryOption`, `findUserRepositoryRows` |
| A method written in its type's file, a field, an interface method | Free |
| A method written in another file | Contains that other file's word |
| A helper in a test file | Contains the word of the source file the test file goes with |
| `TestXxx`, `BenchmarkXxx`, `FuzzXxx`, `ExampleXxx` | Free |

This applies once the package has more than one file with declarations.

- Fit the word where it reads naturally: as a modifier (`userProfile`), as an object (`loadUser`), or later in the phrase (`addToCollection`). Read the name aloud.
- Name the thing accurately first, then make sure the word is in it. Do not bend a name only to carry the word.
- When the word reads badly in a name, the declaration is usually in the wrong file. Move it rather than force the name.
- When a name is fixed, for example by the task, put the declaration in a file whose word that name carries.
- Do not rename existing declarations only to follow this.

## Where a new declaration goes

| The declaration is used by | Put it |
| --- | --- |
| Code in one file | That file |
| Several files, as a helper they use in common on purpose | A file named for its concept, not the file of its first caller. Say in its doc comment which files use it and why |
| Most files of the package, as its common model | The file that already holds the package's common types, if there is one |
| Another package | Wherever it belongs, exported |

**Before you reach into another file, ask whether your code is in the wrong file.** Often the function you are writing belongs beside what it uses.

**Do not create a file just to hold one small declaration** that only its old file uses.

## Common changes

| Change | Do this |
| --- | --- |
| Add a method to a type | Put it in the type's file |
| Add a field that another file reads | Keep it only if the other file really needs it. Say in the field's comment which file reads it |
| Move code between files | Move its tests with it |
| Split a file | Split where the code really divides. Calls between the halves become crossings |

## Do not hide a crossing

| Tempting | Why not |
| --- | --- |
| Exporting a name so another file may use it | It hides the crossing and widens the API. In an `internal/` package, do not export something no other package uses |
| Moving everything into one file | It hides the structure instead of fixing it |
| Widening production code for a test | Move the test instead |

## Order declarations by reach

Inside a file, put what other packages use first, then what other files use, then the file's own internals. Keep a type with its constructors and methods. Do not reorder declarations your change does not touch, and do not reorder where the order is observable, such as `iota` or unkeyed literals.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./internal/storage/v2/api/tracestore
```
