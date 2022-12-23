package main

import (
	"math"

	"github.com/ethansaxenian/advent-of-code/2022/set"
	"github.com/ethansaxenian/advent-of-code/2022/util"
)

type point struct {
	r, c int
}

var inputLines = util.FetchInput(23)

func getElves() set.Set[point] {
	elves := set.NewEmptySet[point]()

	for r, row := range inputLines {
		for c, x := range row {
			if x == '#' {
				elves.Add(point{r, c})
			}
		}
	}
	return elves
}

func emptyNeighbors(p point, dir [2]int, elves set.Set[point]) []point {
	n := []point{}
	var dirs [][2]int
	switch dir {
	case [2]int{1, 1}: // all
		dirs = [][2]int{{-1, 0}, {-1, 1}, {-1, -1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	case [2]int{-1, 0}:
		dirs = [][2]int{{-1, 0}, {-1, 1}, {-1, -1}}
	case [2]int{0, 1}:
		dirs = [][2]int{{0, 1}, {-1, 1}, {1, 1}}
	case [2]int{1, 0}:
		dirs = [][2]int{{1, 0}, {1, 1}, {1, -1}}
	case [2]int{0, -1}:
		dirs = [][2]int{{0, -1}, {-1, -1}, {1, -1}}
	}
	for _, d := range dirs {
		newP := point{p.r + d[0], p.c + d[1]}
		if !elves.Contains(newP) {
			n = append(n, newP)
		}
	}

	return n
}

func propose(elves set.Set[point], directions [][2]int) (map[point]point, map[point]int) {
	proposals := map[point]point{}
	proposalCounts := map[point]int{}
	for _, e := range elves.Items() {
		newElf := e
		if len(emptyNeighbors(e, [2]int{1, 1}, elves)) < 8 {
			for _, dir := range directions {
				if len(emptyNeighbors(e, dir, elves)) == 3 {
					newElf = point{e.r + dir[0], e.c + dir[1]}
					break
				}
			}
		}
		proposals[e] = newElf
		proposalCounts[newElf]++
	}

	return proposals, proposalCounts
}

func move(proposals map[point]point, counts map[point]int) set.Set[point] {
	newElves := set.NewEmptySet[point]()

	for orig, next := range proposals {
		if num := counts[next]; num == 1 {
			newElves.Add(next)
		} else {
			newElves.Add(orig)
		}
	}

	return newElves
}

func round(elves set.Set[point], directions [][2]int) (set.Set[point], [][2]int) {
	proposals, counts := propose(elves, directions)
	newElves := move(proposals, counts)
	newDirectionOrder := append(directions[1:], directions[0])
	return newElves, newDirectionOrder
}

func countEmptyTiles(elves set.Set[point]) int {
	r, c := math.MaxInt, math.MaxInt
	R, C := math.MinInt, math.MinInt

	for _, e := range elves.Items() {
		r = util.Min(r, e.r)
		c = util.Min(c, e.c)
		R = util.Max(R, e.r)
		C = util.Max(C, e.c)
	}

	count := 0

	for y := r; y <= R; y++ {
		for x := c; x <= C; x++ {
			if !elves.Contains(point{y, x}) {
				count++
			}
		}
	}

	return count
}

func part1() int {
	elves := getElves()
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < 10; i++ {
		elves, directions = round(elves, directions)
	}
	return countEmptyTiles(elves)
}

func part2() int {
	elves := getElves()
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	i := 1
	for i > 0 {
		newElves, newDirections := round(elves, directions)

		movement := false
		for _, e := range newElves.Items() {
			if !elves.Contains(e) {
				movement = true
			}
		}

		if !movement {
			return i
		}

		elves = newElves
		directions = newDirections
		i++
	}
	return -1
}

func main() {
	util.Run(part1, part2)
}
