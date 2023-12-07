package main

import (
	"fmt"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2023/util"
)

func part1(lines []string) {
	t := util.StrListToInts(strings.Fields(lines[0])[1:])
	d := util.StrListToInts(strings.Fields(lines[1])[1:])

	prod := 1

	for i := 0; i < len(d); i++ {
		record := d[i]
		time := t[i]
		numWays := 0
		for speed := 1; speed < time; speed++ {
			dist := speed * (time - speed)
			if dist > record {
				numWays++
			}
		}
		prod *= numWays
	}

	fmt.Println(prod)
}

func part2(lines []string) {
	time := util.StrToInt(strings.Join(strings.Fields(lines[0])[1:], ""))
	record := util.StrToInt(strings.Join(strings.Fields(lines[1])[1:], ""))

	numWays := 0
	for speed := 1; speed < time; speed++ {
		dist := speed * (time - speed)
		if dist > record {
			numWays++
		}
	}
	fmt.Println(numWays)

}

func main() {
	lines := util.FetchInput(6)
	part1(lines)
	part2(lines)
}
