line = open("day17_input.txt", "r").readline().strip()
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
            return steps
    return None


paths = []
for x in range(0, xmax):
    for y in range(0, 300):
        path = go(x, y)
        if path is not None:
            paths.append(path)

max_height = 0
for p in paths:
    m = max(p, key=lambda pair: pair[1])[1]
    if m > max_height:
        max_height = m

print(max_height)
