# declscope in practice

[declscope](https://github.com/mpyw/declscope) は、AI エージェントが変更しやすい新しいコードを書く助けになるか？

> [!NOTE]
> 実験は進行中で、まだ結果はない。

オープンソースの Go プロジェクトで実際にマージされた PR を、エージェントに3通りの条件で実装させる。ルールを渡さない、ルールをプロンプトで渡す、ルールをプロンプトで渡したうえで declscope に検査させる、の3通り。続いて、別のエージェントが、それぞれの結果に実際の次の変更を加える。このエージェントは declscope を使わない。あわせて、ブラインドのレビュアーが、結果のコードにまだ必要なリファクタリングを探す。

詳細は [design.md](design.md) にある。進み具合は [Epic #12](https://github.com/mpyw-yattemita/declscope-in-practice/issues/12) で追う。
