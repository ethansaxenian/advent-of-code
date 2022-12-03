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
		set1 := set.NewSetFromIterable(letters[0 : l/2])
		set2 := set.NewSetFromIterable(letters[l/2 : l])

		for _, c := range set.Intersection(set1, set2).Items() {
			total += strings.Index(alphabet, c)
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

		for _, c := range set.Intersection(set1, set2, set3).Items() {
			total += strings.Index(alphabet, c)
		}
	}

	return total
}

func main() {
	util.Run(part1, part2)
}
