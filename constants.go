package main

import "fmt"

type ByteSize float64

func main() {
	const (
		B ByteSize = 1 << (10 * iota)
		KB
		MB
		GB
		TB
		PB
		EB
	)
	fmt.Println("size", B)
	fmt.Println("size", KB)
}
