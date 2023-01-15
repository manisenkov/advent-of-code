use aoc::solution::Solution;

struct Day2022_01 {
    elf_loads: Vec<Vec<u32>>,
}

impl Solution<u32> for Day2022_01 {
    fn new(input: &str) -> Day2022_01 {
        let mut elf_loads = vec![];
        let mut cur_load = vec![];
        for line in input.lines() {
            if line.trim().is_empty() {
                elf_loads.push(cur_load);
                cur_load = Vec::new();
            } else {
                cur_load.push(line.trim().parse().unwrap());
            }
        }
        elf_loads.push(cur_load);
        Day2022_01 { elf_loads }
    }

    fn part_one(&mut self) -> u32 {
        self.elf_loads.iter().map(|v| v.iter().sum()).max().unwrap()
    }

    fn part_two(&mut self) -> u32 {
        let mut sums: Vec<_> = self.elf_loads.iter().map(|v| v.iter().sum()).collect();
        sums.sort();
        sums.iter().rev().take(3).sum()
    }
}

fn main() {
    Day2022_01::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_01;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_01.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_01::new(TEST_INPUT);
        assert_eq!(sol.part_one(), 24000);
        assert_eq!(sol.part_two(), 45000);
    }
}
