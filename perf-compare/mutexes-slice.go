package main

import (
        "fmt"
        "math/rand"
        "sync/atomic"
        "time"
)

func main() {
    var s = make([]int64, 5, 5)

    var readOps uint64
    var writeOps uint64

    for r := 0; r < 100; r++ {
        go func() {
            var total int64 = 0
            for {
                pos := rand.Intn(5)
                total += atomic.LoadInt64(&s[pos])
                atomic.AddUint64(&readOps, 1)
                //time.Sleep(time.Millisecond)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                pos := rand.Intn(5)
                val := rand.Int63()
                atomic.StoreInt64(&s[pos], val)
                atomic.AddUint64(&writeOps, 1)
                //time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    fmt.Println(s)

}
