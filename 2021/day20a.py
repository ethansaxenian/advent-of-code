import copy


SYMBOL_TO_BIT = {
    ".": "0",
    "#": "1"
}


def add_blanks(grid, c):
    for row in grid:
        row.insert(0, c)
        row.append(c)
    blanks = [c for _ in range(len(grid[0]))]
    grid.insert(0, blanks[:])
    grid.append(blanks[:])


def get_output_pixel(grid, x, y, algorithm):
    try:
        sequence = ""
        for xi, yi in [(-1, -1), (0, -1), (1, -1), (-1, 0), (0, 0), (1, 0), (-1, 1), (0, 1), (1, 1)]:
            sequence += SYMBOL_TO_BIT[grid[y + yi][x + xi]]
        return algorithm[int(sequence, 2)]
    except IndexError:
        if grid[y][x] == ".":
            return algorithm[0]
        if grid[y][x] == "#":
            return algorithm[-1]


def enhance(grid, algorithm):
    new_grid = copy.deepcopy(grid)
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            new_pixel = get_output_pixel(grid, x, y, algorithm)
            new_grid[y][x] = new_pixel

    return new_grid


if __name__ == '__main__':
    with open("day20_input.txt", "r") as file:
        lines = [l.strip() for l in file.readlines()]
        algorithm = lines[0]
        output = [list(l) for l in lines[2:]]

    g = output
    on = False
    for _ in range(2):
        for _ in range(3):
            add_blanks(output, "#" if on else ".")
        g = enhance(g, algorithm)
        on = not on

    t = 0
    for l in g:
        for c in l:
            if c == "#":
                t += 1

    print(t)
