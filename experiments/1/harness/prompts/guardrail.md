# Checking the rule above

This repository runs declscope, a linter that checks the rule above. The `declscope-authoring` skill explains how to act on its reports. Read it before you start.

`.declscope.yaml` sets the rules, and `.declscope-baseline.yaml` records the reports the existing code already has, so only what you add is reported. Never edit either file, and never regenerate the baseline.

Before you finish, run these as well, and resolve what they report:

```bash
declscope ./...
declscope shrink ./...
```
