import itertools
from collections import deque

import util


def race(maze, start, end):
    dists = {}

    q = deque([(0, start)])
    while q:
        x, pos = q.popleft()

        if pos == end:
            return dists

        if pos in dists and x >= dists[pos]:
            continue

        dists[pos] = x

        r, c = pos
        for dr, dc in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
            new_pos = r + dr, c + dc
            if maze[new_pos] != "#":
                q.append((x + 1, new_pos))

    raise Exception


def solve(input: str, cheat_length: int) -> int:
    grid = {}
    lines = input.splitlines()
    start = end = (-1, -1)
    for r, line in enumerate(lines):
        for c, char in enumerate(line):
            grid[(r, c)] = char
            if char == "S":
                start = (r, c)
            elif char == "E":
                end = (r, c)

    dists = race(grid, start, end)
    time = len(dists)
    dists[end] = time

    n = 0

    for a, b in itertools.combinations(dists.keys(), 2):
        if (manhattan := abs(a[0] - b[0]) + abs(a[1] - b[1])) <= cheat_length:
            dist = abs(dists[a] - dists[b])
            if time - (time - dist + manhattan) >= 100:
                n += 1

    return n


def part1(input: str) -> int:
    return solve(input, 2)


def part2(input: str) -> int:
    return solve(input, 20)


if __name__ == "__main__":
    util.run(20, part1, part2)
