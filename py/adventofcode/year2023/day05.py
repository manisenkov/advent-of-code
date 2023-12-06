from functools import reduce
from dataclasses import dataclass
from typing import TextIO
from sys import stdin


@dataclass
class Map:
    source: str
    target: str
    ranges: list[tuple[tuple[int, int], tuple[int, int]]]

    def convert_range(self, source: tuple[int, int]) -> list[tuple[int, int]]:
        source_ranges = [source]
        for r in self.ranges:
            [low, high] = source_ranges.pop()
            [target_low, target_high] = r[0]
            if target_low >= high:
                source_ranges.append((low, high))
                break
            elif target_high <= low:
                source_ranges.append((low, high))
                continue
            elif target_low <= low and target_high >= high:
                source_ranges.append((low, high))
                break
            elif target_low <= low and target_high < high:
                source_ranges.extend([(low, target_high), (target_high + 1, high)])
            elif target_low > low and target_high >= high:
                source_ranges.extend([(low, target_low - 1), (target_low, high)])
            else:
                source_ranges.extend(
                    [
                        (low, target_low - 1),
                        (target_low, target_high),
                        (target_high + 1, high),
                    ]
                )
        return [(self.convert(r[0]), self.convert(r[1])) for r in source_ranges]

    def convert(self, from_id: int) -> int:
        for r in self.ranges:
            if from_id >= r[0][0] and from_id <= r[0][1]:
                return r[1][0] + (from_id - r[0][0])
        return from_id


class Day05:
    seeds: list[int]
    maps: dict[str, Map]

    def __init__(self, input: TextIO):
        lines = [s.strip() for s in input.readlines()]
        self.seeds = [int(s) for s in lines[0].split(":")[1].strip().split(" ")]
        self.maps = {}
        i = 2
        while i < len(lines):
            [source, _, target] = lines[i].split(" ")[0].split("-")
            try:
                next_line = lines.index("", i + 1)
            except ValueError:
                next_line = len(lines)
            range_defs = []
            for range_str in lines[i + 1 : next_line]:
                range_defs.append([int(s) for s in range_str.split(" ")])
            self.maps[source] = Map(
                source,
                target,
                sorted(
                    [
                        ((r[1], r[1] + r[2] - 1), (r[0], r[0] + r[2] - 1))
                        for r in range_defs
                    ],
                    key=lambda r: r[0],
                ),
            )
            i = next_line + 1

    def part1(self) -> int:
        cur_num = list(self.seeds)
        cur_type = "seed"
        while cur_type != "location":
            mapper = self.maps[cur_type]
            cur_num = [mapper.convert(x) for x in cur_num]
            cur_type = self.maps[cur_type].target
        return min(cur_num)

    def part2(self) -> int:
        cur_ranges = [
            (self.seeds[i * 2], self.seeds[i * 2] + self.seeds[i * 2 + 1] - 1)
            for i in range(int(len(self.seeds) / 2))
        ]
        cur_type = "seed"
        while cur_type != "location":
            mapper = self.maps[cur_type]
            cur_ranges = sorted(
                reduce(
                    lambda r1, r2: r1 + r2,  # type: ignore[operator]
                    [mapper.convert_range(r) for r in cur_ranges],
                ),
                key=lambda r: r[0],
            )
            cur_type = self.maps[cur_type].target
        cur_ranges.sort(key=lambda r: r[0])
        return cur_ranges[0][0]


def main():
    sol = Day05(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
