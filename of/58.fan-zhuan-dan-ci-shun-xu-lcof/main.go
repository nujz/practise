package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	b := new(strings.Builder)

	lo := 0
	for lo < len(s) && s[lo] == ' ' {
		lo++
	}

	i := len(s) - 1
	for i >= lo {
		for i >= lo && s[i] == ' ' {
			i--
		}
		j := i
		for i >= lo && s[i] != ' ' {
			i--
		}
		b.WriteString(s[i+1 : j+1])
		if i >= lo {
			b.WriteByte(' ')
		}
	}

	return b.String()
}

func main() {
	fmt.Printf("\"%s\"", reverseWords("  a good   example  "))
}
