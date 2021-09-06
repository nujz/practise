package main

import "fmt"

func reverseWords(s string) string {
	n := len(s)
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		j, k := i, i
		for i < n && s[i] != ' ' {
			i++
		}
		for ; j < i; j++ {
			res[j] = s[i-j-1+k]
		}
		if i < n {
			res[i] = s[i]
		}
	}
	return string(res)
}

func main() {
	fmt.Println(reverseWords("abc def ghi"))
}
