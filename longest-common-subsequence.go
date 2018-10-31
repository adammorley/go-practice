package main

import "fmt"

func lcs(s0, s1 string) string {
	var m [][]int = make([][]int, len(s0)+1)
	for i := 0; i < len(s0)+1; i++ {
		m[i] = make([]int, len(s1)+1)
	}
	for s0i, mi := 0, 1; s0i < len(s0); s0i, mi = s0i+1, mi+1 {
		for s1j, mj := 0, 1; s1j < len(s1); s1j, mj = s1j+1, mj+1 {
			if s0[s0i] == s1[s1j] {
				m[mi][mj] = m[mi-1][mj-1] + 1
			} else {
				if m[mi-1][mj] > m[mi][mj-1] {
					m[mi][mj] = m[mi-1][mj]
				} else {
					m[mi][mj] = m[mi][mj-1]
				}
			}
		}
	}
	i := len(m) - 1
	j := len(m[i]) - 1
    var r []rune = make([]rune,m[i][j])
    var c int = len(r)-1
	for i != 0 && j != 0 {
		fmt.Println(i, j)
		if m[i][j-1] == m[i][j] {
			j = j - 1
		} else if m[i-1][j] == m[i][j] {
			i = i - 1
		} else if m[i][j] > m[i][j-1] {
			j = j - 1
            r[c] = rune(s0[i-1])
            c--
		} else {
			panic("foobar")
		}
	}
	return string(r)
}

func main() {
	var s0, s1 string = "DCBCAD", "ACDBC"
	r := lcs(s0, s1)
	fmt.Println(r)
}
