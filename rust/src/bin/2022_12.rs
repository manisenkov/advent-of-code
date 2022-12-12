use std::cmp::Ordering;
use std::collections::{BinaryHeap, HashMap};

use aoc::solution::Solution;

#[derive(Clone, Copy, PartialEq, Eq)]
struct State {
    pos: (isize, isize),
    dist: usize,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        other
            .dist
            .cmp(&self.dist)
            .then_with(|| self.pos.cmp(&other.pos))
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

struct Day2022_12 {
    dist_map: HashMap<(isize, isize), usize>,
    height_map: HashMap<(isize, isize), usize>,
    possible_starts: Vec<(isize, isize)>,
    start_point: (isize, isize),
    end_point: (isize, isize),
}

impl Solution<usize> for Day2022_12 {
    fn new() -> Day2022_12 {
        Day2022_12 {
            dist_map: HashMap::new(),
            height_map: HashMap::new(),
            possible_starts: vec![],
            start_point: (0, 0),
            end_point: (0, 0),
        }
    }

    fn init(&mut self, input: &str) {
        for (row, line) in input.lines().enumerate() {
            self.height_map
                .extend(line.trim().chars().enumerate().map(|(col, c)| {
                    let pos = (row as isize, col as isize);
                    let height = if c == 'S' {
                        self.start_point = pos;
                        1
                    } else if c == 'E' {
                        self.end_point = pos;
                        26
                    } else {
                        (c as usize) - ('a' as usize) + 1
                    };
                    if c == 'a' || c == 'S' {
                        self.possible_starts.push(pos);
                    }
                    (pos, height)
                }));
        }
    }

    fn part_one(&mut self) -> usize {
        let mut states = BinaryHeap::<State>::new();

        self.dist_map.insert(self.end_point, 0);
        states.push(State {
            dist: 0,
            pos: self.end_point,
        });

        while states.len() > 0 {
            let cur_state = states.pop().unwrap();
            let cur_height = *self.height_map.get(&cur_state.pos).unwrap();
            let to_visit: Vec<(isize, isize)> = vec![
                (cur_state.pos.0 - 1, cur_state.pos.1),
                (cur_state.pos.0 + 1, cur_state.pos.1),
                (cur_state.pos.0, cur_state.pos.1 - 1),
                (cur_state.pos.0, cur_state.pos.1 + 1),
            ]
            .iter()
            .filter(|p| self.height_map.contains_key(p))
            .filter(|p| self.height_map.get(p).unwrap() + 1 >= cur_height)
            .map(|p| *p)
            .collect();

            if let Some(d) = self.dist_map.get(&cur_state.pos) {
                if cur_state.dist > *d {
                    continue;
                }
            }

            for target_pos in to_visit.iter() {
                let next = State {
                    dist: cur_state.dist + 1,
                    pos: *target_pos,
                };
                let target_dist = *self.dist_map.get(&target_pos).unwrap_or(&0x7fffffff);
                if next.dist < target_dist {
                    self.dist_map.insert(*target_pos, next.dist);
                    states.push(next);
                }
            }
        }
        *self.dist_map.get(&self.start_point).unwrap()
    }

    fn part_two(&mut self) -> usize {
        self.possible_starts
            .iter()
            .map(|p| *self.dist_map.get(p).unwrap_or(&0x7fffffff))
            .min()
            .unwrap()
    }
}

fn main() {
    let mut sol = Day2022_12::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_12;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_12.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_12::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 31);
        assert_eq!(sol.part_two(), 29);
    }
}
