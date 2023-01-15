use std::collections::HashMap;

use aoc::solution::Solution;

fn from_snafu(src: &String) -> i64 {
    let mut res = 0;
    let mut factor = 1;
    for c in src.chars().rev() {
        match c {
            '=' => res += factor * -2,
            '-' => res += factor * -1,
            '0' => res += factor * 0,
            '1' => res += factor * 1,
            '2' => res += factor * 2,
            _ => panic!("wrong char"),
        }
        factor *= 5;
    }
    res
}

fn to_snafu(src: i64) -> String {
    let char_map = HashMap::from([(0, '0'), (1, '1'), (2, '2'), (3, '='), (4, '-')]);
    let mut left = src;
    let mut res = String::from("");
    while left > 0 {
        let n = left % 5;
        res.push(char_map[&n]);
        left = (if n >= 3 { left + 5 } else { left }) / 5;
    }
    res.chars().rev().collect()
}

struct Day2022_25 {
    nums: Vec<String>,
}

impl Solution<String> for Day2022_25 {
    fn new(input: &str) -> Day2022_25 {
        Day2022_25 {
            nums: input.lines().map(|line| line.trim().to_owned()).collect(),
        }
    }

    fn part_one(&mut self) -> String {
        let res = self.nums.iter().map(from_snafu).sum::<i64>();
        to_snafu(res)
    }

    fn part_two(&mut self) -> String {
        "".to_owned()
    }
}

fn main() {
    Day2022_25::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_25;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_25.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_25::new(TEST_INPUT);
        assert_eq!(sol.part_one(), "2=-1=0");
        // assert_eq!(sol.part_two(), "");
    }
}
