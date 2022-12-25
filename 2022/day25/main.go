package main

import (
	"strconv"

	"github.com/ethansaxenian/advent-of-code/2022/util"
)

var inputLines = util.FetchInput(25)

var digits = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'-': -1,
	'=': -2,
}

func revMap(m map[byte]int) map[int]byte {
	rev := map[int]byte{}
	for k, v := range m {
		rev[v] = k
	}
	return rev
}

func parseSNAFU(snafu string) int {
	b10 := 0
	for i := len(snafu) - 1; i >= 0; i-- {
		b10 += digits[snafu[i]] * util.Pow(5, len(snafu)-i-1)
	}
	return b10
}

func convertToSNAFU(b10 int) string {
	snafu := ""
	for b10 > 0 {
		mod := b10 % 5
		switch mod {
		case 0, 1, 2:
			snafu = strconv.Itoa(mod) + snafu
		default:
			snafu = string(revMap(digits)[mod-5]) + snafu
			b10 += 5
		}
		b10 /= 5
	}
	return snafu
}

func part1() string {
	sum := 0
	for _, snafu := range inputLines {
		sum += parseSNAFU(snafu)
	}

	return convertToSNAFU(sum)
}

func part2() string {
	return ""
}

func main() {
	util.Run(part1, part2)
}
