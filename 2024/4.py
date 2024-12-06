import util

DIRS = [(0, 1), (1, 1), (1, 0), (1, -1), (0, -1), (-1, -1), (-1, 0), (-1, 1)]


def check(grid, r, c, dr, dc) -> bool:
    for i, x in enumerate("MAS", start=1):
        nr, nc = r + dr * i, c + dc * i

        if nr < 0 or nr >= len(grid) or nc < 0 or nc >= len(grid[0]):
            return False

        if grid[r + dr * i][c + dc * i] != x:
            return False

    return True


def part1(input: str) -> int:
    s = 0
    lines = input.splitlines()
    for r, line in enumerate(lines):
        for c, char in enumerate(line):
            if char == "X":
                for x, y in DIRS:
                    if check(lines, r, c, x, y):
                        s += 1

    return s


def part2(input: str) -> int:
    s = 0
    lines = input.splitlines()
    for r in range(1, len(lines) - 1):
        for c in range(1, len(lines[r]) - 1):
            char = lines[r][c]
            if char != "A":
                continue

            x1 = {char, lines[r - 1][c - 1], lines[r + 1][c + 1]}
            x2 = {char, lines[r + 1][c - 1], lines[r - 1][c + 1]}
            if x1 == x2 == set("MAS"):
                s += 1

    return s


if __name__ == "__main__":
    util.run(4, part1, part2)
