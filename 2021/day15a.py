from dataclasses import dataclass

from numpy import infty


@dataclass
class Vertex:
    label: str
    neighbors: dict
    prev = None
    visited: bool = False
    distance: int = infty
    src: bool = False
    dest: bool = False


def get_neighbors(grid, x, y):
    neighbors = []
    if y > 0:
        neighbors.append(grid[y - 1][x])
    if y < len(grid) - 1:
        neighbors.append(grid[y + 1][x])
    if x > 0:
        neighbors.append(grid[y][x - 1])
    if x < len(grid[0]) - 1:
        neighbors.append(grid[y][x + 1])

    return neighbors


def length(u, v):
    if v.label in u.neighbors:
        return u.neighbors[v.label]
    return infty


def get_min_dist(graph):
    return min(graph, key=lambda v: v.distance)


with open("day15_input.txt", "r") as file:
    grid = [line.strip() for line in file.readlines()]
    graph = [[Vertex(label=f"{x}X{y}Y", neighbors={}) for x in range(len(grid))] for y in range(len(grid))]
    for y in range(len(graph)):
        for x in range(len(graph)):
            v = graph[y][x]
            if x == y == 0:
                v.src = True
                source = v
                source.distance = 0
            if x == y == len(graph) - 1:
                v.dest = True
                target = v

            if y > 0:
                v.neighbors[graph[y - 1][x].label] = int(grid[y - 1][x])
            if y < len(grid) - 1:
                v.neighbors[graph[y + 1][x].label] = int(grid[y + 1][x])
            if x > 0:
                v.neighbors[graph[y][x - 1].label] = int(grid[y][x - 1])
            if x < len(grid[0]) - 1:
                v.neighbors[graph[y][x + 1].label] = int(grid[y][x + 1])

    graph = [v for row in graph for v in row]

lookup = {v.label: v for v in graph}


while True:
    print(len(graph))
    u = get_min_dist(graph)
    u.visited = True
    graph.remove(u)
    if u.dest:
        break
    for n in u.neighbors:
        v = lookup[n]
        if not v.visited:
            dist = u.distance + length(u, v)
            if dist < v.distance:
                v.distance = dist
                v.prev = u

path = []
u = target
if u.prev is not None or u.src:
    while u:
        path.append(u)
        u = u.prev

path.reverse()
risk = 0
for i in range(len(path)):
    try:
        risk += length(path[i], path[i+1])
    except IndexError:
        continue

for p in path:
    print(p.label, p.distance)
print(risk)
