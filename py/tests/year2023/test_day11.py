from io import StringIO

from adventofcode.year2023.day11 import Day11


INPUT = """...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#....."""


def test_day11():
    sol = Day11(StringIO(INPUT))
    assert sol.part1() == 374
    assert sol.part2() == 82000210
