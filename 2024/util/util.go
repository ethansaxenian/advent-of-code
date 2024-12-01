package util

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	ROOT_DIR string
	COOKIE   string
	EMAIL    string

	partFlag int
	testFlag bool
)

func init() {
	_, curr_file, _, _ := runtime.Caller(0)
	ROOT_DIR = filepath.Dir(filepath.Dir(curr_file))
	godotenv.Load(filepath.Join(filepath.Dir(ROOT_DIR), ".env"))
	COOKIE = os.Getenv("AOC_COOKIE")
	EMAIL = os.Getenv("AOC_EMAIL")

	flag.IntVar(&partFlag, "part", 1, "which part to run")
	flag.BoolVar(&testFlag, "test", false, "use test input")
	flag.Parse()
}

func Run[T comparable](day int, part1, part2 func([]string) T) {
	var input []string
	if testFlag {
		input = FetchInputFromStdin()
	} else {
		input = FetchInput(day)
	}

	var ans any
	switch partFlag {
	case 1:
		ans = part1(input)
	case 2:
		ans = part2(input)
	default:
		panic("invalid part")
	}

	if !testFlag {
		var stdin string
		switch v := ans.(type) {
		case int:
			stdin = strconv.Itoa(v)
		case string:
			stdin = v
		}
		cmd := exec.Command("pbcopy")
		cmd.Stdin = strings.NewReader(stdin)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(ans)
}

func FetchInput(day int) []string {
	inputPath := filepath.Join(ROOT_DIR, fmt.Sprintf("input/day%d.txt", day))
	if _, err := os.Stat(inputPath); err == nil {
		data, err := os.ReadFile(inputPath)
		if err != nil {
			panic(err)
		}

		return strings.Split(strings.Trim(string(data), " "), "\n")
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	cookieHeader := fmt.Sprintf("session=%s", COOKIE)
	req.Header.Add("cookie", cookieHeader)

	userAgentHeader := fmt.Sprintf("github.com/ethansaxenian/advent-of-code by %s", EMAIL)
	req.Header.Add("cookie", userAgentHeader)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(inputPath, body, 0644); err != nil {
		panic(err)
	}

	return strings.Split(strings.Trim(string(body), " "), "\n")
}

func FetchInputFromStdin() []string {
	var lines []string
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		line := scn.Text()
		if strings.ContainsRune(line, '\x1D') {
			break
		}
		lines = append(lines, line)
	}

	return lines
}
