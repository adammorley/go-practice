package main

import (
        "fmt"
        "math/rand"
        "sync/atomic"
        "time"
)

func main() {
    var ops uint64

    rand.Seed(42)

    for i := 0; i < 50; i++ {
        go func() {
            for {
                atomic.AddUint64(&ops, 1)
                x := rand.Intn(100)
                d := time.Duration(x)
                time.Sleep(d * time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}
