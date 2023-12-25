from io import StringIO

from adventofcode.year2023.day25 import Day25


INPUT = """jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr"""


def test_day25():
    sol = Day25(StringIO(INPUT))
    assert sol.part1() == 54
    assert sol.part2() == "Merry Christmas! 🎄"
