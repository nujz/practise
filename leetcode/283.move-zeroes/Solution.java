import java.util.Arrays;

class Solution {
    public void moveZeroes(int[] nums) {
        int left = 0;
        int right = 0;
        int length = nums.length;
        while (right < length) {
            if (nums[right] != 0) {
                int tmp = nums[left];
                nums[left] = nums[right];
                nums[right] = tmp;
                left++;
            }
            right++;
        }
    }

    public static void main(String[] args) {
        int[] nums = new int[] { 0, 1, 0, 1, 0, 1, 1, 1, 0, 0 };
        new Solution().moveZeroes(nums);
        System.out.println(Arrays.toString(nums));
    }
}