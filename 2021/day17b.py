line = open("day17_input.txt", "r").readline().strip()
# line = "target area: x=20..30, y=-10..-5"
xbound, ybound = line.removeprefix("target area: ").split(", ")
xmin, xmax = map(int, xbound.removeprefix("x=").split(".."))
ymin, ymax = map(int, ybound.removeprefix("y=").split(".."))


def in_target_area(x, y):
    return xmin <= x <= xmax and ymin <= y <= ymax


def past_target_area(x, y, yv):
    if x > xmax:
        return True
    if y < ymin and yv < 0:
        return True
    return False


def go(xinit, yinit):
    steps = [(0, 0)]
    x, y = xinit, yinit
    curr_x = curr_y = 0
    while not past_target_area(curr_x, curr_y, y):
        curr_x += x
        curr_y += y
        x = x + (-1 if x > 0 else (1 if x < 0 else 0))
        y -= 1
        steps.append((curr_x, curr_y))
        if in_target_area(curr_x, curr_y):
            return xinit, yinit
    return None


paths = []
for x in range(1, xmax + 1):
    for y in range(ymin - 1, xmax):
        if past_target_area(x, y, y-1):
            continue
        path = go(x, y)
        if path is not None:
            paths.append(path)

print(len(paths))
