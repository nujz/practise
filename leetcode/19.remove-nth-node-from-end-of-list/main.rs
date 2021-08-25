#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub struct Solution;

impl Solution {
    pub fn remove_nth_from_end(mut head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut len = 0;
        let mut cur = &head;
        while let Some(node) = cur {
            len += 1;
            cur = &node.next;
        }
        if n == len {
            return head.map_or(None, |mut node| node.next.take());
        }

        let mut cur = head.as_mut();
        let mut i = 0;
        while let Some(node) = cur {
            i += 1;
            if i + n == len {
                node.next = node.next.as_mut().map_or(None, |node| node.next.take());
                return head;
            }
            cur = node.next.as_mut();
        }

        head
    }
}

fn main() {
    let mut head = Some(Box::new(ListNode::new(1)));
    head.as_mut()
        .map(|node| node.next.replace(Box::new(ListNode::new(2))));

    let r = Solution::remove_nth_from_end(head, 1);
    println!("{:?}", r);
}
