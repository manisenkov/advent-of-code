from io import StringIO

from adventofcode.year2023.day22 import Day22


INPUT = """1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9"""


def test_day22():
    sol = Day22(StringIO(INPUT))
    assert sol.part1() == 5
    assert sol.part2() == 7
