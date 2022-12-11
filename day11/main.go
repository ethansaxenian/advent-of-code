package main

import (
	"fmt"
	"sort"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

type monkey struct {
	items          []int
	getWorryLevel  func(old int) int
	test           int
	ifTrue         int
	ifFalse        int
	numInspections int
}

// var monkeys = map[int]monkey{
// 	0: {
// 		items:          []int{63, 84, 80, 83, 84, 53, 88, 72},
// 		getWorryLevel:  func(old int) int { return old * 11 },
// 		test:           13,
// 		ifTrue:         4,
// 		ifFalse:        7,
// 		numInspections: 0,
// 	},
// 	1: {
// 		items:          []int{67, 56, 92, 88, 84},
// 		getWorryLevel:  func(old int) int { return old + 4 },
// 		test:           11,
// 		ifTrue:         5,
// 		ifFalse:        3,
// 		numInspections: 0,
// 	},
// 	2: {
// 		items:          []int{52},
// 		getWorryLevel:  func(old int) int { return old * old },
// 		test:           2,
// 		ifTrue:         3,
// 		ifFalse:        1,
// 		numInspections: 0,
// 	},
// 	3: {
// 		items:          []int{59, 53, 60, 92, 69, 72},
// 		getWorryLevel:  func(old int) int { return old + 2 },
// 		test:           5,
// 		ifTrue:         5,
// 		ifFalse:        6,
// 		numInspections: 0,
// 	},
// 	4: {
// 		items:          []int{61, 52, 55, 61},
// 		getWorryLevel:  func(old int) int { return old + 3 },
// 		test:           7,
// 		ifTrue:         7,
// 		ifFalse:        2,
// 		numInspections: 0,
// 	},
// 	5: {
// 		items:          []int{79, 53},
// 		getWorryLevel:  func(old int) int { return old + 1 },
// 		test:           3,
// 		ifTrue:         0,
// 		ifFalse:        6,
// 		numInspections: 0,
// 	},
// 	6: {
// 		items:          []int{59, 86, 67, 95, 92, 77, 91},
// 		getWorryLevel:  func(old int) int { return old + 5 },
// 		test:           19,
// 		ifTrue:         4,
// 		ifFalse:        0,
// 		numInspections: 0,
// 	},
// 	7: {
// 		items:          []int{58, 83, 89},
// 		getWorryLevel:  func(old int) int { return old * 19 },
// 		test:           17,
// 		ifTrue:         2,
// 		ifFalse:        1,
// 		numInspections: 0,
// 	},
// }

var monkeys = map[int]monkey{
	0: {
		items:         []int{79, 98},
		getWorryLevel: func(old int) int { return old * 19 },
		test:          23,
		ifTrue:        2,
		ifFalse:       3,
	},
	1: {
		items:         []int{54, 65, 75, 74},
		getWorryLevel: func(old int) int { return old + 6 },
		test:          19,
		ifTrue:        2,
		ifFalse:       0,
	},
	2: {
		items:         []int{79, 60, 97},
		getWorryLevel: func(old int) int { return old * old },
		test:          13,
		ifTrue:        1,
		ifFalse:       3,
	},
	3: {
		items:         []int{74},
		getWorryLevel: func(old int) int { return old + 3 },
		test:          17,
		ifTrue:        0,
		ifFalse:       1,
	},
}

func (m monkey) throw(item int) {
	if item%m.test == 0 {
		next := monkeys[m.ifTrue]
		next.items = append(next.items, item)
		monkeys[m.ifTrue] = next
	} else {
		next := monkeys[m.ifFalse]
		next.items = append(next.items, item)
		monkeys[m.ifFalse] = next
	}
}

func part1() int {
	for r := 0; r < 20; r++ {
		for i := 0; i < 8; i++ {
			m := monkeys[i]
			for _, item := range m.items {
				worry := m.getWorryLevel(item)
				m.numInspections++
				m.throw(worry / 3)
			}
			m.items = []int{}
			monkeys[i] = m
		}
	}

	inspections := []int{}
	for i := 0; i < 8; i++ {
		inspections = append(inspections, monkeys[i].numInspections)
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func part2() int {
	for r := 1; r <= 10000; r++ {
		for i := 0; i < 8; i++ {
			m := monkeys[i]
			for _, item := range m.items {
				worry := m.getWorryLevel(item)
				m.numInspections++
				m.throw(worry)
			}
			m.items = []int{}
			monkeys[i] = m
		}
	}

	inspections := []int{}
	for i := 0; i < 4; i++ {
		inspections = append(inspections, monkeys[i].numInspections)
	}
	sort.Ints(inspections)
	fmt.Println(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func main() {
	util.Run(part1, part2)
}
