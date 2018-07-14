package main

import (
    "fmt"
    "sort"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-14 14:22
 **/

func main() {
    nums := []int{2, 3, 1}
    nextPermutation(nums)
    fmt.Println(nums)
}
func nextPermutation(nums []int) {
    if len(nums) < 2 {
        return
    }
    if generaeGreater(nums) {
        return
    }
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}

func generaeGreater(nums []int) bool {
    if len(nums) == 2 {
        if nums[0] < nums[1] {
            nums[0], nums[1] = nums[1], nums[0]
            return true
        } else {
            return false
        }
    }
    if generaeGreater(nums[1:]) {
        return true
    }
    min := 1
    for i := 1; i < len(nums); i++ {
        if nums[min] > nums[i] && nums[i] > nums[0] {
            min = i
        }
    }
    if nums[0] < nums[min] {
        nums[0], nums[min] = nums[min], nums[0]
        sort.Ints(nums[1:])
        return true
    }

    return false
}
