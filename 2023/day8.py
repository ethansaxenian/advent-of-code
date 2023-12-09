from util.util import fetch_input
import math
import re

input = fetch_input(8)

# input = [
#     "RL",
#     "",
#     "AAA = (BBB, CCC)",
#     "BBB = (DDD, EEE)",
#     "CCC = (ZZZ, GGG)",
#     "DDD = (DDD, DDD)",
#     "EEE = (EEE, EEE)",
#     "GGG = (GGG, GGG)",
#     "ZZZ = (ZZZ, ZZZ)",
# ]

# input = [
#     "LLR",
#     "",
#     "AAA = (BBB, BBB)",
#     "BBB = (AAA, ZZZ)",
#     "ZZZ = (ZZZ, ZZZ)",
# ]

# input = [
#     "LR",
#     "",
#     "11A = (11B, XXX)",
#     "11B = (XXX, 11Z)",
#     "11Z = (11B, XXX)",
#     "22A = (22B, XXX)",
#     "22B = (22C, 22C)",
#     "22C = (22Z, 22Z)",
#     "22Z = (22B, 22B)",
#     "XXX = (XXX, XXX)",
# ]

instructions = input[0].strip()

nodes = {
    a: [b, c]
    for line in input[2:]
    for a, b, c in re.findall(r"([0-9A-Z]+) = \(([0-9A-Z]+), ([0-9A-Z]+)\)", line)
}


def part1():
    node = "AAA"
    ins = 0
    steps = 0
    while node != "ZZZ":
        if instructions[ins] == "L":
            node = nodes[node][0]
        elif instructions[ins] == "R":
            node = nodes[node][1]
        steps += 1
        ins = (ins + 1) % len(instructions)

    print(steps)


def part2():
    curr_nodes = [node for node in nodes if node[-1] == "A"]
    steps = [0] * len(curr_nodes)
    for i in range(len(curr_nodes)):
        ins = 0
        while curr_nodes[i][-1] != "Z":
            if instructions[ins] == "L":
                curr_nodes[i] = nodes[curr_nodes[i]][0]
            elif instructions[ins] == "R":
                curr_nodes[i] = nodes[curr_nodes[i]][1]
            steps[i] += 1
            ins = (ins + 1) % len(instructions)

    print(math.lcm(*steps))


part1()
part2()
