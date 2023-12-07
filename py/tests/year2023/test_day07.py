from io import StringIO

from adventofcode.year2023.day07 import Day07


INPUT = """32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483"""


def test_day07():
    sol = Day07(StringIO(INPUT))
    assert sol.part1() == 6440
    assert sol.part2() == 5905
