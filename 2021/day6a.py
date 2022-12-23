from tqdm import tqdm


def count_fish(days):

    with open("day6_input.txt", "r") as file:
        fish = [int(x) for x in file.readline().split(",")]

        for _ in tqdm(range(days)):
            newfish = 0
            for i in tqdm(range(len(fish))):
                fish[i] -= 1
                if fish[i] == -1:
                    newfish += 1
                    fish[i] = 6
            fish.extend([8 for _ in range(newfish)])

    return len(fish)


if __name__ == '__main__':
    print(count_fish(80))
