from collections import defaultdict, deque
from typing import TextIO
from sys import stdin


class Day23:
    start_pos: tuple[int, int]
    end_pos: tuple[int, int]
    map: list[str]
    roads: dict[tuple[int, int], dict[tuple[int, int], int]]

    def __init__(self, input: TextIO):
        self.map = [s.strip() for s in input.readlines()]

        # Find crossroads
        self.start_pos = (0, 1)
        self.end_pos = (len(self.map) - 1, len(self.map[0]) - 2)
        crossroads: set[tuple[int, int]] = set()
        for row, line in enumerate(self.map):
            for col, cell in enumerate(line):
                if cell == "#":
                    continue
                if cell == "." and row > 0 and row < len(self.map) - 1:
                    neighbors = sum(
                        1
                        for drow, dcol in {(0, -1), (0, 1), (-1, 0), (1, 0)}
                        if self.map[row + drow][col + dcol] in {".", "<", ">", "^", "v"}
                    )
                    if neighbors > 2:
                        crossroads.add((row, col))
        crossroads.add(self.end_pos)
        self.roads = defaultdict[tuple[int, int], dict[tuple[int, int], int]](dict)
        self.roads[self.start_pos] = self.calc_distances(self.start_pos, crossroads)
        for xr in crossroads:
            res = self.calc_distances(xr, crossroads)
            self.roads[xr].update(res)

    def calc_distances(
        self, start_pos: tuple[int, int], crossroads: set[tuple[int, int]]
    ) -> dict[tuple[int, int], int]:
        res: dict[tuple[int, int], int] = {}
        queue = deque[tuple[int, tuple[int, int]]]([(0, start_pos)])
        visited = {start_pos}
        while queue:
            dist, cur_pos = queue.popleft()
            visited.add(cur_pos)
            for drow, dcol in {(0, -1), (0, 1), (-1, 0), (1, 0)}:
                next_pos = (cur_pos[0] + drow, cur_pos[1] + dcol)
                if (
                    next_pos[0] < 0
                    or next_pos[0] == len(self.map)
                    or next_pos[1] < 0
                    or next_pos[1] == len(self.map[0])
                    or self.map[next_pos[0]][next_pos[1]] == "#"
                    or next_pos in visited
                ):
                    continue
                if next_pos in crossroads:
                    res[next_pos] = dist + 1
                else:
                    queue.append((dist + 1, next_pos))
        return res

    def dfs(
        self,
        cur_pos: tuple[int, int],
        visited: frozenset[tuple[int, int]],
    ) -> int:
        if cur_pos == self.end_pos:
            return 0
        return max(
            (
                next_dist + self.dfs(next_pos, visited | {next_pos})
                for next_pos, next_dist in self.roads[cur_pos].items()
                if next_pos not in visited
            ),
            default=-100000,
        )

    def part1(self) -> int:
        allowed_moves = {
            ".": {(-1, 0), (1, 0), (0, -1), (0, 1)},
            "<": {(0, -1)},
            ">": {(0, 1)},
            "^": {(-1, 0)},
            "v": {(1, 0)},
        }
        queue = deque[tuple[tuple[int, int], frozenset[tuple[int, int]]]](
            [(self.start_pos, frozenset())]
        )
        distances = defaultdict[tuple[int, int], int](lambda: -1, [(self.start_pos, 0)])
        while queue:
            cur_pos, visited = queue.popleft()
            for drow, dcol in allowed_moves[self.map[cur_pos[0]][cur_pos[1]]]:
                next_pos = (cur_pos[0] + drow, cur_pos[1] + dcol)
                next_distance = distances[cur_pos] + 1
                if (
                    next_pos[0] < 0
                    or next_pos[0] == len(self.map)
                    or next_pos[1] < 0
                    or next_pos[1] == len(self.map[0])
                    or self.map[next_pos[0]][next_pos[1]] == "#"
                    or next_pos in visited
                    or distances[next_pos] > next_distance
                ):
                    continue
                distances[next_pos] = next_distance
                queue.append((next_pos, visited | {next_pos}))
        return distances[self.end_pos]

    def part2(self) -> int:
        return self.dfs(self.start_pos, frozenset([self.start_pos]))


def main():
    sol = Day23(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
