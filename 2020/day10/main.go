package main

import (
	"sort"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInputInts(10)

func part1() int {
	sort.Ints(inputLines)
	j := 0
	d1, d3 := 0, 0
	for _, v := range inputLines {
		switch v - j {
		case 1:
			d1++
		case 3:
			d3++
		default:
			panic("what")
		}
		j = v
	}
	return d1 * (d3 + 1)
}

var cache = map[int]int{}

func dfs(i int, adapters []int) int {
	if i == len(adapters)-1 {
		return 1
	}

	if x, ok := cache[i]; ok {
		return x
	}

	n := 0
	for _, j := range [3]int{i + 1, i + 2, i + 3} {
		if j < len(adapters) && adapters[j]-adapters[i] <= 3 {
			n += dfs(j, adapters)
		}
	}
	cache[i] = n
	return n
}

func part2() int {
	sort.Ints(inputLines)
	max := inputLines[len(inputLines)-1] + 3
	inputLines = append([]int{0}, inputLines...)
	inputLines = append(inputLines, max)
	return dfs(0, inputLines)
}

func main() {
	util.Run(part1, part2)
}
