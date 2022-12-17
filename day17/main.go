package main

import (
	"fmt"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

type point struct {
	c, r int
}

var inputLines = util.FetchInput(17)

var jetPattern = inputLines[0]

// var jetPattern = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

var rocks = [][]point{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}},
	{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
}

func buildChamber(h int) [][7]bool {
	chamber := [][7]bool{}
	for i := 0; i < h*3; i++ {
		chamber = append(chamber, [7]bool{})
	}
	for i := 0; i < 7; i++ {
		chamber[0][i] = true
	}
	return chamber[:]
}

func getNewRock(y int, rock []point) []point {
	newRock := []point{}
	for _, r := range rock {
		newRock = append(newRock, point{r.c + 2, r.r + y})
	}
	return newRock
}

func push(rock []point, j int, chamber [][7]bool) []point {
	newRock := []point{}
	if jetPattern[j] == '>' {
		for _, r := range rock {
			p := point{r.c + 1, r.r}
			if p.c >= 7 || chamber[p.r][p.c] {
				return rock
			}
			newRock = append(newRock, p)
		}
	} else if jetPattern[j] == '<' {
		for _, r := range rock {
			p := point{r.c - 1, r.r}
			if p.c < 0 || chamber[p.r][p.c] {
				return rock
			}
			newRock = append(newRock, p)
		}
	}
	return newRock
}

func fall(rock []point, chamber [][7]bool) ([]point, bool) {
	newRock := []point{}
	for _, r := range rock {
		p := point{r.c, r.r - 1}
		if chamber[p.r][p.c] {
			return rock, true
		}
		newRock = append(newRock, p)
	}
	return newRock, false
}

func getHighestRock(rock []point) int {
	h := 0
	for _, p := range rock {
		h = util.Max(h, p.r)
	}
	return h
}

func insertRock(rock []point, chamber [][7]bool) {
	for _, p := range rock {
		chamber[p.r][p.c] = true
	}
}

func printChamber(chamber [][7]bool, t int) {
	for i := t; i >= 0; i-- {
		for _, r := range chamber[i] {
			if r {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part1() int {
	totalNumRocks := 2022
	chamber := buildChamber(totalNumRocks)
	tallestRock := 0
	numRocks := 0
	j := 0
	for numRocks < totalNumRocks {
		for _, r := range rocks {
			rock := getNewRock(tallestRock+4, r)
			for {
				rock = push(rock, j, chamber)
				j = (j + 1) % len(jetPattern)
				done := false
				if rock, done = fall(rock, chamber); done {
					tallestRock = util.Max(tallestRock, getHighestRock(rock))
					insertRock(rock, chamber)
					break
				}
			}
			numRocks++
			if numRocks >= totalNumRocks {
				break
			}
		}
	}
	return tallestRock
}

func part2() int {
	return 0
}

func main() {
	util.Run(part1, part2)
}
