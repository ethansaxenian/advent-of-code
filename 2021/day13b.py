from day13a import parse_input, do_folds, count_dots

if __name__ == '__main__':
    coords, folds = parse_input()
    do_folds(coords, folds)
    xmax = max(coords, key=lambda pair: pair[0])[0] + 1
    ymax = max(coords, key=lambda pair: pair[1])[1] + 1

    grid = [[False for _ in range(xmax)] for _ in range(ymax)]

    for x, y in coords:
        grid[y][x] = True

    for y in range(ymax):
        for x in range(xmax):
            print("#" if grid[y][x] else ".", end=" ")
        print("\n")
