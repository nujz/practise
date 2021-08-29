package main

import (
	"math/big"
)

func uniquePaths1(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

// f(i, j) = f(i-1, j) + f(i, j-1)
func uniquePaths(m, n int) int {
	var f = make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		f[i][0] = 1
	}
	for i := 0; i < n; i++ {
		f[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}

	return f[m-1][n-1]
}

func main() {
	println(uniquePaths(3, 7))
}
