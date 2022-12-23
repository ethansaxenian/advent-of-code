with open("day2_input.txt", "r") as file:
    pos = 0
    depth = 0
    aim = 0
    for line in file:
        if line.startswith("forward"):
            pos += int(line[8:])
            depth += aim * int(line[8:])
        if line.startswith("down"):
            aim += int(line[5:])
        if line.startswith("up"):
            aim -= int(line[3:])

print(pos, depth, pos * depth)
