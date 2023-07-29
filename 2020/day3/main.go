package main

import "github.com/ethansaxenian/advent-of-code/2020/util"

var inputLines = util.FetchInput(3)

func part1() int {
	trees := 0
	x, y := 0, 0
	for y < len(inputLines) {
		if inputLines[y][x] == '#' {
			trees++
		}

		x = (x + 3) % len(inputLines[0])
		y += 1
	}
	return trees
}

func part2() int {
	trees := [5]int{0, 0, 0, 0, 0}
	xs := [5]int{0, 0, 0, 0, 0}
	ys := [5]int{0, 0, 0, 0, 0}
	jumps := [5][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	for ys[0] < len(inputLines) {
		for i, xy := range jumps {
			if ys[i] >= len(inputLines) {
				continue
			}

			if inputLines[ys[i]][xs[i]] == '#' {
				trees[i]++
			}

			dx, dy := xy[0], xy[1]

			xs[i] = (xs[i] + dx) % len(inputLines[0])
			ys[i] = ys[i] + dy

		}
	}

	prod := 1
	for _, t := range trees {
		prod *= t
	}
	return prod
}

func main() {
	util.Run(part1, part2)
}
