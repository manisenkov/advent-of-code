use std::collections::HashSet;

use aoc::solution::Solution;

fn get_prio(c: char) -> u32 {
    match c {
        'a'..='z' => (c as u32) - ('a' as u32) + 1,
        'A'..='Z' => (c as u32) - ('A' as u32) + 27,
        _ => panic!("wrong character"),
    }
}

struct Day2022_03 {
    rucksack_content: Vec<Vec<char>>,
}

impl Solution<u32> for Day2022_03 {
    fn new() -> Day2022_03 {
        Day2022_03 {
            rucksack_content: Vec::new(),
        }
    }

    fn init(&mut self, input: &str) {
        for line in input.lines() {
            self.rucksack_content.push(line.trim().chars().collect());
        }
    }

    fn part_one(&mut self) -> u32 {
        self.rucksack_content
            .iter()
            .map(|content| {
                let part1: HashSet<char> =
                    HashSet::from_iter(content.iter().take(content.len() / 2).map(|c| *c));
                let part2: HashSet<char> =
                    HashSet::from_iter(content.iter().skip(content.len() / 2).map(|c| *c));
                part1
                    .intersection(&part2)
                    .map(|c| get_prio(*c))
                    .sum::<u32>()
            })
            .sum()
    }

    fn part_two(&mut self) -> u32 {
        self.rucksack_content
            .chunks(3)
            .map(|chunk| {
                chunk
                    .iter()
                    .map(|content| HashSet::from_iter(content.iter().map(|c| *c)))
                    .reduce(|acc: HashSet<char>, s| {
                        HashSet::from_iter(acc.intersection(&s).map(|c| *c))
                    })
                    .unwrap()
                    .iter()
                    .map(|c| get_prio(*c))
                    .sum::<u32>()
            })
            .sum()
    }
}

fn main() {
    let mut sol = Day2022_03::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_03;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_03.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_03::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 157);
        assert_eq!(sol.part_two(), 70);
    }
}
