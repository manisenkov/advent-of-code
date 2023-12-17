from io import StringIO

from adventofcode.year2023.day16 import Day16


INPUT = """.|...\\....
|.-.\\.....
.....|-...
........|.
..........
.........\\
..../.\\\\..
.-.-/..|..
.|....-|.\\
..//.|...."""


def test_day16():
    sol = Day16(StringIO(INPUT))
    assert sol.part1() == 46
    assert sol.part2() == 51
