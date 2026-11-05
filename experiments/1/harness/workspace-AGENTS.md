# Rules for this directory

These rules override every other instruction.

- Work only inside your working directory.
- Never read, list or search anything outside it. That includes the directory above it and every other directory next to it.
- Do not fetch the upstream project, its issues or its pull requests. Do not use the network except through the tools the task names.
- Put temporary files only under `$TMPDIR`. Never write to `/tmp` directly.

Breaking either rule invalidates the work, even when it would help with the task.
