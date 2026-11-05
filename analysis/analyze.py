#!/usr/bin/env python3
"""Apply the decision rule of design.md to the results.

  python3 -I analysis/analyze.py --root runs [--human analysis/human/answers.json]

Writes analysis/results.json and analysis/results.md. With --human, the
colleagues' answers are decoded with analysis/human/key.json and their
agreement with the AI judge is measured (design.md, 人による確かめ).
"""

import argparse
import json
import random
from pathlib import Path

REPO = Path(__file__).resolve().parent.parent
PAIRS = [("A", "P"), ("P", "G"), ("A", "G")]
ARM = {"A": "A（なし）", "P": "P（プロンプト）", "G": "G（ガードレール）"}
B = 10000


def bootstrap_ci(values, stat=lambda v: sum(v) / len(v)):
    if not values:
        return None, None, None
    rng = random.Random(20261105)
    point = stat(values)
    samples = sorted(stat([rng.choice(values) for _ in values]) for _ in range(B))
    return point, samples[int(0.025 * B)], samples[int(0.975 * B) - 1]


def load_runs(root):
    tasks = {}
    for tdir in sorted(p for p in root.iterdir() if p.is_dir()):
        for arm in ("A", "P", "G"):
            for rdir in sorted((tdir / arm).glob("*")) if (tdir / arm).is_dir() else []:
                rj = rdir / "run.json"
                if not rj.exists():
                    continue
                r = json.loads(rj.read_text())
                steps = r["steps"]
                s1 = [s for s in steps if s["stage"] == 1]
                s2 = [s for s in steps if s["stage"] == 2]
                rev = rdir / "review.json"
                tasks.setdefault(tdir.name, {}).setdefault(arm, []).append({
                    "rep": rdir.name,
                    "stage1_last_success": bool(s1 and s1[-1]["score"].get("success")),
                    "stage1_successes": [bool(s["score"].get("success")) for s in s1],
                    "stage1_tokens": sum(s["agent"]["tokens"] or 0 for s in s1),
                    "stage2_success": bool(s2 and s2[-1]["score"].get("success")),
                    "stage2_tokens": s2[-1]["agent"]["tokens"] if s2 else None,
                    "stage2_turns": s2[-1]["agent"]["num_turns"] if s2 else None,
                    "directives": sum(s.get("directives") or 0 for s in s1),
                    "audit_ok": all(s["audit"]["ok"] for s in steps),
                    "review": json.loads(rev.read_text()) if rev.exists() else None,
                })
    return tasks


def compare(tasks, judges, x, y):
    fails = {a: sum(1 for t in tasks.values() for r in t.get(a, []) if not r["stage1_last_success"]) for a in (x, y)}
    guard = "下がっていない" if fails[x] - fails[y] <= 1 else "下がっている"
    shares, per_task = [], {}
    for tid, j in judges.items():
        cs = [c for c in j["comparisons"] if {c["x"].split("/")[0], c["y"].split("/")[0]} == {x, y}]
        wins = sum(1 for c in cs if c["winner"] == x)
        losses = sum(1 for c in cs if c["winner"] == y)
        ties = len(cs) - wins - losses
        per_task[tid] = {"wins": wins, "losses": losses, "ties": ties}
        if wins + losses:
            shares.append(wins / (wins + losses))
    share, lo, hi = bootstrap_ci(shares)
    if guard == "下がっている":
        result = f"{x} の負け"
    elif lo is not None and lo > 0.5:
        result = f"{x} の勝ち"
    elif hi is not None and hi < 0.5:
        result = f"{y} の勝ち"
    else:
        result = "差なし"
    ratios = []
    for t in tasks.values():
        a = [r["stage2_tokens"] for r in t.get(x, []) if r["stage2_tokens"]]
        b = [r["stage2_tokens"] for r in t.get(y, []) if r["stage2_tokens"]]
        if a and b:
            ratios.append((sum(a) / len(a)) / (sum(b) / len(b)))
    ratio, rlo, rhi = bootstrap_ci(ratios)
    return {"x": x, "y": y, "failures": fails, "guard": guard, "win_share": share, "ci": [lo, hi],
            "tasks_with_a_decision": len(shares), "per_task": per_task, "result": result,
            "stage2_token_ratio": ratio, "stage2_token_ratio_ci": [rlo, rhi]}


def human_agreement(judges, answers_path):
    key = json.loads((REPO / "analysis" / "human" / "key.json").read_text())
    answers = json.loads(Path(answers_path).read_text())  # list of {judge, answers: [{pair, choice}]}
    votes = {}
    for sheet in answers:
        for a in sheet["answers"]:
            if not a.get("choice"):
                continue
            k = key[a["pair"]]
            arm = "tie" if a["choice"] == "tie" else k[a["choice"]].split("/")[0]
            votes.setdefault(a["pair"], []).append(arm)
    rows, agree, counted = [], 0, 0
    for pid, vs in sorted(votes.items()):
        k = key[pid]
        top = max(set(vs), key=vs.count)
        majority = top if vs.count(top) * 2 > len(vs) else "tie"
        j = judges.get(k["task"], {"comparisons": []})
        pair = {k["1"].split("/")[0], k["2"].split("/")[0]}
        ai = next((c["winner"] for c in j["comparisons"] if {c["x"].split("/")[0], c["y"].split("/")[0]} == pair), None)
        rows.append({"pair": pid, "task": k["task"], "votes": vs, "human": majority, "ai": ai})
        if majority != "tie":
            counted += 1
            agree += ai == majority
    rate = agree / counted if counted else None
    return {"rows": rows, "agreement": rate, "counted": counted, "meets_75": rate is not None and rate >= 0.75}


def fmt(v, d=2):
    return "–" if v is None else f"{v:.{d}f}"


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="runs")
    ap.add_argument("--human")
    args = ap.parse_args()
    root = (REPO / args.root).resolve()
    tasks = load_runs(root)
    judges = {t: json.loads((root / t / "judge.json").read_text()) for t in tasks if (root / t / "judge.json").exists()}
    comps = [compare(tasks, judges, x, y) for x, y in PAIRS]
    out = {"tasks": tasks, "comparisons": comps}
    if args.human:
        out["human"] = human_agreement(judges, args.human)
    (REPO / "analysis" / "results.json").write_text(json.dumps(out, indent=2, ensure_ascii=False) + "\n")

    md = ["# 集計結果", "", "`analysis/analyze.py` が `runs/` から作る。判定ルールは [design.md](../design.md#判定ルール) のとおり。", ""]
    md += ["## 判定", "", "| 比較 | 歯止め（段階1の最後の段の失敗） | 分割の勝率（95% 区間） | 判定が付いた課題 | 結果 | 段階2のトークン比（参考） |",
           "| --- | --- | --- | --- | --- | --- |"]
    for c in comps:
        md.append(f"| {ARM[c['x']]} 対 {ARM[c['y']]} | {c['x']} {c['failures'][c['x']]} / {c['y']} {c['failures'][c['y']]}: {c['guard']} "
                  f"| {fmt(c['win_share'])}（{fmt(c['ci'][0])}〜{fmt(c['ci'][1])}） | {c['tasks_with_a_decision']} / {len(judges)} "
                  f"| **{c['result']}** | {fmt(c['stage2_token_ratio'])}（{fmt(c['stage2_token_ratio_ci'][0])}〜{fmt(c['stage2_token_ratio_ci'][1])}） |")
    md += ["", "勝率は、引き分けを除いた見比べのうち、左の群が勝った割合。トークン比は、左の群の段階2のトークン数を右の群で割ったもの。1 より小さいと左の群が少ない。", ""]
    md += ["## 課題ごとの見比べ", "", "| 課題 | " + " | ".join(f"{ARM[x]} 対 {ARM[y]}" for x, y in PAIRS) + " |", "| --- | --- | --- | --- |"]
    for t in sorted(judges):
        cells = []
        for c in comps:
            p = c["per_task"].get(t, {})
            cells.append(f"{p.get('wins', 0)} 勝 {p.get('losses', 0)} 敗 {p.get('ties', 0)} 分")
        md.append(f"| {t} | " + " | ".join(cells) + " |")
    md += ["", "## run ごとの記録", "", "| 課題 | 群 | 段階1の成否（段ごと） | 段階1のトークン | 段階2の成否 | 段階2のトークン | ディレクティブ | 構造レビューの指摘（1000行あたり） | 監査 |",
           "| --- | --- | --- | --- | --- | --- | --- | --- | --- |"]
    for t, arms in tasks.items():
        for a in ("A", "P", "G"):
            for r in arms.get(a, []):
                rv = r["review"] or {}
                md.append(f"| {t} | {a} | {' '.join('○' if s else '×' for s in r['stage1_successes'])} | {r['stage1_tokens'] // 1000}k "
                          f"| {'○' if r['stage2_success'] else '×'} | {(r['stage2_tokens'] or 0) // 1000}k | {r['directives']} "
                          f"| {rv.get('counted', '–')}（{fmt(rv.get('per_kloc'), 1)}） | {'問題なし' if r['audit_ok'] else '違反あり'} |")
    if "human" in out:
        h = out["human"]
        md += ["", "## 人による確かめ", "", f"人の多数決と AI の判定の一致率: {fmt(h['agreement'])}（{h['counted']} 組）。基準の 75% を{'満たす' if h['meets_75'] else '満たさない'}。", ""]
    (REPO / "analysis" / "results.md").write_text("\n".join(md) + "\n")
    print("\n".join(md[:12]))


if __name__ == "__main__":
    main()
