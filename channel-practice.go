package main

import (
        "crypto/sha1"
        "fmt"
        "time"
)

func hasher(i <-chan []byte, o chan<- [sha1.Size]byte) {
    for b := range i {
        o <- sha1.Sum(b)
        time.Sleep(100 * time.Millisecond)
    }
    close(o)
}

func filler(c chan<- []byte) {
    b := make([]byte, 4, 4)
    for i := 0; i < 10; i++ {
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
    go hasher(c, b)
    go filler(c)

    d := make(chan bool)
    go printer(b, d)

    <-d
}
