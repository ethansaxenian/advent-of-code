package main

import (
	"math"

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

func part2() int {
	target := part1()
	sum := inputLines[0]
	i := 0
	j := 0
	for sum != target {
		if sum < target {
			j++
			sum += inputLines[j]
		} else if sum > target {
			sum -= inputLines[i]
			i++
		}
	}

	max, min := 0, math.MaxInt64
	for i <= j {
		n := inputLines[i]
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		i++
	}
	return min + max
}

func main() {
	util.Run(part1, part2)
}
