package main

import (
	"fmt"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2022/util"
)

var inputLines = util.FetchInput(10)

func addInterestingSignalStrength(cycle, register int) int {
	if util.Contains([]int{20, 60, 100, 140, 180, 220}, cycle) {
		return cycle * register
	}
	return 0
}

func part1() int {
	register, cycle := 1, 1
	total := 0
	i := 0
	for i < len(inputLines) {
		cmd := strings.Split(inputLines[i], " ")
		if cmd[0] == "noop" {
			cycle++
			total += addInterestingSignalStrength(cycle, register)
		} else {
			cycle++
			total += addInterestingSignalStrength(cycle, register)
			cycle++
			register += util.ToInt(cmd[1])
			total += addInterestingSignalStrength(cycle, register)
		}
		i++
	}

	return total
}

func draw(cycle int, sprite [3]int) {
	pos := cycle % 40
	if pos == 0 {
		pos = 40
	}

	if util.Contains(sprite[:], pos) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if pos == 40 {
		fmt.Println()
	}
}

func part2() int {
	cycle := 1
	i := 0
	sprite := [3]int{1, 2, 3}
	for i < len(inputLines) {
		cmd := strings.Split(inputLines[i], " ")
		draw(cycle, sprite)
		if cmd[0] == "noop" {
			cycle++
		} else {
			cycle++
			draw(cycle, sprite)
			cycle++
			x := util.ToInt(cmd[1])
			sprite[0] += x
			sprite[1] += x
			sprite[2] += x
		}
		i++
	}

	return 0
}

func main() {
	util.Run(part1, part2)
}
