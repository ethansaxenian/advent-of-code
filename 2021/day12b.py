from collections import defaultdict

graph = defaultdict(list)

with open("day12_input.txt", "r") as file:
    for line in file:
        src, dest = line.strip().split("-")
        graph[src].append(dest)
        graph[dest].append(src)

paths = []


def dfs(G, vertex, visited, paths, visited_small_cave_twice):
    visited.append(vertex)
    if vertex.islower() and visited.count(vertex) == 2:
        visited_small_cave_twice = True
    if vertex == "end":
        paths.append(visited[:])
        print(",".join(visited))
        visited.pop()
        return
    for v in G[vertex]:
        if v.isupper():
            dfs(G, v, visited, paths, visited_small_cave_twice)
        elif v in ("start", "end"):
            if v not in visited:
                dfs(G, v, visited, paths, visited_small_cave_twice)
        elif visited_small_cave_twice:
            if v not in visited:
                dfs(G, v, visited, paths, visited_small_cave_twice)
        elif visited.count(v) < 2:
            dfs(G, v, visited, paths, visited_small_cave_twice)
    visited.pop()


dfs(graph, "start", [], paths, False)

# print(paths)
print(len(paths))
for i, p in enumerate(paths):
    if p in paths[i+1:]:
        print(p)
