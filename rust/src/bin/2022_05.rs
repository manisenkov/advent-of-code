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
    init_stacks: Vec<Vec<char>>,
    moves: Vec<Move>,
}

impl Solution<String> for Day2022_05 {
    fn new() -> Day2022_05 {
        Day2022_05 {
            init_stacks: Vec::new(),
            moves: Vec::new(),
        }
    }

    fn init(&mut self, input: &str) {
        let mut box_plan = Vec::<String>::new();
        let mut box_part_done = false;
        for line in input.lines() {
            if line.trim() == "" {
                box_part_done = true;
                continue;
            }
            if box_part_done {
                // Parse moves
                self.moves.push(Move::parse(line.trim()))
            } else {
                // Add new line to map
                box_plan.push(String::from(line))
            }
        }

        // Parse initial box map
        let num_col = (box_plan.last().unwrap().len() + 2) / 4;
        self.init_stacks.extend((0..num_col).map(|_| Vec::new()));
        for plan_line in box_plan.iter().take(box_plan.len() - 1).rev() {
            let chars: Vec<char> = plan_line.chars().collect();
            for i in 0..num_col {
                match chars.get((i * 4) + 1) {
                    Some(c) => {
                        if !c.is_whitespace() {
                            self.init_stacks[i].push(*c);
                        }
                    }
                    None => {
                        break;
                    }
                }
            }
        }
    }

    fn part_one(&mut self) -> String {
        let mut state = self.init_stacks.to_vec();
        for mov in &self.moves {
            for _ in 0..mov.count {
                let c = state[mov.from].pop().unwrap();
                state[mov.to].push(c);
            }
        }
        state.iter().map(|v| v.last().unwrap()).collect::<String>()
    }

    fn part_two(&mut self) -> String {
        let mut state = self.init_stacks.to_vec();
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
