package main

func checkInclusion(s1 string, s2 string) bool {
	// Todo
	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt [26]int
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, v := range cnt[:] {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}

		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}

		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}

		if diff == 0 {
			return true
		}
	}
	return false
}

func checkInclusion1(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	var cnt1, cnt2 [26]int
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := l1; i < l2; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-l1]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}

	return false
}

func main() {
}
