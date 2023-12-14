from util.util import fetch_input

input_lines = fetch_input(13)

patterns = []

next_pattern = []
for line in input_lines:
    if line == "":
        patterns.append(next_pattern)
        next_pattern = []
        continue
    next_pattern.append(list(line.strip()))

patterns.append(next_pattern)


def check_vertical(pattern, ignore=-1):
    for c in range(len(pattern[0]) - 1):
        is_reflection_point = True
        for dc in range(c + 1):
            if c + 1 + dc >= len(pattern[0]) or c - dc < 0:
                break
            for r in range(len(pattern)):
                if pattern[r][c - dc] != pattern[r][c + 1 + dc]:
                    is_reflection_point = False
                    break
            if not is_reflection_point:
                break

        if is_reflection_point and c + 1 != ignore:
            return c + 1

    return 0


def check_horizontal(pattern, ignore=-1):
    for r in range(len(pattern) - 1):
        is_reflection_point = True
        for dr in range(r + 1):
            if r + 1 + dr >= len(pattern) or r - dr < 0:
                break
            for c in range(len(pattern[0])):
                if pattern[r - dr][c] != pattern[r + 1 + dr][c]:
                    is_reflection_point = False
                    break
            if not is_reflection_point:
                break

        if is_reflection_point and r + 1 != ignore:
            return r + 1

    return 0


def print_pattern(pattern):
    for r in range(len(pattern)):
        print("".join(pattern[r]))


n = 0
for pattern in patterns:
    n += check_vertical(pattern) + check_horizontal(pattern) * 100
print(n)


n = 0
for pattern in patterns:
    orig_v = check_vertical(pattern)
    orig_h = check_horizontal(pattern)
    fixed = False
    for r in range(len(pattern)):
        for c in range(len(pattern[0])):
            if pattern[r][c] == "#":
                pattern[r][c] = "."
            else:
                pattern[r][c] = "#"

            v = check_vertical(pattern, ignore=orig_v)
            h = check_horizontal(pattern, ignore=orig_h)

            if pattern[r][c] == "#":
                pattern[r][c] = "."
            else:
                pattern[r][c] = "#"

            if v != 0 or h != 0:
                n += v + h * 100
                fixed = True
                break

        if fixed:
            break


print(n)
