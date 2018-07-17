package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-17 19:08
 **/

func main() {
    fmt.Println(uniquePaths(7, 3))
}
func uniquePaths(m int, n int) int {
    if m == 1 || n == 1 {
        return 1
    }
    m -= 1
    n -= 1
    sum := m + n
    if n > m {
        n = m
    }
    res := 1
    for i := 0; i < n; i++ {
        res *= sum - i
    }
    for i := 1; i <= n; i++ {
        res /= i
    }
    return res
}
