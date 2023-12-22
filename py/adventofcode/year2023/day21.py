from functools import cache
from collections import defaultdict
from typing import TextIO
from sys import stdin


class Day21:
    grid: list[list[bool]]
    start_point: tuple[int, int]
    num_steps: tuple[int, int] = (64, 26501365)
    cache: dict[
        frozenset[tuple[int, int]],
        tuple[
            frozenset[tuple[int, int]],
            dict[tuple[int, int], frozenset[tuple[int, int]]],
        ],
    ]
    cache_misses = 0
    cache_hits = 0

    def __init__(self, input: TextIO):
        self.grid = []
        self.cache = {}
        for row, line in enumerate(input.readlines()):
            self.grid.append([])
            for col, c in enumerate(line.strip()):
                self.grid[row].append(c == "#")
                if c == "S":
                    self.start_point = (row, col)

    # @cache
    def step(
        self, cur_steps: frozenset[tuple[int, int]]
    ) -> tuple[
        frozenset[tuple[int, int]], dict[tuple[int, int], frozenset[tuple[int, int]]]
    ]:
        if cur_steps in self.cache:
            self.cache_hits += 1
            return self.cache[cur_steps]
        self.cache_misses += 1
        next_steps = frozenset[tuple[int, int]]()
        outer_steps = defaultdict[tuple[int, int], frozenset[tuple[int, int]]](
            frozenset
        )
        for row, col in cur_steps:
            for drow, dcol in ((0, 1), (0, -1), (1, 0), (-1, 0)):
                next_row, next_col = row + drow, col + dcol
                if 0 <= next_row < len(self.grid) and 0 <= next_col < len(self.grid[0]):
                    if not self.grid[next_row][next_col]:
                        next_steps |= {(next_row, next_col)}
                else:
                    block_row = (
                        -1 if next_row < 0 else (1 if next_row >= len(self.grid) else 0)
                    )
                    block_col = (
                        -1
                        if next_col < 0
                        else (1 if next_col >= len(self.grid[0]) else 0)
                    )
                    # Shift coordinates to block coordinates
                    next_row = (
                        len(self.grid) + next_row
                        if block_row == -1
                        else (next_row - len(self.grid) if block_row == 1 else next_row)
                    )
                    next_col = (
                        len(self.grid[0]) + next_col
                        if block_col == -1
                        else (
                            next_col - len(self.grid[0]) if block_col == 1 else next_col
                        )
                    )
                    outer_steps[(block_row, block_col)] |= {(next_row, next_col)}
        self.cache[cur_steps] = (next_steps, outer_steps)
        return next_steps, outer_steps

    def part1(self) -> int:
        cur_steps = frozenset([self.start_point])
        for _ in range(self.num_steps[0]):
            cur_steps, _ = self.step(cur_steps)
        return len(cur_steps)

    def part2(self) -> int:
        blocks: dict[tuple[int, int], frozenset[tuple[int, int]]] = {
            (0, 0): frozenset([self.start_point])
        }
        block_runs: dict[tuple[int, int], list[int]] = {}
        block_step: dict[tuple[int, int], int] = {(0, 0): 0}
        removed_blocks = set[tuple[int, int]]()
        sum_removed = 0

        for step in range(self.num_steps[1]):
            if step < 1000:
                if (step + 1) % 20 == 0:
                    print(
                        f"Step {step + 1}, num of blocks: {len(blocks)}, size of blocks: {sum(len(r) for r in blocks.values())}, cache ratio: {self.cache_hits / (self.cache_hits + self.cache_misses)}, hits: {self.cache_hits}, cache size: {len(self.cache)}"
                    )
            elif step < 10000:
                if (step + 1) % 500 == 0:
                    print(
                        f"Step {step + 1}, num of blocks: {len(blocks)}, size of blocks: {sum(len(r) for r in blocks.values())}, cache ratio: {self.cache_hits / (self.cache_hits + self.cache_misses)}, hits: {self.cache_hits}, cache size: {len(self.cache)}"
                    )
            elif step < 100000:
                if (step + 1) % 5000 == 0:
                    print(
                        f"Step {step + 1}, num of blocks: {len(blocks)}, size of blocks: {sum(len(r) for r in blocks.values())}, cache ratio: {self.cache_hits / (self.cache_hits + self.cache_misses)}, hits: {self.cache_hits}, cache size: {len(self.cache)}"
                    )
            else:
                if (step + 1) % 50000 == 0:
                    print(
                        f"Step {step + 1}, num of blocks: {len(blocks)}, size of blocks: {sum(len(r) for r in blocks.values())}, cache ratio: {self.cache_hits / (self.cache_hits + self.cache_misses)}, hits: {self.cache_hits}, cache size: {len(self.cache)}"
                    )
            to_update = defaultdict[tuple[int, int], frozenset[tuple[int, int]]](
                frozenset
            )
            for (block_row, block_col), block in blocks.items():
                next_steps, outer_steps = self.step(block)
                to_update[(block_row, block_col)] |= next_steps

                # Add outer blocks
                for (
                    outer_block_row_diff,
                    outer_block_col_diff,
                ), outer_block in outer_steps.items():
                    outer_block_row = block_row + outer_block_row_diff
                    outer_block_col = block_col + outer_block_col_diff
                    to_update[(outer_block_row, outer_block_col)] |= outer_block

            for block_idx, new_block in to_update.items():
                if block_idx in removed_blocks:
                    continue
                blocks[block_idx] = new_block
                new_run = len(new_block)
                if block_idx not in block_runs:
                    block_runs[block_idx] = [new_run]
                    block_step[block_idx] = step
                else:
                    block_runs[block_idx].append(new_run)
                the_run = block_runs[block_idx]
                if len(the_run) > 4 and the_run[-2:] == the_run[-4:-2]:
                    removed_blocks.add(block_idx)
                    tail_size = len(the_run) + block_step[block_idx] - 2
                    d = self.num_steps[1] - tail_size
                    idx = len(the_run) - 2 + (d % 2) - 1
                    sum_removed += the_run[idx]
                    del blocks[block_idx]
                    del block_runs[block_idx]
                    del block_step[block_idx]

        return sum(len(r) for r in blocks.values()) + sum_removed


def main():
    sol = Day21(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
