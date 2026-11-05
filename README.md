# declscope in practice

Does [declscope](https://github.com/mpyw/declscope) help an AI agent write new code that is easier to change?

> [!NOTE]
> This experiment is in progress. No results yet.

An agent implements real pull requests from open-source Go projects in three ways. It gets no rule, the rule in its prompt, or the rule in its prompt with declscope checking it. A second agent then makes the next real change to each result, without declscope. A blind reviewer also looks for refactoring the result still needs. [design.md](design.md) has the details.
