struct Solution;

impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let (mut left, mut right, length) = (0, 0, nums.len());
        while right < length {
            if nums[right] != 0 {
                nums.swap(left, right);
                left += 1;
            }
            right += 1;
        }
    }
}

fn main() {
    let mut nums = vec![0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0];
    Solution::move_zeroes(&mut nums);
    println!("{:?}", nums);
}
