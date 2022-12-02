package main

import (
	"sort"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

func part1() int {
	lines := util.ReadInput(1)
	max := 0
	curr := 0
	for _, line := range lines {
		if line == "" {
			if curr > max {
				max = curr
			}
			curr = 0

		} else {
			curr += util.ToInt(line)
		}
	}
	return max
}

func part2() int {
	lines := util.ReadInput(1)
	calories := make([]int, 0)
	curr := 0
	for _, line := range lines {
		if line == "" {
			calories = append(calories, curr)
			curr = 0
		} else {
			curr += util.ToInt(line)
		}
	}
	sort.Ints(calories)
	return util.Sum(calories[len(calories)-3:])
}

func main() {
	util.Run(part1, part2)
}
