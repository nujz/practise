package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		o := target - numbers[i]
		if j, ok := m[o]; ok {
			return []int{j + 1, i + 1}
		}
		m[numbers[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
