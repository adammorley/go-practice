package main

import "container/list"
import "fmt"

type String string

func (s0 String) lcs(s1 String) *list.List {
	// we'll work with runes
	var r0, r1 []byte = []byte(s0), []byte(s1)
	// output, need to be able to append, so can't stringify yet
	var o *list.List = new(list.List)
	o = list.New()
	// m is the matrix to track which values "align" between s0 and s1
	var m [][]int = make([][]int, len(r0))
	for i := 0; i < len(r0); i++ {
		m[i] = make([]int, len(r1))
	}
	var g int = 0
	for i := 0; i < len(r0); i++ {
		for j := 0; j < len(r1); j++ {
			if r0[i] == r1[j] {
				if i == 0 || j == 0 {
					m[i][j] = 1
				} else {
					m[i][j] = m[i-1][j-1] + 1
				}
				if m[i][j] >= g {
					if m[i][j] > g {
						g = m[i][j]
						o = list.New()
					}
					o.PushBack(buildSubstring(r0, g, i))
				}
			}
		}
	}
	return o
}

func buildSubstring(r0 []byte, z, i int) (t []byte) {
	t = make([]byte, z)
	for k, n := i-z+1, 0; k <= i; k, n = k+1, n+1 {
		t[n] = r0[k]
	}
	return
}

func main() {
	var s0 String = "this is a nice day to run down the lane"
	var s1 String = "nice day this run down lane"
	e := s0.lcs(s1).Front()
    for e != nil {
	    fmt.Println(string(e.Value.([]byte)))
        e = e.Next()
    }
}
