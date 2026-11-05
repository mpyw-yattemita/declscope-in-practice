# Candidate PR chains

Generated 2026-11-05. Each chain is 3-5 consecutive merged PRs on one package: no other commit touched the package's own `.go` files between members except the trivial commits listed. Lines are non-test `.go` lines (added+removed) in the package; `whole` counts every non-test `.go` file in the PRs. Names to specify are the identifiers that a step's added test lines use and that some chain step (up to that one) introduced, checked by token matching and, for unexported ones, by reading the test lines. Offline tests: `go test -count=1 ./<pkg>` under `sandbox-exec` with outbound network denied except localhost, `GOPROXY=off`, `CGO_ENABLED=0`, Go 1.27.0 darwin/arm64, at the seed and at every step's merge commit (so each step's own tests ran where they were merged).

## Ranking

| # | Repo | Package (files/ns) | Steps | Merged | Pkg lines (whole) | Pkg files | Names to specify (unexported) | Offline: seed / steps | Why it gives placement decisions |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | hashicorp/hcl | `hclwrite` (17/18) | [#841](https://github.com/hashicorp/hcl/pull/841) → [#811](https://github.com/hashicorp/hcl/pull/811) → [#820](https://github.com/hashicorp/hcl/pull/820) → [#835](https://github.com/hashicorp/hcl/pull/835) | 2026-10-13 .. 2026-10-19 | 788 (788) | 10 | 7 (1) | pass / pass; pass; pass; pass | Each step adds a new expression kind or API that must be wired through parser.go, ast_expression.go and the node/unwrap plumbing; the agent chooses between new files (ast_object_cons_expr.go, ast_function_call_expr.go) and existing ones, and #820/#835 build on #811's unwrap/variables code. |
| 2 | modelcontextprotocol/go-sdk | `mcp` (25/34) | [#975](https://github.com/modelcontextprotocol/go-sdk/pull/975) → [#989](https://github.com/modelcontextprotocol/go-sdk/pull/989) → [#997](https://github.com/modelcontextprotocol/go-sdk/pull/997) → [#987](https://github.com/modelcontextprotocol/go-sdk/pull/987) | 2026-06-02 .. 2026-06-08 | 617 (646) | 7 | 13 (2) | pass / pass; pass; pass; pass | One protocol feature (SEP-2575 discover / version negotiation) split into client side, retry, logging level and server side; every step touches 3-6 of client.go/server.go/shared.go/protocol.go/streamable.go, and later steps call the earlier steps' Client.discover/usesNewProtocol. |
| 3 | github/github-mcp-server | `pkg/github` (77/113) | [#2474](https://github.com/github/github-mcp-server/pull/2474) → [#2451](https://github.com/github/github-mcp-server/pull/2451) → [#2284](https://github.com/github/github-mcp-server/pull/2284) → [#2445](https://github.com/github/github-mcp-server/pull/2445) → [#2480](https://github.com/github/github-mcp-server/pull/2480) | 2026-05-20 .. 2026-05-21 | 1233 (1233) | 9 | 30 (2) | pass / pass; pass; pass; pass; pass | Five tool features across issues.go, search_utils.go, search.go, minimal_types.go, the three GHAS files and a new issue_fields.go; #2480 reuses #2445's IssueField types and GraphQL query structs, so where #2445 put them matters. |
| 4 | traefik/traefik | `pkg/provider/kubernetes/gateway` (9/9) | [#14130](https://github.com/traefik/traefik/pull/14130) → [#14119](https://github.com/traefik/traefik/pull/14119) → [#14127](https://github.com/traefik/traefik/pull/14127) → [#14123](https://github.com/traefik/traefik/pull/14123) → [#13410](https://github.com/traefik/traefik/pull/13410) | 2026-10-23 .. 2026-10-28 | 597 (970) | 6 | 4 (3) | pass / pass; pass; pass; pass; pass | Five Gateway API features that each touch features.go, httproute.go and some of kubernetes.go/client.go/grpcroute.go/tlsroute.go; backend TLS (#14119) and ClusterTrustBundle CA refs (#13410) share client/TLS loading code. |
| 5 | github/github-mcp-server | `pkg/github` (77/113) | [#683](https://github.com/github/github-mcp-server/pull/683) → [#470](https://github.com/github/github-mcp-server/pull/470) → [#690](https://github.com/github/github-mcp-server/pull/690) → [#774](https://github.com/github/github-mcp-server/pull/774) | 2025-07-21 .. 2025-07-29 | 1243 (1243) | 10 | 13 (0) | pass / pass; pass; pass; pass | #683 adds pagination parameter types used across 9 tool files; #470 (sub-issues), #690 (discussions sort) and #774 (draft PR state) then use those params, so the first step's placement of shared helpers is exercised by all later steps. |
| 6 | modelcontextprotocol/go-sdk | `mcp` (25/34) | [#646](https://github.com/modelcontextprotocol/go-sdk/pull/646) → [#655](https://github.com/modelcontextprotocol/go-sdk/pull/655) → [#663](https://github.com/modelcontextprotocol/go-sdk/pull/663) | 2025-11-23 .. 2025-12-01 | 638 (703) | 9 | 18 (2) | pass / pass; pass; pass | URL-mode elicitation, then error codes / URLElicitationRequired errors built on it, then SSE polling; spread over client.go, server.go, shared.go, protocol.go, transport.go, streamable.go, event.go. |
| 7 | mark3labs/mcp-go | `server` (37/90) | [#841](https://github.com/mark3labs/mcp-go/pull/841) → [#843](https://github.com/mark3labs/mcp-go/pull/843) → [#844](https://github.com/mark3labs/mcp-go/pull/844) | 2026-05-04 .. 2026-05-04 | 973 (973) | 7 | 34 (5) | pass / pass; pass; pass | Three HTTP-transport features (protected-resource metadata, CORS, output-schema validation) that each need hooks in sse.go and streamable_http.go plus a new file; placement of shared HTTP middleware is a real choice. |
| 8 | jesseduffield/lazygit | `pkg/config` (15/13) | [#5563](https://github.com/jesseduffield/lazygit/pull/5563) → [#5565](https://github.com/jesseduffield/lazygit/pull/5565) → [#5571](https://github.com/jesseduffield/lazygit/pull/5571) → [#5634](https://github.com/jesseduffield/lazygit/pull/5634) | 2026-04-30 .. 2026-05-25 | 1206 (3375) | 6 | 4 (0) | pass / pass; pass; pass; pass | Keybinding revamp: #5563 introduces key names/labels, #5565 changes canonical syntax using them, #5634 lets a binding hold several keys; code spreads over keynames.go, user_config.go, app_config.go, keybinding.go, user_config_validation.go. |
| 9 | slack-go/slack | `.` (125/149) | [#1551](https://github.com/slack-go/slack/pull/1551) → [#1553](https://github.com/slack-go/slack/pull/1553) → [#1552](https://github.com/slack-go/slack/pull/1552) → [#1555](https://github.com/slack-go/slack/pull/1555) | 2026-04-17 .. 2026-05-24 | 562 (562) | 11 | 86 (0) | pass except TestRTMDisconnect (dials slack.com) / pass except TestRTMDisconnect (dials slack.com); pass except TestRTMDisconnect (dials slack.com); pass except TestRTMDisconnect (dials slack.com); pass except TestRTMDisconnect (dials slack.com) | New Block Kit blocks registered in block.go/block_conv.go plus a new file each; #1552/#1555 reuse the chunk/card types from #1551. |

## Chains in detail

### 1. hashicorp/hcl `hclwrite`

Seed `a786e4fc5566`, final `4a347c448ee9`. Package lines 788, whole PRs 788. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #841 hclwrite: find traversals inside object-construct expressions | 2026-10-13 | 25 / ast_expression.go, ast_object_cons_expr.go, parser.go | 25 | none | 0 | - | pass |
| 2 | #811 hclwrite: add read and write API for `ObjectConsExpr` | 2026-10-14 | 452 / ast.go, ast_expression.go, ast_object_cons_expr.go, node.go, parser.go, unwrap.go | 452 | `ObjectConsExpr.ItemFor`, `ObjectConsExpr.SetItem`, `ObjectConsExpr.SetItemRaw`, `ObjectConsExpr.SetItemTraversal`, `ObjectConsExpr.SetItemValue`, `ObjectConsExpr.ValueExprFor` | 0 | none found | pass |
| 3 | #820 hclwrite: parse function call and tuple-construct expressions | 2026-10-14 | 221 / ast_expression.go, ast_function_call_expr.go, ast_tuple_cons_expr.go, parser.go | 221 | none | 0 | `Expression.unwrap`, `ObjectConsExpr.variables`, `ObjectConsKeyExpr.unwrap`, `unwrapNode.unwrap` | pass |
| 4 | #835 hclwrite: New function `NewExpression()` and `Body.SetAttribute()` | 2026-10-19 | 90 / ast_body.go, parser.go, public.go | 90 | `parseExpr` | 0 | `FunctionCallExpr.name` | pass |

Trivial commits inside the chain: none.

Placement: Each step adds a new expression kind or API that must be wired through parser.go, ast_expression.go and the node/unwrap plumbing; the agent chooses between new files (ast_object_cons_expr.go, ast_function_call_expr.go) and existing ones, and #820/#835 build on #811's unwrap/variables code.

Caveats: #841 is small (25 lines) but is the seed of the object-construct work; the coordinator's model chain started at #810, but #824 and #841 sit between #810 and #811, so #810 cannot be in a strictly consecutive chain unless #824 (51-line Body feature) is included as a step. #835's test calls unexported `parseExpr(src []byte, filename string, start hcl.Pos)`.

### 2. modelcontextprotocol/go-sdk `mcp`

Seed `d54751a8eec7`, final `dd978160abb3`. Package lines 617, whole PRs 646. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #975 mcp: Implement support for SEP-2575 on client side | 2026-06-02 | 227 / client.go, protocol.go, requests.go, server.go, shared.go, streamable.go | 233 | `CodeUnsupportedProtocolVersion`, `DiscoverParams`, `DiscoverResult`, `DiscoverResult.Capabilities`, `DiscoverResult.Instructions`, `DiscoverResult.ServerInfo`, `DiscoverResult.SupportedVersions`, `methodDiscover` | 0 | - | pass |
| 2 | #989 mcp: Implement retry on advised supportedVersions in UnsupportedProtocolVersion error | 2026-06-08 | 94 / client.go, protocol.go, shared.go | 94 | `CodeUnsupportedProtocolVersion`, `DiscoverParams`, `DiscoverResult`, `DiscoverResult.Capabilities`, `DiscoverResult.ServerInfo`, `DiscoverResult.SupportedVersions`, `UnsupportedProtocolVersionData`, `UnsupportedProtocolVersionData.Requested`, `UnsupportedProtocolVersionData.Supported`, `methodDiscover` | 0 | `Client.discover`, `CodeUnsupportedProtocolVersion`, `DiscoverParams`, `DiscoverResult`, `DiscoverResult.Capabilities` | pass |
| 3 | #997 mcp: add support for logging level (SEP-2575) | 2026-06-08 | 44 / protocol.go, server.go, shared.go, streamable.go | 44 | `MetaKeyLogLevel`, `validatedMeta.logLevel` | 0 | `ClientSession.usesNewProtocol` | pass |
| 4 | #987 mcp: Implement server-side support for discover method (SEP-2575) | 2026-06-08 | 252 / server.go, shared.go, streamable.go, transport.go | 275 | `CodeUnsupportedProtocolVersion`, `DiscoverResult`, `DiscoverResult.Capabilities`, `DiscoverResult.ServerInfo`, `DiscoverResult.SupportedVersions`, `methodDiscover` | 0 | `Client.discover`, `ClientSession.usesNewProtocol`, `CodeUnsupportedProtocolVersion`, `DiscoverParams`, `DiscoverResult` | pass |

Trivial commits inside the chain: none.

Placement: One protocol feature (SEP-2575 discover / version negotiation) split into client side, retry, logging level and server side; every step touches 3-6 of client.go/server.go/shared.go/protocol.go/streamable.go, and later steps call the earlier steps' Client.discover/usesNewProtocol.

Caveats: Tests are partly white-box: they switch on the unexported `methodDiscover` const and read `validatedMeta.logLevel`. #975 and #987 also change internal/jsonrpc2 (outside the package, small).

### 3. github/github-mcp-server `pkg/github`

Seed `e9533330b0ec`, final `13e7f10fc8dd`. Package lines 1233, whole PRs 1233. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #2474 Add field_values to search_issues results | 2026-05-20 | 215 / issues.go, search_utils.go | 215 | `SearchIssueResult.FieldValues`, `SearchIssuesResponse`, `SearchIssuesResponse.IncompleteResults`, `SearchIssuesResponse.Items`, `SearchIssuesResponse.Total` | 0 | - | pass |
| 2 | #2451 feat: add pagination to list GHAS alerts tools | 2026-05-20 | 221 / code_scanning.go, dependabot.go, secret_scanning.go | 221 | none | 3 | none found | pass |
| 3 | #2284 feat: Add search commit tool | 2026-05-20 | 224 / minimal_types.go, search.go, tools.go | 224 | `MinimalCommitSearchItem.Repository`, `MinimalRepoRef.FullName`, `MinimalRepoRef.HTMLURL`, `MinimalRepoRef.Private`, `MinimalSearchCommitsResult`, `MinimalSearchCommitsResult.IncompleteResults`, `MinimalSearchCommitsResult.Items`, `MinimalSearchCommitsResult.TotalCount`, `SearchCommits`, `SearchIssuesResponse.IncompleteResults`, `SearchIssuesResponse.Items`, `SearchIssuesResponse.Total` | 1 | `SearchIssuesResponse.IncompleteResults`, `SearchIssuesResponse.Items`, `SearchIssuesResponse.Total` | pass |
| 4 | #2445 Add list_issue_fields tool | 2026-05-21 | 224 / issue_fields.go, tools.go | 224 | `IssueField`, `IssueField.DataType`, `IssueField.Description`, `IssueField.ID`, `IssueField.Name`, `IssueField.Options`, `IssueField.Visibility`, `IssueSingleSelectFieldOption`, `IssueSingleSelectFieldOption.Color`, `IssueSingleSelectFieldOption.Description`, `IssueSingleSelectFieldOption.ID`, `IssueSingleSelectFieldOption.Name`, `ListIssueFields`, `MinimalReviewThread.ID`, `issueFieldsOrgQuery`, `issueFieldsRepoQuery` | 1 | `MinimalCommitSearchItem.Repository`, `searchIssuesNodesQuery.Nodes` | pass |
| 5 | #2480 Add custom field filtering to list_issues | 2026-05-21 | 349 / issue_fields.go, issues.go | 349 | `issueFieldsRepoQuery` | 1 | `IssueField`, `IssueField.DataType`, `IssueField.Description`, `IssueField.Name`, `IssueField.Options` | pass |

Trivial commits inside the chain: #2515 Fix return Thread node ID in get_review_comments response (3 lines).

Placement: Five tool features across issues.go, search_utils.go, search.go, minimal_types.go, the three GHAS files and a new issue_fields.go; #2480 reuses #2445's IssueField types and GraphQL query structs, so where #2445 put them matters.

Caveats: Snapshot tests: steps 2-5 ship `__toolsnaps__/*.snap` files whose exact tool schema/description the task text must include. #2445/#2480 tests construct the unexported GraphQL query structs `issueFieldsRepoQuery{}`/`issueFieldsOrgQuery{}` (githubv4 mock matching). ~30 exported names (mostly result struct fields).

### 4. traefik/traefik `pkg/provider/kubernetes/gateway`

Seed `90a7b2b411dd`, final `074d02c09fae`. Package lines 597, whole PRs 970. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #14130 Support HTTPRoute backendRequest timeout | 2026-10-23 | 28 / features.go, httproute.go | 28 | none | 0 | - | pass |
| 2 | #14119 Support Gateway backend client certificate | 2026-10-26 | 152 / client.go, features.go, grpcroute.go, httproute.go, kubernetes.go, tlsroute.go | 152 | `Provider.loadGatewayBackendTLS`, `gatewayWithListeners.errResolvedRefsCondition` | 0 | none found | pass |
| 3 | #14127 Support HTTPRoute retry | 2026-10-27 | 54 / features.go, httproute.go, kubernetes.go | 54 | `Provider.RetryMaxRequestBodyBytes` | 0 | none found | pass |
| 4 | #14123 Support HTTPRoute RequestMirror filter | 2026-10-28 | 76 / features.go, httproute.go | 447 | `Provider.loadRequestMirror` | 0 | none found | pass |
| 5 | #13410 Support ClusterTrustBundle CA certificate references | 2026-10-28 | 287 / client.go, features.go, httproute.go, kubernetes.go, tlsroute.go | 289 | none | 0 | none found | pass |

Trivial commits inside the chain: #14125 Bump sigs.k8s.io/gateway-api to v1.7.0-rc.3 (12 lines), #14145 Merge branch v3.7 into master (2 lines).

Placement: Five Gateway API features that each touch features.go, httproute.go and some of kubernetes.go/client.go/grpcroute.go/tlsroute.go; backend TLS (#14119) and ClusterTrustBundle CA refs (#13410) share client/TLS loading code.

Caveats: Package lines are at the low end (597). Tests call unexported `Provider.loadGatewayBackendTLS`, `Provider.loadRequestMirror(...)` and the field `gatewayWithListeners.errResolvedRefsCondition`. #14123 also changes pkg/config/dynamic, a middleware and generated CRD code (whole PR 447 lines); #13410 adds a fixture YAML. Trivial commits in between: #14125 gateway-api bump (12 lines in the package), #14145 merge of v3.7 (2 lines).

### 5. github/github-mcp-server `pkg/github`

Seed `74964520cf15`, final `d5e1f48728dd`. Package lines 1243, whole PRs 1243. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #683 Add pagination support to GraphQL-based tools | 2025-07-21 | 382 / actions.go, discussions.go, issues.go, notifications.go, pullrequests.go, repositories.go, search.go, search_utils.go, server.go | 382 | `CursorPaginationParams.PerPage`, `PaginationParams.Page`, `PaginationParams.PerPage` | 0 | - | pass |
| 2 | #470 Add tools for sub-issue endpoint | 2025-07-22 | 407 / issues.go, tools.go | 407 | `AddSubIssue`, `ListSubIssues`, `RemoveSubIssue`, `ReprioritizeSubIssue` | 4 | `CursorPaginationParams.PerPage`, `PaginationParams.Page`, `PaginationParams.PerPage` | pass |
| 3 | #690 feat: list_discussions sort by updatedAt & createdAt, return updatedAt and author | 2025-07-24 | 306 / discussions.go | 306 | `NodeFragment.Number` | 0 | `CursorPaginationParams.After`, `GraphQLPaginationParams.After`, `GraphQLPaginationParams.First`, `PaginationParams.After` | pass |
| 4 | #774 Add updating draft state to `update_pull_request` tool | 2025-07-29 | 148 / pullrequests.go, tools.go | 148 | `BasicNoOrder.Repository`, `BasicWithOrder.Repository`, `NodeFragment.Number`, `NodeFragment.Title`, `WithCategoryAndOrder.Repository`, `WithCategoryNoOrder.Repository` | 1 | `BasicNoOrder.Repository`, `BasicWithOrder.Repository`, `NodeFragment.Number`, `NodeFragment.Title`, `WithCategoryAndOrder.Repository` | pass |

Trivial commits inside the chain: #740 Changed q to query in search (14 lines).

Placement: #683 adds pagination parameter types used across 9 tool files; #470 (sub-issues), #690 (discussions sort) and #774 (draft PR state) then use those params, so the first step's placement of shared helpers is exercised by all later steps.

Caveats: Merged 2025-07 (before the cutoff). Snapshot files in #470 and #774. Trivial commit #740 (14 lines, rename of a search param) sits between #690 and #774.

### 6. modelcontextprotocol/go-sdk `mcp`

Seed `d9263f7e47ba`, final `623780125cbb`. Package lines 638, whole PRs 703. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #646 mcp: support url mode elicitation | 2025-11-23 | 188 / client.go, protocol.go, requests.go, server.go, shared.go | 236 | `ClientOptions.ElicitationCompleteHandler`, `ClientOptions.ElicitationModes`, `ElicitParams.ElicitationID`, `ElicitParams.Mode`, `ElicitParams.URL`, `ElicitationCapabilities.URL`, `ElicitationCompleteNotificationRequest`, `ElicitationCompleteParams`, `ElicitationCompleteParams.ElicitationID`, `notificationElicitationComplete` | 0 | - | pass |
| 2 | #655 Expose JSON-RPC error codes, and URLElicitationRequired errors | 2025-12-01 | 258 / client.go, resource.go, server.go, shared.go, transport.go | 275 | `ClientOptions.ElicitationModes`, `CodeResourceNotFound`, `CodeURLElicitationRequired`, `ElicitParams.ElicitationID`, `ElicitParams.Mode`, `ElicitParams.URL`, `ElicitationCapabilities.Form`, `ElicitationCapabilities.URL`, `ElicitationCompleteParams`, `ElicitationCompleteParams.ElicitationID`, `FormElicitationCapabilities`, `URLElicitationCapabilities`, `URLElicitationRequiredError`, `notificationElicitationComplete`, `urlElicitationMiddleware` | 0 | `ClientOptions.ElicitationModes`, `ElicitParams.ElicitationID`, `ElicitParams.Mode`, `ElicitationCapabilities.Form`, `ElicitationCompleteParams` | pass |
| 3 | #663 mcp: support SSE polling and server disconnects (SEP-1699) | 2025-12-01 | 192 / event.go, shared.go, streamable.go | 192 | `ElicitParams.URL`, `ElicitationCapabilities.URL`, `RequestExtra.CloseStream` | 0 | none found | pass |

Trivial commits inside the chain: #665 fix: handle Windows CRLF in MCP client(#664) (8 lines).

Placement: URL-mode elicitation, then error codes / URLElicitationRequired errors built on it, then SSE polling; spread over client.go, server.go, shared.go, protocol.go, transport.go, streamable.go, event.go.

Caveats: Merged 2025-11/12. Tests use unexported `notificationElicitationComplete` and `urlElicitationMiddleware()`. #655 also edits jsonrpc/jsonrpc.go. Trivial commit #665 (8 lines, CRLF fix) between #655 and #663. Overlaps no other chain, but #2's package is the same.

### 7. mark3labs/mcp-go `server`

Seed `6200d8197d25`, final `887866786141`. Package lines 973, whole PRs 973. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #841 feat(server): add OAuth Protected Resource Metadata endpoint (RFC 9728) | 2026-05-04 | 243 / protected_resource.go, sse.go, streamable_http.go | 243 | `NewProtectedResourceMetadataHandler`, `ProtectedResourceMetadataConfig`, `ProtectedResourceMetadataConfig.AuthorizationServers`, `ProtectedResourceMetadataConfig.BearerMethodsSupported`, `ProtectedResourceMetadataConfig.JWKSURI`, `ProtectedResourceMetadataConfig.Resource`, `ProtectedResourceMetadataConfig.ResourceDocumentation`, `ProtectedResourceMetadataConfig.ResourceName`, `ProtectedResourceMetadataConfig.ScopesSupported`, `ProtectedResourceMetadataPath`, `WithProtectedResourceMetadata`, `WithSSEProtectedResourceMetadata` | 0 | - | pass |
| 2 | #843 feat: add opt-in CORS support to HTTP transports (#810) | 2026-05-04 | 343 / http_cors.go, sse.go, streamable_http.go | 343 | `CORSConfig`, `CORSConfig.AllowCredentials`, `CORSConfig.AllowedHeaders`, `CORSConfig.AllowedMethods`, `CORSConfig.AllowedOrigins`, `CORSConfig.ExposedHeaders`, `CORSConfig.MaxAge`, `CORSOption`, `WithCORSAllowCredentials`, `WithCORSAllowedHeaders`, `WithCORSAllowedMethods`, `WithCORSAllowedOrigins`, `WithCORSExposedHeaders`, `WithCORSMaxAge`, `WithSSECORS`, `WithStreamableHTTPCORS`, `CORSConfig.clone`, `CORSConfig.enabled`, `CORSConfig.resolveOrigin` | 0 | none found | pass |
| 3 | #844 feat: validate tool output against declared output schemas (#811) | 2026-05-04 | 387 / input_validation.go, output_validation.go, server.go | 387 | `WithOutputSchemaValidation`, `MCPServer.outputValidator`, `compiledSchemaCache.cached` | 0 | none found | pass |

Trivial commits inside the chain: none.

Placement: Three HTTP-transport features (protected-resource metadata, CORS, output-schema validation) that each need hooks in sse.go and streamable_http.go plus a new file; placement of shared HTTP middleware is a real choice.

Caveats: Weak build-on between steps (no step uses an earlier step's names); tests use unexported `CORSConfig.enabled/clone/resolveOrigin` and `MCPServer.outputValidator.cached`. ~34 names in total, most exported config fields.

### 8. jesseduffield/lazygit `pkg/config`

Seed `69635f379c81`, final `c826811eaf36`. Package lines 1206, whole PRs 3375. Seed test: pass.

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #5563 Revamp keybinding mechanism to support richer keybindings (e.g. `ctrl-alt-shift-x `) | 2026-04-30 | 279 / app_config.go, keynames.go, user_config.go | 627 | `KeyFromLabel`, `LabelForKey` | 0 | - | pass |
| 2 | #5565 Change canonical keybinding syntax (`<ctrl+a>` instead of `<c-a>`) | 2026-04-30 | 78 / keynames.go, user_config.go | 103 | none | 0 | `platformKeyBinding` | pass |
| 3 | #5571 Redraw only the spinner when nothing else has changed | 2026-05-06 | 19 / user_config_validation.go | 111 | none | 0 | none found | pass |
| 4 | #5634 Allow keybindings to configure more than one key | 2026-05-25 | 830 / app_config.go, dummies.go, keybinding.go, keynames.go, user_config.go, user_config_validation.go | 2534 | `Keybinding`, `KeybindingConfig.MergeLegacyAltKeybindings` | 0 | `GetValidatedKeyBindingKey`, `KeybindingUniversalConfig.BackspaceWord`, `KeybindingUniversalConfig.ForwardDeleteWord`, `KeybindingUniversalConfig.MoveWordLeft`, `KeybindingUniversalConfig.MoveWordRight` | pass |

Trivial commits inside the chain: #5564 Platform-specific edit bindings for move-by-word and backspace-word (8 lines), #5582 Revert the change of `<ctrl+j>`/`<ctrl+k>` to `<alt+down>`/`<alt+up>` (4 lines).

Placement: Keybinding revamp: #5563 introduces key names/labels, #5565 changes canonical syntax using them, #5634 lets a binding hold several keys; code spreads over keynames.go, user_config.go, app_config.go, keybinding.go, user_config_validation.go.

Caveats: Mostly outside the package: whole PRs total 3375 non-test lines (pkg/config 1206), touching gocui, gui and cheatsheet; the task would have to include those or be scoped down. #5571 contributes only 19 lines here. Trivial commits #5564 (8 lines) and #5582 (4-line revert of a default binding) sit inside the chain.

### 9. slack-go/slack `.`

Seed `4c472cd10a45`, final `8c5ef3c18b29`. Package lines 562, whole PRs 562. Seed test: pass except TestRTMDisconnect (dials slack.com).

| Step | PR | Merged | Pkg lines / files | Whole lines | Names this step's tests need (new in chain) | Snapshot/golden | Uses earlier steps' code | Test at merge |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | #1551 feat(block): add card and carousel blocks | 2026-04-17 | 251 / block.go, block_card.go, block_carousel.go, block_conv.go, chat.go, chat_stream_chunks.go | 251 | `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock`, `CardBlock.Actions`, `CardBlock.BlockType`, `CardBlock.Body`, `CardBlock.HeroImage`, `CardBlock.ID`, `CardBlock.Icon`, `CardBlock.Subtitle`, `CardBlock.Title`, `CardBlock.Type`, `CardBlock.WithActions`, `CardBlock.WithBody`, `CardBlock.WithHeroImage`, `CardBlock.WithIcon`, `CardBlock.WithSubtitle`, `CardBlock.WithTitle`, `CardBlockOptionBlockID`, `CarouselBlock`, `CarouselBlock.AddCard`, `CarouselBlock.BlockType`, `CarouselBlock.Elements`, `CarouselBlock.ID`, `CarouselBlock.Type`, `CarouselBlock.WithBlockID`, `MBTCard`, `MBTCarousel`, `MarkdownTextChunk.Text`, `MarkdownTextChunk.Type`, `NewCardBlock`, `NewCarouselBlock`, `PlanUpdateChunk.Title`, `PlanUpdateChunk.Type`, `TaskUpdateChunk.ID`, `TaskUpdateChunk.Title`, `TaskUpdateChunk.Type` | 0 | - | pass except TestRTMDisconnect (dials slack.com) |
| 2 | #1553 feat(assistant): add username and icon to status update | 2026-04-20 | 15 / assistant.go | 15 | `AssistantThreadsSetStatusParameters.IconEmoji`, `AssistantThreadsSetStatusParameters.IconURL`, `AssistantThreadsSetStatusParameters.Username` | 0 | none found | pass except TestRTMDisconnect (dials slack.com) |
| 3 | #1552 feat(block): add alert block | 2026-04-22 | 73 / block.go, block_alert.go, block_conv.go | 73 | `AlertBlock`, `AlertBlock.BlockType`, `AlertBlock.ID`, `AlertBlock.Level`, `AlertBlock.Text`, `AlertBlock.Type`, `AlertBlockOptionBlockID`, `AlertBlockOptionLevel`, `AlertLevelInfo`, `AlertLevelWarning`, `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock.BlockType`, `CardBlock.ID`, `CardBlock.Type`, `CarouselBlock.BlockType`, `CarouselBlock.ID`, `CarouselBlock.Type`, `MBTAlert`, `MarkdownTextChunk.Text`, `MarkdownTextChunk.Type`, `NewAlertBlock`, `PlanUpdateChunk.Type`, `TaskUpdateChunk.ID`, `TaskUpdateChunk.Type` | 0 | `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock.BlockID`, `CardBlock.BlockType`, `CardBlock.Type` | pass except TestRTMDisconnect (dials slack.com) |
| 4 | #1555 feat: add new block kit block Data Table | 2026-05-24 | 223 / block.go, block_conv.go, block_data_table.go, block_plan.go, block_task_card.go | 223 | `AlertBlock.BlockType`, `AlertBlock.ID`, `AlertBlock.Text`, `AlertBlock.Type`, `BlocksChunk.Blocks`, `BlocksChunk.Type`, `CardBlock.BlockType`, `CardBlock.ID`, `CardBlock.Type`, `CarouselBlock.BlockType`, `CarouselBlock.Elements`, `CarouselBlock.ID`, `CarouselBlock.Type`, `DataTableBlock`, `DataTableBlock.AddRow`, `DataTableBlock.BlockType`, `DataTableBlock.Caption`, `DataTableBlock.ID`, `DataTableBlock.PageSize`, `DataTableBlock.RowHeaderColumnIndex`, `DataTableBlock.Rows`, `DataTableBlock.Type`, `DataTableBlock.WithPageSize`, `DataTableBlock.WithRowHeaderColumnIndex`, `DataTableBlockOptionBlockID`, `DataTableCell.DataTableCellType`, `DataTableCellRawNumber`, `DataTableCellRawText`, `DataTableCellType`, `DataTableRawNumberCell`, `DataTableRawNumberCell.DataTableCellType`, `DataTableRawNumberCell.Text`, `DataTableRawNumberCell.Type`, `DataTableRawNumberCell.Value`, `DataTableRawNumberCell.WithText`, `DataTableRawTextCell`, `DataTableRawTextCell.DataTableCellType`, `DataTableRawTextCell.Text`, `DataTableRawTextCell.Type`, `DataTableRichTextCell`, `DataTableRichTextCell.DataTableCellType`, `DataTableRichTextCell.Elements`, `DataTableRichTextCell.Type`, `MBTDataTable`, `MarkdownTextChunk.Text`, `MarkdownTextChunk.Type`, `NewDataTableBlock`, `NewDataTableRawNumberCell`, `NewDataTableRawTextCell`, `PlanUpdateChunk.Type`, `TaskUpdateChunk.ID`, `TaskUpdateChunk.Type` | 0 | `AlertBlock.BlockID`, `AlertBlock.BlockType`, `AlertBlock.Text`, `AlertBlock.Type`, `BlocksChunk.Blocks` | pass except TestRTMDisconnect (dials slack.com) |

Trivial commits inside the chain: direct commit security: reject empty signing secret for NewSecretsVerifier (19 lines).

Placement: New Block Kit blocks registered in block.go/block_conv.go plus a new file each; #1552/#1555 reuse the chunk/card types from #1551.

Caveats: Name-heavy (about 86 exported names, all dictated by Slack's API) and mostly one-new-file-per-step, so placement choices are limited. #1553 is a 15-line field addition. A direct (no-PR) security commit of 19 lines sits inside the chain. TestRTMDisconnect dials slack.com and fails offline.

## Notable rejected chains

| Repo | Package | Chain | Failed criterion |
| --- | --- | --- | --- |
| cockroachdb/pebble | `.` | #6061 → #6047 → #6065 (2026-05) | Offline test cannot build with Go 1.27 (cockroachdb/swiss: undefined hashFn) at any commit; also steps are a fix, a perf pool and an option, not features |
| cockroachdb/pebble | `sstable` | #4812 → #4907 → #4909 → #4938 → #4908 (2025-06) | Mostly bug fixes; 15 golden datadriven files; 3 trivial commits inside |
| prometheus/prometheus | `web/api/v1, storage, tsdb` | #18624 → #18573 → #18724; #18497 → #18573 → #19135; #19013 → #19016 → #18905 (2026) | White-box tests: 64-68 unexported names to specify (web/api, storage); tsdb chain is two fixes and one feature |
| prometheus/prometheus | `discovery/aws` | #19504 → #19506 → #19524 → #19585 → #19584 (2026-08/09) | Performance and robustness fixes, each in one file (rds.go, elasticache.go, msk.go); no placement decisions |
| open-policy-agent/opa | `v1/ast` | #9318 → #9315 → #9327 → #9331 (2026-10) | Correctness/perf tweaks to compile.go and index.go, not features; 2 trivial commits inside |
| nats-io/nats-server | `server` | #7788 → #7892 → #7886 → #7895 → #7894 (2026-03) | 2164 package lines dominated by one 1852-line error-handling change; mostly fixes; very slow cluster tests (not run) |
| juicedata/juicefs | `pkg/meta` | #6656 → #6695 → #6698 (2026-02/03) | Step 2 is a simplification refactor and step 3 a test fix |
| jaegertracing/jaeger | `internal/storage/elasticsearch/esclient, v2/api/tracestore` | #8944 → #8958 → #8963 → #8967; #9655 → #9669 → #9670 → #9672 (2026) | Mostly migrations; 33-71 names to specify; under half of the lines inside the package |
| mark3labs/mcp-go | `server` | #610 → #617 → #613 → #624 → #626 (2025-10); #273 → #320 → #313 → #341 (2025-05) | 18-23 unexported names that the PRs invent are used by tests |
| ollama/ollama | `cmd/launch, create, server` | #17594 → #17733 → #17899 and others (2026) | 39-42 unexported names; model-support PRs mostly outside the package (5-20% inside) |
| github/github-mcp-server | `pkg/github` | #1113 → #1091 → #1114 → #1145 → #1152 (2025-09); #2964 → #2999 → #3020 → #2941 (2026-08) | 12 and 19 unexported names (GraphQL query structs, duplicate-detection types) |
| github/github-mcp-server | `pkg/github` | #834 → #869 → #866 → #938 → #919 (2025-08) | One new tool per PR in its own file; little building on earlier steps |
| temporalio/sdk-go | `internal` | #2634 → #2617 → #2633; #2394 → #2231 → #2397; #2721 → #2738 → #2781 → #2828 → #2780 (2026) | Mix of refactor/test-coverage PRs; up to 19 unexported names |
| pion/webrtc | `.` | #3548 → #3357 → #3569 (2026-09) | Unrelated changes (SCTP context, Cryptex, media direction) with no building on each other; 9 unexported names |
| traefik/traefik | `pkg/provider/kubernetes/ingress-nginx` | #13945 → #12918 → #13990 → #14012 → #14081 (2026-09/10) | 4 of 5 members are 'Merge branch v3.7 into master' commits |
| restic/restic | `internal/repository, internal/restic` | #21848 → #21851 → #21849 → #21861 → #21862 (2026-06) | Pure refactors (move/unexport types) |
| urfave/cli | `.` | #2449 → #2441 → #2451 → #2453 → #2454 (2026-10) | One 734-line build-tag feature plus four unrelated small fixes |
| open-telemetry/opentelemetry-collector | `confmap/internal` | #15142 → #15175 → #15339 → #15613 (2026) | API stabilisation/removal PRs, not features; half the lines outside the package |
| go-git/go-git | `plumbing/transport/http` | #2366 → #2422 → #2533 (2026-09/10) | Fix, protocol removal, fix; root-package tests need network (see the pairs round) |

## Screening

- Surveyed 61 module roots (40 from the pairs round, 13 skipped repos run serially after `go mod download`, 8 opentelemetry-collector modules); opentelemetry-collector's root module has no packages, so only its sub-modules were surveyed.
- For 372 packages with 8+ non-test files and 3+ namespaces, walked first-parent history since 2025-01-01. A commit is a chain member if it is a merged PR, its title is not lint/bump/chore/docs/refactor/rename/flaky-test/CI, it changes 10+ non-test lines and a test file in the package. A commit of 20 lines or fewer (tests included) may sit between members; anything else breaks the chain.
- 1001 windows of 3-5 members with 450-2600 package lines and 4+ package files, in 88 packages. 539 kept after requiring a later step to touch a file an earlier step touched and 45%+ of the lines inside the package.
- Name check on all 539 (plus 23 hand-picked windows): 358 needed 5+ unexported names that some chain PR invents; 293 needed 20+ names in total; 173 had no later step using an earlier step's new identifiers. 53 passed all three.
- About 30 windows read by hand; 9 kept. 49 offline test runs (seed plus every step's merge commit of the 10 finalists, pebble included).

Unexported names are counted whether or not they come from the step itself; exported names include struct fields, which inflates counts for API-mirroring packages (slack, github-mcp-server).
