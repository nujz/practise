package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	l, r, p := 0, n-1, n-1
	for l <= r {
		if abs(nums[l]) > abs(nums[r]) {
			res[p] = nums[l] * nums[l]
			l++
		} else {
			res[p] = nums[r] * nums[r]
			r--
		}
		p--
	}
	return res
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 2, 5}))
}
