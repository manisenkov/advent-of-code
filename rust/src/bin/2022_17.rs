use std::{
    collections::{HashMap, HashSet},
    hash::Hash,
};

use aoc::solution::Solution;

struct Figure(Vec<(i64, i64)>);

impl Figure {
    fn new(cells: &[(i64, i64)]) -> Self {
        Figure(Vec::from(cells))
    }

    fn to_rock(&self, pos: &(i64, i64)) -> Rock {
        let Figure(cells) = self;
        Rock(cells.clone()).mov(pos)
    }
}

struct Rock(Vec<(i64, i64)>);

impl Rock {
    fn fit(&self, occupied: &HashSet<(i64, i64)>) -> bool {
        let Rock(cells) = self;
        cells.iter().all(|(row, col)| {
            *row >= 0 && *col >= 0 && *col < 7 && !occupied.contains(&(*row, *col))
        })
    }

    fn mov(&self, pos: &(i64, i64)) -> Rock {
        let (d_row, d_col) = pos;
        let Rock(cells) = self;
        Rock(Vec::from_iter(
            cells
                .iter()
                .map(|(row, col)| (*row + *d_row, *col + *d_col)),
        ))
    }
}

#[derive(PartialEq, Eq, Hash)]
struct Step {
    rock_index: usize,
    dir_index: usize,
    fingerprint: i64,
}

fn calc_fingerprint(occupied: &HashSet<(i64, i64)>, top: i64) -> i64 {
    let mut res: i64 = 0;
    let mut idx = 0;
    for row in (top - 8..top).rev() {
        for col in 0..7 {
            if occupied.contains(&(row, col)) {
                res += 1 << idx;
            }
            idx += 1;
        }
    }
    res
}

struct Day2022_17 {
    figures: Vec<Figure>,
    directions: Vec<char>,
}

impl Day2022_17 {
    fn cycle(&self) -> impl Iterator<Item = (Step, i64)> + '_ {
        let mut occupied: HashSet<(i64, i64)> = HashSet::new();
        let mut top = 0;
        let mut dir_index = 0;
        (0..).map(move |i| {
            let mut rock = self.figures[i % self.figures.len()].to_rock(&(top + 3, 2));
            loop {
                let dir = self.directions[dir_index];
                dir_index = (dir_index + 1) % self.directions.len();

                let mut next_rock = rock.mov(if dir == '<' { &(0, -1) } else { &(0, 1) });
                if next_rock.fit(&occupied) {
                    rock = next_rock;
                }

                next_rock = rock.mov(&(-1, 0));
                if next_rock.fit(&occupied) {
                    rock = next_rock;
                } else {
                    let Rock(cells) = rock;
                    occupied.extend(cells.iter());
                    top = top.max(cells.iter().map(|(row, _)| *row + 1).max().unwrap());
                    break;
                }
            }
            (
                Step {
                    rock_index: i % self.figures.len(),
                    dir_index: dir_index,
                    fingerprint: calc_fingerprint(&occupied, top.max(8)),
                },
                top,
            )
        })
    }
}

impl Solution<i64> for Day2022_17 {
    fn new() -> Day2022_17 {
        Day2022_17 {
            figures: Vec::new(),
            directions: Vec::new(),
        }
    }

    fn init(&mut self, input: &str) {
        self.directions.extend(input.trim().chars());
        self.figures
            .push(Figure::new(&[(0, 0), (0, 1), (0, 2), (0, 3)])); // -
        self.figures
            .push(Figure::new(&[(0, 1), (1, 0), (1, 1), (1, 2), (2, 1)])); // +
        self.figures
            .push(Figure::new(&[(0, 0), (0, 1), (0, 2), (1, 2), (2, 2)])); // J
        self.figures
            .push(Figure::new(&[(0, 0), (1, 0), (2, 0), (3, 0)])); // I
        self.figures
            .push(Figure::new(&[(0, 0), (0, 1), (1, 0), (1, 1)])); // o
    }

    fn part_one(&mut self) -> i64 {
        self.cycle().take(2022).last().unwrap().1
    }

    fn part_two(&mut self) -> i64 {
        let mut steps: HashMap<Step, usize> = HashMap::new();
        let mut tops: Vec<i64> = vec![];
        let mut repeated_at: usize = 0;
        for (i, (step, top)) in self.cycle().enumerate() {
            if steps.contains_key(&step) {
                repeated_at = steps[&step];
                break;
            }
            steps.insert(step, i);
            tops.push(top);
        }
        let cycle_size = steps.len() - repeated_at;
        let diff = *tops.last().unwrap() - tops[repeated_at - 1];
        let total = 1_000_000_000_000 - repeated_at;
        let rest = total % cycle_size;

        *tops.last().unwrap()
            + (total / cycle_size - 1) as i64 * diff
            + tops[repeated_at + rest - 1]
            - tops[repeated_at - 1]
    }
}

fn main() {
    let mut sol = Day2022_17::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_17;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_17.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_17::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 3068);
        assert_eq!(sol.part_two(), 1514285714288);
    }
}
