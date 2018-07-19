package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-19 11:35
 **/

func main() {
    //fmt.Println(mySqrt(4))
    fmt.Println(mySqrt(11000))
}
func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    if x < 4 {
        return 1
    }
    var i int
    for i = 2; i*i <= x; i *= 2 {
    }
    for ; ; i -= 1 {
        if i*i <= x {
            return i
        }
    }
    return i - 1
}
