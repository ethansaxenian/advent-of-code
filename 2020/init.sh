#!/usr/bin/env zsh

DAY=$1

if [[ -d "day${DAY}" ]]; then
	echo "Day ${DAY} already exists"
	exit 1
fi

mkdir "day${DAY}"

echo "package main

import \"github.com/ethansaxenian/advent-of-code/2020/util\"

var inputLines = util.FetchInput(${DAY})

func part1() int {
	return 0
}

func part2() int {
	return 0
}

func main() {
	util.Run(part1, part2)
}" > "day${DAY}/main.go"
