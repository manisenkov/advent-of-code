use std::collections::{HashMap, VecDeque};

use aoc::solution::Solution;

type Pos = (isize, isize);

struct Day2022_12 {
    dist_map: HashMap<Pos, usize>,
    height_map: HashMap<Pos, usize>,
    possible_starts: Vec<Pos>,
    start_point: Pos,
    end_point: Pos,
}

impl Solution<usize> for Day2022_12 {
    fn new(input: &str) -> Day2022_12 {
        let mut start_point: Pos = (0, 0);
        let mut end_point: Pos = (0, 0);
        let mut possible_starts: Vec<_> = vec![];
        let mut height_map: HashMap<_, _> = HashMap::new();
        for (row, line) in input.lines().enumerate() {
            for (col, c) in line.trim().chars().enumerate() {
                let pos = (row as isize, col as isize);
                let height = if c == 'S' {
                    start_point = pos;
                    1
                } else if c == 'E' {
                    end_point = pos;
                    26
                } else {
                    (c as usize) - ('a' as usize) + 1
                };
                if c == 'a' || c == 'S' {
                    possible_starts.push(pos);
                }
                height_map.insert(pos, height);
            }
        }
        Day2022_12 {
            height_map,
            possible_starts,
            start_point,
            end_point,
            dist_map: HashMap::new(),
        }
    }

    fn part_one(&mut self) -> usize {
        let mut queue = VecDeque::<Pos>::from([self.end_point]);

        self.dist_map.insert(self.end_point, 0);

        while let Some(cur_pos) = queue.pop_front() {
            let cur_dist = self.dist_map[&cur_pos];
            let cur_height = self.height_map[&cur_pos];
            let to_visit: Vec<Pos> = vec![
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
    Day2022_12::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_12;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_12.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_12::new(TEST_INPUT);
        assert_eq!(sol.part_one(), 31);
        assert_eq!(sol.part_two(), 29);
    }
}
