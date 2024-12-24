import itertools
import operator

import util


def to_int(wires, prefix) -> int:
    bin = sorted([(k, v) for k, v in wires.items() if k[0] == prefix], reverse=True)
    return int("".join(str(v) for _, v in bin), 2)


def part1(input: str) -> int:
    wires_input, gates_input = input.split("\n\n")
    wires = {}
    for w in wires_input.splitlines():
        a, b = w.split(": ")
        wires[a] = int(b)

    gates = {}
    for g in gates_input.splitlines():
        a, op_str, b, _, c = g.split()
        match op_str:
            case "AND":
                op = operator.and_
            case "OR":
                op = operator.or_
            case "XOR":
                op = operator.xor
            case _:
                raise Exception
        gates[c] = (a, op, b)

    while len(gates) > 0:
        for c, (a, op, b) in {**gates}.items():
            if a not in wires or b not in wires:
                continue

            wires[c] = op(wires[a], wires[b])
            del gates[c]

    return to_int(wires, "z")


def part2(input: str) -> str:
    wires_input, gates_input = input.split("\n\n")
    wires = {}
    for w in wires_input.splitlines():
        a, b = w.split(": ")
        wires[a] = int(b)

    gates = {}
    for g in gates_input.splitlines():
        a, op_str, b, _, c = g.split()
        match op_str:
            case "AND":
                op = operator.and_
            case "OR":
                op = operator.or_
            case "XOR":
                op = operator.xor
            case _:
                raise Exception
        gates[c] = (a, op, b)

    # graphviz(gates)

    swaps = [
        ("gws", "nnt"),
        ("hgj", "z33"),
        ("z19", "cph"),
        ("z13", "npf"),
    ]
    for a, b in swaps:
        gates[a], gates[b] = gates[b], gates[a]

    # graphviz(gates)

    while len(gates) > 0:
        for c, (a, op, b) in {**gates}.items():
            if a not in wires or b not in wires:
                continue

            wires[c] = op(wires[a], wires[b])
            del gates[c]

    x = to_int(wires, "x")
    y = to_int(wires, "y")
    z = to_int(wires, "z")
    assert x + y == z

    return ",".join(sorted(list(itertools.chain(*swaps))))


def graphviz(gates):
    # try engine=dot or engine=fdp
    ands = []
    ors = []
    xors = []
    edges = []

    for c, (a, op, b) in gates.items():
        edges.append(f"{a} -> {c}; {b} -> {c};")
        if op is operator.xor:
            xors.append(c)
        if op is operator.and_:
            ands.append(c)
        if op is operator.or_:
            ors.append(c)

    print(f"""digraph G {{
    subgraph and {{
        node [style=filled,color=green];
        {'; '.join(ands)};
    }}
    subgraph or {{
        node [style=filled,color=blue];
        {'; '.join(ors)};
    }}
    subgraph xor {{
        node [style=filled,color=red];
        {'; '.join(xors)};
    }}
    subgraph x {{
        edge [style=dotted];
        {' -> '.join(f'x{i:02}' for i in range(45))};
    }}
    subgraph y {{
        edge [style=dotted];
        {' -> '.join(f'y{i:02}' for i in range(45))};
    }}
    subgraph z {{
        edge [style=dotted];
        {' -> '.join(f'z{i:02}' for i in range(45))};
    }}
    {"\n    ".join(edges)}
}}""")


if __name__ == "__main__":
    util.run(24, part1, part2)
