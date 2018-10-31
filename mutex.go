package main

import "fmt"
import "sync"
import "time"

type hits struct {
	sync.Mutex
	n int
}

func main() {
	var hi *hits = new(hits)
	go func(h *hits) {
		for {
			h.Lock()
			h.n++
			h.Unlock()
		}
	}(hi)
	go func(h *hits) {
		for {
			h.Lock()
			h.n++
			h.Unlock()
		}
	}(hi)
	go func(h *hits) {
		for {
			time.Sleep(time.Second)
			h.Lock()
			fmt.Println(h.n)
			h.Unlock()
		}
	}(hi)
	time.Sleep(10 * time.Second)
}
