from collections import deque
from dataclasses import dataclass
from math import lcm
from enum import Enum
from typing import TextIO
from sys import stdin


class Pulse(Enum):
    LOW = 0
    HIGH = 1


@dataclass
class Module:
    type: str
    targets: list[str]


class Day20:
    modules: dict[str, Module]
    inputs: dict[str, list[str]]

    def __init__(self, input: TextIO):
        self.modules = {}
        self.inputs = {}
        for line in input.readlines():
            name_str, targets_str = line.strip().split(" -> ")
            targets = targets_str.split(", ")
            type = name_str[0]
            name = name_str[1:] if type in "%&" else name_str
            if type == "%":
                self.modules[name] = Module("FLIP-FLOP", targets)
            elif type == "&":
                self.modules[name] = Module("CONJ", targets)
            else:
                self.modules[name] = Module("BROADCASTER", targets)
            for target in targets:
                if target not in self.inputs:
                    self.inputs[target] = []
                self.inputs[target].append(name)

    def press_button(
        self, flip_flop_state: dict[str, bool], conj_state: dict[str, dict[str, Pulse]]
    ) -> dict[tuple[str, Pulse], int]:
        queue = deque[tuple[str, str, Pulse]]([("broadcaster", "button", Pulse.LOW)])
        res: dict[tuple[str, Pulse], int] = {}
        while queue:
            name, source, pulse = queue.popleft()
            res[(name, pulse)] = res.get((name, pulse), 0) + 1
            if name not in self.modules:
                continue
            mod = self.modules[name]
            if mod.type == "FLIP-FLOP":
                if pulse == Pulse.LOW:
                    queue.extend(
                        (
                            target,
                            name,
                            Pulse.LOW if flip_flop_state[name] else Pulse.HIGH,
                        )
                        for target in mod.targets
                    )
                    flip_flop_state[name] = not flip_flop_state[name]
            elif mod.type == "CONJ":
                conj_state[name][source] = pulse
                if all(pulse == Pulse.HIGH for pulse in conj_state[name].values()):
                    queue.extend((target, name, Pulse.LOW) for target in mod.targets)
                elif any(pulse == Pulse.LOW for pulse in conj_state[name].values()):
                    queue.extend((target, name, Pulse.HIGH) for target in mod.targets)
            else:
                queue.extend((target, name, pulse) for target in mod.targets)
        return res

    def part1(self) -> int:
        pulse_counter = {Pulse.LOW: 0, Pulse.HIGH: 0}
        flip_flop_state: dict[str, bool] = {
            name: False for name, mod in self.modules.items() if mod.type == "FLIP-FLOP"
        }
        conj_state: dict[str, dict[str, Pulse]] = {
            name: {input: Pulse.LOW for input in inputs}
            for name, inputs in self.inputs.items()
            if name in self.modules and self.modules[name].type == "CONJ"
        }
        for _ in range(1000):
            for (_, pulse), counter in self.press_button(
                flip_flop_state, conj_state
            ).items():
                pulse_counter[pulse] += counter
        return pulse_counter[Pulse.HIGH] * pulse_counter[Pulse.LOW]

    def part2(self) -> int:
        # Backtrack to find modules that contribute to the target one
        target_modules: list[set[str]] = []
        queue = deque([(0, "rx")])
        while queue:
            i, name = queue.popleft()
            if len(target_modules) <= i:
                target_modules.append(set())
            target_modules[i].add(name)
            for input in self.inputs[name]:
                if self.modules[input].type == "CONJ":
                    queue.append((i + 1, input))

        # Choosing this arbitrarily - seems that third layer is good enough 🤷
        target_modules_idx = 2
        waiting_signal = Pulse.LOW if target_modules_idx % 2 == 0 else Pulse.HIGH

        # Find cycles of the pulses for target modules
        flip_flop_state: dict[str, bool] = {
            name: False for name, mod in self.modules.items() if mod.type == "FLIP-FLOP"
        }
        conj_state: dict[str, dict[str, Pulse]] = {
            name: {input: Pulse.LOW for input in inputs}
            for name, inputs in self.inputs.items()
            if name in self.modules and self.modules[name].type == "CONJ"
        }
        counters: dict[str, list[int]] = {
            name: [] for name in target_modules[target_modules_idx]
        }
        for i in range(100000):
            for (name, pulse), _ in self.press_button(
                flip_flop_state, conj_state
            ).items():
                if pulse == waiting_signal and name in counters:
                    counters[name].append(i)
            if all(
                len(counter) > 2
                and (counter[-1] - counter[-2]) == (counter[-2] - counter[-3])
                for counter in counters.values()
            ):
                break
        return lcm(*[counter[-1] - counter[-2] for counter in counters.values()])


def main():
    sol = Day20(stdin)
    print(f"Part 1: {sol.part1()}")
    print(f"Part 2: {sol.part2()}")


if __name__ == "__main__":
    main()
