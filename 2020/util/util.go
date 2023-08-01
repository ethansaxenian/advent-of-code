package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Run[T comparable](part1, part2 func() T) {
	part := "1"
	if len(os.Args) >= 2 {
		part = os.Args[1]
	}

	var ans any
	if part == "1" {
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

	url := fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day)
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

func FetchInputInts(day int) []int {
	input := FetchInput(day)
	ints := []int{}
	for _, n := range input {
		ints = append(ints, ToInt(n))
	}
	return ints
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func Contains(sl []string, el string) bool {
	for _, s := range sl {
		if s == el {
			return true
		}
	}
	return false
}
