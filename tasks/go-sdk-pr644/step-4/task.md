# fix: handle Windows CRLF in the MCP client stdio transport

On Windows, a Go MCP client that connects over stdio to a Python MCP server (for example `uvx mcp-server-fetch` through `CommandTransport`) fails during initialize:

```
calling "initialize": invalid trailing data at the end of stream
```

The stdio framing accepts only `\n` after a JSON message. The Python server writes Windows line endings (`\r\n`), so the client treats the `\r` as trailing data.

Please accept both Unix (`\n`) and Windows (`\r\n`) line endings after a message, and keep the existing behaviour on Unix. For example, the input `{"jsonrpc":"2.0","id":1,"method":"test","params":{}}` followed by `\r\n` must be read as one message, with no error. Any other trailing data must still be reported as `invalid trailing data at the end of stream`.

Add a test case for the Windows line ending.
