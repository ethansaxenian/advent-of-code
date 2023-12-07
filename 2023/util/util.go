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

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2023/day/%x/input", day), nil)
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
