package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-14 16:53
 **/

func main() {
    //fmt.Println(longestValidParentheses(")()())"))
    fmt.Println(longestValidParentheses("(()"))
}
func longestValidParentheses(s string) int {
    array := make([]int8, len(s))
    for i, v := range s {
        if v == '(' {
            array[i] = -1
        } else {
            array[i] = 1
        }
    }
    formatArray(array)
    var max int
    var start = -1
    for i, v := range array {
        if v != 0 {
            start = i
            continue
        }
        if max < i-start {
            max = i - start
        }

    }
    return max
}

func formatArray(array []int8) {
    left := -1
    for i := 0; i < len(array); i++ {
        if array[i] == -1 {
            left = i
            continue
        } else if left == -1 {
            continue
        }
        array[i], array[left] = 0, 0
        left = getLfet(array, left)
    }
}

func getLfet(array []int8, left int) int {
    if left == -1 {
        return left
    }
    for i := left; i > -1; i-- {
        if array[i] == -1 {
            return i
        } else if array[i] == 1 {
            return -1
        }
    }
    return -1
}
