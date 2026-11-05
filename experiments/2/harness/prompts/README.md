# 各群のプロンプト（実験2）

測られるエージェントに渡す文言。エージェントへの入力なので、中身は英語で書く。

| ファイル | 中身 | 渡す群 |
| --- | --- | --- |
| [self-review.md](self-review.md) | 共通の自己レビュー手順。実験1と一字一句同じ | N、M、R |
| [scope-minimal.md](scope-minimal.md) | 変更を課題に必要な範囲に留め、既存のコードをリファクタリングしない | M |
| [scope-refactor.md](scope-refactor.md) | 課題の package を、変更にあわせてリファクタリングしてよい。振る舞いと export された API は変えない | R |
| [guardrail.md](guardrail.md) | declscope の案内。skill を読むこと、baseline を触らないこと、終える前のチェック | M、R |

各群に渡すものは、上のファイルを次の順につないで作る。`{package}` は課題の package のパスに置き換える。

| 群 | 渡すもの | 終える前に実行するチェック |
| --- | --- | --- |
| N | 課題文、`self-review.md` | `go vet ./...`、`go test ./<pkg>` |
| M | 課題文、`guardrail.md`、`scope-minimal.md`、`self-review.md` | N に加えて `declscope ./...`、`declscope shrink ./...` |
| R | 課題文、`guardrail.md`、`scope-refactor.md`、`self-review.md` | M と同じ |

M と R は、さらに declscope 0.24.2 の `declscope-authoring` と `declscope-adoption` を project skill として受け取る。ワークスペースには実験1と同じ `.declscope.yaml` を置き、baseline もこの設定で記録する。

```yaml
rules:
  naming:
    qualify: ondemand
    exported: true
```

## 実験1からの違い

| 違い | 理由 |
| --- | --- |
| `rule.md`（置き場所のルールを平易な言葉で書いたもの）を渡さない | 実験1でこれが要ったのは、declscope を使わない P 群にルールを伝えるため。実験2の M と R は skill で同じ内容を読む。N にはルールを渡さない |
| `guardrail.md` の冒頭を、`rule.md` を前提にしない書き方にした | 同じ理由 |
| M と R の違いは、スコープの1段落だけにする | 比べたいのは、リファクタリングを頼むかどうかの効果だけだから |

## R の文言の決めごと

- 「`declscope-authoring` の慣習に合わない構造」をリファクタリングの対象に挙げた。declscope 0.24.2 の skill は、ユーザーがリファクタリングを頼んだときに、その範囲で skill の慣習に従うよう書いている。R は、その経路を測る
- 範囲は課題の package にした。ほかの package に広げると、隠しテストと名前合わせの前提が崩れる
- export された API を変えないよう頼んだ。上流の隠しテストは export された名前で呼ぶので、変えると名前合わせ役の仕事が増え、群の差に混ざる
- baseline の再生成は、M と同じく禁じたままにする。skill は、リファクタリングのあとに `declscope baseline -merge` で古い記録を消すよう書いているが、記録が古いまま残っても指摘は増えない。再生成を許すと、新しい指摘を記録して隠す道ができる
