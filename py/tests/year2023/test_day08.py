from io import StringIO

from adventofcode.year2023.day08 import Day08


INPUT_1 = """RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)"""


INPUT_2 = """LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)"""

INPUT_3 = """LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)"""


def test_day08_1():
    sol = Day08(StringIO(INPUT_1))
    assert sol.part1() == 2


def test_day08_2():
    sol = Day08(StringIO(INPUT_2))
    assert sol.part1() == 6


def test_day08_3():
    sol = Day08(StringIO(INPUT_3))
    assert sol.part2() == 6
