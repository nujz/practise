package main

import (
	"fmt"
	"strconv"
)

func solve(s, t string) string {
	carry, m, n := 0, len(s), len(t)

	var res []byte
	var p int
	if m > n {
		res = make([]byte, m+1)
		p = m
	} else {
		res = make([]byte, n+1)
		p = n
	}

	for m > 0 || n > 0 || carry != 0 {
		if m > 0 {
			carry += int(s[m-1] - '0')
			m--
		}
		if n > 0 {
			carry += int(t[n-1] - '0')
			n--
		}
		res[p] = byte(carry%10) + '0'
		p--
		carry /= 10
	}

	if res[0] == 0 {
		res = res[1:]
	}

	return string(res)
}

func solve1(s, t string) string {
	m, n := len(s), len(t)

	o := 0
	res := ""

	for m > 0 || n > 0 || o != 0 {
		if m > 0 {
			o += int(s[m-1] - '0')
		}
		if n > 0 {
			o += int(t[n-1] - '0')
		}
		res = strconv.Itoa(o%10) + res
		o /= 10
		m--
		n--
	}

	return res
}

func main() {
	fmt.Println(solve("0", "0"))
	fmt.Println(solve("9991", "10"))
}
