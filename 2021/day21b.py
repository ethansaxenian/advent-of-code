from collections import defaultdict

lines = [l.strip() for l in open("day21_input.txt", "r").readlines()]

p1 = int(lines[0][-1])
p2 = int(lines[1][-1])


def advance(states, p1_turn):
    x = 0
    new_states = defaultdict(int)
    for (p1, p2, s1, s2), num_occurances in states.items():
        for a in [1, 2, 3]:
            for b in [1, 2, 3]:
                for c in [1, 2, 3]:
                    if p1_turn:
                        new_p1 = (p1+a+b+c - 1) % 10 + 1
                        if s1+new_p1 >= 21:
                            x += num_occurances
                        else:
                            new_states[(new_p1, p2, s1+new_p1, s2)] += num_occurances
                    else:
                        new_p2 = (p2+a+b+c - 1) % 10 + 1
                        if s2+new_p2 >= 21:
                            x += num_occurances
                        else:
                            new_states[(p1, new_p2, s1, s2+new_p2)] += num_occurances
    return new_states, x


game_states = defaultdict(int)
game_states[(p1, p2, 0, 0)] += 1

p1_w = 0
p2_w = 0
p1_turn = True
while True:
    game_states, x = advance(game_states, p1_turn)
    if p1_turn:
        p1_w += x
    else:
        p2_w += x
    if len(game_states) == 0:
        break
    p1_turn = not p1_turn

print(max(p1_w, p2_w))
