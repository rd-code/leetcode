package main

import (
    "sort"
    "fmt"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-10 15:55
 **/

func combinationSum2(candidates []int, target int) [][]int {
    data := make(map[int]int, len(candidates))
    for _, v := range candidates {
        data[v] += 1
    }
    return combinationSum(data, target)
}

func combinationSum(candidates map[int]int, target int) [][]int {
    if target < 1 {
        return nil
    }
    var res [][]int

    for k, v := range candidates {
        if k > target || v < 1 {
            continue
        } else if k == target {
            res = append(res, []int{k})
            continue
        }
        candidates[k] -= 1
        items := combinationSum(candidates, target-k)
        for _, item := range items {
            tmp := []int{k}
            tmp = append(tmp, item...)
            sort.Ints(tmp)
            res = append(res, tmp)
        }
        candidates[k] += 1
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
    candidates := []int{2,5,2,1,2}
    fmt.Println(combinationSum2(candidates, 5))
}
