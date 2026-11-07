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

# Scope of the change

As part of this change, refactor package `pkg/provider/kubernetes/gateway` where its structure makes the change awkward. You may move, rename and split existing declarations and files within that package. Keep the package's behaviour and its exported API unchanged, and keep its tests passing.

Nobody will answer questions while you work on this task. When a decision would normally go to the repository's owner, make the choice you would recommend, carry it out, and say in your final message what you chose and why.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/provider/kubernetes/gateway
```
