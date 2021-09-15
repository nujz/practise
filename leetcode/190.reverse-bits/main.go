package main

import "fmt"

func reverseBits(num uint32) (res uint32) {
	for i := 0; i < 32 && num > 0; i++ {
		res |= num & 1 << (31 - i)
		num >>= 1
	}
	return res
}

func main() {
	fmt.Printf("%b", reverseBits(0b11111111111111111111111111111101))
}
