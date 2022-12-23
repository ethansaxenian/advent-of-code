package main

import (
	"github.com/ethansaxenian/advent-of-code/2022/util"
)

type monkey struct {
	items          []int
	getWorryLevel  func(old int) int
	test           int
	ifTrue         int
	ifFalse        int
	numInspections int
}

var monkeys = []*monkey{
	{
		items:         []int{63, 84, 80, 83, 84, 53, 88, 72},
		getWorryLevel: func(old int) int { return old * 11 },
		test:          13,
		ifTrue:        4,
		ifFalse:       7,
	},
	{
		items:         []int{67, 56, 92, 88, 84},
		getWorryLevel: func(old int) int { return old + 4 },
		test:          11,
		ifTrue:        5,
		ifFalse:       3,
	},
	{
		items:         []int{52},
		getWorryLevel: func(old int) int { return old * old },
		test:          2,
		ifTrue:        3,
		ifFalse:       1,
	},
	{
		items:         []int{59, 53, 60, 92, 69, 72},
		getWorryLevel: func(old int) int { return old + 2 },
		test:          5,
		ifTrue:        5,
		ifFalse:       6,
	},
	{
		items:         []int{61, 52, 55, 61},
		getWorryLevel: func(old int) int { return old + 3 },
		test:          7,
		ifTrue:        7,
		ifFalse:       2,
	},
	{
		items:         []int{79, 53},
		getWorryLevel: func(old int) int { return old + 1 },
		test:          3,
		ifTrue:        0,
		ifFalse:       6,
	},
	{
		items:         []int{59, 86, 67, 95, 92, 77, 91},
		getWorryLevel: func(old int) int { return old + 5 },
		test:          19,
		ifTrue:        4,
		ifFalse:       0,
	},
	{
		items:         []int{58, 83, 89},
		getWorryLevel: func(old int) int { return old * 19 },
		test:          17,
		ifTrue:        2,
		ifFalse:       1,
	},
}

func (m *monkey) inspect(item int) int {
	worry := m.getWorryLevel(item)
	m.numInspections++
	return worry
}

func (m *monkey) throw(item int) {
	var nextMonkey *monkey
	if item%m.test == 0 {
		nextMonkey = monkeys[m.ifTrue]
	} else {
		nextMonkey = monkeys[m.ifFalse]
	}
	nextMonkey.items = append(nextMonkey.items, item)
}

func getBusiness() int {
	first, second := monkeys[0].numInspections, monkeys[1].numInspections
	if second > first {
		first, second = second, first
	}
	for i := 2; i < len(monkeys); i++ {
		if monkeys[i].numInspections > second {
			second = monkeys[i].numInspections
		}
		if second > first {
			first, second = second, first
		}
	}
	return first * second
}

func part1() int {
	for r := 0; r < 20; r++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				worry := m.inspect(item)
				m.throw(worry / 3)
			}
			m.items = []int{}
		}
	}

	return getBusiness()
}

func part2() int {
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.test
	}
	for r := 0; r < 10000; r++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				worry := m.inspect(item)
				m.throw(worry % lcm)
			}
			m.items = []int{}
		}
	}

	return getBusiness()
}

func main() {
	util.Run(part1, part2)
}
