package fibonacci

import "testing"

func TestFib(t *testing.T) {
    x, f := fib(4)
    if x != 3 {
        t.Error("wrong number for fib 4")
    } else if f(6) != 8 {
        t.Error("wrong number for 6")
    } else if f(12) != 144 {
        t.Error("wrong number for 12")
    }
}
