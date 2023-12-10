from enum import Enum
from typing import TextIO
from sys import stdin


class Direction(Enum):
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


POSSIBLE_CONNECTIONS: dict[tuple[str, Direction], set[str]] = {
    ("-", Direction.UP): set(),
    ("-", Direction.DOWN): set(),
    ("-", Direction.LEFT): {"-", "F", "L"},
    ("-", Direction.RIGHT): {"-", "7", "J"},
    ("|", Direction.UP): {"|", "F", "7"},
    ("|", Direction.DOWN): {"|", "L", "J"},
    ("|", Direction.LEFT): set(),
    ("|", Direction.RIGHT): set(),
    ("F", Direction.UP): set(),
    ("F", Direction.DOWN): {"|", "L", "J"},
    ("F", Direction.LEFT): set(),
    ("F", Direction.RIGHT): {"-", "J", "7"},
    ("L", Direction.UP): {"|", "7", "F"},
    ("L", Direction.DOWN): set(),
    ("L", Direction.LEFT): set(),
    ("L", Direction.RIGHT): {"-", "J", "7"},
    ("J", Direction.UP): {"|", "7", "F"},
    ("J", Direction.DOWN): set(),
    ("J", Direction.LEFT): {"-", "F", "L"},
    ("J", Direction.RIGHT): set(),
    ("7", Direction.UP): set(),
    ("7", Direction.DOWN): {"|", "J", "L"},
    ("7", Direction.LEFT): {"-", "L", "F"},
    ("7", Direction.RIGHT): set(),
    ("S", Direction.UP): {"|", "F", "7"},
    ("S", Direction.DOWN): {"|", "J", "L"},
    ("S", Direction.LEFT): {"-", "L", "F"},
    ("S", Direction.RIGHT): {"-", "J", "7"},
}


class Day10:
    map: list[list[str]]
    start_point: tuple[int, int]
    loop: list[tuple[int, int]]

    def __init__(self, input: TextIO):
        self.map = [list(s.strip()) for s in input.readlines()]
        self.loop = []
        for r, row in enumerate(self.map):
            for c, cell in enumerate(row):
                if cell == "S":
                    self.start_point = (r, c)
                    return

    def validate_pos(self, row: int, col: int) -> bool:
        return row >= 0 and col >= 0 and row < len(self.map) and col < len(self.map[0])

    def part1(self) -> int:
        # Find loop
        loop: list[tuple[int, int]] = []
        for dir in [Direction.UP, Direction.RIGHT, Direction.DOWN, Direction.LEFT]:
            (cur_row, cur_col) = self.start_point
            next_row, next_col = dir.move(cur_row, cur_col)
            if (
                not self.validate_pos(next_row, next_col)
                or self.map[next_row][next_col] not in POSSIBLE_CONNECTIONS[("S", dir)]
            ):
                continue
            visited = set()
            cur_dir = dir
            while (next_row, next_col) != self.start_point:
                # Search for next move
                loop.append((next_row, next_col))
                visited.add((next_row, next_col))
                cur_row, cur_col = next_row, next_col
                found = False
                for cell_dir in [
                    d
                    for d in [
                        Direction.UP,
                        Direction.RIGHT,
                        Direction.DOWN,
                        Direction.LEFT,
                    ]
                    if d != cur_dir.opposite()
                ]:
                    (next_row, next_col) = cell_dir.move(cur_row, cur_col)
                    if (
                        (next_row, next_col) not in visited
                        and self.validate_pos(next_row, next_col)
                        and (
                            (
                                self.map[next_row][next_col] == "S"
                                and self.map[cur_row][cur_col]
                                in POSSIBLE_CONNECTIONS[("S", cell_dir.opposite())]
                            )
                            or self.map[next_row][next_col]
                            in POSSIBLE_CONNECTIONS[
                                (self.map[cur_row][cur_col], cell_dir)
                            ]
                        )
                    ):
                        found = True
                        cur_dir = cell_dir
                        break
                if not found:
                    loop = []
                    break
            if len(loop) > 0:
                break
        loop.insert(0, self.start_point)
        self.loop = loop
        return len(loop) // 2

    def part2(self) -> int:
        # Replace "S" with a pipe
        diff = {
            (self.loop[1][0] - self.loop[0][0], self.loop[1][1] - self.loop[0][1]),
            (self.loop[-1][0] - self.loop[0][0], self.loop[-1][1] - self.loop[0][1]),
        }
        (row, col) = self.start_point
        if diff == {(0, 1), (1, 0)}:
            self.map[row][col] = "F"
        elif diff == {(0, 1), (-1, 0)}:
            self.map[row][col] = "L"
        elif diff == {(0, -1), (1, 0)}:
            self.map[row][col] = "7"
        elif diff == {(0, -1), (-1, 0)}:
            self.map[row][col] = "J"
        elif diff == {(-1, 0), (1, 0)}:
            self.map[row][col] = "|"
        else:
            self.map[row][col] = "-"

        min_row = min(row for row, _ in self.loop)
        min_col = min(col for _, col in self.loop)
        max_row = max(row for row, _ in self.loop)
        max_col = max(col for _, col in self.loop)
        scaled_map = [
            ["."] * ((max_col - min_col) * 2 + 3)
            for _ in range((max_row - min_row) * 2 + 3)
        ]

        # Scale up the map
        for row, col in self.loop:
            target_row = (row - min_row) * 2 + 1
            target_col = (col - min_col) * 2 + 1
            cell = self.map[row][col]
            scaled_map[target_row][target_col] = cell
            if cell == "F":
                scaled_map[target_row][target_col + 1] = "-"
                scaled_map[target_row + 1][target_col] = "|"
            elif cell == "-":
                scaled_map[target_row][target_col + 1] = "-"
                scaled_map[target_row + 1][target_col] = "."
            elif cell == "7":
                scaled_map[target_row][target_col + 1] = "."
                scaled_map[target_row + 1][target_col] = "|"
            elif cell == "|":
                scaled_map[target_row][target_col + 1] = "."
                scaled_map[target_row + 1][target_col] = "|"
            elif cell == "J":
                scaled_map[target_row][target_col + 1] = "."
                scaled_map[target_row + 1][target_col] = "."
            else:
                scaled_map[target_row][target_col + 1] = "-"
                scaled_map[target_row + 1][target_col] = "."

        # Fill up outer part
        stack: list[tuple[int, int]] = [(0, 0)]
        while len(stack) > 0:
            row, col = stack.pop()
            scaled_map[row][col] = "O"
            for dir in [Direction.UP, Direction.RIGHT, Direction.DOWN, Direction.LEFT]:
                next_row, next_col = dir.move(row, col)
                if (
                    next_row < 0
                    or next_col < 0
                    or next_row >= len(scaled_map)
                    or next_col >= len(scaled_map[0])
                    or scaled_map[next_row][next_col] != "."
                ):
                    continue
                stack.append((next_row, next_col))

        # Count scaled blocks of four dots
        result = 0
        counted: set[tuple[int, int]] = set()
        for row in range(len(scaled_map)):
            for col in range(len(scaled_map[row])):
                if scaled_map[row][col] != "." or (row, col) in counted:
                    continue
                if (
                    scaled_map[row][col] == "."
                    and scaled_map[row][col + 1] == "."
                    and scaled_map[row + 1][col] == "."
                    and scaled_map[row + 1][col + 1] == "."
                ):
                    counted.update(
                        [(row, col), (row, col + 1), (row + 1, col), (row + 1, col + 1)]
                    )
                    result += 1

        return result


def main():
    sol = Day10(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
