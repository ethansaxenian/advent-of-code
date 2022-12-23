package main

import (
	"container/list"

	"github.com/ethansaxenian/advent-of-code/2022/set"
	"github.com/ethansaxenian/advent-of-code/2022/util"
)

type point struct {
	r int
	c int
}

type qElem struct {
	p     point
	steps int
}

var inputLines = util.FetchInput(12)

func buildGrid() (point, point, [][]rune) {
	var start point
	var end point
	grid := [][]rune{}

	for r, row := range inputLines {
		nextRow := []rune{}
		for c := range row {
			if row[c] == 'S' {
				start = point{r, c}
				nextRow = append(nextRow, 'a')
			} else if row[c] == 'E' {
				end = point{r, c}
				nextRow = append(nextRow, 'z')
			} else {
				nextRow = append(nextRow, rune(row[c]))
			}
		}
		grid = append(grid, nextRow)
	}

	return start, end, grid
}

func bfs(start, end point, grid [][]rune) int {
	q := list.New()
	q.PushBack(qElem{start, 0})
	vis := set.NewEmptySet[point]()
	for q.Len() > 0 {
		front := q.Front()
		q.Remove(front)
		curr := front.Value.(qElem)
		p := curr.p
		steps := curr.steps

		if p == end {
			return steps
		}

		for _, dir := range [4]point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			newP := point{p.r + dir.r, p.c + dir.c}
			if newP.r >= 0 && newP.c >= 0 && newP.r < len(grid) && newP.c < len(grid[0]) && grid[newP.r][newP.c] <= grid[p.r][p.c]+1 {
				if !vis.Contains(newP) {
					vis.Add(newP)
					q.PushBack(qElem{newP, steps + 1})
				}
			}
		}
	}
	return 0
}

func part1() int {
	start, end, grid := buildGrid()
	steps := bfs(start, end, grid)
	return steps
}

func part2() int {
	starts := []point{}
	_, end, grid := buildGrid()
	for r, row := range grid {
		for c := range row {
			if row[c] == 'a' {
				starts = append(starts, point{r, c})
			}
		}
	}
	minSteps := len(grid) * len(grid[0])
	for _, s := range starts {
		if steps := bfs(s, end, grid); steps != 0 && steps < minSteps {
			minSteps = steps
		}
	}
	return minSteps
}

func main() {
	util.Run(part1, part2)
}
