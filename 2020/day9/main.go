package main

import (
	"errors"
	"sort"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = func() []int {
	out := []int{}
	for _, s := range util.FetchInput(9) {
		out = append(out, util.ToInt(s))
	}
	return out
}()

const preambleLength = 25

func loadPreamble() []int {
	preamble := []int{}
	for i := 0; i < preambleLength; i++ {
		preamble = append(preamble, inputLines[i])
	}
	return preamble
}

func isValid(n int, preamble []int) bool {
	for i := 0; i < preambleLength-1; i++ {
		for j := i + 1; j < preambleLength; j++ {
			if preamble[i]+preamble[j] == n {
				return true
			}
		}
	}
	return false
}

func part1() int {
	preamble := loadPreamble()
	for i := preambleLength; i < len(inputLines); i++ {
		n := inputLines[i]
		if !isValid(n, preamble) {
			return n
		}
		preamble = preamble[1:]
		preamble = append(preamble, n)
	}

	return 0
}

func findSubarray(i, target int) ([]int, error) {
	sum := 0
	for j := i; j < len(inputLines); j++ {
		sum += inputLines[j]
		if sum == target {
			return inputLines[i : j+1], nil
		}
		if sum > target {
			return []int{}, errors.New("no subarray")
		}
	}
	return []int{}, errors.New("no subarray")
}

func part2() int {
	target := part1()
	for i := range inputLines {
		sub, err := findSubarray(i, target)
		if err == nil {
			sort.Ints(sub)
			return sub[0] + sub[len(sub)-1]
		}
	}
	return 0
}

func main() {
	util.Run(part1, part2)
}
