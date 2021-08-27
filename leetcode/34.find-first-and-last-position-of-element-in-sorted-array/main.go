package main

func searchRange(nums []int, target int) []int {
	return []int{searchFirstEquals(nums, target), searchLastEquals(nums, target)}
}

func searchFirstEquals(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		m := lo + (hi-lo)/2
		if nums[m] == target {
			if m == 0 || nums[m-1] != target {
				return m
			}
			hi = m - 1
		} else if nums[m] > target {
			hi = m - 1
		} else {
			lo = m + 1
		}
	}
	return -1
}

func searchLastEquals(nums []int, target int) int {
	length := len(nums) - 1
	lo, hi := 0, length
	for lo <= hi {
		m := lo + (hi-lo)/2
		if nums[m] == target {
			if m == length || nums[m+1] != target {
				return m
			}
			lo = m + 1
		} else if nums[m] > target {
			hi = m - 1
		} else {
			lo = m + 1
		}
	}
	return -1
}

// TODO
func SearchFirstGreaterTo(nums []int, target int) int {
	return -1
}

// TODO
func SearchFirstLessTo(nums []int, target int) int {
	return -1
}

func main() {

}
