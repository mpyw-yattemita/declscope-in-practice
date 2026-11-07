# 実験2: declscope は、どう使わせると役に立つか

変更の範囲（最小限の対応だけ／リファクタリングしてよい）と、declscope の有無を掛け合わせた 2×2 で比べる。実験1で、declscope の効果がパッケージの既存の流儀で向きを変えたのを受けたもの。

| | declscope なし | declscope あり |
| --- | --- | --- |
| 最小限の対応だけ | M | M′（コードでは `MD`） |
| リファクタリングしてよい | R | R′（コードでは `RD`） |

主な比較は R′ 対 R（リファクタリングさせるとき、declscope が役に立つか）。M′ と R′ では、declscope の検査が役に立ったかも段ごとに記録する。

設計と判定ルールは [design.md](design.md)、プロンプトは [harness/prompts/](harness/prompts/README.md)、実行のしかたは [harness/README.md](harness/README.md) にある。

## 状態

4課題すべてを実行した。結果は [REPORT.md](REPORT.md) にある（暫定。分割の良さは AI の判定役だけで判定した）。

**分割の良さを一番左右したのは、リファクタリングを頼むかどうかだった。** declscope は、ルールを守らせることには確実に効いた（M′ と R′ は4課題すべてで指摘 0 件）。リファクタリングさせるときは、declscope のない R のほうが判定役に好まれた。最小限の対応では、declscope のある M′ が少し良かった。

## 始め方

```bash
cd experiments/2
python3 -I harness/run_all.py
```

`run_all.py` は、課題ごとに M、M′、R、R′ を1回ずつ順に実行する。1つの run が終わるたびに commit して push し、3群がそろった課題には判定役と構造レビューを走らせる。止まっても、もう一度実行すれば、終わった run を飛ばして続きから始める。使用量の上限に当たったときは、15分待ってからやり直す。

全課題が終わったら、次を実行して集計する。

```bash
python3 -I analysis/compliance.py --root runs   # 段階1の終わりに declscope が報告するもの
python3 -I analysis/guardrail_use.py --root runs  # declscope の検査が役に立ったか（run_all.py も課題ごとに実行する）
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
| 課題の Go ツールチェーン | Go のモジュールキャッシュ（`golang.org/toolchain@v0.0.1-go<版>.<OS>-<arch>`）。課題の `go` の版（いまは 1.27.0） | `GOTOOLCHAIN=go1.27.0 go version`。キャッシュになければ、ハーネスがこのコマンドを示して止まる |
| ワークスペース | `~/ws`、`~/ws-tmp`、`~/ws-archive` | ハーネスが作る |
| 人の判定の並び順の秘密 | `analysis/human/.salt`。git に入れていない | 判定画面を作る前なら、新しく作り直し、そのハッシュで `pairs.json` の `salt_sha256` を置き換えてよい。並び順は、群や AI の結果とは関係がない |
| 人の判定の答えを読み解く対応表 | `analysis/human/key.json`。判定画面を作るときにできる。git に入れていない | 作り直せない。判定を始めたら、バックアップしておく |

> [!CAUTION]
> 第三者のコードを実行する前に、[SAFETY.md](../../SAFETY.md) と、トップの [AGENTS.md](../../AGENTS.md) を読むこと。テストの実行をサブエージェントに任せない。
