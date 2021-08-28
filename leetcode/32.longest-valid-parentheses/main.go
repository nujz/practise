package main

import "fmt"

// 规律：
//   左右括号数量相同的情况下，无论怎么排列，结果都是最大长度
//   接下来可以往里面放入左括号或者右括号
//  如果
//   1. 右括号多余：从左往右处理，多余的右括号忽略
//   2. 左括号多余：从右往左处理，多余的左括号忽略
func longestValidParentheses(s string) int {
	var l, r int
	var m int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}

		if l == r {
			m = max(m, l*2)
		} else if r > l {
			l, r = 0, 0
		}
	}

	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			r++
		} else {
			l++
		}

		if l == r {
			m = max(m, l*2)
		} else if r < l {
			l, r = 0, 0
		}
	}

	return m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(longestValidParentheses("()("))
}
