package main

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return true
	}
	if n == 1 {
		return true
	}
	for n%2 == 0 {
		n = n / 2
		if n == 1 {
			return true
		}
	}
	return false
}

func isPowerOfTwo1(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	println(isPowerOfTwo(9))
}
