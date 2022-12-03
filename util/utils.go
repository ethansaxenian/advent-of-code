package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
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

func FetchInput(day int) []string {
	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)
	client := &http.Client{}
	_, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	cookieHeader := fmt.Sprintf("session=%s", os.Getenv("AOC_COOKIE"))
	req.Header.Add("cookie", cookieHeader)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(body), "\n")
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
