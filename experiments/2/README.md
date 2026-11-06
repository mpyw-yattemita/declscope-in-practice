# 実験2: 既存のコードをどこまで変えさせるか

declscope を入れたエージェントに、最小限の対応だけをさせるのと、リファクタリングまで頼むのとで、新しいコードの分割がどう変わるかを比べる。実験1で、declscope の効果がパッケージの既存の流儀で向きを変えたのを受けたもの。

| 群 | 条件 |
| --- | --- |
| N: 指示なし | 課題文と自己レビュー手順だけ。実験1の A と同じ |
| M: 最小限＋declscope | declscope 0.24.2 と skill を渡し、変更を課題に必要な範囲に留めるよう頼む |
| R: リファクタリング＋declscope | declscope 0.24.2 と skill を渡し、課題の package のリファクタリングも頼む |

設計と判定ルールは [design.md](design.md)、プロンプトは [harness/prompts/](harness/prompts/README.md)、実行のしかたは [harness/README.md](harness/README.md) にある。

## 状態

準備済み。まだ実行していない。

| 済んだこと | 内容 |
| --- | --- |
| ハーネス | 実験1のものを写し、群を N、M、R に変えた。declscope 0.24.2 に固定し、版を確かめる |
| 段のあいだの変更 | 実験1で当てていなかった package の中の変更を当てるようにした（traefik の2段目） |
| プロンプト | M と R の違いを、変更の範囲の1段落だけにした |
| 判定ルール | 実験1と同じルールを、どの run より前に固定した |
| 人が判定する8組 | どの run より前に固定した |

## 始め方

```bash
cd experiments/2
python3 -I harness/run_all.py
```

`run_all.py` は、課題ごとに N、M、R を1回ずつ順に実行する。1つの run が終わるたびに commit して push し、3群がそろった課題には判定役と構造レビューを走らせる。止まっても、もう一度実行すれば、終わった run を飛ばして続きから始める。使用量の上限に当たったときは、15分待ってからやり直す。

全課題が終わったら、次を実行して集計する。

```bash
python3 -I analysis/compliance.py --root runs   # 段階1の終わりに declscope が報告するもの
python3 -I analysis/analyze.py --root runs      # 判定ルールの結果。analysis/results.md に書く
```

報告は、実験1の [REPORT.md](../1/REPORT.md) と同じ形で `REPORT.md` に書く。

## 前提になっているもの

このリポジトリの外にあり、実行に要るもの。どれも用意済み。消えていたら、右の方法で作り直す。

| もの | 場所 | 作り直し方 |
| --- | --- | --- |
| upstream の clone（5本） | `~/ws-clones/<owner>__<repo>`。`--clones` か環境変数 `WS_CLONES` で変えられる | `git clone https://github.com/<owner>/<repo> ~/ws-clones/<owner>__<repo>`。全履歴が要る |
| declscope 0.24.2 と skill | `~/ws-tools/declscope-0.24.2/` | [harness/README.md](harness/README.md) の「準備」 |
| パッケージの地図を作るツール | `~/ws-tools/pkgmap` | `harness/pkgmap/` で `go build -o ~/ws-tools/pkgmap .` |
| 第三者のコードを sandbox で包む hook | `~/.claude/hooks/untrusted-code-sandbox.py`、`~/.claude/hooks/nonet.sb` | [SAFETY.md](../../SAFETY.md) |
| 課題の Go ツールチェーン | Go のモジュールキャッシュ（`golang.org/toolchain`）。課題の `go` の版（いまは 1.27.0） | `GOTOOLCHAIN=go1.27.0 go version`。ハーネスは `GOSUMDB=off` で動くので、キャッシュにないツールチェーンを取りに行けず、`go mod download` で止まる |
| ワークスペース | `~/ws`、`~/ws-tmp`、`~/ws-archive` | ハーネスが作る |
| 人の判定の並び順の秘密 | `analysis/human/.salt`。git に入れていない | 判定画面を作る前なら、新しく作り直し、そのハッシュで `pairs.json` の `salt_sha256` を置き換えてよい。並び順は、群や AI の結果とは関係がない |
| 人の判定の答えを読み解く対応表 | `analysis/human/key.json`。判定画面を作るときにできる。git に入れていない | 作り直せない。判定を始めたら、バックアップしておく |

> [!CAUTION]
> 第三者のコードを実行する前に、[SAFETY.md](../../SAFETY.md) と、トップの [AGENTS.md](../../AGENTS.md) を読むこと。テストの実行をサブエージェントに任せない。
