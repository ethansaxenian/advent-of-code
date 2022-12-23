from dataclasses import dataclass
from math import prod


@dataclass
class Loc:
    val: int
    low: bool = False
    marked: bool = False
    neighbors: list = None


def setup_grid():
    grid = []

    with open("day9_input.txt", "r") as file:
        for line in file:
            grid.append([Loc(val=int(c)) for c in line.strip()])

    for y in range(len(grid)):
        for x in range(len(grid[0])):
            loc = grid[y][x]
            loc.neighbors = get_neighbors(grid, x, y)
            loc.low = all(loc.val < n.val for n in loc.neighbors)

    return grid


def get_neighbors(grid, x, y):
    neighbors = []
    if y > 0:
        neighbors.append(grid[y - 1][x])
    if y < len(grid) - 1:
        neighbors.append(grid[y + 1][x])
    if x > 0:
        neighbors.append(grid[y][x - 1])
    if x < len(grid[0]) - 1:
        neighbors.append(grid[y][x + 1])

    return neighbors


def get_basin_size(loc):
    basin_size = 1
    loc.marked = True
    for l in loc.neighbors:
        if not l.marked and loc.val < l.val < 9:
            basin_size += get_basin_size(l)

    return basin_size


if __name__ == '__main__':
    basins = []

    grid = setup_grid()

    for y in range(len(grid)):
        for x in range(len(grid[0])):
            loc = grid[y][x]
            if loc.low:
                basins.append(get_basin_size(loc))

    print(prod(sorted(basins, reverse=True)[:3]))
