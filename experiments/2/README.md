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
python3 -I harness/run_all.py --clones <upstream の clone を置いたディレクトリ>
```
