use aoc::solution::Solution;

enum Command {
    Noop,
    AddX(isize),
}

impl Command {
    fn parse(input: &str) -> Command {
        if input.trim() == "noop" {
            Command::Noop
        } else {
            let mut s = input.trim().split(" ");
            s.next();
            Command::AddX(s.next().unwrap().parse().unwrap())
        }
    }
}

struct Day2022_10 {
    commands: Vec<Command>,
    display: String,
}

impl Solution<isize, String> for Day2022_10 {
    fn new() -> Day2022_10 {
        Day2022_10 {
            commands: vec![],
            display: String::new(),
        }
    }

    fn init(&mut self, input: &str) {
        self.commands.extend(input.lines().map(Command::parse));
    }

    fn part_one(&mut self) -> isize {
        let mut strength = 1;
        let mut sum: isize = 0;
        let mut cycle = 0;

        let mut draw = |s: &isize| {
            let col = cycle % 40;
            if col == 0 {
                self.display.push_str("\n");
            }
            self.display.push_str(if col >= s - 1 && col <= s + 1 {
                "█"
            } else {
                "."
            });
            cycle += 1;
            if (cycle - 20) % 40 == 0 {
                sum += cycle * s;
            }
        };

        for cmd in self.commands.iter() {
            match cmd {
                Command::Noop => {
                    draw(&strength);
                }
                Command::AddX(x) => {
                    draw(&strength);
                    draw(&strength);
                    strength += x;
                }
            };
        }
        sum
    }

    fn part_two(&mut self) -> String {
        self.display.to_string()
    }
}

fn main() {
    let mut sol = Day2022_10::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_10;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_10.txt");
    const EXPECTED: &str = "
██..██..██..██..██..██..██..██..██..██..
███...███...███...███...███...███...███.
████....████....████....████....████....
█████.....█████.....█████.....█████.....
██████......██████......██████......████
███████.......███████.......███████.....";

    #[test]
    fn test_1() {
        let mut sol = Day2022_10::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 13140);
        assert_eq!(sol.part_two(), EXPECTED);
    }
}
