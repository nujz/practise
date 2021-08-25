package main

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		if j, ok := m[a]; ok {
			return []int{i, j}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
}
