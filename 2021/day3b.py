from collections import Counter

with open("day3_input.txt", "r") as file:
    og_nums = [line.strip() for line in file]
    cs_nums = og_nums[:]

    for i in range(len(og_nums[0])):
        mc_counts = Counter(n[i] for n in og_nums).most_common()
        if len(mc_counts) == 1:
            mc = mc_counts[0][0]
        else:
            mc = "1" if mc_counts[0][1] == mc_counts[-1][1] else mc_counts[0][0]
        # print(mc_counts, mc)
        og_nums = [n for n in og_nums if n[i] == mc]

        cs_counts = Counter(n[i] for n in cs_nums).most_common()
        if len(cs_counts) == 1:
            lc = cs_counts[0][0]
        else:
            lc = "0" if cs_counts[0][1] == cs_counts[-1][1] else cs_counts[-1][0]
        cs_nums = [n for n in cs_nums if n[i] == lc]

print(og_nums, cs_nums)
ogr = og_nums[0]
csr = cs_nums[0]
print(ogr, int(ogr, 2))
print(csr, int(csr, 2))
print(int(csr, 2) * int(ogr, 2))
