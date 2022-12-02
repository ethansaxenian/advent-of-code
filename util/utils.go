package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
