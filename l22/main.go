/**
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */
package main

import "fmt"

func main() {
    fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
    left, right := make([]rune, n), make([]rune, n)
    for i := 0; i < n; i++ {
        left[i] = '('
        right[i] = ')'
    }
    return generate("", left, right)
}

func generate(data string, left, right []rune) []string {
    if len(left) == 0 {
        data += string(right)
        return []string{data}
    }
    if len(left) == len(right) {
        data += string(left[0])
        return generate(data, left[1:], right)
    }
    items := generate(data+string(left[0]), left[1:], right)
    items = append(items, generate(data+string(right[0]), left, right[1:])...)
    return items
}
