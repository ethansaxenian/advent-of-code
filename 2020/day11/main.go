package main

import (
	"fmt"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

type cell struct {
	state     rune
	neighbors [][2]int
}

func (c cell) countOccupiedNeighbors(grid [][]cell) int {
	count := 0
	for _, n := range c.neighbors {
		if grid[n[0]][n[1]].state == '#' {
			count++
		}
	}
	return count
}

func (c cell) String() string {
	return string(c.state)
}

var inputLines = util.FetchInput(11)

var inputLines = []string{
	"L.LL.LL.LL",
	"LLLLLLL.LL",
	"L.L.L..L..",
	"LLLL.LL.LL",
	"L.LL.LL.LL",
	"L.LLLLL.LL",
	"..L.L.....",
	"LLLLLLLLLL",
	"L.LLLLLL.L",
	"L.LLLLL.LL",
}

func occupiedNeighbors(i, j int, grid []string) int {
	num := 0
	for _, di := range []int{-1, 0, 1} {
		for _, dj := range []int{-1, 0, 1} {
			if di == 0 && dj == 0 {
				continue
			}
			xi, xj := i+di, j+dj
			if xi >= 0 && xi < len(grid) && xj >= 0 && xj < len(grid[i]) {
				if i == 0 && j == 3 {
				}
				if grid[xi][xj] == '#' {
					num++
				}
			}
		}
	}
	return num
}

func nextIter(grid []string) []string {
	next := []string{}
	for i := range grid {
		row := ""
		for j := range grid[i] {
			numOccupiedNeighbors := occupiedNeighbors(i, j, grid)
			char := grid[i][j]
			if char == 'L' && numOccupiedNeighbors == 0 {
				char = '#'
			}
			if char == '#' && numOccupiedNeighbors >= 4 {
				char = 'L'
			}
			row += string(char)
		}
		next = append(next, row)
	}

	return next
}

func notEq(g1, g2 []string) bool {
	for i := 0; i < len(g1); i++ {
		if g1[i] != g2[i] {
			return true
		}
	}
	return false
}

func countOccupied(grid []string) int {
	num := 0
	for _, row := range grid {
		for _, c := range row {
			if c == '#' {
				num++
			}
		}
	}
	return num
}

func printGrid(grid []string) {
	return
	fmt.Println()
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}

func part1() int {
	grid := inputLines
	printGrid(grid)
	next := nextIter(grid)
	for notEq(grid, next) {
		printGrid(next)
		grid = next
		next = nextIter(grid)
	}
	printGrid(next)
	return countOccupied(next)
}

func getNeighbors(i, j int, grid []string) [][2]int {
	neighborhood := [][2]int{}
	for _, di := range []int{-1, 0, 1} {
		for _, dj := range []int{-1, 0, 1} {
			if di == 0 && dj == 0 {
				continue
			}
			xi, xj := i, j
			for {
				xi += di
				xj += dj
				if xi < 0 || xi >= len(grid) || xj < 0 || xj >= len(grid[0]) {
					break
				}
				if grid[xi][xj] != '.' {
					neighborhood = append(neighborhood, [2]int{xi, xj})
					break
				}
			}
		}
	}
	return neighborhood
}

func buildGrid(grid []string) [][]cell {
	g := [][]cell{}
	for i, row := range grid {
		r := []cell{}
		for j, c := range row {
			n := getNeighbors(i, j, grid)
			r = append(r, cell{c, n})
		}
		g = append(g, r)
	}
	return g
}

func nextIter2(grid [][]cell) [][]cell {
	next := [][]cell{}
	for _, row := range grid {
		newRow := []cell{}
		for _, c := range row {
			numOccupiedNeighbors := c.countOccupiedNeighbors(grid)
			char := c.state
			if char == 'L' && numOccupiedNeighbors == 0 {
				char = '#'
			}
			if char == '#' && numOccupiedNeighbors >= 5 {
				char = 'L'
			}
			newRow = append(newRow, cell{char, c.neighbors})
		}
		next = append(next, newRow)
	}
	return next
}

func notEq2(g1, g2 [][]cell) bool {
	for i := 0; i < len(g1); i++ {
		for j := 0; j < len(g1[0]); j++ {
			if g1[i][j].state != g2[i][j].state {
				return true
			}
		}
	}
	return false
}

func countOccupied2(grid [][]cell) int {
	num := 0
	for _, row := range grid {
		for _, c := range row {
			if c.state == '#' {
				num++
			}
		}
	}
	return num
}

func printGrid2(grid [][]cell) {
	fmt.Println()
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(string(c.state))
		}
		fmt.Println()
	}
	fmt.Println()
}

func part2() int {
	grid := buildGrid(inputLines)
	printGrid2(grid)
	next := nextIter2(grid)
	for notEq2(grid, next) {
		printGrid2(next)
		grid = next
		next = nextIter2(grid)
	}
	printGrid2(next)
	return countOccupied2(next)
}

func main() {
	util.Run(part1, part2)
}
