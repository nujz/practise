package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	j := 0
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok {
			res = max(res, len(m))
			for ; j <= v; j++ {
				delete(m, s[j])
			}
		}
		m[s[i]] = i
	}
	res = max(res, len(m))
	return res
}

func main() {
	println(lengthOfLongestSubstring("ababc"))
}
