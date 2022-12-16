package main

import (
	"math"
	"regexp"
	"sort"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

type point struct {
	x, y int
}

type sensor struct {
	pos           point
	closestBeacon point
	distToBeacon  int
}

func manhattanDist(a, b point) int {
	return util.Abs(a.x-b.x) + util.Abs(a.y-b.y)
}

var inputLines = util.FetchInput(15)

func parseSensors() ([]sensor, []point) {
	sensors := []sensor{}
	beacons := []point{}

	re := regexp.MustCompile("[-0-9]+")
	for _, line := range inputLines {
		split := re.FindAllString(line, -1)
		sPos := point{util.ToInt(split[0]), util.ToInt(split[1])}
		bPos := point{util.ToInt(split[2]), util.ToInt(split[3])}
		sensors = append(sensors, sensor{sPos, bPos, manhattanDist(sPos, bPos)})
		beacons = append(beacons, bPos)
	}

	return sensors, beacons
}

func getBounds(sensors []sensor) (int, int, int) {
	xMin := math.MaxInt
	xMax := math.MinInt
	maxDist := math.MinInt

	for _, s := range sensors {
		xMin = util.Min(xMin, s.pos.x)
		xMax = util.Max(xMax, s.pos.x)
		maxDist = util.Max(maxDist, s.distToBeacon)
	}

	return maxDist, xMin, xMax
}

func cantContainBeacon(p point, sensors []sensor) bool {
	for _, s := range sensors {
		if manhattanDist(p, s.pos) <= s.distToBeacon {
			return true
		}
	}
	return false
}

func part1() int {
	sensors, beacons := parseSensors()
	maxDist, xMin, xMax := getBounds(sensors)
	numSpaces := 0
	y := 2000000
	for x := xMin - maxDist; x < xMax+maxDist; x++ {
		p := point{x, y}
		if cantContainBeacon(p, sensors) && !util.Contains(beacons, p) {
			numSpaces++
		}
	}
	return numSpaces
}

type sensorArea struct {
	top, bottom, left, right, mid point
	radius                        int
}

func (a sensorArea) findRightEdgeX(y int) int {
	if y < a.left.y {
		dy := y - a.top.y
		return a.top.x + dy
	} else if y > a.left.y {
		dy := a.bottom.y - y
		return a.bottom.x + dy
	} else {
		return a.right.x
	}
}

func part2() int {
	sensors, _ := parseSensors()
	areas := []sensorArea{}
	for _, s := range sensors {
		areas = append(areas, sensorArea{
			top:    point{s.pos.x, s.pos.y - s.distToBeacon},
			bottom: point{s.pos.x, s.pos.y + s.distToBeacon},
			left:   point{s.pos.x - s.distToBeacon, s.pos.y},
			right:  point{s.pos.x + s.distToBeacon, s.pos.y},
			mid:    s.pos,
			radius: s.distToBeacon,
		})
	}

	sort.Slice(areas, func(i, j int) bool { return areas[i].mid.y < areas[j].mid.y })

	upperBound := 4000000
	for y := 0; y <= upperBound; y++ {
		x := 0
		for x <= upperBound {
			isDistressBeacon := true
			for _, a := range areas {
				if manhattanDist(a.mid, point{x, y}) <= a.radius {
					x = a.findRightEdgeX(y) + 1
					isDistressBeacon = false
				}
			}
			if isDistressBeacon {
				return x*upperBound + y
			}
		}
	}
	return 0
}

func main() {
	util.Run(part1, part2)
}
