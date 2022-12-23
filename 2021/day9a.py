grid = []

with open("day9_input.txt", "r") as file:
    for line in file:
        grid.append(line.strip())

risk_level = 0

for y in range(len(grid)):
    for x in range(len(grid[0])):
        neighbors = []
        if y > 0:
            neighbors.append(grid[y-1][x])
        if y < len(grid) - 1:
            neighbors.append(grid[y+1][x])
        if x > 0:
            neighbors.append(grid[y][x-1])
        if x < len(grid[0]) - 1:
            neighbors.append(grid[y][x+1])

        if all(int(grid[y][x]) < int(n) for n in neighbors):
            risk_level += int(grid[y][x]) + 1

print(risk_level)
