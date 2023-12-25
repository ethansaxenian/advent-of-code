from util.util import fetch_input
import networkx as nx

puzzle_input = fetch_input(25)

G = nx.Graph()

for line in puzzle_input:
    n, edges = line.split(":")
    for e in edges.strip().split():
        G.add_edge(n, e)
        G.add_edge(e, n)


cutset = nx.minimum_edge_cut(G)
assert len(cutset) == 3
for edge in cutset:
    G.remove_edge(*edge)

x, y = nx.connected_components(G)
print(len(x) * len(y))
