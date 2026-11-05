# mcp: apply schema defaults to elicitation results (SEP-1034)

[SEP-1034](https://github.com/modelcontextprotocol/modelcontextprotocol/issues/1034) lets an elicitation schema give default values for its properties. Our elicitation schemas are intentionally decoupled from their JSON-RPC implementation, but when we validate an elicitation result against the requested schema, we should also apply the schema's defaults.

## Behaviour

- When the client validates an elicitation result against the requested schema (the existing form validation), it then applies the schema's defaults to `ElicitResult.Content`. A property that the handler left out gets its default value; properties the handler returned are kept.
- `ServerSession.Elicit` does the same on the server side: when `ElicitParams.RequestedSchema` is set, it validates the returned content against the schema and applies its defaults before returning the result. If the content does not match, `Elicit` returns an error whose message starts with `elicitation result content does not match requested schema: `. A nil schema skips validation.
- Example: with the schema `{"type":"object","properties":{"key":{"type":"string","default":"potato"}}}` and a handler that returns `{"default":"response"}`, `Elicit` returns the content `{"key":"potato","default":"response"}`. This works for `boolean`, `string` (including enums), `integer` and `number` properties. Numeric defaults come back as `float64`.

## Validating defaults in the schema

Schema validation already rejects a boolean property whose default is not a boolean. Extend this to the other primitive types, and use these messages (`<name>` is the property name):

| Property type | Bad default | Error message (prefix) |
| --- | --- | --- |
| `boolean` | e.g. `"not-a-boolean"` | `elicit schema property "<name>" has invalid default value, must be a bool` |
| `string` | e.g. `true` | `elicit schema property "<name>" has invalid default value, must be a string` |
| `integer` or `number` | e.g. `true` | `elicit schema property "<name>" has default value that cannot be interpreted as an int or float` |

Note that the boolean message changes from `must be a boolean` to `must be a bool`.

Add tests for applying defaults of each type and for the invalid defaults.
