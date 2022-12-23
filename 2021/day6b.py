def count_fish(days):
    with open("day6_input.txt", "r") as file:
        fish = [int(x) for x in file.readline().split(",")]
        nums = [sum(1 for x in fish if x == i) or -1 for i in range(9)]

        for _ in range(days):
            new_nums = nums[:]
            for i in range(9):
                if i == 8:
                    if nums[0] != -1:
                        new_nums[i] = nums[0]
                    else:
                        new_nums[i] = 0
                else:
                    if nums[i+1] != -1:
                        new_nums[i] = nums[i+1]
                    else:
                        new_nums[i] = 0

            if nums[0] != -1:
                new_nums[6] += nums[0]

            nums = new_nums[:]

    return sum(nums)


if __name__ == '__main__':
    print(count_fish(256))
