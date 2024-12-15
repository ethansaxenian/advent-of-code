import util

move_map = {"<": (0, -1), "^": (-1, 0), ">": (0, 1), "v": (1, 0)}


def part1(input: str) -> int:
    g, moves = input.split("\n\n")
    grid = {}
    fish = (-1, -1)
    for r, line in enumerate(g.splitlines()):
        for c, char in enumerate(line):
            grid[(r, c)] = char
            if char == "@":
                fish = (r, c)

    for dr, dc in [move_map[m] for m in "".join(moves.strip().splitlines())]:
        new_fish = fish[0] + dr, fish[1] + dc

        if grid[new_fish] == "#":
            continue

        elif grid[new_fish] == "O":
            boxes_in_the_way = set()
            xr, xc = new_fish
            while True:
                match grid[(xr, xc)]:
                    case "O":
                        boxes_in_the_way.add((xr, xc))
                    case "#":
                        break
                    case ".":
                        for br, bc in boxes_in_the_way:
                            grid[(br + dr, bc + dc)] = "O"
                        grid[fish] = "."
                        grid[new_fish] = "@"
                        fish = new_fish
                        break
                    case _:
                        pass

                xr += dr
                xc += dc

        elif grid[new_fish] == ".":
            grid[fish] = "."
            grid[new_fish] = "@"
            fish = new_fish

    return sum(100 * r + c for (r, c), v in grid.items() if v == "O")


def get_boxes(grid, dir, pos) -> set[tuple[int, int, str]]:
    dr, dc = dir
    stack = [(pos)]
    boxes = set()
    seen = set()
    while stack:
        r, c = stack.pop()
        x = grid[(r, c)]
        if x == "#":
            return set()

        if x == "." or (r, c) in seen:
            continue
        seen.add((r, c))

        boxes.add((r, c, x))

        stack.append((r + dr, c + dc))
        if dc == 0:
            if x == "[":
                boxes.add((r, c + 1, "]"))
                stack.append((r + dr, c + dc + 1))
            elif x == "]":
                boxes.add((r, c - 1, "["))
                stack.append((r + dr, c + dc - 1))

    return boxes


def part2(input: str) -> int:
    g, moves = input.split("\n\n")
    g = g.replace("#", "##").replace(".", "..").replace("@", "@.").replace("O", "[]")
    grid = {}
    fish = (-1, -1)
    for r, line in enumerate(g.splitlines()):
        for c, char in enumerate(line):
            grid[(r, c)] = char
            if char == "@":
                fish = (r, c)

    for dr, dc in [move_map[m] for m in "".join(moves.strip().splitlines())]:
        new_fish = fish[0] + dr, fish[1] + dc

        if grid[new_fish] == "#":
            continue

        elif grid[new_fish] in "[]":
            if boxes_in_the_way := get_boxes(grid, (dr, dc), new_fish):
                if dc == 0:
                    for br, bc, c in boxes_in_the_way:
                        if c == "[":
                            grid[(br, bc + 1)] = "."
                        if c == "]":
                            grid[(br, bc - 1)] = "."

                for br, bc, c in boxes_in_the_way:
                    grid[(br + dr, bc + dc)] = c

                grid[fish] = "."
                grid[new_fish] = "@"
                fish = new_fish

        elif grid[new_fish] == ".":
            grid[fish] = "."
            grid[new_fish] = "@"
            fish = new_fish

    return sum(100 * r + c for (r, c), v in grid.items() if v == "[")


if __name__ == "__main__":
    util.run(15, part1, part2)
