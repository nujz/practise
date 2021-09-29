package main

func numWays(n int) int {
	x, y := 1, 1
	for i := 2; i <= n; i++ {
		tmp := (x + y) % 1000000007
		x = y
		y = tmp
	}
	return y
}

func main() {
	println(numWays(0))
	println(numWays(1))
	println(numWays(2))
	println(numWays(92))
}
