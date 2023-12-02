from dataclasses import dataclass
from typing import TextIO
from sys import stdin


@dataclass
class Reveal:
    red: int
    green: int
    blue: int


class Day02:
    games: dict[int, list[Reveal]]

    def __init__(self, input: TextIO):
        self.games = {}
        for line in input.readlines():
            game = []
            [name, reveals] = line.split(":")
            for part in reveals.split(";"):
                red = 0
                green = 0
                blue = 0
                for subpart in part.split(","):
                    [count, color] = subpart.strip().split(" ")
                    if color == "red":
                        red = int(count)
                    elif color == "green":
                        green = int(count)
                    elif color == "blue":
                        blue = int(count)
                game.append(Reveal(red, green, blue))
            self.games[int(name.split(" ")[1])] = game

    def part1(self) -> int:
        goal = Reveal(12, 13, 14)
        return sum(
            idx
            for idx, game in self.games.items()
            if all(
                r.red <= goal.red and r.green <= goal.green and r.blue <= goal.blue
                for r in game
            )
        )

    def part2(self) -> int:
        return sum(
            res.red * res.green * res.blue
            for res in (
                Reveal(
                    max(rev.red for rev in game),
                    max(rev.green for rev in game),
                    max(rev.blue for rev in game),
                )
                for game in self.games.values()
            )
        )


def main():
    sol = Day02(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
