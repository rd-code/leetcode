package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-17 21:15
 **/

func main() {
    fmt.Println(isNumber("2e10"))
}

func isNumber(s string) bool {
    if len(s) == 0 {
        return false
    }
    array := []byte(s)
    var i int
    for i = 0; i < len(array); i++ {
        if array[i] != ' ' {
            break
        }
    }
    array = array[i:]
    for i = len(array); i > 0; i-- {
        if array[i-1] != ' ' {
            break
        }
    }
    if i > 0 {
        array = array[:i]
    }
    for i = 0; i < len(array); i++ {
        if array[i] == 'e' || array[i] == 'E' {
            break
        }
    }
    if i == len(array) {
        return check(rm(array), 0)
    }

    return check(rm(array[:i]), 0) && check(rm(array[i+1:]), 1)
}

func rm(array []byte) []byte {
    if len(array) == 0 {
        return array
    }
    if array[0] == '+' || array[0] == '-' {
        return array[1:]
    }
    return array
}

func check(array []byte, count int) bool {
    if len(array) == 0 {
        return false
    }
    for _, v := range array {
        if v == '.' {
            if count > 0 {
                return false
            }
            count += 1
        } else if v < '0' || v > '9' {
            return false
        }
    }
    if len(array) == 1 && array[0] == '.' {
        return false
    }
    return true
}
