from collections import defaultdict
from typing import TextIO
from sys import stdin


def hash(s: str) -> int:
    res = 0
    for c in s:
        res = ((res + ord(c)) * 17) % 256
    return res


class Day15:
    commands: list[str]

    def __init__(self, input: TextIO):
        line = input.readlines()[0].strip()
        self.commands = line.split(",")

    def part1(self) -> int:
        return sum(hash(c) for c in self.commands)

    def part2(self) -> int:
        boxes = defaultdict[int, dict[str, tuple[int, int]]](dict)
        for c in self.commands:
            if c[-1] == "-":  # remove lense
                label = c[:-1]
                box = boxes[hash(label)]
                if label in box:
                    _, idx = box[label]
                    del box[label]
                    box_items = list(box.items())
                    for k, (focal, pos) in box_items:
                        if pos > idx:
                            del box[k]
                            box[k] = (focal, pos - 1)
            else:  # add lense
                [label, focal_str] = c.split("=")
                focal = int(focal_str)
                box = boxes[hash(label)]
                if label in box:
                    _, idx = box[label]
                    box[label] = (focal, idx)
                else:
                    box[label] = (focal, len(box))
        return sum(
            sum((box_idx + 1) * (pos + 1) * focal for (focal, pos) in box.values())
            for box_idx, box in boxes.items()
        )


def main():
    sol = Day15(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
