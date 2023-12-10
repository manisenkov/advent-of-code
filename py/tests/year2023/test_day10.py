from io import StringIO

from adventofcode.year2023.day10 import Day10


INPUT_1 = """-L|F7
7S-7|
L|7||
-L-J|
L|-JF"""

INPUT_2 = """7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ"""

INPUT_3 = """..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
.........."""

INPUT_4 = """.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ..."""

INPUT_5 = """FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L"""


def test_day10_1():
    sol = Day10(StringIO(INPUT_1))
    assert sol.part1() == 4


def test_day10_2():
    sol = Day10(StringIO(INPUT_2))
    assert sol.part1() == 8


def test_day10_3():
    sol = Day10(StringIO(INPUT_3))
    assert sol.part1() == 22
    assert sol.part2() == 4


def test_day10_4():
    sol = Day10(StringIO(INPUT_4))
    assert sol.part1() == 70
    assert sol.part2() == 8


def test_day10_5():
    sol = Day10(StringIO(INPUT_5))
    assert sol.part1() == 80
    assert sol.part2() == 10
