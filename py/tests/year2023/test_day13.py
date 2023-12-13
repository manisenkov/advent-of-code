from io import StringIO

from adventofcode.year2023.day13 import Day13


INPUT = """#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#"""


def test_day13():
    sol = Day13(StringIO(INPUT))
    assert sol.part1() == 405
    assert sol.part2() == 400
