package main

import (
    "strconv"
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-11 19:37
 **/

func main() {
    fmt.Println(countAndSay(5))
}
func countAndSay(n int) string {
    if n < 2 {
        return "1"
    }
    res := "1"
    for i := 1; i < n; i++ {
        tmp := ""
        last := rune(res[0])
        count := 0
        for _, v := range []rune(res) {
            if last == v {
                count++
            } else {
                tmp += strconv.Itoa(count)
                tmp += string(last)
                last = v
                count = 1
            }
        }
        tmp += strconv.Itoa(count)
        tmp += string(last)
        res = tmp
    }
    return res
}
