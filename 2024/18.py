import sys
from collections import deque

import util

H = W = 71
L = 1024


def dijkstras(bytes):
    start = (0, 0)
    end = (H - 1, W - 1)
    q = deque([(0, start, set())])
    dists = {start: sys.maxsize}
    while q:
        x, pos, vis = q.popleft()

        if pos == end:
            return vis

        if pos in dists and x >= dists[pos]:
            continue

        dists[pos] = x

        r, c = pos
        for dr, dc in [(1, 0), (0, -1), (-1, 0), (0, 1)]:
            new_pos = r + dr, c + dc
            if (
                0 <= r + dr < H
                and 0 <= c + dc < W
                and new_pos not in bytes
                and new_pos not in vis
            ):
                q.append((x + 1, new_pos, vis | {pos}))

    return set()


def part1(input: str) -> int:
    bytes = set(
        (int(y), int(x))
        for x, y in map(lambda line: line.split(","), input.splitlines()[:L])
    )

    return len(dijkstras(bytes))


def part2(input: str) -> str:
    bytes = list(
        (int(y), int(x))
        for x, y in map(lambda line: line.split(","), input.splitlines())
    )

    first, rest = set(bytes[:L]), bytes[L:]

    lo = 0
    hi = len(rest)

    while hi - lo > 1:
        curr = (hi + lo) // 2

        path = dijkstras(first | set(rest[:curr]))
        if len(path) > 0:
            lo = curr
        else:
            hi = curr

    return f"{rest[lo][1]},{rest[lo][0]}"


if __name__ == "__main__":
    util.run(18, part1, part2)
