package main

import "fmt"
import "math/rand"

func main() {
    var x []byte = make([]byte, 4)/*
    x = randByteSlice()
    fmt.Println(x)*/
    if x == nil {
        fmt.Println("slice is nil")
    } else {
        fmt.Println("slice is not nil")
    }
}

func randByteSlice() []byte {
    var b []byte = make([]byte, 4)
    rand.Read(b)
    return b
}
