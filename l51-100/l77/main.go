package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-24 16:17
 **/
func main() {
    items := combine(4, 2)
    for _, item := range items {
        fmt.Println(item)
    }
}

func combine(n int, k int) [][]int {
    return generate(n, 1, k)
}

func generate(n int, start int, k int) [][]int {
    if k <= 1 {
        res := make([][]int, 0)
        for i := start; i <= n; i++ {
            res = append(res, []int{i})
        }
        return res
    }
    size := n - start
    res := make([][]int, 0, size*(size-1)/2)
    for i := start; i <= n; i++ {
        tmp := generate(n, i+1, k-1)
        items := make([][]int, 0, len(tmp))
        for _, t := range tmp {
            item := make([]int, 0, len(t)+1)
            item = append(item, i)
            item = append(item, t...)
            items = append(items, item)
        }
        res = append(res, items...)
    }
    return res
}
