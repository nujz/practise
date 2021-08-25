from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(0, len(nums)):
            for j in range(i+1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]
        return []

    def twoSum2(self, nums: List[int], target: int) -> List[int]:
        l = len(nums)
        m = {}
        for i in range(0, l):
            m[nums[i]] = i

        for i in range(0, l):
            a = target - nums[i]
            if a in m and m[a] != i:
                return [i, m[a]]
        return []

    def twoSum3(self, nums: List[int], target: int) -> List[int]:
        l = len(nums)

        m = {}
        for i in range(0, l):
            a = target - nums[i]
            if a in m:
                return [m[a], i]
            else:
                m[nums[i]] = i

        return []


if __name__ == '__main__':
    print(Solution().twoSum3([2, 3, 5, 9], 7))
