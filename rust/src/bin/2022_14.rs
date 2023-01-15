use std::collections::HashSet;

use aoc::solution::Solution;

fn is_occupied(rocks: &HashSet<(i32, i32)>, sand: &HashSet<(i32, i32)>, pos: (i32, i32)) -> bool {
    rocks.contains(&pos) || sand.contains(&pos)
}

fn drop_sand(
    rocks: &HashSet<(i32, i32)>,
    sand: &HashSet<(i32, i32)>,
    max_y: i32,
    start_pos: (i32, i32),
) -> (i32, i32) {
    let mut cur = start_pos;
    while cur.1 < max_y + 2 {
        if !is_occupied(rocks, sand, (cur.0, cur.1 + 1)) {
            cur = (cur.0, cur.1 + 1);
        } else if !is_occupied(rocks, sand, (cur.0 - 1, cur.1 + 1)) {
            cur = (cur.0 - 1, cur.1 + 1);
        } else if !is_occupied(rocks, sand, (cur.0 + 1, cur.1 + 1)) {
            cur = (cur.0 + 1, cur.1 + 1);
        } else {
            return cur;
        }
    }
    (cur.0, max_y + 1)
}

struct Day2022_14 {
    rocks: HashSet<(i32, i32)>,
    max_y: i32,
}

impl Solution<usize> for Day2022_14 {
    fn new(input: &str) -> Day2022_14 {
        let mut rocks = HashSet::new();
        let mut max_y: i32 = 0;
        for line in input.lines() {
            let path: Vec<_> = line
                .split("->")
                .map(|s| {
                    let p: Vec<_> = s.trim().split(",").map(|t| t.parse().unwrap()).collect();
                    (p[0], p[1])
                })
                .collect();
            for i in 0..path.len() - 1 {
                let d_abs = (path[i + 1].0 - path[i].0, path[i + 1].1 - path[i].1);
                let d_tot = ((d_abs.0 + d_abs.1) as i32).abs();
                let d = (d_abs.0 / d_tot, d_abs.1 / d_tot);
                let mut cur = path[i];
                while cur != path[i + 1] {
                    rocks.insert(cur);
                    cur = (cur.0 + d.0, cur.1 + d.1);
                }
                rocks.insert(path[i + 1]);
            }
            max_y = *path
                .iter()
                .map(|(_, y)| y)
                .chain(vec![max_y].iter())
                .max()
                .unwrap();
        }
        Day2022_14 { rocks, max_y }
    }

    fn part_one(&mut self) -> usize {
        let mut sand = HashSet::<(i32, i32)>::new();
        for i in 0.. {
            let grain = drop_sand(&self.rocks, &sand, self.max_y, (500, 0));
            if grain.1 > self.max_y {
                return i;
            }
            sand.insert(grain);
        }
        panic!("you shouldn't be here")
    }

    fn part_two(&mut self) -> usize {
        let mut sand = HashSet::<(i32, i32)>::new();
        for i in 0.. {
            let grain = drop_sand(&self.rocks, &sand, self.max_y, (500, 0));
            if grain == (500, 0) {
                return i + 1;
            }
            sand.insert(grain);
        }
        panic!("you shouldn't be here")
    }
}

fn main() {
    Day2022_14::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_14;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_14.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_14::new(TEST_INPUT);
        assert_eq!(sol.part_one(), 24);
        assert_eq!(sol.part_two(), 93);
    }
}
