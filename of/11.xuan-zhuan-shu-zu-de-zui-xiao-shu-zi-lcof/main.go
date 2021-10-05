package main

func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		mid := low + (high-low)/2
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

func main() {
}
