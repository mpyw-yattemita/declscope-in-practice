# ast: fix parentheses in String() of object, array and set comprehensions

When a comprehension's head (the term of an array or set comprehension, or the key or value of an object comprehension) is an infix operator call, `String()` prints it as a prefix call, for example `[mul(x, 2) | ...]` for `[(x * 2) | ...]`. Printed without parentheses in infix form it would be ambiguous with the `|` separator, and the prefix form is hard to read and differs from what the user wrote.

Please render such heads in infix form wrapped in parentheses, recursively for nested infix calls. The comprehension body keeps its current rendering. Examples (`String()` of the parsed term):

| Input | Output |
| --- | --- |
| `[(x * 2) \| x := data.numbers[_]; x > 5]` | `[(x * 2) \| assign(x, data.numbers[_]); gt(x, 5)]` |
| `[((x + 1) * 2) \| x := data.numbers[_]]` | `[((x + 1) * 2) \| assign(x, data.numbers[_])]` |
| `[count(x) \| x := data.lists[_]]` | `[count(x) \| assign(x, data.lists[_])]` |
| `[(to_number(cpu_str[i])/1000) \| cpu_str[i]]` | `[(to_number(cpu_str[i]) / 1000) \| cpu_str[i]]` |
| `{k: (v * 100) \| some k, v; v > 0}` | `{k: (v * 100) \| some k, v; gt(v, 0)}` |
| `{(k + 1): v \| some k, v; v > 0}` | `{(k + 1): v \| some k, v; gt(v, 0)}` |

Only built-in calls with an infix operator and two operands are rendered this way; the membership operator `in` is not. Other heads (function calls, refs, variables) are printed as before.

`AppendText` must produce the same text as `String()`, and `StringLength()` must return its exact length, so that appending to a buffer pre-sized with `StringLength()` does not allocate.

Add tests for these cases.
