from collections import defaultdict
from enum import Enum
from typing import TextIO
from sys import stdin


class Direction(Enum):
    RIGHT = 0
    DOWN = 1
    LEFT = 2
    UP = 3


def in_range(value: int, checked_range: tuple[int, int]) -> bool:
    return value >= checked_range[0] and value <= checked_range[1]


def dig(commands: list[tuple[Direction, int, str]]) -> int:
    # Calculate horizontal trenches
    pos = (0, 0)
    horizontal_trenches = defaultdict[int, list[tuple[int, int]]](list)
    for dir, dist, _ in commands:
        if dir == Direction.RIGHT:
            next = (pos[0], pos[1] + dist)
            horizontal_trenches[pos[0]].append((pos[1], next[1]))
        elif dir == Direction.DOWN:
            next = (pos[0] + dist, pos[1])
            for i in range(pos[0], next[0] + 1):
                horizontal_trenches[i].append((pos[1], next[1]))
        elif dir == Direction.LEFT:
            next = (pos[0], pos[1] - dist)
            horizontal_trenches[pos[0]].append((next[1], pos[1]))
        elif dir == Direction.UP:
            next = (pos[0] - dist, pos[1])
            for i in range(next[0], pos[0] + 1):
                horizontal_trenches[i].append((next[1], pos[1]))
        pos = next

    # Merge trenches
    for trench in horizontal_trenches.values():
        trench.sort()
        i = 0
        while i < len(trench) - 1:
            if trench[i + 1][0] - trench[i][1] <= 1:
                trench[i] = (trench[i][0], trench[i + 1][1])
                del trench[i + 1]
            else:
                i += 1

    print(f" -- Rows total: {len(horizontal_trenches)}")

    # Calculate area per row
    res = 0
    counter = 1
    for row, trenches in horizontal_trenches.items():
        counter += 1
        if counter % 200000 == 0:
            progress = (counter / len(horizontal_trenches)) * 100
            print(f" -- {progress:.2f}% done...")
        start_pos = None
        for t in trenches:
            is_change = (
                (row - 1) in horizontal_trenches
                and (row + 1) in horizontal_trenches
                and (
                    (
                        any(in_range(t[0], tr) for tr in horizontal_trenches[row - 1])
                        and any(
                            in_range(t[1], tr) for tr in horizontal_trenches[row + 1]
                        )
                    )
                    or (
                        any(in_range(t[0], tr) for tr in horizontal_trenches[row + 1])
                        and any(
                            in_range(t[1], tr) for tr in horizontal_trenches[row - 1]
                        )
                    )
                )
            )
            if is_change:
                if start_pos is None:
                    start_pos = t[0]
                else:
                    res += t[1] - start_pos + 1
                    start_pos = None
            elif start_pos is None:
                res += t[1] - t[0] + 1
    print(" -- Completed!")
    return res


class Day18:
    commands: list[tuple[Direction, int, str]]

    def __init__(self, input: TextIO):
        self.commands = []
        for line in input.readlines():
            dir_str, dist_str, color_str = line.strip().split(" ")
            if dir_str == "U":
                dir = Direction.UP
            elif dir_str == "R":
                dir = Direction.RIGHT
            elif dir_str == "D":
                dir = Direction.DOWN
            else:
                dir = Direction.LEFT
            dist = int(dist_str)
            color = color_str[1:-1]
            self.commands.append((dir, dist, color))

    def part1(self) -> int:
        return dig(self.commands)

    def part2(self) -> int:
        commands = [
            (Direction(int(color[-1])), int(color[1:-1], 16), color)
            for (_, _, color) in self.commands
        ]
        return dig(commands)


def main():
    sol = Day18(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
