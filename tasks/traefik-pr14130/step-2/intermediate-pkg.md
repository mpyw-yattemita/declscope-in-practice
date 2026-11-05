## Also: follow-up of the Gateway API v1.7.0-rc.3 upgrade

The repository has just been upgraded to `sigs.k8s.io/gateway-api` v1.7.0-rc.3. Two small provider changes come with it:

- Add `HTTPRouteBackendURLRewrite` to the extended HTTPRoute features, and `GRPCRouteRequestHeaderModifier` to the extended GRPCRoute features, that the provider reports as supported.
- For a TLSRoute backend that uses a `TCPServersTransport` with a TLS configuration (from a `BackendTLSPolicy`), mark each generated `dynamic.TCPServer` with `TLS: true`, unless the listener is in `Passthrough` mode. The TLS settings of the transport are only used to dial servers marked as TLS.
