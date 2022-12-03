package main

import (
	"strings"

	"github.com/ethansaxenian/advent-of-code-2022/set"
	"github.com/ethansaxenian/advent-of-code-2022/util"
)

var rucksacks = util.ReadInput(3)

var alphabet = "-abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func part1() int {
	total := 0

	for _, sack := range rucksacks {
		letters := strings.Split(sack, "")
		l := len(letters)
		l1 := letters[0 : l/2]
		l2 := letters[l/2 : l]

		set1 := set.NewSetFromIterable(l1)
		set2 := set.NewSetFromIterable(l2)

		for _, c := range set1.Items() {
			if set2.Contains(c) {
				total += strings.Index(alphabet, c)
			}
		}

	}

	return total
}

func part2() int {
	total := 0

	for i := 0; i < len(rucksacks)-1; i += 3 {
		set1 := set.NewSetFromString(rucksacks[i])
		set2 := set.NewSetFromString(rucksacks[i+1])
		set3 := set.NewSetFromString(rucksacks[i+2])

		for _, c := range set1.Items() {
			if set2.Contains(c) && set3.Contains(c) {
				total += strings.Index(alphabet, c)
			}
		}
	}

	return total
}

func main() {
	util.Run(part1, part2)
}
