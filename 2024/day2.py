import itertools

import util


def is_safe(lst: list[int]) -> bool:
    if not (sorted(lst) == lst or sorted(lst, reverse=True) == lst):
        return False

    return all(0 < abs(x - y) <= 3 for x, y in itertools.pairwise(lst))


def part1(input: list[str]) -> int:
    s = 0
    for line in input:
        ns = list(map(int, line.split()))
        if is_safe(ns):
            s += 1

    return s


def part2(input: list[str]) -> int:
    s = 0
    for line in input:
        line = list(map(int, line.split()))
        for i in range(len(line)):
            lst = line[:i] + line[i + 1 :]
            if is_safe(lst):
                s += 1
                break

    return s


if __name__ == "__main__":
    util.run(2, part1, part2)
