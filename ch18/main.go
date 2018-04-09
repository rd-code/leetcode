/**
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note: The solution set must not contain duplicate quadruplets.

For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

 */
package main

import (
    "sort"
    "fmt"
)

func main() {
    fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {
        return nil
    }
    data := map[int]int{}
    for _, num := range nums {
        if _, ok := data[num]; !ok {
            data[num] = 0
        }
        data[num] += 1
    }

    tmp := map[[4]int]struct{}{}

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                t := target - nums[i] - nums[j] - nums[k]
                if v, ok := data[t]; ok {
                    if nums[i] == t {
                        v--
                    }
                    if nums[j] == t {
                        v--
                    }
                    if nums[k] == t {
                        v--
                    }
                    if v > 0 {
                        items := []int{nums[i], nums[j], nums[k], t}
                        sort.Ints(items)
                        tmp[[4]int{items[0], items[1], items[2], items[3]}] = struct{}{}
                    }

                }
            }
        }

    }
    result := make([][]int, 0, len(tmp))
    for k := range tmp {
        result = append(result, []int{k[0], k[1], k[2], k[3]})
    }
    return result
}
