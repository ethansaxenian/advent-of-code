import re

import util

H, W = 103, 101


def part1(input: str) -> int:
    q1, q2, q3, q4 = 0, 0, 0, 0
    for pc, pr, vc, vr in map(
        lambda line: map(int, re.findall(r"-?\d+", line)), input.splitlines()
    ):
        c = (vc * 100 + pc) % W
        r = (vr * 100 + pr) % H

        if r < H // 2 and c < W // 2:
            q1 += 1
        elif r < H // 2 and c > W // 2:
            q2 += 1
        elif r > H // 2 and c > W // 2:
            q3 += 1
        elif r > H // 2 and c < W // 2:
            q4 += 1

    return q1 * q2 * q3 * q4


def part2(input: str) -> int:
    robots = []
    for line in input.splitlines():
        robots.append(map(int, re.findall(r"-?\d+", line)))

    num_robots = len(robots)

    i = 0
    while True:
        i += 1

        new_robots = []
        new_pos = set()
        for pc, pr, vc, vr in robots:
            c = (vc + pc) % W
            r = (vr + pr) % H
            new_robots.append((c, r, vc, vr))
            new_pos.add((c, r))

        if len(new_pos) == num_robots:
            return i

        robots = new_robots


if __name__ == "__main__":
    util.run(14, part1, part2)
