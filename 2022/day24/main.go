package main

import (
	"math"

	"github.com/ethansaxenian/advent-of-code/2022/set"
	"github.com/ethansaxenian/advent-of-code/2022/util"
)

type point struct {
	r, c int
}

type blizzard struct {
	point
	dir string
}

var inputLines = util.FetchInput(24)

var rows = len(inputLines)
var cols = len(inputLines[0])

func buildGrid() ([]blizzard, set.Set[point]) {
	walls := set.NewEmptySet[point]()
	blizzards := []blizzard{}
	for r, row := range inputLines {
		for c, char := range row {
			if char == '>' || char == '<' || char == 'v' || char == '^' {
				blizzards = append(blizzards, blizzard{point{r, c}, string(char)})
			} else if char == '#' {
				walls.Add(point{r, c})
			}
		}
	}
	return blizzards, walls
}

func getNextBlizzards(blizzards []blizzard) ([]blizzard, set.Set[point]) {
	newBlizzards := []blizzard{}
	blizzardLocs := set.NewEmptySet[point]()
	for _, blizz := range blizzards {
		p := getNextBlizzPos(blizz)
		blizzardLocs.Add(p)
		newBlizzards = append(newBlizzards, blizzard{p, blizz.dir})
	}
	return newBlizzards, blizzardLocs
}

func getNextBlizzPos(blizzard blizzard) point {
	r := blizzard.r
	c := blizzard.c
	switch blizzard.dir {
	case ">":
		if c++; c == cols-1 {
			c = 1
		}
	case "<":
		if c--; c == 0 {
			c = cols - 2
		}
	case "v":
		if r++; r == rows-1 {
			r = 1
		}
	case "^":
		if r--; r == 0 {
			r = rows - 2
		}
	}

	return point{r, c}
}

type state struct {
	pos       point
	time      int
	blizzards []blizzard
}

func bfs(start, end point, initTime int, initBlizzards []blizzard, walls set.Set[point]) (int, []blizzard) {
	vis := set.NewEmptySet[[3]int]()
	vis.Add([3]int{start.r, start.c, initTime})

	q := []state{{start, initTime, initBlizzards}}

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		if s.pos == end {
			return s.time, s.blizzards
		}

		newBlizzards, occupied := getNextBlizzards(s.blizzards)

		for _, dir := range [][2]int{{0, 0}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			nextPos := point{s.pos.r + dir[0], s.pos.c + dir[1]}
			nextVis := [3]int{nextPos.r, nextPos.c, s.time + 1}
			if nextPos.r >= 0 && nextPos.r < rows && !walls.Contains(nextPos) && !occupied.Contains(nextPos) && !vis.Contains(nextVis) {
				vis.Add(nextVis)
				q = append(q, state{nextPos, s.time + 1, newBlizzards})
			}
		}

	}
	return math.MaxInt, []blizzard{}
}

func part1() int {
	blizzards, walls := buildGrid()
	start := point{0, 1}
	end := point{rows - 1, cols - 2}
	time, _ := bfs(start, end, 0, blizzards, walls)

	return time
}

func part2() int {
	blizzards, walls := buildGrid()
	start := point{0, 1}
	end := point{rows - 1, cols - 2}

	t1, nextBlizzards := bfs(start, end, 0, blizzards, walls)
	t2, nextBlizzards := bfs(end, start, t1, nextBlizzards, walls)
	t3, _ := bfs(start, end, t2, nextBlizzards, walls)
	return t3
}

func main() {
	util.Run(part1, part2)
}
