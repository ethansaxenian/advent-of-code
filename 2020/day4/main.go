package main

import (
	"aoc2020/util"
	"strings"
)

var inputLines = util.FetchInput(4)

var requiredFields [7]string = [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var validecl []string = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func getPassportStrings(lines []string) []string {
	passports := []string{}
	i := 0
	for i < len(lines) {
		passport := lines[i]
		for lines[i] != "" {
			i++
			if i >= len(lines) {
				break
			}
			passport += " " + lines[i]
		}
		passports = append(passports, passport)
		i++
	}

	return passports
}

func getPassportMap(passportString string) map[string]string {
	passport := map[string]string{}
	passportFields := strings.Split(passportString, " ")
	for _, field := range passportFields {
		if field == "" {
			continue
		}
		kv := strings.Split(field, ":")
		passport[kv[0]] = kv[1]
	}
	return passport
}

func validatePassport(passport map[string]string) bool {
	byr := util.ToInt(passport["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr := util.ToInt(passport["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr := util.ToInt(passport["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt := util.ToInt(strings.TrimRight(passport["hgt"], "cmin"))
	if !(strings.HasSuffix(passport["hgt"], "cm") || strings.HasSuffix(passport["hgt"], "in")) {
		return false
	}
	if strings.HasSuffix(passport["hgt"], "cm") && (hgt < 150 || hgt > 193) {
		return false
	}
	if strings.HasSuffix(passport["hgt"], "in") && (hgt < 59 || hgt > 76) {
		return false
	}

	if strings.TrimRight(passport["hcl"], "0123456789abcdef") != "#" {
		return false
	}

	if !util.Contains(validecl, passport["ecl"]) {
		return false
	}

	if len(passport["pid"]) != 9 || strings.TrimRight(passport["pid"], "0123456789") != "" {
		return false
	}

	return true
}

func part1() int {
	numValid := 0
	passportStrings := getPassportStrings(inputLines)
	for _, p := range passportStrings {
		passport := getPassportMap(p)
		_, hascid := passport["cid"]
		if len(passport) == 8 || len(passport) == 7 && !hascid {
			numValid++
		}
	}
	return numValid
}

func part2() int {
	numValid := 0
	passportStrings := getPassportStrings(inputLines)
	for _, p := range passportStrings {
		passport := getPassportMap(p)
		_, hascid := passport["cid"]
		if len(passport) == 8 || len(passport) == 7 && !hascid {
			if validatePassport(passport) {
				numValid++
			}
		}
	}
	return numValid
}

func main() {
	util.Run(part1, part2)
}
