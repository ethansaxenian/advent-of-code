package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2020/set"
	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var _ = fmt.Println

var inputLines = util.FetchInput(16)

// var inputLines = []string{
// 	"class: 1-3 or 5-7",
// 	"row: 6-11 or 33-44",
// 	"seat: 13-40 or 45-50",
// 	"",
// 	"your ticket:",
// 	"7,1,14",
// 	"",
// 	"nearby tickets:",
// 	"7,3,47",
// 	"40,4,50",
// 	"55,2,20",
// 	"38,6,12",
// }

// var inputLines = []string{
// 	"class: 0-1 or 4-19",
// 	"row: 0-5 or 8-19",
// 	"seat: 0-13 or 16-19",
// 	"",
// 	"your ticket:",
// 	"11,12,13",
// 	"",
// 	"nearby tickets:",
// 	"3,9,18",
// 	"15,1,5",
// 	"5,14,9",
// }

type field struct {
	name   string
	values map[int]bool
}

func (f field) sortedValues() []int {
	values := []int{}
	for k := range f.values {
		values = append(values, k)
	}
	sort.Ints(values)
	return values
}

func parseInput(input []string) ([]field, [][]int) {
	rangeRegex := regexp.MustCompile(`(\d+)-(\d+)`)
	nameRegex := regexp.MustCompile(`([\w\s]+):\s`)
	fields := []field{}
	values := [][]int{}
	for _, line := range input {
		nameMatch := nameRegex.FindStringSubmatch(line)
		if len(nameMatch) > 0 {
			f := field{nameMatch[1], map[int]bool{}}
			rangeMatches := rangeRegex.FindAllStringSubmatch(line, -1)
			for _, m := range rangeMatches {
				for i := util.ToInt(m[1]); i <= util.ToInt(m[2]); i++ {
					f.values[i] = true
				}
			}
			fields = append(fields, f)
		}

		if strings.Contains(line, ",") {
			values = append(values, util.Map[string, int](util.ToInt, strings.Split(line, ",")))
		}
	}
	return fields, values
}

func part1() int {
	fields, tickets := parseInput(inputLines)
	validValues := map[int]bool{}
	for _, f := range fields {
		for v := range f.values {
			validValues[v] = true
		}
	}

	sum := 0

	for _, ticket := range tickets {
		for _, v := range ticket {
			if _, ok := validValues[v]; !ok {
				sum += v
			}
		}
	}
	return sum
}

func part2() int {
	fields, tickets := parseInput(inputLines)
	myTicket := tickets[0]

	validValues := map[int]bool{}
	for _, f := range fields {
		for v := range f.values {
			validValues[v] = true
		}
	}

	validTickets := [][]int{}
	for _, ticket := range tickets {
		isValid := true
		for _, v := range ticket {
			if _, ok := validValues[v]; !ok {
				isValid = false
				break
			}
		}
		if isValid {
			validTickets = append(validTickets, ticket)
		}
	}

	possibleIndexes := map[string]set.Set[int]{}

	for _, f := range fields {
		possibleIndexes[f.name] = set.NewEmptySet[int]()
		for i := range myTicket {
			iIsGood := true
			for _, t := range validTickets {
				v := t[i]
				if _, ok := f.values[v]; !ok {
					iIsGood = false
					break
				}
			}
			if iIsGood {
				possibleIndexes[f.name].Add(i)
			}
		}
	}

	for _, v := range possibleIndexes {
		x := v.Items()
		sort.Ints(x)
	}

	product := 1

	for len(possibleIndexes) > 0 {
		for i := 0; i < 20; i++ {
			count := 0
			var n string
			for k, v := range possibleIndexes {
				if v.Contains(i) {
					count++
					n = k
				}
			}
			if count == 1 {
				if strings.HasPrefix(n, "departure") {
					product *= myTicket[i]
				}
				delete(possibleIndexes, n)
			}
		}
	}

	return product
}

func main() {
	util.Run(part1, part2)
}
