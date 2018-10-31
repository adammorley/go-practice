package main

import "fmt"

func mergeSort(n []int) []int {
	if len(n) == 0 {
		return []int{}
	} else if len(n) == 1 {
		return n
	} else {
		return merge(mergeSort(n[:len(n)/2]), mergeSort(n[len(n)/2:]))
	}
}
func merge(n0, n1 []int) []int {
	r := make([]int, len(n0)+len(n1))
	rI := 0
    i,j:=0,0
	for i < len(n0) && j < len(n1) {
		if n0[i] < n1[j] {
			r[rI] = n0[i]
			i++
		} else if n1[j] <= n0[i] {
			r[rI] = n1[j]
			j++
		}
		rI++
	}
	for j < len(n1) {
		r[rI] = n1[j]
		rI++
		j++
	}
	for i < len(n0) {
		r[rI] = n0[i]
		rI++
		i++
	}
	return r
}

func main() {
	one := []int{5, 12, 7, 9}
	two := []int{3, 88, 2, 4}
	fmt.Println(mergeSort(one))
	fmt.Println(mergeSort(two))
}
