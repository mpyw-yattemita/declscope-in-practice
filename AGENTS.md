# Repository instructions

This repository runs a controlled experiment on declscope. Read [design.md](design.md) before changing the harness, a task or the analysis.

## Contracts to preserve

- Never create a run workspace inside this repository. Its path would show the repository name to the agent under measurement. The workspace root lives outside, and its path must not mention declscope or the experiment.
- Nothing the agent under measurement can read may mention declscope, the arms, or this experiment. That covers the workspace path, `harness/workspace-AGENTS.md`, the task text, and arm A's system prompt. Arm B alone receives declscope, its baseline and the `declscope-authoring` skill.
- Run directories get random names, so the path does not show the arm.
- A seed reaches the workspace through `git archive` and one new commit. Never clone it with history, since history holds the answer.
- Hidden tests enter the workspace only after the agent finishes.
- Audit every transcript before its run counts. Discard a run that touched a path outside its workspace or used the network to reach the upstream project.
- The reviewer rubric must not use declscope's terms such as namespace, private, shared or boundary. Strip directives and declscope files before review.
- Results in `runs/` are raw data. Never edit them by hand. Regenerate them.
