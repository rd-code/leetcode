package main

import (
    "sort"
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-10 15:18
 **/
func combinationSum(candidates []int, target int) [][]int {
    if target < 1 {
        return nil
    }
    var res [][]int
    for _, v := range candidates {
        if v > target {
            continue
        } else if v == target {
            res = append(res, []int{v})
            continue
        }
        items := combinationSum(candidates, target-v)
        for _, item := range items {
            tmp := []int{v}
            tmp = append(tmp, item...)
            sort.Ints(tmp)
            res = append(res, tmp)
        }
    }
    var items [][]int
    for _, t := range res {
        for _, item := range items {
            if len(item) != len(t) {
                continue
            }
            equal := true
            for i := range item {
                if item[i] != t[i] {
                    equal = false
                    break
                }
            }
            if equal {
                goto Label
            }
        }

        items = append(items, t)
    Label:
    }
    return items
}

func main() {
    candidates := []int{2, 3, 6, 7}
    fmt.Println(combinationSum(candidates, 7))

}

//candidates = [2,3,6,7], target = 7,
