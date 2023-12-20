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

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day), nil)
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

func StrListToInts(strList []string) []int {
	intList := []int{}
	for _, s := range strList {
		i, _ := strconv.Atoi(s)
		intList = append(intList, i)
	}
	return intList
}

func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func Contains[T comparable](list []T, item T) bool {
	for _, x := range list {
		if x == item {
			return true
		}
	}
	return false
}

func Combinations[T comparable](list []T, length int) [][]T {
	if length == 1 {
		combs := [][]T{}
		for _, item := range list {
			combs = append(combs, []T{item})
		}
		return combs
	}

	combs := [][]T{}
	for i, item := range list {
		for _, comb := range Combinations[T](list[i+1:], length-1) {
			combs = append(combs, append([]T{item}, comb...))
		}
	}
	return combs
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ShortestPath(a, b [2]int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

func Map[T, U any](list []T, f func(T) U) []U {
	newList := []U{}
	for _, item := range list {
		newList = append(newList, f(item))
	}
	return newList
}

func All[T any](list []T, f func(T) bool) bool {
	for _, item := range list {
		if !f(item) {
			return false
		}
	}
	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
