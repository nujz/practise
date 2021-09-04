
from typing import List


class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        for i in range(0, len(nums)):
            nums[i] = nums[i] * nums[i]
        nums.sort()
        return nums


if __name__ == '__main__':
    print(Solution().sortedSquares([-4, -1, 0, 5, 9]))
