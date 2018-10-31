package main

import (
	"bufio"
    "errors"
	"fmt"
	"math"
	"os"
)

type locations map[int]bool
type distance map[string]locations

func parseFile(f *os.File) (d distance) {
    d = make(distance)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	c := 1
	for scanner.Scan() {
        word :=scanner.Text()
        if _, ok := d[word]; !ok {
            d[word]=make(map[int]bool)
            }
        d[word][c]=true
		c++
	}
	return
}

func (d distance) calc(w0, w1 string) (int, error) {
	i, ok := d[w0]
	if !ok {
		return 0, errors.New(fmt.Sprintf("cannot find %v\n", w0))
	}
	j, ok := d[w1]
	if !ok {
		return 0, errors.New(fmt.Sprintf("cannot find %v\n", w1))
	}
	r := int(math.MaxInt64)
	found := false
	for w0d, _ := range i {
		for w1d, _ := range j {
			if abs(w0d-w1d) < r {
				r = abs(w0d - w1d)
				found = true
			}
		}
	}
	if !found {
		return 0, errors.New("did not find match")
	}
	return r, nil
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
func main() {
    f, e:=os.Open("file")
    if e != nil {
        panic(e)
        }
    d:=parseFile(f)
    r, e:=d.calc("is", "cats")
    fmt.Println(e)
    r, e=d.calc("is", "this")
    fmt.Println(r)
}
