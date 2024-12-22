import sys
from collections import defaultdict, deque

import util


def new_secret(n):
    n ^= n * 64
    n %= 16777216
    n ^= n // 32
    n %= 16777216
    n ^= n * 2048
    n %= 16777216
    return n


def part1(input: str) -> int:
    x = 0
    for secret_number in map(int, input.splitlines()):
        n = secret_number
        for _ in range(2000):
            n = new_secret(n)
        x += n

    return x


def part2(input: str) -> int:
    sequences = defaultdict(int)
    for n in map(int, input.splitlines()):
        digits = deque([n % 10], maxlen=4)
        changes = deque([sys.maxsize], maxlen=4)
        seen = set()
        for _ in range(2000):
            n = new_secret(n)
            last_digit = digits[-1]
            new_last_digit = n % 10
            changes.append(new_last_digit - last_digit)
            digits.append(new_last_digit)

            if (c := tuple(changes)) not in seen:
                sequences[c] += digits[-1]
                seen.add(c)

    return max(sequences.values())


if __name__ == "__main__":
    util.run(22, part1, part2)
