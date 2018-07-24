package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-24 16:37
 **/

func main() {
    nums := []int{1, 2, 3}
    items := subsets(nums)
    for _, item := range items {
        fmt.Println(item)
    }
}

func subsets(nums []int) [][]int {
    res := subset(nums)
    res = append(res, []int{})
    return res
}

func subset(nums []int) [][]int {
    if len(nums) == 0 {
        return nil
    }
    res := make([][]int, 0, len(nums)*(len(nums)-1)/2)
    for i, v := range nums {
        res = append(res, []int{v})
        items := subset(nums[i+1:])
        for _, item := range items {
            tmp := make([]int, 0, len(item)+1)
            tmp = append(tmp, v)
            tmp = append(tmp, item...)
            res = append(res, tmp)
        }
    }
    return res
}
