from util.util import fetch_input

puzzle_input = fetch_input(16)


def energize(r, c, d):
    energized = set()
    beams: list[tuple[int, int, tuple[int, int]]] = [(r, c, d)]
    seen = set()

    while beams:
        r, c, d = beams.pop()

        if r < 0 or c < 0 or r >= len(puzzle_input) or c >= len(puzzle_input[r]):
            continue

        if (r, c, d) in seen:
            continue

        energized.add((r, c))
        seen.add((r, c, d))

        match puzzle_input[r][c], d:
            case ".", (dr, dc):
                beams.append((r + dr, c + dc, d))
            case "|", (dr, 0):
                beams.append((r + dr, c, d))
            case "|", (0, dc):
                beams.append((r - dc, c, (-dc, 0)))
                beams.append((r + dc, c, (dc, 0)))
            case "-", (0, dc):
                beams.append((r, c + dc, d))
            case "-", (dr, 0):
                beams.append((r, c - dr, (0, -dr)))
                beams.append((r, c + dr, (0, dr)))
            case "\\", (dr, dc):
                beams.append((r + dc, c + dr, (dc, dr)))
            case "/", (dr, dc):
                beams.append((r - dc, c - dr, (-dc, -dr)))

    return len(energized)


print(energize(0, 0, (0, 1)))

m = -1
for r in range(len(puzzle_input)):
    m = max(m, energize(r, 0, (0, 1)))
    m = max(m, energize(r, len(puzzle_input[r]) - 1, (0, -1)))

for c in range(len(puzzle_input[0])):
    m = max(m, energize(0, c, (1, 0)))
    m = max(m, energize(len(puzzle_input) - 1, c, (-1, 0)))

print(m)
