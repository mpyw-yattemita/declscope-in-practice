# 課題

[candidates/README.md](../candidates/README.md) の「選んだ課題」の6本の連続を、ハーネスが読める形にしたもの。考え方は [design.md](../design.md) の「課題」「隠しテストの名前合わせ」「2つの段階」「隔離」にある。検証の結果は [VERIFY.md](VERIFY.md) にある。

> [!CAUTION]
> **opa-pr8427 は除外した。** テストが外部のサイトから JSON スキーマを取りに行き、Go 1.25.5 のツールチェーンのダウンロードも要る。[harness/SAFETY.md](../harness/SAFETY.md) のルール3による。判定は `opa-pr8427/prescreen.json` にある。


| 課題 | リポジトリ / package | 段 | Go |
| --- | --- | --- | --- |
| [hcl-pr841](hcl-pr841/) | hashicorp/hcl `hclwrite` | #841 → #811 → #820 → #835 | 1.27.0 |
| [opa-pr8427](opa-pr8427/) | open-policy-agent/opa `v1/ast` | #8427 → #8482 → #8485 → #8468 → #8511 | 1.25.5 |
| [jaeger-pr9256](jaeger-pr9256/) | jaegertracing/jaeger `internal/storage/v2/api/tracestore` | #9256 → #9356 → #9357 | 1.27.0 |
| [github-mcp-server-pr683](github-mcp-server-pr683/) | github/github-mcp-server `pkg/github` | #683 → #470 → #690 → #740 → #774 | 1.27.0 |
| [go-sdk-pr644](go-sdk-pr644/) | modelcontextprotocol/go-sdk `mcp` | #644 → #646 → #655 → #665 → #663 | 1.27.0 |
| [traefik-pr14130](traefik-pr14130/) | traefik/traefik `pkg/provider/kubernetes/gateway` | #14130 → #14119 → #14127 → #14123 → #13410 | 1.27.0 |

## 構成

```text
tasks/<課題>/
  task.yaml                 課題全体と各段の定義
  step-<n>/
    task.md                 エージェントに渡す課題文（英語）
    given/<repo のパス>      課題文と一緒に渡す、外から見える仕様のファイル（あるときだけ）
    outside.patch           package の外の変更。ハーネスが段の前に当てる（treatment: harness のときだけ）
    intermediate.patch      前の段とこの段のあいだの commit の差分。ハーネスが段の前に当てる（あるときだけ）
    hidden/<repo のパス>     採点で使う upstream のテスト一式
    names.yaml              名前合わせ役への手がかり
```

`hidden/` と `given/` の中は、リポジトリのルートからのパスのまま置いてある。そのままワークスペースのルートに重ねればよい。

> [!IMPORTANT]
> `hidden/`、`names.yaml`、`outside.patch`、`intermediate.patch`、`task.yaml` は、測られるエージェントに見せない。エージェントが受け取るのは `task.md` と `given/` だけ。

## task.yaml

| フィールド | 中身 |
| --- | --- |
| `id` | 課題の名前。ディレクトリ名と同じ |
| `repo` | upstream のリポジトリの URL |
| `package` | 対象の package の、module の中のパス |
| `module_dir` | module がリポジトリのルートにないときだけ。6本ともルートにあるので、どの課題にもない |
| `go` | 使う Go のバージョン。`GOTOOLCHAIN=go<go>` で指定する |
| `env` | テストのときの環境変数。6本とも `CGO_ENABLED=0` |
| `seed` | seed の commit。最初の段の parent と同じ |
| `notes` | 課題全体についての注意 |
| `steps` | 段の一覧。実装する順に並ぶ |

`steps` の各要素は次のとおり。

| フィールド | 中身 |
| --- | --- |
| `n`, `pr`, `title` | 段の番号、upstream の PR 番号、PR のタイトル |
| `merge` | その PR のマージ commit。隠しテストはこの commit から取った |
| `parent` | マージ commit の最初の親。upstream でその PR を当てる直前の状態 |
| `stage` | `1` か `2`。最後の段だけ `2` |
| `task` | 課題文のパス |
| `given` | 課題文と一緒に渡すファイル。リポジトリのルートからのパス。ないときは空 |
| `outside.files` | その PR が package の外で変えたファイルと行数 |
| `outside.treatment` | `none`（外の変更がない）、`agent`（エージェントが外も実装する。課題文にそう書いてある）、`harness`（ハーネスが upstream の差分を当てる） |
| `outside.reason` | その扱いにした理由 |
| `outside.patch` | `harness` のときの差分のパス |
| `test_cmd` | 採点で実行するコマンド。module のルートで実行する |
| `known_failing` | 実装と関係なく落ちるので外すテストと、その理由。`test_cmd` の `-skip` に入っている |
| `white_box_tests` | 内部の構造を直接調べるテストと、その理由。機能が正しくても、内部の作りが upstream と違えば落ちる。集計で別に扱うかどうかは、試運転のあとで決める |
| `intermediate.commits` | 前の段のマージ commit からこの段の parent までの commit のうち、package、package が使う module 内の package、`go.mod`、`go.sum` のどれかを変えたもの |
| `intermediate.handling` | その commit の扱い |
| `intermediate.patch` | ハーネスが当てる差分のパス |
| `notes` | その段についての注意 |

## 1つの段の流れ

| 順 | すること | する人 |
| --- | --- | --- |
| 1 | `intermediate.patch` があれば、ワークスペースに当てる | ハーネス |
| 2 | `outside.treatment` が `harness` なら、`outside.patch` を当てる | ハーネス |
| 3 | `given/` をワークスペースに重ねる | ハーネス |
| 4 | `task.md` を渡して実装させる | 測られるエージェント |
| 5 | 下の「採点」をする | ハーネス、名前合わせ役 |

1〜3 はエージェントが始める前に済ませる。段階1か段階2かは `stage` で決まる。段階2では、どの群も A 群の設定で実行する（design.md の「2つの段階」）。

## 採点

隠しテストを差分で当てると、エージェントのテストとぶつかったり、同じ名前のテストが2つできたりする。そこで、package のテストを丸ごと入れ替える。

1. ワークスペースから、package の `_test.go` をすべて一時的に取り除く。対象は package のディレクトリの直下だけ。サブディレクトリの別の package は触らない。
2. `step-<n>/hidden/` をワークスペースのルートに重ねる。中身は、その段のマージ commit での package の `_test.go` すべてと、テストが読むデータ（traefik の `fixtures/` など）。データは、ワークスペースにある同じパスのファイルを上書きする。
3. 1段目から n段目までの `given/` を、もう一度重ねる。エージェントが仕様のファイルを書き換えていても、採点は upstream の仕様で行う。
4. 名前合わせ役に、`names.yaml` を手がかりとして、隠しテストの名前をエージェントのコードに合わせさせる。直してよいものは design.md の「隠しテストの名前合わせ」のとおり。
5. 合わせる前後の差分で、期待する値と assert の行が変わっていないことを機械的に確かめる。
6. `GOTOOLCHAIN=go<go>`、`env`、`GOPROXY=off`、外向きのネットワークを遮断した状態で、`test_cmd` を実行する。
7. エージェントの `_test.go` を元に戻し、2 と 3 で重ねたファイルを取り除く（3 で上書きした `given/` は、エージェントが残した版に戻す）。

隠しテストは、その段までの upstream のテスト一式なので、前の段の機能と seed からある挙動も、毎回確かめ直すことになる。

### 中に入れていないデータ

package の中のデータのディレクトリで、連続のあいだに upstream が一度も変えていないもの（opa の `v1/ast/testdata`、go-sdk の `mcp/testdata`）は、`hidden/` に入れていない。seed の時点でワークスペースにあり、そのままで upstream と同じだから。エージェントがそこを書き換えた場合に備えるなら、採点の前に seed の版に戻す。

## names.yaml

隠しテストが使う名前のうち、seed の package にはなく、連続のどこかの段で upstream が足したものの一覧。

| フィールド | 中身 |
| --- | --- |
| `name` | `型.名前`、または top-level の名前 |
| `kind` | `func`、`method`、`type`、`field`、`const`、`var` など |
| `exported` | export されているか |
| `signature` | その段のマージ commit での upstream の宣言 |
| `used_in` | その名前を使うテストのファイル |
| `risk` | 名前合わせで困りそうな理由。unexported なフィールドを読む、小さな内部の helper を直接呼ぶ、など |
| `false_positive` | 人が確かめて、別の同じつづりの名前だと分かったもの |
| `maybe_false_positive` | テストの中の struct に同じ名前のフィールドがあり、取り違えかもしれないもの |

一覧は、テストの識別子と宣言の名前をつづりで突き合わせて作った。型は見ていない。`Name` のようなありふれた名前は、取り違えが混ざる。

## 作り方

`task.yaml`、`hidden/`、`given/`、`outside.patch`、`intermediate.patch`、`names.yaml` は、upstream の clone からスクリプトで作った。`task.md` は、PR とその issue の説明から手で書いた。課題文には、export される名前と引数、外から見える挙動（エラーの文言、wire format、設定の名前）を書き、ファイル名、unexported な名前、置き場所、内部の設計は書いていない。
