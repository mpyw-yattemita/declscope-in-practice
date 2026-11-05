# Checking the rule above

This repository runs declscope, a linter that checks the rule above. The `declscope-authoring` skill explains how to act on its reports. Read it before you start.

`.declscope-baseline.yaml` records the reports the existing code already has, so only what you add is reported. Never regenerate or edit it.

Before you finish, run these as well, and resolve what they report:

```bash
declscope ./...
declscope shrink ./...
```
