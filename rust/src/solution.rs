use std::fmt::Display;
use std::io::{self, Read};
use std::time::Instant;

pub trait Solution<T1, T2 = T1>
where
    T1: Display + PartialEq<T1>,
    T2: Display + PartialEq<T2>,
{
    fn new() -> Self
    where
        Self: Sized;
    fn init(&mut self, input: &str);
    fn part_one(&mut self) -> T1;
    fn part_two(&mut self) -> T2;

    fn run_on_stdin(&mut self) {
        let mut stdin = io::stdin();
        let mut input = String::new();
        stdin.read_to_string(&mut input).unwrap();
        self.run(input.trim())
    }

    fn run(&mut self, input: &str) {
        let start_instant = Instant::now();

        self.init(input);
        let init_time = start_instant.elapsed().as_nanos();
        println!("Init time {} ns", init_time);

        let t1 = self.part_one();
        let part_one_time = start_instant.elapsed().as_nanos() - init_time;
        println!("Part one: {}", t1);
        println!("Time spent on part one: {} ns", part_one_time);

        let t2 = self.part_two();
        let part_two_time = start_instant.elapsed().as_nanos() - part_one_time;
        println!("Part two: {}", t2);
        println!("Time spent on part two: {} ns", part_two_time);
    }
}
