#!/usr/bin/env python3
"""Run one chain of one task for one arm, and collect the results.

Usage:
  python3 -I harness/run.py --task tasks/<id> --arm A|P|G --rep N --clones DIR
                            [--steps 1,2] [--model claude-opus-5-5]

See harness/README.md for what each phase does.
"""

import argparse
import datetime
import gzip
import json
import os
import re
import secrets
import shutil
import subprocess
import sys
import time
from pathlib import Path

HERE = Path(__file__).resolve().parent
REPO = HERE.parent
PROMPTS = HERE / "prompts"

# Workspaces live outside this repository, under a path that names nothing.
WS_ROOT = Path(os.environ.get("WS_ROOT", Path.home() / "ws"))
TOOLS_DIR = Path(os.environ.get("WS_TOOLS", Path.home() / "ws-tools" / "bin"))
DECLSCOPE_SKILLS = Path(
    os.environ.get("DECLSCOPE_SKILLS", Path.home() / "github.com/mpyw/declscope/skills")
)
CLAUDE = os.environ.get("CLAUDE_BIN", shutil.which("claude") or "claude")
# Every command that runs third-party code goes through this profile, which
# denies outbound network except localhost. See harness/SAFETY.md.
NONET = Path(os.environ.get("NONET_PROFILE", Path.home() / ".claude" / "hooks" / "nonet.sb"))


def sandboxed(cmd):
    if not NONET.is_file():
        raise RuntimeError(f"{NONET} is missing; refusing to run third-party code")
    return ["sandbox-exec", "-f", str(NONET), *cmd]
AGENT_TIMEOUT = int(os.environ.get("AGENT_TIMEOUT", 3600))


def load_yaml(path):
    # PyYAML is not installed here; ruby ships YAML in its standard library.
    out = subprocess.run(
        ["ruby", "-ryaml", "-rjson", "-e", "puts JSON.dump(YAML.load_file(ARGV[0]))", str(path)],
        check=True, capture_output=True, text=True,
    ).stdout
    return json.loads(out)


def sh(cmd, cwd=None, env=None, check=True, capture=True, timeout=None, input=None):
    r = subprocess.run(
        cmd, cwd=cwd, env=env, capture_output=capture, text=True, timeout=timeout, input=input,
        shell=isinstance(cmd, str),
    )
    if check and r.returncode != 0:
        raise RuntimeError(f"command failed ({r.returncode}): {cmd}\n{r.stdout}\n{r.stderr}")
    return r


def go_root():
    return Path(sh(["go", "env", "GOROOT"]).stdout.strip())


def mod_cache():
    return Path(sh(["go", "env", "GOMODCACHE"]).stdout.strip())


def run_tmp(ws):
    """A temporary directory of this run alone, outside its workspace."""
    d = WS_ROOT.parent / (WS_ROOT.name + "-tmp") / ws.name
    d.mkdir(parents=True, exist_ok=True)
    return d


def agent_env(task, arm, offline=True, ws=None):
    """The environment of the agent under measurement.

    PATH is minimal on purpose. The user's shell profile activates mise
    only when mise is on PATH, and mise would put declscope there.
    """
    path = [str(go_root() / "bin"), "/usr/bin", "/bin", "/usr/sbin", "/sbin"]
    if arm == "G":
        path.insert(0, str(TOOLS_DIR))
    env = {
        "HOME": str(Path.home()),
        "USER": os.environ.get("USER", ""),
        "LOGNAME": os.environ.get("LOGNAME", ""),
        "SHELL": os.environ.get("SHELL", "/bin/zsh"),
        "TERM": "dumb",
        "LANG": "en_US.UTF-8",
        "PATH": ":".join(path),
        "GOTOOLCHAIN": "go" + task["go"] if task.get("go") else "local",
        "GOSUMDB": "off",
    }
    if offline:
        env["GOPROXY"] = "off"
    if ws is not None:
        env["TMPDIR"] = str(run_tmp(ws)) + "/"
    for k, v in (task.get("env") or {}).items():
        env[k] = str(v)
    return env


def new_dir(root):
    root.mkdir(parents=True, exist_ok=True)
    while True:
        d = root / secrets.token_hex(6)
        if not d.exists():
            d.mkdir()
            return d


def ensure_ws_root():
    WS_ROOT.mkdir(parents=True, exist_ok=True)
    shutil.copyfile(HERE / "workspace-AGENTS.md", WS_ROOT / "AGENTS.md")


def git(ws, *args, check=True):
    return sh(["git", "-c", "user.name=dev", "-c", "user.email=dev@example.com", *args], cwd=ws, check=check)


def commit_all(ws, msg):
    git(ws, "add", "-A")
    git(ws, "commit", "-q", "--allow-empty", "-m", msg)
    return git(ws, "rev-parse", "HEAD").stdout.strip()


def export_seed(task, clones, ws):
    clone = clones / task["clone"]
    p1 = subprocess.Popen(["git", "-C", str(clone), "archive", task["seed"]], stdout=subprocess.PIPE)
    p2 = subprocess.run(["tar", "-x", "-C", str(ws)], stdin=p1.stdout, capture_output=True)
    p1.wait()
    if p1.returncode or p2.returncode:
        raise RuntimeError(f"git archive failed: {p2.stderr}")
    git(ws, "init", "-q", "-b", "main")


def module_dir(task, ws):
    return ws / task.get("module_dir", ".")


def warm(task, ws):
    """Fetch modules with the network on, without running any code.

    `go test -run '^$'` must never be used here: it still starts the test
    binary, so TestMain and init run with the network on. That is how a
    third-party test once downloaded software that Microsoft Defender flagged.
    `go mod download` and `go list` only fetch and read.
    """
    env = agent_env(task, "A", offline=False)
    md = module_dir(task, ws)
    # Warming must not edit go.mod or go.sum: upstream's patches assume
    # upstream's files. -mod=mod once let go list add lines to go.sum.
    before = {f: (md / f).read_bytes() for f in ("go.mod", "go.sum") if (md / f).exists()}
    sh(["go", "mod", "download"], cwd=md, env=env)
    sh(["go", "list", "-e", "-deps", "-test", "./" + task["package"]], cwd=md, env=env)
    after = {f: (md / f).read_bytes() for f in before}
    if after != before:
        raise RuntimeError("warming changed " + ", ".join(f for f in before if before[f] != after[f]))


# The configuration arm G runs under. The naming rule is on, for exported
# names too, so the stated rule in rule.md covers names as well.
DECLSCOPE_CONFIG = """\
rules:
  naming:
    qualify: ondemand
    exported: true
"""


def setup_guardrail(task, ws):
    env = agent_env(task, "G")
    md = module_dir(task, ws)
    (md / ".declscope.yaml").write_text(DECLSCOPE_CONFIG)
    sh(["declscope", "baseline", "./..."], cwd=md, env=env)
    skills = ws / ".claude" / "skills"
    skills.mkdir(parents=True, exist_ok=True)
    for name in ("declscope-authoring", "declscope-adoption"):
        shutil.copytree(DECLSCOPE_SKILLS / name, skills / name)


def prompt_for(task_dir, step, arm, task):
    pkg = task["package"]
    text = (task_dir / step["task"]).read_text()
    parts = [text.strip()]
    if step.get("given"):
        parts.append(
            "The following files are part of the specification and are already in place:\n"
            + "\n".join(f"- {g}" for g in step["given"])
        )
    if arm in ("P", "G"):
        parts.append((PROMPTS / "rule.md").read_text().strip())
    if arm == "G":
        parts.append((PROMPTS / "guardrail.md").read_text().strip())
    checks = ["go vet ./...", f"go test ./{pkg}"]
    if arm == "G":
        checks += ["declscope ./...", "declscope shrink ./..."]
    parts.append(
        (PROMPTS / "self-review.md").read_text().strip()
        + "\n\n```bash\n" + "\n".join(checks) + "\n```"
    )
    return "\n\n".join(parts) + "\n"


def run_agent(prompt, ws, env, out_dir, model, label):
    out_dir.mkdir(parents=True, exist_ok=True)
    # The prompt goes on stdin: a judge's prompt can exceed the argument limit.
    cmd = [
        CLAUDE, "-p",
        "--output-format", "stream-json", "--verbose",
        "--permission-mode", "bypassPermissions",
        "--disallowedTools", "WebFetch", "WebSearch",
        "--strict-mcp-config",
        "--model", model,
    ]
    t0 = time.time()
    r = subprocess.run(cmd, cwd=ws, env=env, capture_output=True, text=True, timeout=AGENT_TIMEOUT, input=prompt)
    elapsed = time.time() - t0
    with gzip.open(out_dir / f"{label}.jsonl.gz", "wt") as f:
        f.write(r.stdout)
    (out_dir / f"{label}.prompt.md").write_text(prompt)
    if r.stderr:
        (out_dir / f"{label}.stderr.txt").write_text(r.stderr)
    return summarize(r.stdout, elapsed, r.returncode)


LIMIT = re.compile(r"usage limit|rate limit|rate_limit|overloaded|429|5-hour limit|weekly limit|try again later", re.I)


def limited(sdir):
    """Whether the last agent run stopped on a usage or rate limit."""
    try:
        with gzip.open(sdir / "agent.jsonl.gz", "rt") as f:
            tail = f.read()[-4000:]
    except OSError:
        tail = ""
    err = (sdir / "agent.stderr.txt").read_text()[-4000:] if (sdir / "agent.stderr.txt").exists() else ""
    return bool(LIMIT.search(tail + err))


def run_agent_retrying(prompt, ws, before, cwd, env, sdir, model, wait=900, tries=40):
    """Run the agent; on a usage or rate limit, undo its work, wait, and run again."""
    for i in range(tries):
        summary = run_agent(prompt, cwd, env, sdir, model, "agent")
        if not (summary["is_error"] or summary["exit_code"]) or not limited(sdir):
            summary["limit_retries"] = i
            return summary
        print(json.dumps({"limit": True, "retry_in_s": wait, "attempt": i + 1}), flush=True)
        git(ws, "reset", "-q", "--hard", before)
        git(ws, "clean", "-qfdx")
        time.sleep(wait)
    raise RuntimeError("the agent kept hitting the usage limit")


def run_agent_waiting(prompt, cwd, env, out_dir, model, label, wait=900, tries=40):
    """run_agent for helpers that change nothing: on a limit, wait and ask again."""
    for i in range(tries):
        summary = run_agent(prompt, cwd, env, out_dir, model, label)
        text = ""
        try:
            with gzip.open(out_dir / f"{label}.jsonl.gz", "rt") as f:
                text = f.read()[-4000:]
        except OSError:
            pass
        if not (summary["is_error"] or summary["exit_code"]) or not LIMIT.search(text):
            return summary
        print(json.dumps({"limit": True, "label": label, "retry_in_s": wait}), flush=True)
        time.sleep(wait)
    raise RuntimeError("the helper kept hitting the usage limit")


def summarize(stream, elapsed, code):
    result = {}
    for line in stream.splitlines():
        try:
            ev = json.loads(line)
        except json.JSONDecodeError:
            continue
        if ev.get("type") == "result":
            result = ev
    u = result.get("usage") or {}
    tokens = sum(u.get(k, 0) or 0 for k in (
        "input_tokens", "output_tokens", "cache_creation_input_tokens", "cache_read_input_tokens"))
    return {
        "exit_code": code,
        "is_error": result.get("is_error"),
        "num_turns": result.get("num_turns"),
        "tokens": tokens,
        "usage": u,
        "cost_usd": result.get("total_cost_usd"),
        "duration_s": round(elapsed, 1),
    }


# ---- audit -------------------------------------------------------------

NET = re.compile(r"(^|[\s;&|(])(gh|curl|wget)\s|git\s+(fetch|clone|pull|ls-remote|remote\s+add)")
ABS = re.compile(r"(?<![\w.-])(/(?:Users|private|tmp|var|opt|etc|home|Volumes)/[^\s'\"`;|&)]*)")


REL = re.compile(r"(?<![\w/.-])((?:\.\./)+[^\s'\"`;|&)]*|\.\.(?=$|[\s;|&)]))")
CD = re.compile(r"(?:^|[;&|(\n]\s*|&&\s*)cd\s+([^\s;&|)]+)", re.M)


HEREDOC = re.compile(r"<<-?\s*(['\"]?)(\w+)\1[^\n]*\n.*?^\s*\2\s*$", re.S | re.M)


def strip_heredocs(cmd):
    """Drop here-document bodies: they are file contents or code, not shell."""
    return HEREDOC.sub("<<heredoc", cmd)


def escapes(cmd, ws, start, carry=None):
    """Relative paths in a shell command that leave the workspace.

    The command starts in `start`; each `cd` moves it, in order. A path like
    ../../../../.declscope-baseline.yaml from inside the package stays in the
    workspace and is fine.
    """
    cmd = strip_heredocs(cmd)
    cwd, out = Path(start), []
    for part in re.split(r"(?=(?:^|[;&|(\n]\s*|&&\s*)cd\s)", cmd, flags=re.M):
        m = CD.search(part)
        if m:
            target = os.path.expanduser(m.group(1).strip("'\""))
            cwd = Path(os.path.normpath(target if target.startswith("/") else cwd / target))
            # The cd target is judged through cwd below; do not read it again
            # from the new directory.
            part = part[:m.start(1)] + part[m.end(1):]
        for r in REL.findall(part):
            p = os.path.normpath(cwd / r)
            if not (p == str(ws) or p.startswith(str(ws) + "/")):
                out.append(p)
        if not (str(cwd) == str(ws) or str(cwd).startswith(str(ws) + "/")):
            out.append(str(cwd))
    if carry is not None:
        carry["cwd"] = cwd
    return out


def audit(transcript_gz, ws, arm, allowed_extra=(), start=None):
    # Claude Code keeps per-session files for this workspace, named after its
    # path (background output, saved tool results); those belong to the run.
    session_tag = "-" + str(ws).strip("/").replace("/", "-").replace(".", "-")
    allowed = [str(ws), str(mod_cache()), str(Path(mod_cache()).parent), str(go_root()),
               "/dev/null", "/tmp/go-build", "/private/var/folders"]
    allowed += [str(a) for a in allowed_extra]
    issues = []
    # The Bash tool keeps the working directory from one call to the next,
    # so follow it through the transcript in order.
    carry = {"cwd": Path(start or ws)}
    pending, wrapped = {}, set()
    visited = {Path(start or ws)}
    with gzip.open(transcript_gz, "rt") as f:
        for line in f:
            try:
                ev = json.loads(line)
            except json.JSONDecodeError:
                continue
            # A command the sandbox hook wrapped runs in a child shell, so a cd
            # inside it does not carry to the next call.
            for meta in ev.get("tool_result_meta") or []:
                if isinstance(meta, dict) and (meta.get("permission_decision") or {}).get("source") == "hook":
                    wrapped.add(meta.get("id"))
            msg = ev.get("message")
            # Some events carry a plain string as their message; only a
            # message with a list of blocks holds tool calls and results.
            if not isinstance(msg, dict) or not isinstance(msg.get("content"), list):
                continue
            for block in msg["content"]:
                if not isinstance(block, dict):
                    continue
                if block.get("type") == "tool_use":
                    inp = json.dumps(block.get("input"), ensure_ascii=False)
                    name = block.get("name")
                    if name == "Bash" and NET.search(strip_heredocs(block["input"].get("command", ""))):
                        issues.append({"kind": "network", "tool": name, "input": inp[:300]})
                    for p in ABS.findall(inp):
                        if any(p.startswith(a) for a in allowed) or session_tag in p:
                            continue
                        # /tmp is shared by every run, so a file left there can
                        # reach another arm. Recorded apart from other paths.
                        kind = "shared-tmp" if re.match(r"/(private/)?tmp/", p) else "outside-path"
                        issues.append({"kind": kind, "tool": name, "path": p})
                    if name == "Bash":
                        prev = carry["cwd"]
                        command = block["input"].get("command", "")
                        # Whether a command succeeded, and so kept its cd, cannot
                        # be read reliably. A relative path counts only when it
                        # leaves the workspace from every place the agent has been.
                        seen = [p for p in visited if p != prev]
                        hits = escapes(command, ws, prev, carry)
                        for other in seen:
                            if not hits:
                                break
                            alt = set(escapes(command, ws, other))
                            hits = [h for h in hits if h in alt]
                        for p in hits:
                            if any(p.startswith(a) for a in allowed):
                                continue
                            kind = "shared-tmp" if re.match(r"/(private/)?tmp(/|$)", p) else "parent-ref"
                            issues.append({"kind": kind, "tool": name, "path": p})
                        c = carry["cwd"]
                        if str(c) == str(ws) or str(c).startswith(str(ws) + "/"):
                            visited.add(c)
                        pending[block.get("id")] = prev
                if block.get("type") == "tool_result" and block.get("tool_use_id") in pending:
                    # The Bash tool keeps a cd only when the command succeeds and
                    # says so when it does not: "Shell cwd was reset to <dir>".
                    prev = pending.pop(block["tool_use_id"])
                    if block["tool_use_id"] in wrapped:
                        carry["cwd"] = prev
                        continue
                    text = json.dumps(block.get("content"), ensure_ascii=False)
                    m = re.search(r"Shell cwd was reset to ([^\\\"\n]+)", text)
                    if m:
                        carry["cwd"] = Path(m.group(1).strip())
                    elif block.get("is_error"):
                        carry["cwd"] = prev
                if arm in ("A", "P"):
                    text = json.dumps(block, ensure_ascii=False)
                    if "declscope" in text.lower():
                        issues.append({"kind": "leak", "type": block.get("type"), "excerpt": text[:300]})
    # The same path may be reported many times; keep one of each.
    seen, unique = set(), []
    for i in issues:
        k = json.dumps(i, sort_keys=True)
        if k not in seen:
            seen.add(k)
            unique.append(i)
    # Writing to the shared /tmp is a warning, not a breach: it reads nothing
    # of another run. It is still recorded and reported.
    breaches = [i for i in unique if i["kind"] != "shared-tmp"]
    return {"ok": not breaches, "issues": breaches, "warnings": [i for i in unique if i["kind"] == "shared-tmp"]}


# ---- scoring -----------------------------------------------------------

RENAMER_PROMPT = """\
The test files listed below do not compile against the code in this directory.
They were written against an earlier implementation of the same behaviour, which
used different names.

Edit ONLY these test files so that they compile against the current code:

{files}

Rules:
- You may change identifiers: function, method, type, field and constant names,
  and how arguments are passed, so that each call reaches the code that does the
  same job.
- Never change a literal value, an expected result, an assertion, or what a test
  checks. Never add or remove checks.
- Do not edit any other file.
- If a test function has no counterpart in the current code, for example because
  the helper it calls does not exist as a separate function, delete that test
  function and list its name, one per line, in a file named NO_COUNTERPART.txt
  in this directory.

Run `{compile_cmd}` until it compiles. You do not need the tests to pass.
"""

GO_LIT = re.compile(r'`[^`]*`|"(?:\\.|[^"\\])*"|\'(?:\\.|[^\'\\])+\'|\b\d[\d_]*(?:\.\d+)?(?:[eE][+-]?\d+)?\b')
GO_FUNC = re.compile(r"^func\s+(?:\([^)]*\)\s*)?(\w+)", re.M)


def split_funcs(src):
    """Map each top-level func name to its body text (approximate)."""
    out, idx = {}, [(m.start(), m.group(1)) for m in GO_FUNC.finditer(src)]
    for i, (start, name) in enumerate(idx):
        end = idx[i + 1][0] if i + 1 < len(idx) else len(src)
        out[name] = src[start:end]
    return out


def literals_preserved(before, after, removed):
    """Every remaining func keeps the same literals, in the same order."""
    bf, af = split_funcs(before), split_funcs(after)
    problems = []
    for name, body in bf.items():
        if name in removed:
            continue
        if name not in af:
            problems.append(f"{name}: removed without being listed")
            continue
        if GO_LIT.findall(body) != GO_LIT.findall(af[name]):
            problems.append(f"{name}: literals changed")
    return problems


def strip_tool(ws):
    """Remove every trace of declscope, so a reader cannot tell the arm."""
    for name in (".declscope-baseline.yaml", ".declscope.yaml"):
        for p in ws.rglob(name):
            p.unlink()
    shutil.rmtree(ws / ".claude" / "skills", ignore_errors=True)
    for p in ws.rglob("*.go"):
        s = p.read_text()
        if "//declscope:" in s:
            lines = [l for l in s.splitlines(keepends=True) if not l.lstrip().startswith("//declscope:")]
            s2 = "".join(lines)
            s2 = re.sub(r"\s*//declscope:\S+(?: //.*)?$", "", s2, flags=re.M)
            p.write_text(s2)


def score(task, task_dir, step, n, ws, out_dir, model):
    """Swap in upstream's package test suite, rename if needed, run, restore."""
    md = module_dir(task, ws)
    pkg = md / task["package"]
    hidden = task_dir / f"step-{n}" / "hidden"
    stash = new_dir(out_dir.parent / ".stash")
    env = agent_env(task, "A")
    # 1. Stash the agent's own tests in the package.
    agent_tests = sorted(pkg.glob("*_test.go"))
    for t in agent_tests:
        shutil.move(str(t), stash / t.name)
    # 2. Lay upstream's suite and its testdata over the workspace.
    overwritten, added = {}, []
    for src in hidden.rglob("*"):
        if src.is_dir():
            continue
        rel = src.relative_to(hidden)
        dst = ws / rel
        if dst.exists():
            overwritten[str(rel)] = dst.read_bytes()
        else:
            added.append(rel)
        dst.parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(src, dst)
    rel_pkg = "./" + task["package"]
    # -c compiles the test binary without running it.
    compile_cmd = f"go test -c -o /dev/null {rel_pkg}"
    result = {"renamed": False, "no_counterpart": [], "rename_problems": []}
    try:
        c = sh(compile_cmd, cwd=md, env=env, check=False)
        result["compiled_as_is"] = c.returncode == 0
        if c.returncode != 0:
            result.update(rename(task, ws, md, pkg, compile_cmd, out_dir, model, n))
        skip = "|".join(f"^{k['test']}$" for k in step.get("known_failing") or [])
        cmd = sandboxed(["go", "test", "-count=1", "-json", rel_pkg] + (["-skip", skip] if skip else []))
        t = sh(cmd, cwd=md, env=env, check=False, timeout=1800)
        (out_dir / "test.jsonl").write_text(t.stdout)
        result.update(test_summary(t.stdout, t.returncode))
        # Some upstream tests are timing-sensitive (go-sdk's streamable tests
        # failed once under load and passed three times alone). Run the failed
        # tests once more; both outcomes are recorded.
        if result["failed"] and not result["build_failed"]:
            again = "|".join(f"^{n}$" for n in result["failed"])
            t2 = sh(sandboxed(["go", "test", "-count=1", "-json", "-run", again, rel_pkg]),
                    cwd=md, env=env, check=False, timeout=1800)
            (out_dir / "test-retry.jsonl").write_text(t2.stdout)
            r2 = test_summary(t2.stdout, t2.returncode)
            result["first_failed"] = result["failed"]
            result["retry"] = r2
            if r2["success"]:
                result["failed"], result["success"] = [], True
    finally:
        # 3. Restore: remove upstream's files, put back what they covered.
        for rel in added:
            (ws / rel).unlink(missing_ok=True)
        for rel, data in overwritten.items():
            (ws / rel).write_bytes(data)
        for t in pkg.glob("*_test.go"):
            if t.name not in {a.name for a in agent_tests}:
                t.unlink()
        for t in agent_tests:
            shutil.move(str(stash / t.name), t)
        shutil.rmtree(stash, ignore_errors=True)
    if git(ws, "status", "--porcelain").stdout.strip():
        result["restore_dirty"] = git(ws, "status", "--porcelain").stdout
    return result


def rename(task, ws, md, pkg, compile_cmd, out_dir, model, n):
    """Ask a blind agent to match the hidden tests' names to the code."""
    rws = new_dir(WS_ROOT)
    shutil.copytree(ws, rws, dirs_exist_ok=True, ignore=shutil.ignore_patterns(".git"))
    strip_tool(rws)
    rpkg = rws / pkg.relative_to(ws)
    files = sorted(p.name for p in rpkg.glob("*_test.go"))
    before = {f: (rpkg / f).read_text() for f in files}
    prompt = RENAMER_PROMPT.format(
        files="\n".join(f"- {rpkg.relative_to(rws) / f}" for f in files), compile_cmd=compile_cmd)
    env = agent_env(task, "A")
    summary = run_agent_waiting(prompt, rws / md.relative_to(ws), env, out_dir, model, "rename")
    removed = []
    nc = rws / md.relative_to(ws) / "NO_COUNTERPART.txt"
    if nc.exists():
        removed = [l.strip() for l in nc.read_text().splitlines() if l.strip()]
    problems = []
    for f in files:
        after_path = rpkg / f
        after = after_path.read_text() if after_path.exists() else ""
        problems += [f"{f}: {p}" for p in literals_preserved(before[f], after, set(removed))]
        (pkg / f).write_text(after)
    diff = sh(["diff", "-ru", "--exclude=.git", str(ws), str(rws)], check=False).stdout
    (out_dir / "rename.diff").write_text(diff)
    shutil.rmtree(rws, ignore_errors=True)
    return {"renamed": True, "rename_agent": summary, "no_counterpart": removed, "rename_problems": problems}


def test_summary(stream, code):
    passed, failed, skipped = set(), set(), set()
    build_failed = False
    for line in stream.splitlines():
        try:
            ev = json.loads(line)
        except json.JSONDecodeError:
            continue
        a, t = ev.get("Action"), ev.get("Test")
        if a == "build-fail" or (a == "fail" and not t and ev.get("FailedBuild")):
            build_failed = True
        if not t or "/" in t:
            continue
        {"pass": passed, "fail": failed, "skip": skipped}.get(a, set()).add(t)
    return {
        "test_exit": code, "build_failed": build_failed,
        "passed": len(passed), "failed": sorted(failed), "skipped": len(skipped),
        "success": code == 0 and not build_failed,
    }


def require_prescreen(task_dir):
    """Refuse a task whose tests were not checked for external downloads."""
    p = task_dir / "prescreen.json"
    if not p.is_file():
        raise SystemExit(f"{p} is missing. Run harness/prescreen.py and review it first.")
    r = json.loads(p.read_text())
    if r.get("verdict") == "clean":
        return
    review = r.get("review") or {}
    # A reviewer may clear needs-review hits, and only the hits they saw.
    if (r.get("verdict") == "needs-review" and review.get("decision") == "clean"
            and review.get("hit_count") == len(r.get("hits", []))):
        return
    raise SystemExit(f"{task_dir.name}: prescreen verdict is {r.get('verdict')!r}; the task is excluded.")


# ---- main --------------------------------------------------------------

def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--task", required=True)
    ap.add_argument("--arm", required=True, choices=["A", "P", "G"])
    ap.add_argument("--rep", type=int, required=True)
    ap.add_argument("--clones", required=True, help="directory holding upstream clones")
    ap.add_argument("--steps", help="comma-separated step numbers to run (default: all)")
    ap.add_argument("--model", default="claude-opus-5-5")
    ap.add_argument("--out", default="runs", help="output root inside this repository (smoke tests use smoke)")
    args = ap.parse_args()

    task_dir = (REPO / args.task).resolve()
    task = load_yaml(task_dir / "task.yaml")
    # Clones are named owner__name, as the candidate search made them.
    task.setdefault("clone", "__".join(task["repo"].rstrip("/").removesuffix(".git").split("/")[-2:]))
    steps = task["steps"]
    wanted = {int(s) for s in args.steps.split(",")} if args.steps else set(range(1, len(steps) + 1))
    out = REPO / args.out / task["id"] / args.arm / str(args.rep)
    out.mkdir(parents=True, exist_ok=True)

    require_prescreen(task_dir)
    ensure_ws_root()
    ws = new_dir(WS_ROOT)
    meta = {
        "task": task["id"], "arm": args.arm, "rep": args.rep, "model": args.model,
        "workspace": ws.name, "started": datetime.datetime.now().isoformat(timespec="seconds"),
        "steps": [],
    }
    export_seed(task, Path(args.clones), ws)
    warm(task, ws)
    if args.arm == "G":
        setup_guardrail(task, ws)
    base = commit_all(ws, "initial")

    for n, step in enumerate(steps, 1):
        if n not in wanted:
            continue
        sdir = out / f"step-{n}"
        sdir.mkdir(parents=True, exist_ok=True)
        stage = step.get("stage", 1)
        arm = args.arm if stage == 1 else "A"
        if stage == 2:
            # The structure stage 1 left is what the judge compares.
            meta["stage1_end"] = git(ws, "rev-parse", "HEAD").stdout.strip()
            (out / "stage1.patch").write_text(git(ws, "diff", base, "HEAD").stdout)
            (out / "stage1-full.patch").write_text(git(ws, "diff", "-U1000000", base, "HEAD").stdout)
            strip_tool(ws)
            # What the judge and the reviewers read: the package as stage 1
            # left it, with every trace of the tool removed.
            snap = out / "stage1-package"
            shutil.rmtree(snap, ignore_errors=True)
            shutil.copytree(module_dir(task, ws) / task["package"], snap,
                            ignore=shutil.ignore_patterns("testdata"))
            # The history records the tool's files and their removal, so the
            # stage 2 agent would see which arm wrote the code. Keep it where
            # the agent is not pointed, and start a history of one commit.
            archive = WS_ROOT.parent / (WS_ROOT.name + "-archive")
            archive.mkdir(parents=True, exist_ok=True)
            shutil.move(str(ws / ".git"), archive / (ws.name + ".git"))
            git(ws, "init", "-q", "-b", "main")
            commit_all(ws, "initial")
        # Upstream's changes outside the package, which the task text says are
        # already done: first what landed between steps, then this step's own.
        applied = False
        for name in ("intermediate.patch", "outside.patch"):
            patch = task_dir / f"step-{n}" / name
            if patch.exists():
                git(ws, "apply", "--whitespace=nowarn", str(patch))
                applied = True
        if applied:
            warm(task, ws)  # a patch may change go.mod; fetching runs no code
        for g in step.get("given") or []:
            src = task_dir / f"step-{n}" / "given" / g
            (ws / g).parent.mkdir(parents=True, exist_ok=True)
            shutil.copyfile(src, ws / g)
        before = commit_all(ws, f"before step {n}")
        prompt = prompt_for(task_dir, step, arm, task)
        summary = run_agent_retrying(prompt, ws, before, module_dir(task, ws), agent_env(task, arm, ws=ws), sdir, args.model)
        after = commit_all(ws, f"step {n}")
        diff = git(ws, "diff", before, after).stdout
        (sdir / "diff.patch").write_text(diff)
        # The same diff with every line of each touched file, for reading a
        # changed file whole.
        (sdir / "diff-full.patch").write_text(git(ws, "diff", "-U1000000", before, after).stdout)
        stat = git(ws, "diff", "--numstat", before, after).stdout
        files = [l.split("\t") for l in stat.splitlines() if l]
        result = {
            "step": n, "pr": step.get("pr"), "stage": stage, "arm_setup": arm, "agent": summary,
            "files_touched": len(files),
            "lines_touched": sum(int(a) + int(d) for a, d, _ in files if a != "-"),
            "directives": len(re.findall(r"^\+.*//declscope:", diff, re.M)),
        }
        result["score"] = score(task, task_dir, step, n, ws, sdir, args.model)
        result["audit"] = audit(sdir / "agent.jsonl.gz", ws, arm, allowed_extra=[run_tmp(ws)], start=module_dir(task, ws))
        if (sdir / "rename.jsonl.gz").exists():
            result["audit_rename"] = audit(sdir / "rename.jsonl.gz", WS_ROOT, "A")
        (sdir / "metrics.json").write_text(json.dumps(result, indent=2, ensure_ascii=False))
        meta["steps"].append(result)
        print(json.dumps({"step": n, "success": result["score"].get("success"),
                          "tokens": summary["tokens"], "turns": summary["num_turns"],
                          "audit_ok": result["audit"]["ok"]}), flush=True)

    meta["finished"] = datetime.datetime.now().isoformat(timespec="seconds")
    (out / "run.json").write_text(json.dumps(meta, indent=2, ensure_ascii=False))
    print(f"workspace kept at {ws}", file=sys.stderr)


if __name__ == "__main__":
    main()
