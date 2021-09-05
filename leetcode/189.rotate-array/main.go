package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	res := make([]int, n)
	for i := range nums {
		res[(i+k)%n] = nums[i]
	}
	copy(nums, res)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	rotate(a, 2)
	fmt.Println()
}
