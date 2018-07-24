/*
    calculate fibonacci numbers using memoization
*/
package fibonacci

// f_n = f_{n-1} + f_{n-2}
func fib(n int) (int, func(int) int) {
    var f = make(map[int]int)
    f[0] = 0
    f[1] = 1
    for i := 2; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[n], func(x int) int {
        if v, ok := f[x]; ok {
            return v
        } else {
            for i := len(f)-1; i <= x; i++ {
                f[i] = f[i-1] + f[i-2]
            }
            return f[x]
        }
    }
}
