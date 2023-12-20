package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/ethansaxenian/advent-of-code/2023/util"
)

var puzzleInput = util.FetchInput(20)

const (
	LOW  = false
	HIGH = true
)

var conj string
var prevs = map[string]int{}
var done bool = false
var i int = 0

type QueueItem struct {
	pulse  bool
	id, in string
}

var queue = []QueueItem{}

func queuePulse(pulse bool, id, in string) {
	queue = append(queue, QueueItem{pulse, id, in})
}

type Module interface {
	getDestinations() []string
	receivePulse(pulse bool, in string)
	String() string
	ID() string
}

type FlipFlop struct {
	id           string
	on           bool
	destinations []string
}

func (f *FlipFlop) receivePulse(pulse bool, in string) {
	if pulse == HIGH {
		return
	}

	f.on = !f.on
	for _, d := range f.destinations {
		queuePulse(f.on, d, f.id)
	}
}

func (f FlipFlop) getDestinations() []string {
	return f.destinations
}

func (f FlipFlop) String() string {
	return fmt.Sprintf("(%v, %v, %v)", f.id, f.on, f.destinations)
}

func (f FlipFlop) ID() string {
	return f.id
}

type Conjunction struct {
	id           string
	memory       map[string]bool
	destinations []string
}

func (c *Conjunction) receivePulse(pulse bool, in string) {
	c.memory[in] = pulse

	if c.id == conj && pulse == HIGH {
		prevs[in] = i

		all := true
		for _, v := range prevs {
			if v == 0 {
				all = false
				break
			}
		}

		if all {
			done = true
			ints := []int{}
			for _, v := range prevs {
				ints = append(ints, v)
			}
			fmt.Println(util.LCM(ints[0], ints[1], ints[2:]...))
			os.Exit(0)
		}
	}

	nextPulse := LOW

	if len(c.memory) == 0 {
		nextPulse = HIGH
	}

	for _, v := range c.memory {
		if v == LOW {
			nextPulse = HIGH
			break
		}
	}

	for _, d := range c.destinations {
		queuePulse(nextPulse, d, c.id)
	}
}

func (c Conjunction) getDestinations() []string {
	return c.destinations
}

func (c Conjunction) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.id, c.memory, c.destinations)
}

func (c Conjunction) ID() string {
	return c.id
}

func parseInput(input []string) ([]string, map[string]Module) {
	var broadcaster []string
	var modules = map[string]Module{}

	conjunctions := map[string]*Conjunction{}

	for _, line := range input {
		parts := strings.Fields(line)
		prefix := parts[0][0]
		id := strings.TrimLeft(parts[0], "&%")
		destinations := util.Map[string, string](parts[2:], func(s string) string {
			return strings.TrimRight(s, ",")
		})

		if id == "broadcaster" {
			broadcaster = destinations
		} else if prefix == '%' {
			modules[id] = &FlipFlop{id, LOW, destinations}
		} else if prefix == '&' {
			conj := &Conjunction{id, map[string]bool{}, destinations}
			conjunctions[id] = conj
			modules[id] = conj
		}
	}

	for id, module := range modules {
		for _, d := range module.getDestinations() {
			if reflect.TypeOf(modules[d]) == reflect.TypeOf(&Conjunction{}) {
				conjunctions[d].memory[id] = LOW
			}
		}
	}

	return broadcaster, modules
}

func pressButton(broadcaster []string, modules map[string]Module) (int, int) {
	var low, high int
	low++

	for _, d := range broadcaster {
		queuePulse(LOW, d, "broadcaster")
	}

	for len(queue) > 0 {

		if done {
			return low, high
		}

		item := queue[0]
		queue = queue[1:]

		if item.pulse == LOW {
			low++
		} else {
			high++
		}

		if mod, ok := modules[item.id]; ok {
			mod.receivePulse(item.pulse, item.in)
		}
	}

	return low, high
}

func getRelevantInputs(modules map[string]Module) (string, map[string]int) {
	var last string
	for _, mod := range modules {
		for _, d := range mod.getDestinations() {
			if d == "rx" {
				last = mod.ID()
			}
		}
	}

	prevs := map[string]int{}

	for _, mod := range modules {
		for _, d := range mod.getDestinations() {
			if d == last {
				prevs[mod.ID()] = 0
			}
		}
	}

	return last, prevs
}

func main() {
	broadcaster, modules := parseInput(puzzleInput)

	var low, high int
	for i := 0; i < 1000; i++ {
		l, h := pressButton(broadcaster, modules)
		low += l
		high += h
	}

	fmt.Println(low * high)

	broadcaster, modules = parseInput(puzzleInput)

	conj, prevs = getRelevantInputs(modules)

	i = 1
	done = false

	for {
		pressButton(broadcaster, modules)
		i++
	}

}
