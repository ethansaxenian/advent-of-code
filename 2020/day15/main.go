package main

import (
	"strings"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInput(15)

// var inputLines = []string{"0,3,6"}

type num struct {
	first, last, prev int
}

func solve(turns int) int {

	counter := []num{}
	for i := 0; i <= turns; i++ {
		counter = append(counter, num{-1, -1, -1})
	}

	starting := strings.Split(inputLines[0], ",")

	for i, n := range starting {
		counter[util.ToInt(n)] = num{i + 1, i + 1, i + 1}
	}

	turn := len(starting) + 1
	last := util.ToInt(starting[len(starting)-1])

	for turn <= turns {
		n := counter[last]
		if turn-1 == n.first {
			last = 0
			counter[last] = num{counter[last].first, turn, counter[last].last}
		} else {
			last = n.last - n.prev
			n := counter[last]
			if n.first != -1 {
				counter[last] = num{counter[last].first, turn, counter[last].last}
			} else {
				counter[last] = num{turn, turn, turn}
			}
		}
		turn++
	}
	return last
}

func part1() int {
	return solve(2020)
}

func part2() int {
	return solve(30000000)
}

func main() {
	util.Run(part1, part2)
}
