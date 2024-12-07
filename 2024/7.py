import functools
import itertools
import operator
from collections.abc import Callable

import util


def calibrate(lines: list[str], operators: list[Callable[[int, int], int]]) -> int:
    ans = 0
    for line in lines:
        test, rest = line.split(":")
        test = int(test)
        nums = list(map(int, rest.split()))
        for comb in itertools.product(operators, repeat=len(nums) - 1):
            res = nums[0]
            for op, x in zip(comb, nums[1:]):
                if res > test:
                    break
                res = op(res, x)

            if res == test:
                ans += res
                break

    return ans


def part1(input: str) -> int:
    return calibrate(input.splitlines(), [operator.add, operator.mul])


@functools.cache
def concat(a: int, b: int) -> int:
    return int(str(a) + str(b))


def part2(input: str) -> int:
    return calibrate(input.splitlines(), [operator.add, operator.mul, concat])


if __name__ == "__main__":
    util.run(7, part1, part2)
