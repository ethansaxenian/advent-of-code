package main

import (
	"aoc2020/util"
)

var inputLines = util.FetchInput(1)

// var inputLines = []string{
// 	"1721",
// 	"979",
// 	"366",
// 	"299",
// 	"675",
// 	"1456",
// }

var sum = 2020

func toInts(slice []string) []int {
	ints := make([]int, len(slice))
	for i, s := range slice {
		if s != "" {
			ints[i] = util.ToInt(s)
		}
	}
	return ints
}

func part1() int {
	seen := map[int]int{}

	for _, x := range toInts(inputLines) {
		diff := sum - x

		_, check := seen[diff]
		if check {
			return x * diff
		}

		seen[x] = x
	}

	return 0
}

func part2() int {
	seen := map[int]int{}
	intLines := toInts(inputLines)

	for i := 0; i < len(intLines)-1; i++ {
		for j := i + 1; j < len(intLines); j++ {
			x := intLines[i]
			y := intLines[j]

			if x == y {
				panic("x == y")
			}

			diff := sum - (x + y)

			_, check := seen[diff]
			if check {
				return x * y * diff
			}

			seen[x] = x
			seen[y] = y
		}
	}

	return 0
}

func main() {
	util.Run(part1, part2)
}
