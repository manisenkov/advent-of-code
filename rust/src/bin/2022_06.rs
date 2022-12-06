use aoc::solution::Solution;

struct DaySolution {
    signal: Vec<char>,
}

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
        .enumerate()
        .find(|(_, s)| is_all_different(s))
        .unwrap()
        .0
        + size
}

impl Solution<usize> for DaySolution {
    fn new() -> DaySolution {
        DaySolution { signal: Vec::new() }
    }

    fn init(&mut self, input: &str) {
        self.signal.extend(input.chars());
    }

    fn part_one(&mut self) -> usize {
        find_packet_index(&self.signal, 4)
    }

    fn part_two(&mut self) -> usize {
        find_packet_index(&self.signal, 14)
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
