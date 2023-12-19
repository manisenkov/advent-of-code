from functools import reduce
from itertools import combinations
from collections import deque
from dataclasses import dataclass
from typing import Optional, TextIO
from sys import stdin


@dataclass
class Part:
    x: int
    m: int
    a: int
    s: int

    @property
    def rating(self) -> int:
        return self.x + self.m + self.a + self.s


@dataclass
class PartRange:
    x: tuple[int, int]
    m: tuple[int, int]
    a: tuple[int, int]
    s: tuple[int, int]

    @property
    def rating(self) -> int:
        return (
            (self.x[1] - self.x[0] + 1)
            * (self.m[1] - self.m[0] + 1)
            * (self.a[1] - self.a[0] + 1)
            * (self.s[1] - self.s[0] + 1)
        )

    def intersect(self, other: "PartRange") -> "PartRange":
        if (
            self.x[0] > other.x[1]
            or self.x[1] < other.x[0]
            or self.m[0] > other.m[1]
            or self.m[1] < other.m[0]
            or self.a[0] > other.a[1]
            or self.a[1] < other.a[0]
            or self.s[0] > other.s[1]
            or self.s[1] < other.s[0]
        ):
            return PartRange((0, -1), (0, -1), (0, -1), (0, -1))
        return PartRange(
            x=(max(self.x[0], other.x[0]), min(self.x[1], other.x[1])),
            m=(max(self.m[0], other.m[0]), min(self.m[1], other.m[1])),
            a=(max(self.a[0], other.a[0]), min(self.a[1], other.a[1])),
            s=(max(self.s[0], other.s[0]), min(self.s[1], other.s[1])),
        )


@dataclass
class BasicRule:
    target: str

    def check(self, _: Part) -> bool:
        raise NotImplementedError()

    def split(self, part: PartRange) -> tuple[Optional[PartRange], Optional[PartRange]]:
        raise NotImplementedError()


@dataclass
class UnconditionalRule(BasicRule):
    def check(self, _: Part) -> bool:
        return True

    def split(self, part: PartRange) -> tuple[Optional[PartRange], Optional[PartRange]]:
        return part, None


@dataclass
class ConditionalRule(BasicRule):
    prop: str
    is_greater: bool
    value: int

    def check(self, part: Part) -> bool:
        return (
            (getattr(part, self.prop) > self.value)
            if self.is_greater
            else (getattr(part, self.prop) < self.value)
        )

    def split(self, part: PartRange) -> tuple[Optional[PartRange], Optional[PartRange]]:
        part_low, part_high = getattr(part, self.prop)
        if (self.is_greater and part_high < self.value) or (
            not self.is_greater and part_low > self.value
        ):
            return None, part
        new_range = (
            (max(self.value + 1, part_low), part_high)
            if self.is_greater
            else (part_low, min(self.value - 1, part_high))
        )
        new_part = PartRange(
            x=part.x if self.prop != "x" else new_range,
            m=part.m if self.prop != "m" else new_range,
            a=part.a if self.prop != "a" else new_range,
            s=part.s if self.prop != "s" else new_range,
        )

        # Subtract new_part from part
        if new_range == (part_low, part_high):
            next_part = None
        else:
            next_range = (
                new_range[1] + 1 if part_low == new_range[0] else part_low,
                new_range[0] - 1 if part_high == new_range[1] else part_high,
            )
            next_part = PartRange(
                x=part.x if self.prop != "x" else next_range,
                m=part.m if self.prop != "m" else next_range,
                a=part.a if self.prop != "a" else next_range,
                s=part.s if self.prop != "s" else next_range,
            )
        return new_part, next_part


class Day19:
    workflows: dict[str, list[BasicRule]]
    parts: list[Part]

    def __init__(self, input: TextIO):
        lines = [s.strip() for s in input.readlines()]
        rule_lines, part_lines = lines[: lines.index("")], lines[lines.index("") + 1 :]
        self.workflows = {}
        for line in rule_lines:
            name, workflow_str = line.split("{")
            workflow_str = workflow_str[:-1]
            self.workflows[name] = []
            for rule_def in workflow_str.split(","):
                if ":" in rule_def:
                    statement, target = rule_def.split(":")
                    if ">" in statement:
                        prop, value = statement.split(">")
                        self.workflows[name].append(
                            ConditionalRule(target, prop, True, int(value))
                        )
                    else:
                        prop, value = statement.split("<")
                        self.workflows[name].append(
                            ConditionalRule(target, prop, False, int(value))
                        )
                else:
                    self.workflows[name].append(UnconditionalRule(rule_def))
        self.parts = [
            Part(*[int(s.split("=")[1]) for s in line[1:-1].split(",")])
            for line in part_lines
        ]

    def part1(self) -> int:
        res = 0
        for part in self.parts:
            cur_workflow = "in"
            while cur_workflow not in ["A", "R"]:
                for rule in self.workflows[cur_workflow]:
                    if (
                        isinstance(rule, ConditionalRule) and rule.check(part)
                    ) or isinstance(rule, UnconditionalRule):
                        cur_workflow = rule.target
                        break
            if cur_workflow == "A":
                res += part.rating
        return res

    def part2(self) -> int:
        queue = deque[tuple[str, PartRange]](
            [("in", PartRange((1, 4000), (1, 4000), (1, 4000), (1, 4000)))]
        )
        successful_parts: list[PartRange] = []
        while queue:
            next_workflow, cur_part_range = queue.popleft()
            if next_workflow == "A":
                successful_parts.append(cur_part_range)
            elif next_workflow != "R":
                for rule in self.workflows[next_workflow]:
                    new_part, next_part = rule.split(cur_part_range)
                    if new_part is not None:
                        queue.append((rule.target, new_part))
                    if next_part is not None:
                        cur_part_range = next_part
                    else:
                        break
        res = 0
        factor = 1
        for i in range(1, len(successful_parts) + 1):
            next_res = res
            for combination in combinations(successful_parts, i):
                intersect = reduce(lambda x, y: x.intersect(y), combination)
                next_res += intersect.rating * factor
            if next_res == res:
                break
            res = next_res
            factor *= -1
        return res


def main():
    sol = Day19(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
