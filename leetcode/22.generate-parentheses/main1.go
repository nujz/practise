package main

import "fmt"

func generateParenthesis(n int) []string {
	var arr []string
	var generate func(int, int, int, []byte)

	generate = func(l int, r int, n int, pre []byte) {
		if l >= n && r >= n {
			arr = append(arr, string(pre))
			return
		}
		if l <= r {
			pre[l+r] = '('
			generate(l+1, r, n, pre)
		} else if l < n {
			pre[l+r] = '('
			generate(l+1, r, n, pre)

			pre[l+r] = ')'
			generate(l, r+1, n, pre)
		} else {
			pre[l+r] = ')'
			generate(l, r+1, n, pre)
		}
	}

	generate(0, 0, n, make([]byte, n*2))

	return arr
}

func main() {
	fmt.Println(generateParenthesis(1))
}
