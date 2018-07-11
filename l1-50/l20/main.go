package main

import "fmt"

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.


 */

func main() {
    fmt.Println(isValid("()"))
}
func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    if len(s)%2 != 0 {
        return false
    }
    left := map[rune]struct{}{
        '(': {},
        '{': {},
        '[': {},
    }

    right := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    items := make([]rune, 0, len(s))
    for _, r := range s {
        if _, ok := left[r]; ok {
            items = append(items, r)
            fmt.Println(string(r))
        } else {
            var v rune
            if v, ok = right[r]; !ok {
                return false
            }
            if len(items) < 1 || items[len(items)-1] != v {
                return false
            }
            items = items[:len(items)-1]
        }
    }

    return len(items) == 0
}
