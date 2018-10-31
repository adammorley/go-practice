package main

import "container/list"
import "fmt"

func lcs(s0, s1 string) *list.List {
	// build a matrix to score the counts
	var m [][]int= make([][]int, len(s0))
	for i := 0; i < len(s0); i++ {
		m[i] = make([]int, len(s1))
	}
	var b0, b1 []byte = []byte(s0), []byte(s1)
	var l *list.List = list.New()
	var max int
	for i := 0; i < len(b0); i++ {
		for j := 0; j < len(b1); j++ {
			if b0[i] == b1[j] {
				if i == 0 || j == 0 {
					m[i][j] = 1
				} else {
					m[i][j] = m[i-1][j-1] + 1
				}
				if m[i][j] >= max {
					if m[i][j] > max {
						l = list.New()
					}
					l.PushBack(buildSubstring(b0, max, i))
					max = m[i][j]
				}
			}
		}
	}
	return l
}

func buildSubstring(b []byte, max, i int) string {
	var o []byte = make([]byte, max)
	for k, n := i-max+1, 0; k <= i; k, n = k+1, n+1 {
		o[n] = b[k]
	}
	return string(o)
}

func main() {
	var s0, s1 string = "cats and dogs", "dogs and cats"
	l := lcs(s0, s1)
	e := l.Front()
	for e != nil {
		fmt.Println(e.Value.(string))
		e = e.Next()
	}
}
