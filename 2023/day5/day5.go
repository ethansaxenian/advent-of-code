package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2023/util"
)

func parseInput(lines []string) ([]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int, [][3]int) {
	seeds := util.StrListToInts(strings.Split(lines[0], " ")[1:])

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

		splits := util.StrListToInts(strings.Split(line, " "))
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

func part1(input []string) {
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

func part2(input []string) {
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
	input := util.FetchInput(5)
	part1(input)
	part2(input)
}
