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

    seen = {(start, (0, 1)): math.inf}
    stack = [(0, start, (0, 1), {start})]
    heapq.heapify(stack)
    while stack:
        p, pos, dir, vis = heapq.heappop(stack)

        if pos == end:
            return p

        if (pos, dir) in seen and p >= seen[(pos, dir)]:
            continue

        seen[(pos, dir)] = p

        r, c = pos
        dr, dc = dir

        nr, nc = r + dr, c + dc
        if grid[nr][nc] != "#" and (nr, nc) not in vis:
            heapq.heappush(stack, (p + 1, (nr, nc), dir, vis | {(nr, nc)}))

        ndr, ndc = dc * -1, dr
        nr, nc = r + ndr, c + ndc
        if grid[nr][nc] != "#":
            heapq.heappush(stack, (p + 1000, pos, (ndr, ndc), vis))

        ndr, ndc = dc, dr * -1
        nr, nc = r + ndr, c + ndc
        if grid[nr][nc] != "#":
            heapq.heappush(stack, (p + 1000, pos, (ndr, ndc), vis))

    return -1


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
            else:
                return len(path)

            continue

        if (pos, dir) in seen and p > seen[(pos, dir)]:
            continue

        seen[(pos, dir)] = p

        r, c = pos
        dr, dc = dir

        nr, nc = r + dr, c + dc
        if grid[nr][nc] != "#" and (nr, nc) not in vis:
            heapq.heappush(stack, (p + 1, (nr, nc), dir, vis | {(nr, nc)}))

        ndr, ndc = dc * -1, dr
        nr, nc = r + ndr, c + ndc
        if grid[nr][nc] != "#":
            heapq.heappush(stack, (p + 1000, pos, (ndr, ndc), vis))

        ndr, ndc = dc, dr * -1
        nr, nc = r + ndr, c + ndc
        if grid[nr][nc] != "#":
            heapq.heappush(stack, (p + 1000, pos, (ndr, ndc), vis))

    return -1


if __name__ == "__main__":
    util.run(16, part1, part2)
