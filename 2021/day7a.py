from statistics import mean, median, mode


def get_input():
    with open("day7_input.txt", "r") as file:
        return [int(x) for x in file.readline().split(",")]


def calculate_fuel(crabs, target_fn, fuel_modifier):
    target = target_fn(crabs)
    return sum(fuel_modifier(abs(x - target)) for x in crabs)


if __name__ == '__main__':
    crabs = get_input()
    print(calculate_fuel(crabs, median, lambda x: x))
