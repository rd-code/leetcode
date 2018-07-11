package main

import (
    "math"
    "fmt"
)

func main() {
    t := 10
    t = t << 1
    fmt.Println(t)
    fmt.Println(divide(-2147483648, 2))
}

func divide(dividend int, divisor int) int {
    if divisor == 0 {
        return math.MaxInt32
    }
    if dividend > math.MaxInt32 || dividend < math.MinInt32 {
        return math.MaxInt32
    }

    flag := false
    if dividend < 0 {
        dividend = 0 - dividend
        flag = !flag
    }
    if divisor < 0 {
        divisor = 0 - divisor
        flag = !flag
    }
    var result int

    for dividend-divisor >= 0 {
        count, tmp := 1, divisor
        for {
            tmp = tmp << 1
            if dividend-tmp >= 0 {
                count = count << 1
            } else {
                tmp = tmp >> 1
                break
            }
            if tmp == 0 {
                break
            }
        }
        result += count
        dividend = dividend - tmp
    }
    if flag {
        result = 0 - result
    }
    if result > math.MaxInt32 || result < math.MinInt32 {
        return math.MaxInt32
    }
    return result
}
