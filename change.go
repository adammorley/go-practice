package main

import "fmt"

var max int = 100000000
//var coins []int = []int{1, 5, 10, 25}
var coins []int = []int{1, 5, 10, 20,25}
var m map[int]int = make(map[int]int)

func change(v int) int {
    if v == 0 {
        m[0] = 0
        return m[0]
    } else if c, ok := m[v]; ok {
		return c
	} else {
        change(v-1)
    }
	var count int = max
	for i := 0; i <= v; i++ {
		for c := range coins {
			coin := coins[c]
			if coin <= v {
				if i+coin == v {
					if m[i]+1 < count {
						count = m[i] + 1
						m[v] = count
					}
				}
			}
		}
	}
	return m[v]
}
func main() {
	if 1 != change(1) {
		fmt.Println("wrong change for 1")
	} else if 2 != change(2) {
		fmt.Println("wrong change for 2")
	} else if 1 != change(5) {
		fmt.Println("wrong change for 5")
	} else if 3 != change(3) {
		fmt.Println("wrong change for 3")
	} else if 1 != change(10) {
		fmt.Println("wrong change for 10")
	} else if 2 != change(11) {
		fmt.Println("wrong change for 11", change(11))
	} else if 1 != change(25) {
		fmt.Println("wrong change for 25")
	} else if 4 != change(42) {
		fmt.Println("wrong change for 42", change(42))
	}
	fmt.Println("bye")
}
