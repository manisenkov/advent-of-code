use std::collections::HashSet;

use aoc::solution::Solution;

struct DaySolution {
    signal: Vec<char>,
}

impl Solution<usize> for DaySolution {
    fn new() -> DaySolution {
        DaySolution { signal: Vec::new() }
    }

    fn init(&mut self, input: &str) {
        self.signal.extend(input.chars());
    }

    fn part_one(&mut self) -> usize {
        self.signal
            .windows(4)
            .enumerate()
            .find(|(_, w)| HashSet::<&char>::from_iter(w.iter()).len() == 4)
            .unwrap()
            .0
            + 4
    }

    fn part_two(&mut self) -> usize {
        self.signal
            .windows(14)
            .enumerate()
            .find(|(_, w)| HashSet::<&char>::from_iter(w.iter()).len() == 14)
            .unwrap()
            .0
            + 14
    }
}

fn main() {
    let mut sol = DaySolution::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::DaySolution;
    use aoc::solution::Solution;

    #[test]
    fn test_1() {
        let mut sol = DaySolution::new();
        sol.init("mjqjpqmgbljsphdztnvjfqwrcgsmlb");
        assert_eq!(sol.part_one(), 7);
        assert_eq!(sol.part_two(), 19);
    }

    #[test]
    fn test_2() {
        let mut sol = DaySolution::new();
        sol.init("bvwbjplbgvbhsrlpgdmjqwftvncz");
        assert_eq!(sol.part_one(), 5);
        assert_eq!(sol.part_two(), 23);
    }

    #[test]
    fn test_3() {
        let mut sol = DaySolution::new();
        sol.init("nppdvjthqldpwncqszvftbrmjlhg");
        assert_eq!(sol.part_one(), 6);
        assert_eq!(sol.part_two(), 23);
    }

    #[test]
    fn test_4() {
        let mut sol = DaySolution::new();
        sol.init("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg");
        assert_eq!(sol.part_one(), 10);
        assert_eq!(sol.part_two(), 29);
    }

    #[test]
    fn test_5() {
        let mut sol = DaySolution::new();
        sol.init("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw");
        assert_eq!(sol.part_one(), 11);
        assert_eq!(sol.part_two(), 26);
    }
}
