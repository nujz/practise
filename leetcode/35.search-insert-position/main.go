package main

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		m := lo + (hi-lo)/2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			hi = m - 1
		} else {
			lo = m + 1
		}
	}
	if nums[lo] < target {
		return lo + 1
	}
	return lo
}

func main() {

}
