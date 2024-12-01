import argparse
import os
import subprocess
import sys
from collections.abc import Callable
from io import StringIO
from pathlib import Path, PurePath
from urllib.request import Request, urlopen

from dotenv import load_dotenv

CURR_DIR = PurePath(__file__).parent

load_dotenv(PurePath(CURR_DIR.parent / ".env"))

COOKIE = os.environ["AOC_COOKIE"]
EMAIL = os.environ["AOC_EMAIL"]

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


def fetch_input(day: int) -> list[str]:
    input_dir = Path(CURR_DIR / "input")
    input_dir.mkdir(exist_ok=True)

    input_path = Path(input_dir / f"day{day}.txt")
    if input_path.exists():
        return input_path.read_text().strip().splitlines()

    req = Request(f"https://adventofcode.com/2024/day/{day}/input")
    req.add_header("Cookie", f"session={COOKIE}")
    req.add_header("User-Agent", f"github.com/ethansaxenian/advent-of-code by {EMAIL}")

    with urlopen(req) as response:
        puzzle_input = response.read().decode("utf-8").strip()

    input_path.write_text(puzzle_input)

    return puzzle_input.splitlines()


def fetch_input_from_stdin() -> list[str]:
    return sys.stdin.read().splitlines()


def run(
    day: int,
    part1: Callable[[list[str]], str | int],
    part2: Callable[[list[str]], str | int],
):
    if args.test:
        input = sys.stdin.read().splitlines()
    else:
        input = fetch_input(day)

    match int(args.part):
        case 1:
            ans = part1(input)
        case 2:
            ans = part2(input)
        case _:
            raise ValueError("invalid part")

    if not args.test:
        subprocess.run("pbcopy", input=str(ans).encode())

    print(ans)
