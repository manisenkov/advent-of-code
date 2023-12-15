from io import StringIO

from adventofcode.year2023.day15 import Day15, hash


INPUT = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"


def test_day15():
    sol = Day15(StringIO(INPUT))
    assert hash("HASH") == 52
    assert sol.part1() == 1320
    assert sol.part2() == 145
