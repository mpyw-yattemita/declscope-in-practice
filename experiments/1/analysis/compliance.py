#!/usr/bin/env python3
"""Count what declscope reports at the end of stage 1, for every arm.

  python3 -I analysis/compliance.py --root runs --clones DIR [--tasks a,b]

For each task, the seed is exported and a baseline is recorded under arm G's
config. Each run's stage 1 end is then read from its archived history, given
the same config and that baseline, and checked with `declscope` and
`declscope shrink` on the package. Arm G should report nothing. Arms A and P
show how far the stated rule, or no rule, gets without the check.

declscope only loads and type-checks packages; it runs no code of the
project. It still goes through the no-network sandbox, as all commands on
third-party trees do. Results go to <root>/<task>/<arm>/<rep>/compliance.json.
"""

import argparse
import io
import json
import re
import shutil
import subprocess
import sys
import tarfile
from pathlib import Path

REPO = Path(__file__).resolve().parent.parent
# The tasks are shared by every experiment, at the top of the repository.
ROOT = REPO.parent.parent
sys.path.insert(0, str(REPO / "harness"))
import run  # noqa: E402

WORK = Path.home() / "ws-check"
DECLSCOPE = run.TOOLS_DIR / "declscope"
RULE = re.compile(r"^\S+:\d+:\d+: (?P<msg>.*)$")


def export(clone, sha, dest):
    shutil.rmtree(dest, ignore_errors=True)
    dest.mkdir(parents=True)
    data = subprocess.run(["git", "-C", str(clone), "archive", sha], capture_output=True, check=True).stdout
    tarfile.open(fileobj=io.BytesIO(data)).extractall(dest)


def export_git_dir(git_dir, sha, dest):
    shutil.rmtree(dest, ignore_errors=True)
    dest.mkdir(parents=True)
    data = subprocess.run(["git", "--git-dir", str(git_dir), "archive", sha], capture_output=True, check=True).stdout
    tarfile.open(fileobj=io.BytesIO(data)).extractall(dest)


def declscope(args, cwd, task):
    env = run.agent_env(task, "G")
    r = subprocess.run(run.sandboxed([str(DECLSCOPE), *args]), cwd=cwd, env=env, capture_output=True, text=True)
    return r.returncode, r.stdout + r.stderr


def classify(out):
    kinds = {"boundary": 0, "qualify": 0, "surplus": 0, "unused": 0, "overexported": 0, "other": 0}
    for line in out.splitlines():
        m = RULE.match(line)
        if not m or line.startswith("\t"):
            continue
        msg = m.group("msg")
        if msg[:1].isspace():
            continue  # a "used here" line under the report above it
        if "does not carry namespace" in msg:
            kinds["qualify"] += 1
        elif "is private to" in msg or "is declared private" in msg:
            kinds["boundary"] += 1
        elif "nothing outside" in msg or "is exported, but" in msg:
            kinds["overexported"] += 1
        elif msg.startswith("unused") or "unused //declscope" in msg:
            kinds["unused"] += 1
        elif "no use from another namespace" in msg or "takes shared scope" in msg:
            kinds["surplus"] += 1
        else:
            kinds["other"] += 1
    return kinds


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="runs")
    ap.add_argument("--clones", required=True)
    ap.add_argument("--tasks")
    args = ap.parse_args()
    root = (REPO / args.root).resolve()
    tasks = args.tasks.split(",") if args.tasks else sorted(p.name for p in root.iterdir() if p.is_dir())
    for tid in tasks:
        task = run.load_yaml(ROOT / "tasks" / tid / "task.yaml")
        task["id"] = tid
        clone = Path(args.clones) / "__".join(task["repo"].rstrip("/").removesuffix(".git").split("/")[-2:])
        seed_dir = WORK / tid / "seed"
        export(clone, task["seed"], seed_dir)
        (seed_dir / ".declscope.yaml").write_text(run.DECLSCOPE_CONFIG)
        code, out = declscope(["baseline", "./..."], seed_dir, task)
        baseline = (seed_dir / ".declscope-baseline.yaml").read_text() if (seed_dir / ".declscope-baseline.yaml").exists() else None
        for arm in ("A", "P", "G"):
            for rdir in sorted((root / tid / arm).glob("*")) if (root / tid / arm).is_dir() else []:
                rj = rdir / "run.json"
                if not rj.exists():
                    continue
                r = json.loads(rj.read_text())
                git_dir = run.WS_ROOT.parent / (run.WS_ROOT.name + "-archive") / (r["workspace"] + ".git")
                if not git_dir.exists() or not r.get("stage1_end"):
                    continue
                d = WORK / tid / f"{arm}-{rdir.name}"
                export_git_dir(git_dir, r["stage1_end"], d)
                # Arm G keeps its own config and baseline. The others get the same ones.
                (d / ".declscope.yaml").write_text(run.DECLSCOPE_CONFIG)
                if baseline is not None and not (d / ".declscope-baseline.yaml").exists():
                    (d / ".declscope-baseline.yaml").write_text(baseline)
                pkg = "./" + task["package"]
                c1, o1 = declscope([pkg], d, task)
                c2, o2 = declscope(["shrink", pkg], d, task)
                # shrink stands down where it cannot see every importer; that
                # is "not judged", not zero.
                judged = not re.search(r"not judged: \S*/" + re.escape(task["package"]) + r"(:|\s|$)", o2)
                shrink = classify(o2)["overexported"] if judged else None
                res = {"arm": arm, "rep": rdir.name, "analyzer": classify(o1), "shrink": shrink,
                       "analyzer_exit": c1, "shrink_exit": c2,
                       "clean": c1 == 0 and not shrink,
                       "output": (o1 + "\n" + o2)[-6000:]}
                (rdir / "compliance.json").write_text(json.dumps(res, indent=2, ensure_ascii=False) + "\n")
                print(json.dumps({k: res[k] for k in ("arm", "analyzer", "shrink", "clean")} | {"task": tid}), flush=True)
                shutil.rmtree(d, ignore_errors=True)
        shutil.rmtree(seed_dir, ignore_errors=True)


if __name__ == "__main__":
    main()
