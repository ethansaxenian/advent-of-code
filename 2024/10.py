import math

import util


def part1(input: str) -> int:
    trailheads = []
    grid = {}
    for r, line in enumerate(input.splitlines()):
        for c, char in enumerate(line):
            grid[(r, c)] = int(char)
            if char == "0":
                trailheads.append((r, c))
    s = 0
    for th in trailheads:
        vis = set()
        ends = set()
        q = [th]

        while q:
            loc = q.pop()

            h = grid.get(loc, None)

            if h is None:
                continue

            vis.add(loc)
            if h == 9:
                ends.add(loc)
                continue

            r, c = loc
            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                new = (r + dr, c + dc)
                if new not in grid or new in vis or (grid.get(new, math.inf) != h + 1):
                    continue

                q.append(new)

        s += len(ends)

    return s


def part2(input: str) -> int:
    trailheads = []
    grid = {}
    for r, line in enumerate(input.splitlines()):
        for c, char in enumerate(line):
            grid[(r, c)] = int(char)
            if char == "0":
                trailheads.append((r, c))

    s = 0
    for th in trailheads:
        paths = set()
        q = [(th, [th])]

        while q:
            loc, path = q.pop()

            h = grid.get(loc, None)

            if h is None:
                continue

            if h == 9:
                paths.add("".join(map(str, path)))
                continue

            r, c = loc
            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                new = (r + dr, c + dc)
                if new not in grid or (grid.get(new, math.inf) != h + 1):
                    continue

                q.append((new, path + [new]))

        s += len(paths)

    return s


if __name__ == "__main__":
    util.run(10, part1, part2)
