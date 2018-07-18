package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-18 16:12
 **/

func main() {
   // fmt.Println(addBinary("1010", "1011"))
    fmt.Println(addBinary("11", "1"))
    //"1010", b = "1011"
}

func addBinary(a string, b string) string {
    if len(a) == 0 {
        return b
    }
    if len(b) == 0 {
        return a
    }
    n := len(a)
    if n < len(b) {
        n = len(b)
    }
    n1, n2 := len(a), len(b)
    var carry, v1, v2 uint8
    res := make([]byte, n+1)
    for i := 0; i < n; i++ {
        if i < n1 {
            v1 = a[n1-1-i] - '0'
        } else {
            v1 = 0
        }
        if i < n2 {
            v2 = b[n2-1-i] - '0'
        } else {
            v2 = 0
        }
        carry = carry + v1 + v2
        res[n-i] = carry%2 + '0'
        carry = carry / 2
    }
    if carry != 0 {
        res[0] = carry + '0'
    } else {
        res = res[1:]
    }
    return string(res)
}
