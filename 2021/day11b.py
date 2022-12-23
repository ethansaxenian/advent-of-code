from dataclasses import dataclass


@dataclass
class Octopus:
    level: int
    flashed: bool = False
    neighbors: list = None


def get_neighbors(grid, x, y):
    neighbors = []
    for yi in range(-1, 2):
        for xi in range(-1, 2):
            if not (xi == 0 == yi):
                try:
                    if y+yi >= 0 and x+xi >= 0:
                        neighbors.append(grid[y+yi][x+xi])
                except IndexError:
                    continue

    return neighbors


grid = []
with open("day11_input.txt", "r") as file:
    for line in file:
        grid.append([Octopus(level=int(i)) for i in line.strip()])

    for y in range(len(grid)):
        for x in range(len(grid[0])):
            grid[y][x].neighbors = get_neighbors(grid, x, y)


def flash(oct):
    flashes = 0
    if not oct.flashed:
        flashes += 1
        oct.flashed = True
        for o in oct.neighbors:
            o.level += 1
            if o.level > 9:
                flashes += flash(o)

    return flashes


step = 0
not_synchronized = True

while not_synchronized:
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            grid[y][x].level += 1

    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if grid[y][x].level > 9:
                flash(grid[y][x])

    not_synchronized = False
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if not grid[y][x].flashed:
                not_synchronized = True

    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if grid[y][x].flashed:
                grid[y][x].flashed = False
                grid[y][x].level = 0

    step += 1

print(step)
