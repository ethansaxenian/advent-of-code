import json
from pathlib import Path
from urllib.request import Request, urlopen


def fetch_input(day: int) -> list[str]:
    url = f"https://adventofcode.com/2023/day/{day}/input"

    curr_dir = Path(__file__).parent
    with open(f"{curr_dir.parent.parent / 'aoc-cookie.json'}") as f:
        cookie = json.load(f)["aoc_cookie"]

    req = Request(url)
    req.add_header("Cookie", f"session={cookie}")

    with urlopen(req) as response:
        puzzle_input = response.read().decode("utf-8")

    return puzzle_input.split("\n")
