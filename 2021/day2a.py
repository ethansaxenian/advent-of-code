with open("day2_input.txt", "r") as file:
    horizontal = 0
    vertical = 0
    for line in file:
        if line.startswith("forward"):
            horizontal += int(line[8:])
        if line.startswith("down"):
            vertical += int(line[5:])
        if line.startswith("up"):
            vertical -= int(line[3:])

print(horizontal, vertical, horizontal * vertical)
