from typing import TextIO
from sys import stdin


def validate(arrangement_part: str, operational_size: int, broken_size: int) -> bool:
    return all(c in (".", "?") for c in arrangement_part[:operational_size]) and all(
        c in ("#", "?")
        for c in arrangement_part[operational_size : operational_size + broken_size]
    )


def solve(
    arrangement: str, group: tuple[int, ...], pos: int, group_idx: int, table: dict
) -> int:
    key = (pos, group_idx)
    if key in table:
        return table[key]

    res = 0
    if group_idx >= len(group):
        res = 1 if len([c for c in arrangement[pos:] if c == "#"]) == 0 else 0
    else:
        broken_size = group[group_idx]
        res = sum(
            solve(
                arrangement,
                group,
                pos + operational_size + broken_size,
                group_idx + 1,
                table,
            )
            for operational_size in range(
                0 if group_idx == 0 else 1, len(arrangement) - pos - broken_size + 1
            )
            if validate(arrangement[pos:], operational_size, broken_size)
        )

    table[key] = res
    return res


class Day12:
    arrangements: list[str]
    groups: list[tuple[int, ...]]

    def __init__(self, input: TextIO):
        self.arrangements = []
        self.groups = []
        for line in input.readlines():
            arrangement, group_str = line.split(" ")
            self.arrangements.append(arrangement)
            self.groups.append(tuple(int(g) for g in group_str.split(",")))

    def part1(self) -> int:
        return sum(
            solve(arrangement, group, 0, 0, {})
            for arrangement, group in zip(self.arrangements, self.groups)
        )

    def part2(self) -> int:
        return sum(
            solve("?".join([arrangement] * 5), (group * 5), 0, 0, {})
            for arrangement, group in zip(self.arrangements, self.groups)
        )


def main():
    sol = Day12(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
