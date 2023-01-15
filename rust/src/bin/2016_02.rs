use aoc::solution::Solution;

#[derive(Clone, Copy)]
enum Move {
    Left,
    Right,
    Up,
    Down,
}

impl Move {
    fn apply(self, pos: char) -> char {
        match self {
            Move::Left => match pos {
                '1' | '4' | '7' => pos,
                _ => char::from_u32(pos as u32 - 1).unwrap(),
            },
            Move::Right => match pos {
                '3' | '6' | '9' => pos,
                _ => char::from_u32(pos as u32 + 1).unwrap(),
            },
            Move::Up => match pos {
                '1'..='3' => pos,
                _ => char::from_u32(pos as u32 - 3).unwrap(),
            },
            Move::Down => match pos {
                '7'..='9' => pos,
                _ => char::from_u32(pos as u32 + 3).unwrap(),
            },
        }
    }

    fn apply_ext(self, pos: char) -> char {
        match self {
            Move::Left => match pos {
                '3' => '2',
                '4' => '3',
                '6' => '5',
                '7' => '6',
                '8' => '7',
                '9' => '8',
                'B' => 'A',
                'C' => 'B',
                _ => pos,
            },
            Move::Right => match pos {
                '2' => '3',
                '3' => '4',
                '5' => '6',
                '6' => '7',
                '7' => '8',
                '8' => '8',
                'A' => 'B',
                'B' => 'C',
                _ => pos,
            },
            Move::Up => match pos {
                '3' => '1',
                '6' => '2',
                '7' => '3',
                '8' => '4',
                'A' => '6',
                'B' => '7',
                'C' => '8',
                'D' => 'B',
                _ => pos,
            },
            Move::Down => match pos {
                '1' => '3',
                '2' => '6',
                '3' => '7',
                '4' => '8',
                '6' => 'A',
                '7' => 'B',
                '8' => 'C',
                'B' => 'D',
                _ => pos,
            },
        }
    }

    fn from_char(c: char) -> Self {
        match c {
            'L' => Move::Left,
            'R' => Move::Right,
            'U' => Move::Up,
            'D' => Move::Down,
            _ => panic!("invalid character"),
        }
    }
}

struct Day2016_02 {
    instructions: Vec<Vec<Move>>,
}

impl Solution<String> for Day2016_02 {
    fn new(input: &str) -> Day2016_02 {
        Day2016_02 {
            instructions: input
                .lines()
                .map(|line| line.trim().chars().map(Move::from_char).collect())
                .collect(),
        }
    }

    fn part_one(&mut self) -> String {
        let mut pos = '5';
        self.instructions
            .iter()
            .map(|moves| {
                let mut cur_pos = pos;
                for mov in moves {
                    cur_pos = mov.apply(cur_pos);
                }
                pos = cur_pos;
                cur_pos.to_string()
            })
            .collect()
    }

    fn part_two(&mut self) -> String {
        let mut pos = '5';
        self.instructions
            .iter()
            .map(|moves| {
                let mut cur_pos = pos;
                for mov in moves {
                    cur_pos = mov.apply_ext(cur_pos);
                }
                pos = cur_pos;
                cur_pos.to_string()
            })
            .collect()
    }
}

fn main() {
    Day2016_02::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2016_02;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2016_02.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2016_02::new(TEST_INPUT);
        assert_eq!(sol.part_one(), "1985");
        assert_eq!(sol.part_two(), "5DB3");
    }
}
