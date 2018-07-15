package main

import (
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-15 11:07
 **/

func main() {
    // 0,9
    s := "barfoothefoobarman"
    words := []string{"foo", "bar"}

    //s := "wordgoodstudentgoodword"
    //words := []string{"word", "student"}

    // s := "barfoofoobarthefoobarman"
    // words := []string{"bar", "foo", "the"}

    //s := "wordgoodgoodgoodbestword"
    //words := []string{"word", "good", "best", "good"}

    v := findSubstring(s, words)
    fmt.Println(v)
}

func findSubstring(s string, words []string) []int {
    if len(s) == 0 || len(words) == 0 {
        return nil
    }
    data := make(map[string]int, len(words))
    for _, word := range words {
        data[word] += 1
    }
    var res []int
    size := len(words[0]) * len(words)
    worldSize := len(words[0])
    for i := 0; i < len(s)-size+1; i++ {
        if check(s[i:], copy(data), size, worldSize) {
            res = append(res, i)
        }
    }
    return res
}

func copy(data map[string]int) map[string]int {
    res := make(map[string]int)
    for k, v := range data {
        res[k] = v
    }
    return res

}

func check(s string, data map[string]int, size, worldSize int) bool {
    if len(s) < size {
        return false
    }
    for i := 0; i < size; i += worldSize {
        str := s[i : i+worldSize]
        if data[str] != 0 {
            data[str] -= 1
        } else {
            return false
        }
    }
    for _, v := range data {
        if v > 0 {
            return false
        }
    }
    return true

}
