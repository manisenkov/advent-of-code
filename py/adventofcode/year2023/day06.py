from typing import TextIO
from sys import stdin


class Day06:
    times: list[int]
    distances: list[int]

    def __init__(self, input: TextIO):
        lines = [s.strip() for s in input.readlines()]
        self.times = [
            int(s) for s in lines[0].split(":")[1].strip().split(" ") if s != ""
        ]
        self.distances = [
            int(s) for s in lines[1].split(":")[1].strip().split(" ") if s != ""
        ]

    def part1(self) -> int:
        res = 1
        for time, dist in zip(self.times, self.distances):
            d = time**2 - 4 * dist
            ts = sorted([int((-time - d**0.5) / -2), int((-time + d**0.5) / -2)])
            if (time - ts[0]) * ts[0] <= dist:
                ts[0] += 1
            if (time - ts[1]) * ts[1] <= dist:
                ts[1] -= 1
            res *= ts[1] - ts[0] + 1
        return res

    def part2(self) -> int:
        time = int("".join(str(t) for t in self.times))
        dist = int("".join(str(d) for d in self.distances))
        d = time**2 - 4 * dist
        ts = sorted([int((-time - d**0.5) / -2), int((-time + d**0.5) / -2)])
        if (time - ts[0]) * ts[0] <= dist:
            ts[0] += 1
        if (time - ts[1]) * ts[1] <= dist:
            ts[1] -= 1
        return ts[1] - ts[0] + 1


def main():
    sol = Day06(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
