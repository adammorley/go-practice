/*
correcting the error in the channel-practice.go file using re-allocated array for each loop
this generates a 4-byte structure every time the loop runs
*/
package main

import (
        "crypto/sha1"
        "fmt"
)

func hasher(i <-chan [4]byte, o chan<- [sha1.Size]byte) {
    for b := range i {
        o <- sha1.Sum(b[:])
    }
    close(o)
}

func filler(c chan<- [4]byte) {
    for i := 0; i < 100; i++ {
        var b [4]byte
        for j := 0; j < 4; j++ {
            b[j] = byte(i)
        }
        c <- b
    }
    close(c)
}

func printer(i <-chan [sha1.Size]byte, d chan<- bool) {
    for x := range i {
        fmt.Printf("%x\n", x)
    }
    d <- true
}

func main() {
    c := make(chan [4]byte)
    b := make(chan [sha1.Size]byte)
    go hasher(c, b)
    go filler(c)

    d := make(chan bool)
    go printer(b, d)

    <-d
}
