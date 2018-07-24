/*
using a conditional to signal the hasher when it is safe to use the data on the channel
*/
package main

import (
        "crypto/sha1"
        "fmt"
        "sync"
)

var m sync.Mutex
var cond *sync.Cond = sync.NewCond(&m)

func hasher(i <-chan []byte, o chan<- [sha1.Size]byte) {
    for b := range i {
        d := sha1.Sum(b)
        o <- d
        cond.Signal()
    }
    close(o)
}

func filler(c chan<- []byte) {
    b := make([]byte, 4, 4)
    cond.L.Lock()
    for i := 0; i < 100; i++ {
        for j := 0; j < 4; j++ {
            b[j] = byte(i)
        }
        c <- b
        cond.Wait()
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
    c := make(chan []byte, 1)
    b := make(chan [sha1.Size]byte)
    go filler(c)
    go hasher(c, b)

    d := make(chan bool)
    go printer(b, d)

    <-d
}
