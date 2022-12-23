from collections import defaultdict

graph = defaultdict(list)

with open("day12_input.txt", "r") as file:
    for line in file:
        src, dest = line.strip().split("-")
        graph[src].append(dest)
        graph[dest].append(src)

paths = []


def dfs(G, vertex, visited, paths):
    if vertex.islower():
        visited.append(vertex)
    if vertex == "end":
        paths.append(visited[:])
        if vertex.islower():
            visited.pop()
        return
    for v in G[vertex]:
        if v not in visited:
            dfs(G, v, visited, paths)
    if vertex.islower():
        visited.pop()


dfs(graph, "start", [], paths)

print(len(paths))
