use std::collections::HashMap;

use aoc::solution::Solution;

enum ConsoleLine {
    ChangeDir(String),
    List,
    Dir,
    File(usize),
}

impl ConsoleLine {
    fn parse(line: &str) -> ConsoleLine {
        let parts: Vec<_> = line.trim().split(" ").collect();
        if parts[0] == "$" {
            match parts[1] {
                "cd" => ConsoleLine::ChangeDir(parts[2].into()),
                "ls" => ConsoleLine::List,
                _ => panic!("invalid command"),
            }
        } else {
            match parts[0] {
                "dir" => ConsoleLine::Dir,
                size => ConsoleLine::File(size.parse().unwrap()),
            }
        }
    }
}

fn join_paths(path: &str, to_append: &str) -> String {
    if to_append == "/" {
        if path == "" {
            to_append.into()
        } else {
            path.into()
        }
    } else {
        if path == "" {
            String::from("/") + &to_append
        } else {
            path.to_owned() + &"/" + &to_append
        }
    }
}

struct Day2022_07 {
    dir_sizes: HashMap<String, usize>,
}

impl Solution<usize> for Day2022_07 {
    fn new(input: &str) -> Day2022_07 {
        let empty_string = &String::from("");
        let mut path = Vec::<String>::new();
        let mut dir_sizes = HashMap::new();
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
                ConsoleLine::File(file_size) => {
                    for dir_name in path.iter() {
                        dir_sizes
                            .entry(dir_name.to_string())
                            .and_modify(|dir_size| *dir_size += file_size)
                            .or_insert(file_size);
                    }
                }
                _ => (),
            };
        }
        Day2022_07 { dir_sizes }
    }

    fn part_one(&mut self) -> usize {
        self.dir_sizes
            .values()
            .filter(|&&size| size <= 100000)
            .sum()
    }

    fn part_two(&mut self) -> usize {
        let size_left = self.dir_sizes.get("/").unwrap() - 40000000;
        let mut candidates: Vec<_> = self
            .dir_sizes
            .values()
            .filter(|&&size| size >= size_left)
            .map(|&size| size)
            .collect();
        candidates.sort();
        candidates[0]
    }
}

fn main() {
    Day2022_07::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2022_07;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_07.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_07::new(TEST_INPUT);
        assert_eq!(sol.part_one(), 95437);
        assert_eq!(sol.part_two(), 24933642);
    }
}
