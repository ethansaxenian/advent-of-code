package main

import (
	"fmt"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInput(2)

func parsePolicy(policy string) (int, int, rune, string) {
	parts := strings.Split(policy, " ")
	bounds, c, pw := parts[0], parts[1], parts[2]
	boundParts := strings.Split(bounds, "-")
	lb, ub := boundParts[0], boundParts[1]

	return util.ToInt(lb), util.ToInt(ub), rune(c[0]), pw
}

func counter(pw string) map[rune]int {
	m := map[rune]int{}

	for _, c := range pw {
		m[c]++
	}
	return m
}

func part1() int {
	numValid := 0

	for _, line := range inputLines {
		lb, ub, c, pw := parsePolicy(line)
		counts := counter(pw)

		if counts[c] >= lb && counts[c] <= ub {
			numValid++
		}
	}

	return numValid
}

func part2() int {
	numValid := 0

	for _, line := range inputLines {
		firstPos, secondPos, c, pw := parsePolicy(line)
		fmt.Println(firstPos, secondPos, c, pw)
		fmt.Println(pw[firstPos-1], pw[secondPos-1])
		firstChar := pw[firstPos-1]
		secondChar := pw[secondPos-1]
		if (firstChar == byte(c) || secondChar == byte(c)) && (firstChar != secondChar) {
			numValid++
		}
	}

	return numValid
}

func main() {
	util.Run(part1, part2)
}
