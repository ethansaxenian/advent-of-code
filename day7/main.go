package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

type node struct {
	name     string
	parent   *node
	children []*node
	dir      bool
	size     int
}

func getNodeFromParts(parts []string, me *node) *node {
	if parts[0] == "dir" {
		newNode := &node{name: parts[1], dir: true, size: 0, parent: me}
		dirs[newNode] = 0
		return newNode
	} else {
		return &node{name: parts[1], dir: false, size: util.ToInt(parts[0]), parent: me}
	}
}

func getSize(root *node) int {
	if len(root.children) == 0 {
		return (*root).size
	}

	sizeSum := 0
	for _, n := range (*root).children {
		sizeSum += getSize(n)
	}

	return sizeSum
}

func buildFilesystem() *node {
	root := node{name: "/", children: []*node{}, parent: nil, dir: true, size: 0}
	me := &root
	i := 0
	for i < len(inputLines) {
		parts := strings.Split(inputLines[i], " ")
		if parts[0] == "$" {
			if parts[1] == "ls" {
				i += 1
				for i < len(inputLines) {
					fileParts := strings.Split(inputLines[i], " ")
					if fileParts[0] == "$" {
						break
					}

					me.children = append(me.children, getNodeFromParts(fileParts, me))
					i += 1
				}

			} else if parts[1] == "cd" {
				if parts[2] == ".." {
					me = me.parent

				} else {
					for _, n := range me.children {
						if n.name == parts[2] {
							me = n
							break
						}
					}
				}

				i += 1
			}
		} else {
			fmt.Println("SOMETHING WENT REALLY WRONG")
		}

	}

	return &root
}

var inputLines = util.FetchInput(7)[1:]

var dirs = map[*node]int{}

func part1() int {
	buildFilesystem()

	for k := range dirs {
		dirs[k] += getSize(k)
	}

	total := 0
	for _, v := range dirs {
		if v <= 100000 {
			total += v
		}
	}

	return total
}

func part2() int {
	totalAvailableSpace := 70000000
	spaceNeeded := 30000000

	root := buildFilesystem()
	currentUnusedSpace := totalAvailableSpace - getSize(root)
	dirSizeNeeded := spaceNeeded - currentUnusedSpace

	for k := range dirs {
		dirs[k] += getSize(k)
	}

	sizes := []int{}
	for _, v := range dirs {
		sizes = append(sizes, v)
	}
	sort.Ints(sizes)

	for _, s := range sizes {
		if s >= dirSizeNeeded {
			return s
		}
	}

	return 0
}

func main() {
	util.Run(part1, part2)
}
