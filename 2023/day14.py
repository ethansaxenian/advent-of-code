from util.util import fetch_input

puzzle_input = fetch_input(14)

load = 0
for c in range(len(puzzle_input[0])):
    prev_rock = -1
    for r in range(len(puzzle_input)):
        match puzzle_input[r][c]:
            case "#":
                prev_rock = r
            case "O":
                prev_rock += 1
                load += len(puzzle_input) - prev_rock


print(load)

rounds = set()
cubes = set()

for r in range(len(puzzle_input)):
    for c in range(len(puzzle_input[0])):
        match puzzle_input[r][c]:
            case "#":
                cubes.add((r, c))
            case "O":
                rounds.add((r, c))


def tilt_north(rounds, cubes):
    new_rounds = set()
    for c in range(len(puzzle_input[0])):
        prev_rock = -1
        for r in range(len(puzzle_input)):
            if (r, c) in cubes:
                prev_rock = r
            elif (r, c) in rounds:
                prev_rock += 1
                new_rounds.add((prev_rock, c))

    return new_rounds


def tilt_west(rounds, cubes):
    new_rounds = set()
    for r in range(len(puzzle_input)):
        prev_rock = -1
        for c in range(len(puzzle_input[0])):
            if (r, c) in cubes:
                prev_rock = c
            elif (r, c) in rounds:
                prev_rock += 1
                new_rounds.add((r, prev_rock))

    return new_rounds


def tilt_south(rounds, cubes):
    new_rounds = set()
    for c in range(len(puzzle_input[0])):
        prev_rock = len(puzzle_input)
        for r in range(len(puzzle_input) - 1, -1, -1):
            if (r, c) in cubes:
                prev_rock = r
            elif (r, c) in rounds:
                prev_rock -= 1
                new_rounds.add((prev_rock, c))

    return new_rounds


def tilt_east(rounds, cubes):
    new_rounds = set()
    for r in range(len(puzzle_input)):
        prev_rock = len(puzzle_input[0])
        for c in range(len(puzzle_input[0]) - 1, -1, -1):
            if (r, c) in cubes:
                prev_rock = c
            elif (r, c) in rounds:
                prev_rock -= 1
                new_rounds.add((r, prev_rock))

    return new_rounds


def print_grid(rounds, cubes):
    for r in range(len(puzzle_input)):
        for c in range(len(puzzle_input[0])):
            if (r, c) in cubes:
                print("#", end="")
            elif (r, c) in rounds:
                print("O", end="")
            else:
                print(".", end="")
        print()
    print()


seen = {tuple(rounds)}
arr = []


i = 0
while True:
    rounds = tilt_north(rounds, cubes)
    rounds = tilt_west(rounds, cubes)
    rounds = tilt_south(rounds, cubes)
    rounds = tilt_east(rounds, cubes)

    i += 1
    if tuple(rounds) in seen:
        break

    arr.append(rounds)
    seen.add(tuple(rounds))

j = 0
while arr[j] != rounds:
    j += 1


cycle_length = i - (j + 1)
k = j + 1
while k + cycle_length < 1_000_000_000:
    k += cycle_length

rounds = arr[j + 1_000_000_000 - k]


print(sum(len(puzzle_input) - r for r, _ in rounds))
