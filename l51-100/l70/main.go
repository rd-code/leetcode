package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-18 17:20
 **/

func main() {
    fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
    if n < 4 {
        return n
    }
    n1, n2 := 3, 2
    for i := n1; i < n-1; i++ {
        n1, n2 = n1+n2, n1
    }
    return n1 + n2

}
