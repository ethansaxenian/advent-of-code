from collections import deque

import z3

import util


def part1(input: str) -> int:
    res = 0

    for line in input.splitlines():
        raw_lights, *raw_buttons, _ = line.split(" ")

        lights = sum(2**i for i, c in enumerate(raw_lights[1:-1]) if c == "#")

        buttons = [
            sum(2 ** int(n) for n in part[1:-1].split(",")) for part in raw_buttons
        ]

        q = deque([(0, 0)])

        seen = set()

        while q:
            presses, curr = q.popleft()

            if curr == lights:
                res += presses
                break

            if curr in seen:
                continue

            seen.add(curr)

            for b in buttons:
                q.append((presses + 1, curr ^ b))

    return res


def part2(input: str) -> int:
    res = 0
    for i, line in enumerate(input.splitlines()):
        _, *raw_buttons, raw_joltage = line.split(" ")

        buttons = [[int(n) for n in part[1:-1].split(",")] for part in raw_buttons]

        joltage = tuple(
            int(n) for n in (raw_joltage.removeprefix("{").removesuffix("}").split(","))
        )

        s = z3.Optimize()

        # create variable for each button press
        presses = [z3.Int(i) for i in range(len(buttons))]

        # we cannot press a button a negative number of times
        for p in presses:
            s.add(p >= 0)

        # button presses must sum to each joltage value
        for i, jolt in enumerate(joltage):
            sum_presses = 0
            for j, button in enumerate(buttons):
                if i in button:
                    sum_presses += presses[j]
            s.add(sum_presses == jolt)

        s.minimize(sum(presses))
        assert s.check()

        m = s.model()
        for p in presses:
            res += m[p].as_long()  # pyright: ignore

    return res


if __name__ == "__main__":
    util.run(10, 2025, part1, part2)
