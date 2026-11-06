# 各群のプロンプト（実験2）

測られるエージェントに渡す文言。エージェントへの入力なので、中身は英語で書く。

| ファイル | 中身 | 渡す群 |
| --- | --- | --- |
| [self-review.md](self-review.md) | 共通の自己レビュー手順。実験1と一字一句同じ | すべて |
| [scope-minimal.md](scope-minimal.md) | 変更を課題に必要な範囲に留め、既存のコードをリファクタリングしない | M、M′ |
| [scope-refactor.md](scope-refactor.md) | 課題の package を、変更にあわせてリファクタリングしてよい。振る舞いと export された API は変えない | R、R′ |
| [guardrail.md](guardrail.md) | declscope の案内。skill を読むこと、baseline を触らないこと、終える前のチェック | M′、R′ |

各群に渡すものは、上のファイルを次の順につないで作る。`{package}` は課題の package のパスに置き換える。M′ はコードでは `MD`、R′ は `RD`。

| 群 | 渡すもの | 終える前に実行するチェック |
| --- | --- | --- |
| M | 課題文、`scope-minimal.md`、`self-review.md` | `go vet ./...`、`go test ./<pkg>` |
| M′ | 課題文、`guardrail.md`、`scope-minimal.md`、`self-review.md` | M に加えて `declscope ./...`、`declscope shrink ./...` |
| R | 課題文、`scope-refactor.md`、`self-review.md` | M と同じ |
| R′ | 課題文、`guardrail.md`、`scope-refactor.md`、`self-review.md` | M′ と同じ |
| 段階2（すべての群） | 課題文、`self-review.md` | M と同じ |

M′ と R′ は、さらに declscope 0.24.2 の `declscope-authoring` と `declscope-adoption` を project skill として受け取る。ワークスペースには実験1と同じ `.declscope.yaml` を置き、baseline もこの設定で記録する。

```yaml
rules:
  naming:
    qualify: ondemand
    exported: true
```

## 文言の決めごと

| 決めごと | 理由 |
| --- | --- |
| 範囲の段落は、declscope のあり・なしで一字一句同じにする。skill の名前もルールも書かない | R と R′、M と M′ の差を、declscope の有無だけにするため |
| declscope の有無は、`guardrail.md` と skill だけで分ける | 同じ理由 |
| `rule.md`（置き場所のルールを平易な言葉で書いたもの）を渡さない | 実験1でこれが要ったのは、declscope を使わない P 群にルールを伝えるため。実験2の declscope なしの群は、ルールを知らずに書く |
| リファクタリングの範囲は課題の package にする | ほかの package に広げると、隠しテストと名前合わせの前提が崩れる |
| export された API を変えないよう頼む | 上流の隠しテストは export された名前で呼ぶので、変えると名前合わせ役の仕事が増え、群の差に混ざる |
| R′ にも baseline の再生成を禁じる | skill は、リファクタリングのあとに `declscope baseline -merge` で古い記録を消すよう書いているが、記録が古いまま残っても指摘は増えない。再生成を許すと、新しい指摘を記録して隠す道ができる |
