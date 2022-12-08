use std::collections::HashMap;

use aoc::solution::Solution;

enum ConsoleLine {
    ChangeDir(String),
    List,
    Dir(String),
    File(usize, String),
}

impl ConsoleLine {
    fn parse(line: &str) -> ConsoleLine {
        let parts: Vec<&str> = line.trim().split(" ").collect();
        if parts[0] == "$" {
            match parts[1] {
                "cd" => ConsoleLine::ChangeDir(parts[2].to_owned()),
                "ls" => ConsoleLine::List,
                _ => panic!("invalid command"),
            }
        } else {
            match parts[0] {
                "dir" => ConsoleLine::Dir(parts[1].to_owned()),
                size => ConsoleLine::File(size.parse().unwrap(), parts[1].to_owned()),
            }
        }
    }
}

fn join_paths(path: &str, to_append: &str) -> String {
    if to_append == "/" {
        if path == "" {
            to_append.to_owned()
        } else {
            path.to_owned()
        }
    } else {
        if path == "" {
            "/".to_owned() + &to_append.to_owned()
        } else {
            path.to_owned() + &"/".to_owned() + &to_append.to_owned()
        }
    }
}

struct Day2022_07 {
    dir_sizes: HashMap<String, usize>,
}

impl Solution<usize> for Day2022_07 {
    fn new() -> Day2022_07 {
        Day2022_07 {
            dir_sizes: HashMap::new(),
        }
    }

    fn init(&mut self, input: &str) {
        let empty_string = &String::from("");
        let mut path = Vec::<String>::new();
        for line in input.lines().map(ConsoleLine::parse) {
            let cur_dir = path.last().unwrap_or(empty_string);
            match line {
                ConsoleLine::ChangeDir(dir_name) => {
                    if dir_name == ".." {
                        path.pop();
                    } else {
                        path.push(join_paths(&cur_dir, &dir_name));
                    };
                }
                ConsoleLine::File(file_size, _) => {
                    for dir_name in path.iter() {
                        self.dir_sizes
                            .entry(dir_name.to_string())
                            .and_modify(|dir_size| *dir_size += file_size)
                            .or_insert(file_size);
                    }
                }
                _ => (),
            };
        }
    }

    fn part_one(&mut self) -> usize {
        self.dir_sizes
            .values()
            .filter(|size| **size <= 100000)
            .sum()
    }

    fn part_two(&mut self) -> usize {
        let size_left = self.dir_sizes.get("/").unwrap() - 40000000;
        let mut candidates: Vec<usize> = self
            .dir_sizes
            .values()
            .filter(|size| **size >= size_left)
            .map(|size| *size)
            .collect();
        candidates.sort();
        candidates[0]
    }
}

fn main() {
    let mut sol = Day2022_07::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_07;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_07.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_07::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 95437);
        assert_eq!(sol.part_two(), 24933642);
    }
}
