package main

import (
	"aoc2020/util"
	"fmt"
	"sort"
)

var _ = fmt.Printf

var inputLines = util.FetchInput(5)

func findRow(boardingPass string) int {
	min, max := 0, 127
	for _, char := range boardingPass {
		if char == 'F' {
			max = (min + max) / 2
		} else if char == 'B' {
			min = (min + max) / 2
		}
	}

	return max
}

func findCol(boardingPass string) int {
	min, max := 0, 7
	for _, char := range boardingPass {
		if char == 'L' {
			max = (min + max) / 2
		} else if char == 'R' {
			min = (min + max) / 2
		}
	}

	return max
}

func part1() int {
	maxSeatID := 0
	for _, boardingPass := range inputLines {
		seatID := findRow(boardingPass)*8 + findCol(boardingPass)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

func part2() int {
	seatIDs := []int{}
	for _, boardingPass := range inputLines {
		seatID := findRow(boardingPass)*8 + findCol(boardingPass)
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)
	x := seatIDs[0]
	for _, y := range seatIDs {
		if x != y {
			return x
		}
		x++
	}
	return 0

}

func main() {
	util.Run(part1, part2)
}
