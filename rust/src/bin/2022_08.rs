use aoc::solution::Solution;

struct Day2022_08 {
    talls: Vec<Vec<usize>>,
}

impl Solution<usize> for Day2022_08 {
    fn new() -> Day2022_08 {
        Day2022_08 { talls: Vec::new() }
    }

    fn init(&mut self, input: &str) {
        self.talls = input
            .lines()
            .map(|line| {
                line.trim()
                    .chars()
                    .map(|c| c.to_string().parse::<usize>().unwrap())
                    .collect()
            })
            .collect()
    }

    fn part_one(&mut self) -> usize {
        let nrows = self.talls.len();
        let ncols = self.talls[0].len();
        self.talls
            .iter()
            .enumerate()
            .map(|(i, row)| {
                let it = row.iter().enumerate().map(|(j, h)| {
                    if i == 0 || j == 0 || i == nrows - 1 || j == ncols - 1 {
                        1
                    } else {
                        let up = (0..i).all(|n| self.talls[n][j] < *h);
                        let down = (i + 1..nrows).all(|n| self.talls[n][j] < *h);
                        let left = (0..j).all(|n| self.talls[i][n] < *h);
                        let right = (j + 1..ncols).all(|n| self.talls[i][n] < *h);
                        if up || down || left || right {
                            1
                        } else {
                            0
                        }
                    }
                });
                it.sum::<usize>()
            })
            .sum()
    }

    fn part_two(&mut self) -> usize {
        let nrows = self.talls.len();
        let ncols = self.talls[0].len();
        self.talls
            .iter()
            .enumerate()
            .map(|(i, row)| {
                if i == 0 || i == nrows - 1 {
                    return 0;
                }
                row.iter()
                    .enumerate()
                    .map(|(j, h)| {
                        if j == 0 || j == ncols - 1 {
                            return 0;
                        }
                        let up = (0..i).rev().take_while(|n| self.talls[*n][j] < *h).count();
                        let down = (i + 1..nrows)
                            .take_while(|n| self.talls[*n][j] < *h)
                            .count();
                        let left = (0..j).rev().take_while(|n| self.talls[i][*n] < *h).count();
                        let right = (j + 1..ncols)
                            .take_while(|n| self.talls[i][*n] < *h)
                            .count();
                        (if up < i { up + 1 } else { up })
                            * (if down < nrows - i - 1 { down + 1 } else { down })
                            * (if left < j { left + 1 } else { left })
                            * (if right < ncols - j - 1 {
                                right + 1
                            } else {
                                right
                            })
                    })
                    .max()
                    .unwrap()
            })
            .max()
            .unwrap()
    }
}

fn main() {
    let mut sol = Day2022_08::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_08;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_08.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_08::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 21);
        assert_eq!(sol.part_two(), 8);
    }
}
