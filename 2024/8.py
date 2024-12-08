import itertools
from collections import defaultdict

import util


def part1(input: str) -> int:
    grid = {}
    frequencies = defaultdict(set)
    lines = input.splitlines()
    for r, line in enumerate(lines):
        for c, char in enumerate(line):
            grid[(r, c)] = char
            if char != ".":
                frequencies[char].add((r, c))

    antinodes = set()

    for k, v in frequencies.items():
        for (r1, c1), (r2, c2) in itertools.combinations(v, 2):
            dr, dc = r1 - r2, c1 - c2
            for n in [
                (r1 - dr, c1 - dc),
                (r1 + dr, c1 + dc),
                (r2 - dr, c2 - dc),
                (r2 + dr, c2 + dc),
            ]:
                if n != (r1, c1) and n != (r2, c2) and n in grid:
                    antinodes.add(n)

    return len(antinodes)


def part2(input: str) -> int:
    grid = {}
    frequencies = defaultdict(set)
    lines = input.splitlines()
    for r, line in enumerate(lines):
        for c, char in enumerate(line):
            grid[(r, c)] = char
            if char != ".":
                frequencies[char].add((r, c))

    antinodes = set()

    for k, v in frequencies.items():
        for (r1, c1), (r2, c2) in itertools.combinations(v, 2):
            dr, dc = r1 - r2, c1 - c2
            r, c = r1, c1
            while (r, c) in grid:
                antinodes.add((r, c))
                r += dr
                c += dc
            r, c = r1, c1
            while (r, c) in grid:
                antinodes.add((r, c))
                r -= dr
                c -= dc

    return len(antinodes)


if __name__ == "__main__":
    util.run(8, part1, part2)
