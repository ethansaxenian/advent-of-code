from util.util import fetch_input
from collections import defaultdict, deque

puzzle_input = fetch_input(23)

slopes = {
    ">": [(0, 1)],
    "<": [(0, -1)],
    "^": [(-1, 0)],
    "v": [(1, 0)],
    ".": [(-1, 0), (1, 0), (0, -1), (0, 1)],
}

start = next((0, i) for i, c in enumerate(puzzle_input[0]) if c == ".")
end = next(
    (len(puzzle_input) - 1, i) for i, c in enumerate(puzzle_input[-1]) if c == "."
)


intersections = {start, end}
for r in range(len(puzzle_input)):
    for c in range(len(puzzle_input[0])):
        if puzzle_input[r][c] == "#":
            continue

        n = 0
        for dr, dc in [(0, 1), (0, -1), (-1, 0), (1, 0)]:
            if (
                0 <= r + dr < len(puzzle_input)
                and 0 <= c + dc < len(puzzle_input[0])
                and puzzle_input[r + dr][c + dc] != "#"
            ):
                n += 1

        if n >= 3:
            intersections.add((r, c))


def dfs(consider_slopes: bool):
    graph = defaultdict(dict)

    for r, c in intersections:
        q: deque[tuple[int, int, int]] = deque([(r, c, 0)])
        seen = set()

        while q:
            nr, nc, steps = q.popleft()

            if (nr, nc) in seen:
                continue

            seen.add((nr, nc))

            if steps > 0 and (nr, nc) in intersections:
                graph[(r, c)][(nr, nc)] = steps
                continue

            dirs = slopes[puzzle_input[nr][nc]] if consider_slopes else slopes["."]
            for dr, dc in dirs:
                if (
                    0 <= nr + dr < len(puzzle_input)
                    and 0 <= nc + dc < len(puzzle_input[0])
                    and puzzle_input[nr + dr][nc + dc] != "#"
                ):
                    q.append((nr + dr, nc + dc, steps + 1))

    stack = [(start, 0, {start})]

    m = 0

    while stack:
        curr, steps, seen = stack.pop()

        if curr == end:
            m = max(m, steps)
            continue

        for node, dist in graph[curr].items():
            if node not in seen:
                stack.append((node, steps + dist, seen | {node}))

    print(m)


dfs(True)
dfs(False)
