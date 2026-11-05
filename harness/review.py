#!/usr/bin/env python3
"""Blind structural review of each run's stage 1 result (design.md, 構造レビュー). Exploratory.

  python3 -I harness/review.py --root runs --task <id> [--model claude-opus-5-5]

A reviewer reads the package as stage 1 left it (the same two layers the judge
gets) and lists refactoring that should be done now, freely, with a file, a
line and a fix. A second agent checks each finding and sorts the ones it keeps
into the categories of design.md. The reviewer never sees the categories.

One deviation from design.md: the check and the sorting are done by one agent,
not two, to save runs. The reviewer is still kept apart from the categories.

Results go to <root>/<task>/<arm>/<rep>/review.json.
"""

import argparse
import gzip
import json
import re
import shutil
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE))
import judge  # noqa: E402
import run  # noqa: E402

REVIEW = """\
This directory holds one Go package, `pkg/`. Some of its files were just added
or changed; they are marked in the map below, and their full text follows it.
You may read any other file under pkg/.

Review the package as a code reviewer would. List the refactoring that should
be done now. Focus on the new and changed code and how it fits the package.

{map}

{texts}

Answer with one JSON object and nothing else:
{{"findings": [{{"file": "...", "line": 0, "problem": "...", "fix": "..."}}]}}
Give each finding a concrete file, line and fix. Leave out vague remarks.
An empty list is a valid answer.
"""

CHECK = """\
This directory holds one Go package, `pkg/`. A reviewer listed the findings
below. Check each one against the code.

{findings}

For each finding, decide whether it is real and worth fixing now. For each
kept finding, give one category:

- naming: a name does not match its role, or one idea has two names
- duplication: the same logic is written twice
- over-abstraction: a shared helper couples callers that are not related
- mixed-job: one function or type does two unrelated jobs
- dead: a definition is unused
- other: none of the above

Answer with one JSON object and nothing else:
{{"verdicts": [{{"index": 0, "keep": true, "category": "naming", "why": "..."}}]}}
"""


def result_text(path):
    text = ""
    with gzip.open(path, "rt") as f:
        for line in f:
            try:
                ev = json.loads(line)
            except json.JSONDecodeError:
                continue
            if ev.get("type") == "result":
                text = ev.get("result") or ""
    m = re.search(r"\{.*\}", text, re.S)
    try:
        return json.loads(m.group(0)) if m else None
    except json.JSONDecodeError:
        return None


def changed_lines(patch, pkg):
    """New-file line numbers each changed file gained, from a stage 1 patch."""
    out, cur, n = {}, None, 0
    for line in patch.splitlines():
        m = re.match(r"^\+\+\+ b/(.*)", line)
        if m:
            p = m.group(1)
            cur = p[len(pkg) + 1:] if p.startswith(pkg + "/") and "/" not in p[len(pkg) + 1:] else None
            continue
        h = re.match(r"^@@ -\d+(?:,\d+)? \+(\d+)", line)
        if h:
            n = int(h.group(1)) - 1
            continue
        if cur is None or line.startswith("---"):
            continue
        if line.startswith("+"):
            n += 1
            out.setdefault(cur, set()).add(n)
        elif not line.startswith("-"):
            n += 1
    return out


def review_one(root, task, arm, rep, model):
    d = root / task["id"] / arm / rep
    c = judge.candidate(root, task, arm, rep)
    ws = run.new_dir(run.WS_ROOT)
    try:
        shutil.copytree(c["snap"], ws / "pkg")
        texts = "\n\n".join(f"#### pkg/{f}\n\n```go\n{(ws / 'pkg' / f).read_text()}\n```"
                            for f in sorted(c["new"] | c["changed"]) if (ws / "pkg" / f).exists())
        env = run.agent_env({"go": None}, "A", ws=ws)
        out_dir = d / "review"
        prompt = REVIEW.format(map=judge.render_map("pkg", c["files"], c["new"], c["changed"]), texts=texts)
        s1 = run.run_agent_waiting(prompt, ws, env, out_dir, model, "reviewer")
        found = (result_text(out_dir / "reviewer.jsonl.gz") or {}).get("findings", [])
        listing = "\n".join(f"{i}. {f.get('file')}:{f.get('line')}: {f.get('problem')} Fix: {f.get('fix')}"
                            for i, f in enumerate(found)) or "(none)"
        verdicts = []
        if found:
            run.run_agent_waiting(CHECK.format(findings=listing), ws, env, out_dir, model, "checker")
            verdicts = (result_text(out_dir / "checker.jsonl.gz") or {}).get("verdicts", [])
        audits = [run.audit(out_dir / f"{n}.jsonl.gz", ws, "A", start=ws)
                  for n in ("reviewer", "checker") if (out_dir / f"{n}.jsonl.gz").exists()]
    finally:
        shutil.rmtree(ws, ignore_errors=True)
    lines = changed_lines((d / "stage1.patch").read_text(), task["package"])
    kept = []
    for v in verdicts:
        i = v.get("index")
        if v.get("keep") and isinstance(i, int) and 0 <= i < len(found):
            f = found[i]
            name = Path(str(f.get("file", ""))).name
            on_changed = name in lines and (int(f.get("line") or 0) in lines[name] or name in c["new"])
            kept.append({**f, "category": v.get("category"), "why": v.get("why"), "on_changed_lines": on_changed})
    added = sum(len(v) for v in lines.values())
    counted = [k for k in kept if k["on_changed_lines"]]
    res = {"arm": arm, "rep": rep, "findings": found, "kept": kept,
           "counted": len(counted), "changed_lines": added,
           "per_kloc": round(1000 * len(counted) / added, 2) if added else None,
           "by_category": {cat: sum(1 for k in counted if k["category"] == cat)
                           for cat in ("naming", "duplication", "over-abstraction", "mixed-job", "dead", "other")},
           "audit_ok": all(a["ok"] for a in audits), "reviewer_agent": s1}
    (d / "review.json").write_text(json.dumps(res, indent=2, ensure_ascii=False) + "\n")
    return res


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="runs")
    ap.add_argument("--task", required=True)
    ap.add_argument("--model", default="claude-opus-5-5")
    args = ap.parse_args()
    root = (run.REPO / args.root).resolve()
    task = run.load_yaml(run.REPO / "tasks" / args.task / "task.yaml")
    task["id"] = args.task
    for arm in ("A", "P", "G"):
        adir = root / args.task / arm
        if not adir.is_dir():
            continue
        for rep in sorted(p.name for p in adir.iterdir() if (p / "stage1-package").is_dir()):
            if (adir / rep / "review.json").exists():
                continue
            r = review_one(root, task, arm, rep, args.model)
            print(json.dumps({"task": args.task, "arm": arm, "rep": rep, "counted": r["counted"],
                              "per_kloc": r["per_kloc"], "by_category": r["by_category"]}), flush=True)


if __name__ == "__main__":
    main()
