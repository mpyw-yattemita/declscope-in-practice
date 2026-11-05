# mcp: support URL mode elicitation (SEP-1036)

Add initial support for URL mode elicitation from [SEP-1036](https://github.com/modelcontextprotocol/modelcontextprotocol/issues/1036): a server can ask the client to send the user to a URL for a secure out-of-band interaction, instead of showing a form. This change adds the new types and both server-side and client-side validation. Support for the "URL elicitation required" error will come in a follow-up.

The repository's example server and the documentation have already been updated for this change. The example server uses `ElicitParams.URL` and `ElicitParams.ElicitationID`, so it builds only once these fields exist.

## Public API

`ElicitParams` gains three fields (JSON names in parentheses):

```go
// The mode of elicitation to use: "form" or "url".
// If unset, it is inferred from the other fields.
Mode string            // "mode"
// RequestedSchema is now only used for "form" elicitation,
// and is omitted from JSON when empty.
RequestedSchema any    // "requestedSchema,omitempty"
// The URL to present to the user. Only used for "url" elicitation.
URL string             // "url,omitempty"
// The ID of the elicitation. Only used for "url" elicitation.
ElicitationID string   // "elicitationId,omitempty"
```

Capabilities:

```go
// ElicitationCapabilities describes the capabilities for elicitation.
// If neither Form nor URL is set, the form capability is assumed.
type ElicitationCapabilities struct {
	Form *FormElicitationCapabilities
	URL  *URLElicitationCapabilities
}

type FormElicitationCapabilities struct{}
type URLElicitationCapabilities struct{}
```

The completion notification, which the server sends to tell the client that an out-of-band elicitation has finished. Its method is `notifications/elicitation/complete`.

```go
type ElicitationCompleteParams struct {
	Meta `json:"_meta,omitempty"`
	// The ID of the elicitation that has completed. It must match the
	// elicitationId of the original elicitation/create request.
	ElicitationID string `json:"elicitationId"`
}

type ElicitationCompleteNotificationRequest = ClientRequest[*ElicitationCompleteParams]
```

`ClientOptions` gains:

```go
// ElicitationModes lists the elicitation modes the client supports.
// If ElicitationHandler is set and ElicitationModes is empty, it defaults to ["form"].
ElicitationModes []string
// ElicitationCompleteHandler handles notifications/elicitation/complete.
ElicitationCompleteHandler func(context.Context, *ElicitationCompleteNotificationRequest)
```

## Behaviour

Client capabilities: a client with an `ElicitationHandler` advertises `Form` if `ElicitationModes` is empty or contains `"form"`, and `URL` if it contains `"url"`. A client without a handler advertises no elicitation capability, as before.

Client, when it receives `elicitation/create`. Every error here is a JSON-RPC error with the invalid-params code (-32602):

- No `ElicitationHandler`: the error message is `client does not support elicitation`. Before, this used a different code.
- An empty mode means `"form"`.
- `"form"`: if `URL` is set, fail with `URL must not be set for form elicitation`. Otherwise validate the schema, call the handler, and validate the result and apply defaults, exactly as before.
- `"url"`: if `RequestedSchema` is set, fail with `requestedSchema must not be set for URL elicitation`. If `URL` is empty, fail with `URL must be set for URL elicitation`. Otherwise pass the request to the handler and return its result unchanged, with no schema validation.
- Any other mode fails with `unsupported elicitation mode: "<mode>"`.

Client, when it receives `notifications/elicitation/complete`: it calls `ElicitationCompleteHandler`, if set, with the notification. The notification may have no params.

Server, `ServerSession.Elicit`:

- Nil params fail with an error that wraps the JSON-RPC invalid-params error.
- If `Mode` is empty, the server infers it. It sends `"url"` if `URL` or `ElicitationID` is set, and `"form"` otherwise. It does not change the caller's params value.
- Before sending anything, the server checks the capabilities that the client declared when it initialized. These errors are local Go errors, not JSON-RPC errors:
  - No elicitation capability: `client does not support elicitation`.
  - Mode `"form"`, and the client declared only `URL`: `client does not support "form" elicitation`. If the client declared neither `Form` nor `URL`, form is allowed, for backward compatibility.
  - Mode `"url"`, and the client did not declare `URL`: `client does not support "url" elicitation`.
- Form results are still validated against the schema, with defaults applied, as before. URL results are returned as they are.

Add tests for URL mode, for the validation errors and their codes, and for delivery of the completion notification to the client handler.
