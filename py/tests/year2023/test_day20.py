from io import StringIO

from adventofcode.year2023.day20 import Day20


INPUT_1 = """broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a"""

INPUT_2 = """broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output"""


def test_day20_1():
    sol = Day20(StringIO(INPUT_1))
    assert sol.part1() == 32000000


def test_day20_2():
    sol = Day20(StringIO(INPUT_2))
    assert sol.part1() == 11687500
