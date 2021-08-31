package main

import "fmt"

type Holder struct {
	arr []string
}

func generateParenthesis(n int) []string {
	holder := Holder{}
	generate(&holder, 0, 0, n, make([]byte, n*2))
	return holder.arr
}

func generate(holder *Holder, l int, r int, n int, pre []byte) {
	if l >= n && r >= n {
		holder.arr = append(holder.arr, string(pre))
		return
	}
	if l <= r {
		pre[l+r] = '('
		generate(holder, l+1, r, n, pre)
	} else if l < n {
		pre[l+r] = '('
		generate(holder, l+1, r, n, pre)

		pre[l+r] = ')'
		generate(holder, l, r+1, n, pre)
	} else {
		pre[l+r] = ')'
		generate(holder, l, r+1, n, pre)
	}
}

func main() {
	fmt.Println(generateParenthesis(1))
}
