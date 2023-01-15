use std::fmt::Display;
use std::io::{self, Read};
use std::time::{Duration, Instant};

fn format_dur(dur: Duration) -> String {
    if dur.as_secs() > 0 {
        format!("{}.{} s", dur.as_secs(), dur.as_millis())
    } else if dur.as_millis() > 0 {
        format!("{}.{} ms", dur.as_millis(), dur.as_micros())
    } else if dur.as_micros() > 0 {
        format!("{}.{} μs", dur.as_micros(), dur.as_nanos() % 1000)
    } else {
        format!("{} ns", dur.as_nanos())
    }
}

pub trait Solution<T1, T2 = T1>
where
    T1: Display + PartialEq<T1>,
    T2: Display + PartialEq<T2>,
{
    fn new(input: &str) -> Self
    where
        Self: Sized;
    fn part_one(&mut self) -> T1;
    fn part_two(&mut self) -> T2;

    fn run_on_stdin()
    where
        Self: Sized,
    {
        let start_inst = Instant::now();
        let mut stdin = io::stdin();
        let mut input = String::new();
        stdin.read_to_string(&mut input).unwrap();
        let reading_time = start_inst.elapsed();
        println!(" -- Reading time: {}", format_dur(reading_time));

        Self::run(input.as_str());
    }

    fn run(input: &str)
    where
        Self: Sized,
    {
        let start_inst = Instant::now();
        let mut solution = Self::new(input);
        let init_time = start_inst.elapsed();
        println!(" -- Init time: {}", format_dur(init_time));

        let part_one_start_inst = Instant::now();
        let part_one_result = solution.part_one();
        let part_one_time = part_one_start_inst.elapsed();
        println!("Part one: {}", part_one_result);
        println!(" -- Part one time: {}", format_dur(part_one_time));

        let part_two_start_inst = Instant::now();
        let part_two_result = solution.part_two();
        let part_two_time = part_two_start_inst.elapsed();
        println!("Part two: {}", part_two_result);
        println!(" -- Part two time: {}", format_dur(part_two_time));
    }
}
