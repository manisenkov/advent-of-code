use std::collections::HashMap;

use aoc::solution::Solution;

type Pos = (usize, usize);

type PortalMap = HashMap<(Pos, Direction), (Pos, Direction)>;

enum Command {
    Step(usize),
    TurnRight,
    TurnLeft,
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
enum Direction {
    Up,
    Right,
    Down,
    Left,
}

impl Direction {
    fn opposite(&self) -> Direction {
        match self {
            Direction::Up => Direction::Down,
            Direction::Right => Direction::Left,
            Direction::Down => Direction::Up,
            Direction::Left => Direction::Right,
        }
    }

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

fn get_portal_map() -> PortalMap {
    let mut res = HashMap::new();
    let mut ins = |from: Pos, dir_from: Direction, to: Pos, dir_to: Direction| {
        res.insert((from, dir_from), (to, dir_to));
        res.insert((to, dir_to.opposite()), (from, dir_from.opposite()));
    };
    for n in 0..50 {
        // Left 1,1 <-> Top 2,0 corner
        ins((50 + n, 50), Direction::Left, (100, n), Direction::Down);

        // Right 1,1 <-> Bottom 0,2 corner
        ins((99 - n, 99), Direction::Right, (49, 149 - n), Direction::Up);

        // Bottom 2,1 <-> Right 3,0 corner
        ins(
            (149, 99 - n),
            Direction::Down,
            (199 - n, 49),
            Direction::Left,
        );

        // Top 0,1 <-> Left 3,0
        ins((0, 50 + n), Direction::Up, (150 + n, 0), Direction::Right);

        // Top 0,2 <-> Bottom 3,0
        ins((0, 100 + n), Direction::Up, (199, 0 + n), Direction::Up);

        // Left 0,1 <-> Left 2,0 inv
        ins((n, 50), Direction::Left, (149 - n, 0), Direction::Right);

        // Right 0,2 <-> Right 2,1 inv
        ins((n, 149), Direction::Right, (149 - n, 99), Direction::Left);
    }
    res
}

fn mov_simple(map: &Vec<Vec<char>>, start: Pos, dir: Direction, dist: usize) -> Pos {
    let width = map[0].len();
    let height = map.len();
    let mut cur_pos = start;
    for _ in 0..dist {
        let mut next_pos = cur_pos;
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
                cur_pos = next_pos;
                break false;
            }
        };
        if is_wall {
            break;
        }
    }
    cur_pos
}

fn mov_cube(
    map: &Vec<Vec<char>>,
    portals: &PortalMap,
    start: Pos,
    dir: Direction,
    dist: usize,
) -> (Pos, Direction) {
    // let width = map[0].len();
    // let height = map.len();
    let mut cur_pos = start;
    let mut cur_dir = dir;
    for _ in 0..dist {
        let (next_pos, next_dir) = if portals.contains_key(&(cur_pos, cur_dir)) {
            portals[&(cur_pos, cur_dir)]
        } else {
            match cur_dir {
                Direction::Up => ((cur_pos.0 - 1, cur_pos.1), cur_dir),
                Direction::Right => ((cur_pos.0, cur_pos.1 + 1), cur_dir),
                Direction::Down => ((cur_pos.0 + 1, cur_pos.1), cur_dir),
                Direction::Left => ((cur_pos.0, cur_pos.1 - 1), cur_dir),
            }
        };
        if map[next_pos.0][next_pos.1] == '#' {
            break;
        }
        cur_pos = next_pos;
        cur_dir = next_dir;
    }
    (cur_pos, cur_dir)
}

struct Day2022_22 {
    map: Vec<Vec<char>>,
    start_point: Pos,
    commands: Vec<Command>,
    portals: PortalMap,
}

impl Solution<usize> for Day2022_22 {
    fn new() -> Day2022_22 {
        Day2022_22 {
            map: vec![],
            start_point: (0, 0),
            commands: vec![],
            portals: get_portal_map(),
        }
    }

    fn init(&mut self, input: &str) {
        let lines: Vec<&str> = input.lines().collect();
        let map_lines = &lines[..lines.len() - 2];
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
                    pos = mov_simple(&self.map, pos, dir, *dist);
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
        let mut dir = Direction::Right;
        let mut pos = self.start_point;
        for cmd in self.commands.iter() {
            match cmd {
                Command::Step(dist) => {
                    (pos, dir) = mov_cube(&self.map, &self.portals, pos, dir, *dist);
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
}

fn main() {
    let mut sol = Day2022_22::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use std::collections::HashMap;

    use crate::{Day2022_22, Direction, PortalMap, Pos};
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_22.txt");

    fn get_portal_map_example() -> PortalMap {
        let mut res = HashMap::new();
        let mut ins = |from: Pos, dir_from: Direction, to: Pos, dir_to: Direction| {
            res.insert((from, dir_from), (to, dir_to));
            res.insert((to, dir_to.opposite()), (from, dir_from.opposite()));
        };
        for n in 0..4 {
            // Top 0,2 <-> Top 1,0 inv
            ins((0, 8 + n), Direction::Up, (4, 4 - n), Direction::Down);

            // Left 0,2 <-> Top 1,1 corner
            ins((n, 8), Direction::Left, (4, 8 - n), Direction::Down);

            // Right 0,2 <-> Right 2,3 inv
            ins((4 - n, 11), Direction::Right, (8 + n, 15), Direction::Left);

            // Left 1,0 <-> Bottom 2,3 inv
            ins((8 - n, 0), Direction::Left, (12 + n, 11), Direction::Up);

            // Right 1,2 <-> Top 2,3 corner
            ins((8 - n, 11), Direction::Right, (8, 11 + n), Direction::Down);

            // Bottom 1,0 <-> Botton 2,2 inv
            ins((7, 3 - n), Direction::Down, (11, 8 + n), Direction::Up);

            // Bottom 1,1 <-> Left 2,2 corner
            ins((7, 7 - n), Direction::Down, (8 + n, 8), Direction::Left);
        }
        res
    }

    #[test]
    fn test_1() {
        let mut sol = Day2022_22::new();
        sol.init(TEST_INPUT);
        sol.portals = get_portal_map_example();
        assert_eq!(sol.part_one(), 6032);
        assert_eq!(sol.part_two(), 5031);
    }
}
