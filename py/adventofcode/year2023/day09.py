from typing import TextIO
from sys import stdin


def get_seqs(history: list[int]) -> list[list[int]]:
    seqs = [list(history)]
    while any([n != 0 for n in seqs[-1]]):
        seqs.append([seqs[-1][i + 1] - seqs[-1][i] for i in range(len(seqs[-1]) - 1)])
    return seqs


class Day09:
    histories: list[list[int]]

    def __init__(self, input: TextIO):
        self.histories = [
            [int(s) for s in xs]
            for xs in (line.strip().split(" ") for line in input.readlines())
        ]

    def part1(self) -> int:
        res = 0
        for history in self.histories:
            seqs = get_seqs(history)
            seqs[-1].append(0)
            for i in range(len(seqs) - 2, -1, -1):
                seqs[i].append(seqs[i][-1] + seqs[i + 1][-1])
            res += seqs[0][-1]
        return res

    def part2(self) -> int:
        res = 0
        for history in self.histories:
            seqs = get_seqs(history)
            seqs[-1].insert(0, 0)
            for i in range(len(seqs) - 2, -1, -1):
                seqs[i].insert(0, seqs[i][0] - seqs[i + 1][0])
            res += seqs[0][0]
        return res


def main():
    sol = Day09(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
