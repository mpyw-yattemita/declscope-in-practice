# Support Gateway backend client certificate

Please add support for the `Gateway` field `spec.tls.backend.clientCertificateRef` ([GatewayBackendTLS](https://gateway-api.sigs.k8s.io/reference/api-spec/main/spec/#gatewaybackendtls), [GEP-3155](https://gateway-api.sigs.k8s.io/geps/gep-3155/)) in the Kubernetes Gateway API provider. This is the `GatewayBackendClientCertificate` feature (Standard channel, Extended support), which completes backend mTLS with `BackendTLSPolicy`. Declare it in the extended Gateway features that the provider reports as supported.

## Behaviour

- When a `Gateway` references a client certificate, Traefik presents it to the backends targeted by a `BackendTLSPolicy`: for `HTTPRoute` and `GRPCRoute` backends it is set as `Certificates` of the generated `dynamic.ServersTransport`, and for `TLSRoute` backends (`Terminate` mode) as `TLS.Certificates` of the generated `dynamic.TCPServersTransport`. The certificate and key come from the referenced Secret, as for listener certificates.
- Backends without a `BackendTLSPolicy` are not affected: the field does not enable TLS by itself.
- The setting is defined at the Gateway level, so listeners declared by a `ListenerSet` attached to the Gateway use their parent Gateway's client certificate. Two Gateways with different client certificates each present their own.
- The reference is resolved as follows:
  - Only a core `Secret` is supported (an empty or `core` group, kind `Secret` by default). Any other group or kind is reported with reason `InvalidClientCertificateRef` and message `Unsupported ClientCertificateRef group/kind: <group>/<kind>`.
  - The namespace defaults to the Gateway's namespace. A Secret in another namespace needs a `ReferenceGrant` from Gateways in the Gateway's namespace; otherwise the reason is `RefNotPermitted` and the message `Cannot reference ClientCertificateRef <namespace>/<name>: <error>`.
  - A missing or malformed Secret is reported with reason `InvalidClientCertificateRef` and message `Cannot load ClientCertificateRef <namespace>/<name>: <error>`.
  - When the reference cannot be resolved, no client certificate is used.

## Gateway `ResolvedRefs` condition

The Gateway status now always carries a `ResolvedRefs` condition (with the Gateway's generation as `ObservedGeneration`):

- `False` with the reasons and messages above when the Gateway's own client certificate reference cannot be resolved. This takes precedence over the listeners.
- Otherwise `False` with reason `ListenersNotResolved` and message `At least one Listener has unresolved references` when at least one of the Gateway's own listeners (not those of a ListenerSet) has a `ResolvedRefs` condition set to `False`.
- Otherwise `True` with reason `ResolvedRefs` and message `No error found`.

This condition does not change the `Accepted` and `Programmed` conditions. When updating the Gateway status, `ResolvedRefs` is a condition type owned by Traefik, like `Accepted` and `Programmed`: a stale one is replaced or dropped, and conditions owned by others are kept.

Add test cases for HTTPRoute, GRPCRoute and TLSRoute backends with a client certificate, for a ListenerSet, for two Gateways, for each way the reference can fail, and for the Gateway status.

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
go test ./pkg/provider/kubernetes/gateway
```
