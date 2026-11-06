#!/usr/bin/env python3
"""Check a task's code for external downloads and process launches, without running it.

Usage:
  python3 -I harness/prescreen.py --task tasks/<id> --clones DIR

It reads files straight from git objects (`git show <commit>:<path>`), so no
code of the upstream project runs and nothing is checked out. It scans, at the
seed and at every step's merge commit, the task's package and every package of
the same module that it imports, test files included.

The verdict goes to tasks/<id>/prescreen.json:

| verdict      | meaning                                                         |
| ------------ | --------------------------------------------------------------- |
| excluded     | an installer, a download tool or a download URL was found       |
| needs-review | a process launch or an outbound HTTP call was found             |
| clean        | nothing was found, or a reviewer cleared every needs-review hit |

The harness refuses every task whose verdict is not `clean`. A task with any
external tool download is excluded, never cleared.
"""

import argparse
import os
import json
import re
import subprocess
import sys
from pathlib import Path

REPO = Path(__file__).resolve().parent.parent
# The tasks are shared by every experiment, at the top of the repository.
ROOT = REPO.parent.parent

INSTALLER = re.compile(
    r"\b(npm|npx|pnpm|yarn|bunx?|pip3?|pipx|uvx?|brew|apt-get|apk|cargo|curl|wget|"
    r"go\s+install|go\s+get|install\.sh|get\.sh)\b"
)
DOWNLOAD = re.compile(
    r"https?://(?!(localhost|127\.0\.0\.1|0\.0\.0\.0|\[::1\]|example\.(com|org|net)|"
    r"[\w.-]*\.(test|invalid|local|example))(?=[:/\"'`\s]|$))[\w.-]+[^\s\"'`]*"
    r"\.(tar\.gz|tgz|zip|sh|exe|dmg|pkg|deb|rpm|msi|whl|jar)\b"
)
EXEC = re.compile(r"\bexec\.(Command|CommandContext|LookPath)\(|\bos\.StartProcess\(|\bsyscall\.Exec\(")
HTTP_OUT = re.compile(r"\bhttp\.(Get|Post|Head|PostForm)\(|\bhttp\.DefaultClient\b|\bnet\.Dial(er|Timeout)?\b")
IMPORT_BLOCK = re.compile(r"^import\s*\((.*?)^\)", re.S | re.M)
IMPORT_LINE = re.compile(r'^import\s+(?:\w+\s+)?"([^"]+)"', re.M)
QUOTED = re.compile(r'"([^"]+)"')


def git(clone, *args):
    return subprocess.run(["git", "-C", str(clone), *args], capture_output=True, text=True, check=True).stdout


def module_path(clone, commit, module_dir):
    gomod = git(clone, "show", f"{commit}:{(Path(module_dir) / 'go.mod').as_posix()}".lstrip("./"))
    m = re.search(r"^module\s+(\S+)", gomod, re.M)
    return m.group(1)


def go_files(clone, commit, d):
    out = git(clone, "ls-tree", "--name-only", f"{commit}:{d}" if d else commit)
    return [f for f in out.splitlines() if f.endswith(".go")]


def imports(src):
    found = IMPORT_LINE.findall(src)
    for block in IMPORT_BLOCK.findall(src):
        found += QUOTED.findall(block)
    return found


def lex(src):
    """Split Go source into code text and string literals, dropping comments.

    Returns (code, strings): code keeps every character outside comments and
    literals, with literals blanked; strings is a list of (start, text).
    Offsets match src, so a position maps back to its line.
    """
    code, strings, i, n = list(src), [], 0, len(src)
    while i < n:
        c = src[i]
        if src.startswith("//", i):
            j = src.find("\n", i)
            j = n if j < 0 else j
            code[i:j] = " " * (j - i)
            i = j
        elif src.startswith("/*", i):
            j = src.find("*/", i + 2)
            j = n if j < 0 else j + 2
            code[i:j] = [ch if ch == "\n" else " " for ch in src[i:j]]
            i = j
        elif c in "\"'`":
            j = i + 1
            while j < n and src[j] != c:
                if c != "`" and src[j] == "\\":
                    j += 1
                elif c != "`" and src[j] == "\n":
                    break
                j += 1
            j = min(j + 1, n)
            strings.append((i, src[i:j]))
            code[i:j] = [ch if ch == "\n" else " " for ch in src[i:j]]
            i = j
        else:
            i += 1
    return "".join(code), strings


def line_of(src, pos):
    return src.count("\n", 0, pos) + 1


def scan_commit(clone, commit, task):
    mdir = task.get("module_dir", ".")
    mod = module_path(clone, commit, mdir)
    prefix = "" if mdir in (".", "") else mdir.rstrip("/") + "/"
    todo, seen, hits = [task["package"]], set(), []
    while todo:
        pkg = todo.pop()
        if pkg in seen:
            continue
        seen.add(pkg)
        d = prefix + pkg
        try:
            files = go_files(clone, commit, d)
        except subprocess.CalledProcessError:
            continue
        for f in files:
            path = f"{d}/{f}"
            src = git(clone, "show", f"{commit}:{path}")
            for imp in imports(src):
                if imp == mod or imp.startswith(mod + "/"):
                    todo.append(imp[len(mod):].lstrip("/") or ".")
            lines = src.splitlines()
            code, strings = lex(src)

            def add(severity, kind, pos):
                n = line_of(src, pos)
                hits.append({"severity": severity, "kind": kind, "path": path, "line": n,
                             "text": (lines[n - 1] if n <= len(lines) else "").strip()[:200]})

            # Installer names and download URLs count only inside string
            # literals: that is where a command line or a URL is written.
            # Identifiers such as a variable named uv do not count.
            for start, text in strings:
                for kind, rx in (("installer", INSTALLER), ("download-url", DOWNLOAD)):
                    for m in rx.finditer(text):
                        add("excluded", kind, start + m.start())
            for kind, rx in (("exec", EXEC), ("outbound-http", HTTP_OUT)):
                for m in rx.finditer(code):
                    add("needs-review", kind, m.start())
    return sorted(seen), hits


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--task", required=True)
    ap.add_argument("--clones", default=os.environ.get("WS_CLONES", str(Path.home() / "ws-clones")), help="directory holding upstream clones")
    args = ap.parse_args()
    task_dir = (ROOT / args.task).resolve()
    task = json.loads(subprocess.run(
        ["ruby", "-ryaml", "-rjson", "-e", "puts JSON.dump(YAML.load_file(ARGV[0]))", str(task_dir / "task.yaml")],
        capture_output=True, text=True, check=True).stdout)
    clone = Path(args.clones) / "__".join(task["repo"].rstrip("/").removesuffix(".git").split("/")[-2:])
    commits = [("seed", task["seed"])] + [(f"step-{s.get('n', i + 1)}", s["merge"]) for i, s in enumerate(task["steps"])]
    report = {"task": task["id"], "commits": {}, "hits": []}
    for label, c in commits:
        pkgs, hits = scan_commit(clone, c, task)
        report["commits"][label] = {"commit": c, "packages": pkgs}
        for h in hits:
            h["at"] = label
        report["hits"] += hits
    # One entry per distinct finding; the same line recurs at every commit.
    uniq = {}
    for h in report["hits"]:
        uniq.setdefault((h["severity"], h["kind"], h["path"], h["text"]), {**h, "at": []})["at"].append(h["at"])
    report["hits"] = sorted(uniq.values(), key=lambda h: (h["severity"] != "excluded", h["path"], h["line"]))
    sev = {h["severity"] for h in report["hits"]}
    report["verdict"] = "excluded" if "excluded" in sev else "needs-review" if sev else "clean"
    out = task_dir / "prescreen.json"
    prev = json.loads(out.read_text()) if out.exists() else {}
    if prev.get("review") and report["verdict"] == "needs-review":
        report["review"] = prev["review"]
    (out).write_text(json.dumps(report, indent=2, ensure_ascii=False) + "\n")
    counts = {}
    for h in report["hits"]:
        counts[(h["severity"], h["kind"])] = counts.get((h["severity"], h["kind"]), 0) + 1
    print(json.dumps({"task": task["id"], "verdict": report["verdict"],
                      "counts": {f"{a}/{b}": n for (a, b), n in counts.items()}}))


if __name__ == "__main__":
    sys.exit(main())
