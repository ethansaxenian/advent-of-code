from functools import cache

import util


def part1(input: str) -> int:
    lines = input.splitlines()
    patterns = [x.strip() for x in lines[0].split(",")]

    n = 0
    for design in lines[2:]:
        s = [design]
        while s:
            d = s.pop()

            if d == "":
                n += 1
                break

            for p in patterns:
                if d.endswith(p):
                    s.append(d.removesuffix(p))

    return n


def part2(input: str) -> int:
    lines = input.splitlines()
    patterns = [x.strip() for x in lines[0].split(",")]

    @cache
    def _dfs(d) -> int:
        if d == "":
            return 1

        s = 0
        for p in patterns:
            if d.endswith(p):
                s += _dfs(d.removesuffix(p))

        return s

    n = 0
    for design in lines[2:]:
        n += _dfs(design)

    return n


if __name__ == "__main__":
    util.run(19, part1, part2)
