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

    for x in g.keys():
        for y in g[x]:
            if y > x:
                continue
            for z in g[y]:
                if z > y:
                    continue
                if z not in g[x]:
                    continue
                if "t" not in (x[0] + y[0] + z[0]):
                    continue
                n += 1
    return n


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
