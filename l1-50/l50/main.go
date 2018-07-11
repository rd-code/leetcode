package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-11 19:08
 **/

func main() {
    fmt.Println(myPow(0.0001, -2))
}

func myPow(x float64, n int) float64 {
    if n == 0 || x == 1 {
        return 1
    }
    if x == 0 {
        return 0
    }
    if x == -1 {
        if n%2 == 0 {
            return 1
        } else {
            return -1
        }
    }
    var isNegative bool
    if n < 0 {
        isNegative, n = true, 0-n
    }
    var res float64 = 1
    for i := 0; i < n; i++ {
        if res == 0 {
            break
        }
        res *= x
        if isNegative {
            if (1 / res) == 0 {
                return 0
            }
        }
    }
    if isNegative {
        res = 1 / res
    }
    return res
}
