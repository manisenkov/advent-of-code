from sympy.matrices import Matrix
from math import sqrt
from dataclasses import dataclass
from typing import TextIO
from sys import stdin


@dataclass(frozen=True)
class Vec:
    x: int
    y: int
    z: int

    def __abs__(self) -> float:
        return sqrt(self.x**2 + self.y**2 + self.z**2)

    def __add__(self, other: "Vec") -> "Vec":
        return Vec(self.x + other.x, self.y + other.y, self.z + other.z)

    def __sub__(self, other: "Vec") -> "Vec":
        return Vec(self.x - other.x, self.y - other.y, self.z - other.z)

    def __mul__(self, other: int) -> "Vec":
        return Vec(self.x * other, self.y * other, self.z * other)

    def __rmul__(self, other: int) -> "Vec":
        return self.__mul__(other)

    def normalize(self):
        return Vec(self.x / abs(self.x), self.y / abs(self.y), self.z / abs(self.z))

    def codir_2(self, other) -> bool:
        self_norm = self.normalize()
        other_norm = other.normalize()
        return self_norm.x == other_norm.x and self_norm.y == other_norm.y


@dataclass(frozen=True)
class Halestone:
    pos: Vec
    vel: Vec


class Day24:
    test_area: tuple[tuple[int, int], tuple[int, int]] = (
        (200000000000000, 400000000000000),
        (200000000000000, 400000000000000),
    )
    halestones: list[Halestone]

    def __init__(self, input: TextIO):
        self.halestones = []
        for line in input.readlines():
            pos_part, vel_part = line.split(" @ ")
            pos = Vec(*map(int, pos_part.split(", ")))
            vel = Vec(*map(int, vel_part.split(", ")))
            self.halestones.append(Halestone(pos, vel))

    def part1(self) -> int:
        factors: list[tuple[float, float]] = []
        for hs in self.halestones:
            pos_0 = hs.pos
            pos_1 = hs.pos + hs.vel
            factors.append(
                (
                    (pos_1.y - pos_0.y) / (pos_1.x - pos_0.x),
                    (pos_0.y * pos_1.x - pos_0.x * pos_1.y) / (pos_1.x - pos_0.x),
                )
            )
        res = 0
        for i, hs_1 in enumerate(self.halestones[:-1]):
            for j in range(i + 1, len(self.halestones)):
                hs_2 = self.halestones[j]
                if factors[i][0] - factors[j][0] == 0:
                    continue
                cross_x = (factors[j][1] - factors[i][1]) / (
                    factors[i][0] - factors[j][0]
                )
                cross_y = factors[i][0] * cross_x + factors[i][1]
                cross = Vec(int(cross_x), int(cross_y), 0)

                # Check if the cross point is not in the past of halestones
                cross_dir_1 = cross - hs_1.pos
                cross_dir_2 = cross - hs_2.pos
                if (
                    cross_dir_1.codir_2(hs_1.vel)
                    and cross_dir_2.codir_2(hs_2.vel)
                    and cross.x > self.test_area[0][0]
                    and cross.x < self.test_area[0][1]
                    and cross.y > self.test_area[1][0]
                    and cross.y < self.test_area[1][1]
                ):
                    res += 1
        return res

    def part2(self) -> int:
        hs_0, hs_1, hs_2 = self.halestones[:3]
        a = Matrix(
            [
                [
                    hs_1.vel.y - hs_0.vel.y,
                    hs_0.vel.x - hs_1.vel.x,
                    0.0,
                    hs_0.pos.y - hs_1.pos.y,
                    hs_1.pos.x - hs_0.pos.x,
                    0.0,
                ],
                [
                    hs_2.vel.y - hs_0.vel.y,
                    hs_0.vel.x - hs_2.vel.x,
                    0.0,
                    hs_0.pos.y - hs_2.pos.y,
                    hs_2.pos.x - hs_0.pos.x,
                    0.0,
                ],
                [
                    hs_1.vel.z - hs_0.vel.z,
                    0.0,
                    hs_0.vel.x - hs_1.vel.x,
                    hs_0.pos.z - hs_1.pos.z,
                    0.0,
                    hs_1.pos.x - hs_0.pos.x,
                ],
                [
                    hs_2.vel.z - hs_0.vel.z,
                    0.0,
                    hs_0.vel.x - hs_2.vel.x,
                    hs_0.pos.z - hs_2.pos.z,
                    0.0,
                    hs_2.pos.x - hs_0.pos.x,
                ],
                [
                    0.0,
                    hs_1.vel.z - hs_0.vel.z,
                    hs_0.vel.y - hs_1.vel.y,
                    0.0,
                    hs_0.pos.z - hs_1.pos.z,
                    hs_1.pos.y - hs_0.pos.y,
                ],
                [
                    0.0,
                    hs_2.vel.z - hs_0.vel.z,
                    hs_0.vel.y - hs_2.vel.y,
                    0.0,
                    hs_0.pos.z - hs_2.pos.z,
                    hs_2.pos.y - hs_0.pos.y,
                ],
            ],
        )
        b = Matrix(
            [
                (hs_0.pos.y * hs_0.vel.x - hs_1.pos.y * hs_1.vel.x)
                - (hs_0.pos.x * hs_0.vel.y - hs_1.pos.x * hs_1.vel.y),
                (hs_0.pos.y * hs_0.vel.x - hs_2.pos.y * hs_2.vel.x)
                - (hs_0.pos.x * hs_0.vel.y - hs_2.pos.x * hs_2.vel.y),
                (hs_0.pos.z * hs_0.vel.x - hs_1.pos.z * hs_1.vel.x)
                - (hs_0.pos.x * hs_0.vel.z - hs_1.pos.x * hs_1.vel.z),
                (hs_0.pos.z * hs_0.vel.x - hs_2.pos.z * hs_2.vel.x)
                - (hs_0.pos.x * hs_0.vel.z - hs_2.pos.x * hs_2.vel.z),
                (hs_0.pos.z * hs_0.vel.y - hs_1.pos.z * hs_1.vel.y)
                - (hs_0.pos.y * hs_0.vel.z - hs_1.pos.y * hs_1.vel.z),
                (hs_0.pos.z * hs_0.vel.y - hs_2.pos.z * hs_2.vel.y)
                - (hs_0.pos.y * hs_0.vel.z - hs_2.pos.y * hs_2.vel.z),
            ],
        )
        c = a.inv() * b
        return c[0] + c[1] + c[2]


def main():
    sol = Day24(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
