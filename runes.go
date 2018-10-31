package main

import "fmt"

func main() {
    var a string = "abc"
    var b string = "bde"
    var arunes []rune = []rune(a)
    var brunes []rune = []rune(b)
    var aLen int = len(arunes)
    var bLen int = len(brunes)
    var lengths [][]int = make([][]int, aLen+1)
    for i := 0; i <= aLen; i++ {
        lengths[i] = make([]int, bLen+1)
    }
}
