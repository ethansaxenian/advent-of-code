import util


def is_invalid_id_1(id: str) -> bool:
    if len(id) % 2 == 1:
        return False

    a, b = id[: len(id) // 2], id[len(id) // 2 :]
    return a == b


def part1(input: str) -> int:
    s = 0
    for r in input.split(","):
        a, b = map(int, r.split("-"))
        for i in range(a, b + 1):
            if is_invalid_id_1(str(i)):
                s += i

    return s


def is_invalid_id_2(id: str) -> bool:
    l = len(id)
    for i in range(1, l // 2 + 1):
        if l % i != 0:
            continue

        part = id[0:i]
        for j in range(i, l, i):
            if id[j : j + i] != part:
                break
        else:
            return True

    return False


def part2(input: str) -> int:
    s = 0
    for r in input.split(","):
        a, b = map(int, r.split("-"))
        for i in range(a, b + 1):
            if is_invalid_id_2(str(i)):
                s += i
    return s


if __name__ == "__main__":
    util.run(2, part1, part2)
