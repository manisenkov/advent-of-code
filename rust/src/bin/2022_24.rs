use std::collections::{HashMap, HashSet};

use aoc::solution::Solution;

type Pos = (i32, i32);

type State = HashMap<Pos, Vec<(usize, Direction)>>;

#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
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

    // Entrance
    if row == -1 {
        // Down
        if !cur_state.contains_key(&(0, 0)) {
            travel_forward(states, width, height, (0, 0), steps + 1, max_steps, visited)
        }

        // Wait
        travel_forward(
            states,
            width,
            height,
            (-1, 0),
            steps + 1,
            max_steps,
            visited,
        )
    } else {
        // Right
        if col < width - 1 && !cur_state.contains_key(&(row, col + 1)) {
            travel_forward(
                states,
                width,
                height,
                (row, col + 1),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Down
        if row < height - 1 && !cur_state.contains_key(&(row + 1, col)) {
            travel_forward(
                states,
                width,
                height,
                (row + 1, col),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Up
        if row > 0 && !cur_state.contains_key(&(row - 1, col)) {
            travel_forward(
                states,
                width,
                height,
                (row - 1, col),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Left
        if col > 0 && !cur_state.contains_key(&(row, col - 1)) {
            travel_forward(
                states,
                width,
                height,
                (row, col - 1),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Wait
        if !cur_state.contains_key(&(row, col)) {
            travel_forward(
                states,
                width,
                height,
                (row, col),
                steps + 1,
                max_steps,
                visited,
            )
        }
    }
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

    // Entrance
    if row == height {
        // Down
        if !cur_state.contains_key(&(height - 1, width - 1)) {
            travel_backward(
                states,
                width,
                height,
                (height - 1, width - 1),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Wait
        travel_backward(
            states,
            width,
            height,
            (height, width - 1),
            steps + 1,
            max_steps,
            visited,
        )
    } else {
        // Left
        if col > 0 && !cur_state.contains_key(&(row, col - 1)) {
            travel_backward(
                states,
                width,
                height,
                (row, col - 1),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Up
        if row > 0 && !cur_state.contains_key(&(row - 1, col)) {
            travel_backward(
                states,
                width,
                height,
                (row - 1, col),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Right
        if col < width - 1 && !cur_state.contains_key(&(row, col + 1)) {
            travel_backward(
                states,
                width,
                height,
                (row, col + 1),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Down
        if row < height - 1 && !cur_state.contains_key(&(row + 1, col)) {
            travel_backward(
                states,
                width,
                height,
                (row + 1, col),
                steps + 1,
                max_steps,
                visited,
            )
        }

        // Wait
        if !cur_state.contains_key(&(row, col)) {
            travel_backward(
                states,
                width,
                height,
                (row, col),
                steps + 1,
                max_steps,
                visited,
            )
        }
    }
}

struct Day2022_24 {
    width: i32,
    height: i32,
    states: Vec<State>,
    part_one_res: i32,
}

impl Solution<i32> for Day2022_24 {
    fn new() -> Day2022_24 {
        Day2022_24 {
            width: 0,
            height: 0,
            states: vec![],
            part_one_res: 0,
        }
    }

    fn init(&mut self, input: &str) {
        let map: Vec<&str> = input.lines().map(|s| s.trim()).collect();
        let mut idx = 0;
        let mut state: State = State::new();
        for (row, &line) in map[1..map.len() - 1].iter().enumerate() {
            for (col, c) in line.chars().skip(1).enumerate() {
                if let Some(dir) = Direction::from_char(c) {
                    state.insert((row as i32, col as i32), vec![(idx, dir)]);
                    idx += 1;
                }
            }
        }
        self.width = map[0].len() as i32 - 2;
        self.height = map.len() as i32 - 2;
        self.states.push(state);

        for _ in 0..(self.width * self.height * 10) {
            let next_state = get_next_state(&self.states.last().unwrap(), self.width, self.height);
            self.states.push(next_state);
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
    let mut sol = Day2022_24::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_24;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_24.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_24::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 18);
        assert_eq!(sol.part_two(), 54);
    }
}
