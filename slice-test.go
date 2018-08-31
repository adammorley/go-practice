package main

import "fmt"
import "math/rand"

func main() {
    var s []int = []int{1,2,3,4}
    fmt.Println(s)
    s = append(s, 5)
    fmt.Println(s)
    var v int
    s, v = s[:len(s)-1], s[len(s)-1]
    fmt.Println(s)
    fmt.Println(v)
    s = push(s, 6)
    fmt.Println(s)
}

func push(s []int, v int) []int {
    return append(s, v)
}

func pop(s []int) (v int) {
    s, v = s[:len(s)-1], s[len(s)-1]
    return
}

func randByteSlice() []byte {
	var b []byte = make([]byte, 4)
	rand.Read(b)
	return b
}
