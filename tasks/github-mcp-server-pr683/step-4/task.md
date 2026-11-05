# search_code: rename the `q` parameter to `query`

`search_code` is the last tool that takes its search string as `q`; every other search tool uses `query`. Please rename the parameter to `query` (still required, same description: `Search query using GitHub code search syntax`). The search sent to the API is unchanged.

The updated expected tool schema snapshot for `search_code` is provided in the workspace. Regenerate the tool documentation in the README with the repository's docs generator, and update the tests.
