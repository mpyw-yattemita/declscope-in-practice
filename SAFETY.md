# 第三者のコードを実行するときのガードレール

> [!CAUTION]
> 2026-11-05 に、候補探しのサブエージェントが ollama の `cmd/launch` のテストを実行した。このテストは OpenClaw、Cline、Kimi などの外部ツールを「インストールする」処理を含んでいた。テスト本体は sandbox の中だったが、その直前の「温める」段階（`go test -run '^$'`）を sandbox の外、ネットワークありで実行していた。この段階でもテストのバイナリは起動する。ダウンロードが Microsoft Defender に検知され、マシンがネットワークから隔離された。

この文書のルールは、この実験で第三者のコードを扱うすべての場面に当てはまる。例外はない。

## ルール

| # | ルール | 守らせる仕組み |
| --- | --- | --- |
| 1 | **第三者のコードを実行するコマンドは、ネットワークを遮断した sandbox の中でだけ実行する。** `go test`、`go run`、`go generate`、テストのバイナリ、`make`、`npm`、`npx`、`curl` などが対象 | ユーザー全体の hook（下の「hook」）が自動で包み直す。ハーネスは `sandboxed()` を通す |
| 2 | **ネットワークありで実行してよいのは、コードを実行しないコマンドだけ。** `go mod download`、`go list`、`git clone`、`git archive`。`go test -run '^$'` は、テストを1つも実行しなくても TestMain と init が走るので、ここに含めない | ハーネスの `warm()` |
| 3 | **テストを実行する前に、親エージェントが題材を事前チェックする。** 外部ツールのダウンロードやインストールが少しでもあれば、その題材は除外する。除外した題材を許可することはない | 各実験の `harness/prescreen.py`。ハーネスは `prescreen.json` が `clean` でない課題を実行しない |
| 4 | **サブエージェントに第三者のコードのテストを任せない。** テストを実行するのは、事前チェックを通った課題について、親エージェントかハーネスだけ | このルール。サブエージェントに渡す指示にも書く |
| 5 | **コードの実行に使うプロファイルは1つ。** `~/.claude/hooks/nonet.sb`。localhost 以外への外向きの接続と DNS を拒否する | hook とハーネスが同じファイルを使う |

## hook

`~/.claude/settings.json` の `PreToolUse`（`Bash`）に `~/.claude/hooks/untrusted-code-sandbox.py` を登録した。私、サブエージェント、ハーネスが起動する `claude -p` のすべてに効く。

| 条件 | 動き |
| --- | --- |
| 作業ディレクトリかコマンドが第三者のコードの場所を指す（scratchpad、`~/ws`、`~/ws-*`。`~` や `$HOME` の書き方も含む） | 次の行へ |
| コマンドがコードを実行するか、ネットワークから取ってくる | `sandbox-exec -f ~/.claude/hooks/nonet.sb` で包み、`GOPROXY=off` にして実行する |
| hook が入力を読めない、プロファイルがない | 実行を拒否する（`onFailure: block`） |

包み直したコマンドは `~/.claude/hooks/untrusted-code-sandbox.log` に記録される。

## 事前チェック

各実験の `harness/` にある。実験1では、`experiments/1/` で次のように実行する。

```bash
python3 -I harness/prescreen.py --task tasks/<id> --clones <clone の置き場所>
python3 -I harness/prescreen_selftest.py   # チェック自体の自己テスト
```

コードは一切実行しない。git のオブジェクトからファイルを読むだけで、checkout もしない。seed と各段のマージ commit で、課題の package と、それが import する同じ module の package を、テストファイルも含めて調べる。

| 判定 | 意味 | 扱い |
| --- | --- | --- |
| `excluded` | 文字列の中にインストーラーの名前（`npm`、`pip`、`brew`、`curl`、`go install` など）か、ダウンロードの URL がある | 除外する。許可しない |
| `needs-review` | プロセスの起動か、外向きの HTTP 呼び出しがある | 親エージェントが中身を読み、何を起動し、どこに接続するかを `prescreen.json` の `review` に書く。外部ツールのダウンロードでないと確かめたときだけ `clean` にする |
| `clean` | 何も見つからない | 実行してよい |

自己テストは、事故を起こした ollama の形（`exec.Command("npm", "install", "-g", ...)` と、テストの中の複数行のシェルスクリプト）を捕まえることと、Go の変数名 `uv` やコメントの中の `npm` を誤って捕まえないことを確かめる。実際の ollama `cmd/launch` でも、131件の `excluded` になることを確かめた。

## 除外リスト

| リポジトリ / package | 理由 |
| --- | --- |
| ollama/ollama（すべての package） | `cmd/launch` のテストが外部ツールをインストールする。事故の原因 |
