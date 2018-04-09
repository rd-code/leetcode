package main

import (
    "fmt"
)

func main() {
    fmt.Println(isMatch("aa", "a*"))
}

/**
Implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true
 */
func isMatch(s string, p string) bool {
    if len(p) == 0 {
        return len(s) == 0
    }
    if len(p) == 1 {
        return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
    }
    if p[1] != '*' {
        if len(s) == 0 {
            return false
        }
        return (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])

    }
    for len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
        if isMatch(s, p[2:]) {
            return true
        }
        s = s[1:]
    }
    return isMatch(s,p[2:])
}
