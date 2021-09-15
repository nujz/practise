package main

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}

func main() {
	println(singleNumber([]int{1, 2, 1, 2, 3, 5, 6, 6, 5}))
}
