package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-17 21:49
 **/

func main() {
    //nums := []int{1, 2, 3}
    //nums := []int{4,3,2,1}
    nums := []int{9,9}
    fmt.Println(plusOne(nums))
}

func plusOne(digits []int) []int {
    if len(digits) == 0 {
        return digits
    }
    carry := 1
    for i := len(digits) - 1; i > -1; i-- {
        carry = carry + digits[i]
        digits[i] = carry % 10
        carry = carry / 10
    }
    if carry != 0 {
        items := make([]int, 0, len(digits)+1)
        items = append(items, carry)
        items = append(items, digits...)
        digits = items
    }
    return digits
}
