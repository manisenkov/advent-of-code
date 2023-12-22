from typing import TextIO
from sys import stdin


class Day21:
    grid: list[list[bool]]
    start_point: tuple[int, int]
    num_steps: tuple[int, int] = (64, 26501365)
    use_approximation: bool = True

    def __init__(self, input: TextIO):
        self.grid = []
        for row, line in enumerate(input.readlines()):
            self.grid.append([])
            for col, c in enumerate(line.strip()):
                self.grid[row].append(c == "#")
                if c == "S":
                    self.start_point = (row, col)

    def step(self, cur_steps: set[tuple[int, int]]) -> set[tuple[int, int]]:
        next_steps = set[tuple[int, int]]()
        for row, col in cur_steps:
            for drow, dcol in ((0, 1), (0, -1), (1, 0), (-1, 0)):
                next_row, next_col = row + drow, col + dcol
                if not self.grid[next_row % len(self.grid)][
                    next_col % len(self.grid[0])
                ]:
                    next_steps.add((next_row, next_col))
        return next_steps

    def solve(self, steps: int) -> int:
        cur_steps = set([self.start_point])
        for _ in range(steps):
            cur_steps = self.step(cur_steps)
        return len(cur_steps)

    def part1(self) -> int:
        return self.solve(self.num_steps[0])

    def part2(self) -> int:
        if self.use_approximation:
            size = len(self.grid)
            t_1 = self.num_steps[1] % size
            t_2 = self.num_steps[1] % size + size
            t_3 = self.num_steps[1] % size + 2 * size
            d_1 = self.solve(t_1)
            d_2 = self.solve(t_2)
            d_3 = self.solve(t_3)

            def lagrange(t) -> int:
                return (
                    (t - t_2) * (t - t_3) * d_1 // ((t_1 - t_2) * (t_1 - t_3))
                    + (t - t_1) * (t - t_3) * d_2 // ((t_2 - t_1) * (t_2 - t_3))
                    + (t - t_1) * (t - t_2) * d_3 // ((t_3 - t_1) * (t_3 - t_2))
                )

            return lagrange(self.num_steps[1])
        return self.solve(self.num_steps[1])


def main():
    sol = Day21(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
