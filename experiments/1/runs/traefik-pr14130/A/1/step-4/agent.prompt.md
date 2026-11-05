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

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/provider/kubernetes/gateway
```
