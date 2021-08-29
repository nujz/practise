package main

import "fmt"

func moveZeroes(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			j := i + 1
			for j < l && nums[j] == 0 {
				j++
			}
			if j < l {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				break
			}
		}
	}
}

func main() {
	a := []int{0, 1, 1, 0, 1, 1}
	moveZeroes(a)
	fmt.Println(a)
}
