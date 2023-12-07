from collections import Counter
from util.util import fetch_input
import functools

puzzle_input = fetch_input(7)

card_values = {
    "A": 14,
    "K": 13,
    "Q": 12,
    "J": 11,
    "T": 10,
    "9": 9,
    "8": 8,
    "7": 7,
    "6": 6,
    "5": 5,
    "4": 4,
    "3": 3,
    "2": 2,
}


def hand_type(hand: str) -> int:
    num_distinct = len(set(hand))
    if num_distinct == 5:
        return 0
    elif num_distinct == 4:
        return 1
    elif num_distinct == 1:
        return 6

    counts = Counter(hand).most_common()
    if len(counts) == 2:
        if counts[0][1] == 4:
            return 5

        return 4

    if len(counts) == 3 and counts[0][1] == 3:
        return 3

    return 2


def hand_type_with_jokers(hand: str) -> int:
    if "J" not in hand:
        return hand_type(hand)

    max_hand_type = -1
    for c in card_values.keys():
        hand_copy = hand.replace("J", c)
        max_hand_type = max(max_hand_type, hand_type(hand_copy))

    return max_hand_type


@functools.total_ordering
class Hand:
    def __init__(self, hand: str, bid: str, use_jokers: bool = False):
        self.orig = hand
        self.hand = []
        for c in hand:
            if c == "J" and use_jokers:
                self.hand.append(1)
            else:
                self.hand.append(card_values[c])

        self.bid = int(bid)
        if use_jokers:
            self.hand_type = hand_type_with_jokers(hand)
        else:
            self.hand_type = hand_type(hand)

    def __repr__(self):
        return f"{self.orig} {self.bid} - {self.hand_type}"

    def __eq__(self, other):
        return self.hand == other.hand

    def __lt__(self, other):
        if self.hand_type != other.hand_type:
            return self.hand_type < other.hand_type

        for x, y in zip(self.hand, other.hand):
            if x == y:
                continue
            return x < y


def get_hands(use_jokers: bool = False):
    hands = [
        Hand(hand, bid, use_jokers=use_jokers)
        for hand, bid in [line.split(" ") for line in puzzle_input if line != ""]
    ]

    return sorted(hands)


def part1():
    hands = get_hands()
    print(sum([hand.bid * (i + 1) for i, hand in enumerate(hands)]))


def part2():
    hands = get_hands(use_jokers=True)
    print(sum([hand.bid * (i + 1) for i, hand in enumerate(hands)]))


part1()
part2()
