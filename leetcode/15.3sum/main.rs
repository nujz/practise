pub struct Solution;

impl Solution {
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort_unstable();

        let l = nums.len();
        let mut res = vec![];
        for i in 0..l {
            if i > 0 && nums[i - 1] == nums[i] {
                continue;
            }

            let mut k = l - 1;

            for j in i + 1..l {
                if j > i + 1 && nums[j - 1] == nums[j] {
                    continue;
                }

                while j < k && nums[i] + nums[j] + nums[k] > 0 {
                    k -= 1;
                }
                if j == k {
                    break;
                }
                if nums[i] + nums[j] + nums[k] == 0 {
                    res.push(vec![nums[i], nums[j], nums[k]])
                }
            }
        }

        return res;
    }

    pub fn three_sum1(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort_unstable();

        let l = nums.len();
        let mut res = vec![];
        for i in 0..l {
            if i > 0 && nums[i - 1] == nums[i] {
                continue;
            }
            for j in i + 1..l {
                if j > i + 1 && nums[j - 1] == nums[j] {
                    continue;
                }

                if let Ok(k) = nums[j + 1..l].binary_search(&(0 - nums[i] - nums[j])) {
                    res.push(vec![nums[i], nums[j], nums[j + 1 + k]]);
                }
            }
        }

        return res;
    }
}

fn main() {
    println!("{:?}", Solution::three_sum(vec![-1, 0, 1, 2, -1, -4]));
}
