from day7a import get_input, calculate_fuel


def add_to_n(n):
    return n * (n+1) // 2


if __name__ == '__main__':
    crabs = get_input()

    sums = []
    for i in range(0, max(crabs) + 1):
        sums.append(calculate_fuel(crabs, lambda _: i, add_to_n))

    print(min(sums))
