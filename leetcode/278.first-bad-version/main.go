package main

func isBadVersion(n int) bool {
	return false
}

func firstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		m := (lo + hi) / 2
		if isBadVersion(m) {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return lo
}

func main() {

}
