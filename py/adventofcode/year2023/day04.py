# from itertools import chain
from dataclasses import dataclass
from typing import TextIO
from sys import stdin


@dataclass
class Card:
    winning: set[int]
    numbers: set[int]

    def count(self) -> int:
        return len(self.winning.intersection(self.numbers))


class Day04:
    cards: dict[int, Card]
    match_counts: dict[int, int]

    def __init__(self, input: TextIO):
        self.cards = {}
        for line in input.readlines():
            [name, row] = line.split(":")
            [part_winning, part_numbers] = row.strip().split("|")
            self.cards[int(name.split(" ")[-1])] = Card(
                set(int(s) for s in part_winning.strip().split(" ") if s != ""),
                set(int(s) for s in part_numbers.strip().split(" ") if s != ""),
            )

    def part1(self) -> int:
        self.match_counts = dict((i, c.count()) for i, c in self.cards.items())
        return sum(2 ** (n - 1) if n > 0 else 0 for n in self.match_counts.values())

    def part2(self) -> int:
        card_counts = dict((i, 1) for i in self.cards)
        for i in range(1, len(card_counts) + 1):
            matches = self.match_counts[i]
            for j in range(i + 1, i + matches + 1):
                card_counts[j] += card_counts[i]
        return sum(card_counts.values())


def main():
    sol = Day04(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
