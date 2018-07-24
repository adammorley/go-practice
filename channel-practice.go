/*
while this looks conceptually correct, it doesn't work as you'd expect.  why?  because slices are mutable, and the array underlying the slice can be modified while it's being "sent" to the peer goroutine (remember: a slice is just a pointer to an array with some metadata)

the challenge is that the failure is not deterministic --- the failure mode is racey, and depends upon when the byte slice is modified by the sending go routine.  so the first time, it may work fine.  but then the second time, it might not!

one option is to allocate a new byte array/slice for each i in the for loop.  this has the downside of requiring a memory allocation, and hence a cleanup operation.  another is to use an immutable data type; this is a re-allocation, but it also prevents a boo-boo with the slice since it cannot be modified (eg: it's fundamental to the type).

another option is to have the two go-routines note when they have finished with a give piece of data, either using a mutex or a second channel (the second channel seems more go-y)

however, 
*/
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
