lines = [l.strip() for l in open("day21_input.txt", "r").readlines()]

p1 = int(lines[0][-1])
p2 = int(lines[1][-1])

rolls = 0
d = 0
p1_score = 0
p2_score = 0

p = True

while p1_score < 1000 and p2_score < 1000:
    move = 0
    for _ in range(3):
        d += 1
        if d == 1001:
            d = 1
        move += d

    if p:
        p1 = (p1 + move) % 10
        if p1 == 0:
            p1 = 10
        p1_score += p1
    else:
        p2 = (p2 + move) % 10
        if p2 == 0:
            p2 = 10
        p2_score += p2
    rolls += 3

    p = not p

print(p1_score, p2_score, rolls)
print(min(p1_score, p2_score) * rolls)
