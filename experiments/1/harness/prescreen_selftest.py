#!/usr/bin/env python3
"""Self-test of prescreen.py on a synthetic repository. Runs no Go code.

  python3 -I harness/prescreen_selftest.py
"""

import json
import shutil
import subprocess
import sys
import tempfile
from pathlib import Path

HERE = Path(__file__).resolve().parent

FILES = {
    "go.mod": "module example.com/m\n\ngo 1.22\n",
    "pkg/clean.go": """package pkg

// Run npm install by hand if you need the docs site. (A comment: ignored.)
func diff() int {
\tuv := 1 // a variable named uv: ignored
\treturn uv
}
""",
    "pkg/launch.go": """package pkg

import (
\t"os/exec"

\t"example.com/m/internal/helper"
)

func install() error {
\thelper.Do()
\treturn exec.Command("npm", "install", "-g", "tool@latest").Run()
}
""",
    "pkg/launch_test.go": """package pkg

const script = `#!/bin/sh
curl -fsSL https://get.tool.dev/install.sh | sh
`
""",
    "internal/helper/helper.go": """package helper

import "net/http"

func Do() { _, _ = http.Get("http://localhost:1") }
""",
    "other/unrelated.go": """package other

const s = "brew install something"
""",
}

EXPECT = {
    ("excluded", "installer", "pkg/launch.go"),
    ("excluded", "installer", "pkg/launch_test.go"),
    ("excluded", "download-url", "pkg/launch_test.go"),
    ("needs-review", "exec", "pkg/launch.go"),
    ("needs-review", "outbound-http", "internal/helper/helper.go"),
}


def main():
    tmp = Path(tempfile.mkdtemp())
    try:
        repo = tmp / "clones" / "example__m"
        for rel, text in FILES.items():
            (repo / rel).parent.mkdir(parents=True, exist_ok=True)
            (repo / rel).write_text(text)
        git = ["git", "-C", str(repo), "-c", "user.name=t", "-c", "user.email=t@t"]
        subprocess.run([*git, "init", "-q"], check=True)
        subprocess.run([*git, "add", "-A"], check=True)
        subprocess.run([*git, "commit", "-q", "-m", "x"], check=True)
        sha = subprocess.run([*git, "rev-parse", "HEAD"], capture_output=True, text=True, check=True).stdout.strip()
        # The same layout as the repository: the harness inside an
        # experiment, and the tasks at the top.
        root = tmp / "exp"
        harness = root / "experiments" / "1" / "harness"
        harness.mkdir(parents=True)
        shutil.copy(HERE / "prescreen.py", harness)
        (root / "tasks" / "t").mkdir(parents=True)
        (root / "tasks" / "t" / "task.yaml").write_text(
            f"id: t\nrepo: https://github.com/example/m\npackage: pkg\nseed: {sha}\nsteps: []\n")
        subprocess.run([sys.executable, "-I", str(harness / "prescreen.py"),
                        "--task", "tasks/t", "--clones", str(tmp / "clones")], check=True, capture_output=True)
        r = json.loads((root / "tasks" / "t" / "prescreen.json").read_text())
        got = {(h["severity"], h["kind"], h["path"]) for h in r["hits"]}
        ok = got == EXPECT and r["verdict"] == "excluded"
        print(json.dumps({"ok": ok, "verdict": r["verdict"],
                          "missing": sorted(EXPECT - got), "unexpected": sorted(got - EXPECT)}))
        return 0 if ok else 1
    finally:
        shutil.rmtree(tmp, ignore_errors=True)


if __name__ == "__main__":
    sys.exit(main())
