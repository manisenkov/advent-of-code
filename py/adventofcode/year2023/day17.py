from enum import IntEnum
from heapq import heappush, heappop, heapify
from collections import defaultdict
from typing import TextIO
from sys import stdin


class Direction(IntEnum):
    UP = 0
    RIGHT = 1
    DOWN = 2
    LEFT = 3

    def opposite(self):
        if self == Direction.UP:
            return Direction.DOWN
        elif self == Direction.DOWN:
            return Direction.UP
        elif self == Direction.LEFT:
            return Direction.RIGHT
        else:
            return Direction.LEFT

    def move(self, row: int, col: int) -> tuple[int, int]:
        if self == Direction.UP:
            return (row - 1, col)
        elif self == Direction.RIGHT:
            return (row, col + 1)
        elif self == Direction.DOWN:
            return (row + 1, col)
        else:
            return (row, col - 1)


State = tuple[
    int,  # row
    int,  # col
    int,  # count of front moves,
    Direction,  # direction of moving
]


class Day17:
    map: list[list[int]]

    def __init__(self, input: TextIO):
        self.map = [[int(c) for c in line.strip()] for line in input.readlines()]

    def setup(self) -> tuple[list[tuple[int, State]], defaultdict[State, int]]:
        queue: list[tuple[int, State]] = [
            (0, (0, 0, 0, dir))
            for dir in [
                Direction.UP,
                Direction.RIGHT,
                Direction.DOWN,
                Direction.LEFT,
            ]
        ]
        heapify(queue)
        dists = defaultdict[State, int](
            lambda: 0x7FFFFFFF,
            {
                (0, 0, count_front_moves, dir): 0
                for dir in [
                    Direction.UP,
                    Direction.RIGHT,
                    Direction.DOWN,
                    Direction.LEFT,
                ]
                for count_front_moves in range(10)
            },
        )
        return queue, dists

    def part1(self) -> int:
        height, width = len(self.map), len(self.map[0])
        queue, dists = self.setup()
        while queue:
            cur_dist, (row, col, count_front_moves, cur_dir) = heappop(queue)
            for next_dir in [
                Direction.UP,
                Direction.RIGHT,
                Direction.DOWN,
                Direction.LEFT,
            ]:
                if (
                    next_dir == cur_dir and count_front_moves == 3
                ) or next_dir == cur_dir.opposite():
                    continue
                next_row, next_col = next_dir.move(row, col)
                if (
                    next_row < 0
                    or next_col < 0
                    or next_row >= height
                    or next_col >= width
                ):
                    continue
                next_count_front_moves = (
                    count_front_moves + 1 if next_dir == cur_dir else 1
                )
                next_state = (next_row, next_col, next_count_front_moves, next_dir)
                next_dist = cur_dist + self.map[next_row][next_col]
                if next_dist < dists[next_state]:
                    dists[next_state] = next_dist
                    heappush(queue, (next_dist, next_state))
        return min(
            value
            for (row, col, _, _), value in dists.items()
            if row == height - 1 and col == width - 1
        )

    def part2(self) -> int:
        height, width = len(self.map), len(self.map[0])
        queue, dists = self.setup()
        while queue:
            cur_dist, (row, col, count_front_moves, cur_dir) = heappop(queue)
            if count_front_moves < 4:
                add_dist = 0
                next_row, next_col = row, col
                for _ in range(4 - count_front_moves):
                    next_row, next_col = cur_dir.move(next_row, next_col)
                    if (
                        next_row < 0
                        or next_col < 0
                        or next_row >= height
                        or next_col >= width
                    ):
                        add_dist = -1
                        break
                    add_dist += self.map[next_row][next_col]
                if add_dist == -1:
                    continue
                next_state = (next_row, next_col, 4, cur_dir)
                next_dist = cur_dist + add_dist
                if next_dist < dists[next_state]:
                    dists[next_state] = next_dist
                    heappush(queue, (next_dist, next_state))
            else:
                for next_dir in [
                    Direction.UP,
                    Direction.RIGHT,
                    Direction.DOWN,
                    Direction.LEFT,
                ]:
                    if next_dir == cur_dir.opposite():
                        continue
                    if next_dir == cur_dir and count_front_moves == 10:
                        continue
                    next_row, next_col = next_dir.move(row, col)
                    if (
                        next_row < 0
                        or next_col < 0
                        or next_row >= height
                        or next_col >= width
                    ):
                        continue
                    next_count_front_moves = (
                        count_front_moves + 1 if next_dir == cur_dir else 1
                    )
                    next_state = (next_row, next_col, next_count_front_moves, next_dir)
                    next_dist = cur_dist + self.map[next_row][next_col]
                    if next_dist < dists[next_state]:
                        dists[next_state] = next_dist
                        heappush(queue, (next_dist, next_state))
        return min(
            value
            for (row, col, _, _), value in dists.items()
            if row == height - 1 and col == width - 1
        )


def main():
    sol = Day17(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
