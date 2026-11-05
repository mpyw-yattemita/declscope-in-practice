# 各群のプロンプト

測られるエージェントに渡す文言。エージェントへの入力なので、中身は英語で書く。

| ファイル | 中身 | 渡す群 |
| --- | --- | --- |
| [self-review.md](self-review.md) | 共通の自己レビュー手順 | A、P、G。一字一句同じ |
| [rule.md](rule.md) | 置き場所のルール。`declscope-authoring` からツールを取り除いたもの | P、G |
| [guardrail.md](guardrail.md) | declscope の案内。baseline を触らないこと、終える前のチェック | G |

各群に渡すものは、上のファイルを次の順につないで作る。

| 群 | 渡すもの | 終える前に実行するチェック |
| --- | --- | --- |
| A | 課題文、`self-review.md` | `go vet ./...`、`go test ./<pkg>` |
| P | 課題文、`rule.md`、`self-review.md` | A と同じ |
| G | 課題文、`rule.md`、`guardrail.md`、`self-review.md` | A に加えて `declscope ./...`、`declscope shrink ./...` |

G 群は、さらに `declscope-authoring` を project skill として受け取る。ワークスペースには次の `.declscope.yaml` を置き、baseline もこの設定で記録する。

```yaml
rules:
  naming:
    qualify: ondemand
    exported: true
```

## rule.md の作り方

`declscope-authoring`（declscope 0.24.1）の内容から、ツールに関わるものをすべて取り除いた。

| 残したもの | 理由 |
| --- | --- |
| unexported な宣言は宣言したファイルの中だけで使う、というルール | declscope のデフォルト設定が検査するのはこれ |
| フィールド、メソッド、テストファイルの扱い | 同じルールの一部として検査される |
| 新しい宣言の置き場所、よくある変更、隠してはいけないこと | ルールを守るための判断の指針で、skill の中心 |
| `internal/` の中で、ほかの package が使わないものを export しない | `declscope shrink` が検査する内容 |
| 宣言の並べ方 | skill に書いてある慣習で、G 群も読む |
| 命名のルール（名前にファイル名の語を含める） | G 群は `qualify: ondemand`、`exported: true` で検査されるので、P 群にも同じルールを伝える |

| 取り除いたもの | 理由 |
| --- | --- |
| コマンド、ディレクティブ、設定、baseline、診断の読み方 | ツールそのもの |
| namespace、private、shared、boundary、core といった用語 | 用語を見ると、P 群がツールの存在を推測できる。また design.md の隔離の決まりで禁じている |

既存のコードはルールに従っていないことが多い（seed は declscope を使っていないプロジェクト）。そこで rule.md には「既存のコードは課題に必要な範囲を超えて直さず、足したり変えたりするコードでルールに従う」と書いた。G 群で baseline が既存の指摘を吸収するのと同じ扱いになる。

ディレクティブの代わりに、P 群には「わざと共有する helper は、どのファイルがなぜ使うかを doc comment に書く」と伝える。G 群がディレクティブに理由を書くのと同じ判断を、コメントで表す。
