import argparse
import os
import subprocess
import sys
import time
import urllib.request
from collections.abc import Callable
from pathlib import Path, PurePath
from typing import Literal, cast

AOC_EMAIL = os.environ["AOC_EMAIL"]
AOC_COOKIE = os.environ["AOC_COOKIE"]

PROJECT_ROOT = PurePath(__file__).parent.parent

parser = argparse.ArgumentParser()

parser.add_argument(
    "-p",
    "--part",
    type=int,
    help="which part to run",
    default=1,
    choices=[1, 2],
)
parser.add_argument(
    "-t",
    "--test",
    help="use test input",
    action=argparse.BooleanOptionalAction,
)


class Args(argparse.Namespace):
    part: Literal[1, 2]
    test: bool


def fetch_input(day: int, year: int) -> str:
    input_dir = Path(PROJECT_ROOT / str(year) / "input")
    input_dir.mkdir(exist_ok=True)

    input_path = Path(input_dir / f"day{day}.txt")
    if input_path.exists():
        return input_path.read_text().strip()

    req = urllib.request.Request(
        f"https://adventofcode.com/{year}/day/{day}/input",
    )
    req.add_header("Cookie", f"session={AOC_COOKIE}")
    req.add_header(
        "User-Agent",
        f"github.com/ethansaxenian/advent-of-code/tree/main/{year} by {AOC_EMAIL}",
    )

    with urllib.request.urlopen(req) as response:
        puzzle_input = response.read().decode("utf-8").strip()

    input_path.write_text(puzzle_input)

    return puzzle_input.strip()


def run(
    day: int,
    year: int,
    part1: Callable[[str], str | int],
    part2: Callable[[str], str | int],
):
    args = cast(Args, parser.parse_args())

    if args.test:
        input = sys.stdin.read()
    else:
        input = fetch_input(day, year)

    match args.part:
        case 1:
            func = part1
        case 2:
            func = part2

    start = time.perf_counter_ns()
    answer = func(input)
    duration = time.perf_counter_ns() - start

    if not args.test:
        subprocess.run("pbcopy", input=str(answer).encode())

    print(answer)

    print(f"Ran in: {duration / 1e6}ms")
