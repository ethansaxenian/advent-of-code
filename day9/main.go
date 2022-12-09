package main

import (
	"strings"

	"github.com/ethansaxenian/advent-of-code-2022/set"
	"github.com/ethansaxenian/advent-of-code-2022/util"
)

var inputLines = util.FetchInput(9)

var m = map[string][2]int{
	"R": {0, 1},
	"L": {0, -1},
	"U": {1, 0},
	"D": {-1, 0},
}

func getParts(s string) (int, int, int) {
	parts := strings.Split(s, " ")
	direction := m[parts[0]]
	steps := util.ToInt(parts[1])
	return direction[0], direction[1], steps
}

func getPos(t, h [2]int) [2]int {
	dr := h[0] - t[0]
	dc := h[1] - t[1]
	switch [2]int{dr, dc} {
	case [2]int{2, 0}, [2]int{2, 1}, [2]int{2, -1}:
		return [2]int{h[0] - 1, h[1]}
	case [2]int{0, 2}, [2]int{1, 2}, [2]int{-1, 2}:
		return [2]int{h[0], h[1] - 1}
	case [2]int{-2, 0}, [2]int{-2, 1}, [2]int{-2, -1}:
		return [2]int{h[0] + 1, h[1]}
	case [2]int{0, -2}, [2]int{1, -2}, [2]int{-1, -2}:
		return [2]int{h[0], h[1] + 1}
	case [2]int{2, 2}:
		return [2]int{h[0] - 1, h[1] - 1}
	case [2]int{-2, -2}:
		return [2]int{h[0] + 1, h[1] + 1}
	case [2]int{2, -2}:
		return [2]int{h[0] - 1, h[1] + 1}
	case [2]int{-2, 2}:
		return [2]int{h[0] + 1, h[1] - 1}
	default:
		return t
	}
}

func part1() int {
	vis := set.NewEmptySet[[2]int]()
	vis.Add([2]int{0, 0})
	h, t := [2]int{0, 0}, [2]int{0, 0}
	for _, instruction := range inputLines {
		r, c, steps := getParts(instruction)
		for i := 0; i < steps; i++ {
			h[0] += r
			h[1] += c
			t = getPos(t, h)
			vis.Add(t)
		}
	}
	return len(vis)
}

func part2() int {
	vis := set.NewEmptySet[[2]int]()
	vis.Add([2]int{0, 0})
	rope := [10][2]int{}
	for i := 0; i < 10; i++ {
		rope[i] = [2]int{0, 0}
	}

	for _, instruction := range inputLines {
		r, c, steps := getParts(instruction)
		for i := 0; i < steps; i++ {
			rope[0][0] += r
			rope[0][1] += c
			for j := 1; j < 10; j++ {
				rope[j] = getPos(rope[j], rope[j-1])
				if j == 9 {
					vis.Add(rope[j])
				}
			}
		}
	}
	return len(vis)
}

func main() {
	util.Run(part1, part2)
}
