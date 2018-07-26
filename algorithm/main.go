package main

import (
    "strings"
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-26 11:04
 **/

func main() {
    //1.2.3a 和 1.2.4b
    fmt.Println(comp("1.2.3a", "1.2.3b"))
}

func comp(v1 string, v2 string) int {
    strs1 := strings.Split(v1, ".")
    strs2 := strings.Split(v2, ".")
    for i := 0; i < len(strs1) || i < len(strs2); i++ {
        if i == len(strs1) {
            return -1
        }
        if i == len(strs2) {
            return 1
        }
        v := compStr(strs1[i], strs2[i])
        if v != 0 {
            return v
        }
    }
    return 0
}

func compStr(v1, v2 string) int {
    num1 := 0
    str1 := v1
    for i := 0; i < len(v1); i++ {
        if v1[i] < '0' || v1[i] > '9' {
            str1 = v1[i:]
            break
        }
        num1 *= 10
        num1 += int(v1[i] - '0')
    }

    num2 := 0
    str2 := v2
    for i := 0; i < len(v2); i++ {
        if v2[i] < '0' || v2[i] > '9' {
            str2 = v2[i:]
            break
        }
        num2 *= 10
        num2 += int(v2[i] - '0')
    }
    if num1 > num2 {
        return 1
    }
    if num1 < num2 {
        return -1
    }
    for i := 0; i < len(str1) || i < len(str2); i++ {
        if i == len(str1) {
            return -1
        }
        if i == len(str2) {
            return 1
        }
        if str1[i] > str2[i] {
            return 1
        }
        if str1[i] < str2[i] {
            return -1
        }
    }
    return 0

}
