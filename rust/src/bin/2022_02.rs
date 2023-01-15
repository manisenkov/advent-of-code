use aoc::solution::Solution;

#[derive(Clone, Copy)]
enum RPS {
    Rock,
    Paper,
    Scissors,
}

impl RPS {
    fn parse(c: char) -> RPS {
        match c {
            'A' | 'X' => RPS::Rock,
            'B' | 'Y' => RPS::Paper,
            'C' | 'Z' => RPS::Scissors,
            _ => panic!("invalid character"),
        }
    }

    fn to_match(self) -> MatchResult {
        match self {
            RPS::Rock => MatchResult::Lose,
            RPS::Paper => MatchResult::Draw,
            RPS::Scissors => MatchResult::Win,
        }
    }

    fn play(self, opponent: RPS) -> MatchResult {
        match (self, opponent) {
            (RPS::Rock, RPS::Scissors) => MatchResult::Win,
            (RPS::Scissors, RPS::Paper) => MatchResult::Win,
            (RPS::Paper, RPS::Rock) => MatchResult::Win,
            (RPS::Scissors, RPS::Rock) => MatchResult::Lose,
            (RPS::Paper, RPS::Scissors) => MatchResult::Lose,
            (RPS::Rock, RPS::Paper) => MatchResult::Lose,
            _ => MatchResult::Draw,
        }
    }

    fn reverse_play(self, result: MatchResult) -> RPS {
        match (self, result) {
            (RPS::Rock, MatchResult::Win) => RPS::Paper,
            (RPS::Paper, MatchResult::Win) => RPS::Scissors,
            (RPS::Scissors, MatchResult::Win) => RPS::Rock,
            (RPS::Rock, MatchResult::Lose) => RPS::Scissors,
            (RPS::Paper, MatchResult::Lose) => RPS::Rock,
            (RPS::Scissors, MatchResult::Lose) => RPS::Paper,
            (_, MatchResult::Draw) => self,
        }
    }

    fn score(self) -> u32 {
        match self {
            RPS::Rock => 1,
            RPS::Paper => 2,
            RPS::Scissors => 3,
        }
    }
}

#[derive(Clone, Copy)]
enum MatchResult {
    Win,
    Lose,
    Draw,
}

impl MatchResult {
    fn score(self) -> u32 {
        match self {
            MatchResult::Win => 6,
            MatchResult::Lose => 0,
            MatchResult::Draw => 3,
        }
    }
}

#[derive(Clone, Copy)]
struct Round {
    me: RPS,
    opponent: RPS,
}

impl Round {
    fn new(me: RPS, opponent: RPS) -> Self {
        Round { me, opponent }
    }

    fn score(self) -> u32 {
        self.me.play(self.opponent).score() + self.me.score()
    }
}

struct Day2022_02 {
    strategy: Vec<Round>,
}

impl Solution<u32> for Day2022_02 {
    fn new(input: &str) -> Day2022_02 {
        Day2022_02 {
            strategy: input
                .lines()
                .map(|line| {
                    let mut ch = line.chars();
                    let opponent = ch.nth(0).unwrap();
                    let me = ch.nth(1).unwrap();
                    Round::new(RPS::parse(me), RPS::parse(opponent))
                })
                .collect(),
        }
    }

    fn part_one(&mut self) -> u32 {
        self.strategy.iter().map(|r| r.score()).sum()
    }

    fn part_two(&mut self) -> u32 {
        self.strategy
            .iter()
            .map(|r| Round::new(r.opponent.reverse_play(r.me.to_match()), r.opponent).score())
            .sum()
    }
}

fn main() {
    Day2022_02::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_02;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_02.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_02::new(TEST_INPUT);
        assert_eq!(sol.part_one(), 15);
        assert_eq!(sol.part_two(), 12);
    }
}
