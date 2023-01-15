use std::collections::{HashMap, HashSet, VecDeque};

use aoc::solution::Solution;

fn calc_dists(start: &String, tunnels: &HashMap<String, Vec<String>>) -> HashMap<String, u64> {
    let mut queue: VecDeque<String> = VecDeque::from([start.clone()]);
    let mut dists: HashMap<String, u64> = HashMap::from([(start.clone(), 0)]);
    while let Some(cur_pos) = queue.pop_front() {
        let cur_dist = dists[&cur_pos];
        for next in tunnels[&cur_pos].iter() {
            let next_dist = *dists.get(next).unwrap_or(&0x7fffffff);
            if next_dist > cur_dist + 1 {
                dists.insert(next.clone(), cur_dist + 1);
                queue.push_back(next.clone());
            }
        }
    }
    dists
}

fn travel(
    valves: &HashMap<String, u64>,
    dists: &HashMap<(String, String), u64>,
    path: Vec<String>,
    time_left: u64,
) -> u64 {
    let visited: HashSet<String> = HashSet::from_iter(path.iter().map(|s| s.clone()));
    let cur_valve = path.last().unwrap().clone();
    valves
        .iter()
        .filter(|(valve, _)| {
            let dist = dists[&(cur_valve.clone(), (*valve).clone())];
            !visited.contains(*valve) && dist < time_left
        })
        .map(|(valve, flow)| {
            let dist = dists[&(cur_valve.clone(), valve.clone())];
            let next_path: Vec<_> = path
                .iter()
                .chain(vec![valve.clone()].iter())
                .map(|s| s.clone())
                .collect();
            let next_time_left = time_left - dist - 1;
            next_time_left * flow + travel(valves, dists, next_path, next_time_left)
        })
        .max()
        .unwrap_or(0)
}

struct Day2022_16 {
    valves: HashMap<String, u64>,
    dists: HashMap<(String, String), u64>,
}

impl Solution<u64> for Day2022_16 {
    fn new(input: &str) -> Day2022_16 {
        let mut tunnels: HashMap<String, Vec<String>> = HashMap::new();
        let mut all_valves: HashMap<String, u64> = HashMap::new();
        let mut valves = HashMap::new();
        let mut dists = HashMap::new();

        for line in input.lines() {
            let mut s = line.trim().split("; ");
            let valve_input = s.next().unwrap();
            let tunnel_input = s.next().unwrap();
            let valve_name = valve_input[6..8].to_owned();
            let flow_rate: u64 = valve_input[23..].parse().unwrap();
            let destinations: Vec<_> = tunnel_input[22..]
                .split(", ")
                .map(|s| s.trim().to_owned())
                .collect();
            all_valves.insert(valve_name.clone(), flow_rate);
            tunnels.insert(valve_name.clone(), destinations);
            if flow_rate > 0 {
                valves.insert(valve_name.clone(), flow_rate);
            }
        }
        for valve in all_valves.keys() {
            for (k, dist) in calc_dists(&valve, &tunnels).iter() {
                dists.insert((valve.clone(), k.clone()), *dist);
            }
        }
        Day2022_16 { valves, dists }
    }

    fn part_one(&mut self) -> u64 {
        travel(&self.valves, &self.dists, Vec::from(["AA".to_owned()]), 30)
    }

    fn part_two(&mut self) -> u64 {
        let valves_vec = Vec::from_iter(
            self.valves
                .iter()
                .map(|(valve, flow)| (valve.clone(), *flow)),
        );
        (0..1 << (valves_vec.len() - 1))
            .map(|i| {
                let valves_me: HashMap<String, u64> = HashMap::from_iter(
                    valves_vec
                        .iter()
                        .enumerate()
                        .filter(|(j, _)| (1 << j) & i > 0)
                        .map(|(_, (valve, flow))| (valve.clone(), *flow)),
                );
                let valves_elephant: HashMap<String, u64> = HashMap::from_iter(
                    valves_vec
                        .iter()
                        .filter(|(valve, _)| !valves_me.contains_key(valve))
                        .map(|(valve, flow)| (valve.clone(), *flow)),
                );
                let result_me = travel(&valves_me, &self.dists, Vec::from(["AA".to_owned()]), 26);
                let result_elephant = travel(
                    &valves_elephant,
                    &self.dists,
                    Vec::from(["AA".to_owned()]),
                    26,
                );
                result_me + result_elephant
            })
            .max()
            .unwrap()
    }
}

fn main() {
    Day2022_16::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_16;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_16.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_16::new(TEST_INPUT);
        assert_eq!(sol.part_one(), 1651);
        assert_eq!(sol.part_two(), 1707);
    }
}
