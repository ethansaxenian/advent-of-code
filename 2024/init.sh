#!/usr/bin/env zsh

DAY=$1

if [[ -d "day${DAY}" ]]; then
	echo "Day ${DAY} already exists"
	exit 1
fi

mkdir "day${DAY}"

echo "package main

import (
	\"github.com/ethansaxenian/advent-of-code/2024/util\"
)

func part1(input []string) int {
	return 0
}

func part2(input []string) int {
	return 0
}

func main() {
	util.Run(${DAY}, part1, part2)
}" > "day${DAY}/main.go"
