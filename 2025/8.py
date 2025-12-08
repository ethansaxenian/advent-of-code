import itertools
import math
import operator
from collections import defaultdict
from functools import reduce
from typing import cast

import util

type Point = tuple[int, int, int]


def distance(a: Point, b: Point) -> float:
    return math.sqrt((a[0] - b[0]) ** 2 + (a[1] - b[1]) ** 2 + (a[2] - b[2]) ** 2)


def dfs(p: Point, graph: dict[Point, set[Point]]) -> set[Point]:
    vis = set()
    q = [p]
    while q:
        p = q.pop()

        if p in vis:
            continue

        vis.add(p)

        for n in graph[p]:
            q.append(n)
    return vis


def part1(input: str) -> int:
    points = list(
        map(
            lambda line: cast(
                Point,
                tuple(map(int, line.split(","))),
            ),
            input.splitlines(),
        )
    )

    dists = {}
    for a, b in itertools.combinations(points, 2):
        dists[(a, b)] = distance(a, b)

    pairs = sorted(dists.items(), key=lambda item: item[1])

    graph = defaultdict(set)

    for (a, b), _ in itertools.islice(pairs, 1000):
        graph[a].add(b)
        graph[b].add(a)

    circuits = []

    vis = set()
    for p in points:
        if p in vis:
            continue

        circuit = dfs(p, graph)
        circuits.append(len(circuit))
        vis |= circuit

    return reduce(operator.mul, sorted(circuits, reverse=True)[:3])


def part2(input: str) -> int:
    points = list(
        map(
            lambda line: cast(
                Point,
                tuple(map(int, line.split(","))),
            ),
            input.splitlines(),
        )
    )

    dists: dict[tuple[Point, Point], float] = {}
    for a, b in itertools.combinations(points, 2):
        dists[(a, b)] = distance(a, b)

    pairs = sorted(dists.items(), key=lambda item: item[1])

    graph = defaultdict(set)

    for (a, b), _ in pairs:
        graph[a].add(b)
        graph[b].add(a)

        circuit = dfs(a, graph)
        if len(circuit) == len(points):
            return a[0] * b[0]

    return -1


if __name__ == "__main__":
    util.run(8, part1, part2)
