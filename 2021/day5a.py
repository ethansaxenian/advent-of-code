GRID_SIZE = 1000


def init_grid():
    grid = [[0 for _ in range(GRID_SIZE)] for _ in range(GRID_SIZE)]

    with open("day5_input.txt", "r") as file:
        for line in file:
            ((x0, y0), (x1, y1)) = [a.strip().split(",") for a in line.split("->")]
            if (x0 == x1) or (y0 == y1):
                for y in range(min(int(y0), int(y1)), max(int(y0), int(y1)) + 1):
                    for x in range(min(int(x0), int(x1)), max(int(x0), int(x1)) + 1):
                        grid[y][x] += 1

    return grid


def calculate_overlaps(grid):
    n = 0
    for y in range(GRID_SIZE):
        for x in range(GRID_SIZE):
            if grid[y][x] >= 2:
                n += 1

    return n


def print_grid(grid):
    for y in range(GRID_SIZE):
        for x in range(GRID_SIZE):
            if grid[y][x] == 0:
                print(".", end=" ")
                continue
            print(grid[y][x], end=" ")
        print("\n")


if __name__ == '__main__':
    grid = init_grid()
    print(calculate_overlaps(grid))
