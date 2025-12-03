import util


def part1(input: str) -> int:
    ans = 0
    d = 50
    for line in input.splitlines():
        dir, rot = line[0], int(line[1:])
        if dir == "L":
            d -= rot
        if dir == "R":
            d += rot

        d %= 100
        if d == 0:
            ans += 1

    return ans


def part2(input: str) -> int:
    ans = 0
    d = 50
    for line in input.splitlines():
        dir, rot = line[0], int(line[1:])
        for _ in range(rot):
            if dir == "L":
                d -= 1
            if dir == "R":
                d += 1

            d %= 100

            if d == 0:
                ans += 1

    return ans


if __name__ == "__main__":
    util.run(1, part1, part2)
