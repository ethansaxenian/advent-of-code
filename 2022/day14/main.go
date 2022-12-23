package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2022/util"
)

type point struct {
	r, c int
}

var inputLines = util.FetchInput(14)

func buildCave(addFloor bool) ([][]bool, int, int, int) {
	cMax := math.MinInt
	cMin := math.MaxInt
	var rMax int = math.MinInt
	paths := [][]point{}
	for _, line := range inputLines {
		parts := strings.Split(line, " -> ")
		rockPathEndpoints := []point{}
		for _, p := range parts {
			cr := strings.Split(p, ",")
			rock := point{util.ToInt(cr[1]), util.ToInt(cr[0])}
			cMax = util.Max(cMax, rock.c)
			cMin = util.Min(cMin, rock.c)
			rMax = util.Max(rMax, rock.r)
			rockPathEndpoints = append(rockPathEndpoints, rock)
		}
		paths = append(paths, rockPathEndpoints)
	}

	var caveWidth int
	if addFloor {
		caveWidth = 700
	} else {
		caveWidth = cMax + 1
	}

	cave := [][]bool{}
	for i := 0; i <= rMax+1; i++ {
		row := []bool{}
		for j := 0; j <= caveWidth; j++ {
			row = append(row, true)
		}
		cave = append(cave, row)
	}

	if addFloor {
		row := []bool{}
		for j := 0; j <= caveWidth; j++ {
			row = append(row, false)
		}
		cave = append(cave, row)
	}

	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			prev := path[i-1]
			curr := path[i]
			var a int
			var b int
			if prev.r == curr.r {
				if prev.c > curr.c {
					a = curr.c
					b = prev.c
				} else {
					a = prev.c
					b = curr.c
				}
				for c := a; c <= b; c++ {
					cave[prev.r][c] = false
				}
			}
			if prev.c == curr.c {
				if prev.r > curr.r {
					a = curr.r
					b = prev.r
				} else {
					a = prev.r
					b = curr.r
				}
				for r := a; r <= b; r++ {
					cave[r][prev.c] = false
				}
			}
		}
	}

	return cave, rMax, cMax, cMin
}

func printCave(min, max int, cave [][]bool) {
	for r, row := range cave {
		fmt.Print(r, " ")
		for i := min; i <= max; i++ {
			if row[i] {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func getPoint(p point, g [][]bool) bool {
	return g[p.r][p.c]
}

func fall(sand point, cave [][]bool) (point, bool) {
	down := point{sand.r + 1, sand.c}
	if getPoint(down, cave) {
		return down, false
	} else {
		left := point{sand.r + 1, sand.c - 1}
		if getPoint(left, cave) {
			return left, false
		} else {
			right := point{sand.r + 1, sand.c + 1}
			if getPoint(right, cave) {
				return right, false
			} else {
				return sand, true
			}
		}
	}
}

func part1() int {
	cave, rMax, cMax, cMin := buildCave(false)
	grains := 0
	start := point{0, 500}
	sand := start

	for sand.c >= cMin && sand.c <= cMax && sand.r <= rMax {
		newPos, rest := fall(sand, cave)
		if rest {
			grains++
			cave[sand.r][sand.c] = false
			sand = start
		} else {
			sand = newPos
		}
	}

	printCave(cMin, cMax, cave)

	return grains
}

func part2() int {
	cave, _, _, _ := buildCave(true)
	grains := 0
	start := point{0, 500}
	sand := start

	for getPoint(start, cave) {
		newPos, rest := fall(sand, cave)
		if rest {
			grains++
			cave[sand.r][sand.c] = false
			sand = start
		} else {
			sand = newPos
		}
	}

	return grains
}

func main() {
	util.Run(part1, part2)
}
