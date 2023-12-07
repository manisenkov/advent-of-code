from collections import defaultdict
from functools import cache, cmp_to_key
from enum import Enum
from typing import TextIO, Callable
from sys import stdin


CARD_ORDER = dict((c, i) for (i, c) in enumerate("AKQJT98765432"[::-1]))
CARD_ORDER_JOKER = dict((c, i) for (i, c) in enumerate("AKQT98765432J"[::-1]))


class Combination(Enum):
    HIGH = 0
    ONE_PAIR = 1
    TWO_PAIR = 2
    THREE = 3
    FULL_HOUSE = 4
    FOUR = 5
    FIVE = 6


@cache
def get_cards_combination(cards: str) -> Combination:
    groups: defaultdict[str, int] = defaultdict(int)
    for c in cards:
        groups[c] += 1
    vals = list(groups.values())
    if len(vals) == 1:
        return Combination.FIVE
    elif len(vals) == 2:
        if 4 in vals:
            return Combination.FOUR
        else:
            return Combination.FULL_HOUSE
    elif len(vals) == 3:
        if 2 in vals:
            return Combination.TWO_PAIR
        else:
            return Combination.THREE
    elif len(vals) == 4:
        return Combination.ONE_PAIR
    else:
        return Combination.HIGH


@cache
def get_cards_combination_joker(cards: str) -> Combination:
    groups: defaultdict[str, int] = defaultdict(int)
    for c in cards:
        groups[c] += 1
    if "J" in groups:
        # Try to replace jokers with one of other cards
        combinations = sorted(
            [
                get_cards_combination(cards.replace("J", card))
                for card in groups
                if card != "J"
            ]
            + [get_cards_combination(cards)],
            key=lambda c: c.value,
        )
        return combinations[-1]
    else:
        return get_cards_combination(cards)


def get_card_comparator(
    combination_calculator: Callable[[str], Combination], card_order: dict[str, int]
) -> Callable[[tuple[str, int], tuple[str, int]], int]:
    def compare_hands(hand1: tuple[str, int], hand2: tuple[str, int]) -> int:
        cards1, _ = hand1
        cards2, _ = hand2
        combination1 = combination_calculator(cards1)
        combination2 = combination_calculator(cards2)
        if combination1 == combination2:
            for i in range(5):
                if card_order[cards1[i]] > card_order[cards2[i]]:
                    return 1
                elif card_order[cards1[i]] < card_order[cards2[i]]:
                    return -1
            return 0
        elif combination1.value > combination2.value:
            return 1
        else:
            return -1

    return compare_hands


class Day07:
    hands: list[tuple[str, int]]

    def __init__(self, input: TextIO):
        self.hands_and_bids = [
            (c.strip(), int(b.strip()))
            for [c, b] in (line.split(" ") for line in input.readlines())
        ]

    def part1(self) -> int:
        return sum(
            (i + 1) * bid
            for i, (_, bid) in enumerate(
                sorted(
                    self.hands_and_bids,
                    key=cmp_to_key(
                        get_card_comparator(get_cards_combination, CARD_ORDER)
                    ),
                )
            )
        )

    def part2(self) -> int:
        return sum(
            (i + 1) * bid
            for i, (_, bid) in enumerate(
                sorted(
                    self.hands_and_bids,
                    key=cmp_to_key(
                        get_card_comparator(
                            get_cards_combination_joker, CARD_ORDER_JOKER
                        )
                    ),
                )
            )
        )


def main():
    sol = Day07(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
