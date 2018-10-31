package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // arbitrary value
	priority int
	index    int // item index
}

// Priority queue implements heap's interface and holds items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}
func (pq *PriorityQueue) Push(x interface{}) {
	i := x.(*Item)
	i.index = len(*pq)
	*pq = append(*pq, i)
}
func (pq *PriorityQueue) Pop() interface{} {
	i := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return i
}
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	items := map[string]int{"banana": 3, "apple": 2, "pear": 4}
	pq := make(PriorityQueue, len(items))
	i := 0
	for v, p := range items {
		pq[i] = &Item{
			value:    v,
			priority: p,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
