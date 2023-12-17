from util.util import fetch_input
from heapq import heappush, heappop
import math
from collections import defaultdict

puzzle_input = fetch_input(17)


def dijkstras(straight_max: int, straight_min: int = 0):
    vis = set()
    q = []

    heats = defaultdict(lambda: math.inf)
    heats[(0, 0, 1, 0, 1)] = 0
    heats[(0, 0, 0, 1, 1)] = 0

    heappush(q, (0, 0, 0, 1, 0, 1))
    heappush(q, (0, 0, 0, 0, 1, 1))

    while q:
        heat, r, c, dr, dc, consecutive = heappop(q)

        if (r, c, dr, dc, consecutive) in vis:
            continue

        if consecutive > straight_max:
            continue

        if (r, c) == (
            len(puzzle_input) - 1,
            len(puzzle_input[0]) - 1,
        ) and consecutive >= straight_min:
            print(heat)
            break

        vis.add((r, c, dr, dc, consecutive))

        for nr, nc in ((0, 1), (0, -1), (1, 0), (-1, 0)):
            if (nr, nc) == (-dr, -dc):
                continue

            if (nr, nc) != (dr, dc):
                if consecutive < straight_min:
                    continue
                new_consecutive = 1
            else:
                new_consecutive = consecutive + 1

            new_r = r + nr
            new_c = c + nc

            if 0 <= new_r < len(puzzle_input) and 0 <= new_c < len(puzzle_input[0]):
                new_heat = heat + int(puzzle_input[new_r][new_c])
                if new_heat < heats[(new_r, new_c, nr, nc, new_consecutive)]:
                    heats[(new_r, new_c, nr, nc, new_consecutive)] = new_heat
                    heappush(q, (new_heat, new_r, new_c, nr, nc, new_consecutive))


dijkstras(3)
dijkstras(10, 4)
