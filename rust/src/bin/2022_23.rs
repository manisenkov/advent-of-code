use std::collections::{HashMap, HashSet};

use aoc::solution::Solution;

fn mov(pos: &HashSet<(i32, i32)>, round: usize) -> HashSet<(i32, i32)> {
    let mut proposals: HashMap<(i32, i32), Vec<(i32, i32)>> = HashMap::new();
    for &(row, col) in pos.iter() {
        if [
            (row - 1, col - 1),
            (row - 1, col),
            (row - 1, col + 1),
            (row, col - 1),
            (row, col + 1),
            (row + 1, col - 1),
            (row + 1, col),
            (row + 1, col + 1),
        ]
        .iter()
        .all(|p| !pos.contains(p))
        {
            continue;
        }
        let next_positions = &[
            (
                ((row - 1, col - 1), (row - 1, col), (row - 1, col + 1)),
                (row - 1, col),
            ), // North
            (
                ((row + 1, col - 1), (row + 1, col), (row + 1, col + 1)),
                (row + 1, col),
            ), // South
            (
                ((row - 1, col - 1), (row, col - 1), (row + 1, col - 1)),
                (row, col - 1),
            ), // West
            (
                ((row - 1, col + 1), (row, col + 1), (row + 1, col + 1)),
                (row, col + 1),
            ), // East
        ];

        for i in 0..4 {
            let (check_positions, next_pos) = next_positions[(i + round) % next_positions.len()];
            if !(pos.contains(&check_positions.0)
                || pos.contains(&check_positions.1)
                || pos.contains(&check_positions.2))
            {
                proposals
                    .entry(next_pos)
                    .and_modify(|v| v.push((row, col)))
                    .or_insert(vec![(row, col)]);
                break;
            }
        }
    }
    proposals.retain(|_, v| v.len() == 1);
    let mut res: HashSet<(i32, i32)> = HashSet::from_iter(pos.iter().map(|p| *p));
    for (next_pos, elves) in proposals.iter() {
        res.remove(&elves[0]);
        res.insert(*next_pos);
    }
    res
}

struct Day2022_23 {
    init_pos: HashSet<(i32, i32)>,
}

impl Solution<i32, usize> for Day2022_23 {
    fn new() -> Day2022_23 {
        Day2022_23 {
            init_pos: HashSet::new(),
        }
    }

    fn init(&mut self, input: &str) {
        for (row, line) in input.lines().enumerate() {
            for (col, c) in line.trim().chars().enumerate() {
                if c == '#' {
                    self.init_pos.insert((row as i32, col as i32));
                }
            }
        }
    }

    fn part_one(&mut self) -> i32 {
        let mut cur_pos: HashSet<(i32, i32)> = HashSet::from_iter(self.init_pos.iter().map(|p| *p));
        for round in 0..10 {
            cur_pos = mov(&cur_pos, round);
        }
        let min = (
            cur_pos.iter().map(|(row, _)| *row).min().unwrap(),
            cur_pos.iter().map(|(_, col)| *col).min().unwrap(),
        );
        let max = (
            cur_pos.iter().map(|(row, _)| *row).max().unwrap(),
            cur_pos.iter().map(|(_, col)| *col).max().unwrap(),
        );
        (max.0 - min.0 + 1) * (max.1 - min.1 + 1) - cur_pos.len() as i32
    }

    fn part_two(&mut self) -> usize {
        let mut cur_pos: HashSet<(i32, i32)> = HashSet::from_iter(self.init_pos.iter().map(|p| *p));
        for round in 0.. {
            let next_pos = mov(&cur_pos, round);
            if cur_pos
                .intersection(&next_pos)
                .collect::<Vec<&(i32, i32)>>()
                .len()
                == cur_pos.len()
            {
                return round + 1;
            }
            cur_pos = next_pos;
        }
        0
    }
}

fn main() {
    let mut sol = Day2022_23::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_23;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_23.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_23::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 110);
        assert_eq!(sol.part_two(), 20);
    }
}
