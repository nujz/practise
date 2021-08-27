package main

func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			j := i + 1
			for j < len(prices) && prices[j-1] <= prices[j] {
				j++
			}
			sum += prices[j-1] - prices[i-1]
			i = j
		}
	}
	return sum
}

func main() {
}
