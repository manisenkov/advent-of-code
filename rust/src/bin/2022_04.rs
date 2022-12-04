use aoc::solution::Solution;

#[derive(Clone, Copy)]
struct IDRange(u32, u32);

impl IDRange {
    fn parse(input: &str) -> IDRange {
        let p: Vec<&str> = input.split("-").collect();
        IDRange(p[0].parse().unwrap(), p[1].parse().unwrap())
    }

    fn include(&self, other: &IDRange) -> bool {
        self.0 <= other.0 && self.1 >= other.1
    }

    fn overlap(&self, other: &IDRange) -> bool {
        self.0 <= other.1 && self.1 >= other.0
    }
}

#[derive(Clone, Copy)]
struct IDRangePair(IDRange, IDRange);

impl IDRangePair {
    fn parse(input: &str) -> IDRangePair {
        let p: Vec<&str> = input.split(",").collect();
        IDRangePair(IDRange::parse(p[0]), IDRange::parse(p[1]))
    }
}

struct Day2022_04 {
    range_pairs: Vec<IDRangePair>,
}

impl Solution<u32> for Day2022_04 {
    fn new() -> Day2022_04 {
        Day2022_04 {
            range_pairs: Vec::new(),
        }
    }

    fn init(&mut self, input: &str) {
        self.range_pairs
            .extend(input.lines().map(|s| IDRangePair::parse(s.trim())));
    }

    fn part_one(&mut self) -> u32 {
        self.range_pairs
            .iter()
            .filter(|range_pair| {
                range_pair.0.include(&range_pair.1) || range_pair.1.include(&range_pair.0)
            })
            .count() as u32
    }

    fn part_two(&mut self) -> u32 {
        self.range_pairs
            .iter()
            .filter(|range_pair| range_pair.0.overlap(&range_pair.1))
            .count() as u32
    }
}

fn main() {
    let mut sol = Day2022_04::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_04;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_04.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_04::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 2);
        assert_eq!(sol.part_two(), 4);
    }
}
