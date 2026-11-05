# Design

## The question

[declscope](https://github.com/mpyw/declscope) is a Go linter. It gives a flat package a `private` scope per file. Helping an existing codebase is the obvious use. This experiment asks a different question:

> When an AI agent adds new code to an existing package, does running declscope make the result easier to change and harder to criticise?

Models now follow written rules well. So the experiment also asks a second question:

> Does a deterministic check add anything over stating the same rule in the prompt?

## Conditions

Three arms are nested. Each adds one thing to the arm before it.

| Arm | Adds | The difference from the previous arm shows |
| --- | --- | --- |
| **A**: none | The task and the shared self-review step | Nothing. This is the control |
| **P**: prompt | The rule of `declscope-authoring`, rewritten in plain words without the tool | What stating the rule does |
| **G**: guardrail | declscope itself, with a baseline, the lint loop and directives | What a deterministic check adds to the stated rule |

| | A | P | G |
| --- | --- | --- | --- |
| Seed | The same commit, without history | Same | Same |
| Self-review step | Yes | Yes | Yes |
| Placement rule in the prompt | No | Yes | Yes |
| declscope on `PATH`, with a baseline of the seed's findings | No | No | Yes |
| Checks run before finishing | `go vet`, `go test` | `go vet`, `go test` | `go vet`, `go test`, `declscope ./...`, `declscope shrink ./...` |
| `declscope-authoring` as a project skill | No | No | Yes |
| Model and tools | Same | Same | Same |

Every arm gets the same self-review step before its checks:

```text
Before you finish, review the structure of your change and fix anything
you would flag in code review. Then run the checks below until they pass.
```

> [!IMPORTANT]
> The step gives every arm the same chance to look again. Without it, arm G alone would revise its code after a report. Its result would then mix declscope's effect with the effect of one more revision.

Arm P's rule is the content of `declscope-authoring` with the tool removed. It says the same thing as the tool enforces. It names no command, directive or config. P and G must differ only in whether a machine checks the rule.

The baseline absorbs what the seed already has. Arm G is reported only for what the agent adds.

> [!NOTE]
> A baseline key names the declaration and its own namespace, not the user. A new use of a declaration the baseline already records is not reported. Arm G therefore misses some new crossings. This biases the result against declscope, so the measured effect is a lower bound.

Arm G may write `//declscope:` directives, since that is how the tool is meant to be used. It may not regenerate the baseline. Every arm is told not to change existing code beyond what the task needs.

| Outcome | Reading |
| --- | --- |
| A < P, and P ≈ G | Stating the rule is enough for the model under measurement |
| A < P < G | The check catches what the stated rule misses |
| A ≈ P ≈ G | Neither the rule nor the check helps on these tasks |

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
| It adds a feature, touches 2 or more files of the package, and changes 50 to 500 lines | It needs placement decisions, and it fits one agent run |
| Its tests call exported API or observable behaviour | Tests that name unexported identifiers break when the agent names things differently |
| `go test` passes without network or external services | The tests must run in the workspace |
| A later pull request changes the same package | It becomes the follow-up task |
| The follow-up's tests name only what the seed has and what the first task text specifies | A name upstream chose in the first pull request may differ from the agent's. The test would then fail on the name, not on harder code |
| No other commit changes the package between the two pull requests | The follow-up must start from the same ground upstream and in the workspace |

> [!IMPORTANT]
> A good task comes first. The merge date only breaks ties.

The model under measurement may have seen an older pull request in training. It may then recall the upstream code instead of designing its own. That helps no arm over another, but it narrows the gaps between them. So among suitable pull requests, those merged after the model's training cutoff are preferred. For Claude Opus 5.5, that means July 2026 or later.

| Group | Use |
| --- | --- |
| Merged after the cutoff | Preferred |
| Merged before the cutoff | Used when it makes a better task than the later ones |

The two groups are reported apart. A small gap between them shows that recall had little effect.

## Two stages

| Stage | What runs | What it measures |
| --- | --- | --- |
| 1. Implement | Each arm implements the task | Hidden tests passed, tokens, turns, directives written |
| 2. Change | A fresh agent implements the follow-up task on stage 1's result. declscope, its baseline and its directives are removed first, in every arm. The agent gets arm A's setup in every case | How easy the code is to change |

Stage 2 is the main measure. Its agent has no declscope, so any difference comes from the structure stage 1 left.

| Stage 2 metric | Role | Meaning |
| --- | --- | --- |
| Follow-up hidden tests passed | Main | The change succeeded |
| Existing tests still passing | Main | The change broke nothing |
| Tokens and turns | Main | The code was easy to understand |
| Files and lines touched | Reference only | How far the change spread |

> [!NOTE]
> Files touched has no good direction. declscope tends to split code by job into more files. A change may then touch more files and still be easier. So it is recorded, not used to judge.

## Decision rule

This rule was fixed before any run. It decides the result. Every other number is reported as exploratory.

Two stage 2 numbers decide each comparison of two arms.

| Number | Role |
| --- | --- |
| Success rate: follow-up hidden tests pass and existing tests still pass | A guard. The better arm must not lose on it |
| Tokens: every input and output token of the stage 2 run, cached input included | The measure. The better arm uses fewer |

| Success rate of X against Y | Tokens of X against Y | Result |
| --- | --- | --- |
| Not lower | Fewer | X wins |
| Not lower | No clear difference | No difference |
| Lower | Any | X loses |

| Term | Definition |
| --- | --- |
| Not lower | X's success rate is at most 5 percentage points below Y's |
| Lower | X's success rate is more than 5 percentage points below Y's |
| Fewer | Each task gives the ratio of X's median tokens to Y's. The 95% bootstrap interval of the mean ratio over tasks lies wholly below 1 |

A failed stage 2 run still counts its tokens. The same rule applies to A against P, P against G and A against G.

## Structural review

A blind reviewer reads each stage 1 result for refactoring that should be done now.

| Rule | Why |
| --- | --- |
| Directives, the baseline and declscope files are stripped first | The reviewer must not know the arm |
| The rubric never uses declscope's terms or its criterion | Teaching the reviewer declscope's idea would make the measure circular |
| The reviewer writes findings freely. Each has a file, a line and a concrete fix | Vague findings cannot be counted. A fixed list of categories would steer the reviewer |
| A second agent verifies each finding | It removes padding |
| A third agent sorts the verified findings into the categories below | The reviewer never sees the categories |
| Only findings on changed lines, or absent from the seed's review, count | The seed's own problems are the same in every arm |
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
> No category may restate declscope's rule. "A helper meant for one file is used from another" is that rule in other words. A category like that favours arms P and G by construction, so it is not one of the categories. A finding of that kind goes to `other` and is reported separately.

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
> Nothing an agent under measurement can read may mention declscope, the arms, or this experiment. That covers the workspace path, the workspace `AGENTS.md`, the task text, and the prompts of arms A and P. Arm G alone sees declscope.

## Scale

The first round uses 4 or 5 repositories and 2 or 3 pull requests each, about 10 tasks. Each arm runs 5 times per task. That is about 150 stage 1 runs and 150 stage 2 runs. If usage limits make that too slow, the number of runs or tasks is cut once the candidates are known.

## Layout

| Path | Holds |
| --- | --- |
| `candidates/` | `survey` results and the record of how repositories and pull requests were chosen |
| `tasks/<repo>-pr<N>/task.yaml` | The seed commit, task text, hidden tests and follow-up |
| `harness/` | Export, baseline, run, collect and audit |
| `runs/<task>/<arm>/<n>/` | Diffs, `metrics.json` and gzipped transcripts |
| `analysis/` | Aggregation scripts and result tables |
