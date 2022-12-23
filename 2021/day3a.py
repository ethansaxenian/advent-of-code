from collections import Counter

with open("day3_input.txt", "r") as file:
    nums = [line.strip() for line in file]
    most_common = []

    for i in range(len(nums[0])):
        most_common.append(Counter(n[i] for n in nums).most_common(1)[0][0])

gamma = "".join(most_common)
epsilon = "".join(str(int(not int(b))) for b in gamma)
print(gamma, int(gamma, 2))
print(epsilon, int(epsilon, 2))
print(int(gamma, 2) * int(epsilon, 2))
