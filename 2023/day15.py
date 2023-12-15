from collections import defaultdict
from util import util
import re

puzzle_input = util.fetch_input(15)[0]


def hash(s):
    val = 0
    for c in s:
        val += ord(c)
        val *= 17
        val %= 256
    return val


total = 0
for line in puzzle_input.split(","):
    total += hash(line)

print(total)

boxes = defaultdict(list)
lenses = {}


for line in puzzle_input.split(","):
    label, focal_length = re.split("=|-", line)
    box = hash(label)
    if focal_length:
        lenses[label] = (focal_length, box)
        if label not in boxes[box]:
            boxes[box].append(label)
    else:
        if label in boxes[box]:
            boxes[box].remove(label)

power = 0
for box, _lenses in boxes.items():
    for i, lens in enumerate(_lenses, 1):
        power += (1 + box) * i * int(lenses[lens][0])

print(power)
