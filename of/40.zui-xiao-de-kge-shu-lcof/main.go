package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(val interface{}) {
	v := val.(int)
	*h = append(*h, v)
}

func (h *MinHeap) Pop() interface{} {
	lst := len(*h) - 1
	res := (*h)[lst]
	*h = (*h)[:lst]
	return res
}

func getLeastNumbers(nums []int, k int) []int {
	a := MinHeap(nums)
	h := &a
	heap.Init(h)
	var res []int
	for ; k > 0; k-- {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

func main() {
	nums := []int{1, 3, 2, 5, 4, 7, 6}
	fmt.Println(getLeastNumbers(nums, 3))
	fmt.Println(nums)
}
