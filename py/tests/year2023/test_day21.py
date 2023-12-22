from io import StringIO

from adventofcode.year2023.day21 import Day21


INPUT = """...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
..........."""


def test_day21():
    sol = Day21(StringIO(INPUT))
    sol.num_steps = (6, 5000)
    assert sol.part1() == 16
    assert sol.part2() == 16733044
