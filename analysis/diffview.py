#!/usr/bin/env python3
"""Build a page that lays the arms' diffs side by side, task by task and step by step.

  python3 -I analysis/diffview.py --root runs --out analysis/diffview.html
  python3 -I analysis/diffview.py --root smoke --out <somewhere>/diffview.html

It reads <root>/<task>/<arm>/<rep>/step-<n>/{diff.patch,metrics.json} and the
task's task.yaml for the step titles. The page is one self-contained file.
"""

import argparse
import json
import subprocess
from pathlib import Path

REPO = Path(__file__).resolve().parent.parent
TEMPLATE = Path(__file__).resolve().parent / "diffview.template.html"
MAX_DIFF = 400_000  # characters per step and run; longer diffs are cut and say so


def load_yaml(path):
    out = subprocess.run(
        ["ruby", "-ryaml", "-rjson", "-e", "puts JSON.dump(YAML.load_file(ARGV[0]))", str(path)],
        check=True, capture_output=True, text=True).stdout
    return json.loads(out)


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="runs")
    ap.add_argument("--out", required=True)
    args = ap.parse_args()
    root = (REPO / args.root).resolve()
    tasks = []
    for tdir in sorted(p for p in root.iterdir() if p.is_dir() and not p.name.startswith(".")):
        ty = REPO / "tasks" / tdir.name / "task.yaml"
        meta = load_yaml(ty) if ty.exists() else {"steps": []}
        steps = {}
        for arm_dir in sorted(p for p in tdir.iterdir() if p.is_dir()):
            for rep_dir in sorted(p for p in arm_dir.iterdir() if p.is_dir() and not p.name.startswith(".")):
                for sdir in sorted(rep_dir.glob("step-*")):
                    n = int(sdir.name.split("-")[1])
                    diff_p, met_p = sdir / "diff.patch", sdir / "metrics.json"
                    if not diff_p.exists():
                        continue
                    diff = diff_p.read_text(errors="replace")
                    cut = len(diff) > MAX_DIFF
                    m = json.loads(met_p.read_text()) if met_p.exists() else {}
                    score, agent = m.get("score") or {}, m.get("agent") or {}
                    steps.setdefault(n, []).append({
                        "arm": arm_dir.name, "rep": rep_dir.name,
                        "stage": m.get("stage"), "setup": m.get("arm_setup"),
                        "success": score.get("success"), "failed": score.get("failed", []),
                        "no_counterpart": score.get("no_counterpart", []), "renamed": score.get("renamed"),
                        "tokens": agent.get("tokens"), "turns": agent.get("num_turns"),
                        "files": m.get("files_touched"), "lines": m.get("lines_touched"),
                        "directives": m.get("directives"), "audit_ok": (m.get("audit") or {}).get("ok"),
                        "diff": diff[:MAX_DIFF], "cut": cut,
                    })
        if not steps:
            continue
        info = {s.get("n", i + 1): s for i, s in enumerate(meta.get("steps", []))}
        tasks.append({
            "id": tdir.name, "repo": meta.get("repo"), "package": meta.get("package"),
            "steps": [{"n": n, "pr": info.get(n, {}).get("pr"), "title": info.get(n, {}).get("title"),
                       "stage": info.get(n, {}).get("stage"), "runs": runs}
                      for n, runs in sorted(steps.items())],
        })
    html = TEMPLATE.read_text().replace("/*__DATA__*/null", json.dumps({"tasks": tasks}, ensure_ascii=False)
                                        .replace("</", "<\\/"))
    Path(args.out).write_text(html)
    print(f"{args.out}: {len(tasks)} task(s), "
          f"{sum(len(s['runs']) for t in tasks for s in t['steps'])} step run(s)")


if __name__ == "__main__":
    main()
