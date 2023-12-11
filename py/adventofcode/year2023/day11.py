from typing import TextIO
from sys import stdin


def calc_dist_sum(galaxies: list[tuple[int, int]]) -> int:
    return sum(
        sum(
            abs(row_2 - row_1) + abs(col_2 - col_1)
            for (row_2, col_2) in galaxies[i + 1 :]
        )
        for (i, (row_1, col_1)) in enumerate(galaxies[:-1])
    )


class Day11:
    map: list[list[str]]
    empty_rows: list[int]
    empty_cols: list[int]

    def __init__(self, input: TextIO):
        self.map = [list(s.strip()) for s in input.readlines()]
        self.empty_rows = [
            i for i in range(len(self.map)) if all(c == "." for c in self.map[i])
        ]
        self.empty_cols = [
            j
            for j in range(len(self.map[0]))
            if all(c == "." for c in [self.map[i][j] for i in range(len(self.map))])
        ]

    def part1(self) -> int:
        galaxies: list[tuple[int, int]] = []
        for i, row in enumerate(self.map):
            for j, cell in enumerate(row):
                if cell == "#":
                    galaxies.append(
                        (
                            i + len([r for r in self.empty_rows if r < i]),
                            j + len([c for c in self.empty_cols if c < j]),
                        )
                    )
        return calc_dist_sum(galaxies)

    def part2(self) -> int:
        galaxies: list[tuple[int, int]] = []
        for i, row in enumerate(self.map):
            for j, cell in enumerate(row):
                if cell == "#":
                    galaxies.append(
                        (
                            i + 999999 * len([r for r in self.empty_rows if r < i]),
                            j + 999999 * len([c for c in self.empty_cols if c < j]),
                        )
                    )
        return calc_dist_sum(galaxies)


def main():
    sol = Day11(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
