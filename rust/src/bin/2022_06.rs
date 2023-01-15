use aoc::solution::Solution;

fn is_all_different(s: &[char]) -> bool {
    for i in 0..s.len() - 1 {
        for j in (i + 1)..s.len() {
            if s[i] == s[j] {
                return false;
            }
        }
    }
    true
}

fn find_packet_index(signal: &[char], size: usize) -> usize {
    signal
        .windows(size)
        .position(|s| is_all_different(s))
        .unwrap()
        + size
}

struct Day2022_06 {
    signal: Vec<char>,
}

impl Solution<usize> for Day2022_06 {
    fn new(input: &str) -> Day2022_06 {
        Day2022_06 {
            signal: input.chars().collect(),
        }
    }

    fn part_one(&mut self) -> usize {
        find_packet_index(&self.signal, 4)
    }

    fn part_two(&mut self) -> usize {
        find_packet_index(&self.signal, 14)
    }
}

fn main() {
    Day2022_06::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_06;
    use aoc::solution::Solution;

    #[test]
    fn test_1() {
        let mut sol = Day2022_06::new("mjqjpqmgbljsphdztnvjfqwrcgsmlb");
        assert_eq!(sol.part_one(), 7);
        assert_eq!(sol.part_two(), 19);
    }

    #[test]
    fn test_2() {
        let mut sol = Day2022_06::new("bvwbjplbgvbhsrlpgdmjqwftvncz");
        assert_eq!(sol.part_one(), 5);
        assert_eq!(sol.part_two(), 23);
    }

    #[test]
    fn test_3() {
        let mut sol = Day2022_06::new("nppdvjthqldpwncqszvftbrmjlhg");
        assert_eq!(sol.part_one(), 6);
        assert_eq!(sol.part_two(), 23);
    }

    #[test]
    fn test_4() {
        let mut sol = Day2022_06::new("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg");
        assert_eq!(sol.part_one(), 10);
        assert_eq!(sol.part_two(), 29);
    }

    #[test]
    fn test_5() {
        let mut sol = Day2022_06::new("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw");
        assert_eq!(sol.part_one(), 11);
        assert_eq!(sol.part_two(), 26);
    }
}
