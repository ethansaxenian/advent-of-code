import itertools
from collections import defaultdict

import networkx as nx

import util


def part1(input: str) -> int:
    g = defaultdict(set)
    for line in input.splitlines():
        a, b = line.split("-")
        g[a].add(b)
        g[b].add(a)

    n = 0
    for x, y, z in itertools.combinations(g.keys(), 3):
        if "t" not in (x[0] + y[0] + z[0]):
            continue

        if not (x in g[y] and x in g[z]):
            continue

        if not (y in g[x] and y in g[z]):
            continue

        if not (z in g[x] and z in g[y]):
            continue

        n += 1

    return n


def is_lan(g, vs):
    for v in vs:
        e = g[v] & vs
        if len(e | {v}) < len(vs) or not e.issubset(vs - {v}):
            return False

    return True


def part2(input: str) -> str:
    G = nx.Graph()
    for line in input.splitlines():
        a, b = line.split("-")
        G.add_edge(a, b)
        G.add_edge(b, a)

    lan_party = max(nx.find_cliques(G), key=len)
    return ",".join(sorted(lan_party))


if __name__ == "__main__":
    util.run(23, part1, part2)
