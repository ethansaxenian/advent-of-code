package main

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

type session struct {
	Cookie string `json:"aoc_cookie"`
}

func FetchInput() []string {
	_, filename, _, _ := runtime.Caller(0)
	cookiePath := filepath.Join(filename, "../../aoc-cookie.json")
	data, err := os.ReadFile(cookiePath)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2023/day/5/input", nil)
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

func parseInput(lines []string) ([]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int) {
	seeds := StrListToInts(strings.Split(lines[0], " ")[1:])

	m := map[string][][3]int{}
	var curr string

	for _, line := range lines[2:] {
		if line == "" {
			continue
		}

		if string(line[len(line)-1]) == ":" {
			curr = line
			continue
		}

		splits := StrListToInts(strings.Split(line, " "))
		d, s, l := splits[0], splits[1], splits[2]
		m[curr] = append(m[curr], [3]int{s, d, l})
	}

	return seeds, m["seed-to-soil map:"], m["soil-to-fertilizer map:"], m["fertilizer-to-water map:"], m["water-to-light map:"], m["light-to-temperature map:"], m["temperature-to-humidity map:"], m["humidity-to-location map:"]
}

func convertNumber(num int, m [][3]int) int {
	for _, arr := range m {
		s, d, l := arr[0], arr[1], arr[2]
		if num >= s && num < s+l {
			return d + (num - s)
		}
	}
	return num
}

func day1(input []string) {
	seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation := parseInput(input)

	minLoc := math.MaxInt64
	for _, seed := range seeds {
		soil := convertNumber(seed, seedToSoil)
		fertilizer := convertNumber(soil, soilToFertilizer)
		water := convertNumber(fertilizer, fertilizerToWater)
		light := convertNumber(water, waterToLight)
		temperature := convertNumber(light, lightToTemperature)
		humidity := convertNumber(temperature, temperatureToHumidity)
		location := convertNumber(humidity, humidityToLocation)

		if location < minLoc {
			minLoc = location
		}
	}

	fmt.Println(minLoc)

}

func day2(input []string) {
	seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation := parseInput(input)

	minLoc := math.MaxInt64
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed <= seeds[i]+seeds[i+1]; seed++ {
			soil := convertNumber(seed, seedToSoil)
			fertilizer := convertNumber(soil, soilToFertilizer)
			water := convertNumber(fertilizer, fertilizerToWater)
			light := convertNumber(water, waterToLight)
			temperature := convertNumber(light, lightToTemperature)
			humidity := convertNumber(temperature, temperatureToHumidity)
			location := convertNumber(humidity, humidityToLocation)

			if location < minLoc {
				minLoc = location
			}
		}
	}

	fmt.Println(minLoc)
}

func main() {
	input := FetchInput()
	day1(input)
	day2(input)
}
