# Support `ClusterTrustBundle` CA certificate references

Please add support for `ClusterTrustBundle` references in the `caCertificateRefs` of `BackendTLSPolicy` resources and of the Gateway frontend TLS validation, as defined by [GEP-5075](https://gateway-api.sigs.k8s.io/geps/gep-5075/). This fixes the request in the issue below. Declare the `ClusterTrustBundle` extended Gateway feature as supported (unconditionally for now; GEP-5075 asks to advertise it only when the API is served, which can be a later change).

> I expect to be able to reference a `ClusterTrustBundle` object in `BackendTLSPolicy.spec.validation.caCertificateRefs`:
>
> ```yaml
>   validation:
>     caCertificateRefs:
>       - name: my-ca-bundle
>         group: certificates.k8s.io
>         kind: ClusterTrustBundle
>     hostname: my-backend.example.com
> ```
>
> Traefik only resolves `caCertificateRefs` entries with `kind: ConfigMap` or `kind: Secret`. `ClusterTrustBundle` is a cluster-scoped Kubernetes API designed for distributing X.509 trust anchors across a cluster, which avoids duplicating CA bundles in every namespace.

The resource is consumed through its `certificates.k8s.io/v1` version (`k8s.io/api/certificates/v1`). Its trust anchors are the PEM bundle in `spec.trustBundle`. The YAML parser used by the provider's tests already accepts the `ClusterTrustBundle` kind.

## Detecting the API

The API is optional, and an informer that cannot list or watch the resource blocks the cache sync. So when the client starts watching (in `WatchAll`), Traefik first checks:

1. whether the cluster serves the API: the discovery resources for group version `certificates.k8s.io/v1` include `clustertrustbundles`. A "not found" discovery error means "not served"; other errors make `WatchAll` fail.
2. if it is served, whether Traefik may `list` and `watch` `clustertrustbundles` in group `certificates.k8s.io`, checked with a `SelfSubjectAccessReview` for each verb. An error creating the review makes `WatchAll` fail.

Only when both hold does Traefik watch `ClusterTrustBundle` resources (cluster-wide) and read them from the informer cache. Otherwise existing deployments keep working without the new RBAC rule, and a reference is reported as an unsupported kind.

## `BackendTLSPolicy`

For HTTPRoute, GRPCRoute and TLSRoute backends, a `caCertificateRefs` entry with group `certificates.k8s.io` and kind `ClusterTrustBundle` is now accepted, besides core `ConfigMap` and `Secret`. The policy's `ResolvedRefs` condition is `False` with:

| Case | Reason | Message |
| --- | --- | --- |
| Any other group or kind | `InvalidKind` | `Only ConfigMaps, Secrets, and ClusterTrustBundles are supported` (this replaces the former `Only ConfigMaps and Secrets are supported`) |
| API not served | `InvalidKind` | `ClusterTrustBundles are not supported, the cluster does not serve the certificates.k8s.io/v1 API` |
| Not allowed to list and watch | `InvalidKind` | `ClusterTrustBundles are not supported, Traefik is not allowed to list and watch clustertrustbundles` |
| Bundle not found | `InvalidCACertificateRef` | `Cannot get ClusterTrustBundle <name>: <error>` |
| Empty `spec.trustBundle` | `InvalidCACertificateRef` | `ClusterTrustBundle <name> has an empty trustBundle` |

Otherwise the bundle's PEM is used as the root CA of the generated ServersTransport, exactly like the `ca.crt` of a ConfigMap or Secret.

## Gateway frontend TLS validation

A frontend validation `caCertificateRefs` reference (still a single reference) can now be a `ClusterTrustBundle` too. It needs no `ReferenceGrant`. The listener's `ResolvedRefs` condition is `False` with:

| Case | Reason | Message |
| --- | --- | --- |
| API not served | `InvalidCACertificateKind` | `ClusterTrustBundles are not supported, the cluster does not serve the certificates.k8s.io/v1 API` |
| Not allowed to list and watch | `InvalidCACertificateKind` | `ClusterTrustBundles are not supported, Traefik is not allowed to list and watch clustertrustbundles` |
| A namespace is set on the reference | `InvalidCACertificateRef` | `Unsupported namespace on cluster-scoped CACertificateRef: ClusterTrustBundle <name>` |
| Bundle not found | `InvalidCACertificateRef` | `Cannot resolve ClusterTrustBundle: <name>: <error>` |
| Empty `spec.trustBundle` | `InvalidCACertificateRef` | `Cannot find trustBundle: ClusterTrustBundle <name>` |

Otherwise the bundle's PEM is the client CA of the listener's TLS options, exactly like a ConfigMap's `ca.crt`.

Add test cases for HTTPRoute, GRPCRoute and TLSRoute backends with a `ClusterTrustBundle`, with the API not served and with the permissions missing, and for each frontend validation case.
