from io import StringIO

from adventofcode.year2023.day03 import Day03


INPUT = """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."""


def test_day03():
    sol = Day03(StringIO(INPUT))
    assert sol.part1() == 4361
    assert sol.part2() == 467835
