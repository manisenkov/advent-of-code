use aoc::solution::Solution;

#[derive(Clone)]
enum Operation {
    Plus(u64),
    Prod(u64),
    Square,
}

enum Mode {
    Simple,
    Extended(u64),
}

impl Mode {
    fn apply(&self, arg: u64) -> u64 {
        match self {
            Mode::Simple => arg / 3,
            Mode::Extended(n) => arg % n,
        }
    }
}

impl Operation {
    fn parse(input: &str) -> Operation {
        let parts: Vec<&str> = input.split(" ").collect();
        if parts[3] == "+" {
            Operation::Plus(parts[4].parse().unwrap())
        } else if parts[3] == "*" {
            if parts[4] == "old" {
                Operation::Square
            } else {
                Operation::Prod(parts[4].parse().unwrap())
            }
        } else {
            panic!("invalid operation")
        }
    }

    fn run(&self, old: u64) -> u64 {
        match self {
            Operation::Plus(n) => old + *n,
            Operation::Prod(n) => old * *n,
            Operation::Square => old * old,
        }
    }
}

#[derive(Clone)]
struct Monkey {
    items: Vec<u64>,
    operation: Operation,
    test_factor: u64,
    if_true: usize,
    if_false: usize,
}

impl<'a> Monkey {
    fn parse(input: &[&str]) -> Monkey {
        let items: Vec<u64>;
        let operation: Operation;
        let test_factor: u64;
        let if_true: usize;
        let if_false: usize;

        // Items
        {
            let mut it = input[1].trim().split(":");
            it.next();
            items = it
                .next()
                .unwrap()
                .trim()
                .split(",")
                .map(|s| s.trim().parse().unwrap())
                .collect();
        }

        // Operation
        {
            let mut it = input[2].trim().split(":");
            it.next();
            operation = Operation::parse(it.next().unwrap().trim());
        }

        // Test factor
        {
            let v: Vec<&str> = input[3].trim().split(" ").collect();
            test_factor = v[v.len() - 1].parse().unwrap();
        }

        // If true
        {
            let v: Vec<&str> = input[4].trim().split(" ").collect();
            if_true = v[v.len() - 1].parse().unwrap();
        }

        // If false
        {
            let v: Vec<&str> = input[5].trim().split(" ").collect();
            if_false = v[v.len() - 1].parse().unwrap();
        }

        Monkey {
            items,
            operation,
            test_factor,
            if_true,
            if_false,
        }
    }
}

fn turn(index: usize, monkeys: &mut Vec<Monkey>, mode: Mode) -> usize {
    let count = monkeys[index].items.len();
    let to_inspect: Vec<u64> = monkeys[index].items.drain(..count).collect();
    for item in to_inspect {
        let worry_level = mode.apply(monkeys[index].operation.run(item));
        let test_factor = monkeys[index].test_factor;
        let to_index = if worry_level % test_factor == 0 {
            monkeys[index].if_true
        } else {
            monkeys[index].if_false
        };
        monkeys[to_index].items.push(worry_level);
    }
    count
}

struct Day2022_11 {
    monkeys: Vec<Monkey>,
}

impl Solution<usize> for Day2022_11 {
    fn new() -> Day2022_11 {
        Day2022_11 {
            monkeys: Vec::new(),
        }
    }

    fn init(&mut self, input: &str) {
        let s: Vec<&str> = input.lines().collect();
        for i in (0..s.len()).step_by(7) {
            self.monkeys.push(Monkey::parse(&s[i..]));
        }
    }

    fn part_one(&mut self) -> usize {
        let mut monkeys = self.monkeys.clone();
        let mut counters = vec![0; monkeys.len()];
        for _ in 0..20 {
            for i in 0..monkeys.len() {
                counters[i] += turn(i, &mut monkeys, Mode::Simple);
            }
        }
        counters.sort_unstable();
        counters.reverse();
        counters[0] * counters[1]
    }

    fn part_two(&mut self) -> usize {
        let mut monkeys = self.monkeys.clone();
        let mut counters = vec![0; monkeys.len()];
        let mul_factor = monkeys.iter().fold(1, |acc, m| acc * m.test_factor);
        for _ in 0..10000 {
            for i in 0..monkeys.len() {
                counters[i] += turn(i, &mut monkeys, Mode::Extended(mul_factor));
            }
        }
        counters.sort_unstable();
        counters.reverse();
        counters[0] * counters[1]
    }
}

fn main() {
    let mut sol = Day2022_11::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_11;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_11.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_11::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 10605);
        assert_eq!(sol.part_two(), 2713310158);
    }
}
