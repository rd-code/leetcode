package main

import (
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-11 17:50
 **/

func main() {
    a := "123"
    b := "456"
    fmt.Println(multiply(a, b))

}

const (
    zero = '0'
)

func multiply(num1 string, num2 string) string {
    if len(num1) == 0 || len(num2) == 0 || num1 == "0" || num2 == "0" {
        return "0"
    }

    n1, n2 := make([]byte, 0, len(num1)), make([]byte, 0, len(num2))
    for _, num := range num1 {
        n1 = append(n1, byte(num-zero))
    }
    for i := 0; i < len(n1)-i-1; i++ {
        p := len(n1) - i - 1
        n1[i], n1[p] = n1[p], n1[i]
    }

    for _, num := range num2 {
        n2 = append(n2, byte(num-zero))
    }

    for i := 0; i < len(n2)-i-1; i++ {
        p := len(n2) - i - 1
        n2[i], n2[p] = n2[p], n2[i]
    }

    var array []byte
    for i, num := range n1 {
        var items []byte
        var carry byte
        for _, v := range n2 {
            v = num*v + carry
            if v >= 10 {
                carry = v / 10
                v = v % 10
            } else {
                carry = 0
            }
            items = append(items, v)
        }
        if carry != 0 {
            items = append(items, carry)
            carry = 0
        }

        for j := 1; i < i; j++ {
            if len(array) == j {
                array = append(array, 0)
            }
        }
        for j := 0; j < len(items); j++ {
            p := i + j
            if len(array) < p+1 {
                array = append(array, 0)
            }
            v := array[p] + items[j] + carry
            if v >= 10 {
                carry = v / 10
                v = v % 10
            } else {
                carry = 0
            }
            array[p] = v
        }
        if carry != 0 {
            p := i + len(items)
            if len(array) < p+1 {
                array = append(array, 0)
            }
            array[p] = carry
        }
    }

    var res string
    for _, v := range array {
        res = string(rune(v+zero)) + res
    }

    return res
}
