from networkx import Graph, minimum_cut
from typing import TextIO
from sys import stdin


class Day25:
    graph: Graph

    def __init__(self, input: TextIO):
        self.graph = Graph()
        for line in input.readlines():
            start, target_str = line.strip().split(": ")
            targets = target_str.split(" ")
            for target in targets:
                self.graph.add_edge(start, target, capacity=1)

    def part1(self) -> int:
        nodes = list(self.graph.nodes)
        start = nodes[0]
        for end in nodes[1:]:
            min_cut, parts = minimum_cut(self.graph, start, end)
            if min_cut == 3:
                return len(parts[0]) * len(parts[1])
        return 0

    def part2(self) -> str:
        return "Merry Christmas! 🎄"


def main():
    sol = Day25(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
