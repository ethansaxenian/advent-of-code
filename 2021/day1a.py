with open("day1_input.txt", "r") as file:
    n = 0
    prev = 0
    for line in file:
        if int(line) > prev:
            n += 1
        prev = int(line)

print(n-1)
