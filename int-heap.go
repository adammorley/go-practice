package main

import (
	"container/heap"
	"fmt"
)

// an int heap of min-heap ints
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// uses a pointer receiver because modifying slice's length
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	var r int = (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1] // this is weird; operator precedence is funky
	return r
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 4)
	for h.Len() > 0 {
		fmt.Println(h)
		fmt.Println(heap.Pop(h))
	}
}
