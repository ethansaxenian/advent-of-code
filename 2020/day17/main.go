package main

import (
	"fmt"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInput(17)

// var inputLines = []string{
// 	".#.",
// 	"..#",
// 	"###",
// }

type grid map[[3]int]int

func (g grid) print() {
	for k, v := range g {
		fmt.Println(k, v)
	}
}

func (g grid) activeNeighbors(x, y, z int) int {
	active := 0
	for _, dx := range [3]int{-1, 0, 1} {
		for _, dy := range [3]int{-1, 0, 1} {
			for _, dz := range [3]int{-1, 0, 1} {
				if dx != 0 || dy != 0 || dz != 0 {
					active += g[[3]int{x + dx, y + dy, z + dz}]
				}
			}
		}
	}

	return active
}

func parseInput(input []string) grid {
	g := grid{}
	for y, row := range input {
		for x, c := range row {
			point := [3]int{x, y, 0}
			if c == '#' {
				g[point] = 1
			}
		}
	}

	return g
}

func part1() int {
	g := parseInput(inputLines)
	x := [2]int{0, len(inputLines[0]) - 1}
	y := [2]int{0, len(inputLines) - 1}
	z := [2]int{0, 0}

	for i := 0; i < 6; i++ {
		x[0]--
		x[1]++
		y[0]--
		y[1]++
		z[0]--
		z[1]++

		newG := grid{}

		for xi := x[0]; xi <= x[1]; xi++ {
			for yi := y[0]; yi <= y[1]; yi++ {
				for zi := z[0]; zi <= z[1]; zi++ {
					p := [3]int{xi, yi, zi}
					n := g.activeNeighbors(xi, yi, zi)
					if g[p] == 1 && (n == 2 || n == 3) {
						newG[p] = 1
					} else if g[p] == 0 && n == 3 {
						newG[p] = 1
					}
				}
			}
		}

		g = newG
	}

	numActive := 0
	for _, v := range g {
		numActive += v
	}
	return numActive
}

type grid4 map[[4]int]int

func (g grid4) activeNeighbors(w, x, y, z int) int {
	active := 0
	for _, dw := range [3]int{-1, 0, 1} {
		for _, dx := range [3]int{-1, 0, 1} {
			for _, dy := range [3]int{-1, 0, 1} {
				for _, dz := range [3]int{-1, 0, 1} {
					if dw != 0 || dx != 0 || dy != 0 || dz != 0 {
						active += g[[4]int{w + dw, x + dx, y + dy, z + dz}]
					}
				}
			}
		}
	}

	return active
}

func parseInput4(input []string) grid4 {
	g := grid4{}
	for y, row := range input {
		for x, c := range row {
			point := [4]int{0, x, y, 0}
			if c == '#' {
				g[point] = 1
			}
		}
	}

	return g
}

func part2() int {
	g := parseInput4(inputLines)
	x := [2]int{0, len(inputLines[0]) - 1}
	y := [2]int{0, len(inputLines) - 1}
	z := [2]int{0, 0}
	w := [2]int{0, 0}

	for i := 0; i < 6; i++ {
		x[0]--
		x[1]++
		y[0]--
		y[1]++
		z[0]--
		z[1]++
		w[0]--
		w[1]++

		newG := grid4{}

		for wi := w[0]; wi <= w[1]; wi++ {
			for xi := x[0]; xi <= x[1]; xi++ {
				for yi := y[0]; yi <= y[1]; yi++ {
					for zi := z[0]; zi <= z[1]; zi++ {
						p := [4]int{wi, xi, yi, zi}
						n := g.activeNeighbors(wi, xi, yi, zi)
						if g[p] == 1 && (n == 2 || n == 3) {
							newG[p] = 1
						} else if g[p] == 0 && n == 3 {
							newG[p] = 1
						}
					}
				}
			}
		}

		g = newG
	}

	numActive := 0
	for _, v := range g {
		numActive += v
	}
	return numActive
}

func main() {
	util.Run(part1, part2)
}
