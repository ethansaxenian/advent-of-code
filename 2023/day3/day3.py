from pathlib import Path
from urllib.request import Request, urlopen
import json
from dataclasses import dataclass, field

url = "https://adventofcode.com/2023/day/3/input"

curr_dir = Path(__file__).parent
with open(f"{curr_dir.parent.parent / 'aoc-cookie.json'}") as f:
    cookie = json.load(f)["aoc_cookie"]


req = Request(url)
req.add_header("Cookie", f"session={cookie}")

with urlopen(req) as response:
    puzzle_input = response.read().decode("utf-8")

rows = puzzle_input.split("\n")

Coord = tuple[int, int]


@dataclass
class Num:
    num: int
    coords: set[Coord]
    neighbors: set[Coord] = field(init=False, default_factory=set)
    is_part: bool = False

    def __post_init__(self):
        min_i = min(coord[0] for coord in self.coords)
        max_i = max(coord[0] for coord in self.coords)
        min_j = min(coord[1] for coord in self.coords)
        max_j = max(coord[1] for coord in self.coords)

        for i in range(min_i - 1, max_i + 2):
            for j in range(min_j - 1, max_j + 2):
                if (i, j) in self.coords:
                    continue
                try:
                    if not rows[i][j].isnumeric() and rows[i][j] != ".":
                        self.is_part = True
                    self.neighbors.add((i, j))
                except IndexError:
                    continue


numbers: list[Num] = []

for i, row in enumerate(rows):
    j = 0
    while j < len(row):
        if row[j].isnumeric():
            num = ""
            coords = set()
            while j < len(row) and row[j].isnumeric():
                coords.add((i, j))
                num += row[j]
                j += 1
            numbers.append(Num(int(num), coords))
        j += 1


def part1():
    print(sum(num.num for num in numbers if num.is_part))


def part2():
    sum = 0

    for i in range(len(rows)):
        for j in range(len(rows[i])):
            if rows[i][j] == "*":
                gear_coords = (i, j)
                adjacent_numbers: list[int] = []
                for num in numbers:
                    if gear_coords in num.neighbors:
                        adjacent_numbers.append(num.num)

                if len(adjacent_numbers) == 2:
                    sum += adjacent_numbers[0] * adjacent_numbers[1]

    print(sum)


part1()
part2()
