import functools
from util.util import fetch_input

input_lines = fetch_input(12)


@functools.cache
def count_arrangements(row: str, nums: tuple[int]) -> int:
    if row == "" and len(nums) > 0:
        return 0

    if row == "" and len(nums) == 0:
        return 1

    if len(nums) == 0 and "#" in row:
        return 0

    if len(nums) == 0 and "#" not in row:
        return 1

    if row[0] == ".":
        return count_arrangements(row[1:], nums)

    if row[0] == "?":
        return count_arrangements(row[1:], nums) + count_arrangements(
            "#" + row[1:], nums
        )

    curr_group_len, *rest = nums

    if row[0] == "#":
        first_group_works = "." not in row[:curr_group_len]
        hash_is_not_next = (
            len(row) == curr_group_len
            or len(row) > curr_group_len
            and row[curr_group_len] != "#"
        )
        if first_group_works and hash_is_not_next:
            return count_arrangements(row[curr_group_len + 1 :], tuple(rest))

    return 0


def part1():
    n = 0
    for row, nums in [line.split() for line in input_lines]:
        n += count_arrangements(row, tuple(map(int, nums.split(","))))

    print(n)


def part2():
    n = 0
    for row, nums in [line.split() for line in input_lines]:
        row = "?".join([row] * 5)
        nums = tuple(map(int, nums.split(","))) * 5
        n += count_arrangements(row, nums)

    print(n)


part1()
part2()
