package main

import (
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-20 11:30
 **/

func main() {
    //3
    //word1 := "horse"
    //word2 := "ros"
    //word1 := "intention"
    //word2 := "execution"

    // word1 := "pneumonoultramicroscopicsilicovolcanoconiosis"
    //  word2 := "pneumonoconiosis"
    // word1 := "trinitrophenylmethylnitramine"
    //word2 := "dinitrophenylhydrazine"
    // word1 := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdef"
    //..word2 := "bcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefg"
    //  fmt.Println(minDistance(word1, word2))
    fmt.Println(minDistance("", ""))
    fmt.Println(minDistance("a", "ab"))

}

func minDistance(word1, word2 string) int {
    array1, array2 := []byte(word1), []byte(word2)
    data := make([][]int, len(array1))
    for i := range data {
        items := make([]int, len(array2))
        for j := range items {
            items[j] = -1
        }
        data[i] = items
    }
    return calculate(array1, array2, data, len(word1), len(word2))

}

func calculate(word1, word2 []byte, data [][]int, row, column int) int {
    fmt.Println(row, column)
    var min, tmp = -1, 0
    if row == 0 {
        return column
    }
    if column == 0 {
        return row
    }
    if string(word1) == string(word2) {
        min = 0
        goto Label
    }

    if data[row-1][column-1] != -1 {
        return data[row-1][column-1]
    }
    if word1[row-1] == word2[column-1] {
        min = calculate(word1, word2, data, row-1, column-1)
        goto Label
    }
    tmp = calculate(word1, word2, data, row-1, column) + 1
    if min > tmp || min == -1 {
        min = tmp
    }
    tmp = calculate(word1, word2, data, row, column-1) + 1
    if min > tmp || min == -1 {
        min = tmp
    }
    tmp = calculate(word1, word2, data, row-1, column-1) + 1
    if min > tmp || min == -1 {
        min = tmp
    }
Label:
    data[row-1][column-1] = min
    return min
}

/*func minDistance(word1 string, word2 string) int {
    num := 1
    for i := 1; i < len(word1); i++ {
        num *= i
    }
    data := make(map[string]map[string]int, num)
    return calculate(word1, word2, data)
}

func calculate(s1, s2 string, data map[string]map[string]int) int {
    if _, ok := data[s1]; ok {
        if _, ok := data[s1][s2]; ok {
            return data[s1][s2]
        }
    }
    var min = -1
    if s1 == s2 {
        min = 0
        goto Label
    }
    if len(s1) == 0 {
        min = len(s2)
        goto Label
    }
    if len(s2) == 0 {
        min = len(s1)
        goto Label
    }

    for i := 0; i < len(s1); i++ {
        for j := 0; j < len(s2); j++ {
            var num int
            num += calculate(s1[:i], s2[:j], data)
            num += calculate(s1[i+1:], s2[j+1:], data)
            if s1[i] != s2[j] {
                num += 1
            }
            if min > num || min < 0 {
                min = num
            }
        }
    }
Label:
    if _, ok := data[s1]; !ok {
        data[s1] = map[string]int{}
    }
    data[s1][s2] = min
    if _, ok := data[s2]; !ok {
        data[s2] = map[string]int{}
    }
    data[s2][s1] = min
    return min
}
*/
