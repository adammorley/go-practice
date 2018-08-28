package main

import (
	"errors"
	"fmt"
)

func one() {
	f := countMemo()
	for i := 1; i <= 9; i++ {
		if f(i) != countRec(i) {
			panic("this shit is busted")
		}
	}
}

func countMemo() func(int) int {
	var c map[int]int = make(map[int]int)
	c[0] = 0
	c[1] = 1
	c[2] = 2
	c[3] = 4
	return func(n int) int {
		if v, ok := c[n]; ok {
			return v
		} else {
			for i := len(c); i <= n; i++ {
				c[i] = c[i-1] + c[i-2] + c[i-3]
			}
			return c[n]
		}
	}
}

func countRec(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	} else {
		return countRec(n-1) + countRec(n-2) + countRec(n-3)
	}
}

type point struct{ r, c int }
type path map[point]bool
type grid [][]bool

func two() {
	var size int = 4
	var p0 *path = new(path)
	*p0 = make(path)
	var p1 *path = new(path)
	*p1 = make(path)
	var p2 *path = new(path)
	*p2 = make(path)
	var g0 [][]bool = make([][]bool, size)
	var g1 [][]bool = make([][]bool, size)
	var g2 [][]bool = make([][]bool, size)
	for i := 0; i < size; i++ {
		g0[i] = make([]bool, size)
		g1[i] = make([]bool, size)
		g2[i] = make([]bool, size)
	}
	g0[1][3] = true
	g0[3][1] = true
	g2[0][3] = true
	g2[1][2] = true
	g2[2][1] = true
	g2[3][0] = true
	fmt.Println(findRoute(g0, 0, 0, p0), p0)
	fmt.Println(findRoute(g1, 0, 0, p1), p1)
	fmt.Println(findRoute(g2, 0, 0, p2), p2)
}

func findRoute(g grid, r, c int, p *path) bool {
	if c > 3 || r > 3 || g[r][c] {
		return false
	}
	if r == 3 && c == 3 || findRoute(g, r+1, c, p) || findRoute(g, r, c+1, p) {
		(*p)[point{r, c}] = true
		return true
	}
	return false
}

func three() {
	var a0, a1, a2 []int = []int{1, 5, 8, 9}, []int{-12, -3, 44, 27}, []int{-4, 0, 2, 8, 9}
	v, e := findMagic(a0)
	fmt.Println(v, e)
	v, e = findMagic(a1)
	fmt.Println(v, e)
	v, e = findMagic(a2)
	fmt.Println(v, e)
}

func findMagic(A []int) (int, error) {
    if len(A) == 0 {
        return 0, errors.New("did not find it")
    }

    i := midpoint(A)
    if A[i] == i {
		return A[i], nil
	} else if A[i] > i {
		A = A[:i]
		return findMagic(A)
	} else if A[i] < i {
		A = A[i+1:]
		return findMagic(A)
	} else {
		panic("this should not happen")
	}

	/*for i := 0; i < len(A); i++ {
		if i == A[i] {
			return i, nil
		}
	}
	return 0, errors.New("didn't find it") //XXX update to have typed error*/
}

func midpoint(A []int) int {
	return len(A) / 2
}

func main() {
	one() // success
	two() // fail w/o help (two-dimensional)
	three() // success
}
