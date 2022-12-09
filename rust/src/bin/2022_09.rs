use std::collections::HashSet;

use aoc::solution::Solution;

enum Move {
    Left(usize),
    Up(usize),
    Right(usize),
    Down(usize),
}

impl Move {
    fn parse(input: &str) -> Move {
        let mut s = input.trim().split(" ");
        let m = s.next().unwrap();
        let n: usize = s.next().unwrap().parse().unwrap();
        match m {
            "L" => Move::Left(n),
            "U" => Move::Up(n),
            "R" => Move::Right(n),
            "D" => Move::Down(n),
            _ => panic!("wrong move"),
        }
    }

    fn dist(&self) -> usize {
        match self {
            Move::Left(n) => *n,
            Move::Up(n) => *n,
            Move::Right(n) => *n,
            Move::Down(n) => *n,
        }
    }

    fn step(&self, pos: (isize, isize)) -> impl Iterator<Item = (isize, isize)> + '_ {
        let (x, y) = pos;
        let dist = self.dist();
        (1..=dist).map(move |i| match self {
            Move::Left(_) => (x - i as isize, y),
            Move::Up(_) => (x, y + i as isize),
            Move::Right(_) => (x + i as isize, y),
            Move::Down(_) => (x, y - i as isize),
        })
    }
}

fn move_tail(pos_h: &(isize, isize), pos_t: &(isize, isize)) -> (isize, isize) {
    let abs_0 = (pos_h.0 - pos_t.0).abs();
    let abs_1 = (pos_h.1 - pos_t.1).abs();
    if abs_0 <= 1 && abs_1 <= 1 {
        // Don't move tail if head is near
        (pos_t.0, pos_t.1)
    } else if abs_0 == 2 && abs_1 == 2 {
        // Diagonal move
        let d0 = (pos_h.0 - pos_t.0) / 2;
        let d1 = (pos_h.1 - pos_t.1) / 2;
        (pos_t.0 + d0, pos_t.1 + d1)
    } else if pos_h.0 - pos_t.0 == 2 {
        // If tail is on the left side
        (pos_h.0 - 1, pos_h.1)
    } else if pos_t.0 - pos_h.0 == 2 {
        // If tail is on the right side
        (pos_h.0 + 1, pos_h.1)
    } else if pos_h.1 - pos_t.1 == 2 {
        // If tail is on the bottom side
        (pos_h.0, pos_h.1 - 1)
    } else if pos_t.1 - pos_h.1 == 2 {
        // If tail is on the top side
        (pos_h.0, pos_h.1 + 1)
    } else {
        panic!("tail is too far")
    }
}

struct Day2022_09 {
    moves: Vec<Move>,
}

impl Solution<usize> for Day2022_09 {
    fn new() -> Day2022_09 {
        Day2022_09 { moves: Vec::new() }
    }

    fn init(&mut self, input: &str) {
        for line in input.lines() {
            self.moves.push(Move::parse(line));
        }
    }

    fn part_one(&mut self) -> usize {
        let mut visited = HashSet::<(isize, isize)>::from([(0, 0)]);
        let mut pos_h: (isize, isize) = (0, 0);
        let mut pos_t: (isize, isize) = (0, 0);
        for mov in self.moves.iter() {
            for next_pos_h in mov.step(pos_h) {
                pos_t = move_tail(&next_pos_h, &pos_t);
                visited.insert(pos_t);
                pos_h = next_pos_h;
            }
        }
        visited.len()
    }

    fn part_two(&mut self) -> usize {
        let mut visited = HashSet::<(isize, isize)>::from([(0, 0)]);
        let mut pos_h: (isize, isize) = (0, 0);
        let mut pos_ts: Vec<(isize, isize)> = vec![(0, 0); 9];
        for mov in self.moves.iter() {
            for next_pos_h in mov.step(pos_h) {
                for i in 0..9 {
                    pos_ts[i] = move_tail(
                        if i == 0 { &next_pos_h } else { &pos_ts[i - 1] },
                        &pos_ts[i],
                    );
                }
                pos_h = next_pos_h;
                visited.insert(pos_ts[8]);
            }
        }
        visited.len()
    }
}

fn main() {
    let mut sol = Day2022_09::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_09;
    use aoc::solution::Solution;

    const TEST_INPUT_1: &str = include_str!("../../examples/2022_09_1.txt");
    const TEST_INPUT_2: &str = include_str!("../../examples/2022_09_2.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_09::new();
        sol.init(TEST_INPUT_1);
        assert_eq!(sol.part_one(), 13);
        assert_eq!(sol.part_two(), 1);
    }

    #[test]
    fn test_2() {
        let mut sol = Day2022_09::new();
        sol.init(TEST_INPUT_2);
        assert_eq!(sol.part_one(), 88);
        assert_eq!(sol.part_two(), 36);
    }
}
