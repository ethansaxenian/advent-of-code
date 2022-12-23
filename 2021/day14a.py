from collections import Counter


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


def do_insertions(sequence, mapping, iterations=1):
    old_sequence = sequence
    new_sequence = ""
    for _ in range(iterations):
        new_sequence = ""

        for i, e in enumerate(old_sequence):
            new_sequence += e
            try:
                new_sequence += mapping[f"{e}{old_sequence[i+1]}"]
            except IndexError:
                continue

        old_sequence = new_sequence

    return new_sequence


if __name__ == '__main__':
    sequence, mapping = parse_input()
    new_sequence = do_insertions(sequence, mapping, 10)
    counts = Counter(new_sequence)
    print(counts.most_common(1)[0][1] - counts.most_common()[-1][1])
