use aoc::solution::Solution;

#[derive(Debug)]
enum Command {
    Step(usize),
    TurnRight,
    TurnLeft,
}

#[derive(Debug, Clone, Copy)]
enum Direction {
    Up,
    Right,
    Down,
    Left,
}

impl Direction {
    fn turn_left(&self) -> Direction {
        match self {
            Direction::Up => Direction::Left,
            Direction::Right => Direction::Up,
            Direction::Down => Direction::Right,
            Direction::Left => Direction::Down,
        }
    }

    fn turn_right(&self) -> Direction {
        match self {
            Direction::Up => Direction::Right,
            Direction::Right => Direction::Down,
            Direction::Down => Direction::Left,
            Direction::Left => Direction::Up,
        }
    }

    fn to_num(&self) -> usize {
        match self {
            Direction::Right => 0,
            Direction::Down => 1,
            Direction::Left => 2,
            Direction::Up => 3,
        }
    }
}

fn mov(map: &Vec<Vec<char>>, start: (usize, usize), dir: Direction, dist: usize) -> (usize, usize) {
    let width = map[0].len();
    let height = map.len();
    let mut pos = start;
    for _ in 0..dist {
        let mut next_pos = pos;
        let is_wall = loop {
            match dir {
                Direction::Up => {
                    next_pos.0 = (next_pos.0 + height - 1) % height;
                }
                Direction::Right => {
                    next_pos.1 = (next_pos.1 + width + 1) % width;
                }
                Direction::Down => {
                    next_pos.0 = (next_pos.0 + height + 1) % height;
                }
                Direction::Left => {
                    next_pos.1 = (next_pos.1 + width - 1) % width;
                }
            }
            if map[next_pos.0][next_pos.1] == '#' {
                break true;
            } else if map[next_pos.0][next_pos.1] == '.' {
                pos = next_pos;
                break false;
            }
        };
        if is_wall {
            break;
        }
    }
    pos
}

struct Day2022_22 {
    map: Vec<Vec<char>>,
    start_point: (usize, usize),
    commands: Vec<Command>,
}

impl Solution<usize> for Day2022_22 {
    fn new() -> Day2022_22 {
        Day2022_22 {
            map: vec![],
            start_point: (0, 0),
            commands: vec![],
        }
    }

    fn init(&mut self, input: &str) {
        let lines: Vec<&str> = input.lines().collect();
        let map_lines = &lines[..lines.len() - 1];
        let mut width: usize = 0;
        for row in map_lines.iter() {
            let row_chars: Vec<char> = row.trim_end().chars().collect();
            width = width.max(row_chars.len());
            self.map.push(row_chars);
        }
        for row in &mut self.map {
            row.extend(vec![' '; width - row.len()]);
        }

        (self.start_point.1, _) = self.map[0]
            .iter()
            .enumerate()
            .find(|(_, c)| **c == '.')
            .unwrap();

        let last_row: Vec<char> = lines.iter().last().unwrap().trim().chars().collect();
        let mut cur_token = "".to_owned();
        let mut idx = 0;
        loop {
            if idx == last_row.len() {
                if !cur_token.is_empty() {
                    self.commands
                        .push(Command::Step(cur_token.parse().unwrap()));
                }
                break;
            }
            let c = last_row[idx];
            match c {
                'R' => {
                    if !cur_token.is_empty() {
                        self.commands
                            .push(Command::Step(cur_token.parse().unwrap()));
                        cur_token = "".to_owned();
                    }
                    self.commands.push(Command::TurnRight);
                }
                'L' => {
                    if !cur_token.is_empty() {
                        self.commands
                            .push(Command::Step(cur_token.parse().unwrap()));
                        cur_token = "".to_owned();
                    }
                    self.commands.push(Command::TurnLeft);
                }
                _ => {
                    cur_token.push(c);
                }
            }
            idx += 1;
        }
    }

    fn part_one(&mut self) -> usize {
        let mut dir = Direction::Right;
        let mut pos = self.start_point;
        for cmd in self.commands.iter() {
            match cmd {
                Command::Step(dist) => {
                    pos = mov(&self.map, pos, dir, *dist);
                }
                Command::TurnLeft => {
                    dir = dir.turn_left();
                }
                Command::TurnRight => {
                    dir = dir.turn_right();
                }
            }
        }
        (pos.0 + 1) * 1000 + (pos.1 + 1) * 4 + dir.to_num()
    }

    fn part_two(&mut self) -> usize {
        0
    }
}

fn main() {
    let mut sol = Day2022_22::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_22;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_22.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_22::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 6032);
        assert_eq!(sol.part_two(), 0);
    }
}
