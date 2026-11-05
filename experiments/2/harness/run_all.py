#!/usr/bin/env python3
"""Run the whole experiment one run at a time, committing as it goes.

  python3 -I harness/run_all.py --clones DIR [--tasks a,b,...]

Runs are sequential to stay under the usage limits. A finished run (its
run.json exists) is skipped, so the script can be started again after a stop.
After the three arms of a task, it runs the AI judge and the structural review
of that task. It commits and pushes after every run and every evaluation.
"""

import argparse
import datetime
import subprocess
import sys
from pathlib import Path

REPO = Path(__file__).resolve().parent.parent
# The tasks are shared by every experiment, at the top of the repository.
ROOT = REPO.parent.parent
# The four tasks experiment 1 ran. go-sdk-pr644 is left out, as it was in
# experiment 1; pass --tasks to add it.
TASKS = ["jaeger-pr9256", "hcl-pr841", "traefik-pr14130", "github-mcp-server-pr683"]


def log(msg):
    print(f"[{datetime.datetime.now():%H:%M:%S}] {msg}", flush=True)


def sh(cmd, check=True):
    return subprocess.run(cmd, cwd=REPO, check=check)


def commit(paths, msg):
    sh(["git", "add", "-A", *paths])
    if subprocess.run(["git", "diff", "--cached", "--quiet"], cwd=REPO).returncode == 0:
        return
    sh(["git", "commit", "-q", "-m", msg + "\n\nCo-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"])
    for _ in range(3):
        if subprocess.run(["git", "push", "-q"], cwd=REPO).returncode == 0:
            return
        sh(["git", "pull", "-q", "--rebase"], check=False)
    log("push failed; the commit stays local")


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--clones", required=True)
    ap.add_argument("--tasks", default=",".join(TASKS))
    args = ap.parse_args()
    py = sys.executable
    for t in args.tasks.split(","):
        for arm in ("N", "M", "R"):
            out = REPO / "runs" / t / arm / "1"
            if (out / "run.json").exists():
                log(f"{t} {arm}: done already")
                continue
            for attempt in (1, 2):
                log(f"{t} {arm}: start (attempt {attempt})")
                r = subprocess.run([py, "-I", "harness/run.py", "--task", f"tasks/{t}", "--arm", arm, "--rep", "1",
                                    "--clones", args.clones], cwd=REPO)
                if r.returncode == 0 and (out / "run.json").exists():
                    break
                log(f"{t} {arm}: failed with {r.returncode}")
            commit([f"runs/{t}/{arm}"], f"runs: {t}, arm {arm}")
            log(f"{t} {arm}: committed")
        if all((REPO / "runs" / t / a / "1" / "run.json").exists() for a in ("N", "M", "R")):
            log(f"{t}: judge")
            sh([py, "-I", "harness/judge.py", "--root", "runs", "--task", t], check=False)
            log(f"{t}: review")
            sh([py, "-I", "harness/review.py", "--root", "runs", "--task", t], check=False)
            sh([py, "-I", "analysis/diffview.py", "--root", "runs", "--out", "analysis/diffview/runs.html"], check=False)
            commit([f"runs/{t}", "analysis/diffview/runs.html"], f"runs: judge and review {t}")
            log(f"{t}: evaluated and committed")
    log("all done")


if __name__ == "__main__":
    main()
