import java.util.Arrays;

class Solution {
    public int[] searchRange(int[] nums, int target) {
        int lo = 0;
        int hi = nums.length - 1;
        while (lo <= hi) {
            int m = lo + (hi - lo) / 2;
            if (nums[m] == target) {
                lo = hi = m;
                // TODO optimize
                while (--lo >= 0 && nums[lo] == target) {
                }
                while (++hi < nums.length && nums[hi] == target) {
                }
                return new int[] { lo + 1, hi - 1 };
            } else if (nums[m] > target) {
                hi = m - 1;
            } else {
                lo = m + 1;
            }
        }
        return new int[] { -1, -1 };
    }

    public static void main(String[] args) {
        int[] ret = new Solution().searchRange(new int[] { 5, 7, 7, 8, 8, 10 }, 8);
        System.out.println(Arrays.toString(ret));
    }
}