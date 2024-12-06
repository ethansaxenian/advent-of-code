import util


def part1(input: str) -> int:
    lines = input.splitlines()
    start = None
    for r in range(len(lines)):
        for c in range(len(lines[0])):
            if lines[r][c] == "^":
                start = (r, c)

    pos = set()
    r, c = start
    dr, dc = -1, 0
    while True:
        pos.add((r, c))
        try:
            if lines[r + dr][c + dc] == "#":
                dr, dc = dc, -dr
            else:
                r += dr
                c += dc
        except IndexError:
            return len(pos)


def has_loop(grid, h, w, start) -> bool:
    pos = set()
    r, c = start
    dr, dc = -1, 0
    while True:
        if (r, c, dr, dc) in pos:
            return True
        pos.add((r, c, dr, dc))

        match grid.get((r + dr, c + dc)):
            case None:
                return False
            case "#":
                dr, dc = dc, -dr
            case _:
                r += dr
                c += dc


def part2(input: str) -> int:
    lines = input.splitlines()
    h, w = len(lines), len(lines[0])
    grid = {}
    start = None
    for r in range(h):
        for c in range(w):
            if lines[r][c] == "^":
                start = (r, c)
            grid[(r, c)] = lines[r][c]

    pos = set()
    r, c = start
    dr, dc = -1, 0
    while True:
        pos.add((r, c))
        try:
            if lines[r + dr][c + dc] == "#":
                dr, dc = dc, -dr
            else:
                r += dr
                c += dc
        except IndexError:
            break

    x = 0
    for block in pos:
        if block == start:
            continue

        m, grid[block] = grid[block], "#"
        if has_loop(grid, h, w, start):
            x += 1
        grid[block] = m

    return x


if __name__ == "__main__":
    util.run(6, part1, part2)
