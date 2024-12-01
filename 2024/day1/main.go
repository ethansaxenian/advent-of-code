package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2024/util"
)

func part1(input []string) int {
	l1, l2 := []int{}, []int{}
	for _, line := range input {
		parts := strings.SplitN(line, "   ", 2)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		l1 = append(l1, x)
		l2 = append(l2, y)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	sum := 0
	for i := 0; i < len(l1); i++ {
		x, y := l1[i], l2[i]
		sum += util.Abs(x - y)
	}

	return sum
}

func part2(input []string) int {
	l1 := []int{}
	m := map[int]int{}
	for _, line := range input {
		parts := strings.SplitN(line, "   ", 2)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		l1 = append(l1, x)
		m[y]++
	}
	sum := 0
	for _, x := range l1 {
		sum += x * m[x]
	}
	return sum
}

func main() {
	util.Run(1, part1, part2)
}
