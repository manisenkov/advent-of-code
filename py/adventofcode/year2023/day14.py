from enum import Enum
from dataclasses import dataclass
from typing import TextIO
from sys import stdin


class Direction(Enum):
    NORTH = 0
    EAST = 1
    SOUTH = 2
    WEST = 3


@dataclass
class Field:
    size: tuple[int, int]
    round_rocks: set[tuple[int, int]]
    cube_rocks: set[tuple[int, int]]

    def move(self, dir: Direction) -> "Field":
        round_rocks_list = list(self.round_rocks)
        if dir == Direction.NORTH:
            round_rocks_list.sort(key=lambda x: (x[0], x[1]))
        elif dir == Direction.EAST:
            round_rocks_list.sort(key=lambda x: (x[1], x[0]), reverse=True)
        elif dir == Direction.SOUTH:
            round_rocks_list.sort(key=lambda x: (x[0], x[1]), reverse=True)
        elif dir == Direction.WEST:
            round_rocks_list.sort(key=lambda x: (x[1], x[0]))

        next_round_rocks = set(self.round_rocks)
        for i, j in round_rocks_list:
            if dir == Direction.NORTH:
                for n in range(i - 1, -2, -1):
                    if (
                        n == -1
                        or (n, j) in self.cube_rocks
                        or (n, j) in next_round_rocks
                    ):
                        next_round_rocks.remove((i, j))
                        next_round_rocks.add((n + 1, j))
                        break
            elif dir == Direction.EAST:
                for n in range(j + 1, self.size[1] + 1):
                    if (
                        n == self.size[1]
                        or (i, n) in self.cube_rocks
                        or (i, n) in next_round_rocks
                    ):
                        next_round_rocks.remove((i, j))
                        next_round_rocks.add((i, n - 1))
                        break
            elif dir == Direction.SOUTH:
                for n in range(i + 1, self.size[0] + 1):
                    if (
                        n == self.size[0]
                        or (n, j) in self.cube_rocks
                        or (n, j) in next_round_rocks
                    ):
                        next_round_rocks.remove((i, j))
                        next_round_rocks.add((n - 1, j))
                        break
            elif dir == Direction.WEST:
                for n in range(j - 1, -2, -1):
                    if (
                        n == -1
                        or (i, n) in self.cube_rocks
                        or (i, n) in next_round_rocks
                    ):
                        next_round_rocks.remove((i, j))
                        next_round_rocks.add((i, n + 1))
                        break
        return Field(self.size, next_round_rocks, self.cube_rocks)


class Day14:
    initial_field: Field

    def __init__(self, input: TextIO):
        map = [s.strip() for s in input.readlines() if s.strip() != ""]
        size = (len(map), len(map[0]))
        rounded_rocks = set()
        cube_rocks = set()
        for i, line in enumerate(map):
            for j, char in enumerate(line):
                if char == "O":
                    rounded_rocks.add((i, j))
                elif char == "#":
                    cube_rocks.add((i, j))
        self.initial_field = Field(size, rounded_rocks, cube_rocks)

    def part1(self) -> int:
        field = self.initial_field.move(dir=Direction.NORTH)
        height, _ = field.size
        return sum(height - i for (i, j) in field.round_rocks)

    def part2(self) -> int:
        height, _ = self.initial_field.size
        xs = []
        field = self.initial_field
        pattern_size = 0
        for i in range(100000):
            next_field = (
                field.move(dir=Direction.NORTH)
                .move(dir=Direction.WEST)
                .move(dir=Direction.SOUTH)
                .move(dir=Direction.EAST)
            )
            if next_field == field:
                break
            field = next_field
            xs.append(sum(height - i for (i, j) in next_field.round_rocks))
            pattern_size = 0
            for i in range(5, len(xs) // 2 + 1):
                if xs[-i:] == xs[-i * 2 : -i]:
                    pattern_size = i
                    break
            if pattern_size > 0:
                break
        start_pos = len(xs) - pattern_size * 2
        return xs[(1000000000 - start_pos) % pattern_size - 1 + start_pos]


def main():
    sol = Day14(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
