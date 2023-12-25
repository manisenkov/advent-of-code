from io import StringIO

from adventofcode.year2023.day24 import Day24


INPUT = """19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3"""


def test_day24():
    sol = Day24(StringIO(INPUT))
    sol.test_area = ((7, 27), (float("-inf"), float("inf")))
    assert sol.part1() == 2
    assert sol.part2() == 47
