GRID_SIZE = 1000


def init_grid():
    grid = [[0 for _ in range(GRID_SIZE)] for _ in range(GRID_SIZE)]

    with open("day5_input.txt", "r") as file:
        for line in file:
            ((x0, y0), (x1, y1)) = [a.strip().split(",") for a in line.split("->")]
            x0 = int(x0)
            x1 = int(x1)
            y0 = int(y0)
            y1 = int(y1)

            if (x0 == x1) or (y0 == y1):
                x0, x1 = sorted([x0, x1])
                y0, y1 = sorted([y0, y1])

                for y in range(y0, y1 + 1):
                    for x in range(x0, x1 + 1):
                        grid[y][x] += 1
            else:
                xi = -1 if x0 > x1 else 1
                yi = -1 if y0 > y1 else 1

                for i in range(0, (x0-x1 if x0 > x1 else x1-x0) + 1):
                    grid[y0 + i*yi][x0 + i*xi] += 1
    return grid


def calculate_overlaps(grid):
    n = 0
    for y in range(GRID_SIZE):
        for x in range(GRID_SIZE):
            if grid[y][x] >= 2:
                n += 1

    return n


if __name__ == '__main__':
    grid = init_grid()
    print(calculate_overlaps(grid))
