package main

import (
	"container/list"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

var inputLines = util.FetchInput(20)

func getIndex(l *list.List, elem *list.Element) int {
	curr := l.Front()
	i := 0
	for curr != nil {
		if curr == elem {
			return i
		}
		curr = curr.Next()
		i++
	}
	return -1
}

func findAtIndex(l *list.List, index int) *list.Element {
	curr := l.Front()
	for i := 0; i < index; i++ {
		curr = curr.Next()
	}
	return curr
}

func mix(file *list.List, initFile []*list.Element) {
	fileSize := file.Len()

	for _, elem := range initFile {

		i := getIndex(file, elem)
		v := elem.Value.(int)

		newIndex := (i + v) % (fileSize - 1)
		if newIndex < 0 {
			newIndex += (fileSize - 1)
		}

		if newIndex == 0 {
			file.MoveToBack(elem)
		} else if newIndex == fileSize-1 {
			file.MoveToFront(elem)
		} else {

			neighbor := findAtIndex(file, newIndex)

			if i <= newIndex {
				file.MoveAfter(elem, neighbor)
			} else {
				file.MoveBefore(elem, neighbor)
			}
		}
	}
}

func getGroveCoordinates(file *list.List) int {
	fileSize := file.Len()

	curr := file.Front()
	i := 0
	for curr != nil {
		if curr.Value == 0 {
			break
		}
		curr = curr.Next()
		i++
	}

	v1 := findAtIndex(file, (i+1000)%fileSize).Value.(int)
	v2 := findAtIndex(file, (i+2000)%fileSize).Value.(int)
	v3 := findAtIndex(file, (i+3000)%fileSize).Value.(int)
	return v1 + v2 + v3
}

func part1() int {
	file := list.New()
	initFile := []*list.Element{}
	for _, line := range inputLines {
		elem := file.PushBack(util.ToInt(line))
		initFile = append(initFile, elem)
	}

	mix(file, initFile)

	return getGroveCoordinates(file)
}

const decryptionKey = 811589153

func part2() int {
	file := list.New()
	initFile := []*list.Element{}
	for _, line := range inputLines {
		elem := file.PushBack(util.ToInt(line) * decryptionKey)
		initFile = append(initFile, elem)
	}

	for i := 0; i < 10; i++ {
		mix(file, initFile)
	}

	return getGroveCoordinates(file)
}

func main() {
	util.Run(part1, part2)
}
