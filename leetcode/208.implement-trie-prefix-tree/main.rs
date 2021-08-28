pub struct Trie {
    root: Option<Box<Node>>,
}

struct Node {
    is_end: bool,
    children: [Option<Box<Node>>; 26],
}

// TODO 简化代码
impl Trie {
    pub fn new() -> Self {
        Self {
            root: Some(Box::new(Node {
                is_end: false,
                children: Default::default(),
            })),
        }
    }

    pub fn insert(&mut self, word: String) {
        let mut node = self.root.as_mut().unwrap();
        for mut c in word.as_bytes() {
            let c = (c - b'a') as usize;
            if node.children[c].is_none() {
                node.children[c] = Some(Box::new(Node {
                    is_end: false,
                    children: Default::default(),
                }));
            }
            node = node.children[c].as_mut().unwrap();
        }
        node.is_end = true;
    }

    pub fn search(&self, word: String) -> bool {
        let node = self.prefix_node(&word);
        match node {
            None => false,
            Some(n) => n.is_end,
        }
    }

    pub fn starts_with(&self, prefix: String) -> bool {
        self.prefix_node(&prefix).is_some()
    }

    fn prefix_node(&self, prefix: &String) -> Option<&Node> {
        let mut node = self.root.as_ref().unwrap();
        for mut c in prefix.as_bytes() {
            let c = (c - b'a') as usize;
            if node.children[c].is_none() {
                return None;
            }
            node = node.children[c].as_ref().unwrap();
        }
        return Some(node);
    }
}

fn main() {}
