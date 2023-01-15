use std::collections::{HashMap, HashSet};

use aoc::solution::Solution;

type Pos = (i32, i32);

type State = HashMap<Pos, Vec<(usize, Direction)>>;

#[derive(Clone, Copy)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Direction {
    fn from_char(c: char) -> Option<Direction> {
        match c {
            '^' => Some(Direction::Up),
            'v' => Some(Direction::Down),
            '>' => Some(Direction::Right),
            '<' => Some(Direction::Left),
            _ => None,
        }
    }
}

fn get_next_state(state: &State, width: i32, height: i32) -> State {
    let mut res = State::new();
    for (&(row, col), winds) in state.iter() {
        for &(idx, dir) in winds.iter() {
            let next = match dir {
                Direction::Up => ((row + height - 1) % height, col),
                Direction::Down => ((row + 1) % height, col),
                Direction::Left => (row, (col + width - 1) % width),
                Direction::Right => (row, (col + 1) % width),
            };
            res.entry(next)
                .and_modify(|v| v.push((idx, dir)))
                .or_insert(vec![(idx, dir)]);
        }
    }
    res
}

fn travel_forward(
    states: &Vec<State>,
    width: i32,
    height: i32,
    pos: Pos,
    steps: i32,
    max_steps: &mut i32,
    visited: &mut HashSet<(Pos, i32)>,
) {
    if steps >= *max_steps || visited.contains(&(pos, steps)) {
        return;
    }

    if pos == (height - 1, width - 1) && (steps + 1) < *max_steps {
        *max_steps = steps + 1;
        return;
    }

    visited.insert((pos, steps));

    let cur_state = &states[(steps + 1) as usize];
    let (row, col) = pos;

    let next = if row == -1 {
        vec![(0, 0), (-1, 0)]
    } else {
        vec![
            (row, col + 1),
            (row + 1, col),
            (row - 1, col),
            (row, col - 1),
            (row, col),
        ]
    };
    next.iter()
        .filter(|(r, c)| {
            (row == -1 && (*r, *c) == (-1, 0))
                || (*r >= 0
                    && *c >= 0
                    && *r < height
                    && *c < width
                    && !cur_state.contains_key(&(*r, *c)))
        })
        .for_each(|&pos| travel_forward(states, width, height, pos, steps + 1, max_steps, visited));
}

fn travel_backward(
    states: &Vec<State>,
    width: i32,
    height: i32,
    pos: Pos,
    steps: i32,
    max_steps: &mut i32,
    visited: &mut HashSet<(Pos, i32)>,
) {
    if steps >= *max_steps || visited.contains(&(pos, steps)) {
        return;
    }

    if pos == (0, 0) && (steps + 1) < *max_steps {
        *max_steps = steps + 1;
        return;
    }

    visited.insert((pos, steps));

    let cur_state = &states[(steps + 1) as usize];
    let (row, col) = pos;

    let next = if row == height {
        vec![(height - 1, width - 1), (height, width - 1)]
    } else {
        vec![
            (row, col - 1),
            (row - 1, col),
            (row, col + 1),
            (row + 1, col),
            (row, col),
        ]
    };
    next.iter()
        .filter(|(r, c)| {
            (row == height && (*r, *c) == (height, width - 1))
                || (*r >= 0
                    && *c >= 0
                    && *r < height
                    && *c < width
                    && !cur_state.contains_key(&(*r, *c)))
        })
        .for_each(|&pos| {
            travel_backward(states, width, height, pos, steps + 1, max_steps, visited)
        });
}

struct Day2022_24 {
    width: i32,
    height: i32,
    states: Vec<State>,
    part_one_res: i32,
}

impl Solution<i32> for Day2022_24 {
    fn new(input: &str) -> Day2022_24 {
        let map: Vec<_> = input.lines().map(|s| s.trim()).collect();
        let mut idx = 0;
        let mut state: State = State::new();
        let mut states = vec![];
        for (row, &line) in map[1..map.len() - 1].iter().enumerate() {
            for (col, c) in line.chars().skip(1).enumerate() {
                if let Some(dir) = Direction::from_char(c) {
                    state.insert((row as i32, col as i32), vec![(idx, dir)]);
                    idx += 1;
                }
            }
        }
        let width = map[0].len() as i32 - 2;
        let height = map.len() as i32 - 2;
        states.push(state);

        // Precalculte states
        for _ in 0..(width * height * 3) {
            let next_state = get_next_state(&states.last().unwrap(), width, height);
            states.push(next_state);
        }

        Day2022_24 {
            width,
            height,
            states,
            part_one_res: 0,
        }
    }

    fn part_one(&mut self) -> i32 {
        let mut res = self.width * self.height * 3;
        let mut visited = HashSet::new();
        travel_forward(
            &self.states,
            self.width,
            self.height,
            (-1, 0),
            0,
            &mut res,
            &mut visited,
        );
        self.part_one_res = res;
        res
    }

    fn part_two(&mut self) -> i32 {
        let mut visited = HashSet::new();

        let mut steps_backward = self.width * self.height * 10;
        travel_backward(
            &self.states,
            self.width,
            self.height,
            (self.height, self.width - 1),
            self.part_one_res,
            &mut steps_backward,
            &mut visited,
        );

        visited = HashSet::new();
        let mut res = self.width * self.height * 10;
        travel_forward(
            &self.states,
            self.width,
            self.height,
            (-1, 0),
            steps_backward,
            &mut res,
            &mut visited,
        );

        res
    }
}

fn main() {
    Day2022_24::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_24;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_24.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_24::new(TEST_INPUT);
        assert_eq!(sol.part_one(), 18);
        assert_eq!(sol.part_two(), 54);
    }
}
