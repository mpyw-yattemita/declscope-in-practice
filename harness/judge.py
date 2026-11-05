#!/usr/bin/env python3
"""Blind pairwise judgement of how well the new code is divided (design.md, 分割の見比べ).

  python3 -I harness/judge.py --root runs --task <id> [--model claude-opus-5-5]

For each pair of arms (A-P, P-G, A-G) and each pair of their runs, it asks a
judge twice, with the two candidates swapped. Each candidate is the package as
stage 1 left it (runs/.../stage1-package), stripped of every trace of the tool.
The judge gets a map of every file and the full text of the new and changed
files, and may read any other file of the two candidates. Answers that
disagree across the swap count as a tie.

Results go to <root>/<task>/judge.json.
"""

import argparse
import itertools
import json
import re
import shutil
import subprocess
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE))
import run  # noqa: E402

PKGMAP = Path.home() / "ws-tools" / "pkgmap"
PAIRS = [("A", "P"), ("P", "G"), ("A", "G")]

PROMPT = """\
Two versions of the same Go package are in this directory, as `x/` and `y/`.
Both started from the same code and gained the same features, written by
different people. Compare how well each package is divided into files and types.

Judge the package as a whole, and how the new code fits into it:
- Does each file and each type have one clear job?
- Is closely related code kept together?
- Do the boundaries between files follow the boundaries between features?
- Does the new code fit the structure the package already had?

Behaviour and tests are known to be equivalent. Do not judge style, comments,
or naming conventions on their own, except where a name shows how the code is
divided.

Below is a map of each package: every file with its line count and its
top-level declarations. Files marked NEW were added and files marked CHANGED
were modified. The full text of every NEW and CHANGED file follows the maps.
You may read any other file under x/ or y/ if you need to.

{maps}

{texts}

Answer with one JSON object and nothing else:
{{"better": "x" | "y" | "tie", "reasons": ["...", "..."]}}
Give two to five short reasons, each pointing at concrete files or types.
"""


def changed_files(patch_text, pkg):
    new, changed = set(), set()
    for m in re.finditer(r"^diff --git a/(\S+) b/(\S+)\n((?:(?!diff --git).*\n)*)", patch_text, re.M):
        path, body = m.group(2), m.group(3)
        if not path.startswith(pkg + "/") or "/" in path[len(pkg) + 1:]:
            continue
        name = path[len(pkg) + 1:]
        (new if "new file mode" in body else changed).add(name)
    return new, changed


def render_map(label, files, new, changed):
    lines = [f"### {label}/", "", "| file | lines | mark | declarations |", "| --- | --- | --- | --- |"]
    for f in files:
        mark = "NEW" if f["name"] in new else "CHANGED" if f["name"] in changed else ""
        decls = ", ".join(
            (f"{d['recv']}.{d['name']}" if d["kind"] == "method" else d["name"]) + f" ({d['kind']})"
            for d in f.get("decls") or [])
        lines.append(f"| {f['name']} | {f['lines']} | {mark} | {decls} |")
    return "\n".join(lines)


def candidate(root, task, arm, rep):
    d = root / task["id"] / arm / rep
    snap = d / "stage1-package"
    patch = (d / "stage1.patch").read_text()
    new, changed = changed_files(patch, task["package"])
    files = json.loads(subprocess.run([str(PKGMAP), str(snap)], capture_output=True, text=True, check=True).stdout)
    return {"arm": arm, "rep": rep, "snap": snap, "new": new, "changed": changed, "files": files}


def ask(c1, c2, model, out_dir, label):
    ws = run.new_dir(run.WS_ROOT)
    try:
        for name, c in (("x", c1), ("y", c2)):
            shutil.copytree(c["snap"], ws / name)
        maps = "\n\n".join(render_map(n, c["files"], c["new"], c["changed"]) for n, c in (("x", c1), ("y", c2)))
        texts = []
        for n, c in (("x", c1), ("y", c2)):
            for f in sorted(c["new"] | c["changed"]):
                p = ws / n / f
                if p.exists():
                    texts.append(f"#### {n}/{f}\n\n```go\n{p.read_text()}\n```")
        prompt = PROMPT.format(maps=maps, texts="\n\n".join(texts))
        env = run.agent_env({"go": None}, "A", ws=ws)
        summary = run.run_agent_waiting(prompt, ws, env, out_dir, model, label)
        text = ""
        import gzip
        with gzip.open(out_dir / f"{label}.jsonl.gz", "rt") as f:
            for line in f:
                try:
                    ev = json.loads(line)
                except json.JSONDecodeError:
                    continue
                if ev.get("type") == "result":
                    text = ev.get("result") or ""
        m = re.search(r"\{.*\}", text, re.S)
        verdict = json.loads(m.group(0)) if m else {"better": None, "reasons": [text[:500]]}
        verdict["audit"] = run.audit(out_dir / f"{label}.jsonl.gz", ws, "A", start=ws)
        verdict["agent"] = summary
        return verdict
    finally:
        shutil.rmtree(ws, ignore_errors=True)


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="runs")
    ap.add_argument("--task", required=True)
    ap.add_argument("--model", default="claude-opus-5-5")
    args = ap.parse_args()
    root = (run.REPO / args.root).resolve()
    task = run.load_yaml(run.REPO / "tasks" / args.task / "task.yaml")
    task["id"] = args.task
    reps = {a: sorted(p.name for p in (root / args.task / a).iterdir() if (p / "stage1-package").is_dir())
            for a in ("A", "P", "G") if (root / args.task / a).is_dir()}
    out_path = root / args.task / "judge.json"
    result = json.loads(out_path.read_text()) if out_path.exists() else {"task": args.task, "comparisons": []}
    done = {(c["x"], c["y"]) for c in result["comparisons"]}
    jdir = root / args.task / "judge"
    jdir.mkdir(exist_ok=True)
    for a, b in PAIRS:
        for ra, rb in itertools.product(reps.get(a, []), reps.get(b, [])):
            key = (f"{a}/{ra}", f"{b}/{rb}")
            if key in done:
                continue
            ca, cb = candidate(root, task, a, ra), candidate(root, task, b, rb)
            v1 = ask(ca, cb, args.model, jdir, f"{a}{ra}-{b}{rb}-ab")
            v2 = ask(cb, ca, args.model, jdir, f"{a}{ra}-{b}{rb}-ba")
            # Map each answer back to the arms; a disagreement across the swap is a tie.
            w1 = {"x": a, "y": b, "tie": "tie"}.get(v1.get("better"))
            w2 = {"x": b, "y": a, "tie": "tie"}.get(v2.get("better"))
            winner = w1 if w1 == w2 else "tie"
            result["comparisons"].append({"x": key[0], "y": key[1], "winner": winner,
                                          "first": v1, "second": v2})
            out_path.write_text(json.dumps(result, indent=2, ensure_ascii=False) + "\n")
            print(json.dumps({"task": args.task, "pair": key, "winner": winner,
                              "first": v1.get("better"), "second": v2.get("better")}), flush=True)


if __name__ == "__main__":
    main()
