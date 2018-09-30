package main

import (
    "bytes"
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type pair struct {
	s0, s1 string // the delta strings for each buffer
}

func lcs(s0, s1 string) (*list.List, string) {
	// build a matrix to score the counts
	var m [][]int = make([][]int, len(s0))
	for i := 0; i < len(s0); i++ {
		m[i] = make([]int, len(s1))
	}
	var b0, b1 []byte = []byte(s0), []byte(s1)
	for i := 0; i < len(b0); i++ {
		for j := 0; j < len(b1); j++ {
			if b0[i] == b1[j] {
				if i == 0 || j == 0 {
					m[i][j] = 1
				} else {
					m[i][j] = m[i-1][j-1] + 1
				}
			}
		}
	}
	var l *list.List = list.New()
	tb := make([]byte, len(m))
	bb0, bb1 := make([]byte, 0), make([]byte, 0)
    var same bool = true
	for i, j := 0, 0; i < len(m); i, j = i+1, j+1 {
		if m[i][j] == 0 { // the strings don't match, so start accum and replace with * in tb
            same = false
			tb[i] = byte('*')
			bb0 = append(bb0, b0[i])
			bb1 = append(bb1, b1[j])
		} else if m[i][j] > 0 {
            if !same {
			    l.PushBack(&pair{string(bb0), string(bb1)})
			    bb0, bb1 = make([]byte, 0), make([]byte, 0)
                same = true
            }
			tb[i] = b0[i]
		}
	}
	return l, string(tb)
}

func open(filename string) (f *os.File) {
	f, e := os.Open(filename)
	if e != nil {
		panic(e)
	}
	return
}

func main() {
	f0, f1 := open("file0"), open("file1")
	s0, s1 := bufio.NewScanner(f0), bufio.NewScanner(f1)
	// iterate through the scanners, comparing line-by-line
	var lineCount int
	for s0.Scan() && s1.Scan() {
		lineCount++
		t0, t1 := s0.Text(), s1.Text()
		if t0 == t1 {
			// the two lines are equal, so print one
			printLineAndSeparate(t0, lineCount, false)
		} else {
			l, s := lcs(t1, t0)
			line := buildLine(l, s)
			printLineAndSeparate(line, lineCount, true)
		}
	}
}

func buildLine(l *list.List, s string) string {
	text := "lines differ; chars replaced with *;"
	b := make([]byte, 0, len(text)+len(s)*2) // XXX: fix this
    bb := bytes.NewBuffer(b)
    bb.WriteString(text)
    bb.WriteString(s)
    bb.WriteString(";")
	e := l.Front()
	for e != nil {
        fmt.Println(e.Value.(*pair).s0)
        /*bb.WriteString(e.Value.(*pair).s0)
        bb.WriteString(";")
        bb.WriteString(e.Value.(*pair).s1)
        bb.WriteString(";")*/
		e = e.Next()
	}
    return bb.String()
}

func printLineAndSeparate(line0 string, lineCount int, diff bool) {
	if diff {
		fmt.Printf("line: %v differs\n", lineCount)
        fmt.Println(line0)
	} else {
		fmt.Printf("line: %v does not differ\n", lineCount)
		fmt.Println(line0)
	}
	fmt.Println("---------------")
}
