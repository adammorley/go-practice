package main

import "fmt"

func first() {
    // first make a mask to reset the bits from position i to j
    var i uint = 2
    var j uint = 6

    var m uint = 19 // 10011
    m = m << i
    var n uint = 1044 // 10000010100

    var ib uint = ^uint(0) >> (64-i)
    var jb uint = ^uint(0) << j
    var k = jb ^ ib
    var r uint = n & k
    r = r ^ m
    fmt.Printf("%b\n", r)
}

func third() {
    var n uint = 1775
    fmt.Println(flipBits(n))

    /*need to test flip any zero bit
        need to find zero bits*/

}

func countMostBitsOneInSeries(n uint) (x uint) {
    var c, i uint = 0, 0
    for ; i <= 64; i++ {
        if 1 == ((n >> i) & uint(1)) {
            c++
        } else {
            c = 0
        }
        if c > x {
            x = c
        }
    }
    return
}

func flipBits(n uint) (x uint) {
    var i uint = 0
    for ; i <= 64; i++ {
        var r uint = n >> i & uint(1)
        if 0 == (r | uint(0)) {
            z := countMostBitsOneInSeries((1 << i) | n)
            if z > x {
                x = z
            }
        }
    }
    return
}


func fourth() {
    var n uint = 7
    x := countBits(n)
    var i uint = 1
    var y, z uint
    for ; i <= 2000000; i++ {
        if y == 0 && countBits( (n-i) ) == x {
            y = n-i
        }
        if z == 0 && countBits( (n+i) ) == x {
            z = n+i
        }
        if y != 0 && z != 0 {
            break
        }
    }
    fmt.Println(n, y, z)
    fmt.Printf("%b %b %b\n", n, y, z)
}

func countBits(n uint) (r uint) {
    var i uint = 0
    for ; i <= 64; i++ {
        if 1 == (n >> i) & uint(1) {
            r++
        }
    }
    return
}

func sixth() {
    var m, n uint = 29, 15
    var x = m ^ n
    fmt.Println(countBits(x))
}

func seventh() {
    var n uint = 555
    fmt.Printf("%b\n", n)
    n = swap(n)
    fmt.Printf("%b\n", n)
}

func swap(n uint) (x uint) {
    b := n >> 64
    x = n << 1
    x += b
    return
}

func main() {
    first()
    third()
    fourth()
    sixth()
    seventh()
}
