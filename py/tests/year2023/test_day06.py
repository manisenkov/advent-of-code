from io import StringIO

from adventofcode.year2023.day06 import Day06


INPUT = """Time:      7  15   30
Distance:  9  40  200"""


def test_day06():
    sol = Day06(StringIO(INPUT))
    assert sol.part1() == 288
    assert sol.part2() == 71503
