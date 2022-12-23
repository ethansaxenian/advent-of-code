import itertools


def create_list(s):
    l = []
    i = 0
    while i < len(s):
        if not s[i].isnumeric():
            l.append(s[i])
            i += 1
        else:
            n = ""
            while s[i].isnumeric():
                n += s[i]
                i += 1
            l.append(n)

    return l


lines = list(map(create_list, map(str.strip, open("day18_input.txt", "r").readlines())))


def get_pair_indices(s, i):
    depth = 0
    for x in range(i, len(s)):
        if s[x] == "[":
            depth += 1
        if s[x] == "]":
            depth -= 1
        if depth == 0:
            return i+1, x-1


def get_next(s, i, direction):
    x = i + (1 if direction == "r" else -1)
    while x < len(s) and not s[x].isnumeric():
        x += 1 if direction == "r" else -1
    return x


def find_explode(l):
    depth = 0
    for i, c in enumerate(l):
        if c == "[":
            depth += 1
        if c == "]":
            depth -= 1

        if depth >= 5:
            return i

    return None


def explode(l, i):
    new_l = ["0"]
    a, b = get_pair_indices(l, i)
    ai = get_next(l, a, "l")
    bi = get_next(l, b, "r")
    for x in range(a - 2, -1, -1):
        if x == ai:
            new_l.insert(0, str(int(l[x]) + int(l[a])))
        else:
            new_l.insert(0, l[x])

    for x in range(b + 2, len(l)):
        if x == bi:
            new_l.append(str(int(l[x]) + int(l[b])))
        else:
            new_l.append(l[x])

    return new_l


def find_split(l):
    for i, c in enumerate(l):
        if c.isnumeric() and int(c) >= 10:
            return i
    return None


def split(l, i):
    n = int(l[i])
    a = n // 2
    b = a + (1 if n % 2 else 0)
    new_l = []
    for x, c in enumerate(l):
        if x == i:
            new_l.extend(["[", str(a), ",", str(b), "]"])
        else:
            new_l.append(c)
    return new_l


def get_mid(l):
    depth = 1
    i = 1
    while i < len(l):
        if l[i] == "[":
            depth += 1
        if l[i] == "]":
            depth -= 1

        if depth == 1:
            return i + 1

        i += 1


def mag(l):
    if l.isnumeric():
        return int(l)
    i = get_mid(l)
    left, right = l[1:i], l[i+1:-1]
    return 3 * mag(left) + 2 * mag(right)


def sum_snail_numbers(a, b):
    result = ["[", *a, ",", *b, "]"]
    while True:
        x = find_explode(result)
        while x is not None:
            result = explode(result, x)
            x = find_explode(result)

        x = find_split(result)
        if x is not None:
            result = split(result, x)

        if find_explode(result) is None and find_split(result) is None:
            break

    return mag("".join(result))


mags = []
for a, b in itertools.permutations(lines, 2):
    x = sum_snail_numbers(a, b)
    print(x)
    mags.append(x)

print(max(mags))
