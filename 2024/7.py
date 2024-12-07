import operator
from collections.abc import Callable

import util


def dfs(test: int, nums: list[int], operators: list[Callable[[int, int], int]]) -> bool:
    stack = [(nums[0], 1)]
    while stack:
        res, i = stack.pop()
        if res > test:
            continue

        if i == len(nums):
            if res == test:
                return True
            continue

        for op in operators:
            stack.append((op(res, nums[i]), i + 1))

    return False


def calibrate(lines: list[str], operators: list[Callable[[int, int], int]]) -> int:
    ans = 0

    for line in lines:
        test, rest = line.split(":")

        if dfs(test := int(test), list(map(int, rest.split())), operators):
            ans += test

    return ans


def part1(input: str) -> int:
    return calibrate(input.splitlines(), [operator.add, operator.mul])


def part2(input: str) -> int:
    return calibrate(
        input.splitlines(),
        [operator.add, operator.mul, lambda x, y: int(str(x) + str(y))],
    )


if __name__ == "__main__":
    util.run(7, part1, part2)
