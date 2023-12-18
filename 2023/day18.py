from util.util import fetch_input

puzzle_input = fetch_input(18)

holes: set[tuple[int, int]] = {(0, 0)}
y, x = 0, 0
for line in puzzle_input:
    dir, meters, _ = line.split()
    for _ in range(int(meters)):
        match dir:
            case "R":
                y, x = y, x + 1
            case "L":
                y, x = y, x - 1
            case "U":
                y, x = y - 1, x
            case "D":
                y, x = y + 1, x

        holes.add((y, x))

min_r, min_c = min(holes, key=lambda x: x[0])[0], min(holes, key=lambda x: x[1])[1]
max_r, max_c = max(holes, key=lambda x: x[0])[0], max(holes, key=lambda x: x[1])[1]


class OutOfBoundsException(Exception):
    pass


def dfs(r, c):
    vis = set()
    stack = [(r, c)]

    while stack:
        r, c = stack.pop()

        if r < min_r or r > max_r or c < min_c or c > max_c:
            raise OutOfBoundsException()

        if (r, c) in vis or (r, c) in holes:
            continue

        vis.add((r, c))

        for dr, dc in ((0, 1), (0, -1), (1, 0), (-1, 0)):
            stack.append((r + dr, c + dc))

    return vis


for y, x in ((0, 1), (0, -1), (1, 0), (-1, 0), (-1, -1), (-1, 1), (1, -1), (1, 1)):
    if (y, x) in holes or y < min_r or y > max_r or x < min_c or x > max_c:
        continue

    try:
        inner = dfs(y, x)
        print(len(inner) + len(holes))
        break
    except OutOfBoundsException:
        continue


vertices: list[tuple[int, int]] = [(0, 0)]
y, x = 0, 0
perimeter = 0
for line in puzzle_input:
    _, _, color = line.split()
    meters = int(color[2:-2], 16)
    perimeter += meters
    dir = int(color[-2])

    match dir:
        case 0:
            y, x = y, x + meters
        case 2:
            y, x = y, x - meters
        case 3:
            y, x = y - meters, x
        case 1:
            y, x = y + meters, x

    vertices.append((x, y))


def shoelace(vertices, perimiter):
    area = 0
    for i in range(len(vertices) - 1):
        area += vertices[i][0] * vertices[i + 1][1]
        area -= vertices[i][1] * vertices[i + 1][0]
    area += perimiter

    return abs(area // 2 + 1)


print(shoelace(vertices, perimeter))
