#!/usr/bin/env python3
"""Check each step's hidden tests against upstream, inside the no-network sandbox.

  python3 -I harness/verify_task.py --task tasks/<id> --clones DIR --work DIR

For every step it checks that:
- the hidden suite equals upstream's package tests at the merge commit,
- the hidden suite passes at the merge commit,
- the hidden suite fails at the parent commit, so it tests the new behaviour.

It refuses a task whose prescreen is not clean (harness/SAFETY.md). Modules
are downloaded with the network on, which runs no code. Tests run only
through sandbox-exec with the network denied. Results go to
tasks/<id>/verify.json.
"""

import argparse
import json
import shutil
import subprocess
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE))
import run  # noqa: E402


def sh(cmd, cwd, env=None, check=True):
    r = subprocess.run(cmd, cwd=cwd, env=env, capture_output=True, text=True)
    if check and r.returncode:
        raise RuntimeError(f"{cmd}: {r.stderr[-2000:]}")
    return r


def worktree(clone, sha, dest):
    if dest.exists():
        sh(["git", "-C", str(clone), "worktree", "remove", "--force", str(dest)], cwd=clone, check=False)
        shutil.rmtree(dest, ignore_errors=True)
    sh(["git", "-C", str(clone), "worktree", "add", "-q", "-f", "--detach", str(dest), sha], cwd=clone)


def test(task, step, wt, env):
    md = wt / task.get("module_dir", ".")
    sh(["go", "mod", "download"], cwd=md, env={**env, "GOPROXY": "https://proxy.golang.org,direct"}, check=False)
    pkg = md / task["package"]
    for t in pkg.glob("*_test.go"):
        t.unlink()
    hidden = run.REPO / "tasks" / task["id"] / f"step-{step.get('n')}" / "hidden"
    for src in hidden.rglob("*"):
        if src.is_file():
            dst = wt / src.relative_to(hidden)
            dst.parent.mkdir(parents=True, exist_ok=True)
            shutil.copyfile(src, dst)
    skip = "|".join(f"^{k['test']}$" for k in step.get("known_failing") or [])
    cmd = run.sandboxed(["go", "test", "-count=1", "./" + task["package"]] + (["-skip", skip] if skip else []))
    r = sh(cmd, cwd=md, env=env, check=False)
    return r.returncode == 0, (r.stdout + r.stderr)[-1500:]


def same_as_upstream(task, step, clone):
    hidden = run.REPO / "tasks" / task["id"] / f"step-{step.get('n')}" / "hidden"
    prefix = (task.get("module_dir", ".").rstrip("/") + "/") if task.get("module_dir", ".") not in (".", "") else ""
    pkgdir = prefix + task["package"]
    up = sh(["git", "-C", str(clone), "ls-tree", "--name-only", f"{step['merge']}:{pkgdir}"], cwd=clone).stdout.split()
    up_tests = {f for f in up if f.endswith("_test.go")}
    ours = {p.name for p in (hidden / pkgdir).glob("*_test.go")}
    diffs = []
    for f in sorted(up_tests | ours):
        if f not in up_tests or f not in ours:
            diffs.append(f)
            continue
        a = sh(["git", "-C", str(clone), "show", f"{step['merge']}:{pkgdir}/{f}"], cwd=clone).stdout
        if a != (hidden / pkgdir / f).read_text():
            diffs.append(f)
    return diffs


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--task", required=True)
    ap.add_argument("--clones", required=True)
    ap.add_argument("--work", required=True)
    args = ap.parse_args()
    task_dir = (run.REPO / args.task).resolve()
    run.require_prescreen(task_dir)
    task = run.load_yaml(task_dir / "task.yaml")
    task["id"] = task_dir.name
    clone = Path(args.clones) / "__".join(task["repo"].rstrip("/").removesuffix(".git").split("/")[-2:])
    work = Path(args.work)
    work.mkdir(parents=True, exist_ok=True)
    env = run.agent_env(task, "A")
    out = {"task": task_dir.name, "steps": []}
    for i, step in enumerate(task["steps"], 1):
        step.setdefault("n", i)
        rec = {"n": step["n"], "pr": step.get("pr"), "hidden_differs": same_as_upstream(task, step, clone)}
        for label, sha in (("merge", step["merge"]), ("parent", step["parent"])):
            wt = work / f"{task_dir.name}-{step['n']}-{label}"
            worktree(clone, sha, wt)
            ok, tail = test(task, step, wt, env)
            rec[label] = {"passes": ok, "tail": tail if (label == "merge") != ok else ""}
            sh(["git", "-C", str(clone), "worktree", "remove", "--force", str(wt)], cwd=clone, check=False)
        rec["ok"] = rec["merge"]["passes"] and not rec["parent"]["passes"] and not rec["hidden_differs"]
        out["steps"].append(rec)
        print(json.dumps({"task": task_dir.name, "step": step["n"], "ok": rec["ok"],
                          "merge": rec["merge"]["passes"], "parent": rec["parent"]["passes"],
                          "differs": rec["hidden_differs"]}), flush=True)
    out["ok"] = all(s["ok"] for s in out["steps"])
    (task_dir / "verify.json").write_text(json.dumps(out, indent=2, ensure_ascii=False) + "\n")


if __name__ == "__main__":
    main()
