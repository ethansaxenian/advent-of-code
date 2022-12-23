unique_lengths = {
    2: 1,
    3: 7,
    4: 4,
    7: 8
}


with open("day8_input.txt", "r") as file:
    output_values = []
    for line in file:
        out = line.strip().split(" | ")[1].split(" ")
        output_values.extend(out)

    print(sum(1 for v in output_values if len(v) in unique_lengths))
