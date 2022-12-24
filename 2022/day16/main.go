package main

import (
	"regexp"
	"sort"

	"github.com/ethansaxenian/advent-of-code/2022/set"
	"github.com/ethansaxenian/advent-of-code/2022/util"
)

type node struct {
	id       string
	flowRate int
	edges    []string
}

type state struct {
	n      string
	minute int
	open   string
}

var inputLines = util.FetchInput(16)

func buildGraph() map[string]node {
	g := map[string]node{}
	re := regexp.MustCompile("[0-9a-zA-Z]+")
	for _, line := range inputLines {
		parts := re.FindAllString(line, -1)
		n := node{parts[1], util.ToInt(parts[5]), parts[10:]}
		g[parts[1]] = n
	}

	return g
}

func copy(s set.Set[string]) set.Set[string] {
	newSet := set.NewEmptySet[string]()
	for _, v := range s.Items() {
		newSet.Add(v)
	}
	return newSet
}

func hash(s set.Set[string]) string {
	ids := s.Items()
	sort.Strings(ids)
	id := ""
	for _, s := range ids {
		id += s + " "
	}
	return id
}

var memo map[state]int = map[state]int{}

func dfs(n string, minute int, flowRate int, open set.Set[string], graph map[string]node) int {
	if minute == 0 {
		return 0
	}

	h := hash(open)
	s := state{n, minute, h}

	if p, ok := memo[s]; ok {
		return p
	}

	flow := 0

	if _, ok := open[n]; graph[n].flowRate > 0 && !ok {
		s := copy(open)
		s.Add(n)
		flow = dfs(n, minute-1, flowRate+graph[n].flowRate, s, graph)
	}

	neighbors := graph[n].edges
	for _, v := range neighbors {
		flow = util.Max(flow, dfs(v, minute-1, flowRate, copy(open), graph))
	}

	flow += flowRate
	memo[s] = flow

	return memo[s]
}

func part1() int {
	graph := buildGraph()
	maxPressure := dfs("AA", 30, 0, set.NewEmptySet[string](), graph)
	return maxPressure
}

type state2 struct {
	me       string
	elephant string
	minute   int
	open     string
}

var memo2 map[state2]int = map[state2]int{}

func dfs2(me string, elephant string, minute int, flowRate int, open set.Set[string], graph map[string]node) int {
	if minute == 0 {
		return 0
	}

	h := hash(open)
	s := state2{me, elephant, minute, h}

	if p, ok := memo2[s]; ok {
		return p
	}

	me_open := false
	if _, ok := open[me]; graph[me].flowRate > 0 && !ok {
		me_open = true
	}

	elephant_open := false
	if _, ok := open[elephant]; graph[elephant].flowRate > 0 && !ok && me != elephant {
		elephant_open = true
	}

	flow := 0

	if me_open && elephant_open {
		s := copy(open)
		s.Add(me)
		s.Add(elephant)
		flow = dfs2(me, elephant, minute-1, flowRate+graph[me].flowRate+graph[elephant].flowRate, s, graph)
	} else if me_open && !elephant_open {
		s := copy(open)
		s.Add(me)
		for _, e := range graph[elephant].edges {
			flow = util.Max(flow, dfs2(me, e, minute-1, flowRate+graph[me].flowRate, s, graph))
		}
	} else if !me_open && elephant_open {
		s := copy(open)
		s.Add(elephant)
		for _, m := range graph[me].edges {
			flow = util.Max(flow, dfs2(m, elephant, minute-1, flowRate+graph[elephant].flowRate, s, graph))
		}
	} else {
		for _, m := range graph[me].edges {
			for _, e := range graph[elephant].edges {
				flow = util.Max(flow, dfs2(m, e, minute-1, flowRate, copy(open), graph))
			}
		}
	}

	flow += flowRate
	memo2[s] = flow

	return memo2[s]
}

func part2() int {
	graph := buildGraph()
	maxPressure := dfs2("AA", "AA", 26, 0, set.NewEmptySet[string](), graph)
	return maxPressure
}

func main() {
	util.Run(part1, part2)
}
