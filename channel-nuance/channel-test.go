package main

import "fmt"
import "time"

func receiver(i <-chan int) {
    for n := range i {
        fmt.Println("receiver: ", n)
    }
}

func sender(o chan<- int) {
    for i := 0; i < 10; i++ {
        o <- i
        fmt.Println("sender: ", i)
    }
}

func main() {
    c := make(chan int)
    go sender(c)
    go receiver(c)
    time.Sleep(time.Second)
}
