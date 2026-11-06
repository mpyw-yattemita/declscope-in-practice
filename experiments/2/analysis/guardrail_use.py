#!/usr/bin/env python3
"""Record whether the deterministic check (declscope) helped, for every step
of the arms that had it.

  python3 -I analysis/guardrail_use.py --root runs

It reads the transcripts afterwards and asks the agent nothing, so the arms
stay the same. For each step of stage 1 in arms MD and RD, it finds every
shell command that ran the analyzer or `declscope shrink`, and reads the
diagnostics in its output. Type errors that declscope prints for a package
it cannot load are counted apart, and are no finding. Results go to
<root>/<task>/<arm>/<rep>/guardrail.json.

| Verdict | Meaning |
| --- | --- |
| helped | A check reported a finding, and the last check had none: the agent changed its code after a report |
| left | The last check still reported something |
| clean | Every check was clean: nothing for the check to catch |
| not run | The agent never ran a check |

What the code would look like without the check is not in the transcript.
The comparison of M and MD, and of R and RD, in compliance.py answers that.
"""

import argparse
import gzip
import json
import re
from pathlib import Path

HERE = Path(__file__).resolve().parent
REPO = HERE.parent
GUARDED = ("MD", "RD")

DIAG = re.compile(r"^\S+\.go:\d+:\d+: (?P<msg>\S.*)$")
# A command that runs the analyzer or shrink. -V, baseline, inspect, survey
# and skill do not check the code.
CHECK = re.compile(r"(^|[\s;&|(])declscope(\s+-\S+)*\s+(shrink\b|\./|\S+\.\.\.)")


def rule_of(msg):
    if "does not carry namespace" in msg:
        return "qualify"
    if "is private to" in msg or "is declared private" in msg:
        return "boundary"
    if "nothing outside" in msg or "is exported, but" in msg:
        return "overexported"
    if msg.startswith("unused") or "unused //declscope" in msg:
        return "unused"
    if "no use from another namespace" in msg or "takes shared scope" in msg:
        return "surplus"
    # Anything else in that shape is the compiler's: declscope prints the type
    # errors of a package it cannot load. They are no finding of the rule.
    return "type_error"


def text_of(content):
    if isinstance(content, str):
        return content
    if isinstance(content, list):
        return "\n".join(c.get("text", "") for c in content if isinstance(c, dict))
    return ""


def checks_in(transcript):
    """The checks the agent ran, in order: the command and its diagnostics."""
    uses, out = {}, []
    with gzip.open(transcript, "rt") as f:
        for line in f:
            try:
                e = json.loads(line)
            except json.JSONDecodeError:
                continue
            m = e.get("message")
            if not isinstance(m, dict):
                continue
            for b in m.get("content") or []:
                if not isinstance(b, dict):
                    continue
                if b.get("type") == "tool_use" and b.get("name") == "Bash":
                    cmd = (b.get("input") or {}).get("command", "")
                    if CHECK.search(cmd):
                        uses[b["id"]] = cmd
                elif b.get("type") == "tool_result" and b.get("tool_use_id") in uses:
                    diags = [d.group("msg") for d in map(DIAG.match, text_of(b.get("content")).splitlines()) if d]
                    out.append({"command": uses[b["tool_use_id"]][:300], "diagnostics": diags})
    return out


def skill_read(transcript):
    with gzip.open(transcript, "rt") as f:
        return any('"declscope-authoring"' in line and '"Skill"' in line for line in f)


def findings(diags):
    return [d for d in diags if rule_of(d) != "type_error"]


def verdict(checks):
    if not checks:
        return "not run"
    if findings(checks[-1]["diagnostics"]):
        return "left"
    if any(findings(c["diagnostics"]) for c in checks):
        return "helped"
    return "clean"


def step_record(sdir):
    t = sdir / "agent.jsonl.gz"
    checks = checks_in(t)
    seen = sorted({d for c in checks for d in findings(c["diagnostics"])})
    by_rule = {}
    for d in seen:
        by_rule[rule_of(d)] = by_rule.get(rule_of(d), 0) + 1
    type_errors = len({d for c in checks for d in c["diagnostics"] if rule_of(d) == "type_error"})
    diff = (sdir / "diff.patch").read_text() if (sdir / "diff.patch").exists() else ""
    return {
        "step": int(sdir.name.split("-")[1]),
        "verdict": verdict(checks),
        "checks": len(checks),
        "first": len(findings(checks[0]["diagnostics"])) if checks else None,
        "last": len(findings(checks[-1]["diagnostics"])) if checks else None,
        "type_errors": type_errors,
        "distinct": len(seen),
        "by_rule": by_rule,
        "skill_read": skill_read(t),
        "directives_added": len(re.findall(r"^\+.*//declscope:(shared|private|namespace|core)", diff, re.M)),
        "ignores_added": len(re.findall(r"^\+.*//declscope:ignore", diff, re.M)),
        "reported": seen,
    }


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="runs")
    args = ap.parse_args()
    root = (REPO / args.root).resolve()
    for rj in sorted(root.glob("*/*/*/run.json")):
        rdir = rj.parent
        if rdir.parent.name not in GUARDED:
            continue
        run = json.loads(rj.read_text())
        stage1 = {s["step"] for s in run["steps"] if s["stage"] == 1}
        steps = [step_record(rdir / f"step-{n}") for n in sorted(stage1) if (rdir / f"step-{n}" / "agent.jsonl.gz").exists()]
        res = {"task": rdir.parent.parent.name, "arm": rdir.parent.name, "rep": rdir.name, "steps": steps}
        (rdir / "guardrail.json").write_text(json.dumps(res, indent=2, ensure_ascii=False) + "\n")
        print(json.dumps({k: res[k] for k in ("task", "arm", "rep")} | {"verdicts": [s["verdict"] for s in steps]}))


if __name__ == "__main__":
    main()
