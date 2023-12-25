from util.util import fetch_input

puzzle_input = fetch_input(24)

# puzzle_input = [
#     "19, 13, 30 @ -2,  1, -2",
#     "18, 19, 22 @ -1, -1, -2",
#     "20, 25, 34 @ -2, -2, -4",
#     "12, 31, 28 @ -1, -2, -1",
#     "20, 19, 15 @  1, -5, -3",
# ]

LOWER_BOUND = 200000000000000
UPPER_BOUND = 400000000000000

hailstones = [
    tuple(int(c.strip(",")) for c in line.split() if c != "@") for line in puzzle_input
]

n = 0

for i, (x1, y1, _, dx1, dy1, _) in enumerate(hailstones[:-1]):
    m1 = dy1 / dx1
    b1 = y1 - (m1 * x1)
    for x2, y2, _, dx2, dy2, _ in hailstones[i + 1 :]:
        m2 = dy2 / dx2
        b2 = y2 - (m2 * x2)
        # m1x + b1 = m2x + b2
        # m1x - m2x = b2 - b1
        # (m1 - m2)x = b2 - b1
        try:
            x = (b2 - b1) / (m1 - m2)
        except Exception:
            continue
        else:
            y = m1 * x + b1
            if LOWER_BOUND <= x <= UPPER_BOUND and LOWER_BOUND <= y <= UPPER_BOUND:
                # d = rt
                # x - x1 = dx1 * t
                t1 = (x - x1) / dx1
                t2 = (x - x2) / dx2
                if t1 > 0 and t2 > 0:
                    n += 1

print(n)

for x0, y0, z0, dx0, dy0, dz0 in hailstones[:4]:
    # x + dx * t = x0 + dx0 * t
    # (x - x0) / (dx0 - dx) = t
    # y + dy * t = y0 + dy0 * t
    # (y - y0) / (dy0 - dy) = t
    # z + dz * t = z0 + dz0 * t
    # (z- z0) / (dz0 - dz) = t
    print(f"(x-{x0})/({dx0}-t)=(y-{y0})/({dy0}-u)=(z-{z0})/({dz0}-v)")

t = 44
u = 305
v = 75
x = 234382970331570
y = 100887864960615
z = 231102671115832

print(x + y + z)
