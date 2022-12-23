unique_lengths = {
    2: 1,
    3: 7,
    4: 4,
    7: 8
}


with open("day8_input.txt", "r") as file:
    tot_sum = 0

    for line in file:
        signal_patterns = line.strip().split(" | ")[0].split(" ")
        output_values = line.strip().split(" | ")[1].split(" ")
        mapping = {unique_lengths[len(i)]: i for i in signal_patterns if len(i) in unique_lengths}

        for i in signal_patterns:
            if len(i) == 5:
                if all(c in i for c in mapping[1]):
                    mapping[3] = i
            if len(i) == 6:
                if not all(c in i for c in mapping[1]):
                    mapping[6] = i
                elif all(c in i for c in mapping[4]):
                    mapping[9] = i
                else:
                    mapping[0] = i
        for i in signal_patterns:
            if len(i) == 5 and mapping[3] != i:
                if all(c in mapping[9] for c in i):
                    mapping[5] = i
                else:
                    mapping[2] = i

        rev_mapping = {"".join(sorted(p)): i for (i, p) in mapping.items()}

        n = "".join([str(rev_mapping["".join(sorted(w))]) for w in output_values])
        tot_sum += int(n)

    print(tot_sum)
