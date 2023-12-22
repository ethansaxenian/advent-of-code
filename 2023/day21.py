from util.util import fetch_input
import math
from collections import deque, defaultdict

puzzle_input = fetch_input(21)

puzzle_input = [
    "...........",
    ".....###.#.",
    ".###.##..#.",
    "..#.#...#..",
    "....#.#....",
    ".##..S####.",
    ".##..#...#.",
    ".......##..",
    ".##.#.####.",
    ".##..##.##.",
    "...........",
]

rocks = set()

start = (0, 0)

for r, row in enumerate(puzzle_input):
    for c, char in enumerate(row):
        if char == "#":
            rocks.add((r, c))

        if char == "S":
            start = (r, c)


q: deque[tuple[int, int, int]] = deque([(start[0], start[1], 0)])
final = set()

while q:
    r, c, steps = q.popleft()

    if steps == 64:
        final.add((r, c))
        continue

    for dr, dc in [(0, 1), (1, 0), (-1, 0), (0, -1)]:
        nr, nc = r + dr, c + dc

        if nr < 0 or nc < 0:
            continue

        if nr > len(puzzle_input) - 1 or nc > len(puzzle_input[0]) - 1:
            continue

        if (nr, nc) in rocks:
            continue

        if (nr, nc, steps + 1) in q:
            continue

        q.append((nr, nc, steps + 1))

# print(len(final))

q: deque[tuple[int, int, int]] = deque([(start[0], start[1], 0)])
final = set()
seen = defaultdict(lambda: math.inf)

while q:
    r, c, steps = q.popleft()
    print(r, c, steps, len(seen), len(q))

    if steps == 50:
        final.add((r, c))
        continue

    for dr, dc in [(0, 1), (1, 0), (-1, 0), (0, -1)]:
        nr, nc = r + dr, c + dc

        if (nr % len(puzzle_input), nc % len(puzzle_input[0])) in rocks:
            continue

        if seen[(nr, nc)] <= steps + 1:
            continue

        q.append((nr, nc, steps + 1))

print(len(final))
