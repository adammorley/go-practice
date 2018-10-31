package main

import (
	"container/list"
	"fmt"
)

func main() {
	sol := list.New()
	var bs [7]int
	placeQueen(0, bs, sol)
	fmt.Println(sol.Len())
}

func placeQueen(r int, bs [7]int, sol *list.List) {
	if r == 7 {
		sol.PushBack(bs)
		return
	}
	for c := 0; c < 7; c++ {
		if validQueen(r, c, bs) {
			bs[r] = c
			placeQueen(r+1, bs, sol)
		}
	}
}
func validQueen(r, c int, bs [7]int) bool {
	for i := 0; i < r; i++ {
		if c == bs[i] {
			return false
		} else if abs(r-i) == abs(c-bs[i]) {
			return false
		}
	}
	return true
}
func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
