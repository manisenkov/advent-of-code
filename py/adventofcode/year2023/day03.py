from itertools import chain
from typing import TextIO
from sys import stdin


class Day03:
    tool_coords: dict[tuple[int, int], str]
    num_coords: dict[tuple[int, int], int]
    max_num_len: int

    def __init__(self, input: TextIO):
        self.tool_coords = {}
        self.num_coords = {}
        self.max_num_len = 0
        for row, line in enumerate(s.strip() for s in input.readlines()):
            col = 0
            while col < len(line):
                if line[col].isdigit():
                    all_coords = []
                    start_col = col
                    num = ""
                    while col < len(line) and line[col].isdigit():
                        num += line[col]
                        all_coords.append((row, col))
                        col += 1
                    if len(num) > self.max_num_len:
                        self.max_num_len = len(num)
                    self.num_coords[(row, start_col)] = int(num)
                else:
                    if line[col] != ".":
                        self.tool_coords[(row, col)] = line[col]
                    col += 1

    def find_adjacent_tools(self, coord: tuple[int, int]) -> list[tuple[int, int]]:
        num_len = len(str(self.num_coords[coord]))
        row, col = coord
        neighbours = chain(
            ((row - 1, c) for c in range(col - 1, col + num_len + 1)),
            [(row, col - 1), (row, col + num_len)],
            ((row + 1, c) for c in range(col - 1, col + num_len + 1)),
        )
        return list(filter(lambda k: k in self.tool_coords, neighbours))

    def find_adjacent_numbers(self, coord: tuple[int, int]) -> list[tuple[int, int]]:
        row, col = coord
        neighbours_left = chain(
            ((row - 1, c) for c in range(col - self.max_num_len, col)),
            ((row, c) for c in range(col - self.max_num_len, col)),
            ((row + 1, c) for c in range(col - self.max_num_len, col)),
        )
        neighbours_right = chain(
            ((row - 1, c) for c in range(col, col + 2)),
            [(row, col + 1)],
            ((row + 1, c) for c in range(col, col + 2)),
        )
        return list(
            filter(
                lambda k: k in self.num_coords
                and k[1] + len(str(self.num_coords[k])) >= col,
                neighbours_left,
            )
        ) + list(filter(lambda k: k in self.num_coords, neighbours_right))

    def part1(self) -> int:
        return sum(
            num
            for coords, num in self.num_coords.items()
            if len(self.find_adjacent_tools(coords)) > 0
        )

    def part2(self) -> int:
        return sum(
            self.num_coords[adjacent_numbers[0]] * self.num_coords[adjacent_numbers[1]]
            for adjacent_numbers in (
                self.find_adjacent_numbers(gear_coord)
                for gear_coord, tool in self.tool_coords.items()
                if tool == "*"
            )
            if len(adjacent_numbers) == 2
        )


def main():
    sol = Day03(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
