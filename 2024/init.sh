#!/usr/bin/env zsh

DAY=$1

if [[ -a "${DAY}.py" ]]; then
	echo "Day ${DAY} already exists"
	exit 1
fi


echo "import util


def part1(input: str) -> int:
    pass


def part2(input: str) -> int:
    pass


if __name__ == \"__main__\":
    util.run(${DAY}, part1, part2)" > "${DAY}.py"
