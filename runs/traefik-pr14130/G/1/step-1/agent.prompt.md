# Support HTTPRoute `backendRequest` timeout

The Kubernetes Gateway API provider already supports the HTTPRoute rule `timeouts.request` field. Please add support for `timeouts.backendRequest` ([GEP-1742](https://gateway-api.sigs.k8s.io/geps/gep-1742/)), so that the provider can declare the `HTTPRouteBackendTimeout` extended feature as supported.

## Behaviour

- When an HTTPRoute rule sets `timeouts.backendRequest`, every Kubernetes `Service` backend of that rule gets a `dynamic.ServersTransport` whose `ForwardingTimeouts` use the usual defaults (`ForwardingTimeouts.SetDefaults()`: 30s dial, 90s idle, 15s ping) and set `ResponseHeaderTimeout` to the `backendRequest` duration. The timeout applies until the backend's response headers are received; it does not bound reading the response body.
- The ServersTransport is registered under the same name as the backend's service, and the service's load balancer refers to it through `ServersTransport`, as is already done for the ServersTransport created for a `BackendTLSPolicy`. When a backend already gets a ServersTransport (for example from a `BackendTLSPolicy`), the forwarding timeouts are set on that same transport.
- A zero duration means no timeout: no ServersTransport is created just for it.
- The CRD validation guarantees a valid duration string. If parsing still fails, log an error and ignore the field.
- Rules without `timeouts.backendRequest` produce the same configuration as before.
- Add `HTTPRouteBackendTimeout` to the extended HTTPRoute features that the provider reports as supported.

Add a test case for an HTTPRoute that sets both `request` and `backendRequest` timeouts.

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
