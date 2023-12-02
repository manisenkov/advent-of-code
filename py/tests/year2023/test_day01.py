from io import StringIO

from adventofcode.year2023.day01 import Day01


INPUT_1 = """1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet"""


INPUT_2 = """two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen"""


def test_day01_1():
    sol = Day01(StringIO(INPUT_1))
    assert sol.part1() == 142


def test_day01_2():
    sol = Day01(StringIO(INPUT_2))
    assert sol.part2() == 281
