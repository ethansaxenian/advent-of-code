from collections import Counter

import util


def part1(input: str) -> int:
    l1, l2 = [], []
    for line in input.splitlines():
        x, y = map(int, line.split())
        l1.append(x)
        l2.append(y)

    s = 0
    for a, b in zip(*map(sorted, (l1, l2))):
        s += abs(a - b)

    return s


def part2(input: str) -> int:
    l1 = []
    c = Counter()
    for line in input.splitlines():
        x, y = map(int, line.split())
        l1.append(x)
        c[y] += 1

    return sum(x * c[x] for x in l1)


if __name__ == "__main__":
    util.run(1, part1, part2)
