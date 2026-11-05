# ハーネス（実験2）

1つの課題（連続した PR）を、1つの群で1回実行し、結果を `runs/` に回収する。コマンドは `experiments/2/` で実行する。

```bash
python3 -I harness/run.py --task tasks/<id> --arm N|M|R --rep <n> --clones <upstream の clone を置いたディレクトリ>
```

`--steps 1,2` を付けると、指定した段だけを実行する。

## 準備

M 群と R 群が使う declscope と skill を、実験2のために固定して置く。場所は環境変数 `DECLSCOPE_HOME` で変えられる。

```bash
D=~/ws-tools/declscope-0.24.2
GOBIN=$D/bin go install github.com/mpyw/declscope/cmd/declscope@v0.24.2
$D/bin/declscope skill install -dir $D/skills
```

`setup_guardrail` は、実行のたびに `declscope -V=full` が 0.24.2 であることを確かめる。違えば止まる。

全課題をまとめて実行するときは `run_all.py` を使う。1回ずつ順に実行し、run と評価ごとに commit して push する。

```bash
python3 -I harness/run_all.py --clones <upstream の clone を置いたディレクトリ>
```

## 1回の実行の流れ

| 段階 | やること |
| --- | --- |
| 準備 | `~/ws/<ランダム>/` を作り、seed を `git archive` で書き出して、新しいリポジトリに commit する。`~/ws/AGENTS.md` に `workspace-AGENTS.md` を置く |
| モジュールの取得 | ネットワークありで `go mod download` を実行し、キャッシュを温める。以降は `GOPROXY=off` |
| M 群と R 群だけ | declscope が 0.24.2 であることを確かめ、`declscope baseline ./...` で seed の指摘を記録し、`declscope-authoring` と `declscope-adoption` を `.claude/skills/` に置く |
| 段の前 | upstream の段のあいだの変更を当てる。package の中の変更（`intermediate-pkg.patch`）は 3-way merge で当て、当たらなければその説明（`intermediate-pkg.md`）を課題文の先頭に足す。どちらにしたかは `metrics.json` の `intermediate_pkg` に残る。続けて package の外の変更（`intermediate.patch`、`outside.patch`）を当てる |
| 各段 | 課題文と各群のプロンプトをつないで `claude -p` に渡す。終わったら commit し、差分を保存する |
| 採点 | 下の「採点」のとおり |
| 監査 | トランスクリプトを読み、ワークスペースの外のパス、ネットワーク操作、N 群での declscope の語を探す |
| 段階2 | 最後の段の前に、baseline、skill、`//declscope:` の行を取り除く。N 群と同じ設定で実行する |

## 測られるエージェントの環境

| 項目 | 値 | 理由 |
| --- | --- | --- |
| `PATH` | Go の `bin`、`/usr/bin`、`/bin`、`/usr/sbin`、`/sbin`。M 群と R 群だけ `~/ws-tools/declscope-0.24.2/bin` を先頭に足す。このディレクトリには declscope しかない | シェルの設定ファイルは、mise が `PATH` にあるときだけ mise を有効にする。mise は declscope を `PATH` に入れるので、mise を外す |
| `GOPROXY` | `off` | upstream の答えを取りに行けないようにする |
| `GOTOOLCHAIN` | 課題の `go` | opa は Go 1.25.5 でしかテストが通らない |
| ツール | `WebFetch` と `WebSearch` を禁止。MCP サーバーは読み込まない（`--strict-mcp-config`） | ネットワークと外部サービスを使わせない |

最小限の環境で、認証が通り、declscope が見えず、`~/ws/AGENTS.md` が読み込まれることを、試しの実行で確かめた。

## 採点

| 手順 | 内容 |
| --- | --- |
| 1 | package の `_test.go` をすべて退避する（エージェントのテスト） |
| 2 | その段の upstream のマージ commit での、package のテスト一式と testdata を置く |
| 3 | コンパイルできなければ、名前合わせ役を呼ぶ。名前合わせ役は、`//declscope:` の行、baseline、skill を取り除いた別のワークスペースで、テストファイルだけを直す |
| 4 | 名前合わせのあと、各テスト関数のリテラル（文字列、数値）が順番も含めて変わっていないことを確かめる。消してよいのは `NO_COUNTERPART.txt` に書いたテスト関数だけ |
| 5 | `go test -json` を実行する。`known_failing` のテストは `-skip` で外す |
| 6 | upstream のファイルを取り除き、エージェントのテストを元に戻す |

## 出力

`runs/<task>/<arm>/<rep>/` に保存する。

| ファイル | 中身 |
| --- | --- |
| `run.json` | 実行全体のまとめ |
| `step-<n>/metrics.json` | その段の結果。トークン数、ターン数、触ったファイル数と行数、ディレクティブの数、採点、監査 |
| `step-<n>/agent.jsonl.gz` | 測られるエージェントのトランスクリプト |
| `step-<n>/agent.prompt.md` | 渡したプロンプト |
| `step-<n>/diff.patch` | その段の差分 |
| `step-<n>/test.jsonl` | 採点の `go test -json` の出力 |
| `step-<n>/rename.*` | 名前合わせ役のトランスクリプトと差分（呼んだときだけ） |
