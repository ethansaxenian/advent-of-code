package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

const LENGTH = 36

var inputLines = util.FetchInput(14)

func applyMask(val, mask string) (uint64, error) {
	res := ""

	for i := 0; i < LENGTH; i++ {
		if mask[i] == 'X' {
			res += string(val[i])
		} else {
			res += string(mask[i])
		}
	}

	n, err := strconv.ParseUint(res, 2, LENGTH)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func part1() int {
	var mask string
	var memory = map[int]uint64{}

	for _, line := range inputLines {
		parts := strings.Split(line, " ")
		if parts[0] == "mask" {
			mask = parts[2]
			continue
		}

		mem := util.ToInt(strings.Trim(parts[0], "mem[]"))
		val := fmt.Sprintf("%036b", util.ToInt(parts[2]))

		res, err := applyMask(val, mask)
		if err != nil {
			panic("yo")
		}

		memory[mem] = res

	}

	var sum uint64 = 0

	for _, v := range memory {
		sum += v
	}
	return int(sum)
}

func permute(mask string) []uint64 {
	w := strings.Count(mask, "X")
	permutations := []uint64{}

	for i := 0; i < int(math.Pow(2, float64(w))); i++ {
		combo := fmt.Sprintf("%0*b", w, i)
		j := 0
		perm := strings.Map(func(r rune) rune {
			if r == 'X' {
				j++
				return rune(combo[j-1])
			} else {
				return r
			}
		}, mask)
		n, err := strconv.ParseUint(perm, 2, LENGTH)
		if err != nil {
			panic("why")
		}

		permutations = append(permutations, n)
	}

	return permutations
}

func applyMask2(val, mask string) []uint64 {
	res := ""

	for i := 0; i < LENGTH; i++ {
		if mask[i] == '0' {
			res += string(val[i])
		} else {
			res += string(mask[i])
		}
	}

	perms := permute(res)

	return perms
}

func part2() int {
	var mask string
	var memory = map[uint64]int{}

	for _, line := range inputLines {
		parts := strings.Split(line, " ")
		if parts[0] == "mask" {
			mask = parts[2]
			continue
		}

		val := util.ToInt(parts[2])

		mem := util.ToInt(strings.Trim(parts[0], "mem[]"))
		memBinary := fmt.Sprintf("%036b", mem)

		nums := applyMask2(memBinary, mask)
		for _, m := range nums {
			memory[m] = val
		}
	}

	sum := 0

	for _, v := range memory {
		sum += v
	}
	return sum
}

func main() {
	util.Run(part1, part2)
}
