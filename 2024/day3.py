import math
import re

import util


def part1(input: list[str]) -> int:
    s = 0
    for line in input:
        res = re.findall(r"mul\((\d+),(\d+)\)", line)
        for a, b in res:
            s += int(a) * int(b)

    return s


def part2(input: list[str]) -> int:
    s = 0
    do = True
    for line in input:
        for m in re.finditer(r"mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)", line):
            g = m.group(0)
            if g == "do()":
                do = True
            elif g == "don't()":
                do = False
            elif do:
                s += math.prod(map(int, m.group(1, 2)))
    return s


if __name__ == "__main__":
    util.run(3, part1, part2)
