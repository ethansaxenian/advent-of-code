import re

import util


def part1(input: str) -> int:
    pattern = re.compile(r"^(.+)\1$")

    s = 0
    for r in input.split(","):
        a, b = map(int, r.split("-"))
        for i in range(a, b + 1):
            if pattern.match(str(i)) is not None:
                s += i

    return s


def part2(input: str) -> int:
    pattern = re.compile(r"^(.+)\1+$")

    s = 0
    for r in input.split(","):
        a, b = map(int, r.split("-"))
        for i in range(a, b + 1):
            if pattern.match(str(i)) is not None:
                s += i
    return s


if __name__ == "__main__":
    util.run(2, part1, part2)
