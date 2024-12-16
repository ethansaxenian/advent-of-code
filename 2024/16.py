import heapq
import math

import util

DIRS = [(0, 1), (-1, 0), (0, -1), (1, 0)]


def part1(input: str) -> int:
    start = end = (-1, -1)
    grid = input.splitlines()
    for r, line in enumerate(grid):
        for c, char in enumerate(line):
            if char == "S":
                start = (r, c)
            if char == "E":
                end = (r, c)

    min_points = math.inf
    seen = {(start, (0, 1)): math.inf}
    stack = [(0, start, (0, 1), {start})]
    heapq.heapify(stack)
    while stack:
        p, pos, dir, vis = heapq.heappop(stack)

        if pos == end:
            min_points = min(min_points, p)
            continue

        if (pos, dir) in seen and p >= seen[(pos, dir)]:
            continue

        seen[(pos, dir)] = p

        nr, nc = pos[0] + dir[0], pos[1] + dir[1]
        if grid[nr][nc] != "#" and (nr, nc) not in vis:
            heapq.heappush(stack, (p + 1, (nr, nc), dir, vis | {(nr, nc)}))

        for next_dir in DIRS:
            if next_dir == dir:
                continue

            nr, nc = pos[0] + next_dir[0], pos[1] + next_dir[1]
            if grid[nr][nc] == "#":
                continue

            heapq.heappush(stack, (p + 1000, pos, next_dir, vis))

    return int(min_points)


def part2(input: str) -> int:
    start = end = (-1, -1)
    grid = input.splitlines()
    for r, line in enumerate(grid):
        for c, char in enumerate(line):
            if char == "S":
                start = (r, c)
            if char == "E":
                end = (r, c)

    min_points = math.inf
    path = set()
    seen = {(start, (0, 1)): math.inf}
    stack = [(0, start, (0, 1), {start})]
    heapq.heapify(stack)
    while stack:
        p, pos, dir, vis = heapq.heappop(stack)

        if pos == end:
            if p < min_points:
                min_points = p
                path = vis
            elif p == min_points:
                path |= vis

            continue

        if (pos, dir) in seen and p > seen[(pos, dir)]:
            continue

        seen[(pos, dir)] = p

        nr, nc = pos[0] + dir[0], pos[1] + dir[1]
        if grid[nr][nc] != "#" and (nr, nc) not in vis:
            heapq.heappush(stack, (p + 1, (nr, nc), dir, vis | {(nr, nc)}))

        for next_dir in DIRS:
            if next_dir == dir:
                continue

            nr, nc = pos[0] + next_dir[0], pos[1] + next_dir[1]
            if grid[nr][nc] == "#":
                continue

            heapq.heappush(stack, (p + 1000, pos, next_dir, vis))

    return len(path)


if __name__ == "__main__":
    util.run(16, part1, part2)
