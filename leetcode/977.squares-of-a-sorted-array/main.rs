pub struct Solution;

impl Solution {
    pub fn sorted_squares(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut arr = vec![0; n];
        let (mut i, mut j, mut pos) = (0, n, n)
        while i < j {
            pos -= 1;
            if nums[i].abs() > nums[j - 1].abs() {
                arr[pos] = nums[i] * nums[i];
                i += 1;
            } else {
                j -= 1;
                arr[pos] = nums[j] * nums[j];
            }
        }
        arr
    }
}

fn main() {
    println!("{:?}", Solution::sorted_squares(vec![-4, -1, 0, 3, 10]));
}
