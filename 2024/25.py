import util


def part1(input: str) -> int:
    keys = []
    locks = []
    for line in input.split("\n\n"):
        heights = [row.count("#") for row in zip(*line.split())]
        if line[0] == ".":
            keys.append(heights)
        else:
            locks.append(heights)

    n = 0
    for l in locks:
        for k in keys:
            match = True
            for i, j in zip(l, k):
                if i + j > 7:
                    match = False
                    break
            if match:
                n += 1

    return n


def part2(input: str) -> int:
    return -1


if __name__ == "__main__":
    util.run(25, part1, part2)
