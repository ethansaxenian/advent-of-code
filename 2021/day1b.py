with open("day1_input.txt", "r") as file:
    nums = [int(line.strip()) for line in file]

    n = 0
    prev = 0
    for i in range(len(nums)):
        try:
            if int(nums[i]) + int(nums[i+1]) + int(nums[i+2]) > prev:
                n += 1
            prev = int(nums[i]) + int(nums[i+1]) + int(nums[i+2])
        except IndexError:
            continue

print(n-1)
