# Support HTTPRoute retry

Please add support for the HTTPRoute rule `retry` field in the Kubernetes Gateway API provider. A rule with `retry` gets a Retry middleware. Declare the `HTTPRouteRetryCodes` and `HTTPRouteRetryConnectionError` extended features as supported.

## New provider option

Add a field to the provider configuration:

```go
RetryMaxRequestBodyBytes int64 `description:"Defines the maximum request body size (in bytes) buffered for HTTPRoute retries. If negative, the limit is disabled." json:"retryMaxRequestBodyBytes,omitempty" toml:"retryMaxRequestBodyBytes,omitempty" yaml:"retryMaxRequestBodyBytes,omitempty" export:"true"`
```

The option is `providers.kubernetesGateway.retryMaxRequestBodyBytes`. Its default, set in the provider's `SetDefaults`, is `dynamic.RetryDefaultMaxRequestBodyBytes` (2097152). A request with a larger body is forwarded without retry.

## Behaviour

- For a rule with `retry`, each router generated for the rule's matches gets a middleware named `<router name>-retry` holding a `dynamic.Retry`:
  - `Attempts` is the Gateway API `attempts` plus one, because Traefik also counts the initial request. When `attempts` is not set, it defaults to 1, so `Attempts` is 2.
  - `Status` lists the `codes`, each as a decimal string.
  - `InitialInterval` is the `backoff` duration when set. The CRD validation guarantees a valid duration string; if parsing still fails, log an error and ignore the backoff.
  - `MaxRequestBodyBytes` points to the provider's `RetryMaxRequestBodyBytes` value.
- The retry middleware comes last in the router's middleware list, after the rule's filter middlewares, so that only the call to the backend is retried.
- When the rule's filters fail to load (the router then answers with an error), no retry middleware is added.
- An empty `retry: {}` also gets a retry middleware with the defaults above.

Add test cases for a rule with retry settings, for an empty retry, and for a custom `RetryMaxRequestBodyBytes`.

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

# Checking the rule above

This repository runs declscope, a linter that checks the rule above. The `declscope-authoring` skill explains how to act on its reports. Read it before you start.

`.declscope.yaml` sets the rules, and `.declscope-baseline.yaml` records the reports the existing code already has, so only what you add is reported. Never edit either file, and never regenerate the baseline.

Before you finish, run these as well, and resolve what they report:

```bash
declscope ./...
declscope shrink ./...
```

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/provider/kubernetes/gateway
declscope ./...
declscope shrink ./...
```
