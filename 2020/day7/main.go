package main

import (
	"strings"

	"github.com/ethansaxenian/advent-of-code/2020/util"
)

var inputLines = util.FetchInput(7)

// var inputLines = []string{
// 	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
// 	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
// 	"bright white bags contain 1 shiny gold bag.",
// 	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
// 	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
// 	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
// 	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
// 	"faded blue bags contain no other bags.",
// 	"dotted black bags contain no other bags.",
// }

// var inputLines = []string{
// 	"shiny gold bags contain 2 dark red bags.",
// 	"dark red bags contain 2 dark orange bags.",
// 	"dark orange bags contain 2 dark yellow bags.",
// 	"dark yellow bags contain 2 dark green bags.",
// 	"dark green bags contain 2 dark blue bags.",
// 	"dark blue bags contain 2 dark violet bags.",
// 	"dark violet bags contain no other bags.",
// }

type bagWithCount struct {
	c string
	n int
}

var cache = map[string]bool{}

func containsShinyGold(bag string, bags map[string][]string) bool {
	if bag == "shiny gold" {
		return true
	}

	innerBags := bags[bag]

	if len(innerBags) == 0 {
		return false
	}

	if hasShinyGold, ok := cache[bag]; ok {
		return hasShinyGold
	}

	for _, c := range innerBags {
		yeah := containsShinyGold(c, bags)
		cache[c] = yeah
		if yeah {
			return true
		}
	}
	return false
}

func howManyBags(bag string, bags map[string][]bagWithCount) int {
	totNum := 0

	for _, c := range bags[bag] {
		num := howManyBags(c.c, bags)
		if num > 0 {
			totNum += num * c.n
		}
		totNum += c.n
	}

	return totNum
}

func part1() int {
	bags := map[string][]string{}
	for _, line := range inputLines {
		words := strings.Split(line, " ")
		key := strings.Join(words[:2], " ")
		bags[key] = []string{}
		for i, w := range words[4:] {
			if w == "no" {
				break
			}
			if i%4 == 0 {
				innerBag := strings.Join([]string{words[i+5], words[i+6]}, " ")
				bags[key] = append(bags[key], innerBag)
			}
		}
	}

	var count int
	for k := range bags {
		if containsShinyGold(k, bags) {
			count++
		}
	}
	return count - 1
}

func part2() int {
	bags := map[string][]bagWithCount{}
	for _, line := range inputLines {
		words := strings.Split(line, " ")
		key := strings.Join(words[:2], " ")
		bags[key] = []bagWithCount{}
		for i, w := range words[4:] {
			if w == "no" {
				break
			}
			if i%4 == 0 {
				innerBag := strings.Join([]string{words[i+5], words[i+6]}, " ")
				num := util.ToInt(words[i+4])
				bags[key] = append(bags[key], bagWithCount{innerBag, num})
			}
		}
	}

	return howManyBags("shiny gold", bags)
}

func main() {
	util.Run(part1, part2)
}
