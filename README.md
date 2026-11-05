# declscope in practice

[declscope](https://github.com/mpyw/declscope) は、AI エージェントが変更しやすい新しいコードを書く助けになるか？ オープンソースの Go プロジェクトで実際にマージされた PR の連続を、条件を変えてエージェントに実装させ、比べる。

## 実験

| # | 群 | 状態 | 記録 |
| --- | --- | --- | --- |
| 1 | ルールなし / ルールをプロンプトで渡す / ルールをプロンプトで渡し、declscope で検査する | 終了。結果は暫定で、分割の良さは AI の判定役だけで判定した。同僚による人の判定はまだ | [README](experiments/1/README.md)、[設計](experiments/1/design.md)、[報告](experiments/1/REPORT.md) |
| 2 | 指示なし / 最小限の対応を指示し、declscope で検査する / リファクタリングを頼み、declscope で検査する | 準備済み。まだ実行していない | [README](experiments/2/README.md)、[設計](experiments/2/design.md) |

> [!NOTE]
> **実験1の暫定の答え**: declscope の効果は、パッケージの既存の流儀しだいで向きが変わった。機能ごとにファイルを分けたパッケージでは分割が良くなり、層ごとや種類ごとに分けたパッケージでは悪くなった。ルールが守られたかには、はっきり差が出た。declscope で検査した群だけが、4課題すべてでルールを守った。

進み具合は [Epic #12](https://github.com/mpyw-yattemita/declscope-in-practice/issues/12) で追う。

## 構成

| パス | 中身 |
| --- | --- |
| `candidates/` | 課題の候補を探した記録。どの実験も共通で使う |
| `tasks/` | 課題（seed、各段の課題文、隠しテスト）。どの実験も共通で使う |
| `experiments/<n>/` | 実験ごとの設計、ハーネス、run の生データ、集計、報告 |
| [SAFETY.md](SAFETY.md) | 第三者のコードを実行するときのガードレール。どの実験にも当てはまる |

実験ごとのハーネスは、その実験で実際に動かしたものをそのまま残す。次の実験は、前の実験のハーネスを写して始める。実験1を移す前の状態には、tag `experiment-1` が付いている。
