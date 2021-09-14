package main

func hammingWeight(num uint32) (ones int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ones++
		}
	}
	return
}

func hammingWeight1(num uint32) int {
	res := 0
	for num > 0 {
		res += int(num % 2)
		num = num / 2
	}
	return res
}

func main() {
	println(hammingWeight(0b111))
}
