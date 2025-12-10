import itertools
from typing import cast

import util


def part1(input: str) -> int:
    points = map(lambda line: tuple(map(int, line.split(","))), input.splitlines())

    m = -1
    for (x0, y0), (x1, y1) in itertools.combinations(points, 2):
        area = abs(x0 - x1 + 1) * abs(y0 - y1 + 1)
        m = max(m, area)

    return m


def compress(nums: list[int]) -> dict[int, int]:
    return {x: i for i, x in enumerate(sorted(set(nums)))}


def part2(input: str) -> int:
    red = [tuple(map(int, line.split(","))) for line in input.splitlines()]
    red = cast(list[tuple[int, int]], red)

    green = set()
    for (x0, y0), (x1, y1) in itertools.combinations(red, 2):
        if x0 == x1:
            for y in range(min(y0, y1), max(y0, y1) + 1):
                green.add((x0, y))
        if y0 == y1:
            for x in range(min(x0, x1), max(x0, x1) + 1):
                green.add((x, y0))

    m = -1

    for (x0, y0), (x1, y1) in itertools.combinations(red, 2):
        x0, x1 = min(x0, x1), max(x0, x1)
        y0, y1 = min(y0, y1), max(y0, y1)

        area = (x1 - x0 + 1) * (y1 - y0 + 1)

        if area > m:
            valid = True
            for gx, gy in green:
                if x0 < gx < x1 and y0 < gy < y1:
                    valid = False
                    break

            if valid:
                m = area

    return m


if __name__ == "__main__":
    util.run(9, 2025, part1, part2)
