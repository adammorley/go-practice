package main

import "container/heap"
import "fmt"

type heapY []int

func (h *heapY) p(i int) int {
	return (i - 1) / 2
}
func (h *heapY) l(i int) int {
	return 2*i + 1
}
func (h *heapY) r(i int) int {
	return 2*i + 2
}
func (h *heapY) i(i int) {
	*h = append(*h, i)
	n := len(*h) - 1
	p := h.p(n)
	for p >= 0 && (*h)[p] > (*h)[n] {
		h.s(p, n)
		n = p
		p = h.p(n)
	}
}
func (h *heapY) s(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

type MinIntHeap []int
func (h MinIntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinIntHeap) Len() int {
	return len(h)
}
func (h *MinIntHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}
func (h *MinIntHeap) Pop() (r interface{}) {
	r = (*h)[len(*h)-1]
	*h = (*h)[0:len(*h)-1]
	return
}

func main() {
	var a []int = []int{6, 5, 3, 1, 8, 7, 2, 4}
	var h *heapY = new(heapY)
	*h = make(heapY, 0)
	for i := 0; i < len(a); i++ {
		h.i(a[i])
	}
	fmt.Println(h)

    var he *MinIntHeap = &MinIntHeap{6,5,4,2,8,7,2}
    heap.Init(he)
    fmt.Println(he)
    for he.Len() > 0 {
        fmt.Println(heap.Pop(he))
    }
}
