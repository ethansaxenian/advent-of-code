package main

import (
	"strings"

	"github.com/ethansaxenian/advent-of-code/2022/util"
)

var inputLines = util.FetchInput(4)

func processAssignments(assignments string) (int, int, int, int) {
	elves := strings.Split(assignments, ",")
	assignment1 := strings.Split(elves[0], "-")
	assignment2 := strings.Split(elves[1], "-")

	return util.ToInt(assignment1[0]), util.ToInt(assignment1[1]), util.ToInt(assignment2[0]), util.ToInt(assignment2[1])
}

func part1() int {
	total := 0
	for _, assignments := range inputLines {
		if assignments == "" {
			continue
		}

		min1, max1, min2, max2 := processAssignments(assignments)

		if min1 <= min2 && max1 >= max2 || min2 <= min1 && max2 >= max1 {
			total++
		}
	}

	return total
}

func part2() int {
	total := 0
	for _, assignments := range inputLines {
		if assignments == "" {
			continue
		}

		min1, max1, min2, max2 := processAssignments(assignments)

		if min2 <= max1 && max2 >= min1 || min1 <= max2 && max1 >= min2 {
			total++
		}
	}

	return total
}

func main() {
	util.Run(part1, part2)
}
