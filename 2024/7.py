from collections.abc import Callable

import util


def unmultiply(x: int, y: int) -> int:
    if x % y:
        raise ValueError
    return x // y


def unadd(x: int, y: int) -> int:
    return x - y


def unconcat(x: int, y: int) -> int:
    a, b = str(x), str(y)
    if not a.endswith(b):
        raise ValueError
    return int(a.removesuffix(b) or 0)


def dfs(test: int, nums: list[int], operators: list[Callable[[int, int], int]]) -> bool:
    stack = [(test, len(nums) - 1)]
    while stack:
        res, i = stack.pop()
        n = nums[i]
        if res < 0:
            continue

        if i == 0:
            if res == n:
                return True
            continue

        for op in operators:
            try:
                stack.append((op(res, n), i - 1))
            except ValueError:
                continue

    return False


def calibrate(lines: list[str], operators: list[Callable[[int, int], int]]) -> int:
    ans = 0

    for line in lines:
        test, rest = line.split(":")

        if dfs(test := int(test), list(map(int, rest.split())), operators):
            ans += test

    return ans


def part1(input: str) -> int:
    return calibrate(input.splitlines(), [unadd, unmultiply])


def part2(input: str) -> int:
    return calibrate(input.splitlines(), [unadd, unmultiply, unconcat])


if __name__ == "__main__":
    util.run(7, part1, part2)
