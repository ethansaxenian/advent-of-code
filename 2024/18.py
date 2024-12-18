import heapq
import sys
from collections import defaultdict

import util

H = W = 71
L = 1024


def dijkstras(grid, bytes) -> int:
    start = (0, 0)
    end = (H - 1, W - 1)
    q: list[tuple[int, tuple[int, int], set[tuple[int, int]]]] = [(0, start, set())]
    dists: dict[tuple[int, int], int] = {start: sys.maxsize}
    while q:
        x, pos, vis = heapq.heappop(q)

        if pos == end:
            return len(vis)

        if pos in dists and x >= dists[pos]:
            continue

        dists[pos] = x

        r, c = pos
        for dr, dc in [(1, 0), (0, -1), (-1, 0), (0, 1)]:
            new_pos = r + dr, c + dc
            if grid[pos] and new_pos not in bytes and new_pos not in vis:
                heapq.heappush(q, (x + 1, new_pos, vis | {pos}))

    return -1


def part1(input: str) -> int:
    bytes = set(
        (int(y), int(x))
        for x, y in map(lambda line: line.split(","), input.splitlines()[:L])
    )

    grid = defaultdict(lambda: False)
    for r in range(H):
        for c in range(W):
            if (r, c) in bytes:
                grid[(r, c)] = False
            else:
                grid[(r, c)] = True

    return dijkstras(grid, set())


def part2(input: str) -> str:
    extra_bytes = list(
        (int(y), int(x))
        for x, y in map(lambda line: line.split(","), input.splitlines())
    )

    curr_bytes, rest = set(extra_bytes[:L]), extra_bytes[L:]

    grid = defaultdict(lambda: False)
    for r in range(H):
        for c in range(W):
            if (r, c) in curr_bytes:
                grid[(r, c)] = False
            else:
                grid[(r, c)] = True

    extra_bytes = set()
    for byte in rest:
        extra_bytes.add(byte)
        if dijkstras(grid, extra_bytes) == -1:
            return f"{byte[1]},{byte[0]}"

    raise Exception


if __name__ == "__main__":
    util.run(18, part1, part2)
