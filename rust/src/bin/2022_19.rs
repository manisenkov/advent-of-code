use std::collections::HashMap;
use std::sync::mpsc;
use std::thread;

use aoc::solution::Solution;

const ROBOT_CAP: u32 = 10;

#[derive(Clone, Copy)]
enum RobotType {
    Ore,
    Clay,
    Obsidian,
    Geode,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
struct Blueprint {
    ore: u32,
    clay: u32,
    obisidian: (u32, u32), // ore and clay
    geode: (u32, u32),     // ore and obsidian
}

impl Blueprint {
    fn parse(input: &str) -> Self {
        let mut parts = input.split(". ");
        let ore_parts: Vec<&str> = parts.next().unwrap().split(" ").collect();
        let ore: u32 = ore_parts[6].parse().unwrap();
        let clay_parts: Vec<&str> = parts.next().unwrap().split(" ").collect();
        let clay: u32 = clay_parts[4].parse().unwrap();
        let obsidian_parts: Vec<&str> = parts.next().unwrap().split(" ").collect();
        let obsidian_ore: u32 = obsidian_parts[4].parse().unwrap();
        let obsidian_clay: u32 = obsidian_parts[7].parse().unwrap();
        let geode_parts: Vec<&str> = parts.next().unwrap().split(" ").collect();
        let geode_ore: u32 = geode_parts[4].parse().unwrap();
        let geode_obsidian: u32 = geode_parts[7].parse().unwrap();
        Blueprint {
            ore: ore,
            clay: clay,
            obisidian: (obsidian_ore, obsidian_clay),
            geode: (geode_ore, geode_obsidian),
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
struct State {
    ore: u32,
    clay: u32,
    obsidian: u32,
    ore_robot: u32,
    clay_robot: u32,
    obsidian_robot: u32,
}

impl State {
    fn new() -> Self {
        State {
            ore: 0,
            clay: 0,
            obsidian: 0,
            ore_robot: 1,
            clay_robot: 0,
            obsidian_robot: 0,
        }
    }

    fn process(&self, robot_to_make: Option<RobotType>, blueprint: &Blueprint) -> Option<State> {
        let mut new_state = State { ..*self };
        if let Some(robot_type) = robot_to_make {
            match new_state.make_robot(robot_type, blueprint) {
                Some(next) => new_state = next,
                None => return None,
            }
        }
        Some(State {
            ore: new_state.ore + self.ore_robot,
            clay: new_state.clay + self.clay_robot,
            obsidian: new_state.obsidian + self.obsidian_robot,
            ..new_state
        })
    }

    fn make_robot(&self, robot_type: RobotType, blueprint: &Blueprint) -> Option<Self> {
        match robot_type {
            RobotType::Ore => {
                if self.ore_robot >= ROBOT_CAP || self.ore < blueprint.ore {
                    None
                } else {
                    Some(State {
                        ore: self.ore - blueprint.ore,
                        ore_robot: self.ore_robot + 1,
                        ..*self
                    })
                }
            }
            RobotType::Clay => {
                if self.clay_robot >= ROBOT_CAP || self.ore < blueprint.clay {
                    None
                } else {
                    Some(State {
                        ore: self.ore - blueprint.clay,
                        clay_robot: self.clay_robot + 1,
                        ..*self
                    })
                }
            }
            RobotType::Obsidian => {
                if self.obsidian_robot >= ROBOT_CAP
                    || self.ore < blueprint.obisidian.0
                    || self.clay < blueprint.obisidian.1
                {
                    None
                } else {
                    Some(State {
                        ore: self.ore - blueprint.obisidian.0,
                        clay: self.clay - blueprint.obisidian.1,
                        obsidian_robot: self.obsidian_robot + 1,
                        ..*self
                    })
                }
            }
            RobotType::Geode => {
                if self.ore < blueprint.geode.0 || self.obsidian < blueprint.geode.1 {
                    None
                } else {
                    Some(State {
                        ore: self.ore - blueprint.geode.0,
                        obsidian: self.obsidian - blueprint.geode.1,
                        ..*self
                    })
                }
            }
        }
    }

    fn max_geode(
        &self,
        time_left: u32,
        blueprint: &Blueprint,
        cache: &mut HashMap<(u32, State), u32>,
    ) -> u32 {
        if let Some(res) = cache.get(&(time_left, *self)) {
            return *res;
        }

        if time_left == 0 {
            return 0;
        }

        let res = if let Some(next) = self.process(Some(RobotType::Geode), blueprint) {
            (time_left - 1) + next.max_geode(time_left - 1, blueprint, cache)
        } else {
            self.process(None, blueprint)
                .unwrap()
                .max_geode(time_left - 1, blueprint, cache)
                .max(
                    [RobotType::Ore, RobotType::Clay, RobotType::Obsidian]
                        .iter()
                        .map(|robot_type| {
                            if let Some(next_state) = self.process(Some(*robot_type), blueprint) {
                                next_state.max_geode(time_left - 1, blueprint, cache)
                            } else {
                                0
                            }
                        })
                        .max()
                        .unwrap(),
                )
        };
        cache.insert((time_left, *self), res);
        res
    }
}

struct Day2022_19 {
    blueprints: Vec<Blueprint>,
}

impl Solution<u32> for Day2022_19 {
    fn new() -> Day2022_19 {
        Day2022_19 { blueprints: vec![] }
    }

    fn init(&mut self, input: &str) {
        self.blueprints.extend(input.lines().map(Blueprint::parse));
    }

    fn part_one(&mut self) -> u32 {
        let (tx, rx) = mpsc::channel();
        for (i, blueprint) in self.blueprints.iter().enumerate() {
            let bp = blueprint.clone();
            let txp = tx.clone();
            println!("{} started", i + 1);
            thread::spawn(move || {
                let initial_state = State::new();
                let mut mem = HashMap::<(u32, State), u32>::new();
                txp.send((i as u32, initial_state.max_geode(24, &bp, &mut mem)))
                    .unwrap();
            });
        }
        let mut counter = self.blueprints.len();
        let mut res = 0;
        for (i, max_geode) in rx {
            res += (i + 1) * max_geode;
            counter -= 1;
            println!("{} done, {} left", i + 1, counter);
            if counter == 0 {
                break;
            }
        }
        res
    }

    fn part_two(&mut self) -> u32 {
        let (tx, rx) = mpsc::channel();
        for (i, blueprint) in self.blueprints.iter().take(3).enumerate() {
            let bp = blueprint.clone();
            let txp = tx.clone();
            println!("{} started", i + 1);
            thread::spawn(move || {
                let initial_state = State::new();
                let mut mem = HashMap::<(u32, State), u32>::new();
                txp.send(initial_state.max_geode(32, &bp, &mut mem))
                    .unwrap();
            });
        }
        let mut counter = 3.min(self.blueprints.len());
        let mut res = 1;
        for max_geode in rx {
            res *= max_geode;
            counter -= 1;
            println!("{} left", counter);
            if counter == 0 {
                break;
            }
        }
        res
    }
}

fn main() {
    let mut sol = Day2022_19::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_19;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_19.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_19::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 33);
        assert_eq!(sol.part_two(), 3472);
    }
}
