lines = [l.strip().split() for l in open("day22_input.txt", "r").readlines()]

grid = [[[False for _ in range(-50, 51)] for _ in range(-50, 51)] for _ in range(-50, 51)]


for d, cs in lines:
    b = {"on": True, "off": False}[d]
    xs, ys, zs = cs.split(",")
    x1, x2 = map(int, xs.removeprefix("x=").split(".."))
    y1, y2 = map(int, ys.removeprefix("y=").split(".."))
    z1, z2 = map(int, zs.removeprefix("z=").split(".."))
    if any(abs(i) > 50 for i in [x1, x2, y1, y2, z1, z2]):
        continue
    for x in range(x1 + 50, x2 + 51):
        for y in range(y1 + 50, y2 + 51):
            for z in range(z1 + 50, z2 + 51):
                grid[x][y][z] = b

c = 0
for x in range(len(grid)):
    for y in range(len(grid[0])):
        for z in range(len(grid[0][0])):
            if grid[x][y][z]:
                c += 1

print(c)
