package main

import (
	"regexp"
	"strings"

	"github.com/ethansaxenian/advent-of-code-2022/stack"
	"github.com/ethansaxenian/advent-of-code-2022/util"
)

var inputLines = util.FetchInput(5)

var stacks = [9]stack.Stack[string]{
	stack.NewStack("D", "B", "J", "V"),
	stack.NewStack("P", "V", "B", "W", "R", "D", "F"),
	stack.NewStack("R", "G", "F", "L", "D", "C", "W", "Q"),
	stack.NewStack("W", "J", "P", "M", "L", "N", "D", "B"),
	stack.NewStack("H", "N", "B", "P", "C", "S", "Q"),
	stack.NewStack("R", "D", "B", "S", "N", "G"),
	stack.NewStack("Z", "B", "P", "M", "Q", "F", "S", "H"),
	stack.NewStack("W", "L", "F"),
	stack.NewStack("S", "V", "F", "M", "R"),
}

func parseAction(line string) (int, int, int) {
	re := regexp.MustCompile("[0-9]+")
	parts := re.FindAllString(line, -1)
	return util.ToInt(parts[0]), util.ToInt(parts[1]) - 1, util.ToInt(parts[2]) - 1
}

func part1() string {
	for _, action := range inputLines[10:] {
		num, from, to := parseAction(action)
		for i := 0; i < num; i++ {
			poppedStack, elem := stacks[from].Pop()
			pushedStack := stacks[to].Push(elem)
			stacks[from] = poppedStack
			stacks[to] = pushedStack
		}
	}

	ans := [9]string{}
	for i, stack := range stacks {
		ans[i] = stack.Peek()
	}

	return strings.Join(ans[:], "")
}

func part2() string {
	for _, action := range inputLines[10:] {
		num, from, to := parseAction(action)
		tmpStack := stack.NewStack[string]()
		for i := 0; i < num; i++ {
			poppedStack, elem := stacks[from].Pop()
			tmpStack = tmpStack.Push(elem)
			stacks[from] = poppedStack
		}
		for i := 0; i < num; i++ {
			poppedStack, elem := tmpStack.Pop()
			stacks[to] = stacks[to].Push(elem)
			tmpStack = poppedStack
		}
	}

	ans := [9]string{}
	for i, stack := range stacks {
		ans[i] = stack.Peek()
	}

	return strings.Join(ans[:], "")
}

func main() {
	util.Run(part1, part2)
}
