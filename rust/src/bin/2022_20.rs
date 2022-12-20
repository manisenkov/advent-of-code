use std::cell::RefCell;
use std::rc::Rc;

use aoc::solution::Solution;

const DECRYPTION_KEY: i64 = 811589153;

type NodePtr = Rc<RefCell<Node>>;

#[derive(Debug)]
struct Node {
    index: usize,
    step: i64,
    prev_ptr: Option<NodePtr>,
    next_ptr: Option<NodePtr>,
}

impl Node {
    fn to_vector(&self) -> Vec<i64> {
        let mut res = vec![self.step];
        let mut next = Rc::clone(self.next_ptr.as_ref().unwrap());
        loop {
            let next_clone = next;
            let cur = next_clone.borrow();
            if cur.index == self.index {
                break res;
            }
            res.push(cur.step);
            let next_ptr = cur.next_ptr.as_ref().unwrap();
            next = Rc::clone(next_ptr);
        }
    }
}

fn move_at(src: NodePtr, dst: NodePtr) {
    let src_prev = Rc::clone(src.borrow_mut().prev_ptr.as_ref().unwrap());
    let src_next = Rc::clone(src.borrow_mut().next_ptr.as_ref().unwrap());
    let dst_next = Rc::clone(dst.borrow_mut().next_ptr.as_ref().unwrap());

    dst_next.borrow_mut().prev_ptr = Some(Rc::clone(&src));
    dst.borrow_mut().next_ptr = Some(Rc::clone(&src));
    src_prev.borrow_mut().next_ptr = Some(Rc::clone(&src_next));
    src_next.borrow_mut().prev_ptr = Some(Rc::clone(&src_prev));
    src.borrow_mut().next_ptr = Some(Rc::clone(&dst_next));
    src.borrow_mut().prev_ptr = Some(Rc::clone(&dst));
}

fn mix(nodes: &Vec<NodePtr>) {
    for i in 0..nodes.len() {
        let mut dest = Rc::clone(&nodes[i]);
        let step = nodes[i].borrow().step;
        let total = nodes.len() as i64;
        let mut dist_left = step % (total - 1);
        if dist_left < 0 {
            dist_left = total + dist_left - 1;
        }
        if dist_left == 0 {
            continue;
        }
        while dist_left != 0 {
            dest = if dist_left < 0 {
                Rc::clone(dest.borrow().prev_ptr.as_ref().unwrap())
            } else {
                Rc::clone(dest.borrow().next_ptr.as_ref().unwrap())
            };
            dist_left -= dist_left / dist_left.abs();
        }
        move_at(Rc::clone(&nodes[i]), Rc::clone(&dest));
    }
}

fn get_nodes(nums: &[i64], decryption_key: i64) -> (Vec<NodePtr>, usize) {
    let mut nodes = Vec::<NodePtr>::new();
    let mut zero_index = 0;
    for (i, n) in nums.iter().enumerate() {
        if *n == 0 {
            zero_index = i;
        }
        let node = Rc::new(RefCell::new(Node {
            index: i,
            step: *n * decryption_key,
            prev_ptr: if i > 0 {
                Some(Rc::clone(&nodes[i - 1]))
            } else {
                None
            },
            next_ptr: None,
        }));
        if i > 0 {
            let mut prev_node = nodes[i - 1].borrow_mut();
            prev_node.next_ptr = Some(Rc::clone(&node));
        }
        nodes.push(node);
    }
    nodes[0].borrow_mut().prev_ptr = Some(Rc::clone(&nodes[nodes.len() - 1]));
    nodes[nodes.len() - 1].borrow_mut().next_ptr = Some(Rc::clone(&nodes[0]));
    (nodes, zero_index)
}

struct Day2022_20 {
    nums: Vec<i64>,
}

impl Solution<i64> for Day2022_20 {
    fn new() -> Day2022_20 {
        Day2022_20 { nums: vec![] }
    }

    fn init(&mut self, input: &str) {
        for line in input.lines() {
            self.nums.push(line.parse().unwrap())
        }
    }

    fn part_one(&mut self) -> i64 {
        let (mut nodes, zero_index) = get_nodes(&self.nums, 1);
        mix(&mut nodes);
        let v = nodes[zero_index].borrow().to_vector();
        v[1000 % v.len()] + v[2000 % v.len()] + v[3000 % v.len()]
    }

    fn part_two(&mut self) -> i64 {
        let (mut nodes, zero_index) = get_nodes(&self.nums, DECRYPTION_KEY);
        for _ in 0..10 {
            mix(&mut nodes);
        }
        let v = nodes[zero_index].borrow().to_vector();
        v[1000 % v.len()] + v[2000 % v.len()] + v[3000 % v.len()]
    }
}

fn main() {
    let mut sol = Day2022_20::new();
    sol.run_on_stdin()
}

#[cfg(test)]
mod tests {
    use crate::Day2022_20;
    use aoc::solution::Solution;

    const TEST_INPUT: &str = include_str!("../../examples/2022_20.txt");

    #[test]
    fn test_1() {
        let mut sol = Day2022_20::new();
        sol.init(TEST_INPUT);
        assert_eq!(sol.part_one(), 3);
        assert_eq!(sol.part_two(), 1623178306);
    }
}
