use std::collections::{HashMap, HashSet};

use aoc::solution::Solution;

type Pos = (i32, i32, i32);

fn get_directions() -> &'static [Pos] {
    &[
        (-1, 0, 0),
        (1, 0, 0),
        (0, -1, 0),
        (0, 1, 0),
        (0, 0, -1),
        (0, 0, 1),
    ]
}

fn calc_sides(cubes: &Vec<Pos>) -> HashMap<Pos, usize> {
    let directions = get_directions();
    let mut sides = HashMap::new();
    for cube in cubes.iter() {
        sides.insert(*cube, 6);
        for dir in directions.iter() {
            let neighbour = (cube.0 + dir.0, cube.1 + dir.1, cube.2 + dir.2);
            if sides.contains_key(&neighbour) {
                sides.entry(neighbour).and_modify(|n| *n -= 1);
                sides.entry(*cube).and_modify(|n| *n -= 1);
            }
        }
    }
    sides
}

struct Day2022_18 {
    cubes: Vec<Pos>,
}

impl Solution<usize> for Day2022_18 {
    fn new() -> Day2022_18 {
        Day2022_18 { cubes: Vec::new() }
    }

    fn init(&mut self, input: &str) {
        for line in input.lines() {
            let mut it = line.trim().split(",").map(|s| s.parse::<i32>().unwrap());
            self.cubes
                .push((it.next().unwrap(), it.next().unwrap(), it.next().unwrap()));
        }
    }

    fn part_one(&mut self) -> usize {
        calc_sides(&self.cubes).values().sum()
    }

    fn part_two(&mut self) -> usize {
        let min = (
            self.cubes.iter().map(|(x, _, _)| *x).min().unwrap() - 1,
            self.cubes.iter().map(|(_, y, _)| *y).min().unwrap() - 1,
            self.cubes.iter().map(|(_, _, z)| *z).min().unwrap() - 1,
        );
        let max = (
            self.cubes.iter().map(|(x, _, _)| *x).max().unwrap() + 1,
            self.cubes.iter().map(|(_, y, _)| *y).max().unwrap() + 1,
            self.cubes.iter().map(|(_, _, z)| *z).max().unwrap() + 1,
        );
        let directions = get_directions();
        let mut outer_space: HashSet<Pos> = HashSet::new();
        let mut stack: Vec<Pos> = vec![min];
        while let Some(pos) = stack.pop() {
            if outer_space.contains(&pos) {
                continue;
            }
            outer_space.insert(pos);
            for dir in directions {
                let neighbour = (pos.0 + dir.0, pos.1 + dir.1, pos.2 + dir.2);
                if self.cubes.contains(&neighbour)
                    || neighbour.0 < min.0
                    || neighbour.1 < min.1
                    || neighbour.2 < min.2
                    || neighbour.0 > max.0
                    || neighbour.1 > max.1
                    || neighbour.2 > max.2
                {
                    continue;
                }
                stack.push(neighbour);
            }
        }
        let mut sides = calc_sides(&self.cubes);
        for pos in self.cubes.iter() {
            for dir in directions {
                let neighbour = (pos.0 + dir.0, pos.1 + dir.1, pos.2 + dir.2);
                if !self.cubes.contains(&neighbour) && !outer_space.contains(&neighbour) {
                    sides.entry(*pos).and_modify(|n| *n -= 1);
                }
            }
        }
        sides.values().sum()
    }
}

fn main() {
    let mut sol = Day2022_18::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_18;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_18.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_18::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 64);
        assert_eq!(sol.part_two(), 58);
    }
}
