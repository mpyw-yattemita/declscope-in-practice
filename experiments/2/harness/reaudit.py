#!/usr/bin/env python3
"""Re-run the audit over finished runs, after the audit itself was fixed.

  python3 -I harness/reaudit.py --root runs

Only runs with a run.json are touched. Each step's metrics.json and the
run.json get the new audit.
"""

import argparse
import json
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
import run  # noqa: E402


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="runs")
    args = ap.parse_args()
    for rj in sorted((run.REPO / args.root).glob("*/*/*/run.json")):
        r = json.loads(rj.read_text())
        ws = run.WS_ROOT / r["workspace"]
        for s in r["steps"]:
            sdir = rj.parent / f"step-{s['step']}"
            a = run.audit(sdir / "agent.jsonl.gz", ws, s["arm_setup"], allowed_extra=[run.run_tmp(ws)], start=ws)
            s["audit"] = a
            mp = sdir / "metrics.json"
            m = json.loads(mp.read_text())
            m["audit"] = a
            mp.write_text(json.dumps(m, indent=2, ensure_ascii=False))
            if not a["ok"]:
                print(rj.parent, s["step"], a["issues"][:3])
        rj.write_text(json.dumps(r, indent=2, ensure_ascii=False))
    print("re-audited")


if __name__ == "__main__":
    main()
