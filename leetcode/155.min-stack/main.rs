struct MinStack {
    data: Vec<i32>,
    min: Vec<i32>,
}

impl MinStack {
    fn new() -> Self {
        Self {
            data: Vec::new(),
            min: Vec::new(),
        }
    }

    fn push(&mut self, val: i32) {
        self.data.push(val);
        if self.min.len() == 0 {
            self.min.push(val);
        } else {
            let min = self.min.last().unwrap();
            if *min >= val {
                self.min.push(val);
            }
        }
    }

    fn pop(&mut self) {
        self.data.pop().map(|a| {
            let min = self.min.last().unwrap();
            if *min == a {
                self.min.pop();
            }
        });
    }

    fn top(&self) -> i32 {
        *self.data.last().unwrap()
    }

    fn get_min(&self) -> i32 {
        *self.min.last().unwrap()
    }
}

fn main() {}
