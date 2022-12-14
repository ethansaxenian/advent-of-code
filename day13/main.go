package main

import (
	"encoding/json"
	"errors"
	"reflect"
	"sort"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

var inputLines = util.FetchInput(13)

func buildLists() [][]any {
	lists := [][]any{}
	for _, line := range inputLines {
		if line == "" {
			continue
		}
		var list []any
		json.Unmarshal([]byte(line), &list)
		lists = append(lists, list)
	}
	return lists
}

func sameType(a, b any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func compare(left, right []any) (bool, error) {
	for i := 0; i < len(left) && i < len(right); i++ {
		l := left[i]
		r := right[i]

		switch l.(type) {
		case float64:
			if sameType(l, r) {
				if l.(float64) < r.(float64) {
					return true, nil
				} else if l.(float64) > r.(float64) {
					return false, nil
				}
			} else {
				if correct, err := compare([]any{l}, r.([]any)); err == nil {
					return correct, nil
				}
			}
		case []any:
			if sameType(l, r) {
				if correct, err := compare(l.([]any), r.([]any)); err == nil {
					return correct, nil
				}
			} else {
				if correct, err := compare(l.([]any), []any{r}); err == nil {
					return correct, nil
				}
			}
		default:
			panic("help")
		}
	}

	if len(left) < len(right) {
		return true, nil
	}
	if len(left) > len(right) {
		return false, nil
	}

	return false, errors.New("end of lists")
}

func part1() int {
	lists := buildLists()
	s := 0
	p := 1
	for i := 0; i < len(lists); i += 2 {
		if correct, err := compare(lists[i], lists[i+1]); correct && err == nil {
			s += p
		}
		p++
	}
	return s
}

func part2() int {
	lists := buildLists()
	var a, b []any
	json.Unmarshal([]byte("[[2]]"), &a)
	json.Unmarshal([]byte("[[6]]"), &b)
	lists = append(lists, a, b)

	sort.Slice(lists, func(i, j int) bool { correct, _ := compare(lists[i], lists[j]); return correct })

	key := 1
	for i, list := range lists {
		if _, err := compare(list, a); err != nil {
			key *= i + 1
		}
		if _, err := compare(list, b); err != nil {
			key *= i + 1
		}
	}
	return key
}

func main() {
	util.Run(part1, part2)
}
