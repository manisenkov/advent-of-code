use std::collections::HashSet;

use aoc::solution::Solution;

#[derive(Clone, Copy)]
enum Guide {
    Left(i32),
    Right(i32),
}

#[derive(Clone, Copy)]
enum Direction {
    North,
    East,
    South,
    West,
}

impl Direction {
    fn rotate(self, guide: Guide) -> Self {
        match (self, guide) {
            (Direction::North, Guide::Left(_)) => Direction::West,
            (Direction::East, Guide::Left(_)) => Direction::North,
            (Direction::South, Guide::Left(_)) => Direction::East,
            (Direction::West, Guide::Left(_)) => Direction::South,
            (Direction::North, Guide::Right(_)) => Direction::East,
            (Direction::East, Guide::Right(_)) => Direction::South,
            (Direction::South, Guide::Right(_)) => Direction::West,
            (Direction::West, Guide::Right(_)) => Direction::North,
        }
    }
}

#[derive(Clone, Copy, Eq, PartialEq, Hash)]
struct Pos(i32, i32);

impl Pos {
    fn jump(self, direction: Direction, guide: Guide) -> Pos {
        let Pos(c, r) = self;
        let dist = match guide {
            Guide::Left(d) => d,
            Guide::Right(d) => d,
        };
        match direction {
            Direction::North => Pos(c, r + dist),
            Direction::East => Pos(c + dist, r),
            Direction::South => Pos(c, r - dist),
            Direction::West => Pos(c - dist, r),
        }
    }

    fn walk(self, direction: Direction, guide: Guide) -> impl Iterator<Item = Pos> {
        let Pos(c, r) = self;
        let dist = match guide {
            Guide::Left(d) => d,
            Guide::Right(d) => d,
        };
        (1..=dist).map(move |i| match direction {
            Direction::North => Pos(c, r + i),
            Direction::East => Pos(c + i, r),
            Direction::South => Pos(c, r - i),
            Direction::West => Pos(c - i, r),
        })
    }
}

struct Day2016_01 {
    instructions: Vec<Guide>,
}

impl Solution<i32> for Day2016_01 {
    fn new(input: &str) -> Day2016_01 {
        Day2016_01 {
            instructions: input
                .split(", ")
                .map(|r| {
                    let dist: i32 = (&r[1..]).trim_end().parse().unwrap();
                    match r.chars().nth(0) {
                        Some('L') => Guide::Left(dist),
                        Some('R') => Guide::Right(dist),
                        _ => panic!("unknown direction"),
                    }
                })
                .collect(),
        }
    }

    fn part_one(&mut self) -> i32 {
        let mut pos = Pos(0, 0);
        let mut dir: Direction = Direction::North;
        for &guide in self.instructions.iter() {
            dir = dir.rotate(guide);
            pos = pos.jump(dir, guide);
        }
        pos.0.abs() + pos.1.abs()
    }

    fn part_two(&mut self) -> i32 {
        let mut visited = HashSet::<Pos>::new();
        let mut pos = Pos(0, 0);
        let mut dir = Direction::North;
        visited.insert(pos);
        'walking: for &guide in self.instructions.iter() {
            dir = dir.rotate(guide);
            let steps = pos.walk(dir, guide);
            for step in steps {
                pos = step;
                if visited.contains(&step) {
                    break 'walking;
                }
                visited.insert(step);
            }
        }
        pos.0.abs() + pos.1.abs()
    }
}

fn main() {
    Day2016_01::run_on_stdin();
}

#[cfg(test)]
mod tests {
    use crate::Day2016_01;
    use aoc::solution::Solution;

    #[test]
    fn test_1() {
        let mut sol = Day2016_01::new("R2, L3");
        assert_eq!(sol.part_one(), 5);
    }

    #[test]
    fn test_2() {
        let mut sol = Day2016_01::new("R2, R2, R2");
        assert_eq!(sol.part_one(), 2);
    }

    #[test]
    fn test_3() {
        let mut sol = Day2016_01::new("R5, L5, R5, R3");
        assert_eq!(sol.part_one(), 12);
    }

    #[test]
    fn test_4() {
        let mut sol = Day2016_01::new("R8, R4, R4, R8");
        assert_eq!(sol.part_two(), 4);
    }
}
