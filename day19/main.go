package main

import (
	"math"
	"regexp"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

type obsidianCost struct {
	ore, clay int
}

type geodeCost struct {
	ore, obsidian int
}

type blueprint struct {
	id           int
	oreCost      int
	clayCost     int
	obsidianCost obsidianCost
	geodeCost    geodeCost
}

var inputLines = util.FetchInput(19)

func parseBlueprints() []blueprint {
	blueprints := []blueprint{}
	re := regexp.MustCompile("[0-9]+")
	for _, line := range inputLines {
		numbers := re.FindAllString(line, -1)
		blueprints = append(blueprints, blueprint{
			util.ToInt(numbers[0]),
			util.ToInt(numbers[1]),
			util.ToInt(numbers[2]),
			obsidianCost{util.ToInt(numbers[3]), util.ToInt(numbers[4])},
			geodeCost{util.ToInt(numbers[5]), util.ToInt(numbers[6])},
		})
	}
	return blueprints
}

type state struct {
	minute                                             int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
	ore, clay, obsidian, geodes                        int
}

func bfs(bp blueprint, curr state, cache map[state]int) int {
	if curr.minute == 0 {
		return curr.geodes
	}

	if quality, ok := cache[curr]; ok {
		return quality
	}

	newOre := curr.ore + curr.oreRobots
	newClay := curr.clay + curr.clayRobots
	newObsidian := curr.obsidian + curr.obsidianRobots
	newGeodes := curr.geodes + curr.geodeRobots

	var maxGeodes = math.MinInt

	// build a geode robot
	if curr.ore >= bp.geodeCost.ore && curr.obsidian >= bp.geodeCost.obsidian {
		s := state{
			curr.minute - 1,
			curr.oreRobots,
			curr.clayRobots,
			curr.obsidianRobots,
			curr.geodeRobots + 1,
			newOre - bp.geodeCost.ore,
			newClay,
			newObsidian - bp.geodeCost.obsidian,
			newGeodes,
		}
		maxGeodes = util.Max(maxGeodes, bfs(bp, s, cache))

	} else {

		// build an obsidian robot
		if curr.ore >= bp.obsidianCost.ore && curr.clay >= bp.obsidianCost.clay {
			s := state{
				curr.minute - 1,
				curr.oreRobots,
				curr.clayRobots,
				curr.obsidianRobots + 1,
				curr.geodeRobots,
				newOre - bp.obsidianCost.ore,
				newClay - bp.obsidianCost.clay,
				newObsidian,
				newGeodes,
			}
			maxGeodes = util.Max(maxGeodes, bfs(bp, s, cache))
		} else {

			// build a clay robot
			if curr.ore >= bp.clayCost {
				s := state{
					curr.minute - 1,
					curr.oreRobots,
					curr.clayRobots + 1,
					curr.obsidianRobots,
					curr.geodeRobots,
					newOre - bp.clayCost,
					newClay,
					newObsidian,
					newGeodes,
				}
				maxGeodes = util.Max(maxGeodes, bfs(bp, s, cache))
			}

			// build an ore robot
			canBuildOreRobot := curr.ore >= bp.oreCost
			if canBuildOreRobot && curr.oreRobots < bp.clayCost {
				s := state{
					curr.minute - 1,
					curr.oreRobots + 1,
					curr.clayRobots,
					curr.obsidianRobots,
					curr.geodeRobots,
					newOre - bp.oreCost,
					newClay,
					newObsidian,
					newGeodes,
				}
				maxGeodes = util.Max(maxGeodes, bfs(bp, s, cache))
			}

			// don't build new robots
			s := state{
				curr.minute - 1,
				curr.oreRobots,
				curr.clayRobots,
				curr.obsidianRobots,
				curr.geodeRobots,
				newOre,
				newClay,
				newObsidian,
				newGeodes,
			}
			maxGeodes = util.Max(maxGeodes, bfs(bp, s, cache))
		}
	}

	cache[curr] = maxGeodes

	return cache[curr]
}

func part1() int {
	blueprints := parseBlueprints()
	totalQualityLevel := 0
	for _, bp := range blueprints {
		totalQualityLevel += bfs(bp, state{24, 1, 0, 0, 0, 0, 0, 0, 0}, map[state]int{}) * bp.id
	}
	return totalQualityLevel
}

func part2() int {
	blueprints := parseBlueprints()
	geodeProduct := 1
	for _, bp := range blueprints[:3] {
		maxGeodes := bfs(bp, state{32, 1, 0, 0, 0, 0, 0, 0, 0}, map[state]int{})
		geodeProduct *= maxGeodes
	}
	return geodeProduct
}

func main() {
	util.Run(part1, part2)
}
