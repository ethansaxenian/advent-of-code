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
var lcm = util.LCM(rows-2, cols-2)

func getBlizzardsAndWalls() ([]blizzard, set.Set[point]) {
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

func preprocessBlizzards(blizzards []blizzard) map[int]set.Set[point] {
	blizzardMap := map[int]set.Set[point]{}

	for i := 0; i < lcm; i++ {
		newBlizzards, occupied := getNextBlizzards(blizzards)
		blizzards = newBlizzards
		blizzardMap[i] = occupied
	}

	return blizzardMap
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
	pos  point
	time int
}

func bfs(start, end point, initTime int, blizzardMap map[int]set.Set[point], walls set.Set[point]) int {
	vis := set.NewEmptySet[[3]int]()
	vis.Add([3]int{start.r, start.c, initTime})

	q := []state{{start, initTime}}

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		if s.pos == end {
			return s.time
		}

		occupied := blizzardMap[s.time%lcm]

		for _, dir := range [][2]int{{0, 0}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			nextPos := point{s.pos.r + dir[0], s.pos.c + dir[1]}
			nextVis := [3]int{nextPos.r, nextPos.c, s.time + 1}
			if nextPos.r >= 0 && nextPos.r < rows && !walls.Contains(nextPos) && !occupied.Contains(nextPos) && !vis.Contains(nextVis) {
				vis.Add(nextVis)
				q = append(q, state{nextPos, s.time + 1})
			}
		}

	}
	return math.MaxInt
}

func part1() int {
	blizzards, walls := getBlizzardsAndWalls()
	blizzardMap := preprocessBlizzards(blizzards)

	start := point{0, 1}
	end := point{rows - 1, cols - 2}

	return bfs(start, end, 0, blizzardMap, walls)
}

func part2() int {
	blizzards, walls := getBlizzardsAndWalls()
	blizzardMap := preprocessBlizzards(blizzards)

	start := point{0, 1}
	end := point{rows - 1, cols - 2}

	return bfs(
		start,
		end,
		bfs(
			end,
			start,
			bfs(
				start,
				end,
				0,
				blizzardMap,
				walls,
			),
			blizzardMap,
			walls,
		),
		blizzardMap,
		walls,
	)
}

func main() {
	util.Run(part1, part2)
}
