package main

import (
	"strings"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

var rounds = util.FetchInput(2)

func part1() int {
	m := map[rune]rune{
		'X': 'R',
		'Y': 'P',
		'Z': 'S',
		'A': 'R',
		'B': 'P',
		'C': 'S',
	}

	p := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	total := 0

	for _, r := range rounds {
		if r == "" {
			continue
		}

		moves := []rune(r)
		opp := moves[0]
		me := moves[2]

		total += p[string(me)]

		if diff := m[opp] - m[me]; diff == 0 {
			total += 3
		} else if (diff == 1) || (diff == 2) || (diff == -3) {
			total += 6
		}

	}

	return total
}

func part2() int {
	wins := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	losses := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	p := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	total := 0

	for _, r := range rounds {
		if r == "" {
			continue
		}

		moves := strings.Split(r, " ")
		opp := moves[0]
		me := moves[1]

		if me == "Y" {
			total += 3 + p[opp]
		} else if me == "X" {
			total += p[losses[opp]]
		} else if me == "Z" {
			total += 6 + p[wins[opp]]
		}
	}

	return total
}

func main() {
	util.Run(part1, part2)
}
