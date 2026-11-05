#!/usr/bin/env python3
"""Build the site where colleagues judge which package is divided better.

  python3 -I analysis/judgesite.py --root runs --clones DIR --pairs analysis/human/pairs.json \\
      --out site/judge/index.html

Each pair shows two candidates, "コード 1" and "コード 2", in a GitHub-like
file browser: the package's file tree with NEW and CHANGED marks, each file in
full with line numbers, and its diff from the seed. Which arm wrote which
candidate is hidden, and the order is shuffled with a fixed seed. Every trace
of the tool is already stripped from runs/.../stage1-package.

The page keeps answers in the viewer's browser and turns them into a GitHub
issue to submit. It runs no code from the packages.
"""

import argparse
import difflib
import hashlib
import json
import subprocess
import sys
from pathlib import Path

REPO = Path(__file__).resolve().parent.parent
sys.path.insert(0, str(REPO / "harness"))
import judge  # noqa: E402
import run  # noqa: E402

TEMPLATE = Path(__file__).resolve().parent / "judgesite.template.html"
ISSUE_REPO = "mpyw-yattemita/declscope-in-practice"


def seed_text(clone, sha, path):
    r = subprocess.run(["git", "-C", str(clone), "show", f"{sha}:{path}"], capture_output=True, text=True)
    return r.stdout if r.returncode == 0 else None


def build_candidate(root, task, arm, rep, clone):
    c = judge.candidate(root, task, arm, rep)
    prefix = "" if task.get("module_dir", ".") in (".", "") else task["module_dir"].rstrip("/") + "/"
    files = []
    for f in c["files"]:
        text = (c["snap"] / f["name"]).read_text(errors="replace")
        mark = "new" if f["name"] in c["new"] else "changed" if f["name"] in c["changed"] else ""
        diff = None
        if mark:
            before = seed_text(clone, task["seed"], prefix + task["package"] + "/" + f["name"]) or ""
            diff = "".join(difflib.unified_diff(before.splitlines(True), text.splitlines(True),
                                                "seed/" + f["name"], f["name"], n=1_000_000))
        files.append({"name": f["name"], "lines": f["lines"], "test": f["test"], "mark": mark,
                      "decls": f.get("decls") or [], "text": text, "diff": diff})
    return {"files": files}


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="runs")
    ap.add_argument("--clones", required=True)
    ap.add_argument("--pairs", required=True, help="JSON list of {task, x, y}, fixed before the AI results are read")
    ap.add_argument("--out", required=True)
    args = ap.parse_args()
    root = (REPO / args.root).resolve()
    spec = json.loads((REPO / args.pairs).read_text())
    pairs = spec["pairs"]
    salt_path = REPO / "analysis" / "human" / ".salt"
    salt = salt_path.read_text().strip()
    if hashlib.sha256(salt.encode()).hexdigest() != spec["salt_sha256"]:
        raise SystemExit("the salt does not match the hash committed in pairs.json")
    out_pairs = []
    for i, p in enumerate(pairs, 1):
        # A pair whose task was not run (go-sdk, after the run was cut short)
        # is skipped; its number stays, so answers still match key.json.
        if not all((root / p["task"] / s / "stage1-package").is_dir() for s in (p["x"], p["y"])):
            print(f"pair-{i}: {p['task']} was not run; skipped")
            continue
        task = run.load_yaml(REPO / "tasks" / p["task"] / "task.yaml")
        task["id"] = p["task"]
        clone = Path(args.clones) / "__".join(task["repo"].rstrip("/").removesuffix(".git").split("/")[-2:])
        sides = [p["x"], p["y"]]
        # The order comes from a salt kept out of git, so nobody can work out
        # which side is which arm. pairs.json carries the salt's hash, and the
        # salt is published after the judging, which shows the order was fixed.
        if int(hashlib.sha256(f"{salt}|{p['task']}|{p['x']}|{p['y']}".encode()).hexdigest(), 16) % 2:
            sides.reverse()
        cands = [build_candidate(root, task, *s.split("/"), clone) for s in sides]
        out_pairs.append({
            "id": f"pair-{i}", "task": p["task"], "package": task["package"],
            "features": [s.get("title") for s in task["steps"] if s.get("stage", 1) == 1],
            "candidates": cands,
        })
        key_path = REPO / "analysis" / "human" / "key.json"
        keys = json.loads(key_path.read_text()) if key_path.exists() else {}
        keys[f"pair-{i}"] = {"task": p["task"], "1": sides[0], "2": sides[1]}
        key_path.parent.mkdir(parents=True, exist_ok=True)
        key_path.write_text(json.dumps(keys, indent=2, ensure_ascii=False) + "\n")
    data = {"pairs": out_pairs, "issueRepo": ISSUE_REPO}
    html = TEMPLATE.read_text().replace("/*__DATA__*/null", json.dumps(data, ensure_ascii=False).replace("</", "<\\/"))
    Path(args.out).parent.mkdir(parents=True, exist_ok=True)
    Path(args.out).write_text(html)
    print(f"{args.out}: {len(out_pairs)} pair(s), {len(html) // 1024} KiB")


if __name__ == "__main__":
    main()
