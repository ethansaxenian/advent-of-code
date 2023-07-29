package main

import (
	"errors"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInput(8)

func part1() int {
	seen := map[int]bool{}
	acc := 0
	i := 0
	for i < len(inputLines) {
		if _, ok := seen[i]; ok {
			return acc
		}
		seen[i] = true
		instruction := strings.Split(inputLines[i], " ")
		switch instruction[0] {
		case "acc":
			acc += util.ToInt(instruction[1])
			i++
		case "jmp":
			i += util.ToInt(instruction[1])
		case "nop":
			i++
		}
	}
	return 0
}

func possibleChanges() []int {
	changes := []int{}
	for i := range inputLines {
		instruction := strings.Split(inputLines[i], " ")
		if instruction[0] == "nop" || instruction[0] == "jmp" {
			changes = append(changes, i)
		}
	}
	return changes
}

func runWithChange(changeIdx int) (int, error) {
	seen := map[int]bool{}
	acc := 0
	i := 0
	for i < len(inputLines) {
		if _, ok := seen[i]; ok {
			return 0, errors.New("infinite loop")
		}
		seen[i] = true
		instruction := strings.Split(inputLines[i], " ")
		op := instruction[0]
		if i == changeIdx {
			if op == "nop" {
				op = "jmp"
			} else if op == "jmp" {
				op = "nop"
			}
		}
		switch op {
		case "acc":
			acc += util.ToInt(instruction[1])
			i++
		case "jmp":
			i += util.ToInt(instruction[1])
		case "nop":
			i++
		}
	}
	return acc, nil
}

func part2() int {
	for _, idx := range possibleChanges() {
		acc, err := runWithChange(idx)
		if err == nil {
			return acc
		}
	}
	return 0
}

func main() {
	util.Run(part1, part2)
}
