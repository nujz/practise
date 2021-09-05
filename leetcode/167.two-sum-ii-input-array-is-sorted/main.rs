#![feature(test)]

extern crate test;

pub struct Solution;

impl Solution {
    pub fn two_sum2(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut m = fnv::FnvHashMap::<i32, i32>::default();
        for i in 0..numbers.len() {
            let o = target - numbers[i];
            if let Some(j) = m.get(&o) {
                return vec![j + 1, i as i32 + 1];
            }
            m.insert(numbers[i], i as _);
        }
        vec![]
    }

    pub fn two_sum1(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut m = std::collections::HashMap::<i32, i32>::new();
        for i in 0..numbers.len() {
            let o = target - numbers[i];
            if let Some(j) = m.get(&o) {
                return vec![j + 1, i as i32 + 1];
            }
            m.insert(numbers[i], i as _);
        }
        vec![]
    }

    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        for i in 0..numbers.len() {
            for j in i + 1..numbers.len() {
                if numbers[i] + numbers[j] == target {
                    return vec![i as i32 + 1, j as i32 + 1];
                }
            }
        }
        return vec![];
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use test::Bencher;

    #[bench]
    fn bench_add_two(b: &mut Bencher) {
        b.iter(|| Solution::two_sum(vec![1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 2, 7, 11, 15], 9));
    }

    #[bench]
    fn bench_add_two_hash(b: &mut Bencher) {
        b.iter(|| Solution::two_sum1(vec![1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 2, 7, 11, 15], 9));
    }

    #[bench]
    fn bench_add_two_fnv_hash(b: &mut Bencher) {
        b.iter(|| Solution::two_sum2(vec![1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 2, 7, 11, 15], 9));
    }
}
