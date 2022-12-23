

def parse_input():
    coords = []
    folds = []

    with open("day13_input.txt", "r") as file:
        for line in file:
            if line != "\n" and not line.startswith("fold"):
                coords.append([int(i) for i in line.strip().split(",")])
            if line.startswith("fold"):
                folds.append(line.strip().removeprefix("fold along ").split("="))

    return coords, folds


def do_folds(coords, folds):
    for axis, n in folds:
        i = 0 if axis == "x" else 1
        n = int(n)
        for pair in coords:
            value = int(pair[i])
            if value > n:
                pair[i] = value - 2 * (value - n)


def count_dots(coords):
    x = 0
    for i, pair in enumerate(coords):
        if pair in coords[i+1:]:
            x += 1

    return len(coords) - x


if __name__ == '__main__':
    coords, folds = parse_input()
    do_folds(coords, [folds[0]])
    print(count_dots(coords))
