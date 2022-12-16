use std::collections::{HashMap, VecDeque};

use aoc::solution::Solution;

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
        let mut queue = VecDeque::<(isize, isize)>::from([self.end_point]);

        self.dist_map.insert(self.end_point, 0);

        while let Some(cur_pos) = queue.pop_front() {
            let cur_dist = self.dist_map[&cur_pos];
            let cur_height = self.height_map[&cur_pos];
            let to_visit: Vec<(isize, isize)> = vec![
                (cur_pos.0 - 1, cur_pos.1),
                (cur_pos.0 + 1, cur_pos.1),
                (cur_pos.0, cur_pos.1 - 1),
                (cur_pos.0, cur_pos.1 + 1),
            ]
            .iter()
            .filter(|p| self.height_map.contains_key(p))
            .filter(|p| self.height_map[p] + 1 >= cur_height)
            .map(|p| *p)
            .collect();

            for target_pos in to_visit.iter() {
                let next_dist = cur_dist + 1;
                let target_dist = *self.dist_map.get(&target_pos).unwrap_or(&0x7fffffff);
                if next_dist < target_dist {
                    self.dist_map.insert(*target_pos, next_dist);
                    queue.push_back(*target_pos);
                }
            }
        }
        self.dist_map[&self.start_point]
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
