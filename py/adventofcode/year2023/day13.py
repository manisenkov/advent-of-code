from typing import TextIO
from sys import stdin


def find_mirrors(pattern: list[list[str]], is_vertical: bool, is_smudged: bool) -> int:
    numbers = []
    if is_vertical:
        for col in range(len(pattern[0])):
            n = 0
            for row in range(len(pattern)):
                if pattern[row][col] == "#":
                    n += 2**row
            numbers.append(n)
    else:
        for row in range(len(pattern)):
            n = 0
            for col in range(len(pattern[0])):
                if pattern[row][col] == "#":
                    n += 2**col
            numbers.append(n)
    for i in range(1, len(numbers)):
        left = numbers[:i]
        right = numbers[i:]
        size = min(len(left), len(right))
        left = left[-size:]
        right = right[:size]
        if is_smudged:
            xored = [l ^ r for l, r in zip(left, right[::-1]) if l ^ r != 0]
            if len(xored) == 1 and (xored[0] & (xored[0] - 1) == 0):
                return i
        elif left == right[::-1]:
            return i
    return 0


class Day13:
    patterns: list[list[list[str]]]

    def __init__(self, input: TextIO):
        self.patterns = []
        pattern: list[list[str]] = []
        for line in [s.strip() for s in input.readlines()]:
            if line != "":
                pattern.append(list(line))
            else:
                self.patterns.append(pattern)
                pattern = []
        if pattern:
            self.patterns.append(pattern)

    def part1(self) -> int:
        return sum(
            100 * find_mirrors(pattern, is_vertical=False, is_smudged=False)
            + find_mirrors(pattern, is_vertical=True, is_smudged=False)
            for pattern in self.patterns
        )

    def part2(self) -> int:
        return sum(
            100 * find_mirrors(pattern, is_vertical=False, is_smudged=True)
            + find_mirrors(pattern, is_vertical=True, is_smudged=True)
            for pattern in self.patterns
        )


def main():
    sol = Day13(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
