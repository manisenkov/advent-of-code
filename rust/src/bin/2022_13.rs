use std::cmp::Ordering;

use serde_json::{json, Value};

use aoc::solution::Solution;

struct Day2022_13 {
    pairs: Vec<(Value, Value)>,
}

fn compare(left: &Value, right: &Value) -> Ordering {
    match (left, right) {
        (Value::Array(arr_left), Value::Array(arr_right)) => {
            for i in 0.. {
                if i < arr_left.len() && i < arr_right.len() {
                    let res = compare(&arr_left[i], &arr_right[i]);
                    if res != Ordering::Equal {
                        return res;
                    }
                } else if i == arr_left.len() && i < arr_right.len() {
                    return Ordering::Less;
                } else if i < arr_left.len() && i == arr_right.len() {
                    return Ordering::Greater;
                } else {
                    break;
                }
            }
            return Ordering::Equal;
        }
        (Value::Number(n), Value::Array(_)) => {
            return compare(&Value::Array(vec![Value::Number(n.clone())]), right);
        }
        (Value::Array(_), Value::Number(n)) => {
            return compare(left, &Value::Array(vec![Value::Number(n.clone())]));
        }
        (Value::Number(n_left), Value::Number(n_right)) => {
            return n_left.as_u64().unwrap().cmp(&n_right.as_u64().unwrap());
        }
        _ => panic!("wrong values"),
    }
}

impl Solution<usize> for Day2022_13 {
    fn new() -> Day2022_13 {
        Day2022_13 { pairs: vec![] }
    }

    fn init(&mut self, input: &str) {
        let lines: Vec<&str> = input.lines().collect();
        for i in (0..lines.len()).step_by(3) {
            self.pairs.push((
                serde_json::from_str(lines[i]).unwrap(),
                serde_json::from_str(lines[i + 1]).unwrap(),
            ));
        }
    }

    fn part_one(&mut self) -> usize {
        self.pairs
            .iter()
            .enumerate()
            .filter(|(_, pair)| compare(&pair.0, &pair.1).is_le())
            .map(|(i, _)| i + 1)
            .sum()
    }

    fn part_two(&mut self) -> usize {
        let mut values: Vec<Value> = self
            .pairs
            .iter()
            .flat_map(|pair| vec![pair.0.clone(), pair.1.clone()])
            .collect();
        values.push(json!([2]));
        values.push(json!([6]));
        values.sort_by(compare);
        let (index_1, _) = values
            .iter()
            .enumerate()
            .find(|(_, v)| (**v).eq(&json!([2])))
            .unwrap();
        let (index_2, _) = values
            .iter()
            .enumerate()
            .find(|(_, v)| (**v).eq(&json!([6])))
            .unwrap();
        (index_1 + 1) * (index_2 + 1)
    }
}

fn main() {
    let mut sol = Day2022_13::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_13;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_13.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_13::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 13);
        assert_eq!(sol.part_two(), 140);
    }
}
