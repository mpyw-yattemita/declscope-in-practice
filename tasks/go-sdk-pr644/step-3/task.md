# Expose JSON-RPC error codes, and support URLElicitationRequired errors (SEP-1036)

Two related changes.

## 1. Structured JSON-RPC errors

Callers need to see the JSON-RPC error codes behind the errors this SDK returns. The `jsonrpc` package now has this API. It has already been added; do not change it:

```go
// Error is a structured error in a JSON-RPC response
// (fields Code int64, Message string, Data json.RawMessage).
type Error = <the SDK's existing JSON-RPC wire error type>

const (
	CodeParseError     = -32700
	CodeInvalidRequest = -32600
	CodeMethodNotFound = -32601
	CodeInvalidParams  = -32602
	CodeInternalError  = -32603
)
```

In the `mcp` package:

- Every structured protocol error that the package creates or returns is a `*jsonrpc.Error`, so callers can use `errors.As(err, &rpcErr)` and read `rpcErr.Code`. This covers, for example, a tool call with invalid arguments, an unknown tool, an unknown prompt (all `jsonrpc.CodeInvalidParams`), the client's elicitation validation errors (`jsonrpc.CodeInvalidParams`, same messages as before), and `ResourceNotFoundError`. A tool handler that returns a `*jsonrpc.Error` still has it passed through as a protocol error, not wrapped in a tool result.
- Export the MCP-specific codes:

```go
const (
	// CodeResourceNotFound indicates that a requested resource could not be found.
	CodeResourceNotFound = -32002
	// CodeURLElicitationRequired indicates that the server requires URL elicitation
	// before processing the request.
	CodeURLElicitationRequired = -32042
)
```

`ResourceNotFoundError(uri)` keeps its message `Resource not found` and its data `{"uri":"<uri>"}`, with code `CodeResourceNotFound`.

## 2. URLElicitationRequired errors

A server can now refuse a request until the user has completed one or more URL elicitations.

Server side:

```go
// URLElicitationRequiredError returns an error indicating that URL elicitation
// is required before the request can be processed.
func URLElicitationRequiredError(elicitations []*ElicitParams) error
```

- It returns a `*jsonrpc.Error` with code `CodeURLElicitationRequired`, message `URL elicitation required`, and data `{"elicitations":[...]}`, where each element is the JSON form of an `ElicitParams`.
- It panics if any elicitation's mode is not `"url"`. An empty mode counts as `"form"`, so it panics too.
- A tool handler can return it as its error, and it reaches the client as a JSON-RPC error.

Client side: add a client sending middleware that handles these errors automatically. A constructor with no arguments returns a `Middleware`, which a user installs with `Client.AddSendingMiddleware`. Do not export the constructor yet: it is not required by the SEP, and exporting it will be proposed separately. The middleware works like this:

- It passes through every result, every error that does not have code `CodeURLElicitationRequired`, and every error for a notification method (`notifications/...`), which cannot be retried. It also passes the error through if the session's client has no `ElicitationHandler`.
- Otherwise it reads `elicitations` from the error data. If any of them is not URL mode (an empty mode counts as `"form"`), the call fails with an error containing `URLElicitationRequired error must only contain URL mode elicitations, got "<mode>"`. If the data cannot be parsed, the call fails with an error.
- For each elicitation, it runs the client's normal elicitation handling, with the same validation as an incoming `elicitation/create`. If one fails, the call fails with `URL elicitation failed: <err>`.
- It then waits until the client has received `notifications/elicitation/complete` for every elicitation ID. It must start listening before it runs the handlers, because the notification can arrive while a handler is still running. If the context ends first, the call fails with an error that wraps the context's error.
- Then it retries the original request once and returns that result.
- When a completion notification arrives, the client still calls `ClientOptions.ElicitationCompleteHandler`, if set, in addition to waking any waiting middleware.

Capabilities: a client whose `ElicitationModes` lists both `"form"` and `"url"` must advertise both `Form` and `URL`.

Add tests for the error codes of these server errors, for `URLElicitationRequiredError`, including its panics, and for the middleware's automatic retry and its rejection of non-URL modes.
