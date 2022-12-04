use aoc::solution::Solution;

struct DaySolution {
    count: u32,
}

impl Solution<u32> for DaySolution {
    fn new() -> DaySolution {
        DaySolution { count: 0 }
    }

    fn init(&mut self, input: &str) {
        for line in input.lines() {
            println!("{}", line);
            self.count += line.parse::<u32>().unwrap()
        }
    }

    fn part_one(&mut self) -> u32 {
        self.count += 1;
        self.count
    }

    fn part_two(&mut self) -> u32 {
        self.count += 1;
        self.count
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
        sol.init("0");
        assert_eq!(sol.part_one(), 1);
        assert_eq!(sol.part_two(), 2);
    }

    #[test]
    fn test_2() {
        let mut sol = DaySolution::new();
        sol.init("42");
        assert_eq!(sol.part_one(), 43);
        assert_eq!(sol.part_two(), 44);
    }
}
