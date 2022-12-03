#!/usr/bin/env zsh

DAY=$1

mkdir "day${DAY}"

echo "package main

import (
	\"github.com/ethansaxenian/advent-of-code-2022/util\"
)

var rucksacks = util.ReadInput(${DAY})

func part1() int {
	return 0
}

func part2() int {
	return 0
}

func main() {
	util.Run(part1, part2)
}" > "day${DAY}/main.go"

source .env
curl "https://adventofcode.com/2022/day/${DAY}/input" -H "cookie: session=${COOKIE}" -o "day${DAY}/input.txt" 2>/dev/null
