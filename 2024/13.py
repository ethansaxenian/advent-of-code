import re

import util


def solve(a, b, prize) -> int:
    ax, ay = a
    bx, by = b
    px, py = prize
    """
    ax * A + bx * B = px
    ay * A + by * B = py

    A = (px - bx*B) / ax
    A = (py - by*B) / ay
    (px - bx*B) / ax = (py - by*B) / ay
    ay * (px - bx*B) = ax * (py - by*B)
    ay*px - ay*bx*B = ax*py - ax*by*B
    ax*by*B - ay*bx*B = ax*py - ay*px
    B * (ax*by - ay*bx) =ax*py - ay*px
    """
    B = (ax * py - ay * px) / (ax * by - ay * bx)
    A = (px - bx * B) / ax

    if A.is_integer() and B.is_integer():
        return int(3 * A + B)

    return 0


def part1(input: str) -> int:
    cost = 0
    for machine in input.split("\n\n"):
        a, b, prize = map(
            lambda line: map(int, re.findall(r"\d+", line)), machine.splitlines()
        )
        cost += solve(a, b, prize)

    return cost


def part2(input: str) -> int:
    cost = 0
    for machine in input.split("\n\n"):
        a, b, prize = map(
            lambda line: map(int, re.findall(r"\d+", line)), machine.splitlines()
        )
        cost += solve(a, b, map(lambda p: p + 10000000000000, prize))

    return cost


if __name__ == "__main__":
    util.run(13, part1, part2)
