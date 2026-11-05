# mcp: support SSE polling and server disconnects (SEP-1699)

Implement [SEP-1699](https://github.com/modelcontextprotocol/modelcontextprotocol/issues/1699). It lets a streamable HTTP server end the SSE stream of a request it is still processing and tell the client when to reconnect. The client then reconnects and replays the events it missed, so the request still completes.

## Public API

```go
type Event struct {
	Name  string // the "event" field
	ID    string // the "id" field
	Data  []byte // the "data" field
	Retry string // the "retry" field (new)
}

type RequestExtra struct {
	TokenInfo *auth.TokenInfo
	Header    http.Header

	// CloseStream closes the current request stream, if the current transport
	// supports replaying requests.
	//
	// If reconnectAfter is nonzero, it signals to the client to reconnect after
	// the given duration. Otherwise, clients may determine their own
	// reconnection policy.
	CloseStream func(reconnectAfter time.Duration)
}
```

## Behaviour

Server-sent events:

- `Event.Retry` is written as a `retry: <value>` line, after the `id` line and before the data, and is read back when events are parsed. An event with only `Retry` set is not empty (`Event.Empty` reports false).

Streamable server transport:

- For every incoming call (a request, not a notification) received by POST, the transport sets `RequestExtra.CloseStream`, so tool, prompt and other handlers, and middleware, can close the request's stream.
- `CloseStream(d)` with `d > 0` first sends an SSE event with event name `close` and `retry` set to `d` in whole milliseconds (for example `retry: 10`), and then ends the HTTP response for that stream. `CloseStream(0)` ends the response without sending a `close` event. The `close` event is not stored for replay.
- Calling `CloseStream` again later, including after the client has reconnected, closes the current connection of that stream again. Calling it when the stream is already finished or has no connection does nothing.
- Messages sent on a closed stream are not written to the closed response. With an event store they are still stored, so the client receives them when it reconnects. A message that can be neither stored nor delivered is rejected with an error.
- Priming: when the handler has an event store (`StreamableHTTPOptions.EventStore`), the SSE response to a POST starts with an event named `prime` that has an event ID and empty data. It is also appended to the event store, so that event indexes stay aligned. The client can therefore resume the stream even if it is closed before the first real message. Replays skip stored events with empty data. The memory event store must accept empty data, which it currently rejects with a panic.

Streamable client transport:

- Events whose name is neither empty nor `message` (such as `prime`, `close` and `ping`) are not decoded as JSON-RPC messages, but their IDs still count for resumption.
- When a stream ends and the client reconnects with the last event ID, it waits for the delay given by the most recent `retry` field before its first reconnection attempt, instead of its own backoff. Later attempts use the normal backoff. If the connection is closed, or its context is cancelled while waiting, reconnection stops.

Example: a tool that calls `req.Extra.CloseStream(10*time.Millisecond)`, sends a progress notification, waits, calls `CloseStream` again, sends a second notification, and then returns. A client that uses `StreamableClientTransport` against a handler with `NewMemoryEventStore(nil)` receives both notifications in order, and the call succeeds. The same must work for a prompt handler.

Add tests for the disconnect and replay flow, for the priming event, and for `CloseStream` with and without a reconnect delay.
