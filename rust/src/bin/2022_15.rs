use std::collections::HashSet;

use aoc::solution::Solution;

fn manhattan(p1: &(i64, i64), p2: &(i64, i64)) -> i64 {
    (p1.0 - p2.0).abs() + (p1.1 - p2.1).abs()
}

fn merge_range(r1: (i64, i64), r2: (i64, i64)) -> Option<(i64, i64)> {
    if r1.1 + 1 == r2.0 {
        Some((r1.0, r2.1))
    } else if r1.0 <= r2.1 && r1.1 >= r2.0 {
        Some((r1.0.min(r2.0), r1.1.max(r2.1)))
    } else {
        None
    }
}

fn get_coverage(reports: &Vec<((i64, i64), (i64, i64))>, y_to_scan: i64) -> Vec<(i64, i64)> {
    let mut ranges: Vec<(i64, i64)> = vec![];
    for (sensor, beacon) in reports.iter() {
        let dist = manhattan(sensor, beacon);
        if dist < (y_to_scan - sensor.1).abs() {
            continue;
        }
        let dx = dist - (y_to_scan - sensor.1).abs();
        ranges.push((sensor.0 - dx, sensor.0 + dx));
    }

    // Merge ranges
    loop {
        ranges.sort_by(|r1, r2| r1.0.cmp(&r2.0));
        let next: Vec<(i64, i64)> = ranges.iter().fold(vec![], |mut acc, range| {
            match acc.pop() {
                Some(last) => match merge_range(last, *range) {
                    Some(new_range) => {
                        acc.push(new_range);
                    }
                    None => {
                        acc.push(last);
                        acc.push(*range);
                    }
                },
                None => {
                    acc.push(*range);
                }
            };
            acc
        });
        if next.len() == ranges.len() {
            break next;
        }
        ranges = next;
    }
}

struct Day2022_15 {
    y_to_scan: i64,
    reports: Vec<((i64, i64), (i64, i64))>,
}

impl Solution<usize, i64> for Day2022_15 {
    fn new() -> Day2022_15 {
        Day2022_15 {
            y_to_scan: 2_000_000,
            reports: vec![],
        }
    }

    fn init(&mut self, input: &str) {
        for line in input.lines() {
            let mut parts = line.trim().split(": ");
            let sensor_part = parts.next().unwrap();
            let sensor_x: i64 = sensor_part
                [sensor_part.find("x=").unwrap() + 2..sensor_part.find(",").unwrap()]
                .parse()
                .unwrap();
            let sensor_y: i64 = sensor_part[sensor_part.find("y=").unwrap() + 2..]
                .parse()
                .unwrap();
            let beacon_part = parts.next().unwrap();
            let beacon_x: i64 = beacon_part
                [beacon_part.find("x=").unwrap() + 2..beacon_part.find(",").unwrap()]
                .parse()
                .unwrap();
            let beacon_y: i64 = beacon_part[beacon_part.find("y=").unwrap() + 2..]
                .parse()
                .unwrap();
            self.reports
                .push(((sensor_x, sensor_y), (beacon_x, beacon_y)));
        }
    }

    fn part_one(&mut self) -> usize {
        let ranges = get_coverage(&self.reports, self.y_to_scan);
        let to_rm = HashSet::<(i64, i64)>::from_iter(
            self.reports
                .iter()
                .flat_map(|(sensor, beacon)| vec![sensor, beacon])
                .map(|(x, y)| (*x, *y))
                .filter(|(_, y)| *y == self.y_to_scan),
        )
        .len();
        (ranges[0].1 - ranges[0].0 + 1) as usize - to_rm
    }

    fn part_two(&mut self) -> i64 {
        // Find min Y and max Y
        let (min_y, max_y) =
            self.reports
                .iter()
                .fold((0x7fffffff, -0x7fffffff), |acc, (sensor, beacon)| {
                    let dist = manhattan(sensor, beacon);
                    (acc.0.min(sensor.1 - dist), acc.1.max(sensor.1 + dist))
                });

        for y_to_scan in min_y..=max_y {
            let ranges = get_coverage(&self.reports, y_to_scan);

            // Completely random observation 🤷
            if ranges.len() == 2 && ranges[0].1 + 2 == ranges[1].0 {
                return (ranges[0].1 + 1) * 4_000_000 + y_to_scan;
            }
        }
        0
    }
}

fn main() {
    let mut sol = Day2022_15::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_15;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_15.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_15::new();
        sol.init(TEST_INPUT);
        sol.y_to_scan = 10;
        assert_eq!(sol.part_one(), 26);
        assert_eq!(sol.part_two(), 56000011);
    }
}
