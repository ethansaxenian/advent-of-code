from collections import Counter
from collections.abc import Iterable

import util


def blink(initial: Iterable[int], num: int) -> int:
    stones = Counter(initial)
    for i in range(num):
        next_stones = Counter()
        for stone, n in stones.items():
            if stone == 0:
                next_stones[1] += n
            elif (l := len(s := str(stone))) % 2 == 0:
                a, b = map(int, [s[: l // 2], s[l // 2 :]])
                next_stones[a] += n
                next_stones[b] += n
            else:
                next_stones[stone * 2024] += n
        stones = next_stones

    return stones.total()


def part1(input: str) -> int:
    return blink(map(int, input.split()), 25)


def part2(input: str) -> int:
    return blink(map(int, input.split()), 75)


if __name__ == "__main__":
    util.run(11, part1, part2)
