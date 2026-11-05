# Candidate PR chains (v3)

v3 criteria: a fix with tests counts as a step if it fixes externally observable behaviour (judged by reading the PR); unexported test names no longer reject a chain but each step has a rename-risk note; snapshot files of external spec ship with the task text. At most 2 chains per repository.

Generated 2026-11-05. Each chain is 3-5 consecutive merged PRs on one package: no other commit touched the package's own `.go` files between members except the trivial commits listed. Lines are non-test `.go` lines (added+removed) in the package; `whole` counts every non-test `.go` file in the PRs. Chain-introduced names are the identifiers that a step's added test lines use and that some chain step (up to that one) introduced, checked by token matching and, for unexported ones, by reading the test lines. Offline tests: `go test -count=1 ./<pkg>` under `sandbox-exec` with outbound network denied except localhost, `GOPROXY=off`, `CGO_ENABLED=0`, Go 1.27.0 darwin/arm64, at the seed and at every step's merge commit (so each step's own tests ran where they were merged).

## Ranking

| # | Repo | Package (files/ns) | Steps | Merged | Pkg lines (whole) | Pkg files | Chain-introduced names in tests (unexported) | Rename risk: worst / high steps | vs v2 | Offline: seed / steps | Why it gives placement decisions |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | hashicorp/hcl | `hclwrite` (17/18) | [#841](https://github.com/hashicorp/hcl/pull/841) → [#811](https://github.com/hashicorp/hcl/pull/811) → [#820](https://github.com/hashicorp/hcl/pull/820) → [#835](https://github.com/hashicorp/hcl/pull/835) | 2026-10-13 .. 2026-10-19 | 788 (788) | 10 | 7 (1) | medium / 0 | kept | pass / pass; pass; pass; pass | Each step adds a new expression kind or API wired through parser.go, ast_expression.go and the node/unwrap plumbing; the agent chooses between new files and existing ones, and #820/#835 build on #811's unwrap/variables code. |
| 2 | modelcontextprotocol/go-sdk | `mcp` (25/34) | [#975](https://github.com/modelcontextprotocol/go-sdk/pull/975) → [#989](https://github.com/modelcontextprotocol/go-sdk/pull/989) → [#997](https://github.com/modelcontextprotocol/go-sdk/pull/997) → [#987](https://github.com/modelcontextprotocol/go-sdk/pull/987) | 2026-06-02 .. 2026-06-08 | 617 (646) | 7 | 13 (2) | high / 1 | kept | pass / pass; pass; pass; pass | One protocol feature (SEP-2575 discover / version negotiation) split into client side, retry, logging level and server side; each step touches 3-6 of client.go/server.go/shared.go/protocol.go/streamable.go and later steps call earlier steps' discover/usesNewProtocol code. |
| 3 | github/github-mcp-server | `pkg/github` (77/113) | [#2474](https://github.com/github/github-mcp-server/pull/2474) → [#2451](https://github.com/github/github-mcp-server/pull/2451) → [#2284](https://github.com/github/github-mcp-server/pull/2284) → [#2445](https://github.com/github/github-mcp-server/pull/2445) → [#2480](https://github.com/github/github-mcp-server/pull/2480) | 2026-05-20 .. 2026-05-21 | 1233 (1233) | 9 | 30 (2) | high / 2 | kept | pass / pass; pass; pass; pass; pass | Five tool features across issues.go, search_utils.go, search.go, minimal_types.go, the three GHAS files and a new issue_fields.go; #2480 reuses #2445's IssueField types and query structs. |
| 4 | open-policy-agent/opa | `v1/ast` (50/73) | [#8427](https://github.com/open-policy-agent/opa/pull/8427) → [#8482](https://github.com/open-policy-agent/opa/pull/8482) → [#8485](https://github.com/open-policy-agent/opa/pull/8485) → [#8468](https://github.com/open-policy-agent/opa/pull/8468) → [#8511](https://github.com/open-policy-agent/opa/pull/8511) | 2026-03-30 .. 2026-04-14 | 875 (997) | 11 | 5 (1) | medium / 0 | new | pass (Go 1.25.5) / pass (Go 1.25.5); pass (Go 1.25.5); pass (Go 1.25.5); pass (Go 1.25.5); pass (Go 1.25.5) | #8427 adds a `Not` AST node across 10 files (term, compare, visit, compile, parser, string-length, mermaid); later steps are parser, type-checker, compiler and String() changes that touch the same code paths and use the new node's methods. |
| 5 | jaegertracing/jaeger | `internal/storage/v2/api/tracestore` (16/17) | [#9256](https://github.com/jaegertracing/jaeger/pull/9256) → [#9356](https://github.com/jaegertracing/jaeger/pull/9356) → [#9357](https://github.com/jaegertracing/jaeger/pull/9357) | 2026-08-07 .. 2026-08-19 | 565 (851) | 4 | 17 (0) | none / 0 | new | pass / pass; pass; pass | Storage capability declarations grow step by step (search without service name, then filter capabilities, then query shaping in new admission/shape code that consumes them); the agent decides where capability types and query-shape logic live. |
| 6 | github/github-mcp-server | `pkg/github` (77/113) | [#683](https://github.com/github/github-mcp-server/pull/683) → [#470](https://github.com/github/github-mcp-server/pull/470) → [#690](https://github.com/github/github-mcp-server/pull/690) → [#740](https://github.com/github/github-mcp-server/pull/740) → [#774](https://github.com/github/github-mcp-server/pull/774) | 2025-07-21 .. 2025-07-29 | 1247 (1247) | 10 | 13 (0) | none / 0 | kept (now includes #740) | pass / pass; pass; pass; pass; pass | #683 adds pagination parameter types across 9 tool files; #470, #690 and #774 use them, so the first step's placement of shared helpers is exercised by every later step. |
| 7 | modelcontextprotocol/go-sdk | `mcp` (25/34) | [#644](https://github.com/modelcontextprotocol/go-sdk/pull/644) → [#646](https://github.com/modelcontextprotocol/go-sdk/pull/646) → [#655](https://github.com/modelcontextprotocol/go-sdk/pull/655) → [#665](https://github.com/modelcontextprotocol/go-sdk/pull/665) → [#663](https://github.com/modelcontextprotocol/go-sdk/pull/663) | 2025-11-21 .. 2025-12-01 | 701 (766) | 9 | 18 (2) | medium / 0 | kept (now includes #644 and #665) | pass / pass; pass; pass; pass; pass | Elicitation defaults, URL-mode elicitation, error codes / URLElicitationRequired built on it, a CRLF fix, then SSE polling; spread over client.go, server.go, shared.go, protocol.go, transport.go, streamable.go, event.go. |
| 8 | go-git/go-git | `plumbing/protocol/packp` (25/27) | [#2225](https://github.com/go-git/go-git/pull/2225) → [#2227](https://github.com/go-git/go-git/pull/2227) → [#2237](https://github.com/go-git/go-git/pull/2237) → [#2263](https://github.com/go-git/go-git/pull/2263) → [#2264](https://github.com/go-git/go-git/pull/2264) | 2026-06-26 .. 2026-07-23 | 1606 (2357) | 10 | 63 (3) | high / 1 | new | pass / pass; pass; pass; pass; pass | Protocol v2 support built in stages: slice-based AdvRefs, new v2 wire types (capability_adv, command, fetch, ls-refs), switching the default, then input-validation fixes on the new types. |
| 9 | traefik/traefik | `pkg/provider/kubernetes/gateway` (9/9) | [#14130](https://github.com/traefik/traefik/pull/14130) → [#14119](https://github.com/traefik/traefik/pull/14119) → [#14127](https://github.com/traefik/traefik/pull/14127) → [#14123](https://github.com/traefik/traefik/pull/14123) → [#13410](https://github.com/traefik/traefik/pull/13410) | 2026-10-23 .. 2026-10-28 | 597 (970) | 6 | 4 (3) | high / 1 | kept | pass / pass; pass; pass; pass; pass | Five Gateway API features that each touch features.go and httproute.go plus some of kubernetes.go/client.go/grpcroute.go/tlsroute.go; #14119 and #13410 share TLS/CA loading. |
| 10 | slack-go/slack | `.` (125/149) | [#1551](https://github.com/slack-go/slack/pull/1551) → [#1553](https://github.com/slack-go/slack/pull/1553) → [#1552](https://github.com/slack-go/slack/pull/1552) → [#1555](https://github.com/slack-go/slack/pull/1555) → [#1560](https://github.com/slack-go/slack/pull/1560) | 2026-04-17 .. 2026-06-04 | 734 (734) | 12 | 103 (0) | none / 0 | kept (now includes #1560) | pass except TestRTMDisconnect (dials slack.com) / pass except TestRTMDisconnect (dials slack.com); pass except TestRTMDisconnect (dials slack.com); pass except TestRTMDisconnect (dials slack.com); pass except TestRTMDisconnect (dials slack.com); pass except TestRTMDisconnect (dials slack.com) | New Block Kit blocks registered in block.go/block_conv.go plus new files; #1552/#1555/#1560 reuse the chunk/card/table types. |
| 11 | gofiber/fiber | `middleware/cache` (10/17) | [#4735](https://github.com/gofiber/fiber/pull/4735) → [#4736](https://github.com/gofiber/fiber/pull/4736) → [#4744](https://github.com/gofiber/fiber/pull/4744) → [#4590](https://github.com/gofiber/fiber/pull/4590) | 2026-10-03 .. 2026-10-12 | 646 (695) | 9 | 10 (10) | high / 2 | new | pass / pass; pass; pass; pass | Four observable HTTP-cache fixes (OWS parsing, cache key by scheme/host, Vary replay, stale-entry deletion) that touch cache.go, manager.go, config.go and add vary.go/heap.go logic; later steps modify the cache entry structure earlier steps extended. |
| 12 | ollama/ollama | `cmd/launch` (28/37) | [#16516](https://github.com/ollama/ollama/pull/16516) → [#16410](https://github.com/ollama/ollama/pull/16410) → [#16558](https://github.com/ollama/ollama/pull/16558) → [#16683](https://github.com/ollama/ollama/pull/16683) → [#15434](https://github.com/ollama/ollama/pull/15434) | 2026-06-04 .. 2026-06-18 | 698 (698) | 6 | 17 (7) | high / 1 | new | pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken) / pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken); pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken); pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken); pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken); pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken) | Integration launchers (Hermes desktop, oh-my-pi, opencode) in a 28-file flat package; registry.go and launch.go are shared, and #16558 changes #16516's Hermes code. |

## Chains in detail

### 1. hashicorp/hcl `hclwrite`

Seed `a786e4fc5566`, final `4a347c448ee9`. Package lines 788, whole PRs 788. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #841 hclwrite: find traversals inside object-construct expressions | 2026-10-13 | 25 / ast_expression.go, ast_object_cons_expr.go, parser.go | 25 | none | 0 | - | **none**: tests use only exported names or names present at the seed | pass |
| 2 | #811 hclwrite: add read and write API for `ObjectConsExpr` | 2026-10-14 | 452 / ast.go, ast_expression.go, ast_object_cons_expr.go, node.go, parser.go, unwrap.go | 452 | `ObjectConsExpr.ItemFor`, `ObjectConsExpr.SetItem`, `ObjectConsExpr.SetItemRaw`, `ObjectConsExpr.SetItemTraversal`, `ObjectConsExpr.SetItemValue`, `ObjectConsExpr.ValueExprFor` | 0 | none found | **none**: tests use only exported names or names present at the seed | pass |
| 3 | #820 hclwrite: parse function call and tuple-construct expressions | 2026-10-14 | 221 / ast_expression.go, ast_function_call_expr.go, ast_tuple_cons_expr.go, parser.go | 221 | none | 0 | `Expression.unwrap`, `ObjectConsExpr.variables`, `ObjectConsKeyExpr.unwrap`, `unwrapNode.unwrap` | **none**: tests use only exported names or names present at the seed | pass |
| 4 | #835 hclwrite: New function `NewExpression()` and `Body.SetAttribute()` | 2026-10-19 | 90 / ast_body.go, parser.go, public.go | 90 | `parseExpr` | 0 | `FunctionCallExpr.name` | **medium**: `parseExpr(src, filename, pos)` is the expression-parse entry point; the agent almost certainly has an equivalent, but it may be named or shaped differently | pass |

Trivial commits inside the chain: none.

Placement: Each step adds a new expression kind or API wired through parser.go, ast_expression.go and the node/unwrap plumbing; the agent chooses between new files and existing ones, and #820/#835 build on #811's unwrap/variables code.

Caveats: #810 cannot join a strictly consecutive chain without #824 (51-line Body feature) in between; a 5-step variant #824 → #841 → #811 → #820 → #835 (839 lines) also qualifies.

### 2. modelcontextprotocol/go-sdk `mcp`

Seed `d54751a8eec7`, final `dd978160abb3`. Package lines 617, whole PRs 646. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #975 mcp: Implement support for SEP-2575 on client side | 2026-06-02 | 227 / client.go, protocol.go, requests.go, server.go, shared.go, streamable.go | 233 | `CodeUnsupportedProtocolVersion`, `DiscoverParams`, `DiscoverResult`, `DiscoverResult.Capabilities`, `DiscoverResult.Instructions`, `DiscoverResult.ServerInfo`, `DiscoverResult.SupportedVersions`, `methodDiscover` | 0 | - | **low**: only the unexported method-name constant `methodDiscover` ("server/discover"); if the agent inlines the string there is no named counterpart, but the renamer can substitute the literal only by changing a call site, which the rules allow | pass |
| 2 | #989 mcp: Implement retry on advised supportedVersions in UnsupportedProtocolVersion error | 2026-06-08 | 94 / client.go, protocol.go, shared.go | 94 | `CodeUnsupportedProtocolVersion`, `DiscoverParams`, `DiscoverResult`, `DiscoverResult.Capabilities`, `DiscoverResult.ServerInfo`, `DiscoverResult.SupportedVersions`, `UnsupportedProtocolVersionData`, `UnsupportedProtocolVersionData.Requested`, `UnsupportedProtocolVersionData.Supported`, `methodDiscover` | 0 | `Client.discover`, `CodeUnsupportedProtocolVersion`, `DiscoverParams`, `DiscoverResult`, `DiscoverResult.Capabilities` | **low**: same `methodDiscover` constant | pass |
| 3 | #997 mcp: add support for logging level (SEP-2575) | 2026-06-08 | 44 / protocol.go, server.go, shared.go, streamable.go | 44 | `MetaKeyLogLevel`, `validatedMeta.logLevel` | 0 | `ClientSession.usesNewProtocol` | **high**: test reads the unexported field `validatedMeta.logLevel` of an internal per-request struct; no counterpart unless the agent kept the same struct | pass |
| 4 | #987 mcp: Implement server-side support for discover method (SEP-2575) | 2026-06-08 | 252 / server.go, shared.go, streamable.go, transport.go | 275 | `CodeUnsupportedProtocolVersion`, `DiscoverResult`, `DiscoverResult.Capabilities`, `DiscoverResult.ServerInfo`, `DiscoverResult.SupportedVersions`, `methodDiscover` | 0 | `Client.discover`, `ClientSession.usesNewProtocol`, `CodeUnsupportedProtocolVersion`, `DiscoverParams`, `DiscoverResult` | **low**: `methodDiscover` constant | pass |

Trivial commits inside the chain: none.

Placement: One protocol feature (SEP-2575 discover / version negotiation) split into client side, retry, logging level and server side; each step touches 3-6 of client.go/server.go/shared.go/protocol.go/streamable.go and later steps call earlier steps' discover/usesNewProtocol code.

Caveats: #975 and #987 also change internal/jsonrpc2 slightly.

### 3. github/github-mcp-server `pkg/github`

Seed `e9533330b0ec`, final `13e7f10fc8dd`. Package lines 1233, whole PRs 1233. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #2474 Add field_values to search_issues results | 2026-05-20 | 215 / issues.go, search_utils.go | 215 | `SearchIssueResult.FieldValues`, `SearchIssuesResponse`, `SearchIssuesResponse.IncompleteResults`, `SearchIssuesResponse.Items`, `SearchIssuesResponse.Total` | 0 | - | **none**: tests use only exported names or names present at the seed | pass |
| 2 | #2451 feat: add pagination to list GHAS alerts tools | 2026-05-20 | 221 / code_scanning.go, dependabot.go, secret_scanning.go | 221 | none | 3 | none found | **none**: tests use only exported names or names present at the seed | pass |
| 3 | #2284 feat: Add search commit tool | 2026-05-20 | 224 / minimal_types.go, search.go, tools.go | 224 | `MinimalCommitSearchItem.Repository`, `MinimalRepoRef.FullName`, `MinimalRepoRef.HTMLURL`, `MinimalRepoRef.Private`, `MinimalSearchCommitsResult`, `MinimalSearchCommitsResult.IncompleteResults`, `MinimalSearchCommitsResult.Items`, `MinimalSearchCommitsResult.TotalCount`, `SearchCommits`, `SearchIssuesResponse.IncompleteResults`, `SearchIssuesResponse.Items`, `SearchIssuesResponse.Total` | 1 | `SearchIssuesResponse.IncompleteResults`, `SearchIssuesResponse.Items`, `SearchIssuesResponse.Total` | **none**: tests use only exported names or names present at the seed | pass |
| 4 | #2445 Add list_issue_fields tool | 2026-05-21 | 224 / issue_fields.go, tools.go | 224 | `IssueField`, `IssueField.DataType`, `IssueField.Description`, `IssueField.ID`, `IssueField.Name`, `IssueField.Options`, `IssueField.Visibility`, `IssueSingleSelectFieldOption`, `IssueSingleSelectFieldOption.Color`, `IssueSingleSelectFieldOption.Description`, `IssueSingleSelectFieldOption.ID`, `IssueSingleSelectFieldOption.Name`, `ListIssueFields`, `MinimalReviewThread.ID`, `issueFieldsOrgQuery`, `issueFieldsRepoQuery` | 1 | `MinimalCommitSearchItem.Repository`, `searchIssuesNodesQuery.Nodes` | **high**: tests build the unexported githubv4 query structs `issueFieldsRepoQuery{}`/`issueFieldsOrgQuery{}` so the GraphQL mock can match the query shape; the agent's query struct must produce the same GraphQL text | pass |
| 5 | #2480 Add custom field filtering to list_issues | 2026-05-21 | 349 / issue_fields.go, issues.go | 349 | `issueFieldsRepoQuery` | 1 | `IssueField`, `IssueField.DataType`, `IssueField.Description`, `IssueField.Name`, `IssueField.Options` | **high**: same `issueFieldsRepoQuery{}` query struct | pass |

Trivial commits inside the chain: #2515 Fix return Thread node ID in get_review_comments response (3 lines).

Placement: Five tool features across issues.go, search_utils.go, search.go, minimal_types.go, the three GHAS files and a new issue_fields.go; #2480 reuses #2445's IssueField types and query structs.

Caveats: `__toolsnaps__` snapshots ship with the task text (no penalty now). Trivial commit #2515 (3 lines) is a small fix of a review-comment node ID, not a no-op; listed.

### 4. open-policy-agent/opa `v1/ast`

Seed `878cbc79d389`, final `42baad2bc4cb`. Package lines 875, whole PRs 997. Seed test: pass (Go 1.25.5).

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #8427 Add `not` AST node type | 2026-03-30 | 765 / capabilities.go, compare.go, compile.go, mermaid.go, parser.go, policy.go, string_length.go, term.go, term_appenders.go, visit.go | 887 | `Not`, `Not.Body`, `Not.Equal`, `NotExpr`, `mermaidGraph` | 1 | - | **medium**: calls the unexported `mermaidGraph` helper for the mermaid output test | pass (Go 1.25.5) |
| 2 | #8482 Allow back-to-back metadata blocks | 2026-04-02 | 2 / parser.go | 2 | none | 0 | `Not.Location` | **none**: tests use only exported names or names present at the seed | pass (Go 1.25.5) |
| 3 | #8485 type checker: identify compatible type from reference | 2026-04-08 | 26 / check.go | 26 | none | 0 | `Not.Location` | **none**: tests use only exported names or names present at the seed | pass (Go 1.25.5) |
| 4 | #8468 ast: catch functions in dynamic extent of ref head rule | 2026-04-13 | 31 / compile.go | 31 | none | 0 | `Not.IsGround`, `mermaidBuilder.node` | **none**: tests use only exported names or names present at the seed | pass (Go 1.25.5) |
| 5 | #8511 ast: fix parenthesis in String() of {obj,arr,set} comprehensions | 2026-04-14 | 51 / string_length.go, term_appenders.go | 51 | none | 0 | `Not.AppendText`, `Not.Body`, `Not.String`, `Not.StringLength` | **none**: tests use only exported names or names present at the seed | pass (Go 1.25.5) |

Trivial commits inside the chain: #8483 fmt: add new line between METADATA blocks (8 lines), #8463 Add support for days, weeks and years in parse_duration_ns (2 lines).

Placement: #8427 adds a `Not` AST node across 10 files (term, compare, visit, compile, parser, string-length, mermaid); later steps are parser, type-checker, compiler and String() changes that touch the same code paths and use the new node's methods.

Caveats: Weight is front-loaded: #8427 is 765 of 875 lines; the four later steps are small fixes (2-51 lines) of observable parser/checker/compiler/printing behaviour. Trivial commits #8483 (8 lines, fmt newline) and #8463 (2 lines in this package) sit inside.

### 5. jaegertracing/jaeger `internal/storage/v2/api/tracestore`

Seed `35c2a4082608`, final `ea1291f54398`. Package lines 565, whole PRs 851. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #9256 feat(storage): Let backends declare that search can omit the service name | 2026-08-07 | 36 / reader.go | 216 | `Reader.SearchCapabilities`, `SearchCapabilities` | 0 | - | **none**: tests use only exported names or names present at the seed | pass |
| 2 | #9356 feat(storage): Let a reader declare which filters it evaluates | 2026-08-19 | 134 / capabilities.go, reader.go | 134 | `FilterCapabilities`, `FilterCapabilities.IsEmpty`, `FilterCapabilities.Levels`, `FilterCapabilities.Operators`, `FilterCapabilities.SupportsLevel`, `FilterCapabilities.SupportsOperator`, `Reader.SearchCapabilities`, `SearchCapabilities` | 0 | `Reader.SearchCapabilities`, `SearchCapabilities`, `SearchCapabilities.WithoutServiceName` | **none**: tests use only exported names or names present at the seed | pass |
| 3 | #9357 feat(query): Serve a filter query in the shape the backend understands | 2026-08-19 | 395 / admission.go, shape.go | 501 | `ErrFilterInvalid`, `ErrFilterUnsupported`, `FilterCapabilities`, `FilterCapabilities.EnsureSupported`, `FilterCapabilities.Levels`, `FilterCapabilities.Operators`, `Reader.SearchCapabilities`, `SearchCapabilities`, `SearchCapabilities.Filter`, `TraceQueryParams.EnsureFilterStandsAlone`, `TraceQueryParams.Filter`, `TraceQueryParams.ForCapabilities`, `TraceQueryParams.ToFilterShape`, `TraceQueryParams.ToLegacyShape` | 0 | `ErrFilterInvalid`, `ErrFilterUnsupported`, `FilterCapabilities`, `FilterCapabilities.IsEmpty`, `FilterCapabilities.Levels` | **none**: tests use only exported names or names present at the seed | pass |

Trivial commits inside the chain: #9261 fix(memory): Return copies of stored traces, not references (8 lines), #9271 docs(adr): Record the storage capability mechanism and close RFC 0013 (20 lines).

Placement: Storage capability declarations grow step by step (search without service name, then filter capabilities, then query shaping in new admission/shape code that consumes them); the agent decides where capability types and query-shape logic live.

Caveats: 565 package lines, slightly under 600; 66% inside the package (the rest is in query/memory backends). Trivial commits #9261 (8 lines, memory backend returns copies) and #9271 (20 lines of ADR/doc comments) sit inside.

### 6. github/github-mcp-server `pkg/github`

Seed `74964520cf15`, final `d5e1f48728dd`. Package lines 1247, whole PRs 1247. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #683 Add pagination support to GraphQL-based tools | 2025-07-21 | 382 / actions.go, discussions.go, issues.go, notifications.go, pullrequests.go, repositories.go, search.go, search_utils.go, server.go | 382 | `CursorPaginationParams.PerPage`, `PaginationParams.Page`, `PaginationParams.PerPage` | 0 | - | **none**: tests use only exported names or names present at the seed | pass |
| 2 | #470 Add tools for sub-issue endpoint | 2025-07-22 | 407 / issues.go, tools.go | 407 | `AddSubIssue`, `ListSubIssues`, `RemoveSubIssue`, `ReprioritizeSubIssue` | 4 | `CursorPaginationParams.PerPage`, `PaginationParams.Page`, `PaginationParams.PerPage` | **none**: tests use only exported names or names present at the seed | pass |
| 3 | #690 feat: list_discussions sort by updatedAt & createdAt, return updatedAt and author | 2025-07-24 | 306 / discussions.go | 306 | `NodeFragment.Number` | 0 | `CursorPaginationParams.After`, `GraphQLPaginationParams.After`, `GraphQLPaginationParams.First`, `PaginationParams.After` | **none**: tests use only exported names or names present at the seed | pass |
| 4 | #740 Changed q to query in search | 2025-07-24 | 4 / search.go | 4 | none | 1 | none found | **none**: tests use only exported names or names present at the seed | pass |
| 5 | #774 Add updating draft state to `update_pull_request` tool | 2025-07-29 | 148 / pullrequests.go, tools.go | 148 | `BasicNoOrder.Repository`, `BasicWithOrder.Repository`, `NodeFragment.Number`, `NodeFragment.Title`, `WithCategoryAndOrder.Repository`, `WithCategoryNoOrder.Repository` | 1 | `BasicNoOrder.Repository`, `BasicWithOrder.Repository`, `NodeFragment.Number`, `NodeFragment.Title`, `WithCategoryAndOrder.Repository` | **none**: tests use only exported names or names present at the seed | pass |

Trivial commits inside the chain: none.

Placement: #683 adds pagination parameter types across 9 tool files; #470, #690 and #774 use them, so the first step's placement of shared helpers is exercised by every later step.

Caveats: Merged 2025-07 (before the cutoff). #740 (rename of the search parameter `q` to `query`, observable tool schema) is now a step instead of an in-between commit.

### 7. modelcontextprotocol/go-sdk `mcp`

Seed `194378025552`, final `623780125cbb`. Package lines 701, whole PRs 766. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #644 Elicit defaults | 2025-11-21 | 60 / client.go, server.go | 60 | none | 0 | - | **none**: tests use only exported names or names present at the seed | pass |
| 2 | #646 mcp: support url mode elicitation | 2025-11-23 | 188 / client.go, protocol.go, requests.go, server.go, shared.go | 236 | `ClientOptions.ElicitationCompleteHandler`, `ClientOptions.ElicitationModes`, `ElicitParams.ElicitationID`, `ElicitParams.Mode`, `ElicitParams.URL`, `ElicitationCapabilities.URL`, `ElicitationCompleteNotificationRequest`, `ElicitationCompleteParams`, `ElicitationCompleteParams.ElicitationID`, `notificationElicitationComplete` | 0 | none found | **low**: `notificationElicitationComplete` method-name constant | pass |
| 3 | #655 Expose JSON-RPC error codes, and URLElicitationRequired errors | 2025-12-01 | 258 / client.go, resource.go, server.go, shared.go, transport.go | 275 | `ClientOptions.ElicitationModes`, `CodeResourceNotFound`, `CodeURLElicitationRequired`, `ElicitParams.ElicitationID`, `ElicitParams.Mode`, `ElicitParams.URL`, `ElicitationCapabilities.Form`, `ElicitationCapabilities.URL`, `ElicitationCompleteParams`, `ElicitationCompleteParams.ElicitationID`, `FormElicitationCapabilities`, `URLElicitationCapabilities`, `URLElicitationRequiredError`, `notificationElicitationComplete`, `urlElicitationMiddleware` | 0 | `ClientOptions.ElicitationModes`, `ElicitParams.ElicitationID`, `ElicitParams.Mode`, `ElicitationCapabilities.Form`, `ElicitationCompleteParams` | **medium**: installs `urlElicitationMiddleware()` in a test client | pass |
| 4 | #665 fix: handle Windows CRLF in MCP client(#664) | 2025-12-01 | 3 / transport.go | 3 | none | 0 | none found | **none**: tests use only exported names or names present at the seed | pass |
| 5 | #663 mcp: support SSE polling and server disconnects (SEP-1699) | 2025-12-01 | 192 / event.go, shared.go, streamable.go | 192 | `ElicitParams.URL`, `ElicitationCapabilities.URL`, `RequestExtra.CloseStream` | 0 | none found | **none**: tests use only exported names or names present at the seed | pass |

Trivial commits inside the chain: none.

Placement: Elicitation defaults, URL-mode elicitation, error codes / URLElicitationRequired built on it, a CRLF fix, then SSE polling; spread over client.go, server.go, shared.go, protocol.go, transport.go, streamable.go, event.go.

Caveats: Merged 2025-11/12. #665 is a 3-line observable fix (Windows CRLF in SSE parsing).

### 8. go-git/go-git `plumbing/protocol/packp`

Seed `c8de2fac7415`, final `75f92acd2a83`. Package lines 1606, whole PRs 2357. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #2225 plumbing: packp/transport, Slice-based AdvRefs and v1 handling (2/11) | 2026-06-26 | 39 / advrefs.go, advrefs_decode.go, advrefs_encode.go | 46 | `AdvRefs.Version` | 0 | - | **none**: tests use only exported names or names present at the seed | pass |
| 2 | #2227 plumbing: packp, Protocol v2 wire types with caller-owned packfile streaming (4/11) | 2026-06-26 | 1223 / capability_adv.go, command.go, fetch.go, list.go, lsrefs.go | 1237 | `Acknowledgments`, `Acknowledgments.ACKs`, `Acknowledgments.Ready`, `AdvRefs.Version`, `CapabilityAdv`, `CapabilityAdv.Capabilities`, `CapabilityAdv.Decode`, `CapabilityAdv.Encode`, `CapabilityAdv.Version`, `CommandRequest`, `CommandRequest.Args`, `CommandRequest.Capabilities`, `CommandRequest.Command`, `CommandRequest.Decode`, `CommandRequest.Encode`, `FetchArgs`, `FetchArgs.Decode`, `FetchArgs.Deepen`, `FetchArgs.DeepenNot`, `FetchArgs.DeepenRelative`, `FetchArgs.DeepenSince`, `FetchArgs.Done`, `FetchArgs.Encode`, `FetchArgs.Filter`, `FetchArgs.Haves`, `FetchArgs.IncludeTag`, `FetchArgs.NoProgress`, `FetchArgs.OFSDelta`, `FetchArgs.Shallows`, `FetchArgs.ThinPack`, `FetchArgs.WaitForDone`, `FetchArgs.Wants`, `FetchOutput`, `FetchOutput.Acknowledgments`, `FetchOutput.Decode`, `FetchOutput.Encode`, `FetchOutput.Packfile`, `FetchOutput.PackfileURIs`, `FetchOutput.ShallowInfo`, `FetchOutput.WantedRefs`, `LsRefsArgs`, `LsRefsArgs.Decode`, `LsRefsArgs.Encode`, `LsRefsArgs.Peel`, `LsRefsArgs.RefPrefixes`, `LsRefsArgs.Symrefs`, `LsRefsArgs.Unborn`, `LsRefsOutput`, `LsRefsOutput.Decode`, `LsRefsOutput.Encode`, `LsRefsOutput.References`, `PackfileURIs`, `PackfileURIs.URIs`, `ShallowInfo`, `ShallowInfo.Shallows`, `ShallowInfo.Unshallows`, `WantedRefs`, `WantedRefs.Refs` | 0 | `AdvRefs.Version` | **none**: tests use only exported names or names present at the seed | pass |
| 3 | #2237 plumbing: default to wire protocol v2 and close client/server conformance gaps | 2026-07-06 | 287 / advrefs.go, fetch.go, lsrefs.go | 1003 | `CapabilityAdv`, `CapabilityAdv.Decode`, `CapabilityAdv.Encode`, `CommandRequest`, `CommandRequest.Args`, `CommandRequest.Decode`, `CommandRequest.Encode`, `FetchArgs`, `FetchArgs.Decode`, `FetchArgs.Encode`, `FetchOutput`, `FetchOutput.Decode`, `FetchOutput.Encode`, `FetchOutput.Packfile`, `LsRefsArgs`, `LsRefsArgs.Decode`, `LsRefsArgs.Encode`, `LsRefsArgs.Peel`, `LsRefsArgs.RefPrefixes`, `LsRefsArgs.Symrefs`, `LsRefsOutput`, `LsRefsOutput.Decode`, `LsRefsOutput.Encode`, `MalformedResponseError`, `MalformedResponseError.Error`, `maxSectionLines`, `parseLsRefsLine`, `tooManyRefPrefixes` | 0 | `Acknowledgments`, `Acknowledgments.ACKs`, `Acknowledgments.Ready`, `CapabilityAdv`, `CapabilityAdv.Decode` | **high**: unit-tests the small line parser `parseLsRefsLine` and the limits `maxSectionLines`/`tooManyRefPrefixes`; no counterpart if the agent does not extract a line parser | pass |
| 4 | #2263 plumbing: protocol/packp, reject control bytes in daemon request | 2026-07-23 | 46 / gitproto.go | 46 | `CapabilityAdv.Decode`, `CapabilityAdv.Encode`, `CommandRequest.Decode`, `CommandRequest.Encode`, `FetchArgs.Decode`, `FetchArgs.Encode`, `FetchOutput.Decode`, `FetchOutput.Encode`, `LsRefsArgs.Decode`, `LsRefsArgs.Encode`, `LsRefsOutput.Decode`, `LsRefsOutput.Encode` | 0 | `CapabilityAdv.Decode`, `CapabilityAdv.Encode`, `CommandRequest.Decode`, `CommandRequest.Encode`, `FetchArgs.Decode` | **none**: tests use only exported names or names present at the seed | pass |
| 5 | #2264 plumbing: protocol/packp, compare zero object ids by value in Action | 2026-07-23 | 11 / updreq.go | 25 | `CommandRequest.Command`, `MalformedResponseError.Error` | 0 | `CommandRequest.Command`, `MalformedResponseError.Error` | **none**: tests use only exported names or names present at the seed | pass |

Trivial commits inside the chain: none.

Placement: Protocol v2 support built in stages: slice-based AdvRefs, new v2 wire types (capability_adv, command, fetch, ls-refs), switching the default, then input-validation fixes on the new types.

Caveats: #2227 alone is 1223 of 1606 lines (one large step); 68% inside the package. #2264 (compare zero object IDs by value) is borderline: it may only matter for upstream's representation; read before use.

### 9. traefik/traefik `pkg/provider/kubernetes/gateway`

Seed `90a7b2b411dd`, final `074d02c09fae`. Package lines 597, whole PRs 970. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #14130 Support HTTPRoute backendRequest timeout | 2026-10-23 | 28 / features.go, httproute.go | 28 | none | 0 | - | **none**: tests use only exported names or names present at the seed | pass |
| 2 | #14119 Support Gateway backend client certificate | 2026-10-26 | 152 / client.go, features.go, grpcroute.go, httproute.go, kubernetes.go, tlsroute.go | 152 | `Provider.loadGatewayBackendTLS`, `gatewayWithListeners.errResolvedRefsCondition` | 0 | none found | **high**: calls `Provider.loadGatewayBackendTLS` and sets the internal field `gatewayWithListeners.errResolvedRefsCondition` | pass |
| 3 | #14127 Support HTTPRoute retry | 2026-10-27 | 54 / features.go, httproute.go, kubernetes.go | 54 | `Provider.RetryMaxRequestBodyBytes` | 0 | none found | **none**: tests use only exported names or names present at the seed | pass |
| 4 | #14123 Support HTTPRoute RequestMirror filter | 2026-10-28 | 76 / features.go, httproute.go | 447 | `Provider.loadRequestMirror` | 0 | none found | **medium**: calls `Provider.loadRequestMirror(...)` with 10 positional arguments; the agent's builder likely differs in shape | pass |
| 5 | #13410 Support ClusterTrustBundle CA certificate references | 2026-10-28 | 287 / client.go, features.go, httproute.go, kubernetes.go, tlsroute.go | 289 | none | 0 | none found | **none**: tests use only exported names or names present at the seed | pass |

Trivial commits inside the chain: #14125 Bump sigs.k8s.io/gateway-api to v1.7.0-rc.3 (12 lines), #14145 Merge branch v3.7 into master (2 lines).

Placement: Five Gateway API features that each touch features.go and httproute.go plus some of kubernetes.go/client.go/grpcroute.go/tlsroute.go; #14119 and #13410 share TLS/CA loading.

Caveats: Building on earlier steps is only through shared files (no later step's test or code uses a name an earlier step added), 597 package lines, 62% inside. Fixture YAML ships with the hidden tests. Trivial commits #14125 (12-line dependency bump) and #14145 (2-line merge).

### 10. slack-go/slack `.`

Seed `4c472cd10a45`, final `af28090b40b6`. Package lines 734, whole PRs 734. Seed test: pass except TestRTMDisconnect (dials slack.com).

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #1551 feat(block): add card and carousel blocks | 2026-04-17 | 251 / block.go, block_card.go, block_carousel.go, block_conv.go, chat.go, chat_stream_chunks.go | 251 | `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock`, `CardBlock.Actions`, `CardBlock.BlockType`, `CardBlock.Body`, `CardBlock.HeroImage`, `CardBlock.ID`, `CardBlock.Icon`, `CardBlock.Subtitle`, `CardBlock.Title`, `CardBlock.Type`, `CardBlock.WithActions`, `CardBlock.WithBody`, `CardBlock.WithHeroImage`, `CardBlock.WithIcon`, `CardBlock.WithSubtitle`, `CardBlock.WithTitle`, `CardBlockOptionBlockID`, `CarouselBlock`, `CarouselBlock.AddCard`, `CarouselBlock.BlockType`, `CarouselBlock.Elements`, `CarouselBlock.ID`, `CarouselBlock.Type`, `CarouselBlock.WithBlockID`, `MBTCard`, `MBTCarousel`, `MarkdownTextChunk.Text`, `MarkdownTextChunk.Type`, `NewCardBlock`, `NewCarouselBlock`, `PlanUpdateChunk.Title`, `PlanUpdateChunk.Type`, `TaskUpdateChunk.ID`, `TaskUpdateChunk.Title`, `TaskUpdateChunk.Type` | 0 | - | **none**: tests use only exported names or names present at the seed | pass except TestRTMDisconnect (dials slack.com) |
| 2 | #1553 feat(assistant): add username and icon to status update | 2026-04-20 | 15 / assistant.go | 15 | `AssistantThreadsSetStatusParameters.IconEmoji`, `AssistantThreadsSetStatusParameters.IconURL`, `AssistantThreadsSetStatusParameters.Username` | 0 | none found | **none**: tests use only exported names or names present at the seed | pass except TestRTMDisconnect (dials slack.com) |
| 3 | #1552 feat(block): add alert block | 2026-04-22 | 73 / block.go, block_alert.go, block_conv.go | 73 | `AlertBlock`, `AlertBlock.BlockType`, `AlertBlock.ID`, `AlertBlock.Level`, `AlertBlock.Text`, `AlertBlock.Type`, `AlertBlockOptionBlockID`, `AlertBlockOptionLevel`, `AlertLevelInfo`, `AlertLevelWarning`, `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock.BlockType`, `CardBlock.ID`, `CardBlock.Type`, `CarouselBlock.BlockType`, `CarouselBlock.ID`, `CarouselBlock.Type`, `MBTAlert`, `MarkdownTextChunk.Text`, `MarkdownTextChunk.Type`, `NewAlertBlock`, `PlanUpdateChunk.Type`, `TaskUpdateChunk.ID`, `TaskUpdateChunk.Type` | 0 | `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock.BlockID`, `CardBlock.BlockType`, `CardBlock.Type` | **none**: tests use only exported names or names present at the seed | pass except TestRTMDisconnect (dials slack.com) |
| 4 | #1555 feat: add new block kit block Data Table | 2026-05-24 | 223 / block.go, block_conv.go, block_data_table.go, block_plan.go, block_task_card.go | 223 | `AlertBlock.BlockType`, `AlertBlock.ID`, `AlertBlock.Text`, `AlertBlock.Type`, `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock.BlockType`, `CardBlock.ID`, `CardBlock.Type`, `CarouselBlock.BlockType`, `CarouselBlock.Elements`, `CarouselBlock.ID`, `CarouselBlock.Type`, `DataTableBlock`, `DataTableBlock.AddRow`, `DataTableBlock.BlockType`, `DataTableBlock.Caption`, `DataTableBlock.ID`, `DataTableBlock.PageSize`, `DataTableBlock.RowHeaderColumnIndex`, `DataTableBlock.Rows`, `DataTableBlock.Type`, `DataTableBlock.WithPageSize`, `DataTableBlock.WithRowHeaderColumnIndex`, `DataTableBlockOptionBlockID`, `DataTableCell.DataTableCellType`, `DataTableCellRawNumber`, `DataTableCellRawText`, `DataTableCellType`, `DataTableRawNumberCell`, `DataTableRawNumberCell.DataTableCellType`, `DataTableRawNumberCell.Text`, `DataTableRawNumberCell.Type`, `DataTableRawNumberCell.Value`, `DataTableRawNumberCell.WithText`, `DataTableRawTextCell`, `DataTableRawTextCell.DataTableCellType`, `DataTableRawTextCell.Text`, `DataTableRawTextCell.Type`, `DataTableRichTextCell`, `DataTableRichTextCell.DataTableCellType`, `DataTableRichTextCell.Elements`, `DataTableRichTextCell.Type`, `MBTDataTable`, `MarkdownTextChunk.Text`, `MarkdownTextChunk.Type`, `NewDataTableBlock`, `NewDataTableRawNumberCell`, `NewDataTableRawTextCell`, `PlanUpdateChunk.Type`, `TaskUpdateChunk.ID`, `TaskUpdateChunk.Type` | 0 | `AlertBlock.BlockID`, `AlertBlock.BlockType`, `AlertBlock.Text`, `AlertBlock.Type`, `BlocksChunk.Blocks` | **none**: tests use only exported names or names present at the seed | pass except TestRTMDisconnect (dials slack.com) |
| 5 | #1560 fix!: preserve all cell types in `TableBlock` | 2026-06-04 | 172 / block_table.go | 172 | `AlertBlock.BlockType`, `AlertBlock.Text`, `AlertBlock.Type`, `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock.BlockType`, `CardBlock.Type`, `CarouselBlock.BlockType`, `CarouselBlock.Elements`, `CarouselBlock.Type`, `DataTableBlock.AddRow`, `DataTableBlock.BlockType`, `DataTableBlock.Rows`, `DataTableBlock.Type`, `DataTableRawNumberCell.Text`, `DataTableRawNumberCell.Type`, `DataTableRawNumberCell.Value`, `DataTableRawTextCell.Text`, `DataTableRawTextCell.Type`, `DataTableRichTextCell.Elements`, `DataTableRichTextCell.Type`, `MarkdownTextChunk.Text`, `MarkdownTextChunk.Type`, `NewTableRichTextCell`, `PlanUpdateChunk.Type`, `TableCell.TableCellType`, `TableCellRichText`, `TableCellType`, `TableRawNumberCell`, `TableRawNumberCell.TableCellType`, `TableRawNumberCell.Text`, `TableRawNumberCell.Type`, `TableRawNumberCell.Value`, `TableRawTextCell`, `TableRawTextCell.TableCellType`, `TableRawTextCell.Text`, `TableRawTextCell.Type`, `TableRichTextCell`, `TableRichTextCell.Elements`, `TableRichTextCell.TableCellType`, `TableRichTextCell.Type`, `TaskUpdateChunk.Type` | 0 | `AlertBlock.BlockID`, `AlertBlock.BlockType`, `AlertBlock.Text`, `AlertBlock.Type`, `BlocksChunk.Blocks` | **none**: tests use only exported names or names present at the seed | pass except TestRTMDisconnect (dials slack.com) |

Trivial commits inside the chain: direct commit security: reject empty signing secret for NewSecretsVerifier (19 lines).

Placement: New Block Kit blocks registered in block.go/block_conv.go plus new files; #1552/#1555/#1560 reuse the chunk/card/table types.

Caveats: About 100 exported names dictated by Slack's API (low rename risk but long task text); mostly one new file per step. A direct (non-PR) 19-line security commit sits inside. TestRTMDisconnect dials slack.com.

### 11. gofiber/fiber `middleware/cache`

Seed `0923873f4a1b`, final `cec4e4dd4d7f`. Package lines 646, whole PRs 695. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #4735 🐛 bug: accept tab OWS in X-Forwarded-For and share one OWS trim helper | 2026-10-03 | 53 / cachecontrol.go | 83 | none | 0 | - | **none**: tests use only exported names or names present at the seed | pass |
| 2 | #4736 🐛 bug: key cache entries by request scheme and host | 2026-10-05 | 35 / cache.go, config.go, keygen.go | 39 | none | 0 | none found | **none**: tests use only exported names or names present at the seed | pass |
| 3 | #4744 🐛 bug: replay Vary on cache hits and mark credential-keyed entries private | 2026-10-08 | 139 / cache.go, config.go, manager.go, manager_msgp.go, utils.go, vary.go | 139 | `item.vary`, `maxVaryLen`, `varyListsCredential` | 0 | none found | **high**: reads `item.vary` and calls `varyListsCredential` | pass |
| 4 | #4590 🐛 bug: Delete stale entries after uncacheable revalidation | 2026-10-12 | 419 / cache.go, heap.go, manager.go, manager_msgp.go | 434 | `indexedHeap.generation`, `indexedHeap.matches`, `item.heapgen`, `manager.delIf`, `memoryEntryAbsent`, `memoryEntryDeleted`, `memoryEntryReplaced` | 0 | `item.vary` | **high**: reads `item.heapgen`, calls `indexedHeap.generation/matches` and `manager.delIf`; deep internal state of the cache store | pass |

Trivial commits inside the chain: none.

Placement: Four observable HTTP-cache fixes (OWS parsing, cache key by scheme/host, Vary replay, stale-entry deletion) that touch cache.go, manager.go, config.go and add vary.go/heap.go logic; later steps modify the cache entry structure earlier steps extended.

Caveats: All steps are fixes, no feature. Tests inspect cache entry internals (high rename risk).

### 12. ollama/ollama `cmd/launch`

Seed `455f57457d26`, final `07ed752353df`. Package lines 698, whole PRs 698. Seed test: pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken).

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/fixture files | Uses earlier steps' code | Rename risk | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #16516 launch: hermes-desktop app | 2026-06-04 | 156 / hermes.go, launch.go, registry.go | 156 | `HermesDesktop`, `HermesDesktop.Run`, `HermesDesktop.launchArgs`, `hermesDesktopReleaseRoots` | 0 | - | **medium**: calls `HermesDesktop.launchArgs` and `hermesDesktopReleaseRoots` | pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken) |
| 2 | #16410 integrations: oh-my-pi | 2026-06-05 | 473 / launch.go, omp.go, pi.go, registry.go | 473 | `HermesDesktop.Run`, `HermesDesktop.String`, `OMP`, `OMP.ConfigureWithModels`, `OMP.CurrentModel`, `OMP.Paths`, `OMP.RequiresInteractiveOnboarding`, `OMP.Run`, `OMP.String`, `OMP.args`, `OMP.findPath`, `ompExecutableNames`, `ompSetupVersion`, `ompWebSearchPlugin` | 0 | `HermesDesktop.Onboard`, `HermesDesktop.String` | **high**: calls small helpers `OMP.args`, `OMP.findPath` and reads `ompExecutableNames`/`ompSetupVersion`; likely no counterpart if the agent structures the launcher differently | pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken) |
| 3 | #16558 launch: use native Windows Hermes config path | 2026-06-06 | 29 / hermes.go | 29 | `OMP.Paths` | 0 | `OMP.Paths` | **none**: tests use only exported names or names present at the seed | pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken) |
| 4 | #16683 launch: Fix launch provider drift | 2026-06-12 | 3 / launch.go | 3 | none | 0 | none found | **none**: tests use only exported names or names present at the seed | pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken) |
| 5 | #15434 launch: add thinking capability detection to opencode | 2026-06-18 | 37 / opencode.go | 37 | `HermesDesktop.Run`, `OMP.Run` | 0 | none found | **none**: tests use only exported names or names present at the seed | pass except TestShowInModelPicker (go-sqlite3 needs cgo; local cgo linker broken) |

Trivial commits inside the chain: none.

Placement: Integration launchers (Hermes desktop, oh-my-pi, opencode) in a 28-file flat package; registry.go and launch.go are shared, and #16558 changes #16516's Hermes code.

Caveats: #16683 is a 3-line fix; integrations are largely one file each, so placement choices are moderate. Tests call small unexported helpers (args/path builders).

## Notable rejected chains

| Repo | Package | Chain | Failed criterion |
| --- | --- | --- | --- |
| jesseduffield/lazygit | `pkg/config` | #5563 → #5565 → #5571 → #5634 (v2 rank 8) | Dropped: only 36% of the chain's non-test lines are inside the package (rule: 45%+) |
| mark3labs/mcp-go | `server` | #841 → #843 → #844 (v2 rank 7) | Dropped: no later step uses code an earlier step added (three independent transport features) |
| prometheus/prometheus | `promql/parser` | #19821 → #19683 → #19820 → #19827 → #20169 (2026-10) | 1417 of 1611 lines are the generated parser (generated_parser.y.go) in one step |
| prometheus/prometheus | `tsdb, web/api/v1` | #19012 → #19013 → #19016 → #18905; #19910 → ... (2026) | Rename risk 16-20 (white-box tests on internal state); web/api steps do not build on each other |
| temporalio/sdk-go | `internal, converter` | #2522 → #2609 → #2602 → #2634 → #2617; #2703 → #2793 → #2763 (2026) | Steps are 3-4 line fixes plus an 'expose internal APIs' refactor; converter chain is 53% inside and ends in a refactor |
| open-policy-agent/opa | `v1/ast` | #9389 → #9414 → #9062 → #9377; #8298 → ... → #8310 (2026) | Third/fourth opa chain (limit 2 per repo, only one kept); #9377's tests read RegoVersionSpan.min/max; #8304-#8310 test compiler stage internals |
| nats-io/nats-server | `server` | #8540 → #8555 → #8556 and others (2026) | Package tests take far too long to run at every step; rename risk 14-55 |
| cockroachdb/pebble | `.` | #5497 → #5504 → #5521 and others | Does not build with Go 1.27 (cockroachdb/swiss) |
| danielgtaylor/huma | `.` | #876 → #937 → #947 → #951 → #961 (2026-01/02) | Last step is a 12-file 'Fix various issues' grab-bag; little building on earlier steps |
| juicedata/juicefs | `pkg/meta` | #6656 → #6695 → #6698 (2026-02/03) | Step 2 is a refactor; step 3 fixes tests |
| github/github-mcp-server | `pkg/github` | #988 → #1023 → #1062 → #1073 → #1078 (2025-09); #1538 → #2694 → #2601 → #1974 → #2695 (2026-06) | Good but over the 2-per-repository limit; ranked below the two kept github-mcp-server chains |
| quic-go/quic-go | `.` | #5960 → #5961 → #5969 → #5970 → #5976 (2026-10) | Rename risk 18-51; tests drive framer/packer internals |
| autobrr/autobrr | `internal/irc, internal/http` | #2763 → #2765 → #2767; #2530 → #2791 → #2809 | Rename risk 18-74 (tests on handler/service internals) |
| caddyserver/caddy | `reverseproxy, caddyfile` | #7649 → #8135 → #8139 → #8160 → #8159; #8047 → #8163 → #7880 → #8203 | Rename risk 35-64 |
| traefik/traefik | `pkg/provider/kubernetes/gateway` | #13951 → #13960 → #13965 → #13966 → #12909 (2026-09) | Three of five steps are 'Merge branch v3.7' commits; rename risk 50 |

## Screening (v3)

- Reused the 61 survey outputs and the v2 commit data; re-mined 57 module roots with the v3 member rule: a merged PR whose title is not lint/bump/chore/docs/CI/refactor and that changes a test file and at least 1 non-test line in the package (small fixes now count). A commit of 20 lines or fewer may still sit between members; anything else breaks the chain.
- 372 packages; 95 have at least one window (v2: 88). 1456 windows of 3-5 steps with 450-2600 package lines and 4+ files (v2: 1001); 784 after requiring file reuse between steps and 45%+ of lines inside the package (v2: 539).
- Name and rename-risk check on all 784: 269 had no later step using an earlier step's new identifiers; 487 had at least one step whose tests read unexported fields the chain introduced; 546 scored 10+ on the rename-risk heuristic (3 per field read, 2 per helper call, 1 per other unexported name). 108 windows in 17 packages had name-level building, risk under 10 and 600-2000 lines.
- About 45 windows read by hand (titles, files, fix nature, test lines); 12 kept with at most 2 per repository. Fix steps were judged from PR titles/bodies and diffs for observable behaviour.
- Offline tests: 49 runs from v2 reused, 31 new runs for the new chains/steps.
