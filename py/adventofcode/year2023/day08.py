from math import lcm
from typing import TextIO
from sys import stdin


class Day08:
    turns: str
    map: dict[str, tuple[str, str]]

    def __init__(self, input: TextIO):
        lines = input.readlines()
        self.map = {}
        self.turns = lines[0].strip()
        for line in lines[2:]:
            [start, dest] = line.strip().split(" = ")
            [left, right] = dest[1:-1].split(", ")
            self.map[start] = (left, right)

    def get_steps_to_end(self, start: str, end_on_zzz: bool) -> int:
        cur = start
        count = 0
        i = 0
        while cur != "ZZZ" if end_on_zzz else "Z" not in cur:
            count += 1
            cur = self.map[cur][0 if self.turns[i] == "L" else 1]
            i = (i + 1) % len(self.turns)
        return count

    def part1(self) -> int:
        return self.get_steps_to_end("AAA", end_on_zzz=True)

    def part2(self) -> int:
        factors = [
            self.get_steps_to_end(k, end_on_zzz=False)
            for k in self.map.keys()
            if "A" in k
        ]
        return lcm(*factors)


def main():
    sol = Day08(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
