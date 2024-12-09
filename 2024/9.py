import heapq
import math
from collections import defaultdict, deque

import util


def part1(input: str) -> int:
    disk_map = []
    free = deque()
    id = 0
    for i, size in enumerate(map(int, list(input.strip()))):
        if i % 2 == 0:
            disk_map.extend([id] * size)
            id += 1
        else:
            free.extend(range(len(disk_map), len(disk_map) + size))
            disk_map.extend(["."] * size)

    while free:
        c = disk_map.pop()
        if c == ".":
            continue

        idx = free.popleft()
        if idx >= len(disk_map):
            disk_map.append(c)
            break

        disk_map[idx] = c

    return sum(i * int(c) for i, c in enumerate(disk_map))


def part2(input: str) -> int:
    disk_map = []
    id = 0
    # map (length of free space) -> (starting indicies for free spaces of that length)
    space_size_indicies = defaultdict(list)
    file_info = {}
    spaces = []
    for i, size in enumerate(map(int, list(input.strip()))):
        if i % 2 == 0:
            file_info[id] = (len(disk_map), size)
            disk_map.extend([id] * size)
            id += 1
        else:
            heapq.heappush(space_size_indicies[size], len(disk_map))
            spaces.append((len(disk_map), size))
            disk_map.extend(["."] * size)

    max_space_size = max(space_size_indicies.keys())

    for id in range(id - 1, -1, -1):
        file_idx, file_size = file_info[id]

        space_idx, space_size = math.inf, None
        for size in range(file_size, max_space_size + 1):
            si = (space_size_indicies[size] or [math.inf])[0]
            if si < space_idx and si <= file_idx:
                space_idx = si
                space_size = size

        # could not find empty space to fit file
        if space_size is None:
            continue

        heapq.heappop(space_size_indicies[space_size])

        # remove existing file
        disk_map[file_idx : file_idx + file_size] = ["."] * file_size

        # put file in new location
        disk_map[space_idx : space_idx + file_size] = [id] * file_size

        # keep track of new empty space
        leftover_size = space_size - file_size
        heapq.heappush(space_size_indicies[leftover_size], space_idx + file_size)

    return sum(i * c for i, c in enumerate(disk_map) if c != ".")


if __name__ == "__main__":
    util.run(9, part1, part2)
