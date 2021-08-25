struct Solution;

impl Solution {
    fn two_sum1(nums: Vec<i32>, target: i32) -> Vec<i32> {
        for i in 0..nums.len() {
            for j in i + 1..nums.len() {
                if nums[i] + nums[j] == target {
                    return vec![i as _, j as _];
                }
            }
        }
        vec![]
    }

    fn two_sum2(nums: Vec<i32>, target: i32) -> Vec<i32> {
        use std::collections::HashMap;
        let mut m = HashMap::new();

        for i in 0..nums.len() {
            let a = target - nums[i];
            let t = m.get(&a);
            match t {
                Some(j) => return vec![i as _, *j as _],
                None => m.insert(nums[i], i),
            };
        }

        vec![]
    }
}

fn main() {
    println!("{:?}", Solution::two_sum2(vec![2, 7, 11, 15], 9));
}
