# Support HTTPRoute `RequestMirror` filter

Please add support for the HTTPRoute `RequestMirror` filter, with `percent` and `fraction`, in the Kubernetes Gateway API provider. Declare the `HTTPRouteRequestMirror`, `HTTPRouteRequestMultipleMirrors` and `HTTPRouteRequestPercentageMirror` extended features as supported.

The rest of Traefik already has what this needs: the dynamic configuration has a `RequestMirror` middleware, used only by internal providers,

```go
// In the Middleware struct:
RequestMirror *RequestMirror `json:"requestMirror,omitempty" toml:"-" yaml:"-" label:"-" file:"-" kv:"-" export:"true"`

// RequestMirror holds the request mirror middleware configuration.
// Numerator out of Denominator requests are mirrored to Service.
type RequestMirror struct {
	Service     string `json:"service,omitempty"`
	Numerator   int64  `json:"numerator,omitempty"`
	Denominator int64  `json:"denominator,omitempty"`
}
```

and the server runs it. Only the provider side is missing.

## Behaviour

- A `RequestMirror` filter, in a rule's filters or in a backend's filters, becomes a middleware holding a `dynamic.RequestMirror`. Its name follows the existing naming of filter middlewares (`<parent name>-requestmirror-<filter index>`). Because the mirror is a middleware, it keeps its place in the filter order relative to the other filters, and a rule can have several mirrors.
- The mirror's `backendRef` is loaded like a rule backend, with the same rules for kinds, ports, `ReferenceGrant`s, cross-provider and native load-balancing services. The generated service is named as a backend service would be, using the mirror middleware's name as prefix and backend index 0 (for example `<middleware name>-svc-<namespace>-<service name>-0`). The `RequestMirror.Service` field holds that name. The rule's `timeouts.backendRequest` does not apply to mirror backends.
- `Numerator` and `Denominator`:
  - Neither `percent` nor `fraction` set: every request is mirrored, 1/1.
  - `percent: P`: P/100, with P clamped to the range 0 to 100.
  - `fraction`: the denominator defaults to 100. A denominator of zero or less gives 0/1 (nothing is mirrored). Otherwise the numerator is clamped to the range 0 to the denominator, so 3/2 becomes 2/2.
- As the specification says, an invalid mirror (for example a backend Service that does not exist) is dropped without failing the request: the router is generated without that middleware and no mirror service is created. Log it at debug level. Reporting this in the route conditions is not required for now.

Add test cases for a rule with several mirrors mixed with other filters, for a mirror referencing a missing Service, and for the numerator and denominator rules.

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
