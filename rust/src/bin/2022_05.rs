use aoc::solution::Solution;

struct Move {
    count: usize,
    from: usize,
    to: usize,
}

impl Move {
    fn parse(input: &str) -> Move {
        let v: Vec<&str> = input.split(" ").collect();
        Move {
            count: v[1].parse().unwrap(),
            from: v[3].parse::<usize>().unwrap() - 1,
            to: v[5].parse::<usize>().unwrap() - 1,
        }
    }
}

struct Day2022_05 {
    plan: Vec<Vec<char>>,
    moves: Vec<Move>,
}

impl Solution<String> for Day2022_05 {
    fn new() -> Day2022_05 {
        Day2022_05 {
            plan: Vec::new(),
            moves: Vec::new(),
        }
    }

    fn init(&mut self, input: &str) {
        let (plan, moves) = input.split_once("\n\n").unwrap();

        let plan_lines: Vec<&str> = plan.lines().collect();
        let num_col = (plan_lines.last().unwrap().trim().len() + 3) / 4;
        self.plan.extend((0..num_col).map(|_| vec![]));
        for line in plan_lines.iter().take(plan_lines.len() - 1).rev() {
            for (i, c) in line.chars().skip(1).step_by(4).enumerate() {
                if !c.is_whitespace() {
                    self.plan[i].push(c)
                }
            }
        }

        self.moves = moves.lines().map(|s| Move::parse(s.trim())).collect();
    }

    fn part_one(&mut self) -> String {
        let mut state = self.plan.to_vec();
        for mov in &self.moves {
            for _ in 0..mov.count {
                let c = state[mov.from].pop().unwrap();
                state[mov.to].push(c);
            }
        }
        state.iter().map(|v| v.last().unwrap()).collect::<String>()
    }

    fn part_two(&mut self) -> String {
        let mut state = self.plan.to_vec();
        for mov in &self.moves {
            let skip_count = state[mov.from].len() - mov.count;
            let mut target = vec![];
            std::mem::swap(&mut target, &mut state[mov.to]);
            target.extend(
                state[mov.from]
                    .iter()
                    .skip(skip_count)
                    .take(mov.count)
                    .map(|c| *c),
            );
            std::mem::swap(&mut target, &mut state[mov.to]);
            state[mov.from].truncate(skip_count)
        }
        state.iter().map(|v| v.last().unwrap()).collect::<String>()
    }
}

fn main() {
    let mut sol = Day2022_05::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_05;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_05.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_05::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), "CMZ");
        assert_eq!(sol.part_two(), "MCD");
    }
}
