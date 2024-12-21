import itertools
import sys
from collections import deque
from functools import cache

import util

numeric_keypad = {
    "A": {"<": "0", "^": "3"},
    "0": {">": "A", "^": "2"},
    "1": {">": "2", "^": "4"},
    "2": {">": "3", "^": "5", "<": "1", "v": "0"},
    "3": {"^": "6", "<": "2", "v": "A"},
    "4": {">": "5", "^": "7", "v": "1"},
    "5": {">": "6", "^": "8", "<": "4", "v": "2"},
    "6": {"^": "9", "<": "5", "v": "3"},
    "7": {">": "8", "v": "4"},
    "8": {">": "9", "<": "7", "v": "5"},
    "9": {"<": "8", "v": "6"},
}

directional_keypad = {
    "A": {"<": "^", "v": ">"},
    "^": {">": "A", "v": "v"},
    "<": {">": "v"},
    "v": {">": ">", "<": "<", "^": "^"},
    ">": {"<": "v", "^": "A"},
}

KEYPAD = int
NUMERIC_KEYPAD = 0
DIRECTIONAL_KEYPAD = 1


@cache
def bfs(k: KEYPAD, src: str, target: str) -> list[str]:
    keypad = numeric_keypad if k == NUMERIC_KEYPAD else directional_keypad
    paths = []
    seen = {}
    q = deque([(src, "")])
    while q:
        pos, presses = q.popleft()

        if pos == target:
            if not paths or len(presses) <= len(paths[0]):
                paths.append(presses)
            continue

        if pos in seen and len(presses) > seen[pos]:
            continue

        seen[pos] = len(presses)

        for m, next_pos in keypad[pos].items():
            q.append((next_pos, presses + m))

    return paths


@cache
def search(keypad: KEYPAD, target: str, r: int) -> int:
    if r == 0:
        return len(target)

    total_len = 0
    for a, b in itertools.pairwise("A" + target):
        paths = bfs(keypad, a, b)

        min_path_len = sys.maxsize
        for p in paths:
            min_path_len = min(min_path_len, search(DIRECTIONAL_KEYPAD, p + "A", r - 1))
        total_len += min_path_len

    return total_len


def part1(input: str) -> int:
    n = 0
    for line in input.splitlines():
        num_presses = search(NUMERIC_KEYPAD, line, 3)
        n += num_presses * int(line[:-1])
    return n


def part2(input: str) -> int:
    n = 0
    for line in input.splitlines():
        num_presses = search(NUMERIC_KEYPAD, line, 26)
        n += num_presses * int(line[:-1])
    return n


if __name__ == "__main__":
    util.run(21, part1, part2)
