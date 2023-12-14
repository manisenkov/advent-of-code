from io import StringIO

from adventofcode.year2023.day14 import Day14


INPUT = """O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#...."""


def test_day14():
    sol = Day14(StringIO(INPUT))
    assert sol.part1() == 136
    assert sol.part2() == 64
