import functools
import itertools

import util


def part1(input: str) -> int:
    lines = input.splitlines()
    i = 0
    rules = set()
    while (line := lines[i]) != "":
        a, b = line.split("|")
        rules.add((a, b))
        i += 1

    s = 0
    i += 1
    while i < len(lines):
        update = lines[i].split(",")
        if all((a, b) in rules for a, b in itertools.combinations(update, 2)):
            s += int(update[len(update) // 2])
        i += 1

    return s


def part2(input: str) -> int:
    lines = input.splitlines()
    i = 0
    rules = set()
    while lines[i] != "":
        a, b = lines[i].split("|")
        rules.add((a, b))
        i += 1

    def cmp(a, b):
        return -1 if (a, b) in rules else 1

    s = 0
    i += 1
    while i < len(lines):
        update = lines[i].split(",")
        if any((a, b) not in rules for a, b in itertools.combinations(update, 2)):
            update = sorted(
                update,
                key=functools.cmp_to_key(cmp),
            )
            s += int(update[len(update) // 2])

        i += 1

    return s


if __name__ == "__main__":
    util.run(5, part1, part2)
