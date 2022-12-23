import heapq
from dataclasses import dataclass

from numpy import infty


@dataclass
class Vertex:
    label: str
    neighbors: dict
    prev = None
    visited: bool = False
    distance: int = infty

    def __lt__(self, other):
        return self.distance < other.distance


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


grid = [list(map(int, line.strip())) for line in open("day15_input.txt", "r").readlines()]


def wrap(x, i):
    if x + i > 9:
        return (x + i) % 9
    return x + i


for line in grid:
    orig = line[:]
    for i in range(1, 5):
        line.extend(list(map(lambda x: wrap(x, i), orig)))

grid_copy = grid[:]
for i in range(1, 5):
    for line in grid_copy:
        orig = line[:]
        grid.append(list(map(lambda x: wrap(x, i), orig)))

graph = [[Vertex(label=f"{x}X{y}Y", neighbors={}) for x in range(len(grid))] for y in range(len(grid))]
for y in range(len(graph)):
    for x in range(len(graph)):
        v = graph[y][x]
        if y > 0:
            v.neighbors[graph[y - 1][x].label] = int(grid[y - 1][x])
        if y < len(grid) - 1:
            v.neighbors[graph[y + 1][x].label] = int(grid[y + 1][x])
        if x > 0:
            v.neighbors[graph[y][x - 1].label] = int(grid[y][x - 1])
        if x < len(grid[0]) - 1:
            v.neighbors[graph[y][x + 1].label] = int(grid[y][x + 1])

source = graph[0][0]
source.distance = 0

target = graph[len(graph) - 1][len(graph) - 1]

graph = [v for row in graph for v in row]

lookup = {v.label: v for v in graph}

h = []

heapq.heappush(h, source)

while True:
    u = heapq.heappop(h)
    u.visited = True
    if u == target:
        break
    for n in u.neighbors:
        v = lookup[n]
        if v == u.prev:
            continue
        if not v.visited:
            dist = u.distance + length(u, v)
            if dist < v.distance:
                v.distance = dist
                v.prev = u
                heapq.heappush(h, v)

path = []
u = target
if u.prev is not None or u == source:
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


print(risk)
