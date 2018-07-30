/*
    n light switches, all off.  if off, turn on, stop, if on, turn off, goto next switch.  do k times

    thoughts:
        - this is a counting problem.  it turns out it's the n bits of the number k when drawn out:

        0  _ _
        1  | _
        2  _ |
        3  | |
        4  _ _

    so need two mask bits on top of k
*/
package main

// light switches read right to left
func switchFlips(n uint, k uint) uint {
    var mask uint
    var i uint = 0
    for ; i < n; i++ {
        mask = 1 << i | mask
    }
    return mask & k
}

func main() {
    if switchFlips(2, 0) != 0 {
        panic("problem")
    } else if switchFlips(2, 1) != 1 {
        panic("problem")
    } else if switchFlips(2, 2) != 2 {
        panic("problem")
    } else if switchFlips(2, 3) != 3 {
        panic("problem")
    } else if switchFlips(2, 4) != 0 {
        panic("problem")
    } else if switchFlips(2, 5) != 1 {
        panic("problem")
    } else if switchFlips(4, 7) != 7 {
        panic("problem")
    } else if switchFlips(2, 7) != 3 {
        panic("problem")
    }
}
