package main

import "fmt"

func fib(n int) func(int) int {
    //f_n = f_n-1 + f_n-2
    var fib []int = make([]int, 2) // pre-allocate here a bigger size if the number of fibs is known in advance
    fib[0] = 1
    fib[1] = 1
    return func(q int) int {
        if q >= len(fib) {
            for i := len(fib); i <= q; i++ {
                fib = append(fib, fib[i-1] + fib[i-2]) // can probably optimize this append call if pre-allocated:
                /*
                    make capacity bigger; use len and append only when len() exceed (but if there is cap)
                */
            }
        }
        return fib[q]
    }
}

func main() {
    c := fib(0)
    for i := 0; i <= 20; i++ {
        fmt.Println(c(i))
    }
}
