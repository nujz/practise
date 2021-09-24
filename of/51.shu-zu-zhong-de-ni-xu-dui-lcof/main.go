package main

import "fmt"

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, lo, hi int) int {
	if lo >= hi {
		return 0
	}
	mid := lo + (hi-lo)/2
	res := mergeSort(nums, lo, mid) + mergeSort(nums, mid+1, hi)
	var tmp []int
	i, j := lo, mid+1
	for i <= mid && j <= hi {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			res += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, nums[i])
		res += j - (mid + 1)
		i++
	}
	for j <= hi {
		tmp = append(tmp, nums[j])
		j++
	}
	for i := lo; i <= hi; i++ {
		nums[i] = tmp[i-lo]
	}
	return res
}

func main() {
	nums := []int{5, 4, 3, 2, 1}
	println(reversePairs(nums))
	fmt.Println(nums)
}
