package main

import "container/list"
import "fmt"

const boardSize = 8 // make configurable later, if other board sizes

func solve() *list.List {
	var cs [boardSize]int
	var solutions *list.List = list.New()
	placeQueen(0, cs, solutions)
	return solutions
}

func placeQueen(r int, cs [boardSize]int, solutions *list.List) {
	if r == boardSize {
		solutions.PushBack(cs)
	} else {
		for c := 0; c < boardSize; c++ {
			if validQueen(r, c, cs) {
				cs[r] = c
				placeQueen(r+1, cs, solutions)
			}
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func validQueen(r, c int, cs [boardSize]int) bool {
	for i := 0; i < r; i++ {
		if cs[i] == c || abs(r-i) == abs(c-cs[i]) {
			return false
		}
	}
	return true
}

func main() {
	solutions := solve()
	e := solutions.Front()
	for e != nil {
		fmt.Println(e.Value.([boardSize]int))
		e = e.Next()
	}
}
