package main

import (
	"strings"

	"github.com/ethansaxenian/advent-of-code-2022/set"
	"github.com/ethansaxenian/advent-of-code-2022/util"
)

var inputLines = util.FetchInput(8)

func buildGrid() [][]int {
	grid := [][]int{}
	for _, row := range inputLines {
		intRow := []int{}
		chars := strings.Split(row, "")
		for _, c := range chars {
			intRow = append(intRow, util.ToInt(c))
		}
		grid = append(grid, intRow)

	}
	return grid
}

func goRight(grid [][]int, visible set.Set[[2]int]) {
	for r, row := range grid {
		max := -1
		for c, h := range row {
			if h > max {
				visible.Add([2]int{r, c})
				max = h
			}
		}
	}
}

func goLeft(grid [][]int, visible set.Set[[2]int]) {
	for r, row := range grid {
		max := -1
		for c := len(row) - 1; c >= 0; c-- {
			if row[c] > max {
				visible.Add([2]int{r, c})
				max = row[c]
			}
		}
	}
}

func goDown(grid [][]int, visible set.Set[[2]int]) {
	for c := range grid[0] {
		max := -1
		for r, row := range grid {
			if row[c] > max {
				visible.Add([2]int{r, c})
				max = row[c]
			}
		}
	}
}

func goUp(grid [][]int, visible set.Set[[2]int]) {
	for c := range grid[0] {
		max := -1
		for r := len(grid) - 1; r >= 0; r-- {
			if grid[r][c] > max {
				visible.Add([2]int{r, c})
				max = grid[r][c]
			}
		}
	}
}

func part1() int {
	visible := set.NewEmptySet[[2]int]()
	grid := buildGrid()
	goRight(grid, visible)
	goLeft(grid, visible)
	goDown(grid, visible)
	goUp(grid, visible)
	return len(visible)
}

func calculateScenicScore(grid [][]int, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])
	score := 1
	for _, dir := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		y, x := r, c
		dist := 0
		for {
			if y <= 0 || y >= rows-1 || x <= 0 || x >= cols-1 {
				break
			}
			y += dir[0]
			x += dir[1]
			dist++
			if grid[y][x] >= grid[r][c] {
				break
			}
		}
		score *= dist
	}
	return score
}

func part2() int {
	grid := buildGrid()
	max := 0
	for r, row := range grid {
		for c := range row {
			score := calculateScenicScore(grid, r, c)
			if score > max {
				max = score
			}
		}
	}
	return max
}

func main() {
	util.Run(part1, part2)
}
