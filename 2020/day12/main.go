package main

import (
	"fmt"
	"math"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInput(12)

// var inputLines = []string{
// 	"F10",
// 	"N3",
// 	"F7",
// 	"R90",
// 	"F11",
// }

var dToP = map[int][2]int{
	0:   {1, 0},
	90:  {0, 1},
	180: {-1, 0},
	270: {0, -1},
}

func part1() int {
	x := 0
	y := 0
	d := 0
	for _, line := range inputLines {
		n := util.ToInt(line[1:])
		switch line[0] {
		case 'N':
			y += n
		case 'S':
			y -= n
		case 'E':
			x += n
		case 'W':
			x -= n
		case 'F':
			x += dToP[d][0] * n
			y += dToP[d][1] * n
		case 'L':
			d = util.Mod(d+n, 360)
		case 'R':
			d = util.Mod(d-n, 360)
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func rotate(d, x, y int) (int, int) {
	switch d {
	case 0:
		return x, y
	case 90, -270:
		return y, -x
	case 180, -180:
		return -x, -y
	case 270, -90:
		return -y, x
	default:
		panic("eek")
	}
}

func part2() int {
	x, y := 0, 0
	wx, wy := 10, 1
	for _, line := range inputLines {
		n := util.ToInt(line[1:])
		switch line[0] {
		case 'N':
			wy += n
		case 'S':
			wy -= n
		case 'E':
			wx += n
		case 'W':
			wx -= n
		case 'F':
			x += wx * n
			y += wy * n
		case 'L':
			wx, wy = rotate(-n, wx, wy)
		case 'R':
			wx, wy = rotate(n, wx, wy)
		}
		fmt.Printf("%v -> [%v, %v] {%v, %v}\n", line, x, y, wx, wy)
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func main() {
	util.Run(part1, part2)
}
