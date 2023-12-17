from collections import defaultdict
from enum import Enum
from typing import TextIO, Optional
from sys import stdin


class Direction(Enum):
    UP = 0
    RIGHT = 1
    DOWN = 2
    LEFT = 3

    def move(
        self, pos: tuple[int, int], size: tuple[int, int]
    ) -> Optional[tuple[int, int]]:
        i, j = pos
        if self == Direction.UP:
            res = (i - 1, j)
        elif self == Direction.RIGHT:
            res = (i, j + 1)
        elif self == Direction.DOWN:
            res = (i + 1, j)
        elif self == Direction.LEFT:
            res = (i, j - 1)
        if res[0] < 0 or res[1] < 0 or res[0] >= size[0] or res[1] >= size[1]:
            return None
        return res

    def apply(self, cell: str) -> list["Direction"]:
        if cell == "\\":
            if self == Direction.UP:
                return [Direction.LEFT]
            elif self == Direction.LEFT:
                return [Direction.UP]
            elif self == Direction.DOWN:
                return [Direction.RIGHT]
            elif self == Direction.RIGHT:
                return [Direction.DOWN]
        elif cell == "/":
            if self == Direction.UP:
                return [Direction.RIGHT]
            elif self == Direction.RIGHT:
                return [Direction.UP]
            elif self == Direction.DOWN:
                return [Direction.LEFT]
            elif self == Direction.LEFT:
                return [Direction.DOWN]
        elif cell == "|" and self in [Direction.LEFT, Direction.RIGHT]:
            return [Direction.UP, Direction.DOWN]
        elif cell == "-" and [Direction.UP, Direction.DOWN]:
            return [Direction.LEFT, Direction.RIGHT]
        return [self]


class Day16:
    size: tuple[int, int]
    map: dict[tuple[int, int], str]

    def __init__(self, input: TextIO):
        lines = [s.strip() for s in input.readlines() if s.strip() != ""]
        self.size = (len(lines), len(lines[0]))
        self.map = {}
        for i, line in enumerate(lines):
            for j, cell in enumerate(line):
                if cell != ".":
                    self.map[(i, j)] = cell

    def run_beam(
        self, start_pos: tuple[int, int], start_dir: Direction
    ) -> defaultdict[tuple[int, int], set[Direction]]:
        res = defaultdict[tuple[int, int], set[Direction]](set)
        stack = [(start_pos, start_dir)]
        while stack:
            pos, dir = stack.pop()
            if pos in res and dir in res[pos]:
                continue
            res[pos].add(dir)
            if not pos in self.map:
                next_pos = dir.move(pos, self.size)
                if next_pos is not None:
                    stack.append((next_pos, dir))
            else:
                cell = self.map[pos]
                for next_dir in dir.apply(cell):
                    next_pos = next_dir.move(pos, self.size)
                    if next_pos is not None:
                        stack.append((next_pos, next_dir))
        return res

    def part1(self) -> int:
        res = self.run_beam((0, 0), Direction.RIGHT)
        return len(list(res.keys()))

    def part2(self) -> int:
        max_energy = 0
        width, height = self.size
        for i in range(width):
            res = self.run_beam((i, 0), Direction.RIGHT)
            energy = len(list(res.keys()))
            max_energy = max(energy, max_energy)
        for i in range(width):
            res = self.run_beam((i, height - 1), Direction.LEFT)
            energy = len(list(res.keys()))
            max_energy = max(energy, max_energy)
        for j in range(height):
            res = self.run_beam((0, j), Direction.DOWN)
            energy = len(list(res.keys()))
            max_energy = max(energy, max_energy)
        for j in range(height):
            res = self.run_beam((width - 1, j), Direction.UP)
            energy = len(list(res.keys()))
            max_energy = max(energy, max_energy)
        return max_energy


def main():
    sol = Day16(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
