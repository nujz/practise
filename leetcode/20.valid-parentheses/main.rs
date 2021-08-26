struct Solution;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack = vec![];
        for r in s.as_bytes() {
            match *r as _ {
                ']' => match stack.pop() {
                    Some(l) if l == '[' => {}
                    _ => return false,
                },
                '}' => match stack.pop() {
                    Some(l) if l == '{' => {}
                    _ => return false,
                },
                ')' => match stack.pop() {
                    Some(l) if l == '(' => {}
                    _ => return false,
                },
                c => stack.push(c),
            }
        }
        stack.len() == 0
    }
}

fn main() {
    println!("{:?}", Solution::is_valid("([]{})".into()));
}
