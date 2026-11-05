# Design

## The question

[declscope](https://github.com/mpyw/declscope) is a Go linter. It gives a flat package a `private` scope per file. Helping an existing codebase is the obvious use. This experiment asks a different question:

> When an AI agent adds new code to an existing package, does running declscope make the result easier to change and harder to criticise?

## Conditions

| | Arm A | Arm B |
| --- | --- | --- |
| Seed | The same commit, without history | The same commit, without history |
| declscope | Absent | On `PATH`, with a baseline of the seed's existing findings |
| Lint loop | `go vet` and `go test` | `go vet`, `go test`, `declscope ./...` and `declscope shrink ./...` |
| Skill | None | `declscope-authoring`, as a project skill in the workspace |
| Model, prompt, tools | Same | Same |

The baseline absorbs what the seed already has. Arm B is reported only for what the agent adds.

> [!NOTE]
> A baseline key names the declaration and its own namespace, not the user. A new use of a declaration the baseline already records is not reported. Arm B therefore misses some new crossings. This biases the result against declscope, so the measured effect is a lower bound.

Arm B may write `//declscope:` directives, since that is how the tool is meant to be used. It may not regenerate the baseline. Both arms are told not to change existing code beyond what the task needs.

## Tasks

Each task comes from a real merged pull request of a project that does not use declscope.

| Part | Source |
| --- | --- |
| Seed | The base commit of the pull request |
| Task text | The pull request and its issue, plus the signatures the hidden tests call |
| Hidden tests | The tests the pull request added |
| Follow-up task | A later pull request that changes the same package |
| Follow-up hidden tests | The tests the later pull request added |

A task is usable when all of these hold:

| Criterion | Why |
| --- | --- |
| The package has 8 or more non-test files and 3 or more namespaces, by `declscope survey` | Below that, declscope has little to say |
| The pull request was merged after the training cutoff of the model under measurement. For Claude Opus 5.5, that means July 2026 or later | Earlier ones may be in the model's training data |
| It adds a feature, touches 2 or more files of the package, and changes 50 to 500 lines | It needs placement decisions, and it fits one agent run |
| Its tests call exported API or observable behaviour | Tests that name unexported identifiers break when the agent names things differently |
| `go test` passes without network or external services | The tests must run in the workspace |
| A later pull request changes the same package | It becomes the follow-up task |

## Two stages

| Stage | What runs | What it measures |
| --- | --- | --- |
| 1. Implement | Each arm implements the task | Hidden tests passed, tokens, turns, directives written |
| 2. Change | A fresh agent implements the follow-up task on stage 1's result. declscope, its baseline and its directives are removed first, in both arms | How easy the code is to change |

Stage 2 is the main measure. Its agent has no declscope, so any difference comes from the structure stage 1 left.

| Stage 2 metric | Role | Meaning |
| --- | --- | --- |
| Follow-up hidden tests passed | Main | The change succeeded |
| Existing tests still passing | Main | The change broke nothing |
| Tokens and turns | Main | The code was easy to understand |
| Files and lines touched | Reference only | How far the change spread |

> [!NOTE]
> Files touched has no good direction. declscope tends to split code by job into more files. A change may then touch more files and still be easier. So it is recorded, not used to judge.

## Structural review

A blind reviewer reads each stage 1 result for refactoring that should be done now.

| Rule | Why |
| --- | --- |
| Directives, the baseline and declscope files are stripped first | The reviewer must not know the arm |
| The rubric never uses declscope's terms or its criterion | Teaching the reviewer declscope's idea would make the measure circular |
| The reviewer writes findings freely. Each has a file, a line and a concrete fix | Vague findings cannot be counted. A fixed list of categories would steer the reviewer |
| A second agent verifies each finding | It removes padding |
| A third agent sorts the verified findings into the categories below | The reviewer never sees the categories |
| Only findings on changed lines, or absent from the seed's review, count | The seed's own problems are the same in both arms |
| Counts are per 1,000 changed lines | The arms write different amounts of code |

| Category | Example |
| --- | --- |
| `naming` | A name does not match its role, or one idea has two names |
| `duplication` | The same logic is written twice |
| `over-abstraction` | A shared helper couples callers that are not related |
| `mixed-job` | One function or type does two unrelated jobs |
| `dead` | A definition is unused |
| `other` | None of the above |

> [!IMPORTANT]
> No category may restate declscope's rule. "A helper meant for one file is used from another" is that rule in other words. A category like that favours arm B by construction, so it is not one of the categories. A finding of that kind goes to `other` and is reported separately.

The per-category result says where the tool helps and where it does not.

## Isolation

The agents under measurement must not see this repository, the other arm, or the upstream answer.

| Leak | Guard |
| --- | --- |
| This repository's files and its name | Workspaces live outside it, under a directory whose path says nothing about the experiment |
| Which arm a run is | Run directories have random names |
| Other runs | An `AGENTS.md` at the workspace root forbids reading outside the working directory. `harness/workspace-AGENTS.md` is its source |
| The upstream answer in git history | The seed is exported with `git archive` and committed once into a new repository |
| The upstream answer on the network | `WebFetch` and `WebSearch` are disallowed. The instructions forbid fetching. `GOPROXY=off` after the module cache is warmed |
| Hidden tests | They stay out of the workspace until the agent finishes |

The instructions are not enforced by a sandbox. Every transcript is audited instead. A run that touched a path outside its workspace or ran `gh`, `curl`, `wget`, `git fetch` or `git clone` is discarded and run again.

> [!IMPORTANT]
> Nothing an agent under measurement can read may mention declscope, the arms, or this experiment. That covers the workspace path, the workspace `AGENTS.md`, the task text, and arm A's system prompt.

## Scale

The first round uses 4 or 5 repositories and 2 or 3 pull requests each, about 10 tasks. Each arm runs 5 times per task. That is about 100 stage 1 runs and 100 stage 2 runs.

## Layout

| Path | Holds |
| --- | --- |
| `candidates/` | `survey` results and the record of how repositories and pull requests were chosen |
| `tasks/<repo>-pr<N>/task.yaml` | The seed commit, task text, hidden tests and follow-up |
| `harness/` | Export, baseline, run, collect and audit |
| `runs/<task>/<arm>/<n>/` | Diffs, `metrics.json` and gzipped transcripts |
| `analysis/` | Aggregation scripts and result tables |
