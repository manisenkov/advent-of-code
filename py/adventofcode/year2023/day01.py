from typing import TextIO
from sys import stdin


WORD_MAP = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


class Solution:
    def __init__(self, input: TextIO):
        self.lines = [s.strip() for s in input.readlines()]

    def part1(self) -> int:
        nums: list[int] = []
        for line in self.lines:
            first_digit = line[next(i for i, c in enumerate(line) if c.isdigit())]
            last_digit = line[
                len(line) - next(i for i, c in enumerate(line[::-1]) if c.isdigit()) - 1
            ]
            nums.append(int(first_digit + last_digit))
        return sum(nums)

    def part2(self) -> int:
        nums: list[int] = []
        for line in self.lines:
            first_digit = ""
            for i in range(len(line)):
                if line[i].isdigit():
                    first_digit = line[i]
                else:
                    for word, num in WORD_MAP.items():
                        if line[i : i + len(word)] == word:
                            first_digit = num
                            break
                if first_digit != "":
                    break

            last_digit = ""
            for i in range(len(line) - 1, -1, -1):
                if line[i].isdigit():
                    last_digit = line[i]
                else:
                    for word, num in WORD_MAP.items():
                        if line[i : i + len(word)] == word:
                            last_digit = num
                            break
                if last_digit != "":
                    break
            nums.append(int(first_digit + last_digit))
        return sum(nums)


def main():
    sol = Solution(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
