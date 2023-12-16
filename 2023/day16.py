from util.util import fetch_input

puzzle_input = fetch_input(16)

# puzzle_input = [
#     ".|...\\....",
#     "|.-.\\.....",
#     ".....|-...",
#     "........|.",
#     "..........",
#     ".........\\",
#     "..../.\\\\..",
#     ".-.-/..|..",
#     ".|....-|.\\",
#     "..//.|....",
# ]

# print("\n".join(puzzle_input))


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
            case ".", (_, _):
                beams.append((r + d[0], c + d[1], d))
            case "|", (_, 0):
                beams.append((r + d[0], c + d[1], d))
            case "|", (0, _):
                beams.append((r - 1, c, (-1, 0)))
                beams.append((r + 1, c, (1, 0)))
            case "-", (0, _):
                beams.append((r + d[0], c + d[1], d))
            case "-", (_, 0):
                beams.append((r, c - 1, (0, -1)))
                beams.append((r, c + 1, (0, 1)))
            case "\\", (_, _):
                beams.append((r + d[1], c + d[0], (d[1], d[0])))
            case "/", (_, _):
                beams.append((r - d[1], c - d[0], (-d[1], -d[0])))

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
