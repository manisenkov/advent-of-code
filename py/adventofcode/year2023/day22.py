from functools import cache
from collections import namedtuple
from typing import TextIO, Optional
from sys import stdin


Position = namedtuple("Position", ["x", "y", "z"])
Brick = tuple[Position, Position]


@cache
def get_brick_blocks(brick: Brick) -> list[Position]:
    left, right = brick
    blocks = []
    for x in range(left.x, right.x + 1):
        for y in range(left.y, right.y + 1):
            for z in range(left.z, right.z + 1):
                blocks.append(Position(x, y, z))
    return blocks


def settle_bricks(bricks: list[Brick]) -> list[Brick]:
    bricks = bricks.copy()

    # Settle bricks first
    occupied: dict[Position, int] = {}
    for idx, brick in enumerate(bricks):
        for pos in get_brick_blocks(brick):
            occupied[pos] = idx

    # Repeat until all blocks are settled
    while True:
        # Find bricks that are not settled
        not_settled: list[int] = []
        for idx, brick in enumerate(bricks):
            if min(brick[0].z, brick[1].z) == 1:  # Brick is on the ground
                is_settled = True
                continue
            is_settled = False
            for pos in get_brick_blocks(brick):
                if (pos.x, pos.y, pos.z - 1) in occupied and occupied[
                    Position(pos.x, pos.y, pos.z - 1)
                ] != idx:
                    is_settled = True
                    break
            if not is_settled:
                not_settled.append(idx)
        if not not_settled:
            break

        # Move not settled bricks
        for idx in not_settled:
            brick = bricks[idx]
            for pos in get_brick_blocks(brick):
                del occupied[pos]
            cur_brick = brick
            while True:
                next_brick = (
                    Position(cur_brick[0].x, cur_brick[0].y, cur_brick[0].z - 1),
                    Position(cur_brick[1].x, cur_brick[1].y, cur_brick[1].z - 1),
                )
                if next_brick[0].z == 0 or next_brick[1].z == 0:
                    break
                new_pos_found = True
                for pos in get_brick_blocks(next_brick):
                    if pos in occupied:
                        new_pos_found = False
                        break
                if new_pos_found:
                    cur_brick = next_brick
                else:
                    break
            bricks[idx] = cur_brick
            for pos in get_brick_blocks(cur_brick):
                occupied[pos] = idx

    return bricks


class Day22:
    bricks: list[Brick]
    result: Optional[tuple[int, int]] = None

    def __init__(self, input: TextIO):
        self.bricks = []
        for line in input.readlines():
            left_str, right_str = line.strip().split("~")
            left = Position(*(int(c) for c in left_str.split(",")))
            right = Position(*(int(c) for c in right_str.split(",")))
            self.bricks.append((left, right))

    def solve(self):
        bricks = settle_bricks(self.bricks)
        result_bricks = 0
        result_falls = 0
        for idx in range(len(bricks)):
            next_bricks = bricks.copy()
            del next_bricks[idx]
            settled_bricks = settle_bricks(next_bricks)
            fall_bricks_count = sum(
                1 if s != n else 0 for s, n in zip(settled_bricks, next_bricks)
            )
            if fall_bricks_count > 0:
                result_falls += fall_bricks_count
            else:
                result_bricks += 1
        self.result = result_bricks, result_falls

    def part1(self) -> int:
        if self.result is None:
            self.solve()
        return self.result[0]  # type: ignore

    def part2(self) -> int:
        if self.result is None:
            self.solve()
        return self.result[1]  # type: ignore


def main():
    sol = Day22(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
