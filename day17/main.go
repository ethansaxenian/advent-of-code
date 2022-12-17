package main

import (
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

func buildChamber() [][7]bool {
	chamber := [][7]bool{}
	for i := 0; i < 5; i++ {
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

func calculateHash(jetIndex, rockIndex, tallestRock int, chamber [][7]bool) [9]int {
	hash := [9]int{}
	for i := 0; i < 7; i++ {
		j := 0
		for !chamber[tallestRock-j][i] {
			j += 1
		}
		hash[i] = j
	}
	hash[7] = jetIndex
	hash[8] = rockIndex
	return hash
}

var states = map[[9]int][][2]int{}

func tetris(totalNumRocks int) int {
	chamber := buildChamber()
	tallestRock := 0
	numRocks := 0
	j := 0
	for numRocks < totalNumRocks {
		for i, r := range rocks {
			rock := getNewRock(tallestRock+4, r)
			hash := calculateHash(j, i, tallestRock, chamber)
			states[hash] = append(states[hash], [2]int{numRocks, tallestRock})
			for {
				rock = push(rock, j, chamber)
				j = (j + 1) % len(jetPattern)
				done := false
				if rock, done = fall(rock, chamber); done {
					tallestRock = util.Max(tallestRock, getHighestRock(rock))
					insertRock(rock, chamber)
					for len(chamber) < tallestRock+8 {
						chamber = append(chamber, [7]bool{})
					}
					break
				}
			}
			numRocks++
			if numRocks >= totalNumRocks {
				return tallestRock
			}
		}
	}
	return 0
}

func part1() int {
	return tetris(2022)
}

func part2() int {
	totalNumRocks := 1000000000000
	tetris(len(jetPattern))

	var cycles [][2]int
	for _, v := range states {
		if len(v) > 1 {
			cycles = v
			break
		}
	}
	numRocksBeforeFirstCycle := cycles[0][0]
	rocksPerCycle := cycles[1][0] - cycles[0][0]
	tallestRockPerCycle := cycles[1][1] - cycles[0][1]

	numCycles := (totalNumRocks - numRocksBeforeFirstCycle) / rocksPerCycle
	remainingRocks := (totalNumRocks - numRocksBeforeFirstCycle) % rocksPerCycle

	return numCycles*tallestRockPerCycle + (tetris(numRocksBeforeFirstCycle + remainingRocks))
}

func main() {
	util.Run(part1, part2)
}
