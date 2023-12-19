from util.util import fetch_input
import operator

puzzle_input = fetch_input(19)

ops = {
    "<": operator.lt,
    ">": operator.gt,
}


def parse_workflow(line):
    name, rule_strs = line.split("{")
    rule_strs = rule_strs[:-1].split(",")

    rules = []
    for r in rule_strs:
        if ":" in r:
            v, res = r[2:].split(":")
            rules.append((r[0], r[1], int(v), res))
        else:
            rules.append(r)

    return name, rules


def parse_part(line):
    x, m, a, s = line[1:-1].split(",")
    return int(x[2:]), int(m[2:]), int(a[2:]), int(s[2:])


workflows = {}
parts = []

i = 0
while i < len(puzzle_input):
    line = puzzle_input[i]
    i += 1
    if line == "":
        break
    name, rules = parse_workflow(line)
    workflows[name] = rules


while i < len(puzzle_input):
    line = puzzle_input[i]
    if line == "":
        break
    x, m, a, s = parse_part(line)
    parts.append({"x": x, "m": m, "a": a, "s": s})
    i += 1


def run_workflow(part, w):
    for rule in workflows[w]:
        if isinstance(rule, str):
            if rule in "AR":
                return rule == "A"

            return run_workflow(part, rule)

        attr, op, v, res = rule
        if ops[op](part[attr], v):
            if res in "AR":
                return res == "A"

            return run_workflow(part, res)


total = 0
for part in parts:
    if run_workflow(part, "in"):
        total += sum(part.values())

print(total)

bounds_list = []


def get_workflow_bounds(w, curr_rules, bounds):
    if w == "A":
        bounds_list.append(bounds)
        return

    if w == "R":
        return

    if len(curr_rules) == 1:
        if curr_rules[0] == "A":
            bounds_list.append(bounds)
            return
        elif curr_rules[0] == "R":
            return
        else:
            get_workflow_bounds(curr_rules[0], workflows[curr_rules[0]], bounds)
            return

    attr, op, v, res = curr_rules[0]
    curr_min, curr_max = bounds[attr]
    if op == "<":
        accept = (curr_min, min(curr_max, v - 1))
        reject = (max(curr_min, v), curr_max)
        get_workflow_bounds(res, workflows.get(res, []), {**bounds, **{attr: accept}})
        get_workflow_bounds(w, curr_rules[1:], {**bounds, **{attr: reject}})
    elif op == ">":
        accept = (max(curr_min, v + 1), curr_max)
        reject = (curr_min, min(curr_max, v))
        get_workflow_bounds(res, workflows.get(res, []), {**bounds, **{attr: accept}})
        get_workflow_bounds(w, curr_rules[1:], {**bounds, **{attr: reject}})


get_workflow_bounds(
    "in",
    workflows["in"],
    {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (1, 4000)},
)


total = 0
for b in bounds_list:
    x0, x1 = b["x"]
    m0, m1 = b["m"]
    a0, a1 = b["a"]
    s0, s1 = b["s"]

    total += (x1 - x0 + 1) * (m1 - m0 + 1) * (a1 - a0 + 1) * (s1 - s0 + 1)

print(total)
