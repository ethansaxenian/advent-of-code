import re

import util


def part1(input: str) -> int:
    *shapes, regions = input.split("\n\n")

    shape_sizes = [s.count("#") for s in shapes]

    res = 0
    for r in regions.splitlines():
        length, width, *amts = map(int, re.split(r"[^0-9]+", r))
        min_area = sum(s * n for s, n in zip(shape_sizes, amts))
        res += min_area <= length * width

    return res


def part2(input: str) -> int:
    return 0


if __name__ == "__main__":
    util.run(12, 2025, part1, part2)
