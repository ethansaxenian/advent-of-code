import math
from collections import Counter, defaultdict
from functools import cache

import util


@cache
def change(stone: int) -> list[int]:
    if stone == 0:
        return [1]

    digits = int(math.log10(stone) + 1)
    if digits % 2 == 0:
        x = 10 ** (digits // 2)
        return [stone // x, stone % x]

    return [stone * 2024]


def blink(stones: dict[int, int]) -> dict[int, int]:
    _stones = defaultdict(int)
    for stone, count in stones.items():
        for s in change(stone):
            _stones[s] += count
    return _stones


def part1(input: str) -> int:
    stones = Counter(map(int, input.split()))

    for _ in range(25):
        stones = blink(stones)

    return sum(stones.values())


def part2(input: str) -> int:
    stones = Counter(map(int, input.split()))

    for _ in range(75):
        stones = blink(stones)

    return sum(stones.values())


if __name__ == "__main__":
    util.run(11, part1, part2)
