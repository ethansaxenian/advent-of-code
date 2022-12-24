package util

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Run[T comparable](part1, part2 func() T) {
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

type session struct {
	Cookie string `json:"aoc_cookie"`
}

func FetchInput(day int) []string {
	_, filename, _, _ := runtime.Caller(0)
	cookiePath := filepath.Join(filename, "../../../aoc-cookie.json")
	data, err := os.ReadFile(cookiePath)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	var session session

	err = json.Unmarshal(data, &session)
	if err != nil {
		panic(err)
	}

	cookieHeader := fmt.Sprintf("session=%s", session.Cookie)
	req.Header.Add("cookie", cookieHeader)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(body), "\n")
	return lines[:len(lines)-1]
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Contains[T comparable](list []T, item T) bool {
	for _, x := range list {
		if x == item {
			return true
		}
	}
	return false
}

func Max(ints ...int) int {
	m := math.MinInt
	for _, i := range ints {
		if i > m {
			m = i
		}
	}
	return m
}

func Min(ints ...int) int {
	m := math.MaxInt
	for _, i := range ints {
		if i < m {
			m = i
		}
	}
	return m
}

func Mod(x, y int) int {
	m := x % y
	if m < 0 {
		m += y
	}
	return m
}

/*
def gcd(a, b):
    """Return greatest common divisor using Euclid's Algorithm."""
    while b:
        a, b = b, a % b
    return a

def lcm(a, b):
    """Return lowest common multiple."""
    return a * b // gcd(a, b)
*/

func GCD(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
