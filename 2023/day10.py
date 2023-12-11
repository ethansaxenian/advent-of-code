from util.util import fetch_input

puzzle_input = fetch_input(10)


start = (-1, -1)


for r, line in enumerate(puzzle_input):
    for c, char in enumerate(line):
        if char == "S":
            start = (r, c)


dirs = {
    (-1, 0): "|7F",
    (1, 0): "|JL",
    (0, -1): "-FL",
    (0, 1): "-J7",
}

char_dirs = {
    "S": [(-1, 0), (1, 0), (0, -1), (0, 1)],
    "7": [(1, 0), (0, -1)],
    "F": [(1, 0), (0, 1)],
    "L": [(-1, 0), (0, 1)],
    "J": [(-1, 0), (0, -1)],
    "|": [(-1, 0), (1, 0)],
    "-": [(0, -1), (0, 1)],
}


def find_char_for_s(loop):
    r, c = start
    north = puzzle_input[r - 1][c]
    south = puzzle_input[r + 1][c]
    east = puzzle_input[r][c + 1]
    west = puzzle_input[r][c - 1]
    if north in "|7F" and south in "|JL":
        return "|"
    elif east in "-J7" and west in "-FL":
        return "-"
    elif north in "|7F" and east in "-J7":
        return "L"
    elif north in "|7F" and west in "-FL":
        return "J"
    elif south in "|JL" and east in "-J7":
        return "F"
    elif south in "|JL" and west in "-FL":
        return "7"


def get_neighbors(r, c):
    neighbors = []
    char = puzzle_input[r][c]
    for dr, dc in char_dirs[char]:
        if (
            r + dr >= 0
            and c + dc >= 0
            and r + dr < len(puzzle_input)
            and c + dc < len(puzzle_input[0])
        ):
            if puzzle_input[r + dr][c + dc] in dirs[(dr, dc)]:
                neighbors.append((r + dr, c + dc))
    return neighbors


def dfs():
    loop = []
    stack = [start]
    while stack:
        r, c = stack.pop()
        if (r, c) in loop:
            continue
        loop.append((r, c))
        for nr, nc in get_neighbors(r, c):
            if (nr, nc) not in loop:
                stack.append((nr, nc))
    return loop


def enclosed_by_loop(r, c, loop, s):
    if (r, c) in loop:
        return False

    n = 0
    last = puzzle_input[r][c]
    for dr in range(r + 1, len(puzzle_input)):
        char = s if puzzle_input[dr][c] == "S" else puzzle_input[dr][c]
        if (dr, c) in loop and char != "|":
            match (last, char):
                case ("7", "L"):
                    n += 1
                case ("F", "J"):
                    n += 1
                case (_, "-"):
                    n += 1

            last = char

    return n % 2 == 1


def part1():
    loop = dfs()
    print(len(loop) // 2)


def part2():
    loop = dfs()
    s = find_char_for_s(loop)
    n = 0
    for r in range(len(puzzle_input)):
        for c in range(len(puzzle_input[0])):
            if enclosed_by_loop(r, c, loop, s):
                n += 1
    print(n)


part1()
part2()
