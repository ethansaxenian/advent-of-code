package main

import (
	"fmt"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2022/util"
)

var inputLines = util.FetchInput(21)

func extractMonkeys() map[string]string {
	monkeyMap := map[string]string{}
	for _, line := range inputLines {
		parts := strings.Split(line, ": ")
		monkeyMap[parts[0]] = parts[1]
	}
	return monkeyMap
}

func yell(monkey string, monkeyMap map[string]string) int {
	parts := strings.Split(monkeyMap[monkey], " ")

	if len(parts) == 1 {
		return util.ToInt(parts[0])
	}

	left := yell(parts[0], monkeyMap)
	right := yell(parts[2], monkeyMap)

	switch parts[1] {
	case "*":
		return left * right
	case "+":
		return left + right
	case "-":
		return left - right
	case "/":
		return left / right
	default:
		panic("unknown operator")
	}
}

func generateEquation(monkey string, monkeyMap map[string]string) string {
	if monkey == "humn" {
		return "x"
	}

	parts := strings.Split(monkeyMap[monkey], " ")
	if len(parts) == 1 {
		return parts[0]
	}

	return "(" + generateEquation(parts[0], monkeyMap) + "" + parts[1] + "" + generateEquation(parts[2], monkeyMap) + ")"
}

func part1() int {
	monkeyMap := extractMonkeys()
	return yell("root", monkeyMap)
}

func part2() int {
	monkeyMap := extractMonkeys()

	parts := strings.Split(monkeyMap["root"], " ")

	equation := generateEquation(parts[0], monkeyMap) + "=" + generateEquation(parts[2], monkeyMap)
	// plug this into an online equation solver :)
	fmt.Println(equation)
	return 0
}

func main() {
	util.Run(part1, part2)
}
