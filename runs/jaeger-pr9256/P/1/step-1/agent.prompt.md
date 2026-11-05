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
