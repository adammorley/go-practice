/*
use something akin to 2PC to notify when the two goroutines should move forward (to prevent the overwrite problem described in channel-practice.go)

avoiding the memory allocation for each for loop in filler leads to increased complexity; it's fairly wonky too (the sync channel has to be buffered to not cause a deadlock).
*/
package main

import (
        "crypto/sha1"
        "fmt"
)

func hasher(i <-chan []byte, o chan<- [sha1.Size]byte, s chan<- bool) {
    for b := range i {
        o <- sha1.Sum(b)
        s <- true
    }
    close(o)
}

func filler(c chan<- []byte, s <-chan bool) {
    b := make([]byte, 4, 4)
    for i := 0; i < 100; i++ {
        if i > 0 { // allow filler to start
            <-s
        }
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
    c := make(chan []byte)
    b := make(chan [sha1.Size]byte)
    s := make(chan bool, 1)
    go hasher(c, b, s)
    go filler(c, s)

    d := make(chan bool)
    go printer(b, d)

    <-d
}
