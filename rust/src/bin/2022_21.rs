use std::collections::HashMap;

use aoc::solution::Solution;

#[derive(Clone)]
enum Op {
    Number(i64),
    Add(String, String),
    Sub(String, String),
    Mul(String, String),
    Div(String, String),
}

impl Op {
    fn parse(input: &str) -> Self {
        if let Ok(num) = input.parse::<i64>() {
            Op::Number(num)
        } else {
            let mut parts = input.split(" ");
            let monkey_l = parts.next().unwrap().to_owned();
            let op = parts.next().unwrap();
            let monkey_r = parts.next().unwrap().to_owned();
            match op {
                "+" => Op::Add(monkey_l, monkey_r),
                "-" => Op::Sub(monkey_l, monkey_r),
                "*" => Op::Mul(monkey_l, monkey_r),
                "/" => Op::Div(monkey_l, monkey_r),
                _ => panic!("invalid operation"),
            }
        }
    }
}

#[derive(Clone)]
struct Monkey {
    name: String,
    op: Op,
}

impl Monkey {
    fn yell(&self, others: &HashMap<String, Monkey>, mem: &mut HashMap<String, i64>) -> i64 {
        if let Some(num) = mem.get(&self.name) {
            *num
        } else {
            let res = match &self.op {
                Op::Number(num) => *num,
                Op::Add(monkey_l, monkey_r) => {
                    others[monkey_l].yell(others, mem) + others[monkey_r].yell(others, mem)
                }
                Op::Sub(monkey_l, monkey_r) => {
                    others[monkey_l].yell(others, mem) - others[monkey_r].yell(others, mem)
                }
                Op::Mul(monkey_l, monkey_r) => {
                    others[monkey_l].yell(others, mem) * others[monkey_r].yell(others, mem)
                }
                Op::Div(monkey_l, monkey_r) => {
                    others[monkey_l].yell(others, mem) / others[monkey_r].yell(others, mem)
                }
            };
            mem.insert(self.name.clone(), res);
            res
        }
    }

    fn deps(&self) -> Option<(String, String)> {
        match &self.op {
            Op::Number(_) => None,
            Op::Add(monkey_l, monkey_r) => Some((monkey_l.clone(), monkey_r.clone())),
            Op::Sub(monkey_l, monkey_r) => Some((monkey_l.clone(), monkey_r.clone())),
            Op::Mul(monkey_l, monkey_r) => Some((monkey_l.clone(), monkey_r.clone())),
            Op::Div(monkey_l, monkey_r) => Some((monkey_l.clone(), monkey_r.clone())),
        }
    }
}

struct Day2022_21 {
    monkeys: HashMap<String, Monkey>,
}

impl Solution<i64> for Day2022_21 {
    fn new() -> Day2022_21 {
        Day2022_21 {
            monkeys: HashMap::new(),
        }
    }

    fn init(&mut self, input: &str) {
        for line in input.lines() {
            let mut parts = line.trim().split(": ");
            let name = parts.next().unwrap().to_owned();
            let op = Op::parse(parts.next().unwrap());
            self.monkeys.insert(
                name.clone(),
                Monkey {
                    name: name.clone(),
                    op,
                },
            );
        }
    }

    fn part_one(&mut self) -> i64 {
        let mut mem = HashMap::<String, i64>::new();
        self.monkeys[&"root".to_owned()].yell(&self.monkeys, &mut mem)
    }

    fn part_two(&mut self) -> i64 {
        let mut new_rule_monkeys = HashMap::from_iter(
            self.monkeys
                .iter()
                .map(|(name, monkey)| (name.clone(), monkey.clone())),
        );
        let (root_left, root_right) = new_rule_monkeys["root"].deps().unwrap();

        let mut eval = |humn_yell: i64| {
            let mut mem = HashMap::<String, i64>::new();
            new_rule_monkeys.insert(
                "humn".to_owned(),
                Monkey {
                    name: "humn".to_owned(),
                    op: Op::Number(humn_yell),
                },
            );
            let left_yell = new_rule_monkeys[&root_left].yell(&new_rule_monkeys, &mut mem);
            let right_yell = new_rule_monkeys[&root_right].yell(&new_rule_monkeys, &mut mem);
            right_yell - left_yell
        };

        let mut left_point: i64 = -0x80000000000;
        let mut mid_point: i64 = 0;
        let mut right_point: i64 = 0x80000000000;
        loop {
            let left_res = eval(left_point);
            let mid_res = eval(mid_point);
            let right_res = eval(right_point);
            if mid_res != 0 {
                if (left_res > 0 && mid_res > 0 && right_res < 0)
                    || (left_res < 0 && mid_res < 0 && right_res > 0)
                {
                    left_point = mid_point;
                } else {
                    right_point = mid_point;
                }
                mid_point = left_point + (right_point - left_point) / 2;
            } else {
                break loop {
                    if eval(mid_point - 1) == 0 {
                        mid_point -= 1;
                    } else {
                        break mid_point;
                    }
                };
            }
        }
    }
}

fn main() {
    let mut sol = Day2022_21::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_21;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_21.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_21::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 152);
        assert_eq!(sol.part_two(), 301);
    }
}
