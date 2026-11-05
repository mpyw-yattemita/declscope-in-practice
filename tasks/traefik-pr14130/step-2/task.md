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
