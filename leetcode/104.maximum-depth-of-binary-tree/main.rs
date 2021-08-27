use std::{cell::RefCell, rc::Rc};

struct Solution;

#[derive(Debug)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl Solution {
    pub fn max_depth0(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let root = match root {
            Some(r) => r,
            None => return 0,
        };

        let l = Self::max_depth(root.borrow().left.clone()) + 1;
        let r = Self::max_depth(root.borrow().right.clone()) + 1;

        if l > r {
            l
        } else {
            r
        }
    }

    pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let root = match root {
            Some(r) => r,
            None => return 0,
        };

        let mut depth = 0;
        let mut queue = vec![root];
        while queue.len() > 0 {
            let size = queue.len();
            for _ in 0..size {
                let node = queue.remove(0);
                match node.borrow().left.clone() {
                    Some(n) => queue.push(n),
                    None => {}
                };
                match node.borrow().right.clone() {
                    Some(n) => queue.push(n),
                    None => {}
                };
            }
            depth += 1;
        }
        depth
    }
}

fn main() {
    let max_depth = Solution::max_depth(None);
    println!("{:?}", max_depth);
}
