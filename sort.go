package main

import "fmt"

/*func testMergeSort() {
    var a []int = []int{2,66,32,45,1,32,442,21,34,5,6,123,54356,0}
    var r = mergeSort(a[0:len(a)/2
    var r = mergeSort(a)
    var r *[]int mergeSort(&a)
}*/

func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	var b []int = a[:len(a)/2]
	a = a[len(a)/2:]
	return mergeSort(a, b)
}

func mergeSort(a, b []int) []int {
	if len(a) == 1 && len(b) == 1 {
		return merge(a, b)
	} else if len(a) == 1 {
        b = mergeSort(b[:len(b)/2], b[len(b)/2:])
        return merge(a, b)
    } else if len(b) == 1 {
        a = mergeSort(a[:len(a)/2], a[len(a)/2:])
        return merge(a, b)
    }
	var c, d []int = a[0 : len(a)/2], a[len(a)/2:]
	var e, f []int = b[0 : len(b)/2], b[len(b)/2:]
	var g []int = mergeSort(c, d)
	var h []int = mergeSort(e, f)
	return merge(g, h)
}

// merge merges two sorted slices, ensuring the result is sorted
// things to improve:
//      remove pass-by-value (eg: use pointers perhaps)
//      re-use result array
func merge(a, b []int) []int {
	var r []int = make([]int, 0)
	var i, j int = 0, 0
	for {
		// if a has a lower or equal value, put it on the result
		if a[i] < b[j] || a[i] == b[j] {
			r = append(r, a[i])
			i++
		} else if b[j] < a[i] { // same but for b
			r = append(r, b[j])
			j++
		}
		if i == len(a) { // a exhausted
			for ; j < len(b); j++ {
				r = append(r, b[j])
			}
		} else if j == len(b) { // b exhausted
			for ; i < len(a); i++ {
				r = append(r, a[i])
			}
		}
		if j == len(b) && i == len(a) { // global termination condition
			return r
		}
	}
}

func testMerge() {
	var a1 []int = []int{3}
	var b1 []int = []int{4}
	var c1 []int = []int{3, 4}
	compare(merge(a1, b1), c1)
	var a0 []int = []int{3, 4, 12, 44, 53}
	var b0 []int = []int{2, 4, 33, 38, 55}
	var c0 []int = []int{2, 3, 4, 4, 12, 33, 38, 44, 53, 55}
	compare(merge(a0, b0), c0)
}

func testMergeSort() {
	var a []int = []int{33, 1283, 2, 387, 190, 37, 44, 230}
	var b []int = []int{33, 1283, 2, 387, 190, 37, 44, 230, 55}
    compare([]int{2,33,37,44,190,230,387,1283}, MergeSort(a))
	fmt.Println(MergeSort(b))
    fmt.Println(MergeSort([]int{33, 190, 55}))
}

func compare(a, b []int) {
	if len(a) != len(b) {
		panic("wrong length")
	}
	for i := range b {
		if b[i] != a[i] {
			panic("wrong value")
		}
	}
}

func main() {
	fmt.Println("hello")
	//testMergeSort()
	testMerge()
	testMergeSort()
}
