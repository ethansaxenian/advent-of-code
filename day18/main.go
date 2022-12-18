package main

import (
	"math"
	"strings"

	"github.com/ethansaxenian/advent-of-code-2022/set"
	"github.com/ethansaxenian/advent-of-code-2022/util"
)

var inputLines = util.FetchInput(18)

func getDrops() set.Set[[3]int] {
	drops := set.NewEmptySet[[3]int]()
	for _, line := range inputLines {
		xyz := strings.Split(line, ",")
		coord := [3]int{util.ToInt(xyz[0]), util.ToInt(xyz[1]), util.ToInt(xyz[2])}
		drops.Add(coord)
	}

	return drops
}

func getNeighbors(coord [3]int) [][3]int {
	neighbors := [][3]int{}
	for _, dir := range [][3]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}} {
		neighbors = append(neighbors, [3]int{coord[0] + dir[0], coord[1] + dir[1], coord[2] + dir[2]})
	}
	return neighbors
}

func isInBounds(coord [3]int, min, max int) bool {
	return coord[0] >= min && coord[0] <= max && coord[1] >= min && coord[1] <= max && coord[2] >= min && coord[2] <= max
}

func part1() int {
	drops := getDrops()
	surfaceArea := 0
	for _, coord := range drops.Items() {
		for _, newDrop := range getNeighbors(coord) {
			if !drops.Contains(newDrop) {
				surfaceArea++
			}
		}
	}
	return surfaceArea
}

func part2() int {
	drops := getDrops()

	min := math.MaxInt
	max := math.MinInt
	for _, coord := range drops.Items() {
		min = util.Min(min, coord[0], coord[1], coord[2])
		max = util.Max(max, coord[0], coord[1], coord[2])
	}
	min--
	max++

	steamArea := 0
	steam := set.NewEmptySet[[3]int]()
	q := [][3]int{{min, min, min}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, newCoord := range getNeighbors(curr) {
			if drops.Contains(newCoord) {
				steamArea++
			} else if isInBounds(newCoord, min, max) && !steam.Contains(newCoord) {
				q = append(q, newCoord)
				steam.Add(newCoord)
			}
		}
	}

	return steamArea
}

func main() {
	util.Run(part1, part2)
}
