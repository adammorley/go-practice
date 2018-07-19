package main

import (
        "fmt"
        "math/rand"
        "sync"
        "sync/atomic"
        "time"
)

/*
    benchmark using a mutex vs. a channel

    start 1..n go routines which:
        atomic increment the counter
        need mechanism to wait for goroutines
        to finish (channel or waitgroup)

    vs.

    start 1..n go routines which:
        put an increment on a channel
        have a second go routine read from channel
        and do increment

*/

func main() {
    var x uint64 = 0
    var mutex = &sync.Mutex{}

    for i := 1; i < 4; i++ {
