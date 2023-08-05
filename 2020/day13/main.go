package main

import (
	"strings"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInput(13)

// var inputLines = []string{
// 	"111",
// 	"7,13,x,x,59,x,31,19",
// }

func part1() int {
	timestamp := util.ToInt(inputLines[0])
	busIds := util.Map[string, int](util.ToInt, strings.Split(strings.ReplaceAll(inputLines[1], ",x", ""), ","))

	i := timestamp
	for {
		for _, id := range busIds {
			if i%id == 0 {
				return id * (i - timestamp)
			}
		}
		i++
	}
}

func _xi(Ni, mod int) int {
	// Ni * xi = 1 (% mod)
	Ni = Ni % mod
	xi := 1
	for {
		if (Ni*xi)%mod == 1 {
			return xi
		}
		xi++
	}
}

// https://www.youtube.com/watch?v=zIFehsBHB8o
func part2() int {
	busIds := strings.Split(inputLines[1], ",")

	ids := [][2]int{}
	idProduct := 1
	for pos, id := range busIds {
		if id == "x" {
			continue
		}
		ids = append(ids, [2]int{pos, util.ToInt(id)})
		idProduct *= util.ToInt(id)
	}

	sum := 0
	for _, posId := range ids {
		pos, id := posId[0], posId[1]

		bi := id - pos
		Ni := idProduct / id
		xi := _xi(Ni, id)

		sum += (bi * Ni * xi)
	}
	return sum % idProduct
}

func main() {
	util.Run(part1, part2)
}
