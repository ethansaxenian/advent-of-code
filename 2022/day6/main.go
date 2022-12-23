package main

import (
	"github.com/ethansaxenian/advent-of-code/2022/util"
)

var input = util.FetchInput(6)[0]

func loadCounter(num int) map[byte]int {
	counter := map[byte]int{}
	for i := 0; i < num; i++ {
		if val, ok := counter[input[i]]; ok {
			counter[input[i]] = val + 1
		} else {
			counter[input[i]] = 1
		}
	}

	return counter
}

func part1() int {
	counter := loadCounter(4)
	i := 0
	j := 4
	for j < len(input) {
		if v := counter[input[i]]; v == 1 {
			delete(counter, input[i])
		} else {
			counter[input[i]] -= 1
		}

		counter[input[j]]++

		if len(counter) == 4 {
			return j + 1
		}
		i++
		j++
	}
	return 0
}

func part2() int {
	counter := loadCounter(14)
	i := 0
	j := 14
	for j < len(input) {
		if v := counter[input[i]]; v == 1 {
			delete(counter, input[i])
		} else {
			counter[input[i]] -= 1
		}

		counter[input[j]]++

		if len(counter) == 14 {
			return j + 1
		}

		i++
		j++
	}
	return 0
}

func main() {
	util.Run(part1, part2)
}
