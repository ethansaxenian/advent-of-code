package main

import (
	"fmt"

	"github.com/ethansaxenian/advent-of-code/2023/util"
)

var inputLines = util.FetchInput(11)

// var inputLines = []string{
// 	"...#......",
// 	".......#..",
// 	"#.........",
// 	"..........",
// 	"......#...",
// 	".#........",
// 	".........#",
// 	"..........",
// 	".......#..",
// 	"#...#.....",
// }

func findGalaxies(grid []string, expansion int) [][2]int {
	emptyRows := []int{}
	for r, row := range inputLines {
		hasGalaxy := false
		for _, char := range row {
			if char == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			emptyRows = append(emptyRows, r)
		}
	}

	emptyColumns := []int{}
	for c := range inputLines[0] {
		hasGalaxy := false
		for _, row := range inputLines {
			if row[c] == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			emptyColumns = append(emptyColumns, c)
		}
	}

	galaxies := [][2]int{}
	rowAdd := 0
	for r, row := range grid {
		if util.Contains[int](emptyRows, r) {
			rowAdd += expansion - 1
		}
		colAdd := 0
		for c, char := range row {
			if util.Contains[int](emptyColumns, c) {
				colAdd += expansion - 1
			}
			if char == '#' {
				galaxies = append(galaxies, [2]int{r + rowAdd, c + colAdd})
			}
		}
	}

	return galaxies
}

func day1() {
	galaxies := findGalaxies(inputLines, 2)
	combs := util.Combinations[[2]int](galaxies, 2)
	sum := 0
	for _, comb := range combs {
		sum += util.ShortestPath(comb[0], comb[1])
	}
	fmt.Println(sum)
}

func day2() {
	galaxies := findGalaxies(inputLines, 1_000_000)
	combs := util.Combinations[[2]int](galaxies, 2)
	sum := 0
	for _, comb := range combs {
		sum += util.ShortestPath(comb[0], comb[1])
	}
	fmt.Println(sum)
}

func main() {
	day1()
	day2()
}
