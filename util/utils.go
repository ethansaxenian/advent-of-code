package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(part1, part2 func() int) {
	part := 1
	if len(os.Args) >= 2 {
		part = ToInt(os.Args[1])
	}

	var ans any
	if part == 1 {
		ans = part1()
	} else {
		ans = part2()
	}

	fmt.Println(ans)
}

func ReadInput(day int) []string {
	data, err := os.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
