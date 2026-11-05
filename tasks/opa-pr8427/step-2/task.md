# Allow back-to-back METADATA blocks

If a policy has two `METADATA` blocks one right after the other, with no empty line between them, parsing fails with a misleading error:

```rego
package play

# METADATA
# title: Ps
# scope: document
# METADATA
# title: P
p := rego.metadata.chain()
```

```
policy.rego:7: rego_parse_error: yaml: line 4: could not find expected ':'
```

The second `# METADATA` line is read as part of the first block's YAML. The author clearly meant to start a new block.

Please make a `# METADATA` comment line always start a new annotation block, even when it directly follows another block's lines. The example above must parse into two annotation blocks (`title: Ps` with document scope, and `title: P` on the rule).

This also changes the error when a `# METADATA` line directly after a block has no content of its own. For example:

```rego
# METADATA
# scope: rule
# schemas:
#   - data.servers: schema.servers
# METADATA
public_servers[server] {
	server = servers[i]
}
```

The second `# METADATA` line used to be read into the first block and gave a YAML error (`yaml: ... could not find expected ':'`). Now it starts a second, empty block, and the parser reports `rego_parse_error: expected METADATA block, found whitespace`.

Add parser tests, including two blocks for the same rule with no blank line between them.
