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

# Scope of the change

As part of this change, refactor package `pkg/provider/kubernetes/gateway` where its structure makes the change awkward. You may move, rename and split existing declarations and files within that package. Keep the package's behaviour and its exported API unchanged, and keep its tests passing.

Nobody will answer questions while you work on this task. When a decision would normally go to the repository's owner, make the choice you would recommend, carry it out, and say in your final message what you chose and why.

Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.

```bash
go vet ./...
go test ./pkg/provider/kubernetes/gateway
```
