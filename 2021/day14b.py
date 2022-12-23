from collections import defaultdict


def parse_input():
    with open("day14_input.txt", "r") as file:
        sequence = file.readline().strip()

        mapping = {}

        file.readline()
        next_line = file.readline()
        while next_line:
            pair, letter = next_line.strip().split(" -> ")
            mapping[pair] = letter

            next_line = file.readline()

        return sequence, mapping


if __name__ == '__main__':
    sequence, mapping = parse_input()
    counts = defaultdict(int)
    letter_counts = defaultdict(int)
    for i in range(len(sequence) - 1):
        counts[f"{sequence[i]}{sequence[i+1]}"] += 1

    for c in sequence:
        letter_counts[c] += 1

    for _ in range(40):
        for p, i in list(counts.items()):
            counts[p] -= i
            x = mapping[p]
            a, b = p
            counts[f"{a}{x}"] += i
            counts[f"{x}{b}"] += i
            letter_counts[x] += i

    print(max(letter_counts.items(), key=lambda pair: pair[1])[1] - min(letter_counts.items(), key=lambda pair: pair[1])[1])
