package main

import "fmt"

type Heap struct {
	data []int
	cmp  func(int, int) bool
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Push(val int) {
	h.data = append(h.data, val)
	h.up(len(h.data) - 1)
}

func (h *Heap) Pop() int {
	l := len(h.data) - 1
	h.Swap(0, l)
	h.down(0, l)
	res := h.data[l]
	h.data = h.data[:l]
	return res
}

func (h *Heap) Top() int {
	// TODO
	return h.data[0]
}

func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) Cmp(i, j int) bool {
	return h.cmp(h.data[i], h.data[j])
}

func (h *Heap) up(c int) {
	for {
		p := (c - 1) / 2
		if c == p || h.Cmp(p, c) {
			break
		}
		h.Swap(c, p)
		c = p
	}
}

func (h *Heap) down(p, l int) {
	for {
		c1 := 2*p + 1
		c2 := 2*p + 2
		if c1 >= l {
			break
		}
		if c2 < l && !h.Cmp(c1, c2) {
			c1 = c2
		}
		if h.Cmp(p, c1) {
			break
		}
		h.Swap(p, c1)
		p = c1
	}
}

func MaxHeap() *Heap {
	return &Heap{
		cmp: func(a, b int) bool {
			return a > b
		},
	}
}

func MinHeap() *Heap {
	return &Heap{
		cmp: func(a, b int) bool {
			return a < b
		},
	}
}

type MedianFinder struct {
	low  *Heap
	high *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		low:  MaxHeap(),
		high: MinHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	ll := this.low.Len()
	lh := this.high.Len()
	if ll == lh {
		if ll == 0 || num < this.high.Top() {
			this.low.Push(num)
		} else {
			this.high.Push(num)
			this.low.Push(this.high.Pop())
		}
	} else {
		if this.low.Top() < num {
			this.high.Push(num)
		} else {
			this.low.Push(num)
			this.high.Push(this.low.Pop())
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.low.Len() == this.high.Len() {
		return float64(this.low.Top())/2 + float64(this.high.Top())/2
	}
	return float64(this.low.Top())
}

func main() {
	// heap := MaxHeap()
	heap := MinHeap()
	heap.Push(7)
	heap.Push(2)
	heap.Push(1)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5)
	heap.Push(6)
	fmt.Println(heap.data)
	heap.Pop()
	fmt.Println(heap.data)
	heap.Push(8)
	fmt.Println(heap.data)

	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
}
