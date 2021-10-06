package main

import "fmt"

func missingNumber1(nums []int) int {
	n := len(nums)
	if n-1 == nums[n-1] {
		return n
	}
	res := 0
	for i := 0; i < n; i++ {
		res += i - nums[i]
	}
	return res + n
}

func missingNumber(nums []int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] == m {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return l
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2}))
	fmt.Println(missingNumber([]int{}))
}
