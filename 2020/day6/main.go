package main

import (
	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInput(6)

// var inputLines = []string{
// 	"abc",
// 	"",
// 	"a",
// 	"b",
// 	"c",
// 	"",
// 	"ab",
// 	"ac",
// 	"",
// 	"a",
// 	"a",
// 	"a",
// 	"a",
// 	"",
// 	"b",
// }

func part1() int {
	var count int
	group := map[rune]bool{}
	for _, line := range inputLines {
		if line == "" {
			count += len(group)
			group = map[rune]bool{}
		} else {
			for _, char := range line {
				group[char] = true
			}
		}
	}
	count += len(group)
	return count
}

func part2() int {
	var count int
	group := map[rune]int{}
	groupLen := 0
	for _, line := range inputLines {
		if line == "" {
			for _, v := range group {
				if v == groupLen {
					count++
				}
			}
			group = map[rune]int{}
			groupLen = 0
		} else {
			groupLen++
			for _, char := range line {
				group[char]++
			}
		}
	}
	for _, v := range group {
		if v == groupLen {
			count++
		}
	}
	return count
}

func main() {
	util.Run(part1, part2)
}
