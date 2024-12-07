import argparse
import json
import subprocess
import sys
import time
import urllib.request
from collections.abc import Callable
from pathlib import Path, PurePath

CURR_DIR = PurePath(__file__).parent

parser = argparse.ArgumentParser()

parser.add_argument(
    "-p", "--part", type=int, help="which part to run", default=1, choices=[1, 2]
)
parser.add_argument(
    "-t",
    "--test",
    help="use test input",
    action=argparse.BooleanOptionalAction,
)
args = parser.parse_args()


def fetch_input(day: int) -> str:
    with Path(CURR_DIR.parent / "aoc.json").open() as f:
        config = json.load(f)

    input_dir = Path(CURR_DIR / "input")
    input_dir.mkdir(exist_ok=True)

    input_path = Path(input_dir / f"day{day}.txt")
    if input_path.exists():
        return input_path.read_text().strip()

    req = urllib.request.Request(
        f"https://adventofcode.com/2024/day/{day}/input",
    )
    req.add_header("Cookie", f"session={config['cookie']}")
    req.add_header(
        "User-Agent",
        f"github.com/ethansaxenian/advent-of-code/tree/main/2024 by {config['email']}",
    )

    with urllib.request.urlopen(req) as response:
        puzzle_input = response.read().decode("utf-8").strip()

    input_path.write_text(puzzle_input)

    return puzzle_input.strip()


def run(
    day: int,
    part1: Callable[[str], str | int],
    part2: Callable[[str], str | int],
):
    if args.test:
        input = sys.stdin.read()
    else:
        input = fetch_input(day)

    start = time.perf_counter_ns()
    match int(args.part):
        case 1:
            answer = part1(input)
        case 2:
            answer = part2(input)
        case _:
            raise ValueError("invalid part")
    duration = time.perf_counter_ns() - start

    if not args.test:
        subprocess.run("pbcopy", input=str(answer).encode())

    print(answer)

    print(f"Ran in: {duration / 1e6}ms")
