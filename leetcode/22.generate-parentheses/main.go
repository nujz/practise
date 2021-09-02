package main

import "fmt"

func generateParenthesis(n int) []string {
	var arr []string
	generate(&arr, 0, 0, n, make([]byte, n*2))
	return arr
}

func generate(arr *[]string, l int, r int, n int, pre []byte) {
	if l >= n && r >= n {
		*arr = append(*arr, string(pre))
		return
	}
	if l <= r {
		pre[l+r] = '('
		generate(arr, l+1, r, n, pre)
	} else if l < n {
		pre[l+r] = '('
		generate(arr, l+1, r, n, pre)

		pre[l+r] = ')'
		generate(arr, l, r+1, n, pre)
	} else {
		pre[l+r] = ')'
		generate(arr, l, r+1, n, pre)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
