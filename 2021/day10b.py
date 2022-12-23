from statistics import median

strings = [line.strip() for line in open("day10_input.txt", "r")]

pairs = {
    "(": ")",
    "[": "]",
    "{": "}",
    "<": ">"
}

points = {
    ")": 1,
    "]": 2,
    "}": 3,
    ">": 4
}

all_scores = []

for s in strings:
    total_score = 0
    look_for = []
    corrupted = False
    for c in s:
        if c in pairs.keys():
            look_for.append(c)
        else:
            if pairs[look_for[-1]] == c:
                look_for.pop()
            else:
                corrupted = True
                break

    if corrupted:
        continue

    for p in reversed(look_for):
        total_score *= 5
        total_score += points[pairs[p]]

    all_scores.append(total_score)

print(median(all_scores))
