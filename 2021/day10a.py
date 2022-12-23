strings = [line.strip() for line in open("day10_input.txt", "r")]

pairs = {
    "(": ")",
    "[": "]",
    "{": "}",
    "<": ">"
}

points = {
    ")": 3,
    "]": 57,
    "}": 1197,
    ">": 25137
}

errors = []
for s in strings:
    look_for = []
    for c in s:
        if c in pairs.keys():
            look_for.append(c)
        else:
            if pairs[look_for[-1]] == c:
                look_for.pop()
            else:
                errors.append(c)
                break
print(errors)
print(sum(points[e] for e in errors))
