# declscope

This repository runs declscope, a linter that checks where each declaration of a Go package is used and how it is named. The `declscope-authoring` skill explains its rules and how to act on its reports. Read it before you start.

`.declscope.yaml` sets the rules, and `.declscope-baseline.yaml` records the reports the existing code already has, so only new reports are shown. Never edit either file, and never regenerate the baseline.

Before you finish, run these as well, and resolve what they report:

```bash
declscope ./...
declscope shrink ./...
```
