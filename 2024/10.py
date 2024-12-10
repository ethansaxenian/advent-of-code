import util


def part1(input: str) -> int:
    trailheads = []
    grid = {}
    for r, line in enumerate(input.splitlines()):
        for c, char in enumerate(line):
            grid[(r, c)] = int(char)
            if char == "0":
                trailheads.append((r, c))

    score = 0
    for start in trailheads:
        vis = set()
        ends = set()
        stack = [start]

        while stack:
            pos = stack.pop()

            if (height := grid.get(pos)) is None:
                continue

            vis.add(pos)
            if height == 9:
                ends.add(pos)
                continue

            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                new = (pos[0] + dr, pos[1] + dc)
                if new not in vis and grid.get(new) == height + 1:
                    stack.append(new)

        score += len(ends)

    return score


def part2(input: str) -> int:
    trailheads = []
    grid = {}
    for r, line in enumerate(input.splitlines()):
        for c, char in enumerate(line):
            grid[(r, c)] = int(char)
            if char == "0":
                trailheads.append((r, c))

    score = 0
    for start in trailheads:
        paths = []
        stack = [(start, [start])]

        while stack:
            pos, path = stack.pop()

            if (height := grid.get(pos)) is None:
                continue

            if height == 9:
                paths.append(path)
                continue

            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                new = (pos[0] + dr, pos[1] + dc)
                if grid.get(new) == height + 1:
                    stack.append((new, path + [new]))

        score += len(paths)

    return score


if __name__ == "__main__":
    util.run(10, part1, part2)
