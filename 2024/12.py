import util


def search(start, grid) -> tuple[set[tuple[int, int]], int]:
    p = grid[start]
    stack = [start]

    vis = set()
    perimeter = 0

    while stack:
        pos = stack.pop()
        if pos in vis:
            continue

        if grid.get(pos) != p:
            perimeter += 1
            continue

        vis.add(pos)

        r, c = pos
        for dr, dc in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
            stack.append((r + dr, c + dc))

    return vis, perimeter


def part1(input: str) -> int:
    grid = {}
    for r, line in enumerate(input.splitlines()):
        for c, char in enumerate(line):
            grid[(r, c)] = char

    coords = set(grid.keys())
    price = 0
    while coords:
        pos = coords.pop()

        plot, perimeter = search(pos, grid)
        coords -= plot

        price += len(plot) * perimeter

    return price


def part2(input: str) -> int:
    grid = {}
    for r, line in enumerate(input.splitlines()):
        for c, char in enumerate(line):
            grid[(r, c)] = char

    coords = set(grid.keys())
    price = 0
    while coords:
        pos = coords.pop()

        plot, _ = search(pos, grid)
        coords -= plot

        plant = grid[pos]

        vertices = 0
        for r, c in plot:
            for dr, dc in [(1, 1), (1, -1), (-1, 1), (-1, -1)]:
                x = grid.get((r + dr, c))
                y = grid.get((r + dr, c + dc))
                z = grid.get((r, c + dc))

                if (
                    (x != plant and z != plant)
                    or (x != plant and z != plant)
                    or (x == plant and y != plant and z == plant)
                ):
                    vertices += 1

        price += len(plot) * vertices

    return price


if __name__ == "__main__":
    util.run(12, part1, part2)
