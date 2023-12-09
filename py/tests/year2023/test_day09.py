from io import StringIO

from adventofcode.year2023.day09 import Day09


INPUT = """0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45"""


def test_day09():
    sol = Day09(StringIO(INPUT))
    assert sol.part1() == 114
    assert sol.part2() == 2
