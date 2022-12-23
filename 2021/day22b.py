from copy import copy
from dataclasses import dataclass

lines = [l.strip().split() for l in open("day22_input.txt", "r").readlines()]


@dataclass
class Cube:
    x1: int
    x2: int
    y1: int
    y2: int
    z1: int
    z2: int
    on: bool

    def __hash__(self):
        return hash((*self.bounds(), self.on))

    def volume(self):
        return (self.x2 - self.x1 + 1) * (self.y2 - self.y1 + 1) * (self.z2 - self.z1 + 1)

    def bounds(self):
        return self.x1, self.x2, self.y1, self.y2, self.z1, self.z2

    def calculate_intersect(self, other):
        x1, x2, y1, y2, z1, z2 = other.bounds()
        if x1 > self.x2 or x2 < self.x1 or y1 > self.y2 or y2 < self.y1 or z1 > self.z2 or z2 < self.z1:
            return None

        ix1 = max(x1, self.x1)
        ix2 = min(x2, self.x2)
        iy1 = max(y1, self.y1)
        iy2 = min(y2, self.y2)
        iz1 = max(z1, self.z1)
        iz2 = min(z2, self.z2)

        return Cube(ix1, ix2, iy1, iy2, iz1, iz2, not self.on)


cubes = set()

for d, cs in lines:
    on = d == "on"
    xs, ys, zs = cs.split(",")
    x1, x2 = list(map(int, xs.removeprefix("x=").split("..")))
    y1, y2 = list(map(int, ys.removeprefix("y=").split("..")))
    z1, z2 = list(map(int, zs.removeprefix("z=").split("..")))

    new_cube = Cube(x1, x2, y1, y2, z1, z2, on)
    for old_cube in copy(cubes):
        intersect = old_cube.calculate_intersect(new_cube)
        if intersect is not None:
            cubes.add(intersect)

    if new_cube.on:
        cubes.add(new_cube)

total = 0
for cube in cubes:
    total += cube.volume() * (1 if cube.on else -1)

print(total)
