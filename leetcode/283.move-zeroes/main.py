from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        left, right, length = 0, 0, len(nums)
        while right < length:
            if nums[right] != 0:
                nums[left], nums[right] = nums[right], nums[left]
                left += 1
            right += 1


if __name__ == "__main__":
    lst = [1, 0, 1, 0, 1, 1, 1, 0]
    Solution().moveZeroes(lst)
    print(lst)
