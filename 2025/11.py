from collections import defaultdict
from functools import cache

import util


def part1(input: str) -> int:
    graph = {}
    for line in input.splitlines():
        node, *neighbors = line.split()
        graph[node[:-1]] = neighbors

    paths = 0

    q = [("you", set())]
    while q:
        curr, seen = q.pop()

        if curr == "out":
            paths += 1
            continue

        if curr in seen:
            continue

        for n in graph[curr]:
            q.append((n, seen | {curr}))

    return paths


def part2(input: str) -> int:
    graph = defaultdict(list)
    for line in input.splitlines():
        node, *neighbors = line.split()
        graph[node[:-1]] = neighbors

    @cache
    def dfs(curr: str, end: str) -> int:
        if curr == end:
            return 1

        return sum(dfs(n, end) for n in graph[curr])

    return dfs("svr", "fft") * dfs("fft", "dac") * dfs("dac", "out")


if __name__ == "__main__":
    util.run(11, 2025, part1, part2)
